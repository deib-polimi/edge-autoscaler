// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	edgeautoscalerv1alpha1 "github.com/lterrac/edge-autoscaler/pkg/apis/edgeautoscaler/v1alpha1"
	versioned "github.com/lterrac/edge-autoscaler/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/lterrac/edge-autoscaler/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/lterrac/edge-autoscaler/pkg/generated/listers/edgeautoscaler/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CommunityScheduleInformer provides access to a shared informer and lister for
// CommunitySchedules.
type CommunityScheduleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CommunityScheduleLister
}

type communityScheduleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCommunityScheduleInformer constructs a new informer for CommunitySchedule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCommunityScheduleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCommunityScheduleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCommunityScheduleInformer constructs a new informer for CommunitySchedule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCommunityScheduleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdgeautoscalerV1alpha1().CommunitySchedules(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdgeautoscalerV1alpha1().CommunitySchedules(namespace).Watch(context.TODO(), options)
			},
		},
		&edgeautoscalerv1alpha1.CommunitySchedule{},
		resyncPeriod,
		indexers,
	)
}

func (f *communityScheduleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCommunityScheduleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *communityScheduleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&edgeautoscalerv1alpha1.CommunitySchedule{}, f.defaultInformer)
}

func (f *communityScheduleInformer) Lister() v1alpha1.CommunityScheduleLister {
	return v1alpha1.NewCommunityScheduleLister(f.Informer().GetIndexer())
}