package actions

import (
	"sync"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/restmapper"
)

var (
	doOnce      sync.Once
	apiGroupRes []*restmapper.APIGroupResources
)

func GetGroupVersionResource(gr schema.GroupVersionResource, discovery discovery.DiscoveryInterface) (*schema.GroupVersionResource, error) {
	var err error
	doOnce.Do(func() {
		err = InitializeAPIGroupRes(discovery)
	})
	if err != nil {
		return nil, err
	}

	rm := restmapper.NewDiscoveryRESTMapper(apiGroupRes)
	gvr, err := rm.ResourceFor(gr)
	if err != nil {
		return nil, err
	}

	return &gvr, nil
}

func InitializeAPIGroupRes(discovery discovery.DiscoveryInterface) error {
	var err error
	apiGroupRes, err = restmapper.GetAPIGroupResources(discovery)
	if err != nil {
		return err
	}
	return nil
}
