// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdateMTOShipmentParams creates a new UpdateMTOShipmentParams object
// no default values defined in spec.
func NewUpdateMTOShipmentParams() UpdateMTOShipmentParams {

	return UpdateMTOShipmentParams{}
}

// UpdateMTOShipmentParams contains all the bound params for the update m t o shipment operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateMTOShipment
type UpdateMTOShipmentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	IfUnmodifiedSince string
	/*
	  In: body
	*/
	Body UpdateMTOShipmentBody
	/*
	  Required: true
	  In: path
	*/
	MoveTaskOrderID strfmt.UUID
	/*
	  Required: true
	  In: path
	*/
	MtoShipmentID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateMTOShipmentParams() beforehand.
func (o *UpdateMTOShipmentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindIfUnmodifiedSince(r.Header[http.CanonicalHeaderKey("If-Unmodified-Since")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body UpdateMTOShipmentBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = body
			}
		}
	}
	rMoveTaskOrderID, rhkMoveTaskOrderID, _ := route.Params.GetOK("moveTaskOrderID")
	if err := o.bindMoveTaskOrderID(rMoveTaskOrderID, rhkMoveTaskOrderID, route.Formats); err != nil {
		res = append(res, err)
	}

	rMtoShipmentID, rhkMtoShipmentID, _ := route.Params.GetOK("mtoShipmentID")
	if err := o.bindMtoShipmentID(rMtoShipmentID, rhkMtoShipmentID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIfUnmodifiedSince binds and validates parameter IfUnmodifiedSince from header.
func (o *UpdateMTOShipmentParams) bindIfUnmodifiedSince(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("If-Unmodified-Since", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("If-Unmodified-Since", "header", raw); err != nil {
		return err
	}

	o.IfUnmodifiedSince = raw

	return nil
}

// bindMoveTaskOrderID binds and validates parameter MoveTaskOrderID from path.
func (o *UpdateMTOShipmentParams) bindMoveTaskOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("moveTaskOrderID", "path", "strfmt.UUID", raw)
	}
	o.MoveTaskOrderID = *(value.(*strfmt.UUID))

	if err := o.validateMoveTaskOrderID(formats); err != nil {
		return err
	}

	return nil
}

// validateMoveTaskOrderID carries on validations for parameter MoveTaskOrderID
func (o *UpdateMTOShipmentParams) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if err := validate.FormatOf("moveTaskOrderID", "path", "uuid", o.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindMtoShipmentID binds and validates parameter MtoShipmentID from path.
func (o *UpdateMTOShipmentParams) bindMtoShipmentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("mtoShipmentID", "path", "strfmt.UUID", raw)
	}
	o.MtoShipmentID = *(value.(*strfmt.UUID))

	if err := o.validateMtoShipmentID(formats); err != nil {
		return err
	}

	return nil
}

// validateMtoShipmentID carries on validations for parameter MtoShipmentID
func (o *UpdateMTOShipmentParams) validateMtoShipmentID(formats strfmt.Registry) error {

	if err := validate.FormatOf("mtoShipmentID", "path", "uuid", o.MtoShipmentID.String(), formats); err != nil {
		return err
	}
	return nil
}
