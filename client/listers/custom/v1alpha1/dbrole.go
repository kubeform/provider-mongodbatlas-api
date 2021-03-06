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
	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/custom/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DbRoleLister helps list DbRoles.
// All objects returned here must be treated as read-only.
type DbRoleLister interface {
	// List lists all DbRoles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DbRole, err error)
	// DbRoles returns an object that can list and get DbRoles.
	DbRoles(namespace string) DbRoleNamespaceLister
	DbRoleListerExpansion
}

// dbRoleLister implements the DbRoleLister interface.
type dbRoleLister struct {
	indexer cache.Indexer
}

// NewDbRoleLister returns a new DbRoleLister.
func NewDbRoleLister(indexer cache.Indexer) DbRoleLister {
	return &dbRoleLister{indexer: indexer}
}

// List lists all DbRoles in the indexer.
func (s *dbRoleLister) List(selector labels.Selector) (ret []*v1alpha1.DbRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DbRole))
	})
	return ret, err
}

// DbRoles returns an object that can list and get DbRoles.
func (s *dbRoleLister) DbRoles(namespace string) DbRoleNamespaceLister {
	return dbRoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DbRoleNamespaceLister helps list and get DbRoles.
// All objects returned here must be treated as read-only.
type DbRoleNamespaceLister interface {
	// List lists all DbRoles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DbRole, err error)
	// Get retrieves the DbRole from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DbRole, error)
	DbRoleNamespaceListerExpansion
}

// dbRoleNamespaceLister implements the DbRoleNamespaceLister
// interface.
type dbRoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DbRoles in the indexer for a given namespace.
func (s dbRoleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DbRole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DbRole))
	})
	return ret, err
}

// Get retrieves the DbRole from the indexer for a given namespace and name.
func (s dbRoleNamespaceLister) Get(name string) (*v1alpha1.DbRole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dbrole"), name)
	}
	return obj.(*v1alpha1.DbRole), nil
}
