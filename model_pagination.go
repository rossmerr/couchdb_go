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

// checks if the Pagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pagination{}

// Pagination struct for Pagination
type Pagination struct {
	// Offset where the document list started.
	Offset *int32 `json:"offset,omitempty"`
	// Number of documents in the database/view.
	TotalRows *int32 `json:"total_rows,omitempty"`
	// Array of view row objects. By default the information returned contains only the document ID and revision.
	Rows []Row `json:"rows,omitempty"`
	// Current update sequence for the database.
	UpdateSeq map[string]interface{} `json:"update_seq,omitempty"`
}

// NewPagination instantiates a new Pagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagination() *Pagination {
	this := Pagination{}
	return &this
}

// NewPaginationWithDefaults instantiates a new Pagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationWithDefaults() *Pagination {
	this := Pagination{}
	return &this
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *Pagination) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *Pagination) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *Pagination) SetOffset(v int32) {
	o.Offset = &v
}

// GetTotalRows returns the TotalRows field value if set, zero value otherwise.
func (o *Pagination) GetTotalRows() int32 {
	if o == nil || IsNil(o.TotalRows) {
		var ret int32
		return ret
	}
	return *o.TotalRows
}

// GetTotalRowsOk returns a tuple with the TotalRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetTotalRowsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRows) {
		return nil, false
	}
	return o.TotalRows, true
}

// HasTotalRows returns a boolean if a field has been set.
func (o *Pagination) HasTotalRows() bool {
	if o != nil && !IsNil(o.TotalRows) {
		return true
	}

	return false
}

// SetTotalRows gets a reference to the given int32 and assigns it to the TotalRows field.
func (o *Pagination) SetTotalRows(v int32) {
	o.TotalRows = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *Pagination) GetRows() []Row {
	if o == nil || IsNil(o.Rows) {
		var ret []Row
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetRowsOk() ([]Row, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *Pagination) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []Row and assigns it to the Rows field.
func (o *Pagination) SetRows(v []Row) {
	o.Rows = v
}

// GetUpdateSeq returns the UpdateSeq field value if set, zero value otherwise.
func (o *Pagination) GetUpdateSeq() map[string]interface{} {
	if o == nil || IsNil(o.UpdateSeq) {
		var ret map[string]interface{}
		return ret
	}
	return o.UpdateSeq
}

// GetUpdateSeqOk returns a tuple with the UpdateSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetUpdateSeqOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UpdateSeq) {
		return map[string]interface{}{}, false
	}
	return o.UpdateSeq, true
}

// HasUpdateSeq returns a boolean if a field has been set.
func (o *Pagination) HasUpdateSeq() bool {
	if o != nil && !IsNil(o.UpdateSeq) {
		return true
	}

	return false
}

// SetUpdateSeq gets a reference to the given map[string]interface{} and assigns it to the UpdateSeq field.
func (o *Pagination) SetUpdateSeq(v map[string]interface{}) {
	o.UpdateSeq = v
}

func (o Pagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.TotalRows) {
		toSerialize["total_rows"] = o.TotalRows
	}
	if !IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !IsNil(o.UpdateSeq) {
		toSerialize["update_seq"] = o.UpdateSeq
	}
	return toSerialize, nil
}

type NullablePagination struct {
	value *Pagination
	isSet bool
}

func (v NullablePagination) Get() *Pagination {
	return v.value
}

func (v *NullablePagination) Set(val *Pagination) {
	v.value = val
	v.isSet = true
}

func (v NullablePagination) IsSet() bool {
	return v.isSet
}

func (v *NullablePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagination(val *Pagination) *NullablePagination {
	return &NullablePagination{value: val, isSet: true}
}

func (v NullablePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


