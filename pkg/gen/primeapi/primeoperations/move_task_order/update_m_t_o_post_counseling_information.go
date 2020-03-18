// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// UpdateMTOPostCounselingInformationHandlerFunc turns a function with the right signature into a update m t o post counseling information handler
type UpdateMTOPostCounselingInformationHandlerFunc func(UpdateMTOPostCounselingInformationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateMTOPostCounselingInformationHandlerFunc) Handle(params UpdateMTOPostCounselingInformationParams) middleware.Responder {
	return fn(params)
}

// UpdateMTOPostCounselingInformationHandler interface for that can handle valid update m t o post counseling information params
type UpdateMTOPostCounselingInformationHandler interface {
	Handle(UpdateMTOPostCounselingInformationParams) middleware.Responder
}

// NewUpdateMTOPostCounselingInformation creates a new http.Handler for the update m t o post counseling information operation
func NewUpdateMTOPostCounselingInformation(ctx *middleware.Context, handler UpdateMTOPostCounselingInformationHandler) *UpdateMTOPostCounselingInformation {
	return &UpdateMTOPostCounselingInformation{Context: ctx, Handler: handler}
}

/*UpdateMTOPostCounselingInformation swagger:route PATCH /move-task-orders/{moveTaskOrderID}/post-counseling-info moveTaskOrder prime updateMTOPostCounselingInformation

Updates move task order's post counseling information

Updates move task order's post counseling information

*/
type UpdateMTOPostCounselingInformation struct {
	Context *middleware.Context
	Handler UpdateMTOPostCounselingInformationHandler
}

func (o *UpdateMTOPostCounselingInformation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateMTOPostCounselingInformationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateMTOPostCounselingInformationBody update m t o post counseling information body
// swagger:model UpdateMTOPostCounselingInformationBody
type UpdateMTOPostCounselingInformationBody struct {

	// ppm estimated weight
	PpmEstimatedWeight int64 `json:"ppmEstimatedWeight,omitempty"`

	// ppm type
	// Enum: [FULL PARTIAL]
	PpmType string `json:"ppmType,omitempty"`
}

// Validate validates this update m t o post counseling information body
func (o *UpdateMTOPostCounselingInformationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePpmType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateMTOPostCounselingInformationBodyTypePpmTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FULL","PARTIAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateMTOPostCounselingInformationBodyTypePpmTypePropEnum = append(updateMTOPostCounselingInformationBodyTypePpmTypePropEnum, v)
	}
}

const (

	// UpdateMTOPostCounselingInformationBodyPpmTypeFULL captures enum value "FULL"
	UpdateMTOPostCounselingInformationBodyPpmTypeFULL string = "FULL"

	// UpdateMTOPostCounselingInformationBodyPpmTypePARTIAL captures enum value "PARTIAL"
	UpdateMTOPostCounselingInformationBodyPpmTypePARTIAL string = "PARTIAL"
)

// prop value enum
func (o *UpdateMTOPostCounselingInformationBody) validatePpmTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateMTOPostCounselingInformationBodyTypePpmTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *UpdateMTOPostCounselingInformationBody) validatePpmType(formats strfmt.Registry) error {

	if swag.IsZero(o.PpmType) { // not required
		return nil
	}

	// value enum
	if err := o.validatePpmTypeEnum("body"+"."+"ppmType", "body", o.PpmType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateMTOPostCounselingInformationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateMTOPostCounselingInformationBody) UnmarshalBinary(b []byte) error {
	var res UpdateMTOPostCounselingInformationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
