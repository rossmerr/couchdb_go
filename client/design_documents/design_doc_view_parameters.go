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
	"github.com/go-openapi/swag"
)

// NewDesignDocViewParams creates a new DesignDocViewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDesignDocViewParams() *DesignDocViewParams {
	return &DesignDocViewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDesignDocViewParamsWithTimeout creates a new DesignDocViewParams object
// with the ability to set a timeout on a request.
func NewDesignDocViewParamsWithTimeout(timeout time.Duration) *DesignDocViewParams {
	return &DesignDocViewParams{
		timeout: timeout,
	}
}

// NewDesignDocViewParamsWithContext creates a new DesignDocViewParams object
// with the ability to set a context for a request.
func NewDesignDocViewParamsWithContext(ctx context.Context) *DesignDocViewParams {
	return &DesignDocViewParams{
		Context: ctx,
	}
}

// NewDesignDocViewParamsWithHTTPClient creates a new DesignDocViewParams object
// with the ability to set a custom HTTPClient for a request.
func NewDesignDocViewParamsWithHTTPClient(client *http.Client) *DesignDocViewParams {
	return &DesignDocViewParams{
		HTTPClient: client,
	}
}

/* DesignDocViewParams contains all the parameters to send to the API endpoint
   for the design doc view operation.

   Typically these are written to a http.Request.
*/
type DesignDocViewParams struct {

	/* AttEncodingInfo.

	   Include encoding information in attachment stubs if include_docs is true and the particular attachment is compressed. Ignored if include_docs isn’t true. Default is false.
	*/
	AttEncodingInfo *bool

	/* Attachments.

	   Include the Base64-encoded content of attachments in the documents that are included if include_docs is true. Ignored if include_docs isn’t true. Default is false.
	*/
	Attachments *bool

	/* Conflicts.

	   Include conflicts information in response. Ignored if include_docs isn’t true. Default is false.
	*/
	Conflicts *bool

	/* Db.

	   Database name
	*/
	Db string

	/* Ddoc.

	   Design document id
	*/
	Ddoc string

	/* Descending.

	   Return the documents in descending order by key. Default is false.
	*/
	Descending *bool

	/* EndKey.

	   Alias for endkey param
	*/
	EndKey *string

	/* EndKeyDocID.

	   Alias for endkey_docid.
	*/
	EndKeyDocID *string

	/* Endkey.

	   Stop returning records when the specified key is reached.
	*/
	QueryEndKey *string

	/* EndkeyDocid.

	   Stop returning records when the specified document ID is reached. Ignored if endkey is not set.
	*/
	EndkeyDocid *string

	/* Group.

	   Group the results using the reduce function to a group or single row. Implies reduce is true and the maximum group_level. Default is false.
	*/
	Group *bool

	/* GroupLevel.

	   Specify the group level to be used. Implies group is true.
	*/
	GroupLevel *int64

	/* IncludeDocs.

	   Include the associated document with each row. Default is false.
	*/
	IncludeDocs *bool

	/* InclusiveEnd.

	   Specifies whether the specified end key should be included in the result. Default is true.
	*/
	InclusiveEnd *bool

	/* Key.

	   eturn only documents that match the specified key.
	*/
	Key *string

	/* Keys.

	   Return only documents where the key matches one of the keys specified in the array.
	*/
	Keys []string

	/* Limit.

	   Limit the number of the returned documents to the specified number.
	*/
	Limit *int64

	/* Reduce.

	   Use the reduction function. Default is true when a reduce function is defined.
	*/
	Reduce *bool

	/* RevsInfo.

	   Includes detailed information for all known document revisions. Default is false
	*/
	RevsInfo *bool

	/* Skip.

	   Skip this number of records before starting to return the results. Default is 0.
	*/
	Skip *int64

	/* Sorted.

	   Sort returned rows (see Sorting Returned Rows). Setting this to false offers a performance boost. The total_rows and offset fields are not available when this is set to false. Default is true.
	*/
	Sorted *bool

	/* Stable.

	   Whether or not the view results should be returned from a stable set of shards. Default is false.
	*/
	Stable *bool

	/* Stale.

	   Allow the results from a stale view to be used. Supported values: ok and update_after. ok is equivalent to stable=true&update=false. update_after is equivalent to stable=true&update=lazy. The default behavior is equivalent to stable=false&update=true. Note that this parameter is deprecated. Use stable and update instead. See Views Generation for more details.

	*/
	Stale *string

	/* StartKey.

	   Alias for startkey.
	*/
	StartKey *string

	/* StartKeyDocID.

	   Alias for startkey_docid param
	*/
	StartKeyDocID *string

	/* Startkey.

	   Return records starting with the specified key.
	*/
	QueryStartKey *string

	/* StartkeyDocid.

	   Return records starting with the specified document ID. Ignored if startkey is not set.
	*/
	StartkeyDocid *string

	/* Update.

	   Whether or not the view in question should be updated prior to responding to the user. Supported values: true, false, lazy. Default is true.

	*/
	Update *string

	/* UpdateSeq.

	   Whether to include in the response an update_seq value indicating the sequence id of the database the view reflects. Default is false.
	*/
	UpdateSeq *bool

	/* View.

	   View function name
	*/
	View string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the design doc view params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DesignDocViewParams) WithDefaults() *DesignDocViewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the design doc view params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DesignDocViewParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the design doc view params
func (o *DesignDocViewParams) WithTimeout(timeout time.Duration) *DesignDocViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the design doc view params
func (o *DesignDocViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the design doc view params
func (o *DesignDocViewParams) WithContext(ctx context.Context) *DesignDocViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the design doc view params
func (o *DesignDocViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the design doc view params
func (o *DesignDocViewParams) WithHTTPClient(client *http.Client) *DesignDocViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the design doc view params
func (o *DesignDocViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttEncodingInfo adds the attEncodingInfo to the design doc view params
func (o *DesignDocViewParams) WithAttEncodingInfo(attEncodingInfo *bool) *DesignDocViewParams {
	o.SetAttEncodingInfo(attEncodingInfo)
	return o
}

// SetAttEncodingInfo adds the attEncodingInfo to the design doc view params
func (o *DesignDocViewParams) SetAttEncodingInfo(attEncodingInfo *bool) {
	o.AttEncodingInfo = attEncodingInfo
}

// WithAttachments adds the attachments to the design doc view params
func (o *DesignDocViewParams) WithAttachments(attachments *bool) *DesignDocViewParams {
	o.SetAttachments(attachments)
	return o
}

// SetAttachments adds the attachments to the design doc view params
func (o *DesignDocViewParams) SetAttachments(attachments *bool) {
	o.Attachments = attachments
}

// WithConflicts adds the conflicts to the design doc view params
func (o *DesignDocViewParams) WithConflicts(conflicts *bool) *DesignDocViewParams {
	o.SetConflicts(conflicts)
	return o
}

// SetConflicts adds the conflicts to the design doc view params
func (o *DesignDocViewParams) SetConflicts(conflicts *bool) {
	o.Conflicts = conflicts
}

// WithDb adds the db to the design doc view params
func (o *DesignDocViewParams) WithDb(db string) *DesignDocViewParams {
	o.SetDb(db)
	return o
}

// SetDb adds the db to the design doc view params
func (o *DesignDocViewParams) SetDb(db string) {
	o.Db = db
}

// WithDdoc adds the ddoc to the design doc view params
func (o *DesignDocViewParams) WithDdoc(ddoc string) *DesignDocViewParams {
	o.SetDdoc(ddoc)
	return o
}

// SetDdoc adds the ddoc to the design doc view params
func (o *DesignDocViewParams) SetDdoc(ddoc string) {
	o.Ddoc = ddoc
}

// WithDescending adds the descending to the design doc view params
func (o *DesignDocViewParams) WithDescending(descending *bool) *DesignDocViewParams {
	o.SetDescending(descending)
	return o
}

// SetDescending adds the descending to the design doc view params
func (o *DesignDocViewParams) SetDescending(descending *bool) {
	o.Descending = descending
}

// WithEndKey adds the endKey to the design doc view params
func (o *DesignDocViewParams) WithEndKey(endKey *string) *DesignDocViewParams {
	o.SetEndKey(endKey)
	return o
}

// SetEndKey adds the endKey to the design doc view params
func (o *DesignDocViewParams) SetEndKey(endKey *string) {
	o.EndKey = endKey
}

// WithEndKeyDocID adds the endKeyDocID to the design doc view params
func (o *DesignDocViewParams) WithEndKeyDocID(endKeyDocID *string) *DesignDocViewParams {
	o.SetEndKeyDocID(endKeyDocID)
	return o
}

// SetEndKeyDocID adds the endKeyDocId to the design doc view params
func (o *DesignDocViewParams) SetEndKeyDocID(endKeyDocID *string) {
	o.EndKeyDocID = endKeyDocID
}

// WithQueryEndKey adds the endkey to the design doc view params
func (o *DesignDocViewParams) WithQueryEndKey(endkey *string) *DesignDocViewParams {
	o.SetQueryEndKey(endkey)
	return o
}

// SetQueryEndKey adds the endkey to the design doc view params
func (o *DesignDocViewParams) SetQueryEndKey(endkey *string) {
	o.QueryEndKey = endkey
}

// WithEndkeyDocid adds the endkeyDocid to the design doc view params
func (o *DesignDocViewParams) WithEndkeyDocid(endkeyDocid *string) *DesignDocViewParams {
	o.SetEndkeyDocid(endkeyDocid)
	return o
}

// SetEndkeyDocid adds the endkeyDocid to the design doc view params
func (o *DesignDocViewParams) SetEndkeyDocid(endkeyDocid *string) {
	o.EndkeyDocid = endkeyDocid
}

// WithGroup adds the group to the design doc view params
func (o *DesignDocViewParams) WithGroup(group *bool) *DesignDocViewParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the design doc view params
func (o *DesignDocViewParams) SetGroup(group *bool) {
	o.Group = group
}

// WithGroupLevel adds the groupLevel to the design doc view params
func (o *DesignDocViewParams) WithGroupLevel(groupLevel *int64) *DesignDocViewParams {
	o.SetGroupLevel(groupLevel)
	return o
}

// SetGroupLevel adds the groupLevel to the design doc view params
func (o *DesignDocViewParams) SetGroupLevel(groupLevel *int64) {
	o.GroupLevel = groupLevel
}

// WithIncludeDocs adds the includeDocs to the design doc view params
func (o *DesignDocViewParams) WithIncludeDocs(includeDocs *bool) *DesignDocViewParams {
	o.SetIncludeDocs(includeDocs)
	return o
}

// SetIncludeDocs adds the includeDocs to the design doc view params
func (o *DesignDocViewParams) SetIncludeDocs(includeDocs *bool) {
	o.IncludeDocs = includeDocs
}

// WithInclusiveEnd adds the inclusiveEnd to the design doc view params
func (o *DesignDocViewParams) WithInclusiveEnd(inclusiveEnd *bool) *DesignDocViewParams {
	o.SetInclusiveEnd(inclusiveEnd)
	return o
}

// SetInclusiveEnd adds the inclusiveEnd to the design doc view params
func (o *DesignDocViewParams) SetInclusiveEnd(inclusiveEnd *bool) {
	o.InclusiveEnd = inclusiveEnd
}

// WithKey adds the key to the design doc view params
func (o *DesignDocViewParams) WithKey(key *string) *DesignDocViewParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the design doc view params
func (o *DesignDocViewParams) SetKey(key *string) {
	o.Key = key
}

// WithKeys adds the keys to the design doc view params
func (o *DesignDocViewParams) WithKeys(keys []string) *DesignDocViewParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the design doc view params
func (o *DesignDocViewParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WithLimit adds the limit to the design doc view params
func (o *DesignDocViewParams) WithLimit(limit *int64) *DesignDocViewParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the design doc view params
func (o *DesignDocViewParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithReduce adds the reduce to the design doc view params
func (o *DesignDocViewParams) WithReduce(reduce *bool) *DesignDocViewParams {
	o.SetReduce(reduce)
	return o
}

// SetReduce adds the reduce to the design doc view params
func (o *DesignDocViewParams) SetReduce(reduce *bool) {
	o.Reduce = reduce
}

// WithRevsInfo adds the revsInfo to the design doc view params
func (o *DesignDocViewParams) WithRevsInfo(revsInfo *bool) *DesignDocViewParams {
	o.SetRevsInfo(revsInfo)
	return o
}

// SetRevsInfo adds the revsInfo to the design doc view params
func (o *DesignDocViewParams) SetRevsInfo(revsInfo *bool) {
	o.RevsInfo = revsInfo
}

// WithSkip adds the skip to the design doc view params
func (o *DesignDocViewParams) WithSkip(skip *int64) *DesignDocViewParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the design doc view params
func (o *DesignDocViewParams) SetSkip(skip *int64) {
	o.Skip = skip
}

// WithSorted adds the sorted to the design doc view params
func (o *DesignDocViewParams) WithSorted(sorted *bool) *DesignDocViewParams {
	o.SetSorted(sorted)
	return o
}

// SetSorted adds the sorted to the design doc view params
func (o *DesignDocViewParams) SetSorted(sorted *bool) {
	o.Sorted = sorted
}

// WithStable adds the stable to the design doc view params
func (o *DesignDocViewParams) WithStable(stable *bool) *DesignDocViewParams {
	o.SetStable(stable)
	return o
}

// SetStable adds the stable to the design doc view params
func (o *DesignDocViewParams) SetStable(stable *bool) {
	o.Stable = stable
}

// WithStale adds the stale to the design doc view params
func (o *DesignDocViewParams) WithStale(stale *string) *DesignDocViewParams {
	o.SetStale(stale)
	return o
}

// SetStale adds the stale to the design doc view params
func (o *DesignDocViewParams) SetStale(stale *string) {
	o.Stale = stale
}

// WithStartKey adds the startKey to the design doc view params
func (o *DesignDocViewParams) WithStartKey(startKey *string) *DesignDocViewParams {
	o.SetStartKey(startKey)
	return o
}

// SetStartKey adds the startKey to the design doc view params
func (o *DesignDocViewParams) SetStartKey(startKey *string) {
	o.StartKey = startKey
}

// WithStartKeyDocID adds the startKeyDocID to the design doc view params
func (o *DesignDocViewParams) WithStartKeyDocID(startKeyDocID *string) *DesignDocViewParams {
	o.SetStartKeyDocID(startKeyDocID)
	return o
}

// SetStartKeyDocID adds the startKeyDocId to the design doc view params
func (o *DesignDocViewParams) SetStartKeyDocID(startKeyDocID *string) {
	o.StartKeyDocID = startKeyDocID
}

// WithQueryStartKey adds the startkey to the design doc view params
func (o *DesignDocViewParams) WithQueryStartKey(startkey *string) *DesignDocViewParams {
	o.SetQueryStartKey(startkey)
	return o
}

// SetQueryStartKey adds the startkey to the design doc view params
func (o *DesignDocViewParams) SetQueryStartKey(startkey *string) {
	o.QueryStartKey = startkey
}

// WithStartkeyDocid adds the startkeyDocid to the design doc view params
func (o *DesignDocViewParams) WithStartkeyDocid(startkeyDocid *string) *DesignDocViewParams {
	o.SetStartkeyDocid(startkeyDocid)
	return o
}

// SetStartkeyDocid adds the startkeyDocid to the design doc view params
func (o *DesignDocViewParams) SetStartkeyDocid(startkeyDocid *string) {
	o.StartkeyDocid = startkeyDocid
}

// WithUpdate adds the update to the design doc view params
func (o *DesignDocViewParams) WithUpdate(update *string) *DesignDocViewParams {
	o.SetUpdate(update)
	return o
}

// SetUpdate adds the update to the design doc view params
func (o *DesignDocViewParams) SetUpdate(update *string) {
	o.Update = update
}

// WithUpdateSeq adds the updateSeq to the design doc view params
func (o *DesignDocViewParams) WithUpdateSeq(updateSeq *bool) *DesignDocViewParams {
	o.SetUpdateSeq(updateSeq)
	return o
}

// SetUpdateSeq adds the updateSeq to the design doc view params
func (o *DesignDocViewParams) SetUpdateSeq(updateSeq *bool) {
	o.UpdateSeq = updateSeq
}

// WithView adds the view to the design doc view params
func (o *DesignDocViewParams) WithView(view string) *DesignDocViewParams {
	o.SetView(view)
	return o
}

// SetView adds the view to the design doc view params
func (o *DesignDocViewParams) SetView(view string) {
	o.View = view
}

// WriteToRequest writes these params to a swagger request
func (o *DesignDocViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AttEncodingInfo != nil {

		// query param att_encoding_info
		var qrAttEncodingInfo bool

		if o.AttEncodingInfo != nil {
			qrAttEncodingInfo = *o.AttEncodingInfo
		}
		qAttEncodingInfo := swag.FormatBool(qrAttEncodingInfo)
		if qAttEncodingInfo != "" {

			if err := r.SetQueryParam("att_encoding_info", qAttEncodingInfo); err != nil {
				return err
			}
		}
	}

	if o.Attachments != nil {

		// query param attachments
		var qrAttachments bool

		if o.Attachments != nil {
			qrAttachments = *o.Attachments
		}
		qAttachments := swag.FormatBool(qrAttachments)
		if qAttachments != "" {

			if err := r.SetQueryParam("attachments", qAttachments); err != nil {
				return err
			}
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

	// path param ddoc
	if err := r.SetPathParam("ddoc", o.Ddoc); err != nil {
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

	if o.Group != nil {

		// query param group
		var qrGroup bool

		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := swag.FormatBool(qrGroup)
		if qGroup != "" {

			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}
	}

	if o.GroupLevel != nil {

		// query param group_level
		var qrGroupLevel int64

		if o.GroupLevel != nil {
			qrGroupLevel = *o.GroupLevel
		}
		qGroupLevel := swag.FormatInt64(qrGroupLevel)
		if qGroupLevel != "" {

			if err := r.SetQueryParam("group_level", qGroupLevel); err != nil {
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

		// binding items for keys
		joinedKeys := o.bindParamKeys(reg)

		// query array param keys
		if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
			return err
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

	if o.Reduce != nil {

		// query param reduce
		var qrReduce bool

		if o.Reduce != nil {
			qrReduce = *o.Reduce
		}
		qReduce := swag.FormatBool(qrReduce)
		if qReduce != "" {

			if err := r.SetQueryParam("reduce", qReduce); err != nil {
				return err
			}
		}
	}

	if o.RevsInfo != nil {

		// query param revs_info
		var qrRevsInfo bool

		if o.RevsInfo != nil {
			qrRevsInfo = *o.RevsInfo
		}
		qRevsInfo := swag.FormatBool(qrRevsInfo)
		if qRevsInfo != "" {

			if err := r.SetQueryParam("revs_info", qRevsInfo); err != nil {
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

	if o.Sorted != nil {

		// query param sorted
		var qrSorted bool

		if o.Sorted != nil {
			qrSorted = *o.Sorted
		}
		qSorted := swag.FormatBool(qrSorted)
		if qSorted != "" {

			if err := r.SetQueryParam("sorted", qSorted); err != nil {
				return err
			}
		}
	}

	if o.Stable != nil {

		// query param stable
		var qrStable bool

		if o.Stable != nil {
			qrStable = *o.Stable
		}
		qStable := swag.FormatBool(qrStable)
		if qStable != "" {

			if err := r.SetQueryParam("stable", qStable); err != nil {
				return err
			}
		}
	}

	if o.Stale != nil {

		// query param stale
		var qrStale string

		if o.Stale != nil {
			qrStale = *o.Stale
		}
		qStale := qrStale
		if qStale != "" {

			if err := r.SetQueryParam("stale", qStale); err != nil {
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

	if o.Update != nil {

		// query param update
		var qrUpdate string

		if o.Update != nil {
			qrUpdate = *o.Update
		}
		qUpdate := qrUpdate
		if qUpdate != "" {

			if err := r.SetQueryParam("update", qUpdate); err != nil {
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

	// path param view
	if err := r.SetPathParam("view", o.View); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDesignDocView binds the parameter keys
func (o *DesignDocViewParams) bindParamKeys(formats strfmt.Registry) []string {
	keysIR := o.Keys

	var keysIC []string
	for _, keysIIR := range keysIR { // explode []string

		keysIIV := keysIIR // string as string
		keysIC = append(keysIC, keysIIV)
	}

	// items.CollectionFormat: "multi"
	keysIS := swag.JoinByFormat(keysIC, "multi")

	return keysIS
}
