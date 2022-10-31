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

// FakeServiceCatalogControllerManagers implements ServiceCatalogControllerManagerInterface
type FakeServiceCatalogControllerManagers struct {
	Fake *FakeOperatorV1
}

var servicecatalogcontrollermanagersResource = schema.GroupVersionResource{Group: "operator.uccp.io", Version: "v1", Resource: "servicecatalogcontrollermanagers"}

var servicecatalogcontrollermanagersKind = schema.GroupVersionKind{Group: "operator.uccp.io", Version: "v1", Kind: "ServiceCatalogControllerManager"}

// Get takes name of the serviceCatalogControllerManager, and returns the corresponding serviceCatalogControllerManager object, and an error if there is any.
func (c *FakeServiceCatalogControllerManagers) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.ServiceCatalogControllerManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(servicecatalogcontrollermanagersResource, name), &operatorv1.ServiceCatalogControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCatalogControllerManager), err
}

// List takes label and field selectors, and returns the list of ServiceCatalogControllerManagers that match those selectors.
func (c *FakeServiceCatalogControllerManagers) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.ServiceCatalogControllerManagerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(servicecatalogcontrollermanagersResource, servicecatalogcontrollermanagersKind, opts), &operatorv1.ServiceCatalogControllerManagerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.ServiceCatalogControllerManagerList{ListMeta: obj.(*operatorv1.ServiceCatalogControllerManagerList).ListMeta}
	for _, item := range obj.(*operatorv1.ServiceCatalogControllerManagerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceCatalogControllerManagers.
func (c *FakeServiceCatalogControllerManagers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(servicecatalogcontrollermanagersResource, opts))
}

// Create takes the representation of a serviceCatalogControllerManager and creates it.  Returns the server's representation of the serviceCatalogControllerManager, and an error, if there is any.
func (c *FakeServiceCatalogControllerManagers) Create(ctx context.Context, serviceCatalogControllerManager *operatorv1.ServiceCatalogControllerManager, opts v1.CreateOptions) (result *operatorv1.ServiceCatalogControllerManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(servicecatalogcontrollermanagersResource, serviceCatalogControllerManager), &operatorv1.ServiceCatalogControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCatalogControllerManager), err
}

// Update takes the representation of a serviceCatalogControllerManager and updates it. Returns the server's representation of the serviceCatalogControllerManager, and an error, if there is any.
func (c *FakeServiceCatalogControllerManagers) Update(ctx context.Context, serviceCatalogControllerManager *operatorv1.ServiceCatalogControllerManager, opts v1.UpdateOptions) (result *operatorv1.ServiceCatalogControllerManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(servicecatalogcontrollermanagersResource, serviceCatalogControllerManager), &operatorv1.ServiceCatalogControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCatalogControllerManager), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceCatalogControllerManagers) UpdateStatus(ctx context.Context, serviceCatalogControllerManager *operatorv1.ServiceCatalogControllerManager, opts v1.UpdateOptions) (*operatorv1.ServiceCatalogControllerManager, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(servicecatalogcontrollermanagersResource, "status", serviceCatalogControllerManager), &operatorv1.ServiceCatalogControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCatalogControllerManager), err
}

// Delete takes name of the serviceCatalogControllerManager and deletes it. Returns an error if one occurs.
func (c *FakeServiceCatalogControllerManagers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(servicecatalogcontrollermanagersResource, name), &operatorv1.ServiceCatalogControllerManager{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceCatalogControllerManagers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(servicecatalogcontrollermanagersResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.ServiceCatalogControllerManagerList{})
	return err
}

// Patch applies the patch and returns the patched serviceCatalogControllerManager.
func (c *FakeServiceCatalogControllerManagers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.ServiceCatalogControllerManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(servicecatalogcontrollermanagersResource, name, pt, data, subresources...), &operatorv1.ServiceCatalogControllerManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCatalogControllerManager), err
}
