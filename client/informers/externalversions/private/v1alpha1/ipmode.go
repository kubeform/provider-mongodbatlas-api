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

	privatev1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/private/v1alpha1"
	versioned "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-mongodbatlas-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/client/listers/private/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// IpModeInformer provides access to a shared informer and lister for
// IpModes.
type IpModeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.IpModeLister
}

type ipModeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewIpModeInformer constructs a new informer for IpMode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIpModeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIpModeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredIpModeInformer constructs a new informer for IpMode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIpModeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PrivateV1alpha1().IpModes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PrivateV1alpha1().IpModes(namespace).Watch(context.TODO(), options)
			},
		},
		&privatev1alpha1.IpMode{},
		resyncPeriod,
		indexers,
	)
}

func (f *ipModeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIpModeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ipModeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&privatev1alpha1.IpMode{}, f.defaultInformer)
}

func (f *ipModeInformer) Lister() v1alpha1.IpModeLister {
	return v1alpha1.NewIpModeLister(f.Informer().GetIndexer())
}
