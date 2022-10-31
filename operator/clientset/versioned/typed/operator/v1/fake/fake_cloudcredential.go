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

// FakeCloudCredentials implements CloudCredentialInterface
type FakeCloudCredentials struct {
	Fake *FakeOperatorV1
}

var cloudcredentialsResource = schema.GroupVersionResource{Group: "operator.uccp.io", Version: "v1", Resource: "cloudcredentials"}

var cloudcredentialsKind = schema.GroupVersionKind{Group: "operator.uccp.io", Version: "v1", Kind: "CloudCredential"}

// Get takes name of the cloudCredential, and returns the corresponding cloudCredential object, and an error if there is any.
func (c *FakeCloudCredentials) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.CloudCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cloudcredentialsResource, name), &operatorv1.CloudCredential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.CloudCredential), err
}

// List takes label and field selectors, and returns the list of CloudCredentials that match those selectors.
func (c *FakeCloudCredentials) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.CloudCredentialList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cloudcredentialsResource, cloudcredentialsKind, opts), &operatorv1.CloudCredentialList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.CloudCredentialList{ListMeta: obj.(*operatorv1.CloudCredentialList).ListMeta}
	for _, item := range obj.(*operatorv1.CloudCredentialList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudCredentials.
func (c *FakeCloudCredentials) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cloudcredentialsResource, opts))
}

// Create takes the representation of a cloudCredential and creates it.  Returns the server's representation of the cloudCredential, and an error, if there is any.
func (c *FakeCloudCredentials) Create(ctx context.Context, cloudCredential *operatorv1.CloudCredential, opts v1.CreateOptions) (result *operatorv1.CloudCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cloudcredentialsResource, cloudCredential), &operatorv1.CloudCredential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.CloudCredential), err
}

// Update takes the representation of a cloudCredential and updates it. Returns the server's representation of the cloudCredential, and an error, if there is any.
func (c *FakeCloudCredentials) Update(ctx context.Context, cloudCredential *operatorv1.CloudCredential, opts v1.UpdateOptions) (result *operatorv1.CloudCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cloudcredentialsResource, cloudCredential), &operatorv1.CloudCredential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.CloudCredential), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudCredentials) UpdateStatus(ctx context.Context, cloudCredential *operatorv1.CloudCredential, opts v1.UpdateOptions) (*operatorv1.CloudCredential, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(cloudcredentialsResource, "status", cloudCredential), &operatorv1.CloudCredential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.CloudCredential), err
}

// Delete takes name of the cloudCredential and deletes it. Returns an error if one occurs.
func (c *FakeCloudCredentials) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(cloudcredentialsResource, name), &operatorv1.CloudCredential{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudCredentials) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cloudcredentialsResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.CloudCredentialList{})
	return err
}

// Patch applies the patch and returns the patched cloudCredential.
func (c *FakeCloudCredentials) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.CloudCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cloudcredentialsResource, name, pt, data, subresources...), &operatorv1.CloudCredential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.CloudCredential), err
}
