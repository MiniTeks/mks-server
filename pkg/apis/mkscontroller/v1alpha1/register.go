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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	mkscontroller "github.com/MiniTeks/mks-server/pkg/apis/mkscontroller"
)

// Group and Version to uniquely identify the API.
var SchemeGroupVersion = schema.GroupVersion{Group: mkscontroller.GroupName, Version: "v1alpha1"}

// Kind return the GroupKind for the current Group-Version pair.
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Reosurce returns the GroupResource for the current Group-Version pair.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// addKnownTypes add all the defined types to the schema.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&MksTaskRun{},
		&MksTaskRunList{},
		&MksPipelineRun{},
		&MksPipelineRunList{},
		&MksTask{},
		&MksTaskList{},
		&MksPipeline{},
		&MksPipelineList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
