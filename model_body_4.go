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

// checks if the Body4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Body4{}

// Body4 struct for Body4
type Body4 struct {
	Docs []Keys `json:"docs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Body4 Body4

// NewBody4 instantiates a new Body4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBody4() *Body4 {
	this := Body4{}
	return &this
}

// NewBody4WithDefaults instantiates a new Body4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBody4WithDefaults() *Body4 {
	this := Body4{}
	return &this
}

// GetDocs returns the Docs field value if set, zero value otherwise.
func (o *Body4) GetDocs() []Keys {
	if o == nil || IsNil(o.Docs) {
		var ret []Keys
		return ret
	}
	return o.Docs
}

// GetDocsOk returns a tuple with the Docs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Body4) GetDocsOk() ([]Keys, bool) {
	if o == nil || IsNil(o.Docs) {
		return nil, false
	}
	return o.Docs, true
}

// HasDocs returns a boolean if a field has been set.
func (o *Body4) HasDocs() bool {
	if o != nil && !IsNil(o.Docs) {
		return true
	}

	return false
}

// SetDocs gets a reference to the given []Keys and assigns it to the Docs field.
func (o *Body4) SetDocs(v []Keys) {
	o.Docs = v
}

func (o Body4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Body4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Docs) {
		toSerialize["docs"] = o.Docs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Body4) UnmarshalJSON(bytes []byte) (err error) {
	varBody4 := _Body4{}

	if err = json.Unmarshal(bytes, &varBody4); err == nil {
		*o = Body4(varBody4)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "docs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBody4 struct {
	value *Body4
	isSet bool
}

func (v NullableBody4) Get() *Body4 {
	return v.value
}

func (v *NullableBody4) Set(val *Body4) {
	v.value = val
	v.isSet = true
}

func (v NullableBody4) IsSet() bool {
	return v.isSet
}

func (v *NullableBody4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBody4(val *Body4) *NullableBody4 {
	return &NullableBody4{value: val, isSet: true}
}

func (v NullableBody4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBody4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


