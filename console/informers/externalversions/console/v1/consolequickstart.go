// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	consolev1 "github.com/uccps-samples/api/console/v1"
	versioned "github.com/uccps-samples/client-go/console/clientset/versioned"
	internalinterfaces "github.com/uccps-samples/client-go/console/informers/externalversions/internalinterfaces"
	v1 "github.com/uccps-samples/client-go/console/listers/console/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConsoleQuickStartInformer provides access to a shared informer and lister for
// ConsoleQuickStarts.
type ConsoleQuickStartInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ConsoleQuickStartLister
}

type consoleQuickStartInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewConsoleQuickStartInformer constructs a new informer for ConsoleQuickStart type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConsoleQuickStartInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConsoleQuickStartInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredConsoleQuickStartInformer constructs a new informer for ConsoleQuickStart type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConsoleQuickStartInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConsoleV1().ConsoleQuickStarts().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConsoleV1().ConsoleQuickStarts().Watch(context.TODO(), options)
			},
		},
		&consolev1.ConsoleQuickStart{},
		resyncPeriod,
		indexers,
	)
}

func (f *consoleQuickStartInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConsoleQuickStartInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *consoleQuickStartInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&consolev1.ConsoleQuickStart{}, f.defaultInformer)
}

func (f *consoleQuickStartInformer) Lister() v1.ConsoleQuickStartLister {
	return v1.NewConsoleQuickStartLister(f.Informer().GetIndexer())
}
