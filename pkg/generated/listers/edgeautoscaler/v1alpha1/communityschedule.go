// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/lterrac/edge-autoscaler/pkg/apis/edgeautoscaler/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CommunityScheduleLister helps list CommunitySchedules.
// All objects returned here must be treated as read-only.
type CommunityScheduleLister interface {
	// List lists all CommunitySchedules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CommunitySchedule, err error)
	// CommunitySchedules returns an object that can list and get CommunitySchedules.
	CommunitySchedules(namespace string) CommunityScheduleNamespaceLister
	CommunityScheduleListerExpansion
}

// communityScheduleLister implements the CommunityScheduleLister interface.
type communityScheduleLister struct {
	indexer cache.Indexer
}

// NewCommunityScheduleLister returns a new CommunityScheduleLister.
func NewCommunityScheduleLister(indexer cache.Indexer) CommunityScheduleLister {
	return &communityScheduleLister{indexer: indexer}
}

// List lists all CommunitySchedules in the indexer.
func (s *communityScheduleLister) List(selector labels.Selector) (ret []*v1alpha1.CommunitySchedule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CommunitySchedule))
	})
	return ret, err
}

// CommunitySchedules returns an object that can list and get CommunitySchedules.
func (s *communityScheduleLister) CommunitySchedules(namespace string) CommunityScheduleNamespaceLister {
	return communityScheduleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CommunityScheduleNamespaceLister helps list and get CommunitySchedules.
// All objects returned here must be treated as read-only.
type CommunityScheduleNamespaceLister interface {
	// List lists all CommunitySchedules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CommunitySchedule, err error)
	// Get retrieves the CommunitySchedule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CommunitySchedule, error)
	CommunityScheduleNamespaceListerExpansion
}

// communityScheduleNamespaceLister implements the CommunityScheduleNamespaceLister
// interface.
type communityScheduleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CommunitySchedules in the indexer for a given namespace.
func (s communityScheduleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CommunitySchedule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CommunitySchedule))
	})
	return ret, err
}

// Get retrieves the CommunitySchedule from the indexer for a given namespace and name.
func (s communityScheduleNamespaceLister) Get(name string) (*v1alpha1.CommunitySchedule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("communityschedule"), name)
	}
	return obj.(*v1alpha1.CommunitySchedule), nil
}