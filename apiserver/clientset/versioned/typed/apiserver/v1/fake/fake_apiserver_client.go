// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/uccps-samples/client-go/apiserver/clientset/versioned/typed/apiserver/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeApiserverV1 struct {
	*testing.Fake
}

func (c *FakeApiserverV1) APIRequestCounts() v1.APIRequestCountInterface {
	return &FakeAPIRequestCounts{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeApiserverV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
