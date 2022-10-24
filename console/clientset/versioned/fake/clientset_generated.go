// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/uccps-samples/client-go/console/clientset/versioned"
	consolev1 "github.com/uccps-samples/client-go/console/clientset/versioned/typed/console/v1"
	fakeconsolev1 "github.com/uccps-samples/client-go/console/clientset/versioned/typed/console/v1/fake"
	consolev1alpha1 "github.com/uccps-samples/client-go/console/clientset/versioned/typed/console/v1alpha1"
	fakeconsolev1alpha1 "github.com/uccps-samples/client-go/console/clientset/versioned/typed/console/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// ConsoleV1 retrieves the ConsoleV1Client
func (c *Clientset) ConsoleV1() consolev1.ConsoleV1Interface {
	return &fakeconsolev1.FakeConsoleV1{Fake: &c.Fake}
}

// ConsoleV1alpha1 retrieves the ConsoleV1alpha1Client
func (c *Clientset) ConsoleV1alpha1() consolev1alpha1.ConsoleV1alpha1Interface {
	return &fakeconsolev1alpha1.FakeConsoleV1alpha1{Fake: &c.Fake}
}
