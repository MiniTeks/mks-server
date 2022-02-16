// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Satyam Bhardwaj <sabhardw@redhat.com>
// SPDX-FileCopyrightText: 2022 Utkarsh Chaurasia <uchauras@redhat.com>
// SPDX-FileCopyrightText: 2022 Avinal Kumar <avinkuma@redhat.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mkspipelinerun

import (
	"time"

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	clientset "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	appsinformers "github.com/MiniTeks/mks-server/pkg/client/informers/externalversions/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/db"
	"github.com/MiniTeks/mks-server/pkg/tconfig"
	"github.com/go-redis/redis/v8"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
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
	//wait for the informer cache to sync with apiserver
	if !cache.WaitForCacheSync(ch, c.mksPipelineCacheSynced) {
		klog.Info("\twaiting for cache to be synced\n")
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

// Add method of controller for informer function: AddFunc
func (c *controller) handleAdd(obj interface{}) {
	klog.Info("\tAdd mkspipelinerun was called\n")
	db.Increment(rClient, "MKSPIPELINERUNCREATED")
	tp := &tconfig.TektonParam{}
	cs, err := tp.Client()
	if err != nil {
		klog.Fatalf("\tCannot get the tekton client: %v\n", err)
		db.Increment(rClient, "MKSPIPELINERUNFAILED")
		return
	}
	crobj := obj.(*v1alpha1.MksPipelineRun)
	pr, err := Create(cs, crobj, metav1.CreateOptions{}, crobj.Namespace)
	if err != nil {
		klog.Errorf("\tCannot create mkspipelinerun: %v\n", err)
		db.Increment(rClient, "MKSPIPELINERUNFAILED")
		return
	} else {
		klog.Infof("\tMksPipelineRun %s created: %s\n", pr.GetName(), pr.GetUID())
		db.Increment(rClient, "MKSPIPELINERUNCOMPLETED")
	}
	c.queue.Add(obj)
}

// Update method of controller for informer function: UpdateFunc
// It gets new object & resource version can be checked to figure out if resource was edited or not
func (c *controller) handleUpdate(old, new interface{}) {
	klog.Info("\tUpdate mkspipelinerun was called\n")
}

// Delete method of controller for informer function: DeleteFunc
func (c *controller) handleDelete(obj interface{}) {
	klog.Info("\tDelete mkspipelinerun was called\n")
	tp := &tconfig.TektonParam{}
	cs, err := tp.Client()
	if err != nil {
		klog.Fatalf("\tCannot get the tekton client: %v\n", err)
		return
	}
	delObj := obj.(*v1alpha1.MksPipelineRun)
	errDel := Delete(cs, metav1.DeleteOptions{}, delObj.Name, delObj.Namespace)
	if errDel != nil {
		klog.Errorf("\tCannot delete mkspipelinerun: %v", errDel)
		return
	} else {
		klog.Infof("\tMksPipelineRun %s deleted\n", delObj.GetName())
		db.Increment(rClient, "MKSPIPELINERUNDELETED")
	}
	c.queue.Add(obj)
}
