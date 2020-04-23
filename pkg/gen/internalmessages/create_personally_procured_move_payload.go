// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreatePersonallyProcuredMovePayload create personally procured move payload
// swagger:model CreatePersonallyProcuredMovePayload
type CreatePersonallyProcuredMovePayload struct {

	// ZIP code
	// Pattern: ^(\d{5}([\-]\d{4})?)$
	AdditionalPickupPostalCode *string `json:"additional_pickup_postal_code,omitempty"`

	// advance
	Advance *CreateReimbursement `json:"advance,omitempty"`

	// advance worksheet
	AdvanceWorksheet *DocumentPayload `json:"advance_worksheet,omitempty"`

	// How many days do you plan to put your stuff in storage?
	// Maximum: 90
	// Minimum: 0
	DaysInStorage *int64 `json:"days_in_storage,omitempty"`

	// ZIP code
	// Pattern: ^(\d{5}([\-]\d{4})?)$
	DestinationPostalCode *string `json:"destination_postal_code,omitempty"`

	// Estimated Storage Reimbursement
	EstimatedStorageReimbursement *string `json:"estimated_storage_reimbursement,omitempty"`

	// Do you have stuff at another pickup location?
	HasAdditionalPostalCode *bool `json:"has_additional_postal_code,omitempty"`

	// Has Pro-Gear
	// Enum: [NOT SURE YES NO]
	HasProGear *string `json:"has_pro_gear,omitempty"`

	// Has Pro-Gear Over Thousand Pounds
	// Enum: [NOT SURE YES NO]
	HasProGearOverThousand *string `json:"has_pro_gear_over_thousand,omitempty"`

	// Would you like an advance of up to 60% of your PPM incentive?
	HasRequestedAdvance bool `json:"has_requested_advance,omitempty"`

	// Will you put anything in storage?
	HasSit *bool `json:"has_sit,omitempty"`

	// Net Weight
	// Minimum: 1
	NetWeight *int64 `json:"net_weight,omitempty"`

	// When do you plan to move?
	// Format: date
	OriginalMoveDate *strfmt.Date `json:"original_move_date,omitempty"`

	// ZIP code
	// Pattern: ^(\d{5}([\-]\d{4})?)$
	PickupPostalCode *string `json:"pickup_postal_code,omitempty"`

	// size
	Size *TShirtSize `json:"size,omitempty"`

	// Weight Estimate
	// Minimum: 0
	WeightEstimate *int64 `json:"weight_estimate,omitempty"`
}

// Validate validates this create personally procured move payload
func (m *CreatePersonallyProcuredMovePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalPickupPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvanceWorksheet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDaysInStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasProGear(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasProGearOverThousand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalMoveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePickupPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeightEstimate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePersonallyProcuredMovePayload) validateAdditionalPickupPostalCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalPickupPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("additional_pickup_postal_code", "body", string(*m.AdditionalPickupPostalCode), `^(\d{5}([\-]\d{4})?)$`); err != nil {
		return err
	}

	return nil
}

func (m *CreatePersonallyProcuredMovePayload) validateAdvance(formats strfmt.Registry) error {

	if swag.IsZero(m.Advance) { // not required
		return nil
	}

	if m.Advance != nil {
		if err := m.Advance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advance")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePersonallyProcuredMovePayload) validateAdvanceWorksheet(formats strfmt.Registry) error {

	if swag.IsZero(m.AdvanceWorksheet) { // not required
		return nil
	}

	if m.AdvanceWorksheet != nil {
		if err := m.AdvanceWorksheet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advance_worksheet")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePersonallyProcuredMovePayload) validateDaysInStorage(formats strfmt.Registry) error {

	if swag.IsZero(m.DaysInStorage) { // not required
		return nil
	}

	if err := validate.MinimumInt("days_in_storage", "body", int64(*m.DaysInStorage), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("days_in_storage", "body", int64(*m.DaysInStorage), 90, false); err != nil {
		return err
	}

	return nil
}

func (m *CreatePersonallyProcuredMovePayload) validateDestinationPostalCode(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("destination_postal_code", "body", string(*m.DestinationPostalCode), `^(\d{5}([\-]\d{4})?)$`); err != nil {
		return err
	}

	return nil
}

var createPersonallyProcuredMovePayloadTypeHasProGearPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT SURE","YES","NO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createPersonallyProcuredMovePayloadTypeHasProGearPropEnum = append(createPersonallyProcuredMovePayloadTypeHasProGearPropEnum, v)
	}
}

const (

	// CreatePersonallyProcuredMovePayloadHasProGearNOTSURE captures enum value "NOT SURE"
	CreatePersonallyProcuredMovePayloadHasProGearNOTSURE string = "NOT SURE"

	// CreatePersonallyProcuredMovePayloadHasProGearYES captures enum value "YES"
	CreatePersonallyProcuredMovePayloadHasProGearYES string = "YES"

	// CreatePersonallyProcuredMovePayloadHasProGearNO captures enum value "NO"
	CreatePersonallyProcuredMovePayloadHasProGearNO string = "NO"
)

// prop value enum
func (m *CreatePersonallyProcuredMovePayload) validateHasProGearEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createPersonallyProcuredMovePayloadTypeHasProGearPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreatePersonallyProcuredMovePayload) validateHasProGear(formats strfmt.Registry) error {

	if swag.IsZero(m.HasProGear) { // not required
		return nil
	}

	// value enum
	if err := m.validateHasProGearEnum("has_pro_gear", "body", *m.HasProGear); err != nil {
		return err
	}

	return nil
}

var createPersonallyProcuredMovePayloadTypeHasProGearOverThousandPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT SURE","YES","NO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createPersonallyProcuredMovePayloadTypeHasProGearOverThousandPropEnum = append(createPersonallyProcuredMovePayloadTypeHasProGearOverThousandPropEnum, v)
	}
}

const (

	// CreatePersonallyProcuredMovePayloadHasProGearOverThousandNOTSURE captures enum value "NOT SURE"
	CreatePersonallyProcuredMovePayloadHasProGearOverThousandNOTSURE string = "NOT SURE"

	// CreatePersonallyProcuredMovePayloadHasProGearOverThousandYES captures enum value "YES"
	CreatePersonallyProcuredMovePayloadHasProGearOverThousandYES string = "YES"

	// CreatePersonallyProcuredMovePayloadHasProGearOverThousandNO captures enum value "NO"
	CreatePersonallyProcuredMovePayloadHasProGearOverThousandNO string = "NO"
)

// prop value enum
func (m *CreatePersonallyProcuredMovePayload) validateHasProGearOverThousandEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createPersonallyProcuredMovePayloadTypeHasProGearOverThousandPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreatePersonallyProcuredMovePayload) validateHasProGearOverThousand(formats strfmt.Registry) error {

	if swag.IsZero(m.HasProGearOverThousand) { // not required
		return nil
	}

	// value enum
	if err := m.validateHasProGearOverThousandEnum("has_pro_gear_over_thousand", "body", *m.HasProGearOverThousand); err != nil {
		return err
	}

	return nil
}

func (m *CreatePersonallyProcuredMovePayload) validateNetWeight(formats strfmt.Registry) error {

	if swag.IsZero(m.NetWeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("net_weight", "body", int64(*m.NetWeight), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *CreatePersonallyProcuredMovePayload) validateOriginalMoveDate(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginalMoveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("original_move_date", "body", "date", m.OriginalMoveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreatePersonallyProcuredMovePayload) validatePickupPostalCode(formats strfmt.Registry) error {

	if swag.IsZero(m.PickupPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("pickup_postal_code", "body", string(*m.PickupPostalCode), `^(\d{5}([\-]\d{4})?)$`); err != nil {
		return err
	}

	return nil
}

func (m *CreatePersonallyProcuredMovePayload) validateSize(formats strfmt.Registry) error {

	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if m.Size != nil {
		if err := m.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePersonallyProcuredMovePayload) validateWeightEstimate(formats strfmt.Registry) error {

	if swag.IsZero(m.WeightEstimate) { // not required
		return nil
	}

	if err := validate.MinimumInt("weight_estimate", "body", int64(*m.WeightEstimate), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreatePersonallyProcuredMovePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePersonallyProcuredMovePayload) UnmarshalBinary(b []byte) error {
	var res CreatePersonallyProcuredMovePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
