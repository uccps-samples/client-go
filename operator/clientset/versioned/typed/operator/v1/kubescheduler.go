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

// KubeSchedulersGetter has a method to return a KubeSchedulerInterface.
// A group's client should implement this interface.
type KubeSchedulersGetter interface {
	KubeSchedulers() KubeSchedulerInterface
}

// KubeSchedulerInterface has methods to work with KubeScheduler resources.
type KubeSchedulerInterface interface {
	Create(ctx context.Context, kubeScheduler *v1.KubeScheduler, opts metav1.CreateOptions) (*v1.KubeScheduler, error)
	Update(ctx context.Context, kubeScheduler *v1.KubeScheduler, opts metav1.UpdateOptions) (*v1.KubeScheduler, error)
	UpdateStatus(ctx context.Context, kubeScheduler *v1.KubeScheduler, opts metav1.UpdateOptions) (*v1.KubeScheduler, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.KubeScheduler, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.KubeSchedulerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.KubeScheduler, err error)
	KubeSchedulerExpansion
}

// kubeSchedulers implements KubeSchedulerInterface
type kubeSchedulers struct {
	client rest.Interface
}

// newKubeSchedulers returns a KubeSchedulers
func newKubeSchedulers(c *OperatorV1Client) *kubeSchedulers {
	return &kubeSchedulers{
		client: c.RESTClient(),
	}
}

// Get takes name of the kubeScheduler, and returns the corresponding kubeScheduler object, and an error if there is any.
func (c *kubeSchedulers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.KubeScheduler, err error) {
	result = &v1.KubeScheduler{}
	err = c.client.Get().
		Resource("kubeschedulers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KubeSchedulers that match those selectors.
func (c *kubeSchedulers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.KubeSchedulerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.KubeSchedulerList{}
	err = c.client.Get().
		Resource("kubeschedulers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kubeSchedulers.
func (c *kubeSchedulers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("kubeschedulers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a kubeScheduler and creates it.  Returns the server's representation of the kubeScheduler, and an error, if there is any.
func (c *kubeSchedulers) Create(ctx context.Context, kubeScheduler *v1.KubeScheduler, opts metav1.CreateOptions) (result *v1.KubeScheduler, err error) {
	result = &v1.KubeScheduler{}
	err = c.client.Post().
		Resource("kubeschedulers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kubeScheduler).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a kubeScheduler and updates it. Returns the server's representation of the kubeScheduler, and an error, if there is any.
func (c *kubeSchedulers) Update(ctx context.Context, kubeScheduler *v1.KubeScheduler, opts metav1.UpdateOptions) (result *v1.KubeScheduler, err error) {
	result = &v1.KubeScheduler{}
	err = c.client.Put().
		Resource("kubeschedulers").
		Name(kubeScheduler.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kubeScheduler).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *kubeSchedulers) UpdateStatus(ctx context.Context, kubeScheduler *v1.KubeScheduler, opts metav1.UpdateOptions) (result *v1.KubeScheduler, err error) {
	result = &v1.KubeScheduler{}
	err = c.client.Put().
		Resource("kubeschedulers").
		Name(kubeScheduler.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kubeScheduler).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the kubeScheduler and deletes it. Returns an error if one occurs.
func (c *kubeSchedulers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("kubeschedulers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kubeSchedulers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("kubeschedulers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched kubeScheduler.
func (c *kubeSchedulers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.KubeScheduler, err error) {
	result = &v1.KubeScheduler{}
	err = c.client.Patch(pt).
		Resource("kubeschedulers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
