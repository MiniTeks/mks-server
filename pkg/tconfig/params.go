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

package tconfig

import (
	"net/http"

	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	versionedResource "github.com/tektoncd/pipeline/pkg/client/resource/clientset/versioned"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

// Client struct contains Tekton, Kubernetes and Dynamic intefaces.
type Client struct {
	Tekton     tektonclient.Interface
	Kube       kubernetes.Interface
	Resource   versionedResource.Interface
	HTTPclient http.Client
	Dynamic    dynamic.Interface
}

type Params interface {
	SetKubeConfigPath(string)

	SetKubeContext(string)

	Client() (*Client, error)

	SetNamespace(string)
}
