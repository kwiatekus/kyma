// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kyma-project/kyma/components/connectivity-certs-controller/pkg/apis/applicationconnector/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CentralConnectionLister helps list CentralConnections.
type CentralConnectionLister interface {
	// List lists all CentralConnections in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CentralConnection, err error)
	// Get retrieves the CentralConnection from the index for a given name.
	Get(name string) (*v1alpha1.CentralConnection, error)
	CentralConnectionListerExpansion
}

// centralConnectionLister implements the CentralConnectionLister interface.
type centralConnectionLister struct {
	indexer cache.Indexer
}

// NewCentralConnectionLister returns a new CentralConnectionLister.
func NewCentralConnectionLister(indexer cache.Indexer) CentralConnectionLister {
	return &centralConnectionLister{indexer: indexer}
}

// List lists all CentralConnections in the indexer.
func (s *centralConnectionLister) List(selector labels.Selector) (ret []*v1alpha1.CentralConnection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CentralConnection))
	})
	return ret, err
}

// Get retrieves the CentralConnection from the index for a given name.
func (s *centralConnectionLister) Get(name string) (*v1alpha1.CentralConnection, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("centralconnection"), name)
	}
	return obj.(*v1alpha1.CentralConnection), nil
}
