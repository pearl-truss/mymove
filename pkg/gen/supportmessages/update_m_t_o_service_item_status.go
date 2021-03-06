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

// UpdateMTOServiceItemStatus update m t o service item status
// swagger:model UpdateMTOServiceItemStatus
type UpdateMTOServiceItemStatus struct {

	// description
	// Read Only: true
	Description string `json:"description,omitempty"`

	// e tag
	// Read Only: true
	ETag string `json:"eTag,omitempty"`

	// fee type
	// Read Only: true
	// Enum: [COUNSELING CRATING TRUCKING SHUTTLE]
	FeeType string `json:"feeType,omitempty"`

	// id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// move task order ID
	// Read Only: true
	// Format: uuid
	MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

	// mto shipment ID
	// Read Only: true
	// Format: uuid
	MtoShipmentID strfmt.UUID `json:"mtoShipmentID,omitempty"`

	// quantity
	// Read Only: true
	Quantity int64 `json:"quantity,omitempty"`

	// rate
	// Read Only: true
	Rate int64 `json:"rate,omitempty"`

	// re service code
	// Read Only: true
	ReServiceCode string `json:"reServiceCode,omitempty"`

	// re service ID
	// Read Only: true
	// Format: uuid
	ReServiceID strfmt.UUID `json:"reServiceID,omitempty"`

	// re service name
	// Read Only: true
	ReServiceName string `json:"reServiceName,omitempty"`

	// rejection reason
	RejectionReason *string `json:"rejectionReason,omitempty"`

	// status
	// Required: true
	Status MTOServiceItemStatus `json:"status"`

	// total
	// Read Only: true
	Total int64 `json:"total,omitempty"`
}

// Validate validates this update m t o service item status
func (m *UpdateMTOServiceItemStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTaskOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtoShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReServiceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateMTOServiceItemStatusTypeFeeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["COUNSELING","CRATING","TRUCKING","SHUTTLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateMTOServiceItemStatusTypeFeeTypePropEnum = append(updateMTOServiceItemStatusTypeFeeTypePropEnum, v)
	}
}

const (

	// UpdateMTOServiceItemStatusFeeTypeCOUNSELING captures enum value "COUNSELING"
	UpdateMTOServiceItemStatusFeeTypeCOUNSELING string = "COUNSELING"

	// UpdateMTOServiceItemStatusFeeTypeCRATING captures enum value "CRATING"
	UpdateMTOServiceItemStatusFeeTypeCRATING string = "CRATING"

	// UpdateMTOServiceItemStatusFeeTypeTRUCKING captures enum value "TRUCKING"
	UpdateMTOServiceItemStatusFeeTypeTRUCKING string = "TRUCKING"

	// UpdateMTOServiceItemStatusFeeTypeSHUTTLE captures enum value "SHUTTLE"
	UpdateMTOServiceItemStatusFeeTypeSHUTTLE string = "SHUTTLE"
)

// prop value enum
func (m *UpdateMTOServiceItemStatus) validateFeeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateMTOServiceItemStatusTypeFeeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateMTOServiceItemStatus) validateFeeType(formats strfmt.Registry) error {

	if swag.IsZero(m.FeeType) { // not required
		return nil
	}

	// value enum
	if err := m.validateFeeTypeEnum("feeType", "body", m.FeeType); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMTOServiceItemStatus) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMTOServiceItemStatus) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrderID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMTOServiceItemStatus) validateMtoShipmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.MtoShipmentID) { // not required
		return nil
	}

	if err := validate.FormatOf("mtoShipmentID", "body", "uuid", m.MtoShipmentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMTOServiceItemStatus) validateReServiceID(formats strfmt.Registry) error {

	if swag.IsZero(m.ReServiceID) { // not required
		return nil
	}

	if err := validate.FormatOf("reServiceID", "body", "uuid", m.ReServiceID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMTOServiceItemStatus) validateStatus(formats strfmt.Registry) error {

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateMTOServiceItemStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateMTOServiceItemStatus) UnmarshalBinary(b []byte) error {
	var res UpdateMTOServiceItemStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
