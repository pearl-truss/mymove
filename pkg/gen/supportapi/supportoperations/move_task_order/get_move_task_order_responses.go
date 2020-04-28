// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	supportmessages "github.com/transcom/mymove/pkg/gen/supportmessages"
)

// GetMoveTaskOrderOKCode is the HTTP code returned for type GetMoveTaskOrderOK
const GetMoveTaskOrderOKCode int = 200

/*GetMoveTaskOrderOK Successfully retrieve an individual move task order

swagger:response getMoveTaskOrderOK
*/
type GetMoveTaskOrderOK struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.MoveTaskOrder `json:"body,omitempty"`
}

// NewGetMoveTaskOrderOK creates GetMoveTaskOrderOK with default headers values
func NewGetMoveTaskOrderOK() *GetMoveTaskOrderOK {

	return &GetMoveTaskOrderOK{}
}

// WithPayload adds the payload to the get move task order o k response
func (o *GetMoveTaskOrderOK) WithPayload(payload *supportmessages.MoveTaskOrder) *GetMoveTaskOrderOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move task order o k response
func (o *GetMoveTaskOrderOK) SetPayload(payload *supportmessages.MoveTaskOrder) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveTaskOrderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMoveTaskOrderBadRequestCode is the HTTP code returned for type GetMoveTaskOrderBadRequest
const GetMoveTaskOrderBadRequestCode int = 400

/*GetMoveTaskOrderBadRequest The request payload is invalid

swagger:response getMoveTaskOrderBadRequest
*/
type GetMoveTaskOrderBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetMoveTaskOrderBadRequest creates GetMoveTaskOrderBadRequest with default headers values
func NewGetMoveTaskOrderBadRequest() *GetMoveTaskOrderBadRequest {

	return &GetMoveTaskOrderBadRequest{}
}

// WithPayload adds the payload to the get move task order bad request response
func (o *GetMoveTaskOrderBadRequest) WithPayload(payload interface{}) *GetMoveTaskOrderBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move task order bad request response
func (o *GetMoveTaskOrderBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveTaskOrderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMoveTaskOrderUnauthorizedCode is the HTTP code returned for type GetMoveTaskOrderUnauthorized
const GetMoveTaskOrderUnauthorizedCode int = 401

/*GetMoveTaskOrderUnauthorized The request was denied

swagger:response getMoveTaskOrderUnauthorized
*/
type GetMoveTaskOrderUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetMoveTaskOrderUnauthorized creates GetMoveTaskOrderUnauthorized with default headers values
func NewGetMoveTaskOrderUnauthorized() *GetMoveTaskOrderUnauthorized {

	return &GetMoveTaskOrderUnauthorized{}
}

// WithPayload adds the payload to the get move task order unauthorized response
func (o *GetMoveTaskOrderUnauthorized) WithPayload(payload interface{}) *GetMoveTaskOrderUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move task order unauthorized response
func (o *GetMoveTaskOrderUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveTaskOrderUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMoveTaskOrderForbiddenCode is the HTTP code returned for type GetMoveTaskOrderForbidden
const GetMoveTaskOrderForbiddenCode int = 403

/*GetMoveTaskOrderForbidden The request was denied

swagger:response getMoveTaskOrderForbidden
*/
type GetMoveTaskOrderForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetMoveTaskOrderForbidden creates GetMoveTaskOrderForbidden with default headers values
func NewGetMoveTaskOrderForbidden() *GetMoveTaskOrderForbidden {

	return &GetMoveTaskOrderForbidden{}
}

// WithPayload adds the payload to the get move task order forbidden response
func (o *GetMoveTaskOrderForbidden) WithPayload(payload interface{}) *GetMoveTaskOrderForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move task order forbidden response
func (o *GetMoveTaskOrderForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveTaskOrderForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMoveTaskOrderNotFoundCode is the HTTP code returned for type GetMoveTaskOrderNotFound
const GetMoveTaskOrderNotFoundCode int = 404

/*GetMoveTaskOrderNotFound The requested resource wasn't found

swagger:response getMoveTaskOrderNotFound
*/
type GetMoveTaskOrderNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetMoveTaskOrderNotFound creates GetMoveTaskOrderNotFound with default headers values
func NewGetMoveTaskOrderNotFound() *GetMoveTaskOrderNotFound {

	return &GetMoveTaskOrderNotFound{}
}

// WithPayload adds the payload to the get move task order not found response
func (o *GetMoveTaskOrderNotFound) WithPayload(payload interface{}) *GetMoveTaskOrderNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move task order not found response
func (o *GetMoveTaskOrderNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveTaskOrderNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMoveTaskOrderUnprocessableEntityCode is the HTTP code returned for type GetMoveTaskOrderUnprocessableEntity
const GetMoveTaskOrderUnprocessableEntityCode int = 422

/*GetMoveTaskOrderUnprocessableEntity The payload was unprocessable.

swagger:response getMoveTaskOrderUnprocessableEntity
*/
type GetMoveTaskOrderUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.Error `json:"body,omitempty"`
}

// NewGetMoveTaskOrderUnprocessableEntity creates GetMoveTaskOrderUnprocessableEntity with default headers values
func NewGetMoveTaskOrderUnprocessableEntity() *GetMoveTaskOrderUnprocessableEntity {

	return &GetMoveTaskOrderUnprocessableEntity{}
}

// WithPayload adds the payload to the get move task order unprocessable entity response
func (o *GetMoveTaskOrderUnprocessableEntity) WithPayload(payload *supportmessages.Error) *GetMoveTaskOrderUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move task order unprocessable entity response
func (o *GetMoveTaskOrderUnprocessableEntity) SetPayload(payload *supportmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveTaskOrderUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMoveTaskOrderInternalServerErrorCode is the HTTP code returned for type GetMoveTaskOrderInternalServerError
const GetMoveTaskOrderInternalServerErrorCode int = 500

/*GetMoveTaskOrderInternalServerError A server error occurred

swagger:response getMoveTaskOrderInternalServerError
*/
type GetMoveTaskOrderInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetMoveTaskOrderInternalServerError creates GetMoveTaskOrderInternalServerError with default headers values
func NewGetMoveTaskOrderInternalServerError() *GetMoveTaskOrderInternalServerError {

	return &GetMoveTaskOrderInternalServerError{}
}

// WithPayload adds the payload to the get move task order internal server error response
func (o *GetMoveTaskOrderInternalServerError) WithPayload(payload interface{}) *GetMoveTaskOrderInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move task order internal server error response
func (o *GetMoveTaskOrderInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveTaskOrderInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
