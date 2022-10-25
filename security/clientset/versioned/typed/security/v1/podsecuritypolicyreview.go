// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/uccps-samples/api/security/v1"
	scheme "github.com/uccps-samples/client-go/security/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// PodSecurityPolicyReviewsGetter has a method to return a PodSecurityPolicyReviewInterface.
// A group's client should implement this interface.
type PodSecurityPolicyReviewsGetter interface {
	PodSecurityPolicyReviews(namespace string) PodSecurityPolicyReviewInterface
}

// PodSecurityPolicyReviewInterface has methods to work with PodSecurityPolicyReview resources.
type PodSecurityPolicyReviewInterface interface {
	Create(ctx context.Context, podSecurityPolicyReview *v1.PodSecurityPolicyReview, opts metav1.CreateOptions) (*v1.PodSecurityPolicyReview, error)
	PodSecurityPolicyReviewExpansion
}

// podSecurityPolicyReviews implements PodSecurityPolicyReviewInterface
type podSecurityPolicyReviews struct {
	client rest.Interface
	ns     string
}

// newPodSecurityPolicyReviews returns a PodSecurityPolicyReviews
func newPodSecurityPolicyReviews(c *SecurityV1Client, namespace string) *podSecurityPolicyReviews {
	return &podSecurityPolicyReviews{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a podSecurityPolicyReview and creates it.  Returns the server's representation of the podSecurityPolicyReview, and an error, if there is any.
func (c *podSecurityPolicyReviews) Create(ctx context.Context, podSecurityPolicyReview *v1.PodSecurityPolicyReview, opts metav1.CreateOptions) (result *v1.PodSecurityPolicyReview, err error) {
	result = &v1.PodSecurityPolicyReview{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("podsecuritypolicyreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podSecurityPolicyReview).
		Do(ctx).
		Into(result)
	return
}
