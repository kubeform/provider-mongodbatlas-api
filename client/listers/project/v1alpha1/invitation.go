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
	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/project/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InvitationLister helps list Invitations.
// All objects returned here must be treated as read-only.
type InvitationLister interface {
	// List lists all Invitations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Invitation, err error)
	// Invitations returns an object that can list and get Invitations.
	Invitations(namespace string) InvitationNamespaceLister
	InvitationListerExpansion
}

// invitationLister implements the InvitationLister interface.
type invitationLister struct {
	indexer cache.Indexer
}

// NewInvitationLister returns a new InvitationLister.
func NewInvitationLister(indexer cache.Indexer) InvitationLister {
	return &invitationLister{indexer: indexer}
}

// List lists all Invitations in the indexer.
func (s *invitationLister) List(selector labels.Selector) (ret []*v1alpha1.Invitation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Invitation))
	})
	return ret, err
}

// Invitations returns an object that can list and get Invitations.
func (s *invitationLister) Invitations(namespace string) InvitationNamespaceLister {
	return invitationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InvitationNamespaceLister helps list and get Invitations.
// All objects returned here must be treated as read-only.
type InvitationNamespaceLister interface {
	// List lists all Invitations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Invitation, err error)
	// Get retrieves the Invitation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Invitation, error)
	InvitationNamespaceListerExpansion
}

// invitationNamespaceLister implements the InvitationNamespaceLister
// interface.
type invitationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Invitations in the indexer for a given namespace.
func (s invitationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Invitation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Invitation))
	})
	return ret, err
}

// Get retrieves the Invitation from the indexer for a given namespace and name.
func (s invitationNamespaceLister) Get(name string) (*v1alpha1.Invitation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("invitation"), name)
	}
	return obj.(*v1alpha1.Invitation), nil
}
