// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOServiceItemDimension Describes a dimension object for the MTOServiceItem
// swagger:model MTOServiceItemDimension
type MTOServiceItemDimension struct {

	// Height in thousandth inches. 1000 thou = 1 inch.
	// Required: true
	Height *int32 `json:"height"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Length in thousandth inches. 1000 thou = 1 inch.
	// Required: true
	Length *int32 `json:"length"`

	// type
	Type DimensionType `json:"type,omitempty"`

	// Width in thousandth inches. 1000 thou = 1 inch.
	// Required: true
	Width *int32 `json:"width"`
}

// Validate validates this m t o service item dimension
func (m *MTOServiceItemDimension) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MTOServiceItemDimension) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDimension) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDimension) validateLength(formats strfmt.Registry) error {

	if err := validate.Required("length", "body", m.Length); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDimension) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *MTOServiceItemDimension) validateWidth(formats strfmt.Registry) error {

	if err := validate.Required("width", "body", m.Width); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOServiceItemDimension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOServiceItemDimension) UnmarshalBinary(b []byte) error {
	var res MTOServiceItemDimension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
