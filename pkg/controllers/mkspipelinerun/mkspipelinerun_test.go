package mkspipelinerun

import (
	"context"
	"fmt"
	"testing"

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned/fake"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var trGroupResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "taskruns"}

func TestConvertToTekton(t *testing.T) {
	mpr := &v1alpha1.MksPipelineRun{
		Spec: v1alpha1.MksPipelineRunSpec{
			PipelineRef: v1alpha1.MksPipelineRunRef{
				Name: "test",
			},
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      "test-converttotekton",
			Namespace: "default",
		},
	}

	converted := ConvertToTekton(mpr)

	expected := &v1beta1.PipelineRun{
		Spec: v1beta1.PipelineRunSpec{
			PipelineRef: &v1beta1.PipelineRef{
				Name: "test",
			},
			PodTemplate: &v1beta1.PodTemplate{},
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      "test-converttotekton",
			Namespace: "default",
		},
	}

	expected.Kind = "PipelineRun"
	expected.APIVersion = "tekton.dev/v1beta1"

	if compare(expected, converted) == false {
		t.Errorf("Cannot convert to tekton resource")
	}
}

func compare(pr1 *v1beta1.PipelineRun, pr2 *v1beta1.PipelineRun) bool {
	return pr1.Kind == pr2.Kind && pr1.Spec.PipelineRef.Name == pr2.Spec.PipelineRef.Name &&
		pr1.ObjectMeta.Name == pr2.ObjectMeta.Name && pr1.ObjectMeta.Namespace == pr2.ObjectMeta.Namespace
}

func TestCreateMksPipelineRun(t *testing.T) {
	prs := []*v1alpha1.MksPipelineRun{
		{
			TypeMeta: v1.TypeMeta{
				Kind: "MksPipelineRun",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      "banzaaaa",
				Namespace: "default",
			},
			Spec: v1alpha1.MksPipelineRunSpec{
				PipelineRef: v1alpha1.MksPipelineRunRef{
					Name: "hello1",
				},
			},
		},
	}

	mksClient := fake.NewSimpleClientset()
	fmt.Println(mksClient)
	for _, obj := range prs {
		crt, err := mksClient.MkscontrollerV1alpha1().MksPipelineRuns("default").Create(context.TODO(), obj, v1.CreateOptions{})
		if err != nil {
			t.Fatalf("Create MksPipelineRun failed! %s", err.Error())
		} else {
			if crt.Name != obj.Name {
				t.Fatalf("Expected resource name as %s but got %s", obj.Name, crt.Name)
			} else {
				fmt.Println(crt)
				fmt.Println("Mks Pipeline created with Name: ", crt.Name)
			}

		}
	}

}

func TestDeleteMksPipelineRun(t *testing.T) {
	prs := &v1alpha1.MksPipelineRun{
		TypeMeta: v1.TypeMeta{
			Kind: "MksPipelineRun",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      "banzaaaa",
			Namespace: "default",
		},
		Spec: v1alpha1.MksPipelineRunSpec{
			PipelineRef: v1alpha1.MksPipelineRunRef{
				Name: "hello1",
			},
		},
	}
	mksClient := fake.NewSimpleClientset(prs)
	err := mksClient.MkscontrollerV1alpha1().MksPipelineRuns(prs.Namespace).Delete(context.TODO(), prs.Name, v1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Delete MksPipelineRun failed! %s", err.Error())
	} else {
		fmt.Println("Deleted MksPipelineRun named: ", prs.Name)
	}

}
