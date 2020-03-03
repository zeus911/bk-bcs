/*
Copyright 2020 The Kubernetes Authors.

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

package v1beta1

import (
	v1beta1 "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedRoleLister helps list FederatedRoles.
type FederatedRoleLister interface {
	// List lists all FederatedRoles in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.FederatedRole, err error)
	// FederatedRoles returns an object that can list and get FederatedRoles.
	FederatedRoles(namespace string) FederatedRoleNamespaceLister
	FederatedRoleListerExpansion
}

// federatedRoleLister implements the FederatedRoleLister interface.
type federatedRoleLister struct {
	indexer cache.Indexer
}

// NewFederatedRoleLister returns a new FederatedRoleLister.
func NewFederatedRoleLister(indexer cache.Indexer) FederatedRoleLister {
	return &federatedRoleLister{indexer: indexer}
}

// List lists all FederatedRoles in the indexer.
func (s *federatedRoleLister) List(selector labels.Selector) (ret []*v1beta1.FederatedRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedRole))
	})
	return ret, err
}

// FederatedRoles returns an object that can list and get FederatedRoles.
func (s *federatedRoleLister) FederatedRoles(namespace string) FederatedRoleNamespaceLister {
	return federatedRoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedRoleNamespaceLister helps list and get FederatedRoles.
type FederatedRoleNamespaceLister interface {
	// List lists all FederatedRoles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.FederatedRole, err error)
	// Get retrieves the FederatedRole from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.FederatedRole, error)
	FederatedRoleNamespaceListerExpansion
}

// federatedRoleNamespaceLister implements the FederatedRoleNamespaceLister
// interface.
type federatedRoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedRoles in the indexer for a given namespace.
func (s federatedRoleNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedRole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedRole))
	})
	return ret, err
}

// Get retrieves the FederatedRole from the indexer for a given namespace and name.
func (s federatedRoleNamespaceLister) Get(name string) (*v1beta1.FederatedRole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federatedrole"), name)
	}
	return obj.(*v1beta1.FederatedRole), nil
}
