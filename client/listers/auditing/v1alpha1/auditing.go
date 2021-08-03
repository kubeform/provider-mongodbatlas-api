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
	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/auditing/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AuditingLister helps list Auditings.
// All objects returned here must be treated as read-only.
type AuditingLister interface {
	// List lists all Auditings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Auditing, err error)
	// Auditings returns an object that can list and get Auditings.
	Auditings(namespace string) AuditingNamespaceLister
	AuditingListerExpansion
}

// auditingLister implements the AuditingLister interface.
type auditingLister struct {
	indexer cache.Indexer
}

// NewAuditingLister returns a new AuditingLister.
func NewAuditingLister(indexer cache.Indexer) AuditingLister {
	return &auditingLister{indexer: indexer}
}

// List lists all Auditings in the indexer.
func (s *auditingLister) List(selector labels.Selector) (ret []*v1alpha1.Auditing, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Auditing))
	})
	return ret, err
}

// Auditings returns an object that can list and get Auditings.
func (s *auditingLister) Auditings(namespace string) AuditingNamespaceLister {
	return auditingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AuditingNamespaceLister helps list and get Auditings.
// All objects returned here must be treated as read-only.
type AuditingNamespaceLister interface {
	// List lists all Auditings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Auditing, err error)
	// Get retrieves the Auditing from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Auditing, error)
	AuditingNamespaceListerExpansion
}

// auditingNamespaceLister implements the AuditingNamespaceLister
// interface.
type auditingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Auditings in the indexer for a given namespace.
func (s auditingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Auditing, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Auditing))
	})
	return ret, err
}

// Get retrieves the Auditing from the indexer for a given namespace and name.
func (s auditingNamespaceLister) Get(name string) (*v1alpha1.Auditing, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("auditing"), name)
	}
	return obj.(*v1alpha1.Auditing), nil
}