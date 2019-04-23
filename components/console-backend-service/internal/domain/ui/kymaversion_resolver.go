package ui

import (
	"context"
)

// //go:generate mockery -name=backendModuleLister -output=automock -outpkg=automock -case=underscore
// type backendModuleLister interface {
// 	List() ([]*v1alpha1.BackendModule, error)
// }

// //go:generate mockery -name=gqlBackendModuleConverter -output=automock -outpkg=automock -case=underscore
// type gqlBackendModuleConverter interface {
// 	ToGQL(in *v1alpha1.BackendModule) (*gqlschema.BackendModule, error)
// 	ToGQLs(in []*v1alpha1.BackendModule) ([]gqlschema.BackendModule, error)
// }

type kymaVersionResolver struct {
	// backendModuleLister    backendModuleLister
	// backendModuleConverter gqlBackendModuleConverter
	deploymentService deploymentService
}

func newKymaVersionResolver() *kymaVersionResolver {
	return &kymaVersionResolver{
		// backendModuleLister:    backendModuleLister,
		// backendModuleConverter: &backendModuleConverter{},
	}
}

func (r *kymaVersionResolver) KymaVersionQuery(ctx context.Context) (string, error) {
	// var items []*v1alpha1.BackendModule
	// var err error

	// items, err = r.backendModuleLister.List()

	// if err != nil {
	// 	glog.Error(errors.Wrapf(err, "while listing %s", pretty.BackendModules))
	// 	return nil, gqlerror.New(err, pretty.BackendModules)
	// }

	// serviceInstances, err := r.backendModuleConverter.ToGQLs(items)
	// if err != nil {
	// 	glog.Error(errors.Wrapf(err, "while converting %s", pretty.BackendModules))
	// 	return nil, gqlerror.New(err, pretty.BackendModules)
	// }

	return "dupa", nil
}
