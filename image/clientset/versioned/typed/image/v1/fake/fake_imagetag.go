// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	imagev1 "github.com/uccps-samples/api/image/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeImageTags implements ImageTagInterface
type FakeImageTags struct {
	Fake *FakeImageV1
	ns   string
}

var imagetagsResource = schema.GroupVersionResource{Group: "image.uccp.io", Version: "v1", Resource: "imagetags"}

var imagetagsKind = schema.GroupVersionKind{Group: "image.uccp.io", Version: "v1", Kind: "ImageTag"}

// Get takes name of the imageTag, and returns the corresponding imageTag object, and an error if there is any.
func (c *FakeImageTags) Get(ctx context.Context, name string, options v1.GetOptions) (result *imagev1.ImageTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(imagetagsResource, c.ns, name), &imagev1.ImageTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*imagev1.ImageTag), err
}

// List takes label and field selectors, and returns the list of ImageTags that match those selectors.
func (c *FakeImageTags) List(ctx context.Context, opts v1.ListOptions) (result *imagev1.ImageTagList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(imagetagsResource, imagetagsKind, c.ns, opts), &imagev1.ImageTagList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &imagev1.ImageTagList{ListMeta: obj.(*imagev1.ImageTagList).ListMeta}
	for _, item := range obj.(*imagev1.ImageTagList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Create takes the representation of a imageTag and creates it.  Returns the server's representation of the imageTag, and an error, if there is any.
func (c *FakeImageTags) Create(ctx context.Context, imageTag *imagev1.ImageTag, opts v1.CreateOptions) (result *imagev1.ImageTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(imagetagsResource, c.ns, imageTag), &imagev1.ImageTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*imagev1.ImageTag), err
}

// Update takes the representation of a imageTag and updates it. Returns the server's representation of the imageTag, and an error, if there is any.
func (c *FakeImageTags) Update(ctx context.Context, imageTag *imagev1.ImageTag, opts v1.UpdateOptions) (result *imagev1.ImageTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(imagetagsResource, c.ns, imageTag), &imagev1.ImageTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*imagev1.ImageTag), err
}

// Delete takes name of the imageTag and deletes it. Returns an error if one occurs.
func (c *FakeImageTags) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(imagetagsResource, c.ns, name), &imagev1.ImageTag{})

	return err
}
