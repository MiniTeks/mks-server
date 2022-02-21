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

package mkspipeline

import (
	"github.com/MiniTeks/mks-server/pkg/actions"
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/tconfig"
	pbeta "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Group and Version to uniquely identify the Tekton API.
var prGroupResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "pipelines"}

// ConvertToTekton converts a mksresource into corresponding Tekton resource
// definition using field to field copy from mks object to tekton object.
func ConvertToTekton(mp *v1alpha1.MksPipeline) *pbeta.Pipeline {
	res := &pbeta.Pipeline{}
	res.Kind = "Pipeline"
	res.APIVersion = "tekton.dev/v1beta1"
	res.ObjectMeta = metav1.ObjectMeta{
		Name:      mp.ObjectMeta.Name,
		Namespace: mp.ObjectMeta.Namespace,
	}
	// res.Spec = pbeta.PipelineSpec{

	// 	PipelineTaskRef: &pbeta.MksPipelineTaskRef{Name: mp.Spec.PipelineTaskRef.Name},
	// }
	res.Spec = pbeta.PipelineSpec{
		Tasks: []pbeta.PipelineTask{
			{
				Name:    mp.Spec.Task.Name,
				TaskRef: &pbeta.TaskRef{Name: mp.Spec.Task.PipelineTaskRef.Name},
			},
		},
	}
	return res
}

// Create takes a mksresource object, converts it to tekton resource object
// using ConvertToTekton function and then calls the Create fucntion defined in
// the actions package to create resource on Kubernetes/OpenShift cluster using
// Tekton API.
func Create(cl *tconfig.Client, mp *v1alpha1.MksPipeline,
	opt metav1.CreateOptions, ns string) (*pbeta.Pipeline, error) {

	tkp := ConvertToTekton(mp)

	object, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(tkp)
	ustp := &unstructured.Unstructured{
		Object: object,
	}
	nustpr, err := actions.Create(prGroupResource, cl, ustp, ns, opt)
	if err != nil {
		return nil, err
	}
	var pipeline *pbeta.Pipeline
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(nustpr.UnstructuredContent(), &pipeline); err != nil {
		return nil, err
	}
	return pipeline, nil

}

func Delete(cl *tconfig.Client, opt metav1.DeleteOptions, mpName string, ns string) error {
	err := actions.Delete(prGroupResource, cl, mpName, ns, opt)
	if err != nil {
		return err
	}
	return nil
}
