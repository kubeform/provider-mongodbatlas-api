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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/data/v1alpha1"
	scheme "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LakesGetter has a method to return a LakeInterface.
// A group's client should implement this interface.
type LakesGetter interface {
	Lakes(namespace string) LakeInterface
}

// LakeInterface has methods to work with Lake resources.
type LakeInterface interface {
	Create(ctx context.Context, lake *v1alpha1.Lake, opts v1.CreateOptions) (*v1alpha1.Lake, error)
	Update(ctx context.Context, lake *v1alpha1.Lake, opts v1.UpdateOptions) (*v1alpha1.Lake, error)
	UpdateStatus(ctx context.Context, lake *v1alpha1.Lake, opts v1.UpdateOptions) (*v1alpha1.Lake, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Lake, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.LakeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Lake, err error)
	LakeExpansion
}

// lakes implements LakeInterface
type lakes struct {
	client rest.Interface
	ns     string
}

// newLakes returns a Lakes
func newLakes(c *DataV1alpha1Client, namespace string) *lakes {
	return &lakes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lake, and returns the corresponding lake object, and an error if there is any.
func (c *lakes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Lake, err error) {
	result = &v1alpha1.Lake{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lakes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Lakes that match those selectors.
func (c *lakes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LakeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LakeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lakes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lakes.
func (c *lakes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("lakes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a lake and creates it.  Returns the server's representation of the lake, and an error, if there is any.
func (c *lakes) Create(ctx context.Context, lake *v1alpha1.Lake, opts v1.CreateOptions) (result *v1alpha1.Lake, err error) {
	result = &v1alpha1.Lake{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("lakes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lake).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a lake and updates it. Returns the server's representation of the lake, and an error, if there is any.
func (c *lakes) Update(ctx context.Context, lake *v1alpha1.Lake, opts v1.UpdateOptions) (result *v1alpha1.Lake, err error) {
	result = &v1alpha1.Lake{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lakes").
		Name(lake.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lake).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *lakes) UpdateStatus(ctx context.Context, lake *v1alpha1.Lake, opts v1.UpdateOptions) (result *v1alpha1.Lake, err error) {
	result = &v1alpha1.Lake{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lakes").
		Name(lake.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lake).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the lake and deletes it. Returns an error if one occurs.
func (c *lakes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lakes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lakes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lakes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched lake.
func (c *lakes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Lake, err error) {
	result = &v1alpha1.Lake{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("lakes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
