// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	configv1 "github.com/uccps-samples/api/config/v1"
	versioned "github.com/uccps-samples/client-go/config/clientset/versioned"
	internalinterfaces "github.com/uccps-samples/client-go/config/informers/externalversions/internalinterfaces"
	v1 "github.com/uccps-samples/client-go/config/listers/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterOperatorInformer provides access to a shared informer and lister for
// ClusterOperators.
type ClusterOperatorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterOperatorLister
}

type clusterOperatorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterOperatorInformer constructs a new informer for ClusterOperator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterOperatorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterOperatorInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterOperatorInformer constructs a new informer for ClusterOperator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterOperatorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1().ClusterOperators().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1().ClusterOperators().Watch(context.TODO(), options)
			},
		},
		&configv1.ClusterOperator{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterOperatorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterOperatorInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterOperatorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configv1.ClusterOperator{}, f.defaultInformer)
}

func (f *clusterOperatorInformer) Lister() v1.ClusterOperatorLister {
	return v1.NewClusterOperatorLister(f.Informer().GetIndexer())
}
