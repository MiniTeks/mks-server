package mkstaskrun

import (
	"context"
	"fmt"
	"time"

	// "github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/klog/v2"

	// "k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	// "k8s.io/client-go/util/workqueue"

	// "k8s.io/client-go/tools/clientcmd"

	// "k8s.io/client-go/tools/clientcmd"
	// kubeclient "knative.dev/pkg/client/injection/kube/client"
	// "knative.dev/pkg/configmap"
	// "knative.dev/pkg/controller"

	// "github.com/MiniTeks/mks-server/pkg/apis/mkscontroller"
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	clientset "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	mktrinformer "github.com/MiniTeks/mks-server/pkg/client/informers/externalversions/mkscontroller/v1alpha1"
)

type Controller struct {
	kubeclientset kubernetes.Interface
	mksclientset  clientset.Interface
	// workqueue     workqueue.RateLimitingInterface
}

func NewController(kubecltst kubernetes.Interface,
	mkscltset clientset.Interface, mksinfor mktrinformer.MksTaskRunInformer) *Controller {
	// kubeconfig = "/home/avinkuma/.kube/config"
	controller := &Controller{
		kubeclientset: kubecltst,
		mksclientset:  mkscltset,
	}

	mksinfor.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("MksTaskRun has been created")
			fmt.Println(obj)
		},
		UpdateFunc: func(old, new interface{}) {
			fmt.Println("MksTaskRun has been updated")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("MksTaskRun has been deleted")
		},
	})
	return controller

}

func (c *Controller) Run(workers int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	// defer c.workqueue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	klog.Info("Starting Foo controller")

	klog.Info("Started workers")
	<-stopCh
	klog.Info("Shutting down workers")

	return nil
}

// FilteringResourceEventHandler{
// 	FilterFunc: ,
// 	ListFunc: func(lo metav1.ListOptions)(watch.Interface, error){
// 		return clientset.Interface.MkscontrollerV1alpha1().MksTaskRuns().Watch(context.Background(),lo)
// 	},

func WatchResources(clientSet clientset.Interface) cache.Store {
	projectStore, projectController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return clientSet.MkscontrollerV1alpha1().MksTaskRuns("default").List(context.Background(), lo)
			},

			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return clientSet.MkscontrollerV1alpha1().MksTaskRuns("default").Watch(context.Background(), lo)
			},
		},
		&v1alpha1.MksTaskRun{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Print("MksTaskRun has been created")
			},
			UpdateFunc: func(old, new interface{}) {
				fmt.Print("MksTaskRun has been updated")
			},
			DeleteFunc: func(obj interface{}) {
				fmt.Print("MksTaskRun has been deleted")
			},
		},
	)

	go projectController.Run(wait.NeverStop)
	return projectStore
}
