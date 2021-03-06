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

	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/project/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIpAccessLists implements IpAccessListInterface
type FakeIpAccessLists struct {
	Fake *FakeProjectV1alpha1
	ns   string
}

var ipaccesslistsResource = schema.GroupVersionResource{Group: "project.mongodbatlas.kubeform.com", Version: "v1alpha1", Resource: "ipaccesslists"}

var ipaccesslistsKind = schema.GroupVersionKind{Group: "project.mongodbatlas.kubeform.com", Version: "v1alpha1", Kind: "IpAccessList"}

// Get takes name of the ipAccessList, and returns the corresponding ipAccessList object, and an error if there is any.
func (c *FakeIpAccessLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IpAccessList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ipaccesslistsResource, c.ns, name), &v1alpha1.IpAccessList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpAccessList), err
}

// List takes label and field selectors, and returns the list of IpAccessLists that match those selectors.
func (c *FakeIpAccessLists) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IpAccessListList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ipaccesslistsResource, ipaccesslistsKind, c.ns, opts), &v1alpha1.IpAccessListList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IpAccessListList{ListMeta: obj.(*v1alpha1.IpAccessListList).ListMeta}
	for _, item := range obj.(*v1alpha1.IpAccessListList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ipAccessLists.
func (c *FakeIpAccessLists) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ipaccesslistsResource, c.ns, opts))

}

// Create takes the representation of a ipAccessList and creates it.  Returns the server's representation of the ipAccessList, and an error, if there is any.
func (c *FakeIpAccessLists) Create(ctx context.Context, ipAccessList *v1alpha1.IpAccessList, opts v1.CreateOptions) (result *v1alpha1.IpAccessList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ipaccesslistsResource, c.ns, ipAccessList), &v1alpha1.IpAccessList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpAccessList), err
}

// Update takes the representation of a ipAccessList and updates it. Returns the server's representation of the ipAccessList, and an error, if there is any.
func (c *FakeIpAccessLists) Update(ctx context.Context, ipAccessList *v1alpha1.IpAccessList, opts v1.UpdateOptions) (result *v1alpha1.IpAccessList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ipaccesslistsResource, c.ns, ipAccessList), &v1alpha1.IpAccessList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpAccessList), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIpAccessLists) UpdateStatus(ctx context.Context, ipAccessList *v1alpha1.IpAccessList, opts v1.UpdateOptions) (*v1alpha1.IpAccessList, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ipaccesslistsResource, "status", c.ns, ipAccessList), &v1alpha1.IpAccessList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpAccessList), err
}

// Delete takes name of the ipAccessList and deletes it. Returns an error if one occurs.
func (c *FakeIpAccessLists) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ipaccesslistsResource, c.ns, name), &v1alpha1.IpAccessList{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIpAccessLists) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ipaccesslistsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IpAccessListList{})
	return err
}

// Patch applies the patch and returns the patched ipAccessList.
func (c *FakeIpAccessLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IpAccessList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ipaccesslistsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IpAccessList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpAccessList), err
}
