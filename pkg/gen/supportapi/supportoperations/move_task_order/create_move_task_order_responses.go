// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	supportmessages "github.com/transcom/mymove/pkg/gen/supportmessages"
)

// CreateMoveTaskOrderCreatedCode is the HTTP code returned for type CreateMoveTaskOrderCreated
const CreateMoveTaskOrderCreatedCode int = 201

/*CreateMoveTaskOrderCreated Successfully created MoveTaskOrder object.

swagger:response createMoveTaskOrderCreated
*/
type CreateMoveTaskOrderCreated struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.MoveTaskOrder `json:"body,omitempty"`
}

// NewCreateMoveTaskOrderCreated creates CreateMoveTaskOrderCreated with default headers values
func NewCreateMoveTaskOrderCreated() *CreateMoveTaskOrderCreated {

	return &CreateMoveTaskOrderCreated{}
}

// WithPayload adds the payload to the create move task order created response
func (o *CreateMoveTaskOrderCreated) WithPayload(payload *supportmessages.MoveTaskOrder) *CreateMoveTaskOrderCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create move task order created response
func (o *CreateMoveTaskOrderCreated) SetPayload(payload *supportmessages.MoveTaskOrder) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMoveTaskOrderCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMoveTaskOrderBadRequestCode is the HTTP code returned for type CreateMoveTaskOrderBadRequest
const CreateMoveTaskOrderBadRequestCode int = 400

/*CreateMoveTaskOrderBadRequest The request payload is invalid.

swagger:response createMoveTaskOrderBadRequest
*/
type CreateMoveTaskOrderBadRequest struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewCreateMoveTaskOrderBadRequest creates CreateMoveTaskOrderBadRequest with default headers values
func NewCreateMoveTaskOrderBadRequest() *CreateMoveTaskOrderBadRequest {

	return &CreateMoveTaskOrderBadRequest{}
}

// WithPayload adds the payload to the create move task order bad request response
func (o *CreateMoveTaskOrderBadRequest) WithPayload(payload *supportmessages.ClientError) *CreateMoveTaskOrderBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create move task order bad request response
func (o *CreateMoveTaskOrderBadRequest) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMoveTaskOrderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMoveTaskOrderUnauthorizedCode is the HTTP code returned for type CreateMoveTaskOrderUnauthorized
const CreateMoveTaskOrderUnauthorizedCode int = 401

/*CreateMoveTaskOrderUnauthorized The request was unauthorized.

swagger:response createMoveTaskOrderUnauthorized
*/
type CreateMoveTaskOrderUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewCreateMoveTaskOrderUnauthorized creates CreateMoveTaskOrderUnauthorized with default headers values
func NewCreateMoveTaskOrderUnauthorized() *CreateMoveTaskOrderUnauthorized {

	return &CreateMoveTaskOrderUnauthorized{}
}

// WithPayload adds the payload to the create move task order unauthorized response
func (o *CreateMoveTaskOrderUnauthorized) WithPayload(payload interface{}) *CreateMoveTaskOrderUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create move task order unauthorized response
func (o *CreateMoveTaskOrderUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMoveTaskOrderUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// CreateMoveTaskOrderForbiddenCode is the HTTP code returned for type CreateMoveTaskOrderForbidden
const CreateMoveTaskOrderForbiddenCode int = 403

/*CreateMoveTaskOrderForbidden The client doesn't have permissions to perform the request.

swagger:response createMoveTaskOrderForbidden
*/
type CreateMoveTaskOrderForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewCreateMoveTaskOrderForbidden creates CreateMoveTaskOrderForbidden with default headers values
func NewCreateMoveTaskOrderForbidden() *CreateMoveTaskOrderForbidden {

	return &CreateMoveTaskOrderForbidden{}
}

// WithPayload adds the payload to the create move task order forbidden response
func (o *CreateMoveTaskOrderForbidden) WithPayload(payload interface{}) *CreateMoveTaskOrderForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create move task order forbidden response
func (o *CreateMoveTaskOrderForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMoveTaskOrderForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// CreateMoveTaskOrderNotFoundCode is the HTTP code returned for type CreateMoveTaskOrderNotFound
const CreateMoveTaskOrderNotFoundCode int = 404

/*CreateMoveTaskOrderNotFound The requested resource wasn't found.

swagger:response createMoveTaskOrderNotFound
*/
type CreateMoveTaskOrderNotFound struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewCreateMoveTaskOrderNotFound creates CreateMoveTaskOrderNotFound with default headers values
func NewCreateMoveTaskOrderNotFound() *CreateMoveTaskOrderNotFound {

	return &CreateMoveTaskOrderNotFound{}
}

// WithPayload adds the payload to the create move task order not found response
func (o *CreateMoveTaskOrderNotFound) WithPayload(payload *supportmessages.ClientError) *CreateMoveTaskOrderNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create move task order not found response
func (o *CreateMoveTaskOrderNotFound) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMoveTaskOrderNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMoveTaskOrderUnprocessableEntityCode is the HTTP code returned for type CreateMoveTaskOrderUnprocessableEntity
const CreateMoveTaskOrderUnprocessableEntityCode int = 422

/*CreateMoveTaskOrderUnprocessableEntity The payload was unprocessable.

swagger:response createMoveTaskOrderUnprocessableEntity
*/
type CreateMoveTaskOrderUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ValidationError `json:"body,omitempty"`
}

// NewCreateMoveTaskOrderUnprocessableEntity creates CreateMoveTaskOrderUnprocessableEntity with default headers values
func NewCreateMoveTaskOrderUnprocessableEntity() *CreateMoveTaskOrderUnprocessableEntity {

	return &CreateMoveTaskOrderUnprocessableEntity{}
}

// WithPayload adds the payload to the create move task order unprocessable entity response
func (o *CreateMoveTaskOrderUnprocessableEntity) WithPayload(payload *supportmessages.ValidationError) *CreateMoveTaskOrderUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create move task order unprocessable entity response
func (o *CreateMoveTaskOrderUnprocessableEntity) SetPayload(payload *supportmessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMoveTaskOrderUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMoveTaskOrderInternalServerErrorCode is the HTTP code returned for type CreateMoveTaskOrderInternalServerError
const CreateMoveTaskOrderInternalServerErrorCode int = 500

/*CreateMoveTaskOrderInternalServerError A server error occurred.

swagger:response createMoveTaskOrderInternalServerError
*/
type CreateMoveTaskOrderInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.Error `json:"body,omitempty"`
}

// NewCreateMoveTaskOrderInternalServerError creates CreateMoveTaskOrderInternalServerError with default headers values
func NewCreateMoveTaskOrderInternalServerError() *CreateMoveTaskOrderInternalServerError {

	return &CreateMoveTaskOrderInternalServerError{}
}

// WithPayload adds the payload to the create move task order internal server error response
func (o *CreateMoveTaskOrderInternalServerError) WithPayload(payload *supportmessages.Error) *CreateMoveTaskOrderInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create move task order internal server error response
func (o *CreateMoveTaskOrderInternalServerError) SetPayload(payload *supportmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMoveTaskOrderInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
