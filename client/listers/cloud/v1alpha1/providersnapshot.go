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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/cloud/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProviderSnapshotLister helps list ProviderSnapshots.
// All objects returned here must be treated as read-only.
type ProviderSnapshotLister interface {
	// List lists all ProviderSnapshots in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProviderSnapshot, err error)
	// ProviderSnapshots returns an object that can list and get ProviderSnapshots.
	ProviderSnapshots(namespace string) ProviderSnapshotNamespaceLister
	ProviderSnapshotListerExpansion
}

// providerSnapshotLister implements the ProviderSnapshotLister interface.
type providerSnapshotLister struct {
	indexer cache.Indexer
}

// NewProviderSnapshotLister returns a new ProviderSnapshotLister.
func NewProviderSnapshotLister(indexer cache.Indexer) ProviderSnapshotLister {
	return &providerSnapshotLister{indexer: indexer}
}

// List lists all ProviderSnapshots in the indexer.
func (s *providerSnapshotLister) List(selector labels.Selector) (ret []*v1alpha1.ProviderSnapshot, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProviderSnapshot))
	})
	return ret, err
}

// ProviderSnapshots returns an object that can list and get ProviderSnapshots.
func (s *providerSnapshotLister) ProviderSnapshots(namespace string) ProviderSnapshotNamespaceLister {
	return providerSnapshotNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProviderSnapshotNamespaceLister helps list and get ProviderSnapshots.
// All objects returned here must be treated as read-only.
type ProviderSnapshotNamespaceLister interface {
	// List lists all ProviderSnapshots in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProviderSnapshot, err error)
	// Get retrieves the ProviderSnapshot from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ProviderSnapshot, error)
	ProviderSnapshotNamespaceListerExpansion
}

// providerSnapshotNamespaceLister implements the ProviderSnapshotNamespaceLister
// interface.
type providerSnapshotNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProviderSnapshots in the indexer for a given namespace.
func (s providerSnapshotNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProviderSnapshot, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProviderSnapshot))
	})
	return ret, err
}

// Get retrieves the ProviderSnapshot from the indexer for a given namespace and name.
func (s providerSnapshotNamespaceLister) Get(name string) (*v1alpha1.ProviderSnapshot, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("providersnapshot"), name)
	}
	return obj.(*v1alpha1.ProviderSnapshot), nil
}
