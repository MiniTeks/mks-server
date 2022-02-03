package main

import (
	"flag"
	"fmt"
	"time"

	mksclientset "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	informers "github.com/MiniTeks/mks-server/pkg/client/informers/externalversions"
	mprcontroller "github.com/MiniTeks/mks-server/pkg/controllers/mkspipelinerun"
	mtcontroller "github.com/MiniTeks/mks-server/pkg/controllers/mkstask"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

var (
	kuberconfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master      = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {

	fmt.Println("Hello mks-server")
	klog.InitFlags(nil)
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kuberconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %v", err)
	}

	mksClient, err := mksclientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %v", err)
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %v", err)
	}

	ch := make(chan struct{})

	/* creating new instance of NewSharedInformerFactory instead of Informer to reduce the load on apiserver
	   in case on n GVRs
	   if resources of only a particular namespace is required then NewFilteredSharedInformerFactory can be
	   used
	*/
	// sync in memory cache with kubernetes cluster state in every 10 min
	informers := informers.NewSharedInformerFactory(mksClient, 10*time.Minute)
	mprc := mprcontroller.NewController(kubeClient, mksClient, informers.Mkscontroller().V1alpha1().MksPipelineRuns())
	mtc := mtcontroller.NewController(*mksClient, informers.Mkscontroller().V1alpha1().MksTasks())
	// starting informers
	informers.Start(ch)

	// starting controller by calling run() and passing channel ch
	mprc.Run(ch)
	mtc.Run(ch)
	fmt.Println(informers)
}
