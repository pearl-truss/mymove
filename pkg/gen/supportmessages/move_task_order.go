// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MoveTaskOrder move task order
// swagger:model MoveTaskOrder
type MoveTaskOrder struct {

	// Date the MoveTaskOrder was created on.
	// Format: date
	CreatedAt strfmt.Date `json:"createdAt,omitempty"`

	// Uniquely identifies the state of the MoveTaskOrder object (but not the nested objects)
	//
	// It will change everytime the object is updated. Client should store the value.
	// Updates to this MoveTaskOrder will require that this eTag be passed in with the If-Match header.
	//
	ETag string `json:"eTag,omitempty"`

	// ID of the MoveTaskOrder object.
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Indicates this MoveTaskOrder is available for Prime API handling.
	//
	// In production, only MoveTaskOrders for which this is true will be available to the API.
	//
	IsAvailableToPrime *bool `json:"isAvailableToPrime,omitempty"`

	// Indicated this MoveTaskOrder has been canceled.
	IsCanceled *bool `json:"isCanceled,omitempty"`

	// MoveOrder associated with this MoveTaskOrder.
	// Required: true
	MoveOrder *MoveOrder `json:"moveOrder"`

	// ID of the MoveOrder object
	// Format: uuid
	MoveOrderID strfmt.UUID `json:"moveOrderID,omitempty"`

	// Array of MTOServiceItems associated with this MoveTaskOrder.
	MtoServiceItems []*MTOServiceItem `json:"mtoServiceItems"`

	// array of MTOShipments associated with the MoveTaskOrder.
	MtoShipments MTOShipments `json:"mtoShipments,omitempty"`

	// Array of PaymentRequests associated with this MoveTaskOrder.
	PaymentRequests PaymentRequests `json:"paymentRequests,omitempty"`

	// If the move is a PPM, this is the estimated weight in lbs.
	PpmEstimatedWeight int64 `json:"ppmEstimatedWeight,omitempty"`

	// If the move is a PPM, indicates whether it is full or partial.
	// Enum: [FULL PARTIAL]
	PpmType string `json:"ppmType,omitempty"`

	// Unique ID associated with this MoveOrder.
	//
	// No two MoveTaskOrders may have the same ID.
	// Attempting to create a MoveTaskOrder may fail if this referenceId has been used already.
	//
	ReferenceID string `json:"referenceId,omitempty"`

	// Date on which this MoveTaskOrder was last updated.
	// Format: date
	UpdatedAt strfmt.Date `json:"updatedAt,omitempty"`
}

// Validate validates this move task order
func (m *MoveTaskOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtoServiceItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtoShipments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePpmType(formats); err != nil {
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

func (m *MoveTaskOrder) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateMoveOrder(formats strfmt.Registry) error {

	if err := validate.Required("moveOrder", "body", m.MoveOrder); err != nil {
		return err
	}

	if m.MoveOrder != nil {
		if err := m.MoveOrder.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moveOrder")
			}
			return err
		}
	}

	return nil
}

func (m *MoveTaskOrder) validateMoveOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveOrderID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveOrderID", "body", "uuid", m.MoveOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateMtoServiceItems(formats strfmt.Registry) error {

	if swag.IsZero(m.MtoServiceItems) { // not required
		return nil
	}

	for i := 0; i < len(m.MtoServiceItems); i++ {
		if swag.IsZero(m.MtoServiceItems[i]) { // not required
			continue
		}

		if m.MtoServiceItems[i] != nil {
			if err := m.MtoServiceItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mtoServiceItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoveTaskOrder) validateMtoShipments(formats strfmt.Registry) error {

	if swag.IsZero(m.MtoShipments) { // not required
		return nil
	}

	if err := m.MtoShipments.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mtoShipments")
		}
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validatePaymentRequests(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentRequests) { // not required
		return nil
	}

	if err := m.PaymentRequests.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("paymentRequests")
		}
		return err
	}

	return nil
}

var moveTaskOrderTypePpmTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FULL","PARTIAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		moveTaskOrderTypePpmTypePropEnum = append(moveTaskOrderTypePpmTypePropEnum, v)
	}
}

const (

	// MoveTaskOrderPpmTypeFULL captures enum value "FULL"
	MoveTaskOrderPpmTypeFULL string = "FULL"

	// MoveTaskOrderPpmTypePARTIAL captures enum value "PARTIAL"
	MoveTaskOrderPpmTypePARTIAL string = "PARTIAL"
)

// prop value enum
func (m *MoveTaskOrder) validatePpmTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, moveTaskOrderTypePpmTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MoveTaskOrder) validatePpmType(formats strfmt.Registry) error {

	if swag.IsZero(m.PpmType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePpmTypeEnum("ppmType", "body", m.PpmType); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveTaskOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveTaskOrder) UnmarshalBinary(b []byte) error {
	var res MoveTaskOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
