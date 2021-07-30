/*
Copyright AppsCode Inc. and Contributors

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

	privatelinkv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/privatelink/v1alpha1"
	versioned "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/listers/privatelink/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EndpointInformer provides access to a shared informer and lister for
// Endpoints.
type EndpointInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.EndpointLister
}

type endpointInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEndpointInformer constructs a new informer for Endpoint type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEndpointInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEndpointInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEndpointInformer constructs a new informer for Endpoint type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEndpointInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PrivatelinkV1alpha1().Endpoints(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PrivatelinkV1alpha1().Endpoints(namespace).Watch(context.TODO(), options)
			},
		},
		&privatelinkv1alpha1.Endpoint{},
		resyncPeriod,
		indexers,
	)
}

func (f *endpointInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEndpointInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *endpointInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&privatelinkv1alpha1.Endpoint{}, f.defaultInformer)
}

func (f *endpointInformer) Lister() v1alpha1.EndpointLister {
	return v1alpha1.NewEndpointLister(f.Informer().GetIndexer())
}
