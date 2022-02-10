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

package mkstask

import (
	"fmt"
	"os"

	"github.com/MiniTeks/mks-server/pkg/actions"
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/tconfig"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Group and Version to uniquely identify the Tekton API.
var TaskGroupResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "task"}

// ConvertToTekton converts a mksresource into corresponding Tekton resource
// definition using field to field copy from mks object to tekton object.
func ConvertToTekton(mt *v1alpha1.MksTask) *v1beta1.Task {
	res := &v1beta1.Task{}
	res.Kind = "Task"
	res.APIVersion = "tekton.dev/v1beta1"
	res.ObjectMeta = metav1.ObjectMeta{
		Name:      mt.ObjectMeta.Name,
		Namespace: mt.ObjectMeta.Namespace,
	}
	res.Spec = v1beta1.TaskSpec{
		Steps: []v1beta1.Step{
			{
				Container: v1.Container{
					Image:   mt.Spec.Image,
					Name:    mt.Spec.Name,
					Command: []string{mt.Spec.Command},
					Args:    []string{mt.Spec.Args},
				},
			},
		},
	}
	return res
}

// Create takes a mksresource object, converts it to tekton resource object
// using ConvertToTekton function and then calls the Create function defined in
// the actions package to create resource on Kubernetes/OpenShift cluster using
// Tekton API.
func Create(cl *tconfig.Client, mt *v1alpha1.MksTask, opt metav1.CreateOptions, ns string) (*v1beta1.Task, error) {
	tktr := ConvertToTekton(mt)

	object, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(tktr)
	usttr := &unstructured.Unstructured{
		Object: object,
	}
	nusttr, err := actions.Create(TaskGroupResource, cl, usttr, ns, opt)
	if err != nil {
		return nil, err
	}
	var task *v1beta1.Task
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(nusttr.UnstructuredContent(), &task); err != nil {
		return nil, err
	}
	return task, nil

}

// Get returns a Tekton object given the name of the resource.
func Get(c *tconfig.Client, taskname string, opts metav1.GetOptions, ns string) (*v1beta1.Task, error) {
	unstructuredT, err := actions.Get(TaskGroupResource, c, taskname, ns, opts)
	if err != nil {
		return nil, err
	}

	var task *v1beta1.Task
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredT.UnstructuredContent(), &task); err != nil {
		fmt.Fprintf(os.Stderr, "failed to get task from %s namespace \n", ns)
		return nil, err
	}
	return task, nil
}

// List returns an array of Tekton object for the particular resource.
func List(tcl *tconfig.Client, opt metav1.ListOptions, ns string) ([]*v1beta1.Task, error) {
	objlist, err := actions.List(TaskGroupResource, tcl, ns, opt)
	if err != nil {
		return nil, err
	}
	var tasks []*v1beta1.Task
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(objlist.UnstructuredContent(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// Delete deletes the Tekton resource given the name of the resource.
func Delete(tcl *tconfig.Client, mtrname string, opt metav1.DeleteOptions, ns string) error {
	if err := actions.Delete(TaskGroupResource, tcl, mtrname, ns, opt); err != nil {
		return err
	}
	return nil
}

// Update takes a mksresource object, converts it to tekton resource object
// using ConvertToTekton function and then calls the Update function defined in
// the actions package to update resource on Kubernetes/OpenShift cluster using
// Tekton API.
func Update(cl *tconfig.Client, mt *v1alpha1.MksTask, opt metav1.UpdateOptions,
	ns string) (*v1beta1.Task, error) {
	tktr := ConvertToTekton(mt)

	object, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(tktr)
	usttr := &unstructured.Unstructured{
		Object: object,
	}
	nusttr, err := actions.Update(TaskGroupResource, cl, usttr, ns, opt)
	if err != nil {
		return nil, err
	}
	var task *v1beta1.Task
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(nusttr.UnstructuredContent(), &task); err != nil {
		return nil, err
	}
	return task, nil

}
