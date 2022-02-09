package mkspipelinerun

import (
	"fmt"

	"github.com/MiniTeks/mks-server/pkg/actions"
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/db"
	"github.com/MiniTeks/mks-server/pkg/tconfig"
	prbeta "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var prGroupResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "pipelineruns"}

func ConvertToTekton(mpr *v1alpha1.MksPipelineRun) *prbeta.PipelineRun {
	res := &prbeta.PipelineRun{}
	res.Kind = "PipelineRun"
	res.APIVersion = "tekton.dev/v1beta1"
	res.ObjectMeta = metav1.ObjectMeta{
		Name:      mpr.ObjectMeta.Name,
		Namespace: mpr.ObjectMeta.Namespace,
	}
	res.Spec = prbeta.PipelineRunSpec{
		PipelineRef: &prbeta.PipelineRef{Name: mpr.Spec.PipelineRef.Name},
	}
	return res

}

func Create(cl *tconfig.Client, mpr *v1alpha1.MksPipelineRun, opt metav1.CreateOptions, ns string) (*prbeta.PipelineRun, error) {

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

// Add method of controller for informer function: AddFunc
func (c *controller) handleAdd(obj interface{}) {
	fmt.Println("add was called")
	fmt.Println(obj)
	tp := &tconfig.TektonParam{}
	cs, err := tp.Client()
	if err != nil {
		fmt.Errorf("Cannot get client %v", err)
	}
	pr, err := Create(cs, obj.(*v1alpha1.MksPipelineRun), metav1.CreateOptions{}, obj.(*v1alpha1.MksPipelineRun).Namespace)
	if err != nil {
		db.Increment(rClient, "mksPipelineRunfailed")
		return
	} else {
		db.Increment(rClient, "mksPipelineRuncreated")
		fmt.Println("PipelineRun created")
		fmt.Printf("uid %s", pr.UID)
	}

	c.queue.Add(obj)
}

// Update method of controller for informer function: UpdateFunc
// It gets new object & resource version can be checked to figure out if resource was edited or not
func (c *controller) handleUpdate(old, new interface{}) {
	fmt.Println("Update was called")

}

// Delete method of controller for informer function: DeleteFunc
func (c *controller) handleDelete(obj interface{}) {
	tp := &tconfig.TektonParam{}
	cs, err := tp.Client()
	if err != nil {
		fmt.Errorf("Cannot connect to Tekton client: %v", err)
		return
	}
	errDel := Delete(cs, metav1.DeleteOptions{}, obj.(*v1alpha1.MksPipelineRun).Name, obj.(*v1alpha1.MksPipelineRun).Namespace)
	if errDel != nil {
		fmt.Errorf("Cannot delete MksPipelineRun: %v", errDel)
		return
	} else {
		fmt.Println("MksPipelineRun has been deleted")

		db.Increment(rClient, "mksPipelineRundeleted")
	}

	c.queue.Add(obj)
}
