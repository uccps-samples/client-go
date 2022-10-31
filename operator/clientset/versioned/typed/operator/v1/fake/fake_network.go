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

// FakeNetworks implements NetworkInterface
type FakeNetworks struct {
	Fake *FakeOperatorV1
}

var networksResource = schema.GroupVersionResource{Group: "operator.uccp.io", Version: "v1", Resource: "networks"}

var networksKind = schema.GroupVersionKind{Group: "operator.uccp.io", Version: "v1", Kind: "Network"}

// Get takes name of the network, and returns the corresponding network object, and an error if there is any.
func (c *FakeNetworks) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(networksResource, name), &operatorv1.Network{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Network), err
}

// List takes label and field selectors, and returns the list of Networks that match those selectors.
func (c *FakeNetworks) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.NetworkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(networksResource, networksKind, opts), &operatorv1.NetworkList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.NetworkList{ListMeta: obj.(*operatorv1.NetworkList).ListMeta}
	for _, item := range obj.(*operatorv1.NetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networks.
func (c *FakeNetworks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(networksResource, opts))
}

// Create takes the representation of a network and creates it.  Returns the server's representation of the network, and an error, if there is any.
func (c *FakeNetworks) Create(ctx context.Context, network *operatorv1.Network, opts v1.CreateOptions) (result *operatorv1.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(networksResource, network), &operatorv1.Network{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Network), err
}

// Update takes the representation of a network and updates it. Returns the server's representation of the network, and an error, if there is any.
func (c *FakeNetworks) Update(ctx context.Context, network *operatorv1.Network, opts v1.UpdateOptions) (result *operatorv1.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(networksResource, network), &operatorv1.Network{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Network), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworks) UpdateStatus(ctx context.Context, network *operatorv1.Network, opts v1.UpdateOptions) (*operatorv1.Network, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(networksResource, "status", network), &operatorv1.Network{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Network), err
}

// Delete takes name of the network and deletes it. Returns an error if one occurs.
func (c *FakeNetworks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(networksResource, name), &operatorv1.Network{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(networksResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.NetworkList{})
	return err
}

// Patch applies the patch and returns the patched network.
func (c *FakeNetworks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(networksResource, name, pt, data, subresources...), &operatorv1.Network{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Network), err
}
