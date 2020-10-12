// Code generated by go-swagger; DO NOT EDIT.

package document

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

// NewDocParams creates a new DocParams object
// with the default values initialized.
func NewDocParams() *DocParams {
	var ()
	return &DocParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDocParamsWithTimeout creates a new DocParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDocParamsWithTimeout(timeout time.Duration) *DocParams {
	var ()
	return &DocParams{

		timeout: timeout,
	}
}

// NewDocParamsWithContext creates a new DocParams object
// with the default values initialized, and the ability to set a context for a request
func NewDocParamsWithContext(ctx context.Context) *DocParams {
	var ()
	return &DocParams{

		Context: ctx,
	}
}

// NewDocParamsWithHTTPClient creates a new DocParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDocParamsWithHTTPClient(client *http.Client) *DocParams {
	var ()
	return &DocParams{
		HTTPClient: client,
	}
}

/*DocParams contains all the parameters to send to the API endpoint
for the doc operation typically these are written to a http.Request
*/
type DocParams struct {

	/*Db
	  Database name

	*/
	Db string
	/*Docid
	  DDocument ID

	*/
	Docid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the doc params
func (o *DocParams) WithTimeout(timeout time.Duration) *DocParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the doc params
func (o *DocParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the doc params
func (o *DocParams) WithContext(ctx context.Context) *DocParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the doc params
func (o *DocParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the doc params
func (o *DocParams) WithHTTPClient(client *http.Client) *DocParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the doc params
func (o *DocParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDb adds the db to the doc params
func (o *DocParams) WithDb(db string) *DocParams {
	o.SetDb(db)
	return o
}

// SetDb adds the db to the doc params
func (o *DocParams) SetDb(db string) {
	o.Db = db
}

// WithDocid adds the docid to the doc params
func (o *DocParams) WithDocid(docid string) *DocParams {
	o.SetDocid(docid)
	return o
}

// SetDocid adds the docid to the doc params
func (o *DocParams) SetDocid(docid string) {
	o.Docid = docid
}

// WriteToRequest writes these params to a swagger request
func (o *DocParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param db
	if err := r.SetPathParam("db", o.Db); err != nil {
		return err
	}

	// path param docid
	if err := r.SetPathParam("docid", o.Docid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}