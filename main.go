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

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	mksclientset "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	informers "github.com/MiniTeks/mks-server/pkg/client/informers/externalversions"
	mprcontroller "github.com/MiniTeks/mks-server/pkg/controllers/mkspipelinerun"
	mtcontroller "github.com/MiniTeks/mks-server/pkg/controllers/mkstask"
	mtrcontroller "github.com/MiniTeks/mks-server/pkg/controllers/mkstaskrun"
	"github.com/MiniTeks/mks-server/pkg/db"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	klog "k8s.io/klog/v2"
)

var (
	kuberconfig string
	master      string
	dbAddr      string
	password    string
)

func main() {

	fmt.Println("Hello mks-server")
	klog.InitFlags(nil)
	flag.Parse()
	cfg, err := clientcmd.BuildConfigFromFlags(master, kuberconfig)
	if err != nil {
		fmt.Printf("\nCouldn't build kubeconfig from user's-local: %v\n", err.Error())
		fmt.Println("Building kubeconfig from InClusterConfig")
		cfg, err = rest.InClusterConfig()
		if err != nil {
			klog.Fatalf("Error %s, getting inclusterconfig", err.Error())
		}
	}

	mksClient, err := mksclientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building mks client: %v", err)
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error getting kube client: %v", err)
	}

	// redis-db client
	cred := db.RClient{
		Addr: dbAddr,
		Pass: password,
		Db:   0,
	}
	redisClient := db.GetRedisClient(&cred)

	/* creating new instance of NewSharedInformerFactory instead of Informer to reduce the load on apiserver
	   in case on n GVRs
	   if resources of only a particular namespace is required then NewFilteredSharedInformerFactory can be
	   used
	*/
	// sync in memory cache with kubernetes cluster state in every 10 min
	ch := make(chan struct{})
	informers := informers.NewSharedInformerFactory(mksClient, 10*time.Minute)
	mprc := mprcontroller.NewController(kubeClient, mksClient, informers.Mkscontroller().V1alpha1().MksPipelineRuns(), redisClient)
	mtc := mtcontroller.NewController(*mksClient, informers.Mkscontroller().V1alpha1().MksTasks(), redisClient)
	mtrc := mtrcontroller.NewController(kubeClient, mksClient, informers.Mkscontroller().V1alpha1().MksTaskRuns(), redisClient)

	informers.Start(ch)

	// starting controller by calling run() and passing channel ch
	mprc.Run(ch)
	mtc.Run(ch)
	mtrc.Run(ch)
	fmt.Println(informers)
}

// initialize the flags.
func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Errorf("Couldn't read user's home directory!!")
	}
	home = home + "/.kube/config"
	flag.StringVar(&kuberconfig, "kubeconfig", home, "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&master, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&dbAddr, "addr", "127.0.0.1:6379", "The address of the redis server")
	flag.StringVar(&password, "password", "12345", "The password of the redis database.")
}
