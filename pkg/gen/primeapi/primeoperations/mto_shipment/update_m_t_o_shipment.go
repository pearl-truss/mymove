// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// UpdateMTOShipmentHandlerFunc turns a function with the right signature into a update m t o shipment handler
type UpdateMTOShipmentHandlerFunc func(UpdateMTOShipmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateMTOShipmentHandlerFunc) Handle(params UpdateMTOShipmentParams) middleware.Responder {
	return fn(params)
}

// UpdateMTOShipmentHandler interface for that can handle valid update m t o shipment params
type UpdateMTOShipmentHandler interface {
	Handle(UpdateMTOShipmentParams) middleware.Responder
}

// NewUpdateMTOShipment creates a new http.Handler for the update m t o shipment operation
func NewUpdateMTOShipment(ctx *middleware.Context, handler UpdateMTOShipmentHandler) *UpdateMTOShipment {
	return &UpdateMTOShipment{Context: ctx, Handler: handler}
}

/*UpdateMTOShipment swagger:route PUT /move_task_orders/{moveTaskOrderID}/mto_shipments/{mtoShipmentID} mtoShipment prime updateMTOShipment

Updates mto shipment

*/
type UpdateMTOShipment struct {
	Context *middleware.Context
	Handler UpdateMTOShipmentHandler
}

func (o *UpdateMTOShipment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateMTOShipmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateMTOShipmentBody update m t o shipment body
// swagger:model UpdateMTOShipmentBody
type UpdateMTOShipmentBody struct {

	// customer remarks
	CustomerRemarks string `json:"customerRemarks,omitempty"`

	// delivery address
	DeliveryAddress *primemessages.Address `json:"deliveryAddress,omitempty"`

	// pickup address
	PickupAddress *primemessages.Address `json:"pickupAddress,omitempty"`

	// requested pickup date
	// Format: date
	RequestedPickupDate strfmt.Date `json:"requestedPickupDate,omitempty"`

	// scheduled pickup date
	// Format: date
	ScheduledPickupDate strfmt.Date `json:"scheduledPickupDate,omitempty"`

	// secondary delivery address
	SecondaryDeliveryAddress *primemessages.Address `json:"secondaryDeliveryAddress,omitempty"`

	// secondary pickup address
	SecondaryPickupAddress *primemessages.Address `json:"secondaryPickupAddress,omitempty"`

	// shipment type
	// Enum: [HHG INTERNATIONAL_HHG INTERNATIONAL_UB]
	ShipmentType interface{} `json:"shipmentType,omitempty"`
}

// Validate validates this update m t o shipment body
func (o *UpdateMTOShipmentBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDeliveryAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePickupAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRequestedPickupDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScheduledPickupDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecondaryDeliveryAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecondaryPickupAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateMTOShipmentBody) validateDeliveryAddress(formats strfmt.Registry) error {

	if swag.IsZero(o.DeliveryAddress) { // not required
		return nil
	}

	if o.DeliveryAddress != nil {
		if err := o.DeliveryAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "deliveryAddress")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateMTOShipmentBody) validatePickupAddress(formats strfmt.Registry) error {

	if swag.IsZero(o.PickupAddress) { // not required
		return nil
	}

	if o.PickupAddress != nil {
		if err := o.PickupAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pickupAddress")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateMTOShipmentBody) validateRequestedPickupDate(formats strfmt.Registry) error {

	if swag.IsZero(o.RequestedPickupDate) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"requestedPickupDate", "body", "date", o.RequestedPickupDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UpdateMTOShipmentBody) validateScheduledPickupDate(formats strfmt.Registry) error {

	if swag.IsZero(o.ScheduledPickupDate) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"scheduledPickupDate", "body", "date", o.ScheduledPickupDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *UpdateMTOShipmentBody) validateSecondaryDeliveryAddress(formats strfmt.Registry) error {

	if swag.IsZero(o.SecondaryDeliveryAddress) { // not required
		return nil
	}

	if o.SecondaryDeliveryAddress != nil {
		if err := o.SecondaryDeliveryAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "secondaryDeliveryAddress")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateMTOShipmentBody) validateSecondaryPickupAddress(formats strfmt.Registry) error {

	if swag.IsZero(o.SecondaryPickupAddress) { // not required
		return nil
	}

	if o.SecondaryPickupAddress != nil {
		if err := o.SecondaryPickupAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "secondaryPickupAddress")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateMTOShipmentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateMTOShipmentBody) UnmarshalBinary(b []byte) error {
	var res UpdateMTOShipmentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
