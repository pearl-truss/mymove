// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MoveOrder move order
// swagger:model MoveOrder
type MoveOrder struct {

	// agency
	Agency string `json:"agency,omitempty"`

	// confirmation number
	ConfirmationNumber string `json:"confirmation_number,omitempty"`

	// customer ID
	// Format: uuid
	CustomerID strfmt.UUID `json:"customerID,omitempty"`

	// destination duty station
	DestinationDutyStation *DutyStation `json:"destinationDutyStation,omitempty"`

	// entitlement
	Entitlement *Entitlements `json:"entitlement,omitempty"`

	// first name
	// Read Only: true
	FirstName string `json:"first_name,omitempty"`

	// grade
	Grade string `json:"grade,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// last name
	// Read Only: true
	LastName string `json:"last_name,omitempty"`

	// move task order ID
	// Format: uuid
	MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

	// origin duty station
	OriginDutyStation *DutyStation `json:"originDutyStation,omitempty"`
}

// Validate validates this move order
func (m *MoveOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitlement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTaskOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveOrder) validateCustomerID(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerID) { // not required
		return nil
	}

	if err := validate.FormatOf("customerID", "body", "uuid", m.CustomerID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveOrder) validateDestinationDutyStation(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationDutyStation) { // not required
		return nil
	}

	if m.DestinationDutyStation != nil {
		if err := m.DestinationDutyStation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationDutyStation")
			}
			return err
		}
	}

	return nil
}

func (m *MoveOrder) validateEntitlement(formats strfmt.Registry) error {

	if swag.IsZero(m.Entitlement) { // not required
		return nil
	}

	if m.Entitlement != nil {
		if err := m.Entitlement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entitlement")
			}
			return err
		}
	}

	return nil
}

func (m *MoveOrder) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveOrder) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrderID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveOrder) validateOriginDutyStation(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginDutyStation) { // not required
		return nil
	}

	if m.OriginDutyStation != nil {
		if err := m.OriginDutyStation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originDutyStation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveOrder) UnmarshalBinary(b []byte) error {
	var res MoveOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
