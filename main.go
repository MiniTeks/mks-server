package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	informers "github.com/MiniTeks/mks-server/pkg/client/informers/externalversions"
	"github.com/MiniTeks/mks-server/pkg/reconciler/mkstask"
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

	exampleClient, err := versioned.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %v", err)
	}
	ch := make(chan struct{})
	informers := informers.NewSharedInformerFactory(exampleClient, 10*time.Minute)
	c := mkstask.NewController(*exampleClient, informers.Mkscontroller().V1alpha1().MksTasks())
	informers.Start(ch)
	c.Run(ch)
	fmt.Println(informers)
}
