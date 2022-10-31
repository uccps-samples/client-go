// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/uccps-samples/api/user/v1"
	scheme "github.com/uccps-samples/client-go/user/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GroupsGetter has a method to return a GroupInterface.
// A group's client should implement this interface.
type GroupsGetter interface {
	Groups() GroupInterface
}

// GroupInterface has methods to work with Group resources.
type GroupInterface interface {
	Create(ctx context.Context, group *v1.Group, opts metav1.CreateOptions) (*v1.Group, error)
	Update(ctx context.Context, group *v1.Group, opts metav1.UpdateOptions) (*v1.Group, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Group, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.GroupList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Group, err error)
	GroupExpansion
}

// groups implements GroupInterface
type groups struct {
	client rest.Interface
}

// newGroups returns a Groups
func newGroups(c *UserV1Client) *groups {
	return &groups{
		client: c.RESTClient(),
	}
}

// Get takes name of the group, and returns the corresponding group object, and an error if there is any.
func (c *groups) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Group, err error) {
	result = &v1.Group{}
	err = c.client.Get().
		Resource("groups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Groups that match those selectors.
func (c *groups) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.GroupList{}
	err = c.client.Get().
		Resource("groups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested groups.
func (c *groups) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("groups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a group and creates it.  Returns the server's representation of the group, and an error, if there is any.
func (c *groups) Create(ctx context.Context, group *v1.Group, opts metav1.CreateOptions) (result *v1.Group, err error) {
	result = &v1.Group{}
	err = c.client.Post().
		Resource("groups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(group).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a group and updates it. Returns the server's representation of the group, and an error, if there is any.
func (c *groups) Update(ctx context.Context, group *v1.Group, opts metav1.UpdateOptions) (result *v1.Group, err error) {
	result = &v1.Group{}
	err = c.client.Put().
		Resource("groups").
		Name(group.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(group).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the group and deletes it. Returns an error if one occurs.
func (c *groups) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("groups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *groups) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("groups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched group.
func (c *groups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Group, err error) {
	result = &v1.Group{}
	err = c.client.Patch(pt).
		Resource("groups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
