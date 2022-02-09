package actions

import (
	"context"

	"github.com/MiniTeks/mks-server/pkg/tconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func Delete(gr schema.GroupVersionResource, clients *tconfig.Client,
	objname, ns string, opt metav1.DeleteOptions) error {
	gvr, err := GetGroupVersionResource(gr, clients.Tekton.Discovery())
	if err != nil {
		return err
	}

	err = clients.Dynamic.Resource(*gvr).Namespace(ns).Delete(context.Background(), objname, opt)
	if err != nil {
		return err
	}
	return nil
}
