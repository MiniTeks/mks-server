package mkstask

import (
	"fmt"
	"time"

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	examplecomclientset "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	appsinformers "github.com/MiniTeks/mks-server/pkg/client/informers/externalversions/mkscontroller/v1alpha1"
	appslisters "github.com/MiniTeks/mks-server/pkg/client/listers/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/tconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type controller struct {
	clientset          examplecomclientset.Clientset
	mksTaskLister      appslisters.MksTaskLister
	mksTaskCacheSynced cache.InformerSynced
	queue              workqueue.RateLimitingInterface
}

func NewController(clientset examplecomclientset.Clientset, mksTaskInformer appsinformers.MksTaskInformer) *controller {
	c := &controller{
		clientset:          clientset,
		mksTaskLister:      mksTaskInformer.Lister(),
		mksTaskCacheSynced: mksTaskInformer.Informer().HasSynced,
		queue:              workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "mks-task-controller"),
	}

	mksTaskInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    c.handleAdd,
			UpdateFunc: c.handleUpdate,
			DeleteFunc: c.handleDel,
		},
	)
	return c
}

func (c *controller) Run(ch <-chan struct{}) {
	fmt.Println("starting controller")
	if !cache.WaitForCacheSync(ch, c.mksTaskCacheSynced) {
		fmt.Print("waiting for cache to be synced\n")
	}

	go wait.Until(c.worker, 1*time.Second, ch)

	<-ch
}

func (c *controller) worker() {
	for c.processItem() {

	}
}

func (c *controller) processItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Forget(item)
	return true
}

func (c *controller) handleAdd(obj interface{}) {
	fmt.Println("add was called")
	fmt.Println(obj)
	tp := &tconfig.TektonParam{}
	cs, err := tp.Client()
	if err != nil {
		fmt.Errorf("Cannot get tekton client", err)
	}
	Create(cs, obj.(*v1alpha1.MksTask), metav1.CreateOptions{}, "default")
	c.queue.Add(obj)
}

func (c *controller) handleUpdate(old, obj interface{}) {
	fmt.Println("update was called")
	// fmt.Println(obj)
	c.queue.Add(obj)
}

func (c *controller) handleDel(obj interface{}) {
	fmt.Println("del was called")
	c.queue.Add(obj)
}