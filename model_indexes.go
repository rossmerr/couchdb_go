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

// checks if the Indexes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Indexes{}

// Indexes struct for Indexes
type Indexes struct {
	TotalRows *float32 `json:"total_rows,omitempty"`
	Indexes []IndexDefinitions `json:"indexes,omitempty"`
}

// NewIndexes instantiates a new Indexes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexes() *Indexes {
	this := Indexes{}
	return &this
}

// NewIndexesWithDefaults instantiates a new Indexes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexesWithDefaults() *Indexes {
	this := Indexes{}
	return &this
}

// GetTotalRows returns the TotalRows field value if set, zero value otherwise.
func (o *Indexes) GetTotalRows() float32 {
	if o == nil || IsNil(o.TotalRows) {
		var ret float32
		return ret
	}
	return *o.TotalRows
}

// GetTotalRowsOk returns a tuple with the TotalRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Indexes) GetTotalRowsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalRows) {
		return nil, false
	}
	return o.TotalRows, true
}

// HasTotalRows returns a boolean if a field has been set.
func (o *Indexes) HasTotalRows() bool {
	if o != nil && !IsNil(o.TotalRows) {
		return true
	}

	return false
}

// SetTotalRows gets a reference to the given float32 and assigns it to the TotalRows field.
func (o *Indexes) SetTotalRows(v float32) {
	o.TotalRows = &v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *Indexes) GetIndexes() []IndexDefinitions {
	if o == nil || IsNil(o.Indexes) {
		var ret []IndexDefinitions
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Indexes) GetIndexesOk() ([]IndexDefinitions, bool) {
	if o == nil || IsNil(o.Indexes) {
		return nil, false
	}
	return o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *Indexes) HasIndexes() bool {
	if o != nil && !IsNil(o.Indexes) {
		return true
	}

	return false
}

// SetIndexes gets a reference to the given []IndexDefinitions and assigns it to the Indexes field.
func (o *Indexes) SetIndexes(v []IndexDefinitions) {
	o.Indexes = v
}

func (o Indexes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Indexes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalRows) {
		toSerialize["total_rows"] = o.TotalRows
	}
	if !IsNil(o.Indexes) {
		toSerialize["indexes"] = o.Indexes
	}
	return toSerialize, nil
}

type NullableIndexes struct {
	value *Indexes
	isSet bool
}

func (v NullableIndexes) Get() *Indexes {
	return v.value
}

func (v *NullableIndexes) Set(val *Indexes) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexes) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexes(val *Indexes) *NullableIndexes {
	return &NullableIndexes{value: val, isSet: true}
}

func (v NullableIndexes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

