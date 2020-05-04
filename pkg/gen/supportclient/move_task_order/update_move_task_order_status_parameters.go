// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	supportmessages "github.com/transcom/mymove/pkg/gen/supportmessages"
)

// NewUpdateMoveTaskOrderStatusParams creates a new UpdateMoveTaskOrderStatusParams object
// with the default values initialized.
func NewUpdateMoveTaskOrderStatusParams() *UpdateMoveTaskOrderStatusParams {
	var ()
	return &UpdateMoveTaskOrderStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMoveTaskOrderStatusParamsWithTimeout creates a new UpdateMoveTaskOrderStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMoveTaskOrderStatusParamsWithTimeout(timeout time.Duration) *UpdateMoveTaskOrderStatusParams {
	var ()
	return &UpdateMoveTaskOrderStatusParams{

		timeout: timeout,
	}
}

// NewUpdateMoveTaskOrderStatusParamsWithContext creates a new UpdateMoveTaskOrderStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMoveTaskOrderStatusParamsWithContext(ctx context.Context) *UpdateMoveTaskOrderStatusParams {
	var ()
	return &UpdateMoveTaskOrderStatusParams{

		Context: ctx,
	}
}

// NewUpdateMoveTaskOrderStatusParamsWithHTTPClient creates a new UpdateMoveTaskOrderStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMoveTaskOrderStatusParamsWithHTTPClient(client *http.Client) *UpdateMoveTaskOrderStatusParams {
	var ()
	return &UpdateMoveTaskOrderStatusParams{
		HTTPClient: client,
	}
}

/*UpdateMoveTaskOrderStatusParams contains all the parameters to send to the API endpoint
for the update move task order status operation typically these are written to a http.Request
*/
type UpdateMoveTaskOrderStatusParams struct {

	/*IfMatch
	  Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error.


	*/
	IfMatch string
	/*Body*/
	Body *supportmessages.UpdateMoveTaskOrderStatus
	/*MoveTaskOrderID
	  UUID of move task order.

	*/
	MoveTaskOrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update move task order status params
func (o *UpdateMoveTaskOrderStatusParams) WithTimeout(timeout time.Duration) *UpdateMoveTaskOrderStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update move task order status params
func (o *UpdateMoveTaskOrderStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update move task order status params
func (o *UpdateMoveTaskOrderStatusParams) WithContext(ctx context.Context) *UpdateMoveTaskOrderStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update move task order status params
func (o *UpdateMoveTaskOrderStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update move task order status params
func (o *UpdateMoveTaskOrderStatusParams) WithHTTPClient(client *http.Client) *UpdateMoveTaskOrderStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update move task order status params
func (o *UpdateMoveTaskOrderStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the update move task order status params
func (o *UpdateMoveTaskOrderStatusParams) WithIfMatch(ifMatch string) *UpdateMoveTaskOrderStatusParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the update move task order status params
func (o *UpdateMoveTaskOrderStatusParams) SetIfMatch(ifMatch string) {
	o.IfMatch = ifMatch
}

// WithBody adds the body to the update move task order status params
func (o *UpdateMoveTaskOrderStatusParams) WithBody(body *supportmessages.UpdateMoveTaskOrderStatus) *UpdateMoveTaskOrderStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update move task order status params
func (o *UpdateMoveTaskOrderStatusParams) SetBody(body *supportmessages.UpdateMoveTaskOrderStatus) {
	o.Body = body
}

// WithMoveTaskOrderID adds the moveTaskOrderID to the update move task order status params
func (o *UpdateMoveTaskOrderStatusParams) WithMoveTaskOrderID(moveTaskOrderID string) *UpdateMoveTaskOrderStatusParams {
	o.SetMoveTaskOrderID(moveTaskOrderID)
	return o
}

// SetMoveTaskOrderID adds the moveTaskOrderId to the update move task order status params
func (o *UpdateMoveTaskOrderStatusParams) SetMoveTaskOrderID(moveTaskOrderID string) {
	o.MoveTaskOrderID = moveTaskOrderID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMoveTaskOrderStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param If-Match
	if err := r.SetHeaderParam("If-Match", o.IfMatch); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param moveTaskOrderID
	if err := r.SetPathParam("moveTaskOrderID", o.MoveTaskOrderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
