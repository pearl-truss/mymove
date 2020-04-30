// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// UpdateMTOShipmentOKCode is the HTTP code returned for type UpdateMTOShipmentOK
const UpdateMTOShipmentOKCode int = 200

/*UpdateMTOShipmentOK updated instance of mto shipment

swagger:response updateMTOShipmentOK
*/
type UpdateMTOShipmentOK struct {

	/*
	  In: Body
	*/
	Payload *primemessages.MTOShipment `json:"body,omitempty"`
}

// NewUpdateMTOShipmentOK creates UpdateMTOShipmentOK with default headers values
func NewUpdateMTOShipmentOK() *UpdateMTOShipmentOK {

	return &UpdateMTOShipmentOK{}
}

// WithPayload adds the payload to the update m t o shipment o k response
func (o *UpdateMTOShipmentOK) WithPayload(payload *primemessages.MTOShipment) *UpdateMTOShipmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment o k response
func (o *UpdateMTOShipmentOK) SetPayload(payload *primemessages.MTOShipment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentBadRequestCode is the HTTP code returned for type UpdateMTOShipmentBadRequest
const UpdateMTOShipmentBadRequestCode int = 400

/*UpdateMTOShipmentBadRequest invalid request

swagger:response updateMTOShipmentBadRequest
*/
type UpdateMTOShipmentBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOShipmentBadRequest creates UpdateMTOShipmentBadRequest with default headers values
func NewUpdateMTOShipmentBadRequest() *UpdateMTOShipmentBadRequest {

	return &UpdateMTOShipmentBadRequest{}
}

// WithPayload adds the payload to the update m t o shipment bad request response
func (o *UpdateMTOShipmentBadRequest) WithPayload(payload interface{}) *UpdateMTOShipmentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment bad request response
func (o *UpdateMTOShipmentBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOShipmentUnauthorizedCode is the HTTP code returned for type UpdateMTOShipmentUnauthorized
const UpdateMTOShipmentUnauthorizedCode int = 401

/*UpdateMTOShipmentUnauthorized The request was denied

swagger:response updateMTOShipmentUnauthorized
*/
type UpdateMTOShipmentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOShipmentUnauthorized creates UpdateMTOShipmentUnauthorized with default headers values
func NewUpdateMTOShipmentUnauthorized() *UpdateMTOShipmentUnauthorized {

	return &UpdateMTOShipmentUnauthorized{}
}

// WithPayload adds the payload to the update m t o shipment unauthorized response
func (o *UpdateMTOShipmentUnauthorized) WithPayload(payload interface{}) *UpdateMTOShipmentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment unauthorized response
func (o *UpdateMTOShipmentUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOShipmentForbiddenCode is the HTTP code returned for type UpdateMTOShipmentForbidden
const UpdateMTOShipmentForbiddenCode int = 403

/*UpdateMTOShipmentForbidden The request was denied

swagger:response updateMTOShipmentForbidden
*/
type UpdateMTOShipmentForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOShipmentForbidden creates UpdateMTOShipmentForbidden with default headers values
func NewUpdateMTOShipmentForbidden() *UpdateMTOShipmentForbidden {

	return &UpdateMTOShipmentForbidden{}
}

// WithPayload adds the payload to the update m t o shipment forbidden response
func (o *UpdateMTOShipmentForbidden) WithPayload(payload interface{}) *UpdateMTOShipmentForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment forbidden response
func (o *UpdateMTOShipmentForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOShipmentNotFoundCode is the HTTP code returned for type UpdateMTOShipmentNotFound
const UpdateMTOShipmentNotFoundCode int = 404

/*UpdateMTOShipmentNotFound The requested resource wasn't found

swagger:response updateMTOShipmentNotFound
*/
type UpdateMTOShipmentNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOShipmentNotFound creates UpdateMTOShipmentNotFound with default headers values
func NewUpdateMTOShipmentNotFound() *UpdateMTOShipmentNotFound {

	return &UpdateMTOShipmentNotFound{}
}

// WithPayload adds the payload to the update m t o shipment not found response
func (o *UpdateMTOShipmentNotFound) WithPayload(payload interface{}) *UpdateMTOShipmentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment not found response
func (o *UpdateMTOShipmentNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMTOShipmentPreconditionFailedCode is the HTTP code returned for type UpdateMTOShipmentPreconditionFailed
const UpdateMTOShipmentPreconditionFailedCode int = 412

/*UpdateMTOShipmentPreconditionFailed Precondition failed, likely due to a stale eTag (If-Match). Fetch the request again to get the updated eTag value.

swagger:response updateMTOShipmentPreconditionFailed
*/
type UpdateMTOShipmentPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *primemessages.Error `json:"body,omitempty"`
}

// NewUpdateMTOShipmentPreconditionFailed creates UpdateMTOShipmentPreconditionFailed with default headers values
func NewUpdateMTOShipmentPreconditionFailed() *UpdateMTOShipmentPreconditionFailed {

	return &UpdateMTOShipmentPreconditionFailed{}
}

// WithPayload adds the payload to the update m t o shipment precondition failed response
func (o *UpdateMTOShipmentPreconditionFailed) WithPayload(payload *primemessages.Error) *UpdateMTOShipmentPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment precondition failed response
func (o *UpdateMTOShipmentPreconditionFailed) SetPayload(payload *primemessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMTOShipmentInternalServerErrorCode is the HTTP code returned for type UpdateMTOShipmentInternalServerError
const UpdateMTOShipmentInternalServerErrorCode int = 500

/*UpdateMTOShipmentInternalServerError internal server error

swagger:response updateMTOShipmentInternalServerError
*/
type UpdateMTOShipmentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMTOShipmentInternalServerError creates UpdateMTOShipmentInternalServerError with default headers values
func NewUpdateMTOShipmentInternalServerError() *UpdateMTOShipmentInternalServerError {

	return &UpdateMTOShipmentInternalServerError{}
}

// WithPayload adds the payload to the update m t o shipment internal server error response
func (o *UpdateMTOShipmentInternalServerError) WithPayload(payload interface{}) *UpdateMTOShipmentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update m t o shipment internal server error response
func (o *UpdateMTOShipmentInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMTOShipmentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
