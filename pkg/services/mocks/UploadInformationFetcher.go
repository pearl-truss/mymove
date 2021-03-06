// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	services "github.com/transcom/mymove/pkg/services"

	uuid "github.com/gofrs/uuid"
)

// UploadInformationFetcher is an autogenerated mock type for the UploadInformationFetcher type
type UploadInformationFetcher struct {
	mock.Mock
}

// FetchUploadInformation provides a mock function with given fields: _a0
func (_m *UploadInformationFetcher) FetchUploadInformation(_a0 uuid.UUID) (services.UploadInformation, error) {
	ret := _m.Called(_a0)

	var r0 services.UploadInformation
	if rf, ok := ret.Get(0).(func(uuid.UUID) services.UploadInformation); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(services.UploadInformation)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
