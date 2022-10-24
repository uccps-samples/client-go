// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	securityv1 "github.com/uccps-samples/api/security/v1"
	versioned "github.com/uccps-samples/client-go/security/clientset/versioned"
	internalinterfaces "github.com/uccps-samples/client-go/security/informers/externalversions/internalinterfaces"
	v1 "github.com/uccps-samples/client-go/security/listers/security/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RangeAllocationInformer provides access to a shared informer and lister for
// RangeAllocations.
type RangeAllocationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.RangeAllocationLister
}

type rangeAllocationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewRangeAllocationInformer constructs a new informer for RangeAllocation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRangeAllocationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRangeAllocationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredRangeAllocationInformer constructs a new informer for RangeAllocation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRangeAllocationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecurityV1().RangeAllocations().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecurityV1().RangeAllocations().Watch(context.TODO(), options)
			},
		},
		&securityv1.RangeAllocation{},
		resyncPeriod,
		indexers,
	)
}

func (f *rangeAllocationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRangeAllocationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *rangeAllocationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&securityv1.RangeAllocation{}, f.defaultInformer)
}

func (f *rangeAllocationInformer) Lister() v1.RangeAllocationLister {
	return v1.NewRangeAllocationLister(f.Informer().GetIndexer())
}
