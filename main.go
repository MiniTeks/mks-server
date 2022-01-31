package main

import (
	"flag"
	"fmt"
	"time"

	clientset "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	mksinformer "github.com/MiniTeks/mks-server/pkg/client/informers/externalversions"
	mkstrcontroller "github.com/MiniTeks/mks-server/pkg/controllers/mkstaskrun"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"knative.dev/pkg/signals"
)

var (
	kubeconfig = flag.String("kubeconfig", "", "Path to kubeconfig")
	masterUrl  = flag.String("master", "", "API server URL")
)

func main() {
	fmt.Println("Hello mks-server")

	klog.InitFlags(nil)
	flag.Parse()

	stopC := signals.SetupSignalHandler()
	cfg, err := clientcmd.BuildConfigFromFlags(*masterUrl, *kubeconfig)
	if err != nil {
		klog.Fatalf("Error getting kubeConfig: %v", err)
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error getting kube client: %v", err)
	}

	mksClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error getting mks client: %v", err)
	}

	informerfactory := mksinformer.NewSharedInformerFactory(mksClient, 10*time.Minute)

	ctrl := mkstrcontroller.NewController(kubeClient, mksClient, informerfactory.Mkscontroller().V1alpha1().MksTaskRuns())

	informerfactory.Start(stopC)

	ctrl.Run(stopC)
}
