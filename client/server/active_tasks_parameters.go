// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewActiveTasksParams creates a new ActiveTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActiveTasksParams() *ActiveTasksParams {
	return &ActiveTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActiveTasksParamsWithTimeout creates a new ActiveTasksParams object
// with the ability to set a timeout on a request.
func NewActiveTasksParamsWithTimeout(timeout time.Duration) *ActiveTasksParams {
	return &ActiveTasksParams{
		timeout: timeout,
	}
}

// NewActiveTasksParamsWithContext creates a new ActiveTasksParams object
// with the ability to set a context for a request.
func NewActiveTasksParamsWithContext(ctx context.Context) *ActiveTasksParams {
	return &ActiveTasksParams{
		Context: ctx,
	}
}

// NewActiveTasksParamsWithHTTPClient creates a new ActiveTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewActiveTasksParamsWithHTTPClient(client *http.Client) *ActiveTasksParams {
	return &ActiveTasksParams{
		HTTPClient: client,
	}
}

/* ActiveTasksParams contains all the parameters to send to the API endpoint
   for the active tasks operation.

   Typically these are written to a http.Request.
*/
type ActiveTasksParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the active tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActiveTasksParams) WithDefaults() *ActiveTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the active tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActiveTasksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the active tasks params
func (o *ActiveTasksParams) WithTimeout(timeout time.Duration) *ActiveTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the active tasks params
func (o *ActiveTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the active tasks params
func (o *ActiveTasksParams) WithContext(ctx context.Context) *ActiveTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the active tasks params
func (o *ActiveTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the active tasks params
func (o *ActiveTasksParams) WithHTTPClient(client *http.Client) *ActiveTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the active tasks params
func (o *ActiveTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ActiveTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
