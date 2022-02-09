package actions

import (
	"context"

	"github.com/MiniTeks/mks-server/pkg/tconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func Get(gr schema.GroupVersionResource, clients *tconfig.Client, objname, ns string, opt metav1.GetOptions) (*unstructured.Unstructured, error) {
	gvr, err := GetGroupVersionResource(gr, clients.Tekton.Discovery())
	if err != nil {
		return nil, err
	}

	obj, err := clients.Dynamic.Resource(*gvr).Namespace(ns).Get(context.Background(), objname, opt)
	if err != nil {
		return nil, err
	}
	return obj, err
}
