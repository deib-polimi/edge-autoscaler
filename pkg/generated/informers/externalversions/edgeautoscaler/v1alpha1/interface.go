// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/lterrac/edge-autoscaler/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CommunitySchedules returns a CommunityScheduleInformer.
	CommunitySchedules() CommunityScheduleInformer
	// CommunitySettingses returns a CommunitySettingsInformer.
	CommunitySettingses() CommunitySettingsInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// CommunitySchedules returns a CommunityScheduleInformer.
func (v *version) CommunitySchedules() CommunityScheduleInformer {
	return &communityScheduleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CommunitySettingses returns a CommunitySettingsInformer.
func (v *version) CommunitySettingses() CommunitySettingsInformer {
	return &communitySettingsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}