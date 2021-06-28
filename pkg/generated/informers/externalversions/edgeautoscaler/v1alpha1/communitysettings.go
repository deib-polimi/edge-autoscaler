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

// CommunitySettingsInformer provides access to a shared informer and lister for
// CommunitySettingses.
type CommunitySettingsInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CommunitySettingsLister
}

type communitySettingsInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCommunitySettingsInformer constructs a new informer for CommunitySettings type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCommunitySettingsInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCommunitySettingsInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCommunitySettingsInformer constructs a new informer for CommunitySettings type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCommunitySettingsInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdgeautoscalerV1alpha1().CommunitySettingses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdgeautoscalerV1alpha1().CommunitySettingses(namespace).Watch(context.TODO(), options)
			},
		},
		&edgeautoscalerv1alpha1.CommunitySettings{},
		resyncPeriod,
		indexers,
	)
}

func (f *communitySettingsInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCommunitySettingsInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *communitySettingsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&edgeautoscalerv1alpha1.CommunitySettings{}, f.defaultInformer)
}

func (f *communitySettingsInformer) Lister() v1alpha1.CommunitySettingsLister {
	return v1alpha1.NewCommunitySettingsLister(f.Informer().GetIndexer())
}