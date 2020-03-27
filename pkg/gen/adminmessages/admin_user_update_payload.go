// Code generated by go-swagger; DO NOT EDIT.

package adminmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AdminUserUpdatePayload admin user update payload
// swagger:model AdminUserUpdatePayload
type AdminUserUpdatePayload struct {

	// active
	Active *bool `json:"active,omitempty"`

	// First Name
	FirstName *string `json:"firstName,omitempty"`

	// Last Name
	LastName *string `json:"lastName,omitempty"`
}

// Validate validates this admin user update payload
func (m *AdminUserUpdatePayload) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdminUserUpdatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminUserUpdatePayload) UnmarshalBinary(b []byte) error {
	var res AdminUserUpdatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
