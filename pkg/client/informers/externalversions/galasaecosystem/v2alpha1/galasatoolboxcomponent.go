/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by informer-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	time "time"

	galasaecosystemv2alpha1 "github.com/galasa-dev/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	versioned "github.com/galasa-dev/galasa-kubernetes-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/galasa-dev/galasa-kubernetes-operator/pkg/client/informers/externalversions/internalinterfaces"
	v2alpha1 "github.com/galasa-dev/galasa-kubernetes-operator/pkg/client/listers/galasaecosystem/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GalasaToolboxComponentInformer provides access to a shared informer and lister for
// GalasaToolboxComponents.
type GalasaToolboxComponentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2alpha1.GalasaToolboxComponentLister
}

type galasaToolboxComponentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGalasaToolboxComponentInformer constructs a new informer for GalasaToolboxComponent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGalasaToolboxComponentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGalasaToolboxComponentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGalasaToolboxComponentInformer constructs a new informer for GalasaToolboxComponent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGalasaToolboxComponentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GalasaV2alpha1().GalasaToolboxComponents(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GalasaV2alpha1().GalasaToolboxComponents(namespace).Watch(context.TODO(), options)
			},
		},
		&galasaecosystemv2alpha1.GalasaToolboxComponent{},
		resyncPeriod,
		indexers,
	)
}

func (f *galasaToolboxComponentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGalasaToolboxComponentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *galasaToolboxComponentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&galasaecosystemv2alpha1.GalasaToolboxComponent{}, f.defaultInformer)
}

func (f *galasaToolboxComponentInformer) Lister() v2alpha1.GalasaToolboxComponentLister {
	return v2alpha1.NewGalasaToolboxComponentLister(f.Informer().GetIndexer())
}
