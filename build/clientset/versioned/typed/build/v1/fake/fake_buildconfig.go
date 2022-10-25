// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	buildv1 "github.com/uccps-samples/api/build/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBuildConfigs implements BuildConfigInterface
type FakeBuildConfigs struct {
	Fake *FakeBuildV1
	ns   string
}

var buildconfigsResource = schema.GroupVersionResource{Group: "build.uccp.io", Version: "v1", Resource: "buildconfigs"}

var buildconfigsKind = schema.GroupVersionKind{Group: "build.uccp.io", Version: "v1", Kind: "BuildConfig"}

// Get takes name of the buildConfig, and returns the corresponding buildConfig object, and an error if there is any.
func (c *FakeBuildConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *buildv1.BuildConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(buildconfigsResource, c.ns, name), &buildv1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*buildv1.BuildConfig), err
}

// List takes label and field selectors, and returns the list of BuildConfigs that match those selectors.
func (c *FakeBuildConfigs) List(ctx context.Context, opts v1.ListOptions) (result *buildv1.BuildConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(buildconfigsResource, buildconfigsKind, c.ns, opts), &buildv1.BuildConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &buildv1.BuildConfigList{ListMeta: obj.(*buildv1.BuildConfigList).ListMeta}
	for _, item := range obj.(*buildv1.BuildConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested buildConfigs.
func (c *FakeBuildConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(buildconfigsResource, c.ns, opts))

}

// Create takes the representation of a buildConfig and creates it.  Returns the server's representation of the buildConfig, and an error, if there is any.
func (c *FakeBuildConfigs) Create(ctx context.Context, buildConfig *buildv1.BuildConfig, opts v1.CreateOptions) (result *buildv1.BuildConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(buildconfigsResource, c.ns, buildConfig), &buildv1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*buildv1.BuildConfig), err
}

// Update takes the representation of a buildConfig and updates it. Returns the server's representation of the buildConfig, and an error, if there is any.
func (c *FakeBuildConfigs) Update(ctx context.Context, buildConfig *buildv1.BuildConfig, opts v1.UpdateOptions) (result *buildv1.BuildConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(buildconfigsResource, c.ns, buildConfig), &buildv1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*buildv1.BuildConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBuildConfigs) UpdateStatus(ctx context.Context, buildConfig *buildv1.BuildConfig, opts v1.UpdateOptions) (*buildv1.BuildConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(buildconfigsResource, "status", c.ns, buildConfig), &buildv1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*buildv1.BuildConfig), err
}

// Delete takes name of the buildConfig and deletes it. Returns an error if one occurs.
func (c *FakeBuildConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(buildconfigsResource, c.ns, name), &buildv1.BuildConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuildConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(buildconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &buildv1.BuildConfigList{})
	return err
}

// Patch applies the patch and returns the patched buildConfig.
func (c *FakeBuildConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *buildv1.BuildConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildconfigsResource, c.ns, name, pt, data, subresources...), &buildv1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*buildv1.BuildConfig), err
}

// Instantiate takes the representation of a buildRequest and creates it.  Returns the server's representation of the build, and an error, if there is any.
func (c *FakeBuildConfigs) Instantiate(ctx context.Context, buildConfigName string, buildRequest *buildv1.BuildRequest, opts v1.CreateOptions) (result *buildv1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateSubresourceAction(buildconfigsResource, buildConfigName, "instantiate", c.ns, buildRequest), &buildv1.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*buildv1.Build), err
}
