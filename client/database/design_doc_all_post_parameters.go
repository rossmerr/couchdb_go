// Code generated by go-swagger; DO NOT EDIT.

package database

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

	"github.com/rossmerr/couchdb_go/models"
)

// NewDesignDocAllPostParams creates a new DesignDocAllPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDesignDocAllPostParams() *DesignDocAllPostParams {
	return &DesignDocAllPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDesignDocAllPostParamsWithTimeout creates a new DesignDocAllPostParams object
// with the ability to set a timeout on a request.
func NewDesignDocAllPostParamsWithTimeout(timeout time.Duration) *DesignDocAllPostParams {
	return &DesignDocAllPostParams{
		timeout: timeout,
	}
}

// NewDesignDocAllPostParamsWithContext creates a new DesignDocAllPostParams object
// with the ability to set a context for a request.
func NewDesignDocAllPostParamsWithContext(ctx context.Context) *DesignDocAllPostParams {
	return &DesignDocAllPostParams{
		Context: ctx,
	}
}

// NewDesignDocAllPostParamsWithHTTPClient creates a new DesignDocAllPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewDesignDocAllPostParamsWithHTTPClient(client *http.Client) *DesignDocAllPostParams {
	return &DesignDocAllPostParams{
		HTTPClient: client,
	}
}

/* DesignDocAllPostParams contains all the parameters to send to the API endpoint
   for the design doc all post operation.

   Typically these are written to a http.Request.
*/
type DesignDocAllPostParams struct {

	// Body.
	Body *models.Body4

	/* Conflicts.

	   Include conflicts information in response. Ignored if include_docs isn’t true. Default is false.
	*/
	Conflicts *bool

	/* Db.

	   Database name
	*/
	Db string

	/* Descending.

	   Return the documents in descending order by key. Default is false.
	*/
	Descending *bool

	/* EndKey.

	   Alias for endkey param.
	*/
	EndKey *string

	/* EndKeyDocID.

	   Alias for endkey_docid param.
	*/
	EndKeyDocID *string

	/* Endkey.

	   Stop returning records when the specified key is reached. Optional.
	*/
	QueryEndKey *string

	/* EndkeyDocid.

	   Stop returning records when the specified design document ID is reached. Optional.
	*/
	EndkeyDocid *string

	/* IncludeDocs.

	   Include the full content of the design documents in the return. Default is false.
	*/
	IncludeDocs *bool

	/* InclusiveEnd.

	   Specifies whether the specified end key should be included in the result. Default is true.
	*/
	InclusiveEnd *bool

	/* Key.

	   Return only design documents that match the specified key. Optional.
	*/
	Key *string

	/* Keys.

	   Return only design documents that match the specified keys. Optional.
	*/
	Keys *string

	/* Limit.

	   Limit the number of the returned design documents to the specified number. Optional.
	*/
	Limit *int64

	/* Skip.

	   Skip this number of records before starting to return the results. Default is 0.
	*/
	Skip *int64

	/* StartKey.

	   Alias for startkey param.
	*/
	StartKey *string

	/* StartKeyDocID.

	   Alias for startkey_docid param
	*/
	StartKeyDocID *string

	/* Startkey.

	   Return records starting with the specified key. Optional.
	*/
	QueryStartKey *string

	/* StartkeyDocid.

	   Return records starting with the specified design document ID. Optional.
	*/
	StartkeyDocid *string

	/* UpdateSeq.

	   Whether to include in the response an update_seq value indicating the sequence id of the database the view reflects. Default is false.
	*/
	UpdateSeq *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the design doc all post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DesignDocAllPostParams) WithDefaults() *DesignDocAllPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the design doc all post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DesignDocAllPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the design doc all post params
func (o *DesignDocAllPostParams) WithTimeout(timeout time.Duration) *DesignDocAllPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the design doc all post params
func (o *DesignDocAllPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the design doc all post params
func (o *DesignDocAllPostParams) WithContext(ctx context.Context) *DesignDocAllPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the design doc all post params
func (o *DesignDocAllPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the design doc all post params
func (o *DesignDocAllPostParams) WithHTTPClient(client *http.Client) *DesignDocAllPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the design doc all post params
func (o *DesignDocAllPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the design doc all post params
func (o *DesignDocAllPostParams) WithBody(body *models.Body4) *DesignDocAllPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the design doc all post params
func (o *DesignDocAllPostParams) SetBody(body *models.Body4) {
	o.Body = body
}

// WithConflicts adds the conflicts to the design doc all post params
func (o *DesignDocAllPostParams) WithConflicts(conflicts *bool) *DesignDocAllPostParams {
	o.SetConflicts(conflicts)
	return o
}

// SetConflicts adds the conflicts to the design doc all post params
func (o *DesignDocAllPostParams) SetConflicts(conflicts *bool) {
	o.Conflicts = conflicts
}

// WithDb adds the db to the design doc all post params
func (o *DesignDocAllPostParams) WithDb(db string) *DesignDocAllPostParams {
	o.SetDb(db)
	return o
}

// SetDb adds the db to the design doc all post params
func (o *DesignDocAllPostParams) SetDb(db string) {
	o.Db = db
}

// WithDescending adds the descending to the design doc all post params
func (o *DesignDocAllPostParams) WithDescending(descending *bool) *DesignDocAllPostParams {
	o.SetDescending(descending)
	return o
}

// SetDescending adds the descending to the design doc all post params
func (o *DesignDocAllPostParams) SetDescending(descending *bool) {
	o.Descending = descending
}

// WithEndKey adds the endKey to the design doc all post params
func (o *DesignDocAllPostParams) WithEndKey(endKey *string) *DesignDocAllPostParams {
	o.SetEndKey(endKey)
	return o
}

// SetEndKey adds the endKey to the design doc all post params
func (o *DesignDocAllPostParams) SetEndKey(endKey *string) {
	o.EndKey = endKey
}

// WithEndKeyDocID adds the endKeyDocID to the design doc all post params
func (o *DesignDocAllPostParams) WithEndKeyDocID(endKeyDocID *string) *DesignDocAllPostParams {
	o.SetEndKeyDocID(endKeyDocID)
	return o
}

// SetEndKeyDocID adds the endKeyDocId to the design doc all post params
func (o *DesignDocAllPostParams) SetEndKeyDocID(endKeyDocID *string) {
	o.EndKeyDocID = endKeyDocID
}

// WithQueryEndKey adds the endkey to the design doc all post params
func (o *DesignDocAllPostParams) WithQueryEndKey(endkey *string) *DesignDocAllPostParams {
	o.SetQueryEndKey(endkey)
	return o
}

// SetQueryEndKey adds the endkey to the design doc all post params
func (o *DesignDocAllPostParams) SetQueryEndKey(endkey *string) {
	o.QueryEndKey = endkey
}

// WithEndkeyDocid adds the endkeyDocid to the design doc all post params
func (o *DesignDocAllPostParams) WithEndkeyDocid(endkeyDocid *string) *DesignDocAllPostParams {
	o.SetEndkeyDocid(endkeyDocid)
	return o
}

// SetEndkeyDocid adds the endkeyDocid to the design doc all post params
func (o *DesignDocAllPostParams) SetEndkeyDocid(endkeyDocid *string) {
	o.EndkeyDocid = endkeyDocid
}

// WithIncludeDocs adds the includeDocs to the design doc all post params
func (o *DesignDocAllPostParams) WithIncludeDocs(includeDocs *bool) *DesignDocAllPostParams {
	o.SetIncludeDocs(includeDocs)
	return o
}

// SetIncludeDocs adds the includeDocs to the design doc all post params
func (o *DesignDocAllPostParams) SetIncludeDocs(includeDocs *bool) {
	o.IncludeDocs = includeDocs
}

// WithInclusiveEnd adds the inclusiveEnd to the design doc all post params
func (o *DesignDocAllPostParams) WithInclusiveEnd(inclusiveEnd *bool) *DesignDocAllPostParams {
	o.SetInclusiveEnd(inclusiveEnd)
	return o
}

// SetInclusiveEnd adds the inclusiveEnd to the design doc all post params
func (o *DesignDocAllPostParams) SetInclusiveEnd(inclusiveEnd *bool) {
	o.InclusiveEnd = inclusiveEnd
}

// WithKey adds the key to the design doc all post params
func (o *DesignDocAllPostParams) WithKey(key *string) *DesignDocAllPostParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the design doc all post params
func (o *DesignDocAllPostParams) SetKey(key *string) {
	o.Key = key
}

// WithKeys adds the keys to the design doc all post params
func (o *DesignDocAllPostParams) WithKeys(keys *string) *DesignDocAllPostParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the design doc all post params
func (o *DesignDocAllPostParams) SetKeys(keys *string) {
	o.Keys = keys
}

// WithLimit adds the limit to the design doc all post params
func (o *DesignDocAllPostParams) WithLimit(limit *int64) *DesignDocAllPostParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the design doc all post params
func (o *DesignDocAllPostParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithSkip adds the skip to the design doc all post params
func (o *DesignDocAllPostParams) WithSkip(skip *int64) *DesignDocAllPostParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the design doc all post params
func (o *DesignDocAllPostParams) SetSkip(skip *int64) {
	o.Skip = skip
}

// WithStartKey adds the startKey to the design doc all post params
func (o *DesignDocAllPostParams) WithStartKey(startKey *string) *DesignDocAllPostParams {
	o.SetStartKey(startKey)
	return o
}

// SetStartKey adds the startKey to the design doc all post params
func (o *DesignDocAllPostParams) SetStartKey(startKey *string) {
	o.StartKey = startKey
}

// WithStartKeyDocID adds the startKeyDocID to the design doc all post params
func (o *DesignDocAllPostParams) WithStartKeyDocID(startKeyDocID *string) *DesignDocAllPostParams {
	o.SetStartKeyDocID(startKeyDocID)
	return o
}

// SetStartKeyDocID adds the startKeyDocId to the design doc all post params
func (o *DesignDocAllPostParams) SetStartKeyDocID(startKeyDocID *string) {
	o.StartKeyDocID = startKeyDocID
}

// WithQueryStartKey adds the startkey to the design doc all post params
func (o *DesignDocAllPostParams) WithQueryStartKey(startkey *string) *DesignDocAllPostParams {
	o.SetQueryStartKey(startkey)
	return o
}

// SetQueryStartKey adds the startkey to the design doc all post params
func (o *DesignDocAllPostParams) SetQueryStartKey(startkey *string) {
	o.QueryStartKey = startkey
}

// WithStartkeyDocid adds the startkeyDocid to the design doc all post params
func (o *DesignDocAllPostParams) WithStartkeyDocid(startkeyDocid *string) *DesignDocAllPostParams {
	o.SetStartkeyDocid(startkeyDocid)
	return o
}

// SetStartkeyDocid adds the startkeyDocid to the design doc all post params
func (o *DesignDocAllPostParams) SetStartkeyDocid(startkeyDocid *string) {
	o.StartkeyDocid = startkeyDocid
}

// WithUpdateSeq adds the updateSeq to the design doc all post params
func (o *DesignDocAllPostParams) WithUpdateSeq(updateSeq *bool) *DesignDocAllPostParams {
	o.SetUpdateSeq(updateSeq)
	return o
}

// SetUpdateSeq adds the updateSeq to the design doc all post params
func (o *DesignDocAllPostParams) SetUpdateSeq(updateSeq *bool) {
	o.UpdateSeq = updateSeq
}

// WriteToRequest writes these params to a swagger request
func (o *DesignDocAllPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Conflicts != nil {

		// query param conflicts
		var qrConflicts bool

		if o.Conflicts != nil {
			qrConflicts = *o.Conflicts
		}
		qConflicts := swag.FormatBool(qrConflicts)
		if qConflicts != "" {

			if err := r.SetQueryParam("conflicts", qConflicts); err != nil {
				return err
			}
		}
	}

	// path param db
	if err := r.SetPathParam("db", o.Db); err != nil {
		return err
	}

	if o.Descending != nil {

		// query param descending
		var qrDescending bool

		if o.Descending != nil {
			qrDescending = *o.Descending
		}
		qDescending := swag.FormatBool(qrDescending)
		if qDescending != "" {

			if err := r.SetQueryParam("descending", qDescending); err != nil {
				return err
			}
		}
	}

	if o.EndKey != nil {

		// query param end_key
		var qrEndKey string

		if o.EndKey != nil {
			qrEndKey = *o.EndKey
		}
		qEndKey := qrEndKey
		if qEndKey != "" {

			if err := r.SetQueryParam("end_key", qEndKey); err != nil {
				return err
			}
		}
	}

	if o.EndKeyDocID != nil {

		// query param end_key_doc_id
		var qrEndKeyDocID string

		if o.EndKeyDocID != nil {
			qrEndKeyDocID = *o.EndKeyDocID
		}
		qEndKeyDocID := qrEndKeyDocID
		if qEndKeyDocID != "" {

			if err := r.SetQueryParam("end_key_doc_id", qEndKeyDocID); err != nil {
				return err
			}
		}
	}

	if o.QueryEndKey != nil {

		// query param endkey
		var qrEndkey string

		if o.QueryEndKey != nil {
			qrEndkey = *o.QueryEndKey
		}
		qEndkey := qrEndkey
		if qEndkey != "" {

			if err := r.SetQueryParam("endkey", qEndkey); err != nil {
				return err
			}
		}
	}

	if o.EndkeyDocid != nil {

		// query param endkey_docid
		var qrEndkeyDocid string

		if o.EndkeyDocid != nil {
			qrEndkeyDocid = *o.EndkeyDocid
		}
		qEndkeyDocid := qrEndkeyDocid
		if qEndkeyDocid != "" {

			if err := r.SetQueryParam("endkey_docid", qEndkeyDocid); err != nil {
				return err
			}
		}
	}

	if o.IncludeDocs != nil {

		// query param include_docs
		var qrIncludeDocs bool

		if o.IncludeDocs != nil {
			qrIncludeDocs = *o.IncludeDocs
		}
		qIncludeDocs := swag.FormatBool(qrIncludeDocs)
		if qIncludeDocs != "" {

			if err := r.SetQueryParam("include_docs", qIncludeDocs); err != nil {
				return err
			}
		}
	}

	if o.InclusiveEnd != nil {

		// query param inclusive_end
		var qrInclusiveEnd bool

		if o.InclusiveEnd != nil {
			qrInclusiveEnd = *o.InclusiveEnd
		}
		qInclusiveEnd := swag.FormatBool(qrInclusiveEnd)
		if qInclusiveEnd != "" {

			if err := r.SetQueryParam("inclusive_end", qInclusiveEnd); err != nil {
				return err
			}
		}
	}

	if o.Key != nil {

		// query param key
		var qrKey string

		if o.Key != nil {
			qrKey = *o.Key
		}
		qKey := qrKey
		if qKey != "" {

			if err := r.SetQueryParam("key", qKey); err != nil {
				return err
			}
		}
	}

	if o.Keys != nil {

		// query param keys
		var qrKeys string

		if o.Keys != nil {
			qrKeys = *o.Keys
		}
		qKeys := qrKeys
		if qKeys != "" {

			if err := r.SetQueryParam("keys", qKeys); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Skip != nil {

		// query param skip
		var qrSkip int64

		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt64(qrSkip)
		if qSkip != "" {

			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}
	}

	if o.StartKey != nil {

		// query param start_key
		var qrStartKey string

		if o.StartKey != nil {
			qrStartKey = *o.StartKey
		}
		qStartKey := qrStartKey
		if qStartKey != "" {

			if err := r.SetQueryParam("start_key", qStartKey); err != nil {
				return err
			}
		}
	}

	if o.StartKeyDocID != nil {

		// query param start_key_doc_id
		var qrStartKeyDocID string

		if o.StartKeyDocID != nil {
			qrStartKeyDocID = *o.StartKeyDocID
		}
		qStartKeyDocID := qrStartKeyDocID
		if qStartKeyDocID != "" {

			if err := r.SetQueryParam("start_key_doc_id", qStartKeyDocID); err != nil {
				return err
			}
		}
	}

	if o.QueryStartKey != nil {

		// query param startkey
		var qrStartkey string

		if o.QueryStartKey != nil {
			qrStartkey = *o.QueryStartKey
		}
		qStartkey := qrStartkey
		if qStartkey != "" {

			if err := r.SetQueryParam("startkey", qStartkey); err != nil {
				return err
			}
		}
	}

	if o.StartkeyDocid != nil {

		// query param startkey_docid
		var qrStartkeyDocid string

		if o.StartkeyDocid != nil {
			qrStartkeyDocid = *o.StartkeyDocid
		}
		qStartkeyDocid := qrStartkeyDocid
		if qStartkeyDocid != "" {

			if err := r.SetQueryParam("startkey_docid", qStartkeyDocid); err != nil {
				return err
			}
		}
	}

	if o.UpdateSeq != nil {

		// query param update_seq
		var qrUpdateSeq bool

		if o.UpdateSeq != nil {
			qrUpdateSeq = *o.UpdateSeq
		}
		qUpdateSeq := swag.FormatBool(qrUpdateSeq)
		if qUpdateSeq != "" {

			if err := r.SetQueryParam("update_seq", qUpdateSeq); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
