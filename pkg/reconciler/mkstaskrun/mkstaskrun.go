package mkstaskrun

import (
	"github.com/MiniTeks/mks-server/pkg/actions"
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/tface"

	// 	"context"
	// 	"fmt"

	// 	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller"
	// 	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	tbeta "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"

	//v1beta1
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// 	// "k8s.io/apimachinery/pkg/labels"
// 	"k8s.io/client-go/kubernetes"

var trGroupResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "taskruns"}

// const taskrunlabel = mkscontroller.GroupName + "/mkstaskrun"

// type Reconciler struct {
// 	kubeclient kubernetes.Interface
// }

// var _ mkstaskrunrecoiniler.Interface = (*Reconciler)(nil)

// func (r *Reconciler) ReconcileKind(ctx context.Context, d *v1alpha1.MksTaskRun) reconciler.Event {
// 	// selector := labels.SelectorFromSet(labels.Set{
// 	// 	taskrunlabel: d.Name,
// 	// })

// 	list, err := r.kubeclient.CoreV1().Pods(d.Namespace).Watch(ctx, metav1.ListOptions{})
// 	if err != nil {
// 		return fmt.Errorf("Fail to watch resource")
// 	}
// 	fmt.Println(list)
// 	return nil
// }

func ConvertToTekton(mtr *v1alpha1.MksTaskRun) *tbeta.TaskRun {
	res := &tbeta.TaskRun{}
	res.Kind = "TaskRun"
	res.APIVersion = "tekton.dev/v1beta1"
	res.ObjectMeta = metav1.ObjectMeta{
		Name:      mtr.ObjectMeta.Name,
		Namespace: mtr.ObjectMeta.Namespace,
	}
	res.Spec = tbeta.TaskRunSpec{
		TaskRef:     &tbeta.TaskRef{Name: mtr.Spec.TaskRef.Name},
		PodTemplate: &tbeta.PodTemplate{},
	}
	return res

}

func Create(cl *tface.Client, mtr *v1alpha1.MksTaskRun, opt metav1.CreateOptions, ns string) (*tbeta.TaskRun, error) {
	// trgvr , err := actions.GetGroupVersionResource(trGroupResource, cl.Tekton.Discovery())
	// if err != nil {
	// 	return nil, err
	// }
	tktr := ConvertToTekton(mtr)

	object, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(tktr)
	usttr := &unstructured.Unstructured{
		Object: object,
	}
	nusttr, err := actions.Create(trGroupResource, cl, usttr, ns, opt)
	if err != nil {
		return nil, err
	}
	var taskrun *tbeta.TaskRun
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(nusttr.UnstructuredContent(), &taskrun); err != nil {
		return nil, err
	}
	return taskrun, nil

}
