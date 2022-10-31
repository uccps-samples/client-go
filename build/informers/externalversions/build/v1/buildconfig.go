// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	buildv1 "github.com/uccps-samples/api/build/v1"
	versioned "github.com/uccps-samples/client-go/build/clientset/versioned"
	internalinterfaces "github.com/uccps-samples/client-go/build/informers/externalversions/internalinterfaces"
	v1 "github.com/uccps-samples/client-go/build/listers/build/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BuildConfigInformer provides access to a shared informer and lister for
// BuildConfigs.
type BuildConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.BuildConfigLister
}

type buildConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBuildConfigInformer constructs a new informer for BuildConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBuildConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBuildConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBuildConfigInformer constructs a new informer for BuildConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBuildConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BuildV1().BuildConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BuildV1().BuildConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&buildv1.BuildConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *buildConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBuildConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *buildConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&buildv1.BuildConfig{}, f.defaultInformer)
}

func (f *buildConfigInformer) Lister() v1.BuildConfigLister {
	return v1.NewBuildConfigLister(f.Informer().GetIndexer())
}
