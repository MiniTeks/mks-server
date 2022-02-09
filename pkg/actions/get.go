package actions

import (
	"context"

	"github.com/MiniTeks/mks-server/pkg/tconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func Get(gr schema.GroupVersionResource, clients *tconfig.Client, objname, ns string, opt metav1.GetOptions) (*unstructured.Unstructured, error) {
	obj, err := clients.Dynamic.Resource(gr).Namespace(ns).Get(context.Background(), objname, opt)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
