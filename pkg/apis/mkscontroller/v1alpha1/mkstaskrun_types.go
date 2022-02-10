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
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
/*
MksTaskRun struct corresponds to the following CustomResourceDefinition:

	apiVersion: apiextensions.k8s.io/v1
	kind: CustomResourceDefinition
	metadata:
	name: mkstaskruns.mkscontroller.example.mks
	spec:
	group: mkscontroller.example.mks
	versions:
		- name: v1alpha1
		served: true
		storage: true
		schema:
			# schema used for validation
			openAPIV3Schema:
			type: object
			properties:
				spec:
				type: object
				properties:
					taskRef:
					type: object
					properties:
						name:
						type: string

	names:
		kind: MksTaskRun
		plural: mkstaskruns
	scope: Namespaced
*/
type MksTaskRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MksTaskRunSpec `json:"spec"`
}

// MksTaskRunSpec struct defines the Spec field of MksTaskRun struct.
type MksTaskRunSpec struct {
	TaskRef MksTaskRef `json:"taskRef"`
}

// MksTaskRef struct defines the TaskRef field of the MkstaskRunSpec.
// - Name 	name of the task to be referenced
type MksTaskRef struct {
	Name string `json:"name"` // name of the task to be referenced
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MksTaskRunList struct defines a list for the MksTaskRun type.
type MksTaskRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MksTaskRun `json:"items"`
}
