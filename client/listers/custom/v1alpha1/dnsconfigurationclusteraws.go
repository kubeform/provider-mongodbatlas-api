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

// DnsConfigurationClusterAwsLister helps list DnsConfigurationClusterAwses.
// All objects returned here must be treated as read-only.
type DnsConfigurationClusterAwsLister interface {
	// List lists all DnsConfigurationClusterAwses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DnsConfigurationClusterAws, err error)
	// DnsConfigurationClusterAwses returns an object that can list and get DnsConfigurationClusterAwses.
	DnsConfigurationClusterAwses(namespace string) DnsConfigurationClusterAwsNamespaceLister
	DnsConfigurationClusterAwsListerExpansion
}

// dnsConfigurationClusterAwsLister implements the DnsConfigurationClusterAwsLister interface.
type dnsConfigurationClusterAwsLister struct {
	indexer cache.Indexer
}

// NewDnsConfigurationClusterAwsLister returns a new DnsConfigurationClusterAwsLister.
func NewDnsConfigurationClusterAwsLister(indexer cache.Indexer) DnsConfigurationClusterAwsLister {
	return &dnsConfigurationClusterAwsLister{indexer: indexer}
}

// List lists all DnsConfigurationClusterAwses in the indexer.
func (s *dnsConfigurationClusterAwsLister) List(selector labels.Selector) (ret []*v1alpha1.DnsConfigurationClusterAws, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsConfigurationClusterAws))
	})
	return ret, err
}

// DnsConfigurationClusterAwses returns an object that can list and get DnsConfigurationClusterAwses.
func (s *dnsConfigurationClusterAwsLister) DnsConfigurationClusterAwses(namespace string) DnsConfigurationClusterAwsNamespaceLister {
	return dnsConfigurationClusterAwsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DnsConfigurationClusterAwsNamespaceLister helps list and get DnsConfigurationClusterAwses.
// All objects returned here must be treated as read-only.
type DnsConfigurationClusterAwsNamespaceLister interface {
	// List lists all DnsConfigurationClusterAwses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DnsConfigurationClusterAws, err error)
	// Get retrieves the DnsConfigurationClusterAws from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DnsConfigurationClusterAws, error)
	DnsConfigurationClusterAwsNamespaceListerExpansion
}

// dnsConfigurationClusterAwsNamespaceLister implements the DnsConfigurationClusterAwsNamespaceLister
// interface.
type dnsConfigurationClusterAwsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DnsConfigurationClusterAwses in the indexer for a given namespace.
func (s dnsConfigurationClusterAwsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DnsConfigurationClusterAws, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsConfigurationClusterAws))
	})
	return ret, err
}

// Get retrieves the DnsConfigurationClusterAws from the indexer for a given namespace and name.
func (s dnsConfigurationClusterAwsNamespaceLister) Get(name string) (*v1alpha1.DnsConfigurationClusterAws, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dnsconfigurationclusteraws"), name)
	}
	return obj.(*v1alpha1.DnsConfigurationClusterAws), nil
}
