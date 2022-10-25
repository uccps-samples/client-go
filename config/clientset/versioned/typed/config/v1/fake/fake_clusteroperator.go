// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	configv1 "github.com/uccps-samples/api/config/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterOperators implements ClusterOperatorInterface
type FakeClusterOperators struct {
	Fake *FakeConfigV1
}

var clusteroperatorsResource = schema.GroupVersionResource{Group: "config.uccp.io", Version: "v1", Resource: "clusteroperators"}

var clusteroperatorsKind = schema.GroupVersionKind{Group: "config.uccp.io", Version: "v1", Kind: "ClusterOperator"}

// Get takes name of the clusterOperator, and returns the corresponding clusterOperator object, and an error if there is any.
func (c *FakeClusterOperators) Get(ctx context.Context, name string, options v1.GetOptions) (result *configv1.ClusterOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusteroperatorsResource, name), &configv1.ClusterOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.ClusterOperator), err
}

// List takes label and field selectors, and returns the list of ClusterOperators that match those selectors.
func (c *FakeClusterOperators) List(ctx context.Context, opts v1.ListOptions) (result *configv1.ClusterOperatorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusteroperatorsResource, clusteroperatorsKind, opts), &configv1.ClusterOperatorList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &configv1.ClusterOperatorList{ListMeta: obj.(*configv1.ClusterOperatorList).ListMeta}
	for _, item := range obj.(*configv1.ClusterOperatorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterOperators.
func (c *FakeClusterOperators) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusteroperatorsResource, opts))
}

// Create takes the representation of a clusterOperator and creates it.  Returns the server's representation of the clusterOperator, and an error, if there is any.
func (c *FakeClusterOperators) Create(ctx context.Context, clusterOperator *configv1.ClusterOperator, opts v1.CreateOptions) (result *configv1.ClusterOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusteroperatorsResource, clusterOperator), &configv1.ClusterOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.ClusterOperator), err
}

// Update takes the representation of a clusterOperator and updates it. Returns the server's representation of the clusterOperator, and an error, if there is any.
func (c *FakeClusterOperators) Update(ctx context.Context, clusterOperator *configv1.ClusterOperator, opts v1.UpdateOptions) (result *configv1.ClusterOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusteroperatorsResource, clusterOperator), &configv1.ClusterOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.ClusterOperator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterOperators) UpdateStatus(ctx context.Context, clusterOperator *configv1.ClusterOperator, opts v1.UpdateOptions) (*configv1.ClusterOperator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusteroperatorsResource, "status", clusterOperator), &configv1.ClusterOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.ClusterOperator), err
}

// Delete takes name of the clusterOperator and deletes it. Returns an error if one occurs.
func (c *FakeClusterOperators) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusteroperatorsResource, name), &configv1.ClusterOperator{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterOperators) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusteroperatorsResource, listOpts)

	_, err := c.Fake.Invokes(action, &configv1.ClusterOperatorList{})
	return err
}

// Patch applies the patch and returns the patched clusterOperator.
func (c *FakeClusterOperators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *configv1.ClusterOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusteroperatorsResource, name, pt, data, subresources...), &configv1.ClusterOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.ClusterOperator), err
}
