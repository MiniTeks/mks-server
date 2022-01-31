package mkstaskrun

import (
	"fmt"

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	clientset "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	informer "github.com/MiniTeks/mks-server/pkg/client/informers/externalversions/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/tconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type Controller struct {
	kubeclientset kubernetes.Interface
	mksclientset  clientset.Interface
}

func NewController(kubeclientset kubernetes.Interface,
	mksclientset clientset.Interface,
	mksinformer informer.MksTaskRunInformer) *Controller {
	controller := &Controller{
		kubeclientset: kubeclientset,
		mksclientset:  mksclientset,
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

func (c *Controller) Run(stopC <-chan struct{}) error {
	defer runtime.HandleCrash()
	fmt.Println("Starting MksTaskRun Controller")

	<-stopC
	fmt.Println("Shutting down")
	return nil
}
