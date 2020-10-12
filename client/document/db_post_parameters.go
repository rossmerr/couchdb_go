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

	"github.com/RossMerr/couchdb_go/models"
)

// NewDbPostParams creates a new DbPostParams object
// with the default values initialized.
func NewDbPostParams() *DbPostParams {
	var ()
	return &DbPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDbPostParamsWithTimeout creates a new DbPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDbPostParamsWithTimeout(timeout time.Duration) *DbPostParams {
	var ()
	return &DbPostParams{

		timeout: timeout,
	}
}

// NewDbPostParamsWithContext creates a new DbPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewDbPostParamsWithContext(ctx context.Context) *DbPostParams {
	var ()
	return &DbPostParams{

		Context: ctx,
	}
}

// NewDbPostParamsWithHTTPClient creates a new DbPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDbPostParamsWithHTTPClient(client *http.Client) *DbPostParams {
	var ()
	return &DbPostParams{
		HTTPClient: client,
	}
}

/*DbPostParams contains all the parameters to send to the API endpoint
for the db post operation typically these are written to a http.Request
*/
type DbPostParams struct {

	/*Batch
	  Stores document in batch mode Possible values: ok. Optional


	*/
	Batch *string
	/*Body*/
	Body models.Document
	/*Db
	  Database name

	*/
	Db string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the db post params
func (o *DbPostParams) WithTimeout(timeout time.Duration) *DbPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the db post params
func (o *DbPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the db post params
func (o *DbPostParams) WithContext(ctx context.Context) *DbPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the db post params
func (o *DbPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the db post params
func (o *DbPostParams) WithHTTPClient(client *http.Client) *DbPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the db post params
func (o *DbPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBatch adds the batch to the db post params
func (o *DbPostParams) WithBatch(batch *string) *DbPostParams {
	o.SetBatch(batch)
	return o
}

// SetBatch adds the batch to the db post params
func (o *DbPostParams) SetBatch(batch *string) {
	o.Batch = batch
}

// WithBody adds the body to the db post params
func (o *DbPostParams) WithBody(body models.Document) *DbPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the db post params
func (o *DbPostParams) SetBody(body models.Document) {
	o.Body = body
}

// WithDb adds the db to the db post params
func (o *DbPostParams) WithDb(db string) *DbPostParams {
	o.SetDb(db)
	return o
}

// SetDb adds the db to the db post params
func (o *DbPostParams) SetDb(db string) {
	o.Db = db
}

// WriteToRequest writes these params to a swagger request
func (o *DbPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Batch != nil {

		// query param batch
		var qrBatch string
		if o.Batch != nil {
			qrBatch = *o.Batch
		}
		qBatch := qrBatch
		if qBatch != "" {
			if err := r.SetQueryParam("batch", qBatch); err != nil {
				return err
			}
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param db
	if err := r.SetPathParam("db", o.Db); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}