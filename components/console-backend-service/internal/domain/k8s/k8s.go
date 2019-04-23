package k8s

import (
	"time"

	"github.com/kyma-project/kyma/components/console-backend-service/internal/domain/shared"
	"github.com/kyma-project/kyma/components/console-backend-service/pkg/dynamic/dynamicinformer"

	"github.com/kyma-project/kyma/components/application-operator/pkg/apis/applicationconnector/v1alpha1"
	"github.com/pkg/errors"
	"k8s.io/client-go/informers"
	k8sClientset "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

type ApplicationLister interface {
	ListInNamespace(namespace string) ([]*v1alpha1.Application, error)
	ListNamespacesFor(reName string) ([]string, error)
}

type Resolver struct {
	*resourceResolver
	*namespaceResolver
	*secretResolver
	*deploymentResolver
	*resourceQuotaResolver
	*resourceQuotaStatusResolver
	*limitRangeResolver
	*podResolver
	*serviceResolver
	*replicaSetResolver
	*configMapResolver
}

type k8sRetriever struct {
	DeploymentGetter shared.DeploymentGetter
}

func (r *k8sRetriever) Deployment() shared.DeploymentGetter {
	return r.DeploymentGetter
}

type Container struct {
	Resolver        Resolver
	K8sRetriever    *k8sRetriever
	informerFactory dynamicinformer.DynamicSharedInformerFactory
}

func New(restConfig *rest.Config, informerResyncPeriod time.Duration, applicationRetriever shared.ApplicationRetriever, scRetriever shared.ServiceCatalogRetriever, scaRetriever shared.ServiceCatalogAddonsRetriever) (*Container, error) {
	client, err := v1.NewForConfig(restConfig)
	if err != nil {
		return nil, errors.Wrap(err, "while creating K8S Client")
	}

	clientset, err := k8sClientset.NewForConfig(restConfig)
	if err != nil {
		return nil, errors.Wrap(err, "while creating K8S Clientset")
	}

	informerFactory := informers.NewSharedInformerFactory(clientset, informerResyncPeriod)

	namespaceService := newNamespaceService(client.Namespaces())
	deploymentService, err := newDeploymentService(informerFactory.Apps().V1beta2().Deployments().Informer())
	if err != nil {
		return nil, errors.Wrap(err, "while creating deployment service")
	}
	limitRangeService := newLimitRangeService(informerFactory.Core().V1().LimitRanges().Informer())

	podService := newPodService(informerFactory.Core().V1().Pods().Informer(), client)
	resourceService := newResourceService(clientset.Discovery())
	secretService := newSecretService(informerFactory.Core().V1().Secrets().Informer(), client)

	replicaSetService := newReplicaSetService(informerFactory.Apps().V1().ReplicaSets().Informer(), clientset.AppsV1())
	resourceQuotaService := newResourceQuotaService(informerFactory.Core().V1().ResourceQuotas().Informer(),
		informerFactory.Apps().V1().ReplicaSets().Informer(), informerFactory.Apps().V1().StatefulSets().Informer(), client)
	resourceQuotaStatusService := newResourceQuotaStatusService(resourceQuotaService, resourceQuotaService, resourceQuotaService, limitRangeService)
	configMapService := newConfigMapService(informerFactory.Core().V1().ConfigMaps().Informer(), clientset.CoreV1())
	serviceSvc := newServiceService(informerFactory.Core().V1().Services().Informer(), client)

	container := &Container{
		Resolver: Resolver{
			resourceResolver:            newResourceResolver(resourceService),
			namespaceResolver:           newNamespaceResolver(namespaceService, applicationRetriever),
			secretResolver:              newSecretResolver(*secretService),
			deploymentResolver:          newDeploymentResolver(deploymentService, scRetriever, scaRetriever),
			podResolver:                 newPodResolver(podService),
			serviceResolver:             newServiceResolver(serviceSvc),
			replicaSetResolver:          newReplicaSetResolver(replicaSetService),
			limitRangeResolver:          newLimitRangeResolver(limitRangeService),
			resourceQuotaResolver:       newResourceQuotaResolver(resourceQuotaService),
			resourceQuotaStatusResolver: newResourceQuotaStatusResolver(resourceQuotaStatusService),
			configMapResolver:           newConfigMapResolver(configMapService),
		},
		K8sRetriever:    &k8sRetriever{},
		informerFactory: informerFactory,
	}

	r.K8sRetriever.Deployment = deploymentService

	return container, nil
}

func (r *Resolver) WaitForCacheSync(stopCh <-chan struct{}) {
	r.informerFactory.Start(stopCh)
	r.informerFactory.WaitForCacheSync(stopCh)
}
