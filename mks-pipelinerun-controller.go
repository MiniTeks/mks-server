package main

import (
	examplecomclientset "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	appsinformers "github.com/MiniTeks/mks-server/pkg/client/informers/externalversions/mkscontroller/v1alpha1"
	appslisters "github.com/MiniTeks/mks-server/pkg/client/listers/mkscontroller/v1alpha1"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type controller struct {
	clientset              examplecomclientset.Clientset
	mksPipelineLister      appslisters.MksPipelineRunLister
	mksPipelineCacheSynced cache.InformerSynced
	queue                  workqueue.RateLimitingInterface
}

// Creates a new controller and returns
func newController(clientset examplecomclientset.Clientset, mksPipelineInformer appsinformers.MksPipelineRunInformer) *controller {
	c := &controller{
		clientset:              clientset,
		mksPipelineLister:      mksPipelineInformer.Lister(),
		mksPipelineCacheSynced: mksPipelineInformer.Informer().HasSynced,
		queue:                  workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "mks-pipeline-controller"),
	}

	//registering informer functions
	mksPipelineInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    c.handleAdd,    //AddFunc will be called whenever a new resource is created
			UpdateFunc: c.handleUpdate, //UpdateFunc will be called whenever a new resource is updated
			DeleteFunc: c.handleDelete, //DeleteFunc will be called whenever a new resource is deleted
		},
	)
	return c
}
