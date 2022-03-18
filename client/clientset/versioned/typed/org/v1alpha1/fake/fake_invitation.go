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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/org/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInvitations implements InvitationInterface
type FakeInvitations struct {
	Fake *FakeOrgV1alpha1
	ns   string
}

var invitationsResource = schema.GroupVersionResource{Group: "org.mongodbatlas.kubeform.com", Version: "v1alpha1", Resource: "invitations"}

var invitationsKind = schema.GroupVersionKind{Group: "org.mongodbatlas.kubeform.com", Version: "v1alpha1", Kind: "Invitation"}

// Get takes name of the invitation, and returns the corresponding invitation object, and an error if there is any.
func (c *FakeInvitations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Invitation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(invitationsResource, c.ns, name), &v1alpha1.Invitation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Invitation), err
}

// List takes label and field selectors, and returns the list of Invitations that match those selectors.
func (c *FakeInvitations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InvitationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(invitationsResource, invitationsKind, c.ns, opts), &v1alpha1.InvitationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InvitationList{ListMeta: obj.(*v1alpha1.InvitationList).ListMeta}
	for _, item := range obj.(*v1alpha1.InvitationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested invitations.
func (c *FakeInvitations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(invitationsResource, c.ns, opts))

}

// Create takes the representation of a invitation and creates it.  Returns the server's representation of the invitation, and an error, if there is any.
func (c *FakeInvitations) Create(ctx context.Context, invitation *v1alpha1.Invitation, opts v1.CreateOptions) (result *v1alpha1.Invitation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(invitationsResource, c.ns, invitation), &v1alpha1.Invitation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Invitation), err
}

// Update takes the representation of a invitation and updates it. Returns the server's representation of the invitation, and an error, if there is any.
func (c *FakeInvitations) Update(ctx context.Context, invitation *v1alpha1.Invitation, opts v1.UpdateOptions) (result *v1alpha1.Invitation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(invitationsResource, c.ns, invitation), &v1alpha1.Invitation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Invitation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInvitations) UpdateStatus(ctx context.Context, invitation *v1alpha1.Invitation, opts v1.UpdateOptions) (*v1alpha1.Invitation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(invitationsResource, "status", c.ns, invitation), &v1alpha1.Invitation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Invitation), err
}

// Delete takes name of the invitation and deletes it. Returns an error if one occurs.
func (c *FakeInvitations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(invitationsResource, c.ns, name), &v1alpha1.Invitation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInvitations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(invitationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InvitationList{})
	return err
}

// Patch applies the patch and returns the patched invitation.
func (c *FakeInvitations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Invitation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(invitationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Invitation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Invitation), err
}