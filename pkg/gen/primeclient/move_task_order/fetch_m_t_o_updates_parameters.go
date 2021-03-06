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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFetchMTOUpdatesParams creates a new FetchMTOUpdatesParams object
// with the default values initialized.
func NewFetchMTOUpdatesParams() *FetchMTOUpdatesParams {
	var ()
	return &FetchMTOUpdatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFetchMTOUpdatesParamsWithTimeout creates a new FetchMTOUpdatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFetchMTOUpdatesParamsWithTimeout(timeout time.Duration) *FetchMTOUpdatesParams {
	var ()
	return &FetchMTOUpdatesParams{

		timeout: timeout,
	}
}

// NewFetchMTOUpdatesParamsWithContext creates a new FetchMTOUpdatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewFetchMTOUpdatesParamsWithContext(ctx context.Context) *FetchMTOUpdatesParams {
	var ()
	return &FetchMTOUpdatesParams{

		Context: ctx,
	}
}

// NewFetchMTOUpdatesParamsWithHTTPClient creates a new FetchMTOUpdatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFetchMTOUpdatesParamsWithHTTPClient(client *http.Client) *FetchMTOUpdatesParams {
	var ()
	return &FetchMTOUpdatesParams{
		HTTPClient: client,
	}
}

/*FetchMTOUpdatesParams contains all the parameters to send to the API endpoint
for the fetch m t o updates operation typically these are written to a http.Request
*/
type FetchMTOUpdatesParams struct {

	/*Since
	  Only return move task orders updated since this time.

	*/
	Since *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the fetch m t o updates params
func (o *FetchMTOUpdatesParams) WithTimeout(timeout time.Duration) *FetchMTOUpdatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fetch m t o updates params
func (o *FetchMTOUpdatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fetch m t o updates params
func (o *FetchMTOUpdatesParams) WithContext(ctx context.Context) *FetchMTOUpdatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fetch m t o updates params
func (o *FetchMTOUpdatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fetch m t o updates params
func (o *FetchMTOUpdatesParams) WithHTTPClient(client *http.Client) *FetchMTOUpdatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fetch m t o updates params
func (o *FetchMTOUpdatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSince adds the since to the fetch m t o updates params
func (o *FetchMTOUpdatesParams) WithSince(since *int64) *FetchMTOUpdatesParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the fetch m t o updates params
func (o *FetchMTOUpdatesParams) SetSince(since *int64) {
	o.Since = since
}

// WriteToRequest writes these params to a swagger request
func (o *FetchMTOUpdatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Since != nil {

		// query param since
		var qrSince int64
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := swag.FormatInt64(qrSince)
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
