// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	internalmessages "github.com/transcom/mymove/pkg/gen/internalmessages"
)

// ShowPPMEstimateOKCode is the HTTP code returned for type ShowPPMEstimateOK
const ShowPPMEstimateOKCode int = 200

/*ShowPPMEstimateOK Made estimate of PPM cost range

swagger:response showPPMEstimateOK
*/
type ShowPPMEstimateOK struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.PPMEstimateRange `json:"body,omitempty"`
}

// NewShowPPMEstimateOK creates ShowPPMEstimateOK with default headers values
func NewShowPPMEstimateOK() *ShowPPMEstimateOK {

	return &ShowPPMEstimateOK{}
}

// WithPayload adds the payload to the show p p m estimate o k response
func (o *ShowPPMEstimateOK) WithPayload(payload *internalmessages.PPMEstimateRange) *ShowPPMEstimateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show p p m estimate o k response
func (o *ShowPPMEstimateOK) SetPayload(payload *internalmessages.PPMEstimateRange) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowPPMEstimateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowPPMEstimateBadRequestCode is the HTTP code returned for type ShowPPMEstimateBadRequest
const ShowPPMEstimateBadRequestCode int = 400

/*ShowPPMEstimateBadRequest invalid request

swagger:response showPPMEstimateBadRequest
*/
type ShowPPMEstimateBadRequest struct {
}

// NewShowPPMEstimateBadRequest creates ShowPPMEstimateBadRequest with default headers values
func NewShowPPMEstimateBadRequest() *ShowPPMEstimateBadRequest {

	return &ShowPPMEstimateBadRequest{}
}

// WriteResponse to the client
func (o *ShowPPMEstimateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ShowPPMEstimateUnauthorizedCode is the HTTP code returned for type ShowPPMEstimateUnauthorized
const ShowPPMEstimateUnauthorizedCode int = 401

/*ShowPPMEstimateUnauthorized request requires user authentication

swagger:response showPPMEstimateUnauthorized
*/
type ShowPPMEstimateUnauthorized struct {
}

// NewShowPPMEstimateUnauthorized creates ShowPPMEstimateUnauthorized with default headers values
func NewShowPPMEstimateUnauthorized() *ShowPPMEstimateUnauthorized {

	return &ShowPPMEstimateUnauthorized{}
}

// WriteResponse to the client
func (o *ShowPPMEstimateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ShowPPMEstimateForbiddenCode is the HTTP code returned for type ShowPPMEstimateForbidden
const ShowPPMEstimateForbiddenCode int = 403

/*ShowPPMEstimateForbidden user is not authorized

swagger:response showPPMEstimateForbidden
*/
type ShowPPMEstimateForbidden struct {
}

// NewShowPPMEstimateForbidden creates ShowPPMEstimateForbidden with default headers values
func NewShowPPMEstimateForbidden() *ShowPPMEstimateForbidden {

	return &ShowPPMEstimateForbidden{}
}

// WriteResponse to the client
func (o *ShowPPMEstimateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ShowPPMEstimateNotFoundCode is the HTTP code returned for type ShowPPMEstimateNotFound
const ShowPPMEstimateNotFoundCode int = 404

/*ShowPPMEstimateNotFound ppm discount not found for provided postal codes and original move date

swagger:response showPPMEstimateNotFound
*/
type ShowPPMEstimateNotFound struct {
}

// NewShowPPMEstimateNotFound creates ShowPPMEstimateNotFound with default headers values
func NewShowPPMEstimateNotFound() *ShowPPMEstimateNotFound {

	return &ShowPPMEstimateNotFound{}
}

// WriteResponse to the client
func (o *ShowPPMEstimateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ShowPPMEstimateUnprocessableEntityCode is the HTTP code returned for type ShowPPMEstimateUnprocessableEntity
const ShowPPMEstimateUnprocessableEntityCode int = 422

/*ShowPPMEstimateUnprocessableEntity cannot process request with given information

swagger:response showPPMEstimateUnprocessableEntity
*/
type ShowPPMEstimateUnprocessableEntity struct {
}

// NewShowPPMEstimateUnprocessableEntity creates ShowPPMEstimateUnprocessableEntity with default headers values
func NewShowPPMEstimateUnprocessableEntity() *ShowPPMEstimateUnprocessableEntity {

	return &ShowPPMEstimateUnprocessableEntity{}
}

// WriteResponse to the client
func (o *ShowPPMEstimateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(422)
}

// ShowPPMEstimateInternalServerErrorCode is the HTTP code returned for type ShowPPMEstimateInternalServerError
const ShowPPMEstimateInternalServerErrorCode int = 500

/*ShowPPMEstimateInternalServerError internal server error

swagger:response showPPMEstimateInternalServerError
*/
type ShowPPMEstimateInternalServerError struct {
}

// NewShowPPMEstimateInternalServerError creates ShowPPMEstimateInternalServerError with default headers values
func NewShowPPMEstimateInternalServerError() *ShowPPMEstimateInternalServerError {

	return &ShowPPMEstimateInternalServerError{}
}

// WriteResponse to the client
func (o *ShowPPMEstimateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}