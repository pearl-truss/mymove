// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	time "time"

	unit "github.com/transcom/mymove/pkg/unit"
)

// ManagementServicesPricer is an autogenerated mock type for the ManagementServicesPricer type
type ManagementServicesPricer struct {
	mock.Mock
}

// Price provides a mock function with given fields: contractCode, mtoAvailableToPrimeAt
func (_m *ManagementServicesPricer) Price(contractCode string, mtoAvailableToPrimeAt time.Time) (unit.Cents, error) {
	ret := _m.Called(contractCode, mtoAvailableToPrimeAt)

	var r0 unit.Cents
	if rf, ok := ret.Get(0).(func(string, time.Time) unit.Cents); ok {
		r0 = rf(contractCode, mtoAvailableToPrimeAt)
	} else {
		r0 = ret.Get(0).(unit.Cents)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, time.Time) error); ok {
		r1 = rf(contractCode, mtoAvailableToPrimeAt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PriceUsingParams provides a mock function with given fields: params
func (_m *ManagementServicesPricer) PriceUsingParams(params models.PaymentServiceItemParams) (unit.Cents, error) {
	ret := _m.Called(params)

	var r0 unit.Cents
	if rf, ok := ret.Get(0).(func(models.PaymentServiceItemParams) unit.Cents); ok {
		r0 = rf(params)
	} else {
		r0 = ret.Get(0).(unit.Cents)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.PaymentServiceItemParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
