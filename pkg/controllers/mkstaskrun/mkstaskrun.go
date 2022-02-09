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

var trGroupResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "taskruns"}

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

func Delete(tcl *tconfig.Client, mtrname string, opt metav1.DeleteOptions, ns string) error {
	if err := actions.Delete(trGroupResource, tcl, mtrname, ns, opt); err != nil {
		return err
	}
	return nil
}
