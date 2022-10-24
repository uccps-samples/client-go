// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/uccps-samples/api/config/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProxyLister helps list Proxies.
// All objects returned here must be treated as read-only.
type ProxyLister interface {
	// List lists all Proxies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Proxy, err error)
	// Get retrieves the Proxy from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Proxy, error)
	ProxyListerExpansion
}

// proxyLister implements the ProxyLister interface.
type proxyLister struct {
	indexer cache.Indexer
}

// NewProxyLister returns a new ProxyLister.
func NewProxyLister(indexer cache.Indexer) ProxyLister {
	return &proxyLister{indexer: indexer}
}

// List lists all Proxies in the indexer.
func (s *proxyLister) List(selector labels.Selector) (ret []*v1.Proxy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Proxy))
	})
	return ret, err
}

// Get retrieves the Proxy from the index for a given name.
func (s *proxyLister) Get(name string) (*v1.Proxy, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("proxy"), name)
	}
	return obj.(*v1.Proxy), nil
}
