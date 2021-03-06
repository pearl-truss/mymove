// Code generated by go-swagger; DO NOT EDIT.

package move_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// ListMoveOrdersOKCode is the HTTP code returned for type ListMoveOrdersOK
const ListMoveOrdersOKCode int = 200

/*ListMoveOrdersOK Successfully retrieved all move orders

swagger:response listMoveOrdersOK
*/
type ListMoveOrdersOK struct {

	/*
	  In: Body
	*/
	Payload ghcmessages.MoveOrders `json:"body,omitempty"`
}

// NewListMoveOrdersOK creates ListMoveOrdersOK with default headers values
func NewListMoveOrdersOK() *ListMoveOrdersOK {

	return &ListMoveOrdersOK{}
}

// WithPayload adds the payload to the list move orders o k response
func (o *ListMoveOrdersOK) WithPayload(payload ghcmessages.MoveOrders) *ListMoveOrdersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list move orders o k response
func (o *ListMoveOrdersOK) SetPayload(payload ghcmessages.MoveOrders) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMoveOrdersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = ghcmessages.MoveOrders{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListMoveOrdersBadRequestCode is the HTTP code returned for type ListMoveOrdersBadRequest
const ListMoveOrdersBadRequestCode int = 400

/*ListMoveOrdersBadRequest The request payload is invalid

swagger:response listMoveOrdersBadRequest
*/
type ListMoveOrdersBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListMoveOrdersBadRequest creates ListMoveOrdersBadRequest with default headers values
func NewListMoveOrdersBadRequest() *ListMoveOrdersBadRequest {

	return &ListMoveOrdersBadRequest{}
}

// WithPayload adds the payload to the list move orders bad request response
func (o *ListMoveOrdersBadRequest) WithPayload(payload interface{}) *ListMoveOrdersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list move orders bad request response
func (o *ListMoveOrdersBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMoveOrdersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListMoveOrdersUnauthorizedCode is the HTTP code returned for type ListMoveOrdersUnauthorized
const ListMoveOrdersUnauthorizedCode int = 401

/*ListMoveOrdersUnauthorized The request was denied

swagger:response listMoveOrdersUnauthorized
*/
type ListMoveOrdersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListMoveOrdersUnauthorized creates ListMoveOrdersUnauthorized with default headers values
func NewListMoveOrdersUnauthorized() *ListMoveOrdersUnauthorized {

	return &ListMoveOrdersUnauthorized{}
}

// WithPayload adds the payload to the list move orders unauthorized response
func (o *ListMoveOrdersUnauthorized) WithPayload(payload interface{}) *ListMoveOrdersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list move orders unauthorized response
func (o *ListMoveOrdersUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMoveOrdersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListMoveOrdersForbiddenCode is the HTTP code returned for type ListMoveOrdersForbidden
const ListMoveOrdersForbiddenCode int = 403

/*ListMoveOrdersForbidden The request was denied

swagger:response listMoveOrdersForbidden
*/
type ListMoveOrdersForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListMoveOrdersForbidden creates ListMoveOrdersForbidden with default headers values
func NewListMoveOrdersForbidden() *ListMoveOrdersForbidden {

	return &ListMoveOrdersForbidden{}
}

// WithPayload adds the payload to the list move orders forbidden response
func (o *ListMoveOrdersForbidden) WithPayload(payload interface{}) *ListMoveOrdersForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list move orders forbidden response
func (o *ListMoveOrdersForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMoveOrdersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListMoveOrdersNotFoundCode is the HTTP code returned for type ListMoveOrdersNotFound
const ListMoveOrdersNotFoundCode int = 404

/*ListMoveOrdersNotFound The requested resource wasn't found

swagger:response listMoveOrdersNotFound
*/
type ListMoveOrdersNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListMoveOrdersNotFound creates ListMoveOrdersNotFound with default headers values
func NewListMoveOrdersNotFound() *ListMoveOrdersNotFound {

	return &ListMoveOrdersNotFound{}
}

// WithPayload adds the payload to the list move orders not found response
func (o *ListMoveOrdersNotFound) WithPayload(payload interface{}) *ListMoveOrdersNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list move orders not found response
func (o *ListMoveOrdersNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMoveOrdersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListMoveOrdersInternalServerErrorCode is the HTTP code returned for type ListMoveOrdersInternalServerError
const ListMoveOrdersInternalServerErrorCode int = 500

/*ListMoveOrdersInternalServerError A server error occurred

swagger:response listMoveOrdersInternalServerError
*/
type ListMoveOrdersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListMoveOrdersInternalServerError creates ListMoveOrdersInternalServerError with default headers values
func NewListMoveOrdersInternalServerError() *ListMoveOrdersInternalServerError {

	return &ListMoveOrdersInternalServerError{}
}

// WithPayload adds the payload to the list move orders internal server error response
func (o *ListMoveOrdersInternalServerError) WithPayload(payload interface{}) *ListMoveOrdersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list move orders internal server error response
func (o *ListMoveOrdersInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMoveOrdersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
