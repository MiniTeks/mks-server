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

package mkstaskrun

import (
	"github.com/MiniTeks/mks-server/pkg/actions"
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/tconfig"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Group and Version to uniquely identify the Tekton API.
var trGroupResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "taskruns"}

// ConvertToTekton converts a mksresource into corresponding Tekton resource
// definition using field to field copy from mks object to tekton object.
func ConvertToTekton(mtr *v1alpha1.MksTaskRun) *v1beta1.TaskRun {
	res := &v1beta1.TaskRun{
		Spec: v1beta1.TaskRunSpec{
			TaskRef:     &v1beta1.TaskRef{Name: mtr.Spec.TaskRef.Name},
			PodTemplate: &v1beta1.PodTemplate{},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      mtr.ObjectMeta.Name,
			Namespace: mtr.ObjectMeta.Namespace,
		},
	}
	res.Kind = "TaskRun"
	res.APIVersion = "tekton.dev/v1beta1"

	return res
}

// Create takes a mksresource object, converts it to tekton resource object
// using ConvertToTekton function and then calls the Create function defined in
// the actions package to create resource on Kubernetes/OpenShift cluster using
// Tekton API.
func Create(tcl *tconfig.Client, mtr *v1alpha1.MksTaskRun,
	opt metav1.CreateOptions, ns string) (*v1beta1.TaskRun, error) {
	converted := ConvertToTekton(mtr)

	object, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(converted)
	unstructuredObj := &unstructured.Unstructured{
		Object: object,
	}

	newUnstrcturedObj, err := actions.Create(trGroupResource, tcl, unstructuredObj, ns, opt)
	if err != nil {
		return nil, err
	}

	var taskrun *v1beta1.TaskRun

	if err :=
		runtime.DefaultUnstructuredConverter.FromUnstructured(newUnstrcturedObj.UnstructuredContent(), &taskrun); err != nil {
		return nil, err
	}

	return taskrun, nil
}

// Get returns a Tekton object given the name of the resource.
func Get(tcl *tconfig.Client, mtrname string, opt metav1.GetOptions, ns string) (*v1beta1.TaskRun, error) {
	obj, err := actions.Get(trGroupResource, tcl, mtrname, ns, opt)
	if err != nil {
		return nil, err
	}
	var taskrun *v1beta1.TaskRun
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj.UnstructuredContent(), &taskrun); err != nil {
		return nil, err
	}
	return taskrun, nil
}

// List returns an array of Tekton object for the particular resource.
func List(tcl *tconfig.Client, opt metav1.ListOptions, ns string) ([]*v1beta1.TaskRun, error) {
	objlist, err := actions.List(trGroupResource, tcl, ns, opt)
	if err != nil {
		return nil, err
	}
	var taskruns []*v1beta1.TaskRun
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(objlist.UnstructuredContent(), &taskruns); err != nil {
		return nil, err
	}
	return taskruns, nil
}

// Delete deletes the Tekton resource given the name of the resource.
func Delete(tcl *tconfig.Client, mtrname string, opt metav1.DeleteOptions, ns string) error {
	if err := actions.Delete(trGroupResource, tcl, mtrname, ns, opt); err != nil {
		return err
	}
	return nil
}
