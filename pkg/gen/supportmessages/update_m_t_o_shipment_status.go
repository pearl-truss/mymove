// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateMTOShipmentStatus update m t o shipment status
// swagger:model UpdateMTOShipmentStatus
type UpdateMTOShipmentStatus struct {

	// rejection reason
	RejectionReason *string `json:"rejectionReason,omitempty"`

	// status
	// Enum: [REJECTED APPROVED SUBMITTED]
	Status string `json:"status,omitempty"`
}

// Validate validates this update m t o shipment status
func (m *UpdateMTOShipmentStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateMTOShipmentStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REJECTED","APPROVED","SUBMITTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateMTOShipmentStatusTypeStatusPropEnum = append(updateMTOShipmentStatusTypeStatusPropEnum, v)
	}
}

const (

	// UpdateMTOShipmentStatusStatusREJECTED captures enum value "REJECTED"
	UpdateMTOShipmentStatusStatusREJECTED string = "REJECTED"

	// UpdateMTOShipmentStatusStatusAPPROVED captures enum value "APPROVED"
	UpdateMTOShipmentStatusStatusAPPROVED string = "APPROVED"

	// UpdateMTOShipmentStatusStatusSUBMITTED captures enum value "SUBMITTED"
	UpdateMTOShipmentStatusStatusSUBMITTED string = "SUBMITTED"
)

// prop value enum
func (m *UpdateMTOShipmentStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateMTOShipmentStatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateMTOShipmentStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateMTOShipmentStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateMTOShipmentStatus) UnmarshalBinary(b []byte) error {
	var res UpdateMTOShipmentStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
