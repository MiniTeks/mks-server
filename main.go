package main

import (
	// "context"
	"flag"
	"time"

	// "fmt"
	// "github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	examplecomclientset "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/util/runtime"
	"knative.dev/pkg/signals"
	// v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	informers "github.com/MiniTeks/mks-server/pkg/client/informers/externalversions"
	conr "github.com/MiniTeks/mks-server/pkg/reconciler/mkstaskrun"
	"k8s.io/client-go/kubernetes"
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
	stopc := signals.SetupSignalHandler()
	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kuberconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %v", err)
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %v", err)
	}

	exampleClient, err := examplecomclientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %v", err)
	}
	infac := informers.NewSharedInformerFactory(exampleClient, 10*time.Minute)

	 conr.NewController(kubeClient, exampleClient, infac.Mkscontroller().V1alpha1().MksTaskRuns())

	infac.Start(stopc)

	defer runtime.HandleCrash()
	<-stopc

	// controller.Run(2, stopc)
}

// 	deployment := &v1alpha1.MksTask{
// 		TypeMeta:   v1.TypeMeta{Kind: "MksTask"},
// 		ObjectMeta: v1.ObjectMeta{Name: "hello"},
// 		Spec:       v1alpha1.MksTaskSpec{Name: "hello4", Image: "image", Command: "ls", Args: "args"},
// 	}

// 	crt, err := exampleClient.MkscontrollerV1alpha1().MksTasks("default").Create(context.TODO(), deployment, v1.CreateOptions{})
// 	if err != nil {
// 		klog.Fatalf("Error creating all resources: %v", err)
// 	}
// 	fmt.Println(crt)
// 	deployment1 := &v1alpha1.MksTask{
// 		TypeMeta:   v1.TypeMeta{Kind: "MksTask"},
// 		ObjectMeta: v1.ObjectMeta{Name: "hello1"},
// 		Spec:       v1alpha1.MksTaskSpec{Name: "hello1", Image: "image", Command: "ls", Args: "args"},
// 	}
// 	crt1, err := exampleClient.MkscontrollerV1alpha1().MksTasks("default").Create(context.TODO(), deployment1, v1.CreateOptions{})
// 	if err != nil {
// 		klog.Fatalf("Error creating all resources: %v", err)
// 	}
// 	fmt.Println(crt1)

// 	deployment2 := &v1alpha1.MksTask{
// 		TypeMeta:   v1.TypeMeta{Kind: "MksTask"},
// 		ObjectMeta: v1.ObjectMeta{Name: "hello2"},
// 		Spec:       v1alpha1.MksTaskSpec{Name: "hello2", Image: "image", Command: "ls", Args: "args"},
// 	}
// 	crt2, err := exampleClient.MkscontrollerV1alpha1().MksTasks("default").Create(context.TODO(), deployment2, v1.CreateOptions{})
// 	if err != nil {
// 		klog.Fatalf("Error creating all resources: %v", err)
// 	}
// 	fmt.Println(crt2)

// 	deployment3 := &v1alpha1.MksTask{
// 		TypeMeta:   v1.TypeMeta{Kind: "MksTask"},
// 		ObjectMeta: v1.ObjectMeta{Name: "hello3"},
// 		Spec:       v1alpha1.MksTaskSpec{Name: "hello3", Image: "image", Command: "ls", Args: "args"},
// 	}
// 	crt3, err3 := exampleClient.MkscontrollerV1alpha1().MksTasks("default").Create(context.TODO(), deployment3, v1.CreateOptions{})
// 	if err3 != nil {
// 		klog.Fatalf("Error creating all resources: %v", err)
// 	}

// 	fmt.Println(crt3)
// 	list, err := exampleClient.MkscontrollerV1alpha1().MksTasks("default").Get(context.TODO(), "hello", v1.GetOptions{})
// 	if err != nil {
// 		klog.Fatalf("Error listing all databases: %v", err)
// 	}
// 	fmt.Println(list)

// 	gt, err := exampleClient.MkscontrollerV1alpha1().MksTasks("default").List(context.TODO(), v1.ListOptions{})
// 	if err != nil {
// 		klog.Fatalf("Error listing all databases: %v", err)
// 	}
// 	for _, name := range gt.Items {
// 		fmt.Println(name)

// 	}

// }

// package main

// import (
// 	"github.com/MiniTeks/mks-server/pkg/reconciler/mkstaskrun"
// 	"knative.dev/pkg/injection/sharedmain"
// )

// func main() {
// 	sharedmain.Main("controller", mkstaskrun.NewController)
// }
