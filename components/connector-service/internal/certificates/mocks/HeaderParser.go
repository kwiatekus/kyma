// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apperrors "github.com/kyma-project/kyma/components/connector-service/internal/apperrors"
import certificates "github.com/kyma-project/kyma/components/connector-service/internal/certificates"
import http "net/http"
import mock "github.com/stretchr/testify/mock"

// HeaderParser is an autogenerated mock type for the HeaderParser type
type HeaderParser struct {
	mock.Mock
}

// ParseCertificateHeader provides a mock function with given fields: r
func (_m *HeaderParser) ParseCertificateHeader(r http.Request) (certificates.CertInfo, apperrors.AppError) {
	ret := _m.Called(r)

	var r0 certificates.CertInfo
	if rf, ok := ret.Get(0).(func(http.Request) certificates.CertInfo); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(certificates.CertInfo)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(http.Request) apperrors.AppError); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}
