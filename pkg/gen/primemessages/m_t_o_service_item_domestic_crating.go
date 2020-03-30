// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOServiceItemDomesticCrating Describes a domestic crating/uncrating service item subtype of a MTOServiceItem
// swagger:model MTOServiceItemDomesticCrating
type MTOServiceItemDomesticCrating struct {
	eTagField string

	idField strfmt.UUID

	moveTaskOrderIdField strfmt.UUID

	mtoShipmentIdField strfmt.UUID

	reServiceIdField strfmt.UUID

	reServiceNameField string

	// crate
	// Required: true
	Crate *MTOServiceItemDimension `json:"crate"`

	// description
	// Required: true
	Description *string `json:"description"`

	// item
	// Required: true
	Item *MTOServiceItemDimension `json:"item"`

	// Describes available re service code for domestic crating
	// Required: true
	// Enum: [DCRT DUCRT]
	ReServiceCode *string `json:"reServiceCode"`
}

// ETag gets the e tag of this subtype
func (m *MTOServiceItemDomesticCrating) ETag() string {
	return m.eTagField
}

// SetETag sets the e tag of this subtype
func (m *MTOServiceItemDomesticCrating) SetETag(val string) {
	m.eTagField = val
}

// ID gets the id of this subtype
func (m *MTOServiceItemDomesticCrating) ID() strfmt.UUID {
	return m.idField
}

// SetID sets the id of this subtype
func (m *MTOServiceItemDomesticCrating) SetID(val strfmt.UUID) {
	m.idField = val
}

// ModelType gets the model type of this subtype
func (m *MTOServiceItemDomesticCrating) ModelType() MTOServiceItemModelType {
	return "MTOServiceItemDomesticCrating"
}

// SetModelType sets the model type of this subtype
func (m *MTOServiceItemDomesticCrating) SetModelType(val MTOServiceItemModelType) {

}

// MoveTaskOrderID gets the move task order ID of this subtype
func (m *MTOServiceItemDomesticCrating) MoveTaskOrderID() strfmt.UUID {
	return m.moveTaskOrderIdField
}

// SetMoveTaskOrderID sets the move task order ID of this subtype
func (m *MTOServiceItemDomesticCrating) SetMoveTaskOrderID(val strfmt.UUID) {
	m.moveTaskOrderIdField = val
}

// MtoShipmentID gets the mto shipment ID of this subtype
func (m *MTOServiceItemDomesticCrating) MtoShipmentID() strfmt.UUID {
	return m.mtoShipmentIdField
}

// SetMtoShipmentID sets the mto shipment ID of this subtype
func (m *MTOServiceItemDomesticCrating) SetMtoShipmentID(val strfmt.UUID) {
	m.mtoShipmentIdField = val
}

// ReServiceID gets the re service ID of this subtype
func (m *MTOServiceItemDomesticCrating) ReServiceID() strfmt.UUID {
	return m.reServiceIdField
}

// SetReServiceID sets the re service ID of this subtype
func (m *MTOServiceItemDomesticCrating) SetReServiceID(val strfmt.UUID) {
	m.reServiceIdField = val
}

// ReServiceName gets the re service name of this subtype
func (m *MTOServiceItemDomesticCrating) ReServiceName() string {
	return m.reServiceNameField
}

// SetReServiceName sets the re service name of this subtype
func (m *MTOServiceItemDomesticCrating) SetReServiceName(val string) {
	m.reServiceNameField = val
}

// Crate gets the crate of this subtype

// Description gets the description of this subtype

// Item gets the item of this subtype

// ReServiceCode gets the re service code of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *MTOServiceItemDomesticCrating) UnmarshalJSON(raw []byte) error {
	var data struct {

		// crate
		// Required: true
		Crate *MTOServiceItemDimension `json:"crate"`

		// description
		// Required: true
		Description *string `json:"description"`

		// item
		// Required: true
		Item *MTOServiceItemDimension `json:"item"`

		// Describes available re service code for domestic crating
		// Required: true
		// Enum: [DCRT DUCRT]
		ReServiceCode *string `json:"reServiceCode"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		ETag string `json:"eTag,omitempty"`

		ID strfmt.UUID `json:"id,omitempty"`

		ModelType MTOServiceItemModelType `json:"modelType"`

		MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

		MtoShipmentID strfmt.UUID `json:"mtoShipmentID,omitempty"`

		ReServiceID strfmt.UUID `json:"reServiceID,omitempty"`

		ReServiceName string `json:"reServiceName,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result MTOServiceItemDomesticCrating

	result.eTagField = base.ETag

	result.idField = base.ID

	if base.ModelType != result.ModelType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid modelType value: %q", base.ModelType)
	}

	result.moveTaskOrderIdField = base.MoveTaskOrderID

	result.mtoShipmentIdField = base.MtoShipmentID

	result.reServiceIdField = base.ReServiceID

	result.reServiceNameField = base.ReServiceName

	result.Crate = data.Crate

	result.Description = data.Description

	result.Item = data.Item

	result.ReServiceCode = data.ReServiceCode

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m MTOServiceItemDomesticCrating) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// crate
		// Required: true
		Crate *MTOServiceItemDimension `json:"crate"`

		// description
		// Required: true
		Description *string `json:"description"`

		// item
		// Required: true
		Item *MTOServiceItemDimension `json:"item"`

		// Describes available re service code for domestic crating
		// Required: true
		// Enum: [DCRT DUCRT]
		ReServiceCode *string `json:"reServiceCode"`
	}{

		Crate: m.Crate,

		Description: m.Description,

		Item: m.Item,

		ReServiceCode: m.ReServiceCode,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ETag string `json:"eTag,omitempty"`

		ID strfmt.UUID `json:"id,omitempty"`

		ModelType MTOServiceItemModelType `json:"modelType"`

		MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

		MtoShipmentID strfmt.UUID `json:"mtoShipmentID,omitempty"`

		ReServiceID strfmt.UUID `json:"reServiceID,omitempty"`

		ReServiceName string `json:"reServiceName,omitempty"`
	}{

		ETag: m.ETag(),

		ID: m.ID(),

		ModelType: m.ModelType(),

		MoveTaskOrderID: m.MoveTaskOrderID(),

		MtoShipmentID: m.MtoShipmentID(),

		ReServiceID: m.ReServiceID(),

		ReServiceName: m.ReServiceName(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this m t o service item domestic crating
func (m *MTOServiceItemDomesticCrating) Validate(formats strfmt.Registry) error {
	var res []error

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

	if err := m.validateCrate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReServiceCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MTOServiceItemDomesticCrating) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID()) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDomesticCrating) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrderID()) { // not required
		return nil
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDomesticCrating) validateMtoShipmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.MtoShipmentID()) { // not required
		return nil
	}

	if err := validate.FormatOf("mtoShipmentID", "body", "uuid", m.MtoShipmentID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDomesticCrating) validateReServiceID(formats strfmt.Registry) error {

	if swag.IsZero(m.ReServiceID()) { // not required
		return nil
	}

	if err := validate.FormatOf("reServiceID", "body", "uuid", m.ReServiceID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDomesticCrating) validateCrate(formats strfmt.Registry) error {

	if err := validate.Required("crate", "body", m.Crate); err != nil {
		return err
	}

	if m.Crate != nil {
		if err := m.Crate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("crate")
			}
			return err
		}
	}

	return nil
}

func (m *MTOServiceItemDomesticCrating) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDomesticCrating) validateItem(formats strfmt.Registry) error {

	if err := validate.Required("item", "body", m.Item); err != nil {
		return err
	}

	if m.Item != nil {
		if err := m.Item.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item")
			}
			return err
		}
	}

	return nil
}

var mTOServiceItemDomesticCratingTypeReServiceCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DCRT","DUCRT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mTOServiceItemDomesticCratingTypeReServiceCodePropEnum = append(mTOServiceItemDomesticCratingTypeReServiceCodePropEnum, v)
	}
}

// property enum
func (m *MTOServiceItemDomesticCrating) validateReServiceCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, mTOServiceItemDomesticCratingTypeReServiceCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MTOServiceItemDomesticCrating) validateReServiceCode(formats strfmt.Registry) error {

	if err := validate.Required("reServiceCode", "body", m.ReServiceCode); err != nil {
		return err
	}

	// value enum
	if err := m.validateReServiceCodeEnum("reServiceCode", "body", *m.ReServiceCode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOServiceItemDomesticCrating) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOServiceItemDomesticCrating) UnmarshalBinary(b []byte) error {
	var res MTOServiceItemDomesticCrating
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
