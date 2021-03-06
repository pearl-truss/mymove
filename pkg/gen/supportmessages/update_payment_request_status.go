// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpdatePaymentRequestStatus update payment request status
// swagger:model UpdatePaymentRequestStatus
type UpdatePaymentRequestStatus struct {

	// Attribute of the payment request object that automatically changes when the request is updated. This matches the value passed in the header for `If-Match`. Required when sending PUT or PATCH requests to prevent updating stale data.
	ETag string `json:"eTag,omitempty"`

	// A written reason to provide context for the status.
	RejectionReason *string `json:"rejectionReason,omitempty"`

	// status
	Status PaymentRequestStatus `json:"status,omitempty"`
}

// Validate validates this update payment request status
func (m *UpdatePaymentRequestStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdatePaymentRequestStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePaymentRequestStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePaymentRequestStatus) UnmarshalBinary(b []byte) error {
	var res UpdatePaymentRequestStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
