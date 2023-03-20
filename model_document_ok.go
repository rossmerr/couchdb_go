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

// checks if the DocumentOK type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentOK{}

// DocumentOK struct for DocumentOK
type DocumentOK struct {
	// Operation status
	Ok *bool `json:"ok,omitempty"`
	Id *string `json:"id,omitempty"`
	Rev *string `json:"rev,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DocumentOK DocumentOK

// NewDocumentOK instantiates a new DocumentOK object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentOK() *DocumentOK {
	this := DocumentOK{}
	return &this
}

// NewDocumentOKWithDefaults instantiates a new DocumentOK object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentOKWithDefaults() *DocumentOK {
	this := DocumentOK{}
	return &this
}

// GetOk returns the Ok field value if set, zero value otherwise.
func (o *DocumentOK) GetOk() bool {
	if o == nil || IsNil(o.Ok) {
		var ret bool
		return ret
	}
	return *o.Ok
}

// GetOkOk returns a tuple with the Ok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentOK) GetOkOk() (*bool, bool) {
	if o == nil || IsNil(o.Ok) {
		return nil, false
	}
	return o.Ok, true
}

// HasOk returns a boolean if a field has been set.
func (o *DocumentOK) HasOk() bool {
	if o != nil && !IsNil(o.Ok) {
		return true
	}

	return false
}

// SetOk gets a reference to the given bool and assigns it to the Ok field.
func (o *DocumentOK) SetOk(v bool) {
	o.Ok = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentOK) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentOK) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentOK) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentOK) SetId(v string) {
	o.Id = &v
}

// GetRev returns the Rev field value if set, zero value otherwise.
func (o *DocumentOK) GetRev() string {
	if o == nil || IsNil(o.Rev) {
		var ret string
		return ret
	}
	return *o.Rev
}

// GetRevOk returns a tuple with the Rev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentOK) GetRevOk() (*string, bool) {
	if o == nil || IsNil(o.Rev) {
		return nil, false
	}
	return o.Rev, true
}

// HasRev returns a boolean if a field has been set.
func (o *DocumentOK) HasRev() bool {
	if o != nil && !IsNil(o.Rev) {
		return true
	}

	return false
}

// SetRev gets a reference to the given string and assigns it to the Rev field.
func (o *DocumentOK) SetRev(v string) {
	o.Rev = &v
}

func (o DocumentOK) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentOK) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ok) {
		toSerialize["ok"] = o.Ok
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Rev) {
		toSerialize["rev"] = o.Rev
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DocumentOK) UnmarshalJSON(bytes []byte) (err error) {
	varDocumentOK := _DocumentOK{}

	if err = json.Unmarshal(bytes, &varDocumentOK); err == nil {
		*o = DocumentOK(varDocumentOK)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ok")
		delete(additionalProperties, "id")
		delete(additionalProperties, "rev")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDocumentOK struct {
	value *DocumentOK
	isSet bool
}

func (v NullableDocumentOK) Get() *DocumentOK {
	return v.value
}

func (v *NullableDocumentOK) Set(val *DocumentOK) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentOK) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentOK) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentOK(val *DocumentOK) *NullableDocumentOK {
	return &NullableDocumentOK{value: val, isSet: true}
}

func (v NullableDocumentOK) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentOK) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


