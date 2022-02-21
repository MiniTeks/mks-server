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
	Task MksPipelineTask `json:"tasks"`
}

type MksPipelineTask struct {
	Name            string             `json:"name"`
	PipelineTaskRef MksPipelineTaskRef `json:"taskRef"`
}

type MksPipelineTaskRef struct {
	Name string `json:"name"` // name of the task to be referenced
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MksPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MksPipeline `json:"items"`
}
