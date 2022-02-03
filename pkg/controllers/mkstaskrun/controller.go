package mkstaskrun

import (
	"fmt"
	"time"

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	clientset "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	informer "github.com/MiniTeks/mks-server/pkg/client/informers/externalversions/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/tconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	kubeclientset  kubernetes.Interface
	mksclientset   clientset.Interface
	mksTaskRunSync cache.InformerSynced
	queue          workqueue.RateLimitingInterface
}

func NewController(kubeclientset kubernetes.Interface,
	mksclientset clientset.Interface,
	mksinformer informer.MksTaskRunInformer) *Controller {
	controller := &Controller{
		kubeclientset:  kubeclientset,
		mksclientset:   mksclientset,
		mksTaskRunSync: mksinformer.Informer().HasSynced,
		queue:          workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "mks-taskrun-controller"),
	}

	mksinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    addController,
		UpdateFunc: updateController,
		DeleteFunc: deleteController,
	})

	return controller
}

func addController(obj interface{}) {
	fmt.Println("MksTaskRun has been created")

	tp := &tconfig.TektonParam{}
	tcl, err := tp.Client()
	if err != nil {
		fmt.Errorf("Cannot connect to Tekton client: %w", err)
	}
	ttr, err := Create(tcl, obj.(*v1alpha1.MksTaskRun), metav1.CreateOptions{}, "default")
	if err != nil {
		fmt.Errorf("Cannot create Tekton TaskRun: %w", err)
	}

	fmt.Printf("Successfully created Tekton TaskRun: %s\n", ttr.Name)
}

func updateController(oldObj, newObj interface{}) {
	fmt.Println("MksTaskRun has been updated")
}

func deleteController(obj interface{}) {
	fmt.Println("MksTaskRun has been deleted")
}

// func (c *Controller) Run(stopC <-chan struct{}) error {
// 	defer runtime.HandleCrash()
// 	fmt.Println("Starting MksTaskRun Controller")

// 	<-stopC
// 	fmt.Println("Shutting down")
// 	return nil
// }

func (c *Controller) Run(ch <-chan struct{}) {
	fmt.Println("starting controller")
	if !cache.WaitForCacheSync(ch, c.mksTaskRunSync) {
		fmt.Print("waiting for cache to be synced\n")
	}

	go wait.Until(c.worker, 1*time.Second, ch)

	<-ch
}

func (c *Controller) worker() {
	for c.processItem() {

	}
}

func (c *Controller) processItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Forget(item)
	return true
}
