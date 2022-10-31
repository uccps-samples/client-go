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

// FakeServiceCAs implements ServiceCAInterface
type FakeServiceCAs struct {
	Fake *FakeOperatorV1
}

var servicecasResource = schema.GroupVersionResource{Group: "operator.uccp.io", Version: "v1", Resource: "servicecas"}

var servicecasKind = schema.GroupVersionKind{Group: "operator.uccp.io", Version: "v1", Kind: "ServiceCA"}

// Get takes name of the serviceCA, and returns the corresponding serviceCA object, and an error if there is any.
func (c *FakeServiceCAs) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.ServiceCA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(servicecasResource, name), &operatorv1.ServiceCA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCA), err
}

// List takes label and field selectors, and returns the list of ServiceCAs that match those selectors.
func (c *FakeServiceCAs) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.ServiceCAList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(servicecasResource, servicecasKind, opts), &operatorv1.ServiceCAList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.ServiceCAList{ListMeta: obj.(*operatorv1.ServiceCAList).ListMeta}
	for _, item := range obj.(*operatorv1.ServiceCAList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceCAs.
func (c *FakeServiceCAs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(servicecasResource, opts))
}

// Create takes the representation of a serviceCA and creates it.  Returns the server's representation of the serviceCA, and an error, if there is any.
func (c *FakeServiceCAs) Create(ctx context.Context, serviceCA *operatorv1.ServiceCA, opts v1.CreateOptions) (result *operatorv1.ServiceCA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(servicecasResource, serviceCA), &operatorv1.ServiceCA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCA), err
}

// Update takes the representation of a serviceCA and updates it. Returns the server's representation of the serviceCA, and an error, if there is any.
func (c *FakeServiceCAs) Update(ctx context.Context, serviceCA *operatorv1.ServiceCA, opts v1.UpdateOptions) (result *operatorv1.ServiceCA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(servicecasResource, serviceCA), &operatorv1.ServiceCA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCA), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceCAs) UpdateStatus(ctx context.Context, serviceCA *operatorv1.ServiceCA, opts v1.UpdateOptions) (*operatorv1.ServiceCA, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(servicecasResource, "status", serviceCA), &operatorv1.ServiceCA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCA), err
}

// Delete takes name of the serviceCA and deletes it. Returns an error if one occurs.
func (c *FakeServiceCAs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(servicecasResource, name), &operatorv1.ServiceCA{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceCAs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(servicecasResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.ServiceCAList{})
	return err
}

// Patch applies the patch and returns the patched serviceCA.
func (c *FakeServiceCAs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.ServiceCA, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(servicecasResource, name, pt, data, subresources...), &operatorv1.ServiceCA{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCA), err
}
