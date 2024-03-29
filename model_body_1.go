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

// checks if the Body1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Body1{}

// Body1 struct for Body1
type Body1 struct {
	// Type of analyzer
	Field *string `json:"field,omitempty"`
	// Analyzer token you want to test
	Text *string `json:"text,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Body1 Body1

// NewBody1 instantiates a new Body1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBody1() *Body1 {
	this := Body1{}
	return &this
}

// NewBody1WithDefaults instantiates a new Body1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBody1WithDefaults() *Body1 {
	this := Body1{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *Body1) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Body1) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *Body1) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *Body1) SetField(v string) {
	o.Field = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Body1) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Body1) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Body1) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *Body1) SetText(v string) {
	o.Text = &v
}

func (o Body1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Body1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Body1) UnmarshalJSON(bytes []byte) (err error) {
	varBody1 := _Body1{}

	if err = json.Unmarshal(bytes, &varBody1); err == nil {
		*o = Body1(varBody1)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "field")
		delete(additionalProperties, "text")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBody1 struct {
	value *Body1
	isSet bool
}

func (v NullableBody1) Get() *Body1 {
	return v.value
}

func (v *NullableBody1) Set(val *Body1) {
	v.value = val
	v.isSet = true
}

func (v NullableBody1) IsSet() bool {
	return v.isSet
}

func (v *NullableBody1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBody1(val *Body1) *NullableBody1 {
	return &NullableBody1{value: val, isSet: true}
}

func (v NullableBody1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBody1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


