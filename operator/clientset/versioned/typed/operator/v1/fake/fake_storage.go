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

// FakeStorages implements StorageInterface
type FakeStorages struct {
	Fake *FakeOperatorV1
}

var storagesResource = schema.GroupVersionResource{Group: "operator.uccp.io", Version: "v1", Resource: "storages"}

var storagesKind = schema.GroupVersionKind{Group: "operator.uccp.io", Version: "v1", Kind: "Storage"}

// Get takes name of the storage, and returns the corresponding storage object, and an error if there is any.
func (c *FakeStorages) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.Storage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(storagesResource, name), &operatorv1.Storage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Storage), err
}

// List takes label and field selectors, and returns the list of Storages that match those selectors.
func (c *FakeStorages) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.StorageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(storagesResource, storagesKind, opts), &operatorv1.StorageList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.StorageList{ListMeta: obj.(*operatorv1.StorageList).ListMeta}
	for _, item := range obj.(*operatorv1.StorageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storages.
func (c *FakeStorages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(storagesResource, opts))
}

// Create takes the representation of a storage and creates it.  Returns the server's representation of the storage, and an error, if there is any.
func (c *FakeStorages) Create(ctx context.Context, storage *operatorv1.Storage, opts v1.CreateOptions) (result *operatorv1.Storage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(storagesResource, storage), &operatorv1.Storage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Storage), err
}

// Update takes the representation of a storage and updates it. Returns the server's representation of the storage, and an error, if there is any.
func (c *FakeStorages) Update(ctx context.Context, storage *operatorv1.Storage, opts v1.UpdateOptions) (result *operatorv1.Storage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(storagesResource, storage), &operatorv1.Storage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Storage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorages) UpdateStatus(ctx context.Context, storage *operatorv1.Storage, opts v1.UpdateOptions) (*operatorv1.Storage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(storagesResource, "status", storage), &operatorv1.Storage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Storage), err
}

// Delete takes name of the storage and deletes it. Returns an error if one occurs.
func (c *FakeStorages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(storagesResource, name, opts), &operatorv1.Storage{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(storagesResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.StorageList{})
	return err
}

// Patch applies the patch and returns the patched storage.
func (c *FakeStorages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.Storage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(storagesResource, name, pt, data, subresources...), &operatorv1.Storage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.Storage), err
}
