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

// ProviderAccessAuthorizationLister helps list ProviderAccessAuthorizations.
// All objects returned here must be treated as read-only.
type ProviderAccessAuthorizationLister interface {
	// List lists all ProviderAccessAuthorizations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProviderAccessAuthorization, err error)
	// ProviderAccessAuthorizations returns an object that can list and get ProviderAccessAuthorizations.
	ProviderAccessAuthorizations(namespace string) ProviderAccessAuthorizationNamespaceLister
	ProviderAccessAuthorizationListerExpansion
}

// providerAccessAuthorizationLister implements the ProviderAccessAuthorizationLister interface.
type providerAccessAuthorizationLister struct {
	indexer cache.Indexer
}

// NewProviderAccessAuthorizationLister returns a new ProviderAccessAuthorizationLister.
func NewProviderAccessAuthorizationLister(indexer cache.Indexer) ProviderAccessAuthorizationLister {
	return &providerAccessAuthorizationLister{indexer: indexer}
}

// List lists all ProviderAccessAuthorizations in the indexer.
func (s *providerAccessAuthorizationLister) List(selector labels.Selector) (ret []*v1alpha1.ProviderAccessAuthorization, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProviderAccessAuthorization))
	})
	return ret, err
}

// ProviderAccessAuthorizations returns an object that can list and get ProviderAccessAuthorizations.
func (s *providerAccessAuthorizationLister) ProviderAccessAuthorizations(namespace string) ProviderAccessAuthorizationNamespaceLister {
	return providerAccessAuthorizationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProviderAccessAuthorizationNamespaceLister helps list and get ProviderAccessAuthorizations.
// All objects returned here must be treated as read-only.
type ProviderAccessAuthorizationNamespaceLister interface {
	// List lists all ProviderAccessAuthorizations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProviderAccessAuthorization, err error)
	// Get retrieves the ProviderAccessAuthorization from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ProviderAccessAuthorization, error)
	ProviderAccessAuthorizationNamespaceListerExpansion
}

// providerAccessAuthorizationNamespaceLister implements the ProviderAccessAuthorizationNamespaceLister
// interface.
type providerAccessAuthorizationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProviderAccessAuthorizations in the indexer for a given namespace.
func (s providerAccessAuthorizationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProviderAccessAuthorization, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProviderAccessAuthorization))
	})
	return ret, err
}

// Get retrieves the ProviderAccessAuthorization from the indexer for a given namespace and name.
func (s providerAccessAuthorizationNamespaceLister) Get(name string) (*v1alpha1.ProviderAccessAuthorization, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("provideraccessauthorization"), name)
	}
	return obj.(*v1alpha1.ProviderAccessAuthorization), nil
}