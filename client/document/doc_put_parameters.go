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
	"github.com/go-openapi/swag"
)

// NewDocPutParams creates a new DocPutParams object
// with the default values initialized.
func NewDocPutParams() *DocPutParams {
	var ()
	return &DocPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDocPutParamsWithTimeout creates a new DocPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDocPutParamsWithTimeout(timeout time.Duration) *DocPutParams {
	var ()
	return &DocPutParams{

		timeout: timeout,
	}
}

// NewDocPutParamsWithContext creates a new DocPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewDocPutParamsWithContext(ctx context.Context) *DocPutParams {
	var ()
	return &DocPutParams{

		Context: ctx,
	}
}

// NewDocPutParamsWithHTTPClient creates a new DocPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDocPutParamsWithHTTPClient(client *http.Client) *DocPutParams {
	var ()
	return &DocPutParams{
		HTTPClient: client,
	}
}

/*DocPutParams contains all the parameters to send to the API endpoint
for the doc put operation typically these are written to a http.Request
*/
type DocPutParams struct {

	/*IfMatch
	  Document’s revision. Alternative to rev query parameter or document key. Optional

	*/
	IfMatch *string
	/*Batch
	  Stores document in batch mode. Possible values: ok. Optional


	*/
	Batch *string
	/*Db
	  Database name

	*/
	Db string
	/*Docid
	  DDocument ID

	*/
	Docid string
	/*NewEdits
	  Prevents insertion of a conflicting document. Possible values: true (default) and false. If false,
	a well-formed _rev must be included in the document. new_edits=false is used by the replicator
	to insert documents into the target database even if that leads to the creation of conflicts. Optional


	*/
	NewEdits *bool
	/*Rev
	  Document’s revision if updating an existing document. Alternative to If-Match header or document key. Optional

	*/
	Rev *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the doc put params
func (o *DocPutParams) WithTimeout(timeout time.Duration) *DocPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the doc put params
func (o *DocPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the doc put params
func (o *DocPutParams) WithContext(ctx context.Context) *DocPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the doc put params
func (o *DocPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the doc put params
func (o *DocPutParams) WithHTTPClient(client *http.Client) *DocPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the doc put params
func (o *DocPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the doc put params
func (o *DocPutParams) WithIfMatch(ifMatch *string) *DocPutParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the doc put params
func (o *DocPutParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithBatch adds the batch to the doc put params
func (o *DocPutParams) WithBatch(batch *string) *DocPutParams {
	o.SetBatch(batch)
	return o
}

// SetBatch adds the batch to the doc put params
func (o *DocPutParams) SetBatch(batch *string) {
	o.Batch = batch
}

// WithDb adds the db to the doc put params
func (o *DocPutParams) WithDb(db string) *DocPutParams {
	o.SetDb(db)
	return o
}

// SetDb adds the db to the doc put params
func (o *DocPutParams) SetDb(db string) {
	o.Db = db
}

// WithDocid adds the docid to the doc put params
func (o *DocPutParams) WithDocid(docid string) *DocPutParams {
	o.SetDocid(docid)
	return o
}

// SetDocid adds the docid to the doc put params
func (o *DocPutParams) SetDocid(docid string) {
	o.Docid = docid
}

// WithNewEdits adds the newEdits to the doc put params
func (o *DocPutParams) WithNewEdits(newEdits *bool) *DocPutParams {
	o.SetNewEdits(newEdits)
	return o
}

// SetNewEdits adds the newEdits to the doc put params
func (o *DocPutParams) SetNewEdits(newEdits *bool) {
	o.NewEdits = newEdits
}

// WithRev adds the rev to the doc put params
func (o *DocPutParams) WithRev(rev *string) *DocPutParams {
	o.SetRev(rev)
	return o
}

// SetRev adds the rev to the doc put params
func (o *DocPutParams) SetRev(rev *string) {
	o.Rev = rev
}

// WriteToRequest writes these params to a swagger request
func (o *DocPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfMatch != nil {

		// header param If-Match
		if err := r.SetHeaderParam("If-Match", *o.IfMatch); err != nil {
			return err
		}

	}

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

	// path param db
	if err := r.SetPathParam("db", o.Db); err != nil {
		return err
	}

	// path param docid
	if err := r.SetPathParam("docid", o.Docid); err != nil {
		return err
	}

	if o.NewEdits != nil {

		// query param new_edits
		var qrNewEdits bool
		if o.NewEdits != nil {
			qrNewEdits = *o.NewEdits
		}
		qNewEdits := swag.FormatBool(qrNewEdits)
		if qNewEdits != "" {
			if err := r.SetQueryParam("new_edits", qNewEdits); err != nil {
				return err
			}
		}

	}

	if o.Rev != nil {

		// query param rev
		var qrRev string
		if o.Rev != nil {
			qrRev = *o.Rev
		}
		qRev := qrRev
		if qRev != "" {
			if err := r.SetQueryParam("rev", qRev); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
