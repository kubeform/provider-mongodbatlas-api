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
	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/privatelink/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EndpointLister helps list Endpoints.
// All objects returned here must be treated as read-only.
type EndpointLister interface {
	// List lists all Endpoints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Endpoint, err error)
	// Endpoints returns an object that can list and get Endpoints.
	Endpoints(namespace string) EndpointNamespaceLister
	EndpointListerExpansion
}

// endpointLister implements the EndpointLister interface.
type endpointLister struct {
	indexer cache.Indexer
}

// NewEndpointLister returns a new EndpointLister.
func NewEndpointLister(indexer cache.Indexer) EndpointLister {
	return &endpointLister{indexer: indexer}
}

// List lists all Endpoints in the indexer.
func (s *endpointLister) List(selector labels.Selector) (ret []*v1alpha1.Endpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Endpoint))
	})
	return ret, err
}

// Endpoints returns an object that can list and get Endpoints.
func (s *endpointLister) Endpoints(namespace string) EndpointNamespaceLister {
	return endpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EndpointNamespaceLister helps list and get Endpoints.
// All objects returned here must be treated as read-only.
type EndpointNamespaceLister interface {
	// List lists all Endpoints in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Endpoint, err error)
	// Get retrieves the Endpoint from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Endpoint, error)
	EndpointNamespaceListerExpansion
}

// endpointNamespaceLister implements the EndpointNamespaceLister
// interface.
type endpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Endpoints in the indexer for a given namespace.
func (s endpointNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Endpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Endpoint))
	})
	return ret, err
}

// Get retrieves the Endpoint from the indexer for a given namespace and name.
func (s endpointNamespaceLister) Get(name string) (*v1alpha1.Endpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("endpoint"), name)
	}
	return obj.(*v1alpha1.Endpoint), nil
}
