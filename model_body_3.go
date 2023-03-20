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

// checks if the Body3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Body3{}

// Body3 List of documents objects
type Body3 struct {
	Docs []map[string]interface{} `json:"docs,omitempty"`
}

// NewBody3 instantiates a new Body3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBody3() *Body3 {
	this := Body3{}
	return &this
}

// NewBody3WithDefaults instantiates a new Body3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBody3WithDefaults() *Body3 {
	this := Body3{}
	return &this
}

// GetDocs returns the Docs field value if set, zero value otherwise.
func (o *Body3) GetDocs() []map[string]interface{} {
	if o == nil || IsNil(o.Docs) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Docs
}

// GetDocsOk returns a tuple with the Docs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Body3) GetDocsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Docs) {
		return nil, false
	}
	return o.Docs, true
}

// HasDocs returns a boolean if a field has been set.
func (o *Body3) HasDocs() bool {
	if o != nil && !IsNil(o.Docs) {
		return true
	}

	return false
}

// SetDocs gets a reference to the given []map[string]interface{} and assigns it to the Docs field.
func (o *Body3) SetDocs(v []map[string]interface{}) {
	o.Docs = v
}

func (o Body3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Body3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Docs) {
		toSerialize["docs"] = o.Docs
	}
	return toSerialize, nil
}

type NullableBody3 struct {
	value *Body3
	isSet bool
}

func (v NullableBody3) Get() *Body3 {
	return v.value
}

func (v *NullableBody3) Set(val *Body3) {
	v.value = val
	v.isSet = true
}

func (v NullableBody3) IsSet() bool {
	return v.isSet
}

func (v *NullableBody3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBody3(val *Body3) *NullableBody3 {
	return &NullableBody3{value: val, isSet: true}
}

func (v NullableBody3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBody3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


