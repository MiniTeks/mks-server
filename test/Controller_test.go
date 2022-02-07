package test

import (
	"context"
	"flag"
	"os/user"
	"path"
	"testing"

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned/fake"
	versioned "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned/typed/mkscontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var flags = initializeFlags()

type environmentFlags struct {
	Cluster    string
	Kubeconfig string
}

func initializeFlags() *environmentFlags {
	var f environmentFlags
	flag.StringVar(&f.Cluster, "cluster", "", "Cluster to test against")
	var defaultKubeconfig string
	if usr, err := user.Current(); err == nil {
		defaultKubeconfig = path.Join(usr.HomeDir, ".kube/config")
	}
	flag.StringVar(&f.Kubeconfig, "kubeconfig", defaultKubeconfig, "KubeConfig default path is ${HOME}/.kube/config")
	return &f
}

func setupClients(t *testing.T) (*kubernetes.Clientset, versioned.MksTaskInterface, string) {
	cfg, err := clientcmd.BuildConfigFromFlags(flags.Cluster, flags.Kubeconfig)
	if err != nil {
		t.Fatalf("Error creating config from fiel %q with cluster override %q: %s", flags.Kubeconfig, flags.Cluster, err)
	}
	// creating kube client
	k, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		t.Fatalf("Error creating kube client object from file %q with cluster override %q: %s", flags.Kubeconfig, flags.Cluster, err)
	}

	// creating mks crd client
	// mksClient, err := mksclientset.NewForConfig(cfg)
	mksClient := fake.NewSimpleClientset()
	if err != nil {
		t.Fatalf("Error creating mksClient object from file %q with cluster override %q: %s", flags.Kubeconfig, flags.Cluster, err)
	}

	// m := mksClient.MkscontrollerV1alpha1().MksTasks("default")
	m := mksClient.MkscontrollerV1alpha1().MksTasks("default")

	return k, m, "default"
}

func getMksTask(name, namespace, image, args, command string) *v1alpha1.MksTask {
	return &v1alpha1.MksTask{
		TypeMeta: v1.TypeMeta{Kind: "MksTask"},
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1alpha1.MksTaskSpec{
			Name:    name,
			Image:   image,
			Command: command,
			Args:    args,
		},
	}
}

func TestMksTaskCreate(t *testing.T) {
	_, mksClient, _ := setupClients(t)

	name := "hellot6"
	tsk := getMksTask(name, "default", "ubuntu", "-l", "ls")

	t.Logf("Creating mks-task with name %q: ", name)
	if _, err := mksClient.Create(context.Background(), tsk, v1.CreateOptions{}); err != nil {
		t.Fatalf("Failed to create mksTask %q: %s", name, err)
	}

	// tp := &tconfig.TektonParam{}
	// cs, err := tp.Client()
	// if err != nil {
	// 	t.Fatal("Couldn't create tekton client")
	// }

	// if obj, er := mkstask.GetV1beta1(cs, name, v1.GetOptions{}, namespace); er != nil {
	// 	t.Log(obj)
	// 	t.Errorf("Expected deployment %q for mksTask %q was not created: %s", name, name, er)
	// }
}

func TestMksTaskDelete(t *testing.T) {
	_, mksClient, _ := setupClients(t)

	name := "hello2"

	tsk := getMksTask(name, "default", "ubuntu", "-l", "ls")

	t.Logf("Creating mks-task with name %q: ", name)
	if _, err := mksClient.Create(context.Background(), tsk, v1.CreateOptions{}); err != nil {
		t.Fatalf("Failed to create mksTask %q: %s", name, err)
	}
	t.Logf("Deleting mks-task with name %q: ", name)
	if err := mksClient.Delete(context.Background(), name, v1.DeleteOptions{}); err != nil {
		t.Fatalf("Failed to delete mksTask %q: %s", name, err)
	}
}

func TestMksTaskUpdate(t *testing.T) {
	_, mksClient, _ := setupClients(t)

	name := "hellot6"
	tsk := getMksTask(name, "default", "ubuntu", "-l", "ls")

	t.Logf("Creating mks-task with name %q: ", name)
	if _, err := mksClient.Create(context.Background(), tsk, v1.CreateOptions{}); err != nil {
		t.Fatalf("Failed to create mksTask %q: %s", name, err)
	}

	tsk = getMksTask(name, "default", "ubuntu", "Hello World", "echo")
	t.Logf("Updating mks-task with name %q: ", name)
	if _, err := mksClient.Update(context.Background(), tsk, v1.UpdateOptions{}); err != nil {
		t.Fatalf("Failed to update mksTask %q: %s", name, err)
	}

}
