// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MoveTaskOrder move task order
// swagger:model MoveTaskOrder
type MoveTaskOrder struct {

	// Indicates this MoveTaskOrder is available for Prime API handling.
	//
	// In production, only MoveTaskOrders for which this is set will be available to the API.
	//
	// Format: date-time
	AvailableToPrimeAt *strfmt.DateTime `json:"availableToPrimeAt,omitempty"`

	// ID associated with the contractor, in this case Prime
	//
	// Format: uuid
	ContractorID strfmt.UUID `json:"contractorID,omitempty"`

	// Date the MoveTaskOrder was created on.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

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

	// Indicated this MoveTaskOrder has been canceled.
	IsCanceled *bool `json:"isCanceled,omitempty"`

	// MoveOrder associated with this MoveTaskOrder.
	// Required: true
	MoveOrder *MoveOrder `json:"moveOrder"`

	// ID of the MoveOrder object
	// Format: uuid
	MoveOrderID strfmt.UUID `json:"moveOrderID,omitempty"`

	mtoServiceItemsField []MTOServiceItem

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
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// MtoServiceItems gets the mto service items of this base type
func (m *MoveTaskOrder) MtoServiceItems() []MTOServiceItem {
	return m.mtoServiceItemsField
}

// SetMtoServiceItems sets the mto service items of this base type
func (m *MoveTaskOrder) SetMtoServiceItems(val []MTOServiceItem) {
	m.mtoServiceItemsField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *MoveTaskOrder) UnmarshalJSON(raw []byte) error {
	var data struct {
		AvailableToPrimeAt *strfmt.DateTime `json:"availableToPrimeAt,omitempty"`

		ContractorID strfmt.UUID `json:"contractorID,omitempty"`

		CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

		ETag string `json:"eTag,omitempty"`

		ID strfmt.UUID `json:"id,omitempty"`

		IsCanceled *bool `json:"isCanceled,omitempty"`

		MoveOrder *MoveOrder `json:"moveOrder"`

		MoveOrderID strfmt.UUID `json:"moveOrderID,omitempty"`

		MtoServiceItems json.RawMessage `json:"mtoServiceItems"`

		MtoShipments MTOShipments `json:"mtoShipments,omitempty"`

		PaymentRequests PaymentRequests `json:"paymentRequests,omitempty"`

		PpmEstimatedWeight int64 `json:"ppmEstimatedWeight,omitempty"`

		PpmType string `json:"ppmType,omitempty"`

		ReferenceID string `json:"referenceId,omitempty"`

		UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propMtoServiceItems []MTOServiceItem
	if string(data.MtoServiceItems) != "null" {
		mtoServiceItems, err := UnmarshalMTOServiceItemSlice(bytes.NewBuffer(data.MtoServiceItems), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propMtoServiceItems = mtoServiceItems
	}

	var result MoveTaskOrder

	// availableToPrimeAt
	result.AvailableToPrimeAt = data.AvailableToPrimeAt

	// contractorID
	result.ContractorID = data.ContractorID

	// createdAt
	result.CreatedAt = data.CreatedAt

	// eTag
	result.ETag = data.ETag

	// id
	result.ID = data.ID

	// isCanceled
	result.IsCanceled = data.IsCanceled

	// moveOrder
	result.MoveOrder = data.MoveOrder

	// moveOrderID
	result.MoveOrderID = data.MoveOrderID

	// mtoServiceItems
	result.mtoServiceItemsField = propMtoServiceItems

	// mtoShipments
	result.MtoShipments = data.MtoShipments

	// paymentRequests
	result.PaymentRequests = data.PaymentRequests

	// ppmEstimatedWeight
	result.PpmEstimatedWeight = data.PpmEstimatedWeight

	// ppmType
	result.PpmType = data.PpmType

	// referenceId
	result.ReferenceID = data.ReferenceID

	// updatedAt
	result.UpdatedAt = data.UpdatedAt

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m MoveTaskOrder) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AvailableToPrimeAt *strfmt.DateTime `json:"availableToPrimeAt,omitempty"`

		ContractorID strfmt.UUID `json:"contractorID,omitempty"`

		CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

		ETag string `json:"eTag,omitempty"`

		ID strfmt.UUID `json:"id,omitempty"`

		IsCanceled *bool `json:"isCanceled,omitempty"`

		MoveOrder *MoveOrder `json:"moveOrder"`

		MoveOrderID strfmt.UUID `json:"moveOrderID,omitempty"`

		MtoShipments MTOShipments `json:"mtoShipments,omitempty"`

		PaymentRequests PaymentRequests `json:"paymentRequests,omitempty"`

		PpmEstimatedWeight int64 `json:"ppmEstimatedWeight,omitempty"`

		PpmType string `json:"ppmType,omitempty"`

		ReferenceID string `json:"referenceId,omitempty"`

		UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
	}{

		AvailableToPrimeAt: m.AvailableToPrimeAt,

		ContractorID: m.ContractorID,

		CreatedAt: m.CreatedAt,

		ETag: m.ETag,

		ID: m.ID,

		IsCanceled: m.IsCanceled,

		MoveOrder: m.MoveOrder,

		MoveOrderID: m.MoveOrderID,

		MtoShipments: m.MtoShipments,

		PaymentRequests: m.PaymentRequests,

		PpmEstimatedWeight: m.PpmEstimatedWeight,

		PpmType: m.PpmType,

		ReferenceID: m.ReferenceID,

		UpdatedAt: m.UpdatedAt,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		MtoServiceItems []MTOServiceItem `json:"mtoServiceItems"`
	}{

		MtoServiceItems: m.mtoServiceItemsField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this move task order
func (m *MoveTaskOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableToPrimeAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractorID(formats); err != nil {
		res = append(res, err)
	}

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

func (m *MoveTaskOrder) validateAvailableToPrimeAt(formats strfmt.Registry) error {

	if swag.IsZero(m.AvailableToPrimeAt) { // not required
		return nil
	}

	if err := validate.FormatOf("availableToPrimeAt", "body", "date-time", m.AvailableToPrimeAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateContractorID(formats strfmt.Registry) error {

	if swag.IsZero(m.ContractorID) { // not required
		return nil
	}

	if err := validate.FormatOf("contractorID", "body", "uuid", m.ContractorID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
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

	if swag.IsZero(m.MtoServiceItems()) { // not required
		return nil
	}

	for i := 0; i < len(m.MtoServiceItems()); i++ {

		if err := m.mtoServiceItemsField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mtoServiceItems" + "." + strconv.Itoa(i))
			}
			return err
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

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
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
