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
MksTask struct corresponds to the following CustomResourceDefinition:

	apiVersion: apiextensions.k8s.io/v1
	kind: CustomResourceDefinition
	metadata:
	name: mkstasks.mkscontroller.example.mks
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
					name:
					type: string
					image:
					type: string
					command:
					type: string
					args:
					type: string

	names:
		kind: MksTask
		plural: mkstasks
	scope: Namespaced
*/
type MksTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MksTaskSpec `json:"spec"`
}

// MksTaskSpec struct defines the Spec field of the MksTask struct.
// - Name 		name of the task.
// - Image 		docker image for the task.
// - Command 	command to be run.
// - Args 		arguments to the command.
type MksTaskSpec struct {
	Steps  []MksTaskSteps `json:"steps"`
	Params []MksParamSpec `json:"params,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MksTaskList struct defines a list for the MksTask type.
type MksTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MksTask `json:"items"`
}

type MksTaskSteps struct {
	Name       string `json:"name"`    // name of the task
	Image      string `json:"image"`   // docker image for the task
	Command    string `json:"command"` // command to be run
	WorkingDir string `json:"workingDir,omitempty"`
	Args       string `json:"args"` // arguments to the command
}
