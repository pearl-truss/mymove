// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	uuid "github.com/gofrs/uuid"
)

// MTOServiceItemUpdater is an autogenerated mock type for the MTOServiceItemUpdater type
type MTOServiceItemUpdater struct {
	mock.Mock
}

// UpdateMTOServiceItemStatus provides a mock function with given fields: mtoServiceItemID, status, reason, eTag
func (_m *MTOServiceItemUpdater) UpdateMTOServiceItemStatus(mtoServiceItemID uuid.UUID, status models.MTOServiceItemStatus, reason *string, eTag string) (*models.MTOServiceItem, error) {
	ret := _m.Called(mtoServiceItemID, status, reason, eTag)

	var r0 *models.MTOServiceItem
	if rf, ok := ret.Get(0).(func(uuid.UUID, models.MTOServiceItemStatus, *string, string) *models.MTOServiceItem); ok {
		r0 = rf(mtoServiceItemID, status, reason, eTag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MTOServiceItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, models.MTOServiceItemStatus, *string, string) error); ok {
		r1 = rf(mtoServiceItemID, status, reason, eTag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
