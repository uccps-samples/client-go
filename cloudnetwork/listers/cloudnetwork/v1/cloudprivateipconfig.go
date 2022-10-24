// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/uccps-samples/api/cloudnetwork/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CloudPrivateIPConfigLister helps list CloudPrivateIPConfigs.
// All objects returned here must be treated as read-only.
type CloudPrivateIPConfigLister interface {
	// List lists all CloudPrivateIPConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CloudPrivateIPConfig, err error)
	// Get retrieves the CloudPrivateIPConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.CloudPrivateIPConfig, error)
	CloudPrivateIPConfigListerExpansion
}

// cloudPrivateIPConfigLister implements the CloudPrivateIPConfigLister interface.
type cloudPrivateIPConfigLister struct {
	indexer cache.Indexer
}

// NewCloudPrivateIPConfigLister returns a new CloudPrivateIPConfigLister.
func NewCloudPrivateIPConfigLister(indexer cache.Indexer) CloudPrivateIPConfigLister {
	return &cloudPrivateIPConfigLister{indexer: indexer}
}

// List lists all CloudPrivateIPConfigs in the indexer.
func (s *cloudPrivateIPConfigLister) List(selector labels.Selector) (ret []*v1.CloudPrivateIPConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CloudPrivateIPConfig))
	})
	return ret, err
}

// Get retrieves the CloudPrivateIPConfig from the index for a given name.
func (s *cloudPrivateIPConfigLister) Get(name string) (*v1.CloudPrivateIPConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("cloudprivateipconfig"), name)
	}
	return obj.(*v1.CloudPrivateIPConfig), nil
}
