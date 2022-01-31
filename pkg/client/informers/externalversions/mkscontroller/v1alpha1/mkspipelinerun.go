/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	mkscontrollerv1alpha1 "github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	versioned "github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	internalinterfaces "github.com/MiniTeks/mks-server/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/MiniTeks/mks-server/pkg/client/listers/mkscontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MksPipelineRunInformer provides access to a shared informer and lister for
// MksPipelineRuns.
type MksPipelineRunInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MksPipelineRunLister
}

type mksPipelineRunInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMksPipelineRunInformer constructs a new informer for MksPipelineRun type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMksPipelineRunInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMksPipelineRunInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMksPipelineRunInformer constructs a new informer for MksPipelineRun type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMksPipelineRunInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MkscontrollerV1alpha1().MksPipelineRuns(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MkscontrollerV1alpha1().MksPipelineRuns(namespace).Watch(context.TODO(), options)
			},
		},
		&mkscontrollerv1alpha1.MksPipelineRun{},
		resyncPeriod,
		indexers,
	)
}

func (f *mksPipelineRunInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMksPipelineRunInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mksPipelineRunInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mkscontrollerv1alpha1.MksPipelineRun{}, f.defaultInformer)
}

func (f *mksPipelineRunInformer) Lister() v1alpha1.MksPipelineRunLister {
	return v1alpha1.NewMksPipelineRunLister(f.Informer().GetIndexer())
}