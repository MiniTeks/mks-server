package mkspipelinerun

import (
	"fmt"
	"time"

	clientset "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	appsinformers "github.com/MiniTeks/mks-server/pkg/client/informers/externalversions/mkscontroller/v1alpha1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"github.com/go-redis/redis/v8"
)

type controller struct {
	kubeclientset          kubernetes.Interface
	mksclientset           clientset.Interface
	queue                  workqueue.RateLimitingInterface
	mksPipelineCacheSynced cache.InformerSynced
}

var rClient *redis.Client

// Creates a new controller and returns
func NewController(kubectlst kubernetes.Interface, clientset clientset.Interface, mksPipelineRunInformer appsinformers.MksPipelineRunInformer, redisClient *redis.Client) *controller {
	rClient = redisClient
	c := &controller{
		mksclientset:           clientset,
		kubeclientset:          kubectlst,
		mksPipelineCacheSynced: mksPipelineRunInformer.Informer().HasSynced,
		queue:                  workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "mks-pipeline-controller"),
	}

	//registering informer functions
	mksPipelineRunInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    c.handleAdd,    //AddFunc will be called whenever a new resource is created
			UpdateFunc: c.handleUpdate, //UpdateFunc will be called whenever a new resource is updated
			DeleteFunc: c.handleDelete, //DeleteFunc will be called whenever a new resource is deleted
		},
	)
	return c
}

func (c *controller) Run(ch <-chan struct{}) {
	fmt.Println("starting controller")

	//wait for the informer cache to sync with apiserver
	if !cache.WaitForCacheSync(ch, c.mksPipelineCacheSynced) {
		fmt.Print("waiting for cache to be synced\n")
	}

	// Calls worker() in every 1 sec until channel ch is closed
	go wait.Until(c.worker, 1*time.Second, ch)

	// Wait for something to be there in ch & if empty then above wait.Until would not return
	<-ch
}

// If processItem returns false then worker() would return and then wait.Until again calls worker() after 1sec
func (c *controller) worker() {
	for c.processItem() {

	}
}

func (c *controller) processItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	// To make sure the particular item is not processed again
	defer c.queue.Forget(item)
	return true
}
