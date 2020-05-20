// Code generated by go-swagger; DO NOT EDIT.

package payment_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	supportmessages "github.com/transcom/mymove/pkg/gen/supportmessages"
)

// UpdatePaymentRequestStatusOKCode is the HTTP code returned for type UpdatePaymentRequestStatusOK
const UpdatePaymentRequestStatusOKCode int = 200

/*UpdatePaymentRequestStatusOK Successfully updated payment request status.

swagger:response updatePaymentRequestStatusOK
*/
type UpdatePaymentRequestStatusOK struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.PaymentRequest `json:"body,omitempty"`
}

// NewUpdatePaymentRequestStatusOK creates UpdatePaymentRequestStatusOK with default headers values
func NewUpdatePaymentRequestStatusOK() *UpdatePaymentRequestStatusOK {

	return &UpdatePaymentRequestStatusOK{}
}

// WithPayload adds the payload to the update payment request status o k response
func (o *UpdatePaymentRequestStatusOK) WithPayload(payload *supportmessages.PaymentRequest) *UpdatePaymentRequestStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request status o k response
func (o *UpdatePaymentRequestStatusOK) SetPayload(payload *supportmessages.PaymentRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePaymentRequestStatusBadRequestCode is the HTTP code returned for type UpdatePaymentRequestStatusBadRequest
const UpdatePaymentRequestStatusBadRequestCode int = 400

/*UpdatePaymentRequestStatusBadRequest The request payload is invalid.

swagger:response updatePaymentRequestStatusBadRequest
*/
type UpdatePaymentRequestStatusBadRequest struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewUpdatePaymentRequestStatusBadRequest creates UpdatePaymentRequestStatusBadRequest with default headers values
func NewUpdatePaymentRequestStatusBadRequest() *UpdatePaymentRequestStatusBadRequest {

	return &UpdatePaymentRequestStatusBadRequest{}
}

// WithPayload adds the payload to the update payment request status bad request response
func (o *UpdatePaymentRequestStatusBadRequest) WithPayload(payload *supportmessages.ClientError) *UpdatePaymentRequestStatusBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request status bad request response
func (o *UpdatePaymentRequestStatusBadRequest) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestStatusBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePaymentRequestStatusUnauthorizedCode is the HTTP code returned for type UpdatePaymentRequestStatusUnauthorized
const UpdatePaymentRequestStatusUnauthorizedCode int = 401

/*UpdatePaymentRequestStatusUnauthorized The request was unauthorized.

swagger:response updatePaymentRequestStatusUnauthorized
*/
type UpdatePaymentRequestStatusUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdatePaymentRequestStatusUnauthorized creates UpdatePaymentRequestStatusUnauthorized with default headers values
func NewUpdatePaymentRequestStatusUnauthorized() *UpdatePaymentRequestStatusUnauthorized {

	return &UpdatePaymentRequestStatusUnauthorized{}
}

// WithPayload adds the payload to the update payment request status unauthorized response
func (o *UpdatePaymentRequestStatusUnauthorized) WithPayload(payload interface{}) *UpdatePaymentRequestStatusUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request status unauthorized response
func (o *UpdatePaymentRequestStatusUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdatePaymentRequestStatusForbiddenCode is the HTTP code returned for type UpdatePaymentRequestStatusForbidden
const UpdatePaymentRequestStatusForbiddenCode int = 403

/*UpdatePaymentRequestStatusForbidden The client doesn't have permissions to perform the request.

swagger:response updatePaymentRequestStatusForbidden
*/
type UpdatePaymentRequestStatusForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdatePaymentRequestStatusForbidden creates UpdatePaymentRequestStatusForbidden with default headers values
func NewUpdatePaymentRequestStatusForbidden() *UpdatePaymentRequestStatusForbidden {

	return &UpdatePaymentRequestStatusForbidden{}
}

// WithPayload adds the payload to the update payment request status forbidden response
func (o *UpdatePaymentRequestStatusForbidden) WithPayload(payload interface{}) *UpdatePaymentRequestStatusForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request status forbidden response
func (o *UpdatePaymentRequestStatusForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestStatusForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdatePaymentRequestStatusNotFoundCode is the HTTP code returned for type UpdatePaymentRequestStatusNotFound
const UpdatePaymentRequestStatusNotFoundCode int = 404

/*UpdatePaymentRequestStatusNotFound The requested resource wasn't found.

swagger:response updatePaymentRequestStatusNotFound
*/
type UpdatePaymentRequestStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewUpdatePaymentRequestStatusNotFound creates UpdatePaymentRequestStatusNotFound with default headers values
func NewUpdatePaymentRequestStatusNotFound() *UpdatePaymentRequestStatusNotFound {

	return &UpdatePaymentRequestStatusNotFound{}
}

// WithPayload adds the payload to the update payment request status not found response
func (o *UpdatePaymentRequestStatusNotFound) WithPayload(payload *supportmessages.ClientError) *UpdatePaymentRequestStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request status not found response
func (o *UpdatePaymentRequestStatusNotFound) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePaymentRequestStatusPreconditionFailedCode is the HTTP code returned for type UpdatePaymentRequestStatusPreconditionFailed
const UpdatePaymentRequestStatusPreconditionFailedCode int = 412

/*UpdatePaymentRequestStatusPreconditionFailed Precondition failed, likely due to a stale eTag (If-Match). Fetch the request again to get the updated eTag value.

swagger:response updatePaymentRequestStatusPreconditionFailed
*/
type UpdatePaymentRequestStatusPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewUpdatePaymentRequestStatusPreconditionFailed creates UpdatePaymentRequestStatusPreconditionFailed with default headers values
func NewUpdatePaymentRequestStatusPreconditionFailed() *UpdatePaymentRequestStatusPreconditionFailed {

	return &UpdatePaymentRequestStatusPreconditionFailed{}
}

// WithPayload adds the payload to the update payment request status precondition failed response
func (o *UpdatePaymentRequestStatusPreconditionFailed) WithPayload(payload *supportmessages.ClientError) *UpdatePaymentRequestStatusPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request status precondition failed response
func (o *UpdatePaymentRequestStatusPreconditionFailed) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestStatusPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePaymentRequestStatusUnprocessableEntityCode is the HTTP code returned for type UpdatePaymentRequestStatusUnprocessableEntity
const UpdatePaymentRequestStatusUnprocessableEntityCode int = 422

/*UpdatePaymentRequestStatusUnprocessableEntity The payload was unprocessable.

swagger:response updatePaymentRequestStatusUnprocessableEntity
*/
type UpdatePaymentRequestStatusUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ValidationError `json:"body,omitempty"`
}

// NewUpdatePaymentRequestStatusUnprocessableEntity creates UpdatePaymentRequestStatusUnprocessableEntity with default headers values
func NewUpdatePaymentRequestStatusUnprocessableEntity() *UpdatePaymentRequestStatusUnprocessableEntity {

	return &UpdatePaymentRequestStatusUnprocessableEntity{}
}

// WithPayload adds the payload to the update payment request status unprocessable entity response
func (o *UpdatePaymentRequestStatusUnprocessableEntity) WithPayload(payload *supportmessages.ValidationError) *UpdatePaymentRequestStatusUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request status unprocessable entity response
func (o *UpdatePaymentRequestStatusUnprocessableEntity) SetPayload(payload *supportmessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestStatusUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePaymentRequestStatusInternalServerErrorCode is the HTTP code returned for type UpdatePaymentRequestStatusInternalServerError
const UpdatePaymentRequestStatusInternalServerErrorCode int = 500

/*UpdatePaymentRequestStatusInternalServerError A server error occurred.

swagger:response updatePaymentRequestStatusInternalServerError
*/
type UpdatePaymentRequestStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.Error `json:"body,omitempty"`
}

// NewUpdatePaymentRequestStatusInternalServerError creates UpdatePaymentRequestStatusInternalServerError with default headers values
func NewUpdatePaymentRequestStatusInternalServerError() *UpdatePaymentRequestStatusInternalServerError {

	return &UpdatePaymentRequestStatusInternalServerError{}
}

// WithPayload adds the payload to the update payment request status internal server error response
func (o *UpdatePaymentRequestStatusInternalServerError) WithPayload(payload *supportmessages.Error) *UpdatePaymentRequestStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment request status internal server error response
func (o *UpdatePaymentRequestStatusInternalServerError) SetPayload(payload *supportmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentRequestStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
