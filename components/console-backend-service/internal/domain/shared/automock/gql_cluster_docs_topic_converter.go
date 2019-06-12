// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
import mock "github.com/stretchr/testify/mock"

import v1alpha1 "github.com/kyma-project/kyma/components/cms-controller-manager/pkg/apis/cms/v1alpha1"

// GqlClusterDocsTopicConverter is an autogenerated mock type for the GqlClusterDocsTopicConverter type
type GqlClusterDocsTopicConverter struct {
	mock.Mock
}

// ToGQL provides a mock function with given fields: item
func (_m *GqlClusterDocsTopicConverter) ToGQL(item *v1alpha1.ClusterDocsTopic) (*gqlschema.ClusterDocsTopic, error) {
	ret := _m.Called(item)

	var r0 *gqlschema.ClusterDocsTopic
	if rf, ok := ret.Get(0).(func(*v1alpha1.ClusterDocsTopic) *gqlschema.ClusterDocsTopic); ok {
		r0 = rf(item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.ClusterDocsTopic)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.ClusterDocsTopic) error); ok {
		r1 = rf(item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQLs provides a mock function with given fields: in
func (_m *GqlClusterDocsTopicConverter) ToGQLs(in []*v1alpha1.ClusterDocsTopic) ([]gqlschema.ClusterDocsTopic, error) {
	ret := _m.Called(in)

	var r0 []gqlschema.ClusterDocsTopic
	if rf, ok := ret.Get(0).(func([]*v1alpha1.ClusterDocsTopic) []gqlschema.ClusterDocsTopic); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.ClusterDocsTopic)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*v1alpha1.ClusterDocsTopic) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
