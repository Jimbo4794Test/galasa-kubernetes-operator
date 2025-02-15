/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by lister-gen. DO NOT EDIT.

package v2alpha1

import (
	v2alpha1 "github.com/galasa-dev/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GalasaApiComponentLister helps list GalasaApiComponents.
// All objects returned here must be treated as read-only.
type GalasaApiComponentLister interface {
	// List lists all GalasaApiComponents in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2alpha1.GalasaApiComponent, err error)
	// GalasaApiComponents returns an object that can list and get GalasaApiComponents.
	GalasaApiComponents(namespace string) GalasaApiComponentNamespaceLister
	GalasaApiComponentListerExpansion
}

// galasaApiComponentLister implements the GalasaApiComponentLister interface.
type galasaApiComponentLister struct {
	indexer cache.Indexer
}

// NewGalasaApiComponentLister returns a new GalasaApiComponentLister.
func NewGalasaApiComponentLister(indexer cache.Indexer) GalasaApiComponentLister {
	return &galasaApiComponentLister{indexer: indexer}
}

// List lists all GalasaApiComponents in the indexer.
func (s *galasaApiComponentLister) List(selector labels.Selector) (ret []*v2alpha1.GalasaApiComponent, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.GalasaApiComponent))
	})
	return ret, err
}

// GalasaApiComponents returns an object that can list and get GalasaApiComponents.
func (s *galasaApiComponentLister) GalasaApiComponents(namespace string) GalasaApiComponentNamespaceLister {
	return galasaApiComponentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GalasaApiComponentNamespaceLister helps list and get GalasaApiComponents.
// All objects returned here must be treated as read-only.
type GalasaApiComponentNamespaceLister interface {
	// List lists all GalasaApiComponents in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2alpha1.GalasaApiComponent, err error)
	// Get retrieves the GalasaApiComponent from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2alpha1.GalasaApiComponent, error)
	GalasaApiComponentNamespaceListerExpansion
}

// galasaApiComponentNamespaceLister implements the GalasaApiComponentNamespaceLister
// interface.
type galasaApiComponentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GalasaApiComponents in the indexer for a given namespace.
func (s galasaApiComponentNamespaceLister) List(selector labels.Selector) (ret []*v2alpha1.GalasaApiComponent, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.GalasaApiComponent))
	})
	return ret, err
}

// Get retrieves the GalasaApiComponent from the indexer for a given namespace and name.
func (s galasaApiComponentNamespaceLister) Get(name string) (*v2alpha1.GalasaApiComponent, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2alpha1.Resource("galasaapicomponent"), name)
	}
	return obj.(*v2alpha1.GalasaApiComponent), nil
}
