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
	"sync"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/restmapper"
)

var (
	doOnce      sync.Once
	apiGroupRes []*restmapper.APIGroupResources
)

// GetGroupVersionResource returns a Group version resource given a schema
func GetGroupVersionResource(gr schema.GroupVersionResource,
	discovery discovery.DiscoveryInterface) (*schema.GroupVersionResource, error) {
	var err error
	doOnce.Do(func() {
		err = InitializeAPIGroupRes(discovery)
	})
	if err != nil {
		return nil, err
	}

	rm := restmapper.NewDiscoveryRESTMapper(apiGroupRes)
	gvr, err := rm.ResourceFor(gr)
	if err != nil {
		return nil, err
	}

	return &gvr, nil
}

// InitializeAPIGroupRes checks if the provided discovery client is available to
// provide discovery information.
func InitializeAPIGroupRes(discovery discovery.DiscoveryInterface) error {
	var err error
	apiGroupRes, err = restmapper.GetAPIGroupResources(discovery)
	if err != nil {
		return err
	}
	return nil
}
