/*
CouchDB API

*Note* This is not a definitive implementation of the CouchDB API, it's missing a lot of the endpoints for server/database managment and everything for attachments all COPY operations and probably a few other parts.  It also targets golang, as such the '#/definitions/Document' is intentionally left empty to generate a go `interface`, which you can then cast to a `map[string]interface{}`. 

API version: 4.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchdb_go

import (
	"encoding/json"
)

// checks if the ReplicationHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplicationHistory{}

// ReplicationHistory struct for ReplicationHistory
type ReplicationHistory struct {
	// Number of document write failures
	DocWriteFailures *int32 `json:"doc_write_failures,omitempty"`
	// Number of documents read
	DocsRead *int32 `json:"docs_read,omitempty"`
	// Number of documents written to target
	DocsWritten *int32 `json:"docs_written,omitempty"`
	// Last sequence number in changes stream
	EndLastSeq *int32 `json:"end_last_seq,omitempty"`
	// Date/Time replication operation completed in RFC 2822 format
	EndTime *string `json:"end_time,omitempty"`
	// Number of missing documents checked
	MissingChecked *int32 `json:"missing_checked,omitempty"`
	// Number of missing documents found
	MissingFound *int32 `json:"missing_found,omitempty"`
	// Last recorded sequence number
	RecordedSeq *int32 `json:"recorded_seq,omitempty"`
	// Session ID for this replication operation
	SessionId *string `json:"session_id,omitempty"`
	// First sequence number in changes stream
	StartLastSeq *int32 `json:"start_last_seq,omitempty"`
	// Date/Time replication operation started in RFC 2822 format
	StartTime *string `json:"start_time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReplicationHistory ReplicationHistory

// NewReplicationHistory instantiates a new ReplicationHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationHistory() *ReplicationHistory {
	this := ReplicationHistory{}
	return &this
}

// NewReplicationHistoryWithDefaults instantiates a new ReplicationHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationHistoryWithDefaults() *ReplicationHistory {
	this := ReplicationHistory{}
	return &this
}

// GetDocWriteFailures returns the DocWriteFailures field value if set, zero value otherwise.
func (o *ReplicationHistory) GetDocWriteFailures() int32 {
	if o == nil || IsNil(o.DocWriteFailures) {
		var ret int32
		return ret
	}
	return *o.DocWriteFailures
}

// GetDocWriteFailuresOk returns a tuple with the DocWriteFailures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationHistory) GetDocWriteFailuresOk() (*int32, bool) {
	if o == nil || IsNil(o.DocWriteFailures) {
		return nil, false
	}
	return o.DocWriteFailures, true
}

// HasDocWriteFailures returns a boolean if a field has been set.
func (o *ReplicationHistory) HasDocWriteFailures() bool {
	if o != nil && !IsNil(o.DocWriteFailures) {
		return true
	}

	return false
}

// SetDocWriteFailures gets a reference to the given int32 and assigns it to the DocWriteFailures field.
func (o *ReplicationHistory) SetDocWriteFailures(v int32) {
	o.DocWriteFailures = &v
}

// GetDocsRead returns the DocsRead field value if set, zero value otherwise.
func (o *ReplicationHistory) GetDocsRead() int32 {
	if o == nil || IsNil(o.DocsRead) {
		var ret int32
		return ret
	}
	return *o.DocsRead
}

// GetDocsReadOk returns a tuple with the DocsRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationHistory) GetDocsReadOk() (*int32, bool) {
	if o == nil || IsNil(o.DocsRead) {
		return nil, false
	}
	return o.DocsRead, true
}

// HasDocsRead returns a boolean if a field has been set.
func (o *ReplicationHistory) HasDocsRead() bool {
	if o != nil && !IsNil(o.DocsRead) {
		return true
	}

	return false
}

// SetDocsRead gets a reference to the given int32 and assigns it to the DocsRead field.
func (o *ReplicationHistory) SetDocsRead(v int32) {
	o.DocsRead = &v
}

// GetDocsWritten returns the DocsWritten field value if set, zero value otherwise.
func (o *ReplicationHistory) GetDocsWritten() int32 {
	if o == nil || IsNil(o.DocsWritten) {
		var ret int32
		return ret
	}
	return *o.DocsWritten
}

// GetDocsWrittenOk returns a tuple with the DocsWritten field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationHistory) GetDocsWrittenOk() (*int32, bool) {
	if o == nil || IsNil(o.DocsWritten) {
		return nil, false
	}
	return o.DocsWritten, true
}

// HasDocsWritten returns a boolean if a field has been set.
func (o *ReplicationHistory) HasDocsWritten() bool {
	if o != nil && !IsNil(o.DocsWritten) {
		return true
	}

	return false
}

// SetDocsWritten gets a reference to the given int32 and assigns it to the DocsWritten field.
func (o *ReplicationHistory) SetDocsWritten(v int32) {
	o.DocsWritten = &v
}

// GetEndLastSeq returns the EndLastSeq field value if set, zero value otherwise.
func (o *ReplicationHistory) GetEndLastSeq() int32 {
	if o == nil || IsNil(o.EndLastSeq) {
		var ret int32
		return ret
	}
	return *o.EndLastSeq
}

// GetEndLastSeqOk returns a tuple with the EndLastSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationHistory) GetEndLastSeqOk() (*int32, bool) {
	if o == nil || IsNil(o.EndLastSeq) {
		return nil, false
	}
	return o.EndLastSeq, true
}

// HasEndLastSeq returns a boolean if a field has been set.
func (o *ReplicationHistory) HasEndLastSeq() bool {
	if o != nil && !IsNil(o.EndLastSeq) {
		return true
	}

	return false
}

// SetEndLastSeq gets a reference to the given int32 and assigns it to the EndLastSeq field.
func (o *ReplicationHistory) SetEndLastSeq(v int32) {
	o.EndLastSeq = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ReplicationHistory) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationHistory) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ReplicationHistory) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *ReplicationHistory) SetEndTime(v string) {
	o.EndTime = &v
}

// GetMissingChecked returns the MissingChecked field value if set, zero value otherwise.
func (o *ReplicationHistory) GetMissingChecked() int32 {
	if o == nil || IsNil(o.MissingChecked) {
		var ret int32
		return ret
	}
	return *o.MissingChecked
}

// GetMissingCheckedOk returns a tuple with the MissingChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationHistory) GetMissingCheckedOk() (*int32, bool) {
	if o == nil || IsNil(o.MissingChecked) {
		return nil, false
	}
	return o.MissingChecked, true
}

// HasMissingChecked returns a boolean if a field has been set.
func (o *ReplicationHistory) HasMissingChecked() bool {
	if o != nil && !IsNil(o.MissingChecked) {
		return true
	}

	return false
}

// SetMissingChecked gets a reference to the given int32 and assigns it to the MissingChecked field.
func (o *ReplicationHistory) SetMissingChecked(v int32) {
	o.MissingChecked = &v
}

// GetMissingFound returns the MissingFound field value if set, zero value otherwise.
func (o *ReplicationHistory) GetMissingFound() int32 {
	if o == nil || IsNil(o.MissingFound) {
		var ret int32
		return ret
	}
	return *o.MissingFound
}

// GetMissingFoundOk returns a tuple with the MissingFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationHistory) GetMissingFoundOk() (*int32, bool) {
	if o == nil || IsNil(o.MissingFound) {
		return nil, false
	}
	return o.MissingFound, true
}

// HasMissingFound returns a boolean if a field has been set.
func (o *ReplicationHistory) HasMissingFound() bool {
	if o != nil && !IsNil(o.MissingFound) {
		return true
	}

	return false
}

// SetMissingFound gets a reference to the given int32 and assigns it to the MissingFound field.
func (o *ReplicationHistory) SetMissingFound(v int32) {
	o.MissingFound = &v
}

// GetRecordedSeq returns the RecordedSeq field value if set, zero value otherwise.
func (o *ReplicationHistory) GetRecordedSeq() int32 {
	if o == nil || IsNil(o.RecordedSeq) {
		var ret int32
		return ret
	}
	return *o.RecordedSeq
}

// GetRecordedSeqOk returns a tuple with the RecordedSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationHistory) GetRecordedSeqOk() (*int32, bool) {
	if o == nil || IsNil(o.RecordedSeq) {
		return nil, false
	}
	return o.RecordedSeq, true
}

// HasRecordedSeq returns a boolean if a field has been set.
func (o *ReplicationHistory) HasRecordedSeq() bool {
	if o != nil && !IsNil(o.RecordedSeq) {
		return true
	}

	return false
}

// SetRecordedSeq gets a reference to the given int32 and assigns it to the RecordedSeq field.
func (o *ReplicationHistory) SetRecordedSeq(v int32) {
	o.RecordedSeq = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *ReplicationHistory) GetSessionId() string {
	if o == nil || IsNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationHistory) GetSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *ReplicationHistory) HasSessionId() bool {
	if o != nil && !IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *ReplicationHistory) SetSessionId(v string) {
	o.SessionId = &v
}

// GetStartLastSeq returns the StartLastSeq field value if set, zero value otherwise.
func (o *ReplicationHistory) GetStartLastSeq() int32 {
	if o == nil || IsNil(o.StartLastSeq) {
		var ret int32
		return ret
	}
	return *o.StartLastSeq
}

// GetStartLastSeqOk returns a tuple with the StartLastSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationHistory) GetStartLastSeqOk() (*int32, bool) {
	if o == nil || IsNil(o.StartLastSeq) {
		return nil, false
	}
	return o.StartLastSeq, true
}

// HasStartLastSeq returns a boolean if a field has been set.
func (o *ReplicationHistory) HasStartLastSeq() bool {
	if o != nil && !IsNil(o.StartLastSeq) {
		return true
	}

	return false
}

// SetStartLastSeq gets a reference to the given int32 and assigns it to the StartLastSeq field.
func (o *ReplicationHistory) SetStartLastSeq(v int32) {
	o.StartLastSeq = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ReplicationHistory) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationHistory) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ReplicationHistory) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *ReplicationHistory) SetStartTime(v string) {
	o.StartTime = &v
}

func (o ReplicationHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplicationHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocWriteFailures) {
		toSerialize["doc_write_failures"] = o.DocWriteFailures
	}
	if !IsNil(o.DocsRead) {
		toSerialize["docs_read"] = o.DocsRead
	}
	if !IsNil(o.DocsWritten) {
		toSerialize["docs_written"] = o.DocsWritten
	}
	if !IsNil(o.EndLastSeq) {
		toSerialize["end_last_seq"] = o.EndLastSeq
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.MissingChecked) {
		toSerialize["missing_checked"] = o.MissingChecked
	}
	if !IsNil(o.MissingFound) {
		toSerialize["missing_found"] = o.MissingFound
	}
	if !IsNil(o.RecordedSeq) {
		toSerialize["recorded_seq"] = o.RecordedSeq
	}
	if !IsNil(o.SessionId) {
		toSerialize["session_id"] = o.SessionId
	}
	if !IsNil(o.StartLastSeq) {
		toSerialize["start_last_seq"] = o.StartLastSeq
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReplicationHistory) UnmarshalJSON(bytes []byte) (err error) {
	varReplicationHistory := _ReplicationHistory{}

	if err = json.Unmarshal(bytes, &varReplicationHistory); err == nil {
		*o = ReplicationHistory(varReplicationHistory)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "doc_write_failures")
		delete(additionalProperties, "docs_read")
		delete(additionalProperties, "docs_written")
		delete(additionalProperties, "end_last_seq")
		delete(additionalProperties, "end_time")
		delete(additionalProperties, "missing_checked")
		delete(additionalProperties, "missing_found")
		delete(additionalProperties, "recorded_seq")
		delete(additionalProperties, "session_id")
		delete(additionalProperties, "start_last_seq")
		delete(additionalProperties, "start_time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReplicationHistory struct {
	value *ReplicationHistory
	isSet bool
}

func (v NullableReplicationHistory) Get() *ReplicationHistory {
	return v.value
}

func (v *NullableReplicationHistory) Set(val *ReplicationHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicationHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationHistory(val *ReplicationHistory) *NullableReplicationHistory {
	return &NullableReplicationHistory{value: val, isSet: true}
}

func (v NullableReplicationHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


