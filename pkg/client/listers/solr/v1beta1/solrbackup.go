/*
Copyright 2019 Bloomberg Finance LP.

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
	"github.com/bloomberg/solr-operator/pkg/apis/solr/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SolrBackupLister helps list SolrBackups.
type SolrBackupLister interface {
	// List lists all SolrBackups in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.SolrBackup, err error)
	// SolrBackups returns an object that can list and get SolrBackups.
	SolrBackups(namespace string) SolrBackupNamespaceLister
	SolrBackupListerExpansion
}

// solrBackupLister implements the SolrBackupLister interface.
type solrBackupLister struct {
	indexer cache.Indexer
}

// NewSolrBackupLister returns a new SolrBackupLister.
func NewSolrBackupLister(indexer cache.Indexer) SolrBackupLister {
	return &solrBackupLister{indexer: indexer}
}

// List lists all SolrBackups in the indexer.
func (s *solrBackupLister) List(selector labels.Selector) (ret []*v1beta1.SolrBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.SolrBackup))
	})
	return ret, err
}

// SolrBackups returns an object that can list and get SolrBackups.
func (s *solrBackupLister) SolrBackups(namespace string) SolrBackupNamespaceLister {
	return solrBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SolrBackupNamespaceLister helps list and get SolrBackups.
type SolrBackupNamespaceLister interface {
	// List lists all SolrBackups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.SolrBackup, err error)
	// Get retrieves the SolrBackup from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.SolrBackup, error)
	SolrBackupNamespaceListerExpansion
}

// solrBackupNamespaceLister implements the SolrBackupNamespaceLister
// interface.
type solrBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SolrBackups in the indexer for a given namespace.
func (s solrBackupNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.SolrBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.SolrBackup))
	})
	return ret, err
}

// Get retrieves the SolrBackup from the indexer for a given namespace and name.
func (s solrBackupNamespaceLister) Get(name string) (*v1beta1.SolrBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("solrbackup"), name)
	}
	return obj.(*v1beta1.SolrBackup), nil
}