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

// FakeServiceCatalogAPIServers implements ServiceCatalogAPIServerInterface
type FakeServiceCatalogAPIServers struct {
	Fake *FakeOperatorV1
}

var servicecatalogapiserversResource = schema.GroupVersionResource{Group: "operator.uccp.io", Version: "v1", Resource: "servicecatalogapiservers"}

var servicecatalogapiserversKind = schema.GroupVersionKind{Group: "operator.uccp.io", Version: "v1", Kind: "ServiceCatalogAPIServer"}

// Get takes name of the serviceCatalogAPIServer, and returns the corresponding serviceCatalogAPIServer object, and an error if there is any.
func (c *FakeServiceCatalogAPIServers) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.ServiceCatalogAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(servicecatalogapiserversResource, name), &operatorv1.ServiceCatalogAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCatalogAPIServer), err
}

// List takes label and field selectors, and returns the list of ServiceCatalogAPIServers that match those selectors.
func (c *FakeServiceCatalogAPIServers) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.ServiceCatalogAPIServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(servicecatalogapiserversResource, servicecatalogapiserversKind, opts), &operatorv1.ServiceCatalogAPIServerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.ServiceCatalogAPIServerList{ListMeta: obj.(*operatorv1.ServiceCatalogAPIServerList).ListMeta}
	for _, item := range obj.(*operatorv1.ServiceCatalogAPIServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceCatalogAPIServers.
func (c *FakeServiceCatalogAPIServers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(servicecatalogapiserversResource, opts))
}

// Create takes the representation of a serviceCatalogAPIServer and creates it.  Returns the server's representation of the serviceCatalogAPIServer, and an error, if there is any.
func (c *FakeServiceCatalogAPIServers) Create(ctx context.Context, serviceCatalogAPIServer *operatorv1.ServiceCatalogAPIServer, opts v1.CreateOptions) (result *operatorv1.ServiceCatalogAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(servicecatalogapiserversResource, serviceCatalogAPIServer), &operatorv1.ServiceCatalogAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCatalogAPIServer), err
}

// Update takes the representation of a serviceCatalogAPIServer and updates it. Returns the server's representation of the serviceCatalogAPIServer, and an error, if there is any.
func (c *FakeServiceCatalogAPIServers) Update(ctx context.Context, serviceCatalogAPIServer *operatorv1.ServiceCatalogAPIServer, opts v1.UpdateOptions) (result *operatorv1.ServiceCatalogAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(servicecatalogapiserversResource, serviceCatalogAPIServer), &operatorv1.ServiceCatalogAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCatalogAPIServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceCatalogAPIServers) UpdateStatus(ctx context.Context, serviceCatalogAPIServer *operatorv1.ServiceCatalogAPIServer, opts v1.UpdateOptions) (*operatorv1.ServiceCatalogAPIServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(servicecatalogapiserversResource, "status", serviceCatalogAPIServer), &operatorv1.ServiceCatalogAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCatalogAPIServer), err
}

// Delete takes name of the serviceCatalogAPIServer and deletes it. Returns an error if one occurs.
func (c *FakeServiceCatalogAPIServers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(servicecatalogapiserversResource, name, opts), &operatorv1.ServiceCatalogAPIServer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceCatalogAPIServers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(servicecatalogapiserversResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.ServiceCatalogAPIServerList{})
	return err
}

// Patch applies the patch and returns the patched serviceCatalogAPIServer.
func (c *FakeServiceCatalogAPIServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.ServiceCatalogAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(servicecatalogapiserversResource, name, pt, data, subresources...), &operatorv1.ServiceCatalogAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ServiceCatalogAPIServer), err
}
