// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/uccps-samples/api/servicecertsigner/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceCertSignerOperatorConfigs implements ServiceCertSignerOperatorConfigInterface
type FakeServiceCertSignerOperatorConfigs struct {
	Fake *FakeServicecertsignerV1alpha1
}

var servicecertsigneroperatorconfigsResource = schema.GroupVersionResource{Group: "servicecertsigner.config.uccp.io", Version: "v1alpha1", Resource: "servicecertsigneroperatorconfigs"}

var servicecertsigneroperatorconfigsKind = schema.GroupVersionKind{Group: "servicecertsigner.config.uccp.io", Version: "v1alpha1", Kind: "ServiceCertSignerOperatorConfig"}

// Get takes name of the serviceCertSignerOperatorConfig, and returns the corresponding serviceCertSignerOperatorConfig object, and an error if there is any.
func (c *FakeServiceCertSignerOperatorConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceCertSignerOperatorConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(servicecertsigneroperatorconfigsResource, name), &v1alpha1.ServiceCertSignerOperatorConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCertSignerOperatorConfig), err
}

// List takes label and field selectors, and returns the list of ServiceCertSignerOperatorConfigs that match those selectors.
func (c *FakeServiceCertSignerOperatorConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceCertSignerOperatorConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(servicecertsigneroperatorconfigsResource, servicecertsigneroperatorconfigsKind, opts), &v1alpha1.ServiceCertSignerOperatorConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceCertSignerOperatorConfigList{ListMeta: obj.(*v1alpha1.ServiceCertSignerOperatorConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceCertSignerOperatorConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceCertSignerOperatorConfigs.
func (c *FakeServiceCertSignerOperatorConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(servicecertsigneroperatorconfigsResource, opts))
}

// Create takes the representation of a serviceCertSignerOperatorConfig and creates it.  Returns the server's representation of the serviceCertSignerOperatorConfig, and an error, if there is any.
func (c *FakeServiceCertSignerOperatorConfigs) Create(ctx context.Context, serviceCertSignerOperatorConfig *v1alpha1.ServiceCertSignerOperatorConfig, opts v1.CreateOptions) (result *v1alpha1.ServiceCertSignerOperatorConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(servicecertsigneroperatorconfigsResource, serviceCertSignerOperatorConfig), &v1alpha1.ServiceCertSignerOperatorConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCertSignerOperatorConfig), err
}

// Update takes the representation of a serviceCertSignerOperatorConfig and updates it. Returns the server's representation of the serviceCertSignerOperatorConfig, and an error, if there is any.
func (c *FakeServiceCertSignerOperatorConfigs) Update(ctx context.Context, serviceCertSignerOperatorConfig *v1alpha1.ServiceCertSignerOperatorConfig, opts v1.UpdateOptions) (result *v1alpha1.ServiceCertSignerOperatorConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(servicecertsigneroperatorconfigsResource, serviceCertSignerOperatorConfig), &v1alpha1.ServiceCertSignerOperatorConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCertSignerOperatorConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceCertSignerOperatorConfigs) UpdateStatus(ctx context.Context, serviceCertSignerOperatorConfig *v1alpha1.ServiceCertSignerOperatorConfig, opts v1.UpdateOptions) (*v1alpha1.ServiceCertSignerOperatorConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(servicecertsigneroperatorconfigsResource, "status", serviceCertSignerOperatorConfig), &v1alpha1.ServiceCertSignerOperatorConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCertSignerOperatorConfig), err
}

// Delete takes name of the serviceCertSignerOperatorConfig and deletes it. Returns an error if one occurs.
func (c *FakeServiceCertSignerOperatorConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(servicecertsigneroperatorconfigsResource, name), &v1alpha1.ServiceCertSignerOperatorConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceCertSignerOperatorConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(servicecertsigneroperatorconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceCertSignerOperatorConfigList{})
	return err
}

// Patch applies the patch and returns the patched serviceCertSignerOperatorConfig.
func (c *FakeServiceCertSignerOperatorConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceCertSignerOperatorConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(servicecertsigneroperatorconfigsResource, name, pt, data, subresources...), &v1alpha1.ServiceCertSignerOperatorConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCertSignerOperatorConfig), err
}
