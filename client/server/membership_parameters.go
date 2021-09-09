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

// NewMembershipParams creates a new MembershipParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMembershipParams() *MembershipParams {
	return &MembershipParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMembershipParamsWithTimeout creates a new MembershipParams object
// with the ability to set a timeout on a request.
func NewMembershipParamsWithTimeout(timeout time.Duration) *MembershipParams {
	return &MembershipParams{
		timeout: timeout,
	}
}

// NewMembershipParamsWithContext creates a new MembershipParams object
// with the ability to set a context for a request.
func NewMembershipParamsWithContext(ctx context.Context) *MembershipParams {
	return &MembershipParams{
		Context: ctx,
	}
}

// NewMembershipParamsWithHTTPClient creates a new MembershipParams object
// with the ability to set a custom HTTPClient for a request.
func NewMembershipParamsWithHTTPClient(client *http.Client) *MembershipParams {
	return &MembershipParams{
		HTTPClient: client,
	}
}

/* MembershipParams contains all the parameters to send to the API endpoint
   for the membership operation.

   Typically these are written to a http.Request.
*/
type MembershipParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the membership params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MembershipParams) WithDefaults() *MembershipParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the membership params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MembershipParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the membership params
func (o *MembershipParams) WithTimeout(timeout time.Duration) *MembershipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the membership params
func (o *MembershipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the membership params
func (o *MembershipParams) WithContext(ctx context.Context) *MembershipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the membership params
func (o *MembershipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the membership params
func (o *MembershipParams) WithHTTPClient(client *http.Client) *MembershipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the membership params
func (o *MembershipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MembershipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
