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
	"context"
	"testing"

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	fake "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned/fake"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConvertToTekton(t *testing.T) {
	mtr := &v1alpha1.MksTaskRun{
		Spec: v1alpha1.MksTaskRunSpec{
			TaskRef: v1alpha1.MksTaskRef{
				Name: "test-mkstask",
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-mkstaskrun",
			Namespace: "testspace",
		},
	}
	converted := ConvertToTekton(mtr)

	expt := &v1beta1.TaskRun{
		Spec: v1beta1.TaskRunSpec{
			TaskRef: &v1beta1.TaskRef{
				Name: "test-mkstask",
			},
			PodTemplate: &v1beta1.PodTemplate{},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-mkstaskrun",
			Namespace: "testspace",
		},
	}
	expt.Kind = "TaskRun"
	expt.APIVersion = "tekton.dev/v1beta1"

	if !compare(expt, converted) {
		t.Errorf("Cannot convert to tekton successfully")
	}
}

func compare(tr1 *v1beta1.TaskRun, tr2 *v1beta1.TaskRun) bool {
	return tr1.Kind == tr2.Kind && tr1.Spec.TaskRef.Name == tr2.Spec.TaskRef.Name && tr1.ObjectMeta.Name == tr2.ObjectMeta.Name && tr1.ObjectMeta.Namespace == tr2.ObjectMeta.Namespace
}

func TestCreate(t *testing.T) {
	obj := &v1alpha1.MksTaskRun{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "createtest2",
			Namespace: "namespace",
		},
		Spec: v1alpha1.MksTaskRunSpec{
			TaskRef: v1alpha1.MksTaskRef{
				Name: "test-create",
			},
		},
	}
	mkstaskrunClient := fake.NewSimpleClientset()
	crt, err := mkstaskrunClient.MkscontrollerV1alpha1().MksTaskRuns("namespace").Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		t.Errorf("Cannot create mkstaskrun: %v", err)
	}
	t.Logf("Successfully created mkstaskrun: %s", crt.ObjectMeta.Name)
}
