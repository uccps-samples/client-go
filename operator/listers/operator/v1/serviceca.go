// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/uccps-samples/api/operator/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceCALister helps list ServiceCAs.
// All objects returned here must be treated as read-only.
type ServiceCALister interface {
	// List lists all ServiceCAs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ServiceCA, err error)
	// Get retrieves the ServiceCA from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ServiceCA, error)
	ServiceCAListerExpansion
}

// serviceCALister implements the ServiceCALister interface.
type serviceCALister struct {
	indexer cache.Indexer
}

// NewServiceCALister returns a new ServiceCALister.
func NewServiceCALister(indexer cache.Indexer) ServiceCALister {
	return &serviceCALister{indexer: indexer}
}

// List lists all ServiceCAs in the indexer.
func (s *serviceCALister) List(selector labels.Selector) (ret []*v1.ServiceCA, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ServiceCA))
	})
	return ret, err
}

// Get retrieves the ServiceCA from the index for a given name.
func (s *serviceCALister) Get(name string) (*v1.ServiceCA, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("serviceca"), name)
	}
	return obj.(*v1.ServiceCA), nil
}
