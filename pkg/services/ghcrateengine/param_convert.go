package ghcrateengine

import (
	"fmt"
	"time"

	"github.com/transcom/mymove/pkg/models"
)

// Constants for formatting Time-based parameters
const (
	DateParamFormat      = "2006-01-02"
	TimestampParamFormat = time.RFC3339
)

func getParamString(params models.PaymentServiceItemParams, name models.ServiceItemParamName) (string, error) {
	paymentServiceItemParam := getPaymentServiceItemParam(params, name)
	if paymentServiceItemParam == nil {
		return "", fmt.Errorf("could not find param with key %s", name)
	}

	paramType := paymentServiceItemParam.ServiceItemParamKey.Type
	if paramType != models.ServiceItemParamTypeString {
		return "", fmt.Errorf("trying to convert %s to a string, but param is of type %s", name, paramType)
	}

	return paymentServiceItemParam.Value, nil
}

func getParamTime(params models.PaymentServiceItemParams, name models.ServiceItemParamName) (time.Time, error) {
	paymentServiceItemParam := getPaymentServiceItemParam(params, name)
	if paymentServiceItemParam == nil {
		return time.Time{}, fmt.Errorf("could not find param with key %s", name)
	}

	paramType := paymentServiceItemParam.ServiceItemParamKey.Type
	stringValue := paymentServiceItemParam.Value
	var timeValue time.Time
	var err error
	if paramType == models.ServiceItemParamTypeDate {
		timeValue, err = time.Parse(DateParamFormat, stringValue)
		if err != nil {
			return timeValue, fmt.Errorf("could not convert %s to date: %w", stringValue, err)
		}
	} else if paramType == models.ServiceItemParamTypeTimestamp {
		timeValue, err = time.Parse(TimestampParamFormat, stringValue)
		if err != nil {
			return timeValue, fmt.Errorf("could not convert %s to timestamp: %w", stringValue, err)
		}
	} else {
		return timeValue, fmt.Errorf("trying to convert %s to a time, but param is of type %s", name, paramType)
	}

	return timeValue, nil
}

func getPaymentServiceItemParam(params models.PaymentServiceItemParams, name models.ServiceItemParamName) *models.PaymentServiceItemParam {
	for _, param := range params {
		if param.ServiceItemParamKey.Key == name {
			return &param
		}
	}

	return nil
}