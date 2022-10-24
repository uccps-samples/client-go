// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	operatorv1 "github.com/uccps-samples/api/operator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAuthentications implements AuthenticationInterface
type FakeAuthentications struct {
	Fake *FakeOperatorV1
}

var authenticationsResource = schema.GroupVersionResource{Group: "operator.uccp.io", Version: "v1", Resource: "authentications"}

var authenticationsKind = schema.GroupVersionKind{Group: "operator.uccp.io", Version: "v1", Kind: "Authentication"}

// Get takes name of the authentication, and returns the corresponding authentication object, and an error if there is any.
func (c *FakeAuthentications) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.Authentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(authenticationsResource, name), &operatorv1.Authentication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Authentication), err
}

// List takes label and field selectors, and returns the list of Authentications that match those selectors.
func (c *FakeAuthentications) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.AuthenticationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(authenticationsResource, authenticationsKind, opts), &operatorv1.AuthenticationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.AuthenticationList{ListMeta: obj.(*operatorv1.AuthenticationList).ListMeta}
	for _, item := range obj.(*operatorv1.AuthenticationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested authentications.
func (c *FakeAuthentications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(authenticationsResource, opts))
}

// Create takes the representation of a authentication and creates it.  Returns the server's representation of the authentication, and an error, if there is any.
func (c *FakeAuthentications) Create(ctx context.Context, authentication *operatorv1.Authentication, opts v1.CreateOptions) (result *operatorv1.Authentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(authenticationsResource, authentication), &operatorv1.Authentication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Authentication), err
}

// Update takes the representation of a authentication and updates it. Returns the server's representation of the authentication, and an error, if there is any.
func (c *FakeAuthentications) Update(ctx context.Context, authentication *operatorv1.Authentication, opts v1.UpdateOptions) (result *operatorv1.Authentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(authenticationsResource, authentication), &operatorv1.Authentication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Authentication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAuthentications) UpdateStatus(ctx context.Context, authentication *operatorv1.Authentication, opts v1.UpdateOptions) (*operatorv1.Authentication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(authenticationsResource, "status", authentication), &operatorv1.Authentication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Authentication), err
}

// Delete takes name of the authentication and deletes it. Returns an error if one occurs.
func (c *FakeAuthentications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(authenticationsResource, name, opts), &operatorv1.Authentication{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAuthentications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(authenticationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.AuthenticationList{})
	return err
}

// Patch applies the patch and returns the patched authentication.
func (c *FakeAuthentications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.Authentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(authenticationsResource, name, pt, data, subresources...), &operatorv1.Authentication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Authentication), err
}
