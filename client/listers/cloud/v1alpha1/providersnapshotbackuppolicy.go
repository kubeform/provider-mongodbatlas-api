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

// ProviderSnapshotBackupPolicyLister helps list ProviderSnapshotBackupPolicies.
// All objects returned here must be treated as read-only.
type ProviderSnapshotBackupPolicyLister interface {
	// List lists all ProviderSnapshotBackupPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProviderSnapshotBackupPolicy, err error)
	// ProviderSnapshotBackupPolicies returns an object that can list and get ProviderSnapshotBackupPolicies.
	ProviderSnapshotBackupPolicies(namespace string) ProviderSnapshotBackupPolicyNamespaceLister
	ProviderSnapshotBackupPolicyListerExpansion
}

// providerSnapshotBackupPolicyLister implements the ProviderSnapshotBackupPolicyLister interface.
type providerSnapshotBackupPolicyLister struct {
	indexer cache.Indexer
}

// NewProviderSnapshotBackupPolicyLister returns a new ProviderSnapshotBackupPolicyLister.
func NewProviderSnapshotBackupPolicyLister(indexer cache.Indexer) ProviderSnapshotBackupPolicyLister {
	return &providerSnapshotBackupPolicyLister{indexer: indexer}
}

// List lists all ProviderSnapshotBackupPolicies in the indexer.
func (s *providerSnapshotBackupPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ProviderSnapshotBackupPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProviderSnapshotBackupPolicy))
	})
	return ret, err
}

// ProviderSnapshotBackupPolicies returns an object that can list and get ProviderSnapshotBackupPolicies.
func (s *providerSnapshotBackupPolicyLister) ProviderSnapshotBackupPolicies(namespace string) ProviderSnapshotBackupPolicyNamespaceLister {
	return providerSnapshotBackupPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProviderSnapshotBackupPolicyNamespaceLister helps list and get ProviderSnapshotBackupPolicies.
// All objects returned here must be treated as read-only.
type ProviderSnapshotBackupPolicyNamespaceLister interface {
	// List lists all ProviderSnapshotBackupPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProviderSnapshotBackupPolicy, err error)
	// Get retrieves the ProviderSnapshotBackupPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ProviderSnapshotBackupPolicy, error)
	ProviderSnapshotBackupPolicyNamespaceListerExpansion
}

// providerSnapshotBackupPolicyNamespaceLister implements the ProviderSnapshotBackupPolicyNamespaceLister
// interface.
type providerSnapshotBackupPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProviderSnapshotBackupPolicies in the indexer for a given namespace.
func (s providerSnapshotBackupPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProviderSnapshotBackupPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProviderSnapshotBackupPolicy))
	})
	return ret, err
}

// Get retrieves the ProviderSnapshotBackupPolicy from the indexer for a given namespace and name.
func (s providerSnapshotBackupPolicyNamespaceLister) Get(name string) (*v1alpha1.ProviderSnapshotBackupPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("providersnapshotbackuppolicy"), name)
	}
	return obj.(*v1alpha1.ProviderSnapshotBackupPolicy), nil
}
