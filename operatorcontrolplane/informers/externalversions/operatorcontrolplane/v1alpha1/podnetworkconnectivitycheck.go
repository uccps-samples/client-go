// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	operatorcontrolplanev1alpha1 "github.com/uccps-samples/api/operatorcontrolplane/v1alpha1"
	versioned "github.com/uccps-samples/client-go/operatorcontrolplane/clientset/versioned"
	internalinterfaces "github.com/uccps-samples/client-go/operatorcontrolplane/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/uccps-samples/client-go/operatorcontrolplane/listers/operatorcontrolplane/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodNetworkConnectivityCheckInformer provides access to a shared informer and lister for
// PodNetworkConnectivityChecks.
type PodNetworkConnectivityCheckInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PodNetworkConnectivityCheckLister
}

type podNetworkConnectivityCheckInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodNetworkConnectivityCheckInformer constructs a new informer for PodNetworkConnectivityCheck type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodNetworkConnectivityCheckInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodNetworkConnectivityCheckInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodNetworkConnectivityCheckInformer constructs a new informer for PodNetworkConnectivityCheck type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodNetworkConnectivityCheckInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ControlplaneV1alpha1().PodNetworkConnectivityChecks(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ControlplaneV1alpha1().PodNetworkConnectivityChecks(namespace).Watch(context.TODO(), options)
			},
		},
		&operatorcontrolplanev1alpha1.PodNetworkConnectivityCheck{},
		resyncPeriod,
		indexers,
	)
}

func (f *podNetworkConnectivityCheckInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodNetworkConnectivityCheckInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podNetworkConnectivityCheckInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operatorcontrolplanev1alpha1.PodNetworkConnectivityCheck{}, f.defaultInformer)
}

func (f *podNetworkConnectivityCheckInformer) Lister() v1alpha1.PodNetworkConnectivityCheckLister {
	return v1alpha1.NewPodNetworkConnectivityCheckLister(f.Informer().GetIndexer())
}
