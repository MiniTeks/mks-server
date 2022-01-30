package mkstaskrun

// import (
// 	"context"
// 	"fmt"

// 	// "github.com/MiniTeks/mks-server/pkg/actions"
// 	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller"
// 	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
// 	// tbeta "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1" //v1beta1
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	// "k8s.io/apimachinery/pkg/labels"
// 	"k8s.io/client-go/kubernetes"
// )

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

// // func ConvertToTekton(mtr *v1alpha1.MksTaskRun) *talpha.TaskRun {
// // 	res := &talpha.TaskRun{}
// // 	res.Kind = "TaskRun"
// // 	res.APIVersion = "tekton.dev/v1alpha1"
// // 	res.ObjectMeta = mtr.ObjectMeta
// // 	res.Spec.TaskRef.Name = mtr.Spec.TaskRef.Name
// // 	return res

// // }
