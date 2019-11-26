// Code generated by go-swagger; DO NOT EDIT.

package adminmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Move move
// swagger:model Move
type Move struct {

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// locator
	// Required: true
	Locator *string `json:"locator"`

	// orders id
	// Required: true
	// Format: uuid
	OrdersID *strfmt.UUID `json:"orders_id"`

	// service member id
	// Read Only: true
	// Format: uuid
	ServiceMemberID strfmt.UUID `json:"service_member_id,omitempty"`

	// show
	Show bool `json:"show,omitempty"`

	// status
	Status MoveStatus `json:"status,omitempty"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this move
func (m *Move) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrdersID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceMemberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Move) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Move) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Move) validateLocator(formats strfmt.Registry) error {

	if err := validate.Required("locator", "body", m.Locator); err != nil {
		return err
	}

	return nil
}

func (m *Move) validateOrdersID(formats strfmt.Registry) error {

	if err := validate.Required("orders_id", "body", m.OrdersID); err != nil {
		return err
	}

	if err := validate.FormatOf("orders_id", "body", "uuid", m.OrdersID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Move) validateServiceMemberID(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceMemberID) { // not required
		return nil
	}

	if err := validate.FormatOf("service_member_id", "body", "uuid", m.ServiceMemberID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Move) validateStatus(formats strfmt.Registry) error {

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

func (m *Move) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Move) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Move) UnmarshalBinary(b []byte) error {
	var res Move
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
