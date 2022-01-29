package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
)

func (c *controller) run(ch <-chan struct{}) {
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

func (c *controller) worker() {
	for c.processItem() {

	}
}

func (c *controller) processItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Forget(item)
	return true
}

// Add method of controller for informer function: AddFunc
func (c *controller) handleAdd(obj interface{}) {
	fmt.Println("add was called")
	fmt.Println(obj)
	c.queue.Add(obj)
}

// Update method of controller for informer function: UpdateFunc
// It gets new object & resource version can be checked to figure out if resource was edited or not
func (c *controller) handleUpdate(old, new interface{}) {
	fmt.Println("Update was called")

}

// Delete method of controller for informer function: DeleteFunc
func (c *controller) handleDelete(obj interface{}) {
	fmt.Println("del was called")
	c.queue.Add(obj)
}
