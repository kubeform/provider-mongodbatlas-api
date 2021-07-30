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

	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/encryption/v1alpha1"
	scheme "kubeform.dev/provider-mongodbatlas-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AtRestsGetter has a method to return a AtRestInterface.
// A group's client should implement this interface.
type AtRestsGetter interface {
	AtRests(namespace string) AtRestInterface
}

// AtRestInterface has methods to work with AtRest resources.
type AtRestInterface interface {
	Create(ctx context.Context, atRest *v1alpha1.AtRest, opts v1.CreateOptions) (*v1alpha1.AtRest, error)
	Update(ctx context.Context, atRest *v1alpha1.AtRest, opts v1.UpdateOptions) (*v1alpha1.AtRest, error)
	UpdateStatus(ctx context.Context, atRest *v1alpha1.AtRest, opts v1.UpdateOptions) (*v1alpha1.AtRest, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AtRest, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AtRestList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AtRest, err error)
	AtRestExpansion
}

// atRests implements AtRestInterface
type atRests struct {
	client rest.Interface
	ns     string
}

// newAtRests returns a AtRests
func newAtRests(c *EncryptionV1alpha1Client, namespace string) *atRests {
	return &atRests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the atRest, and returns the corresponding atRest object, and an error if there is any.
func (c *atRests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AtRest, err error) {
	result = &v1alpha1.AtRest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("atrests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AtRests that match those selectors.
func (c *atRests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AtRestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AtRestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("atrests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested atRests.
func (c *atRests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("atrests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a atRest and creates it.  Returns the server's representation of the atRest, and an error, if there is any.
func (c *atRests) Create(ctx context.Context, atRest *v1alpha1.AtRest, opts v1.CreateOptions) (result *v1alpha1.AtRest, err error) {
	result = &v1alpha1.AtRest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("atrests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(atRest).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a atRest and updates it. Returns the server's representation of the atRest, and an error, if there is any.
func (c *atRests) Update(ctx context.Context, atRest *v1alpha1.AtRest, opts v1.UpdateOptions) (result *v1alpha1.AtRest, err error) {
	result = &v1alpha1.AtRest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("atrests").
		Name(atRest.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(atRest).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *atRests) UpdateStatus(ctx context.Context, atRest *v1alpha1.AtRest, opts v1.UpdateOptions) (result *v1alpha1.AtRest, err error) {
	result = &v1alpha1.AtRest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("atrests").
		Name(atRest.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(atRest).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the atRest and deletes it. Returns an error if one occurs.
func (c *atRests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("atrests").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *atRests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("atrests").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched atRest.
func (c *atRests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AtRest, err error) {
	result = &v1alpha1.AtRest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("atrests").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
