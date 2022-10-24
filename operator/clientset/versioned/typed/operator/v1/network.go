// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/uccps-samples/api/operator/v1"
	scheme "github.com/uccps-samples/client-go/operator/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NetworksGetter has a method to return a NetworkInterface.
// A group's client should implement this interface.
type NetworksGetter interface {
	Networks() NetworkInterface
}

// NetworkInterface has methods to work with Network resources.
type NetworkInterface interface {
	Create(ctx context.Context, network *v1.Network, opts metav1.CreateOptions) (*v1.Network, error)
	Update(ctx context.Context, network *v1.Network, opts metav1.UpdateOptions) (*v1.Network, error)
	UpdateStatus(ctx context.Context, network *v1.Network, opts metav1.UpdateOptions) (*v1.Network, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Network, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.NetworkList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Network, err error)
	NetworkExpansion
}

// networks implements NetworkInterface
type networks struct {
	client rest.Interface
}

// newNetworks returns a Networks
func newNetworks(c *OperatorV1Client) *networks {
	return &networks{
		client: c.RESTClient(),
	}
}

// Get takes name of the network, and returns the corresponding network object, and an error if there is any.
func (c *networks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Network, err error) {
	result = &v1.Network{}
	err = c.client.Get().
		Resource("networks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Networks that match those selectors.
func (c *networks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NetworkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NetworkList{}
	err = c.client.Get().
		Resource("networks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networks.
func (c *networks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("networks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a network and creates it.  Returns the server's representation of the network, and an error, if there is any.
func (c *networks) Create(ctx context.Context, network *v1.Network, opts metav1.CreateOptions) (result *v1.Network, err error) {
	result = &v1.Network{}
	err = c.client.Post().
		Resource("networks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(network).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a network and updates it. Returns the server's representation of the network, and an error, if there is any.
func (c *networks) Update(ctx context.Context, network *v1.Network, opts metav1.UpdateOptions) (result *v1.Network, err error) {
	result = &v1.Network{}
	err = c.client.Put().
		Resource("networks").
		Name(network.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(network).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *networks) UpdateStatus(ctx context.Context, network *v1.Network, opts metav1.UpdateOptions) (result *v1.Network, err error) {
	result = &v1.Network{}
	err = c.client.Put().
		Resource("networks").
		Name(network.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(network).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the network and deletes it. Returns an error if one occurs.
func (c *networks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("networks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("networks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched network.
func (c *networks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Network, err error) {
	result = &v1.Network{}
	err = c.client.Patch(pt).
		Resource("networks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
