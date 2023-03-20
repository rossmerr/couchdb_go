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

// checks if the ServerVendor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerVendor{}

// ServerVendor struct for ServerVendor
type ServerVendor struct {
	Name *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewServerVendor instantiates a new ServerVendor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerVendor() *ServerVendor {
	this := ServerVendor{}
	return &this
}

// NewServerVendorWithDefaults instantiates a new ServerVendor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerVendorWithDefaults() *ServerVendor {
	this := ServerVendor{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServerVendor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerVendor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServerVendor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServerVendor) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ServerVendor) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerVendor) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ServerVendor) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ServerVendor) SetVersion(v string) {
	o.Version = &v
}

func (o ServerVendor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerVendor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableServerVendor struct {
	value *ServerVendor
	isSet bool
}

func (v NullableServerVendor) Get() *ServerVendor {
	return v.value
}

func (v *NullableServerVendor) Set(val *ServerVendor) {
	v.value = val
	v.isSet = true
}

func (v NullableServerVendor) IsSet() bool {
	return v.isSet
}

func (v *NullableServerVendor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerVendor(val *ServerVendor) *NullableServerVendor {
	return &NullableServerVendor{value: val, isSet: true}
}

func (v NullableServerVendor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerVendor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

