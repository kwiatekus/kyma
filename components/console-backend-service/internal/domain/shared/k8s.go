package shared

import (
	api "k8s.io/api/apps/v1beta2"
)

//go:generate mockery -name=K8sRetriever -output=automock -outpkg=automock -case=underscore
type K8sRetriever interface {
	Deployment() DeploymentGetter
}

//go:generate mockery -name=DeploymentGetter -output=automock -outpkg=automock -case=underscore
type DeploymentGetter interface {
	Find(name string, namespace string) (*api.Deployment, error)
}
