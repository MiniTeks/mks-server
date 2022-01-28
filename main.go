package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	examplecomclientset "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

var (
	kuberconfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master      = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kuberconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %v", err)
	}

	exampleClient, err := examplecomclientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %v", err)
	}

	deployment := &v1alpha1.MksPipeline{
		TypeMeta:   v1.TypeMeta{Kind: "MksPipeline"},
		ObjectMeta: v1.ObjectMeta{Name: "hello"},
		Spec:       v1alpha1.MksPipelineSpec{Name: "hello4", Image: "image", Command: "ls", Args: "args"},
	}

	crt, err := exampleClient.MkscontrollerV1alpha1().MksPipelines("default").Create(context.TODO(), deployment, v1.CreateOptions{})
	if err != nil {
		klog.Fatalf("Error creating all resources: %v", err)
	}
	fmt.Println(crt)
	
	

	
	list, err := exampleClient.MkscontrollerV1alpha1().MksPipelines("default").Get(context.TODO(), "hello", v1.GetOptions{})
	if err != nil {
		klog.Fatalf("Error listing all databases: %v", err)
	}
	fmt.Println(list)

	gt, err := exampleClient.MkscontrollerV1alpha1().MksPipelines("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		klog.Fatalf("Error listing all databases: %v", err)
	}
	for _, name := range gt.Items {
		fmt.Println(name)

	}

}
