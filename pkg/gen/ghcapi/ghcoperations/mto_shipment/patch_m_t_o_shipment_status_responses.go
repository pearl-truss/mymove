// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// PatchMTOShipmentStatusOKCode is the HTTP code returned for type PatchMTOShipmentStatusOK
const PatchMTOShipmentStatusOKCode int = 200

/*PatchMTOShipmentStatusOK Successfully updated shipment

swagger:response patchMTOShipmentStatusOK
*/
type PatchMTOShipmentStatusOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.MTOShipmentWithEtag `json:"body,omitempty"`
}

// NewPatchMTOShipmentStatusOK creates PatchMTOShipmentStatusOK with default headers values
func NewPatchMTOShipmentStatusOK() *PatchMTOShipmentStatusOK {

	return &PatchMTOShipmentStatusOK{}
}

// WithPayload adds the payload to the patch m t o shipment status o k response
func (o *PatchMTOShipmentStatusOK) WithPayload(payload *ghcmessages.MTOShipmentWithEtag) *PatchMTOShipmentStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch m t o shipment status o k response
func (o *PatchMTOShipmentStatusOK) SetPayload(payload *ghcmessages.MTOShipmentWithEtag) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMTOShipmentStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchMTOShipmentStatusNotFoundCode is the HTTP code returned for type PatchMTOShipmentStatusNotFound
const PatchMTOShipmentStatusNotFoundCode int = 404

/*PatchMTOShipmentStatusNotFound The requested resource wasn't found

swagger:response patchMTOShipmentStatusNotFound
*/
type PatchMTOShipmentStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPatchMTOShipmentStatusNotFound creates PatchMTOShipmentStatusNotFound with default headers values
func NewPatchMTOShipmentStatusNotFound() *PatchMTOShipmentStatusNotFound {

	return &PatchMTOShipmentStatusNotFound{}
}

// WithPayload adds the payload to the patch m t o shipment status not found response
func (o *PatchMTOShipmentStatusNotFound) WithPayload(payload interface{}) *PatchMTOShipmentStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch m t o shipment status not found response
func (o *PatchMTOShipmentStatusNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMTOShipmentStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchMTOShipmentStatusConflictCode is the HTTP code returned for type PatchMTOShipmentStatusConflict
const PatchMTOShipmentStatusConflictCode int = 409

/*PatchMTOShipmentStatusConflict Conflict error

swagger:response patchMTOShipmentStatusConflict
*/
type PatchMTOShipmentStatusConflict struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPatchMTOShipmentStatusConflict creates PatchMTOShipmentStatusConflict with default headers values
func NewPatchMTOShipmentStatusConflict() *PatchMTOShipmentStatusConflict {

	return &PatchMTOShipmentStatusConflict{}
}

// WithPayload adds the payload to the patch m t o shipment status conflict response
func (o *PatchMTOShipmentStatusConflict) WithPayload(payload interface{}) *PatchMTOShipmentStatusConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch m t o shipment status conflict response
func (o *PatchMTOShipmentStatusConflict) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMTOShipmentStatusConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchMTOShipmentStatusPreconditionFailedCode is the HTTP code returned for type PatchMTOShipmentStatusPreconditionFailed
const PatchMTOShipmentStatusPreconditionFailedCode int = 412

/*PatchMTOShipmentStatusPreconditionFailed Precondition failed

swagger:response patchMTOShipmentStatusPreconditionFailed
*/
type PatchMTOShipmentStatusPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPatchMTOShipmentStatusPreconditionFailed creates PatchMTOShipmentStatusPreconditionFailed with default headers values
func NewPatchMTOShipmentStatusPreconditionFailed() *PatchMTOShipmentStatusPreconditionFailed {

	return &PatchMTOShipmentStatusPreconditionFailed{}
}

// WithPayload adds the payload to the patch m t o shipment status precondition failed response
func (o *PatchMTOShipmentStatusPreconditionFailed) WithPayload(payload interface{}) *PatchMTOShipmentStatusPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch m t o shipment status precondition failed response
func (o *PatchMTOShipmentStatusPreconditionFailed) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMTOShipmentStatusPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchMTOShipmentStatusUnprocessableEntityCode is the HTTP code returned for type PatchMTOShipmentStatusUnprocessableEntity
const PatchMTOShipmentStatusUnprocessableEntityCode int = 422

/*PatchMTOShipmentStatusUnprocessableEntity Validation error

swagger:response patchMTOShipmentStatusUnprocessableEntity
*/
type PatchMTOShipmentStatusUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.ValidationError `json:"body,omitempty"`
}

// NewPatchMTOShipmentStatusUnprocessableEntity creates PatchMTOShipmentStatusUnprocessableEntity with default headers values
func NewPatchMTOShipmentStatusUnprocessableEntity() *PatchMTOShipmentStatusUnprocessableEntity {

	return &PatchMTOShipmentStatusUnprocessableEntity{}
}

// WithPayload adds the payload to the patch m t o shipment status unprocessable entity response
func (o *PatchMTOShipmentStatusUnprocessableEntity) WithPayload(payload *ghcmessages.ValidationError) *PatchMTOShipmentStatusUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch m t o shipment status unprocessable entity response
func (o *PatchMTOShipmentStatusUnprocessableEntity) SetPayload(payload *ghcmessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMTOShipmentStatusUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchMTOShipmentStatusInternalServerErrorCode is the HTTP code returned for type PatchMTOShipmentStatusInternalServerError
const PatchMTOShipmentStatusInternalServerErrorCode int = 500

/*PatchMTOShipmentStatusInternalServerError A server error occurred

swagger:response patchMTOShipmentStatusInternalServerError
*/
type PatchMTOShipmentStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPatchMTOShipmentStatusInternalServerError creates PatchMTOShipmentStatusInternalServerError with default headers values
func NewPatchMTOShipmentStatusInternalServerError() *PatchMTOShipmentStatusInternalServerError {

	return &PatchMTOShipmentStatusInternalServerError{}
}

// WithPayload adds the payload to the patch m t o shipment status internal server error response
func (o *PatchMTOShipmentStatusInternalServerError) WithPayload(payload interface{}) *PatchMTOShipmentStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch m t o shipment status internal server error response
func (o *PatchMTOShipmentStatusInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMTOShipmentStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
