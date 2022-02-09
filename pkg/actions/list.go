package actions

import (
	"context"

	"github.com/MiniTeks/mks-server/pkg/tconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func List(gr schema.GroupVersionResource, clients *tconfig.Client, ns string, opt metav1.ListOptions) (*unstructured.UnstructuredList, error) {
	gvr, err := GetGroupVersionResource(gr, clients.Tekton.Discovery())
	if err != nil {
		return nil, err
	}

	objlist, err := clients.Dynamic.Resource(*gvr).Namespace(ns).List(context.Background(), opt)
	if err != nil {
		return nil, err
	}
	return objlist, nil
}
