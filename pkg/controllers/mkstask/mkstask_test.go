package mkstask

import (
	"context"
	"testing"

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned/fake"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"gotest.tools/v3/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestConvertToTekton(t *testing.T) {
	mt := &v1alpha1.MksTask{
		TypeMeta: metav1.TypeMeta{Kind: "MksTask"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-mkstask",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskSpec{
			Name:    "test-mkstask",
			Image:   "ubuntu",
			Command: "ls",
			Args:    "-l",
		},
	}
	converted := ConvertToTekton(mt)
	//expected output
	res := &v1beta1.Task{}
	res.Kind = "Task"
	res.APIVersion = "tekton.dev/v1beta1"
	res.ObjectMeta = metav1.ObjectMeta{
		Name:      "test-mkstask",
		Namespace: "default",
	}
	res.Spec = v1beta1.TaskSpec{
		Steps: []v1beta1.Step{
			{
				Container: v1.Container{
					Image:   "ubuntu",
					Name:    "test-mkstask",
					Command: []string{"ls"},
					Args:    []string{"-l"},
				},
			},
		},
	}

	if !compare(res, converted) {
		t.Errorf("Cannot convert to tekton successfully")
	}
}

func compare(tr1 *v1beta1.Task, tr2 *v1beta1.Task) bool {
	return tr1.Kind == tr2.Kind && tr1.Spec.Steps[0].Name == tr2.Spec.Steps[0].Name && tr1.ObjectMeta.Name == tr2.ObjectMeta.Name && tr1.ObjectMeta.Namespace == tr2.ObjectMeta.Namespace
}

func TestCreate(t *testing.T) {
	mt := &v1alpha1.MksTask{
		TypeMeta: metav1.TypeMeta{Kind: "MksTask"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-mkstask",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskSpec{
			Name:    "test-mkstask",
			Image:   "ubuntu",
			Command: "ls",
			Args:    "-l",
		},
	}
	mkstaskClient := fake.NewSimpleClientset()
	crt, err := mkstaskClient.MkscontrollerV1alpha1().MksTasks("default").Create(context.Background(), mt, metav1.CreateOptions{})
	if err != nil {
		t.Errorf("Cannot create %s mkstask: %v", crt.ObjectMeta.Name, err)
	}
	t.Logf("Successfully created mkstask: %s", crt.ObjectMeta.Name)
}

func TestDelete(t *testing.T) {
	mt := &v1alpha1.MksTask{
		TypeMeta: metav1.TypeMeta{Kind: "MksTask"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-mkstask",
			Namespace: "default",
		},
		Spec: v1alpha1.MksTaskSpec{
			Name:    "test-mkstask",
			Image:   "ubuntu",
			Command: "ls",
			Args:    "-l",
		},
	}
	mkstaskClient := fake.NewSimpleClientset(runtime.Object(mt))
	err := mkstaskClient.MkscontrollerV1alpha1().MksTasks("default").Delete(context.Background(), mt.Name, metav1.DeleteOptions{})
	if err != nil {
		t.Errorf("Cannot delete %s mkstask: %v", mt.ObjectMeta.Name, err)
	}
	t.Logf("Successfully deleted mkstask: %s", mt.ObjectMeta.Name)
}

func TestList(t *testing.T) {
	mtl := []*v1alpha1.MksTask{
		{
			TypeMeta: metav1.TypeMeta{Kind: "MksTask"},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-mkstask",
				Namespace: "default",
			},
			Spec: v1alpha1.MksTaskSpec{
				Name:    "test-mkstask",
				Image:   "ubuntu",
				Command: "ls",
				Args:    "-l",
			},
		},
		{
			TypeMeta: metav1.TypeMeta{Kind: "MksTask"},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-mkstask2",
				Namespace: "default",
			},
			Spec: v1alpha1.MksTaskSpec{
				Name:    "test-mkstask2",
				Image:   "ubuntu",
				Command: "ls",
				Args:    "-l",
			},
		},
	}
	mkstaskClient := fake.NewSimpleClientset(mtl[0], mtl[1])
	crt, err := mkstaskClient.MkscontrollerV1alpha1().MksTasks("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		t.Errorf("Cannot list mkstasks: %v", err)
	}
	assert.Assert(t, mtl[0].ObjectMeta.Name == crt.Items[0].ObjectMeta.Name)
	assert.Assert(t, mtl[1].ObjectMeta.Name == crt.Items[1].ObjectMeta.Name, 1)
	t.Logf("Successfully listed mkstasks: %s, %s ", crt.Items[0].ObjectMeta.Name, crt.Items[1].ObjectMeta.Name)
}
