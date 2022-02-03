package mkstask

import (
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

var TrGroupResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "task"}

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

func Create(cl *tconfig.Client, mt *v1alpha1.MksTask, opt metav1.CreateOptions, ns string) (*v1beta1.Task, error) {
	tktr := ConvertToTekton(mt)

	object, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(tktr)
	usttr := &unstructured.Unstructured{
		Object: object,
	}
	nusttr, err := actions.Create(TrGroupResource, cl, usttr, ns, opt)
	if err != nil {
		return nil, err
	}
	var task *v1beta1.Task
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(nusttr.UnstructuredContent(), &task); err != nil {
		return nil, err
	}
	return task, nil

}