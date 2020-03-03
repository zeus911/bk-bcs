/*
Copyright 2020 The Kubernetes Authors.

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
	multiclusterdns_v1alpha1 "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/multiclusterdns/v1alpha1"
	versioned "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned"
	internalinterfaces "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/informers/externalversions/internalinterfaces"
	v1alpha1 "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/listers/multiclusterdns/v1alpha1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceDNSRecordInformer provides access to a shared informer and lister for
// ServiceDNSRecords.
type ServiceDNSRecordInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ServiceDNSRecordLister
}

type serviceDNSRecordInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceDNSRecordInformer constructs a new informer for ServiceDNSRecord type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceDNSRecordInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceDNSRecordInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceDNSRecordInformer constructs a new informer for ServiceDNSRecord type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceDNSRecordInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MulticlusterdnsV1alpha1().ServiceDNSRecords(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MulticlusterdnsV1alpha1().ServiceDNSRecords(namespace).Watch(options)
			},
		},
		&multiclusterdns_v1alpha1.ServiceDNSRecord{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceDNSRecordInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceDNSRecordInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceDNSRecordInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&multiclusterdns_v1alpha1.ServiceDNSRecord{}, f.defaultInformer)
}

func (f *serviceDNSRecordInformer) Lister() v1alpha1.ServiceDNSRecordLister {
	return v1alpha1.NewServiceDNSRecordLister(f.Informer().GetIndexer())
}
