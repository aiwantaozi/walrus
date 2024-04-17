// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package v1

import (
	v1 "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResourceDefinitionMatchingRuleLister helps list ResourceDefinitionMatchingRules.
// All objects returned here must be treated as read-only.
type ResourceDefinitionMatchingRuleLister interface {
	// List lists all ResourceDefinitionMatchingRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ResourceDefinitionMatchingRule, err error)
	// ResourceDefinitionMatchingRules returns an object that can list and get ResourceDefinitionMatchingRules.
	ResourceDefinitionMatchingRules(namespace string) ResourceDefinitionMatchingRuleNamespaceLister
	ResourceDefinitionMatchingRuleListerExpansion
}

// resourceDefinitionMatchingRuleLister implements the ResourceDefinitionMatchingRuleLister interface.
type resourceDefinitionMatchingRuleLister struct {
	indexer cache.Indexer
}

// NewResourceDefinitionMatchingRuleLister returns a new ResourceDefinitionMatchingRuleLister.
func NewResourceDefinitionMatchingRuleLister(indexer cache.Indexer) ResourceDefinitionMatchingRuleLister {
	return &resourceDefinitionMatchingRuleLister{indexer: indexer}
}

// List lists all ResourceDefinitionMatchingRules in the indexer.
func (s *resourceDefinitionMatchingRuleLister) List(selector labels.Selector) (ret []*v1.ResourceDefinitionMatchingRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ResourceDefinitionMatchingRule))
	})
	return ret, err
}

// ResourceDefinitionMatchingRules returns an object that can list and get ResourceDefinitionMatchingRules.
func (s *resourceDefinitionMatchingRuleLister) ResourceDefinitionMatchingRules(namespace string) ResourceDefinitionMatchingRuleNamespaceLister {
	return resourceDefinitionMatchingRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceDefinitionMatchingRuleNamespaceLister helps list and get ResourceDefinitionMatchingRules.
// All objects returned here must be treated as read-only.
type ResourceDefinitionMatchingRuleNamespaceLister interface {
	// List lists all ResourceDefinitionMatchingRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ResourceDefinitionMatchingRule, err error)
	// Get retrieves the ResourceDefinitionMatchingRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ResourceDefinitionMatchingRule, error)
	ResourceDefinitionMatchingRuleNamespaceListerExpansion
}

// resourceDefinitionMatchingRuleNamespaceLister implements the ResourceDefinitionMatchingRuleNamespaceLister
// interface.
type resourceDefinitionMatchingRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceDefinitionMatchingRules in the indexer for a given namespace.
func (s resourceDefinitionMatchingRuleNamespaceLister) List(selector labels.Selector) (ret []*v1.ResourceDefinitionMatchingRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ResourceDefinitionMatchingRule))
	})
	return ret, err
}

// Get retrieves the ResourceDefinitionMatchingRule from the indexer for a given namespace and name.
func (s resourceDefinitionMatchingRuleNamespaceLister) Get(name string) (*v1.ResourceDefinitionMatchingRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.SchemeResource("resourcedefinitionmatchingrule"), name)
	}
	return obj.(*v1.ResourceDefinitionMatchingRule), nil
}