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

// ArrayOrString is a type that can hold a single string or string array.
// Used in JSON unmarshalling so that a single JSON field can accept
// either an individual string or an array of strings.
type ArrayOrString struct {
	Type      string   `json:"type"` // Represents the stored type of ArrayOrString.
	StringVal string   `json:"stringVal"`
	ArrayVal  []string `json:"arrayVal"`
}

// ParamSpec defines an arbitrary named  input whose value can be supplied by a
// `Param`.
type ParamSpec struct {
	// Name declares the name by which a parameter is referenced.
	Name string `json:"name"`

	// Type is the user-specified type of the parameter. The possible types
	// are currently "string" and "array", and "string" is the default.
	// +optional
	Type string `json:"type,omitempty"`
	// Description is a user-facing description of the parameter that may be
	// used to populate a UI.
	// +optional
	Description string `json:"description,omitempty"`
	// Default is the value a parameter takes if no input value via a Param is supplied.
	// +optional
	Default string `json:"default,omitempty"`
}

// Param defines a string value to be used for a ParamSpec with the same name.
type Param struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
