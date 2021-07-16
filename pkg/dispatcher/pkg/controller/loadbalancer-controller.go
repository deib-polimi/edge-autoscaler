package controller

import (
	"fmt"
	"net/url"
	"time"

	"github.com/lterrac/edge-autoscaler/pkg/dispatcher/pkg/balancer"
	eaclientset "github.com/lterrac/edge-autoscaler/pkg/generated/clientset/versioned"
	eascheme "github.com/lterrac/edge-autoscaler/pkg/generated/clientset/versioned/scheme"
	"github.com/lterrac/edge-autoscaler/pkg/informers"
	"github.com/lterrac/edge-autoscaler/pkg/queue"
	corev1 "k8s.io/api/core/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
)

const (
	controllerAgentName string = "loadbalancer-controller"

	// SuccessSynced is used as part of the Event 'reason' when a podScale is synced
	SuccessSynced string = "Synced"

	// MessageResourceSynced is the message used for an Event fired when a configmap
	// is synced successfully
	MessageResourceSynced string = "Community Settings synced successfully"
)

// LoadBalancerController works at node level to forward an incoming request for a function
// to the right backend, implementing load balancing policies.
type LoadBalancerController struct {
	// saClientSet is a clientset for our own API group
	edgeAutoscalerClientSet eaclientset.Interface

	// kubernetesCLientset is the client-go of kubernetes
	kubernetesClientset kubernetes.Interface

	// balancers keeps track of the load balancers associated to a function
	// TODO: change key with the openfaas reference?
	balancers map[url.URL]*balancer.LoadBalancer

	listers informers.Listers

	nodeSynced cache.InformerSynced

	//TODO: change with proper CRD
	configmapSynced cache.InformerSynced

	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder

	// workqueue contains all the communityconfigurations to sync
	workqueue queue.Queue
}

// NewController returns a new SystemController
func NewController(
	kubernetesClientset *kubernetes.Clientset,
	eaClientSet eaclientset.Interface,
	informers informers.Informers,
) *LoadBalancerController {

	// Create event broadcaster
	// Add system-controller types to the default Kubernetes Scheme so Events can be
	// logged for system-controller types.
	utilruntime.Must(eascheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	// Instantiate the Controller
	controller := &LoadBalancerController{
		edgeAutoscalerClientSet: eaClientSet,
		kubernetesClientset:     kubernetesClientset,
		recorder:                recorder,
		listers:                 informers.GetListers(),
		nodeSynced:              informers.Node.Informer().HasSynced,
		//TODO: change with proper CRD
		configmapSynced: informers.ConfigMap.Informer().HasSynced,
		workqueue:       queue.NewQueue("ConfigMapQueue"),
	}

	klog.Info("Setting up event handlers")
	// Set up an event handler for when ServiceLevelAgreements resources change
	informers.CommunityConfiguration.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    controller.workqueue.Add,
		UpdateFunc: controller.workqueue.Update,
		DeleteFunc: controller.workqueue.Deletion,
	})

	return controller
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *LoadBalancerController) Run(threadiness int, stopCh <-chan struct{}) error {

	// Start the informer factories to begin populating the informer caches
	klog.Info("Starting system level controller")

	// Wait for the caches to be synced before starting workers
	klog.Info("Waiting for informer caches to sync")

	if ok := cache.WaitForCacheSync(
		stopCh,
		//TODO: change with proper CRD
		c.configmapSynced,
		c.nodeSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.Info("Starting system controller workers")

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runStandardWorker, time.Second, stopCh)
	}

	return nil
}

// handles standard partitioning (e.g. first partioning and cache sync)
func (c *LoadBalancerController) runStandardWorker() {
	for c.workqueue.ProcessNextItem(c.syncConfigMap) {
	}
}

// control loop to handle performance degradation inside communities
func (c *LoadBalancerController) runPerformanceDegradationObserver() {
}

// control loop to handle cluster topology changes
func (c *LoadBalancerController) runTopologyObserver() {
}

// Shutdown is called when the controller has finished its work
func (c *LoadBalancerController) Shutdown() {
	utilruntime.HandleCrash()
}