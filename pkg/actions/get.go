package actions

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
)

func Get(gr schema.GroupVersionResource, dynamic dynamic.Interface, discovery discovery.DiscoveryInterface, objname, ns string, op metav1.GetOptions) (*unstructured.Unstructured, error) {
	gvr, err := GetGroupVersionResource(gr, discovery)
	if err != nil {
		return nil, err
	}

	obj, err := dynamic.Resource(*gvr).Namespace(ns).Get(context.Background(), objname, op)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
