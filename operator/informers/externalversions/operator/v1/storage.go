// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	operatorv1 "github.com/uccps-samples/api/operator/v1"
	versioned "github.com/uccps-samples/client-go/operator/clientset/versioned"
	internalinterfaces "github.com/uccps-samples/client-go/operator/informers/externalversions/internalinterfaces"
	v1 "github.com/uccps-samples/client-go/operator/listers/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StorageInformer provides access to a shared informer and lister for
// Storages.
type StorageInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.StorageLister
}

type storageInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewStorageInformer constructs a new informer for Storage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorageInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredStorageInformer constructs a new informer for Storage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().Storages().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().Storages().Watch(context.TODO(), options)
			},
		},
		&operatorv1.Storage{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorageInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *storageInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operatorv1.Storage{}, f.defaultInformer)
}

func (f *storageInformer) Lister() v1.StorageLister {
	return v1.NewStorageLister(f.Informer().GetIndexer())
}
