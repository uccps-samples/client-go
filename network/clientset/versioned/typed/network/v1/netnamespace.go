// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/uccps-samples/api/network/v1"
	scheme "github.com/uccps-samples/client-go/network/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NetNamespacesGetter has a method to return a NetNamespaceInterface.
// A group's client should implement this interface.
type NetNamespacesGetter interface {
	NetNamespaces() NetNamespaceInterface
}

// NetNamespaceInterface has methods to work with NetNamespace resources.
type NetNamespaceInterface interface {
	Create(ctx context.Context, netNamespace *v1.NetNamespace, opts metav1.CreateOptions) (*v1.NetNamespace, error)
	Update(ctx context.Context, netNamespace *v1.NetNamespace, opts metav1.UpdateOptions) (*v1.NetNamespace, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.NetNamespace, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.NetNamespaceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetNamespace, err error)
	NetNamespaceExpansion
}

// netNamespaces implements NetNamespaceInterface
type netNamespaces struct {
	client rest.Interface
}

// newNetNamespaces returns a NetNamespaces
func newNetNamespaces(c *NetworkV1Client) *netNamespaces {
	return &netNamespaces{
		client: c.RESTClient(),
	}
}

// Get takes name of the netNamespace, and returns the corresponding netNamespace object, and an error if there is any.
func (c *netNamespaces) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.NetNamespace, err error) {
	result = &v1.NetNamespace{}
	err = c.client.Get().
		Resource("netnamespaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetNamespaces that match those selectors.
func (c *netNamespaces) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NetNamespaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NetNamespaceList{}
	err = c.client.Get().
		Resource("netnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested netNamespaces.
func (c *netNamespaces) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("netnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a netNamespace and creates it.  Returns the server's representation of the netNamespace, and an error, if there is any.
func (c *netNamespaces) Create(ctx context.Context, netNamespace *v1.NetNamespace, opts metav1.CreateOptions) (result *v1.NetNamespace, err error) {
	result = &v1.NetNamespace{}
	err = c.client.Post().
		Resource("netnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(netNamespace).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a netNamespace and updates it. Returns the server's representation of the netNamespace, and an error, if there is any.
func (c *netNamespaces) Update(ctx context.Context, netNamespace *v1.NetNamespace, opts metav1.UpdateOptions) (result *v1.NetNamespace, err error) {
	result = &v1.NetNamespace{}
	err = c.client.Put().
		Resource("netnamespaces").
		Name(netNamespace.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(netNamespace).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the netNamespace and deletes it. Returns an error if one occurs.
func (c *netNamespaces) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("netnamespaces").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *netNamespaces) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("netnamespaces").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched netNamespace.
func (c *netNamespaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetNamespace, err error) {
	result = &v1.NetNamespace{}
	err = c.client.Patch(pt).
		Resource("netnamespaces").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
