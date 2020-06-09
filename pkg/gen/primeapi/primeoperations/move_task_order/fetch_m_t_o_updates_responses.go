// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// FetchMTOUpdatesOKCode is the HTTP code returned for type FetchMTOUpdatesOK
const FetchMTOUpdatesOKCode int = 200

/*FetchMTOUpdatesOK Successfully retrieved move task orders where `availableToPrimeAt` has been set.

swagger:response fetchMTOUpdatesOK
*/
type FetchMTOUpdatesOK struct {

	/*
	  In: Body
	*/
	Payload primemessages.MoveTaskOrders `json:"body,omitempty"`
}

// NewFetchMTOUpdatesOK creates FetchMTOUpdatesOK with default headers values
func NewFetchMTOUpdatesOK() *FetchMTOUpdatesOK {

	return &FetchMTOUpdatesOK{}
}

// WithPayload adds the payload to the fetch m t o updates o k response
func (o *FetchMTOUpdatesOK) WithPayload(payload primemessages.MoveTaskOrders) *FetchMTOUpdatesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fetch m t o updates o k response
func (o *FetchMTOUpdatesOK) SetPayload(payload primemessages.MoveTaskOrders) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FetchMTOUpdatesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = primemessages.MoveTaskOrders{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// FetchMTOUpdatesBadRequestCode is the HTTP code returned for type FetchMTOUpdatesBadRequest
const FetchMTOUpdatesBadRequestCode int = 400

/*FetchMTOUpdatesBadRequest The request payload is invalid.

swagger:response fetchMTOUpdatesBadRequest
*/
type FetchMTOUpdatesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewFetchMTOUpdatesBadRequest creates FetchMTOUpdatesBadRequest with default headers values
func NewFetchMTOUpdatesBadRequest() *FetchMTOUpdatesBadRequest {

	return &FetchMTOUpdatesBadRequest{}
}

// WithPayload adds the payload to the fetch m t o updates bad request response
func (o *FetchMTOUpdatesBadRequest) WithPayload(payload *primemessages.ClientError) *FetchMTOUpdatesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fetch m t o updates bad request response
func (o *FetchMTOUpdatesBadRequest) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FetchMTOUpdatesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FetchMTOUpdatesUnauthorizedCode is the HTTP code returned for type FetchMTOUpdatesUnauthorized
const FetchMTOUpdatesUnauthorizedCode int = 401

/*FetchMTOUpdatesUnauthorized The request was denied.

swagger:response fetchMTOUpdatesUnauthorized
*/
type FetchMTOUpdatesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewFetchMTOUpdatesUnauthorized creates FetchMTOUpdatesUnauthorized with default headers values
func NewFetchMTOUpdatesUnauthorized() *FetchMTOUpdatesUnauthorized {

	return &FetchMTOUpdatesUnauthorized{}
}

// WithPayload adds the payload to the fetch m t o updates unauthorized response
func (o *FetchMTOUpdatesUnauthorized) WithPayload(payload *primemessages.ClientError) *FetchMTOUpdatesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fetch m t o updates unauthorized response
func (o *FetchMTOUpdatesUnauthorized) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FetchMTOUpdatesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FetchMTOUpdatesForbiddenCode is the HTTP code returned for type FetchMTOUpdatesForbidden
const FetchMTOUpdatesForbiddenCode int = 403

/*FetchMTOUpdatesForbidden The request was denied.

swagger:response fetchMTOUpdatesForbidden
*/
type FetchMTOUpdatesForbidden struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewFetchMTOUpdatesForbidden creates FetchMTOUpdatesForbidden with default headers values
func NewFetchMTOUpdatesForbidden() *FetchMTOUpdatesForbidden {

	return &FetchMTOUpdatesForbidden{}
}

// WithPayload adds the payload to the fetch m t o updates forbidden response
func (o *FetchMTOUpdatesForbidden) WithPayload(payload *primemessages.ClientError) *FetchMTOUpdatesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fetch m t o updates forbidden response
func (o *FetchMTOUpdatesForbidden) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FetchMTOUpdatesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FetchMTOUpdatesNotFoundCode is the HTTP code returned for type FetchMTOUpdatesNotFound
const FetchMTOUpdatesNotFoundCode int = 404

/*FetchMTOUpdatesNotFound The requested resource wasn't found.

swagger:response fetchMTOUpdatesNotFound
*/
type FetchMTOUpdatesNotFound struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewFetchMTOUpdatesNotFound creates FetchMTOUpdatesNotFound with default headers values
func NewFetchMTOUpdatesNotFound() *FetchMTOUpdatesNotFound {

	return &FetchMTOUpdatesNotFound{}
}

// WithPayload adds the payload to the fetch m t o updates not found response
func (o *FetchMTOUpdatesNotFound) WithPayload(payload *primemessages.ClientError) *FetchMTOUpdatesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fetch m t o updates not found response
func (o *FetchMTOUpdatesNotFound) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FetchMTOUpdatesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FetchMTOUpdatesInternalServerErrorCode is the HTTP code returned for type FetchMTOUpdatesInternalServerError
const FetchMTOUpdatesInternalServerErrorCode int = 500

/*FetchMTOUpdatesInternalServerError A server error occurred.

swagger:response fetchMTOUpdatesInternalServerError
*/
type FetchMTOUpdatesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *primemessages.Error `json:"body,omitempty"`
}

// NewFetchMTOUpdatesInternalServerError creates FetchMTOUpdatesInternalServerError with default headers values
func NewFetchMTOUpdatesInternalServerError() *FetchMTOUpdatesInternalServerError {

	return &FetchMTOUpdatesInternalServerError{}
}

// WithPayload adds the payload to the fetch m t o updates internal server error response
func (o *FetchMTOUpdatesInternalServerError) WithPayload(payload *primemessages.Error) *FetchMTOUpdatesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fetch m t o updates internal server error response
func (o *FetchMTOUpdatesInternalServerError) SetPayload(payload *primemessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FetchMTOUpdatesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
