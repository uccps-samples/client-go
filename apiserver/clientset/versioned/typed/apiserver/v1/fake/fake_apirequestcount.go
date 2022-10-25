// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	apiserverv1 "github.com/uccps-samples/api/apiserver/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAPIRequestCounts implements APIRequestCountInterface
type FakeAPIRequestCounts struct {
	Fake *FakeApiserverV1
}

var apirequestcountsResource = schema.GroupVersionResource{Group: "apiserver.uccp.io", Version: "v1", Resource: "apirequestcounts"}

var apirequestcountsKind = schema.GroupVersionKind{Group: "apiserver.uccp.io", Version: "v1", Kind: "APIRequestCount"}

// Get takes name of the aPIRequestCount, and returns the corresponding aPIRequestCount object, and an error if there is any.
func (c *FakeAPIRequestCounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *apiserverv1.APIRequestCount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(apirequestcountsResource, name), &apiserverv1.APIRequestCount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiserverv1.APIRequestCount), err
}

// List takes label and field selectors, and returns the list of APIRequestCounts that match those selectors.
func (c *FakeAPIRequestCounts) List(ctx context.Context, opts v1.ListOptions) (result *apiserverv1.APIRequestCountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(apirequestcountsResource, apirequestcountsKind, opts), &apiserverv1.APIRequestCountList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apiserverv1.APIRequestCountList{ListMeta: obj.(*apiserverv1.APIRequestCountList).ListMeta}
	for _, item := range obj.(*apiserverv1.APIRequestCountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPIRequestCounts.
func (c *FakeAPIRequestCounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(apirequestcountsResource, opts))
}

// Create takes the representation of a aPIRequestCount and creates it.  Returns the server's representation of the aPIRequestCount, and an error, if there is any.
func (c *FakeAPIRequestCounts) Create(ctx context.Context, aPIRequestCount *apiserverv1.APIRequestCount, opts v1.CreateOptions) (result *apiserverv1.APIRequestCount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(apirequestcountsResource, aPIRequestCount), &apiserverv1.APIRequestCount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiserverv1.APIRequestCount), err
}

// Update takes the representation of a aPIRequestCount and updates it. Returns the server's representation of the aPIRequestCount, and an error, if there is any.
func (c *FakeAPIRequestCounts) Update(ctx context.Context, aPIRequestCount *apiserverv1.APIRequestCount, opts v1.UpdateOptions) (result *apiserverv1.APIRequestCount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(apirequestcountsResource, aPIRequestCount), &apiserverv1.APIRequestCount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiserverv1.APIRequestCount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAPIRequestCounts) UpdateStatus(ctx context.Context, aPIRequestCount *apiserverv1.APIRequestCount, opts v1.UpdateOptions) (*apiserverv1.APIRequestCount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(apirequestcountsResource, "status", aPIRequestCount), &apiserverv1.APIRequestCount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiserverv1.APIRequestCount), err
}

// Delete takes name of the aPIRequestCount and deletes it. Returns an error if one occurs.
func (c *FakeAPIRequestCounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(apirequestcountsResource, name), &apiserverv1.APIRequestCount{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPIRequestCounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(apirequestcountsResource, listOpts)

	_, err := c.Fake.Invokes(action, &apiserverv1.APIRequestCountList{})
	return err
}

// Patch applies the patch and returns the patched aPIRequestCount.
func (c *FakeAPIRequestCounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apiserverv1.APIRequestCount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(apirequestcountsResource, name, pt, data, subresources...), &apiserverv1.APIRequestCount{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiserverv1.APIRequestCount), err
}
