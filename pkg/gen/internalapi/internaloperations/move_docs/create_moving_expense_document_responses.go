// Code generated by go-swagger; DO NOT EDIT.

package move_docs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	internalmessages "github.com/transcom/mymove/pkg/gen/internalmessages"
)

// CreateMovingExpenseDocumentOKCode is the HTTP code returned for type CreateMovingExpenseDocumentOK
const CreateMovingExpenseDocumentOKCode int = 200

/*CreateMovingExpenseDocumentOK returns new moving expense document object

swagger:response createMovingExpenseDocumentOK
*/
type CreateMovingExpenseDocumentOK struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.MoveDocumentPayload `json:"body,omitempty"`
}

// NewCreateMovingExpenseDocumentOK creates CreateMovingExpenseDocumentOK with default headers values
func NewCreateMovingExpenseDocumentOK() *CreateMovingExpenseDocumentOK {

	return &CreateMovingExpenseDocumentOK{}
}

// WithPayload adds the payload to the create moving expense document o k response
func (o *CreateMovingExpenseDocumentOK) WithPayload(payload *internalmessages.MoveDocumentPayload) *CreateMovingExpenseDocumentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create moving expense document o k response
func (o *CreateMovingExpenseDocumentOK) SetPayload(payload *internalmessages.MoveDocumentPayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMovingExpenseDocumentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMovingExpenseDocumentBadRequestCode is the HTTP code returned for type CreateMovingExpenseDocumentBadRequest
const CreateMovingExpenseDocumentBadRequestCode int = 400

/*CreateMovingExpenseDocumentBadRequest invalid request

swagger:response createMovingExpenseDocumentBadRequest
*/
type CreateMovingExpenseDocumentBadRequest struct {
}

// NewCreateMovingExpenseDocumentBadRequest creates CreateMovingExpenseDocumentBadRequest with default headers values
func NewCreateMovingExpenseDocumentBadRequest() *CreateMovingExpenseDocumentBadRequest {

	return &CreateMovingExpenseDocumentBadRequest{}
}

// WriteResponse to the client
func (o *CreateMovingExpenseDocumentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateMovingExpenseDocumentUnauthorizedCode is the HTTP code returned for type CreateMovingExpenseDocumentUnauthorized
const CreateMovingExpenseDocumentUnauthorizedCode int = 401

/*CreateMovingExpenseDocumentUnauthorized must be authenticated to use this endpoint

swagger:response createMovingExpenseDocumentUnauthorized
*/
type CreateMovingExpenseDocumentUnauthorized struct {
}

// NewCreateMovingExpenseDocumentUnauthorized creates CreateMovingExpenseDocumentUnauthorized with default headers values
func NewCreateMovingExpenseDocumentUnauthorized() *CreateMovingExpenseDocumentUnauthorized {

	return &CreateMovingExpenseDocumentUnauthorized{}
}

// WriteResponse to the client
func (o *CreateMovingExpenseDocumentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// CreateMovingExpenseDocumentForbiddenCode is the HTTP code returned for type CreateMovingExpenseDocumentForbidden
const CreateMovingExpenseDocumentForbiddenCode int = 403

/*CreateMovingExpenseDocumentForbidden not authorized to modify this move

swagger:response createMovingExpenseDocumentForbidden
*/
type CreateMovingExpenseDocumentForbidden struct {
}

// NewCreateMovingExpenseDocumentForbidden creates CreateMovingExpenseDocumentForbidden with default headers values
func NewCreateMovingExpenseDocumentForbidden() *CreateMovingExpenseDocumentForbidden {

	return &CreateMovingExpenseDocumentForbidden{}
}

// WriteResponse to the client
func (o *CreateMovingExpenseDocumentForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// CreateMovingExpenseDocumentInternalServerErrorCode is the HTTP code returned for type CreateMovingExpenseDocumentInternalServerError
const CreateMovingExpenseDocumentInternalServerErrorCode int = 500

/*CreateMovingExpenseDocumentInternalServerError server error

swagger:response createMovingExpenseDocumentInternalServerError
*/
type CreateMovingExpenseDocumentInternalServerError struct {
}

// NewCreateMovingExpenseDocumentInternalServerError creates CreateMovingExpenseDocumentInternalServerError with default headers values
func NewCreateMovingExpenseDocumentInternalServerError() *CreateMovingExpenseDocumentInternalServerError {

	return &CreateMovingExpenseDocumentInternalServerError{}
}

// WriteResponse to the client
func (o *CreateMovingExpenseDocumentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}