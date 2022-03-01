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
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MksPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MksPipelineSpec `json:"spec"`
}

type MksPipelineSpec struct {
	Workspaces []MksPipelineWorkspace `json:"workspaces,omitempty"`
	Param      []ParamSpec            `json:"params,omitempty"`
	Task       []MksPipelineTask      `json:"tasks"`
}

type MksPipelineWorkspace struct {
	Name      string `json:"name"`
	Workspace string `json:"workspace,omitempty"`
}

type MksPipelineTask struct {
	Name            string                 `json:"name"`
	PipelineTaskRef MksPipelineTaskRef     `json:"taskRef"`
	Workspaces      []MksPipelineWorkspace `json:"workspaces,omitempty"`
	Param           []Param                `json:"params,omitempty"`
	RunAfter        []string               `json:"runAfter,omitempty"`
}

type MksPipelineTaskRef struct {
	Name string `json:"name"` // name of the task to be referenced
	Kind string `json:"kind,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MksPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MksPipeline `json:"items"`
}
