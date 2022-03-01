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

package mkspipelinerun

import (
	"github.com/MiniTeks/mks-server/pkg/actions"
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/tconfig"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	prbeta "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Group and Version to uniquely identify the Tekton API.
var prGroupResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "pipelineruns"}

// ConvertToTekton converts a mksresource into corresponding Tekton resource
// definition using field to field copy from mks object to tekton object.
func ConvertToTekton(mpr *v1alpha1.MksPipelineRun) *prbeta.PipelineRun {
	res := &prbeta.PipelineRun{}
	res.Kind = "PipelineRun"
	res.APIVersion = "tekton.dev/v1beta1"
	res.ObjectMeta = metav1.ObjectMeta{
		Name:      mpr.ObjectMeta.Name,
		Namespace: mpr.ObjectMeta.Namespace,
	}
	var tparam []v1beta1.Param
	for _, prm := range mpr.Spec.Params {
		tparam = append(tparam, v1beta1.Param{
			Name:  prm.Name,
			Value: *v1beta1.NewArrayOrString(prm.Value),
		})
	}
	var twork []v1beta1.WorkspaceBinding
	for _, wrk := range mpr.Spec.Workspaces {
		twork = append(twork, v1beta1.WorkspaceBinding{
			Name:                  wrk.Name,
			PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{ClaimName: wrk.PersistentVolumeClaim.ClaimName},
			ConfigMap:             &v1.ConfigMapVolumeSource{LocalObjectReference: v1.LocalObjectReference{Name: wrk.ConfigMap.Name}},
		})
	}
	res.Spec = prbeta.PipelineRunSpec{
		PipelineRef: &prbeta.PipelineRef{Name: mpr.Spec.PipelineRef.Name},
		Params:      tparam,
		Workspaces:  twork,
	}
	return res
}

// Create takes a mksresource object, converts it to tekton resource object
// using ConvertToTekton function and then calls the Create fucntion defined in
// the actions package to create resource on Kubernetes/OpenShift cluster using
// Tekton API.
func Create(cl *tconfig.Client, mpr *v1alpha1.MksPipelineRun,
	opt metav1.CreateOptions, ns string) (*prbeta.PipelineRun, error) {

	tkpr := ConvertToTekton(mpr)

	object, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(tkpr)
	ustpr := &unstructured.Unstructured{
		Object: object,
	}
	nustpr, err := actions.Create(prGroupResource, cl, ustpr, ns, opt)
	if err != nil {
		return nil, err
	}
	var pipelinerun *prbeta.PipelineRun
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(nustpr.UnstructuredContent(), &pipelinerun); err != nil {
		return nil, err
	}
	return pipelinerun, nil

}

func Delete(cl *tconfig.Client, opt metav1.DeleteOptions, mprName string, ns string) error {
	err := actions.Delete(prGroupResource, cl, mprName, ns, opt)
	if err != nil {
		return err
	}
	return nil
}
