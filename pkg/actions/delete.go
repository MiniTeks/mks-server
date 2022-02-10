// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Satyam Bhardwaj <sabhardw@redhat.com>
// SPDX-FileCopyrightText: 2022 Utkarsh Chaurasia <uchauras@redhat.com>
// SPDX-FileCopyrightText: 2022 Avinal Kumar <avinkuma@redhat.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package actions

import (
	"context"

	"github.com/MiniTeks/mks-server/pkg/tconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Delete deletes the given Tekton resource on the Kubernetes/Openshift cluster
// via Tekton API and returns the result as error.
func Delete(gr schema.GroupVersionResource, clients *tconfig.Client,
	objname, ns string, opt metav1.DeleteOptions) error {
	gvr, err := GetGroupVersionResource(gr, clients.Tekton.Discovery())
	if err != nil {
		return err
	}

	err = clients.Dynamic.Resource(*gvr).Namespace(ns).Delete(context.Background(), objname, opt)
	if err != nil {
		return err
	}
	return nil
}
