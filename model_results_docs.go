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

// checks if the ResultsDocs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultsDocs{}

// ResultsDocs struct for ResultsDocs
type ResultsDocs struct {
	Ok map[string]interface{} `json:"ok,omitempty"`
	Error *ErrorBulkGetResponse `json:"error,omitempty"`
}

// NewResultsDocs instantiates a new ResultsDocs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultsDocs() *ResultsDocs {
	this := ResultsDocs{}
	return &this
}

// NewResultsDocsWithDefaults instantiates a new ResultsDocs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultsDocsWithDefaults() *ResultsDocs {
	this := ResultsDocs{}
	return &this
}

// GetOk returns the Ok field value if set, zero value otherwise.
func (o *ResultsDocs) GetOk() map[string]interface{} {
	if o == nil || IsNil(o.Ok) {
		var ret map[string]interface{}
		return ret
	}
	return o.Ok
}

// GetOkOk returns a tuple with the Ok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsDocs) GetOkOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Ok) {
		return map[string]interface{}{}, false
	}
	return o.Ok, true
}

// HasOk returns a boolean if a field has been set.
func (o *ResultsDocs) HasOk() bool {
	if o != nil && !IsNil(o.Ok) {
		return true
	}

	return false
}

// SetOk gets a reference to the given map[string]interface{} and assigns it to the Ok field.
func (o *ResultsDocs) SetOk(v map[string]interface{}) {
	o.Ok = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ResultsDocs) GetError() ErrorBulkGetResponse {
	if o == nil || IsNil(o.Error) {
		var ret ErrorBulkGetResponse
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsDocs) GetErrorOk() (*ErrorBulkGetResponse, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ResultsDocs) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ErrorBulkGetResponse and assigns it to the Error field.
func (o *ResultsDocs) SetError(v ErrorBulkGetResponse) {
	o.Error = &v
}

func (o ResultsDocs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultsDocs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ok) {
		toSerialize["ok"] = o.Ok
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableResultsDocs struct {
	value *ResultsDocs
	isSet bool
}

func (v NullableResultsDocs) Get() *ResultsDocs {
	return v.value
}

func (v *NullableResultsDocs) Set(val *ResultsDocs) {
	v.value = val
	v.isSet = true
}

func (v NullableResultsDocs) IsSet() bool {
	return v.isSet
}

func (v *NullableResultsDocs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultsDocs(val *ResultsDocs) *NullableResultsDocs {
	return &NullableResultsDocs{value: val, isSet: true}
}

func (v NullableResultsDocs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultsDocs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

