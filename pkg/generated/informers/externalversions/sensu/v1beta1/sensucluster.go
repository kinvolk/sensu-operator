/*
Copyright 2018 The sensu-operator Authors

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

package v1beta1

import (
	time "time"

	sensu_v1beta1 "github.com/sensu/sensu-operator/pkg/apis/sensu/v1beta1"
	versioned "github.com/sensu/sensu-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/sensu/sensu-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/sensu/sensu-operator/pkg/generated/listers/sensu/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SensuClusterInformer provides access to a shared informer and lister for
// SensuClusters.
type SensuClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.SensuClusterLister
}

type sensuClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSensuClusterInformer constructs a new informer for SensuCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSensuClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSensuClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSensuClusterInformer constructs a new informer for SensuCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSensuClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SensuV1beta1().SensuClusters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SensuV1beta1().SensuClusters(namespace).Watch(options)
			},
		},
		&sensu_v1beta1.SensuCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *sensuClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSensuClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sensuClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sensu_v1beta1.SensuCluster{}, f.defaultInformer)
}

func (f *sensuClusterInformer) Lister() v1beta1.SensuClusterLister {
	return v1beta1.NewSensuClusterLister(f.Informer().GetIndexer())
}
