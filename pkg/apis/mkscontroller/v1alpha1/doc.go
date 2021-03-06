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

// +k8s:deepcopy-gen=package
// +groupName=mkscontroller.example.mks

/*
Package v1alpha1 containes the MKS CustomResourceDefinition and its structs.

Currently there are three custom mks resource:-

	- MksTask
	- MksTaskRun
	- MksPipelineRun

Basics

A CustomResourceDefinition defined in a YAML file can be registered by providing
its corresponding structs. For example let us take the following YAML:

	apiVersion: apiextensions.k8s.io/v1
	kind: CustomResourceDefinition
	metadata:
	# name must match the spec fields below, and be in the form: <plural>.<group>
	name: resource.mks.example.com
	spec:
	# group name to use for REST API: /apis/<group>/<version>
	group: mks.example.com
	# list of versions supported by this CustomResourceDefinition
	versions:
		- name: v1
		# Each version can be enabled/disabled by Served flag.
		served: true
		# One and only one version must be marked as the storage version.
		storage: true
		schema:
			openAPIV3Schema:
			type: object
			properties:
				spec:
				type: object
				properties:
					name:
					type: string
	# either Namespaced or Cluster
	scope: Namespaced
	names:
		# plural name to be used in the URL: /apis/<group>/<version>/<plural>
		plural: resources
		# singular name to be used as an alias on the CLI and for display
		singular: resource
		# kind is normally the CamelCased singular type. Your resource manifests use this.
		kind: Resource
		# shortNames allow shorter string to match your resource on the CLI
		shortNames:
		- rs


The above CRD can be translated into following structs:

	import (
		metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	)

	type Resource struct {
		metav1.TypeMeta 	`json:",inline"`
		metav1.ObjectMeta 	`json:"metadata,omitempty"`

		Spec ResourceSpec 	`json:"spec"`
	}

	type ResourceSpec struct {
		Name string `json:"name"`
	}

	type ResourceList struct {
		metav1.TypeMeta 	`json:",inline"`
		metav1.ListMeta 	`json:"metadata"`

		Items []Resource 	`json:"items"`
	}

Working

Notice how the 'schema' field is translated into different struct based on its
type. For each 'object' you should have a dedicated struct, for other types you
should have fields inside the struct.

Code generation

Later Deepcopygen is used to generate cleintset and deepcopy definitions for
these structs. They can be denoted in the file by adding following comment
above the structs.

	// +genclient
	// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

See https://github.com/kubernetes/code-generator for more information on code
generation.

Register

Every defined type must be registered. See register.go in this package for more
information on this.
*/
package v1alpha1
