/*
Copyright 2022 Google LLC

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

	scheme "github.com/openconfig/lemming/operator/api/clientset/scheme"
	v1alpha1 "github.com/openconfig/lemming/operator/api/lemming/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LemmingsGetter has a method to return a LemmingInterface.
// A group's client should implement this interface.
type LemmingsGetter interface {
	Lemmings(namespace string) LemmingInterface
}

// LemmingInterface has methods to work with Lemming resources.
type LemmingInterface interface {
	Create(ctx context.Context, lemming *v1alpha1.Lemming, opts v1.CreateOptions) (*v1alpha1.Lemming, error)
	Update(ctx context.Context, lemming *v1alpha1.Lemming, opts v1.UpdateOptions) (*v1alpha1.Lemming, error)
	UpdateStatus(ctx context.Context, lemming *v1alpha1.Lemming, opts v1.UpdateOptions) (*v1alpha1.Lemming, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Lemming, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.LemmingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Lemming, err error)
	LemmingExpansion
}

// lemmings implements LemmingInterface
type lemmings struct {
	client rest.Interface
	ns     string
}

// newLemmings returns a Lemmings
func newLemmings(c *LemmingV1alpha1Client, namespace string) *lemmings {
	return &lemmings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lemming, and returns the corresponding lemming object, and an error if there is any.
func (c *lemmings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Lemming, err error) {
	result = &v1alpha1.Lemming{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lemmings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Lemmings that match those selectors.
func (c *lemmings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LemmingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LemmingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lemmings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lemmings.
func (c *lemmings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("lemmings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a lemming and creates it.  Returns the server's representation of the lemming, and an error, if there is any.
func (c *lemmings) Create(ctx context.Context, lemming *v1alpha1.Lemming, opts v1.CreateOptions) (result *v1alpha1.Lemming, err error) {
	result = &v1alpha1.Lemming{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("lemmings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lemming).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a lemming and updates it. Returns the server's representation of the lemming, and an error, if there is any.
func (c *lemmings) Update(ctx context.Context, lemming *v1alpha1.Lemming, opts v1.UpdateOptions) (result *v1alpha1.Lemming, err error) {
	result = &v1alpha1.Lemming{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lemmings").
		Name(lemming.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lemming).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *lemmings) UpdateStatus(ctx context.Context, lemming *v1alpha1.Lemming, opts v1.UpdateOptions) (result *v1alpha1.Lemming, err error) {
	result = &v1alpha1.Lemming{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lemmings").
		Name(lemming.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lemming).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the lemming and deletes it. Returns an error if one occurs.
func (c *lemmings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lemmings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lemmings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lemmings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched lemming.
func (c *lemmings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Lemming, err error) {
	result = &v1alpha1.Lemming{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("lemmings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
