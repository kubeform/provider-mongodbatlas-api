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

	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/encryption/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAtRests implements AtRestInterface
type FakeAtRests struct {
	Fake *FakeEncryptionV1alpha1
	ns   string
}

var atrestsResource = schema.GroupVersionResource{Group: "encryption.mongodbatlas.kubeform.com", Version: "v1alpha1", Resource: "atrests"}

var atrestsKind = schema.GroupVersionKind{Group: "encryption.mongodbatlas.kubeform.com", Version: "v1alpha1", Kind: "AtRest"}

// Get takes name of the atRest, and returns the corresponding atRest object, and an error if there is any.
func (c *FakeAtRests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AtRest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(atrestsResource, c.ns, name), &v1alpha1.AtRest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AtRest), err
}

// List takes label and field selectors, and returns the list of AtRests that match those selectors.
func (c *FakeAtRests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AtRestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(atrestsResource, atrestsKind, c.ns, opts), &v1alpha1.AtRestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AtRestList{ListMeta: obj.(*v1alpha1.AtRestList).ListMeta}
	for _, item := range obj.(*v1alpha1.AtRestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested atRests.
func (c *FakeAtRests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(atrestsResource, c.ns, opts))

}

// Create takes the representation of a atRest and creates it.  Returns the server's representation of the atRest, and an error, if there is any.
func (c *FakeAtRests) Create(ctx context.Context, atRest *v1alpha1.AtRest, opts v1.CreateOptions) (result *v1alpha1.AtRest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(atrestsResource, c.ns, atRest), &v1alpha1.AtRest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AtRest), err
}

// Update takes the representation of a atRest and updates it. Returns the server's representation of the atRest, and an error, if there is any.
func (c *FakeAtRests) Update(ctx context.Context, atRest *v1alpha1.AtRest, opts v1.UpdateOptions) (result *v1alpha1.AtRest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(atrestsResource, c.ns, atRest), &v1alpha1.AtRest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AtRest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAtRests) UpdateStatus(ctx context.Context, atRest *v1alpha1.AtRest, opts v1.UpdateOptions) (*v1alpha1.AtRest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(atrestsResource, "status", c.ns, atRest), &v1alpha1.AtRest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AtRest), err
}

// Delete takes name of the atRest and deletes it. Returns an error if one occurs.
func (c *FakeAtRests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(atrestsResource, c.ns, name), &v1alpha1.AtRest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAtRests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(atrestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AtRestList{})
	return err
}

// Patch applies the patch and returns the patched atRest.
func (c *FakeAtRests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AtRest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(atrestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AtRest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AtRest), err
}
