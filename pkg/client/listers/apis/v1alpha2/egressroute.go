/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// EgressRouteLister helps list EgressRoutes.
// All objects returned here must be treated as read-only.
type EgressRouteLister interface {
	// List lists all EgressRoutes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.EgressRoute, err error)
	// EgressRoutes returns an object that can list and get EgressRoutes.
	EgressRoutes(namespace string) EgressRouteNamespaceLister
	EgressRouteListerExpansion
}

// egressRouteLister implements the EgressRouteLister interface.
type egressRouteLister struct {
	indexer cache.Indexer
}

// NewEgressRouteLister returns a new EgressRouteLister.
func NewEgressRouteLister(indexer cache.Indexer) EgressRouteLister {
	return &egressRouteLister{indexer: indexer}
}

// List lists all EgressRoutes in the indexer.
func (s *egressRouteLister) List(selector labels.Selector) (ret []*v1alpha2.EgressRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.EgressRoute))
	})
	return ret, err
}

// EgressRoutes returns an object that can list and get EgressRoutes.
func (s *egressRouteLister) EgressRoutes(namespace string) EgressRouteNamespaceLister {
	return egressRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EgressRouteNamespaceLister helps list and get EgressRoutes.
// All objects returned here must be treated as read-only.
type EgressRouteNamespaceLister interface {
	// List lists all EgressRoutes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.EgressRoute, err error)
	// Get retrieves the EgressRoute from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha2.EgressRoute, error)
	EgressRouteNamespaceListerExpansion
}

// egressRouteNamespaceLister implements the EgressRouteNamespaceLister
// interface.
type egressRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EgressRoutes in the indexer for a given namespace.
func (s egressRouteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.EgressRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.EgressRoute))
	})
	return ret, err
}

// Get retrieves the EgressRoute from the indexer for a given namespace and name.
func (s egressRouteNamespaceLister) Get(name string) (*v1alpha2.EgressRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("egressroute"), name)
	}
	return obj.(*v1alpha2.EgressRoute), nil
}