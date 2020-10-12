// Code generated by go-swagger; DO NOT EDIT.

package design_documents

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

// NewDesignHeadDocParams creates a new DesignHeadDocParams object
// with the default values initialized.
func NewDesignHeadDocParams() *DesignHeadDocParams {
	var ()
	return &DesignHeadDocParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDesignHeadDocParamsWithTimeout creates a new DesignHeadDocParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDesignHeadDocParamsWithTimeout(timeout time.Duration) *DesignHeadDocParams {
	var ()
	return &DesignHeadDocParams{

		timeout: timeout,
	}
}

// NewDesignHeadDocParamsWithContext creates a new DesignHeadDocParams object
// with the default values initialized, and the ability to set a context for a request
func NewDesignHeadDocParamsWithContext(ctx context.Context) *DesignHeadDocParams {
	var ()
	return &DesignHeadDocParams{

		Context: ctx,
	}
}

// NewDesignHeadDocParamsWithHTTPClient creates a new DesignHeadDocParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDesignHeadDocParamsWithHTTPClient(client *http.Client) *DesignHeadDocParams {
	var ()
	return &DesignHeadDocParams{
		HTTPClient: client,
	}
}

/*DesignHeadDocParams contains all the parameters to send to the API endpoint
for the design head doc operation typically these are written to a http.Request
*/
type DesignHeadDocParams struct {

	/*Db
	  Database name

	*/
	Db string
	/*Ddoc
	  Design document id

	*/
	Ddoc string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the design head doc params
func (o *DesignHeadDocParams) WithTimeout(timeout time.Duration) *DesignHeadDocParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the design head doc params
func (o *DesignHeadDocParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the design head doc params
func (o *DesignHeadDocParams) WithContext(ctx context.Context) *DesignHeadDocParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the design head doc params
func (o *DesignHeadDocParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the design head doc params
func (o *DesignHeadDocParams) WithHTTPClient(client *http.Client) *DesignHeadDocParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the design head doc params
func (o *DesignHeadDocParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDb adds the db to the design head doc params
func (o *DesignHeadDocParams) WithDb(db string) *DesignHeadDocParams {
	o.SetDb(db)
	return o
}

// SetDb adds the db to the design head doc params
func (o *DesignHeadDocParams) SetDb(db string) {
	o.Db = db
}

// WithDdoc adds the ddoc to the design head doc params
func (o *DesignHeadDocParams) WithDdoc(ddoc string) *DesignHeadDocParams {
	o.SetDdoc(ddoc)
	return o
}

// SetDdoc adds the ddoc to the design head doc params
func (o *DesignHeadDocParams) SetDdoc(ddoc string) {
	o.Ddoc = ddoc
}

// WriteToRequest writes these params to a swagger request
func (o *DesignHeadDocParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param db
	if err := r.SetPathParam("db", o.Db); err != nil {
		return err
	}

	// path param ddoc
	if err := r.SetPathParam("ddoc", o.Ddoc); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
