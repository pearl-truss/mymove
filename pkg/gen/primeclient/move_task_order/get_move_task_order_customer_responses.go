// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// GetMoveTaskOrderCustomerReader is a Reader for the GetMoveTaskOrderCustomer structure.
type GetMoveTaskOrderCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMoveTaskOrderCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMoveTaskOrderCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetMoveTaskOrderCustomerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMoveTaskOrderCustomerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMoveTaskOrderCustomerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetMoveTaskOrderCustomerUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMoveTaskOrderCustomerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMoveTaskOrderCustomerOK creates a GetMoveTaskOrderCustomerOK with default headers values
func NewGetMoveTaskOrderCustomerOK() *GetMoveTaskOrderCustomerOK {
	return &GetMoveTaskOrderCustomerOK{}
}

/*GetMoveTaskOrderCustomerOK handles this case with default header values.

Successfully retrieved customer associated with move task order
*/
type GetMoveTaskOrderCustomerOK struct {
	Payload *primemessages.Customer
}

func (o *GetMoveTaskOrderCustomerOK) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/{moveTaskOrderID}/customer][%d] getMoveTaskOrderCustomerOK  %+v", 200, o.Payload)
}

func (o *GetMoveTaskOrderCustomerOK) GetPayload() *primemessages.Customer {
	return o.Payload
}

func (o *GetMoveTaskOrderCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMoveTaskOrderCustomerUnauthorized creates a GetMoveTaskOrderCustomerUnauthorized with default headers values
func NewGetMoveTaskOrderCustomerUnauthorized() *GetMoveTaskOrderCustomerUnauthorized {
	return &GetMoveTaskOrderCustomerUnauthorized{}
}

/*GetMoveTaskOrderCustomerUnauthorized handles this case with default header values.

The request was denied
*/
type GetMoveTaskOrderCustomerUnauthorized struct {
	Payload interface{}
}

func (o *GetMoveTaskOrderCustomerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/{moveTaskOrderID}/customer][%d] getMoveTaskOrderCustomerUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMoveTaskOrderCustomerUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *GetMoveTaskOrderCustomerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMoveTaskOrderCustomerForbidden creates a GetMoveTaskOrderCustomerForbidden with default headers values
func NewGetMoveTaskOrderCustomerForbidden() *GetMoveTaskOrderCustomerForbidden {
	return &GetMoveTaskOrderCustomerForbidden{}
}

/*GetMoveTaskOrderCustomerForbidden handles this case with default header values.

The request was denied
*/
type GetMoveTaskOrderCustomerForbidden struct {
	Payload interface{}
}

func (o *GetMoveTaskOrderCustomerForbidden) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/{moveTaskOrderID}/customer][%d] getMoveTaskOrderCustomerForbidden  %+v", 403, o.Payload)
}

func (o *GetMoveTaskOrderCustomerForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetMoveTaskOrderCustomerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMoveTaskOrderCustomerNotFound creates a GetMoveTaskOrderCustomerNotFound with default headers values
func NewGetMoveTaskOrderCustomerNotFound() *GetMoveTaskOrderCustomerNotFound {
	return &GetMoveTaskOrderCustomerNotFound{}
}

/*GetMoveTaskOrderCustomerNotFound handles this case with default header values.

The requested resource wasn't found
*/
type GetMoveTaskOrderCustomerNotFound struct {
	Payload interface{}
}

func (o *GetMoveTaskOrderCustomerNotFound) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/{moveTaskOrderID}/customer][%d] getMoveTaskOrderCustomerNotFound  %+v", 404, o.Payload)
}

func (o *GetMoveTaskOrderCustomerNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetMoveTaskOrderCustomerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMoveTaskOrderCustomerUnprocessableEntity creates a GetMoveTaskOrderCustomerUnprocessableEntity with default headers values
func NewGetMoveTaskOrderCustomerUnprocessableEntity() *GetMoveTaskOrderCustomerUnprocessableEntity {
	return &GetMoveTaskOrderCustomerUnprocessableEntity{}
}

/*GetMoveTaskOrderCustomerUnprocessableEntity handles this case with default header values.

The payload was unprocessable.
*/
type GetMoveTaskOrderCustomerUnprocessableEntity struct {
	Payload *primemessages.ValidationError
}

func (o *GetMoveTaskOrderCustomerUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/{moveTaskOrderID}/customer][%d] getMoveTaskOrderCustomerUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetMoveTaskOrderCustomerUnprocessableEntity) GetPayload() *primemessages.ValidationError {
	return o.Payload
}

func (o *GetMoveTaskOrderCustomerUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMoveTaskOrderCustomerInternalServerError creates a GetMoveTaskOrderCustomerInternalServerError with default headers values
func NewGetMoveTaskOrderCustomerInternalServerError() *GetMoveTaskOrderCustomerInternalServerError {
	return &GetMoveTaskOrderCustomerInternalServerError{}
}

/*GetMoveTaskOrderCustomerInternalServerError handles this case with default header values.

A server error occurred
*/
type GetMoveTaskOrderCustomerInternalServerError struct {
	Payload interface{}
}

func (o *GetMoveTaskOrderCustomerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/{moveTaskOrderID}/customer][%d] getMoveTaskOrderCustomerInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMoveTaskOrderCustomerInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *GetMoveTaskOrderCustomerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
