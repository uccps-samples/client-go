// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/uccps-samples/api/console/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConsolePluginLister helps list ConsolePlugins.
// All objects returned here must be treated as read-only.
type ConsolePluginLister interface {
	// List lists all ConsolePlugins in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConsolePlugin, err error)
	// Get retrieves the ConsolePlugin from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ConsolePlugin, error)
	ConsolePluginListerExpansion
}

// consolePluginLister implements the ConsolePluginLister interface.
type consolePluginLister struct {
	indexer cache.Indexer
}

// NewConsolePluginLister returns a new ConsolePluginLister.
func NewConsolePluginLister(indexer cache.Indexer) ConsolePluginLister {
	return &consolePluginLister{indexer: indexer}
}

// List lists all ConsolePlugins in the indexer.
func (s *consolePluginLister) List(selector labels.Selector) (ret []*v1alpha1.ConsolePlugin, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConsolePlugin))
	})
	return ret, err
}

// Get retrieves the ConsolePlugin from the index for a given name.
func (s *consolePluginLister) Get(name string) (*v1alpha1.ConsolePlugin, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("consoleplugin"), name)
	}
	return obj.(*v1alpha1.ConsolePlugin), nil
}
