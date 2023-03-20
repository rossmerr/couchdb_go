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

// checks if the DatabaseSizes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseSizes{}

// DatabaseSizes struct for DatabaseSizes
type DatabaseSizes struct {
	// The size of live data inside the database, in bytes.
	Active *int32 `json:"active,omitempty"`
	// The uncompressed size of database contents in bytes. sizes.file (number) – The size of the database file on disk in bytes. Views indexes are not included in the calculation.
	External *int32 `json:"external,omitempty"`
	// An opaque string that describes the state of the database. Do not rely on this string for counting the number of updates.
	File *int32 `json:"file,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DatabaseSizes DatabaseSizes

// NewDatabaseSizes instantiates a new DatabaseSizes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseSizes() *DatabaseSizes {
	this := DatabaseSizes{}
	return &this
}

// NewDatabaseSizesWithDefaults instantiates a new DatabaseSizes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseSizesWithDefaults() *DatabaseSizes {
	this := DatabaseSizes{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *DatabaseSizes) GetActive() int32 {
	if o == nil || IsNil(o.Active) {
		var ret int32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseSizes) GetActiveOk() (*int32, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *DatabaseSizes) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given int32 and assigns it to the Active field.
func (o *DatabaseSizes) SetActive(v int32) {
	o.Active = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *DatabaseSizes) GetExternal() int32 {
	if o == nil || IsNil(o.External) {
		var ret int32
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseSizes) GetExternalOk() (*int32, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *DatabaseSizes) HasExternal() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given int32 and assigns it to the External field.
func (o *DatabaseSizes) SetExternal(v int32) {
	o.External = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *DatabaseSizes) GetFile() int32 {
	if o == nil || IsNil(o.File) {
		var ret int32
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseSizes) GetFileOk() (*int32, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *DatabaseSizes) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given int32 and assigns it to the File field.
func (o *DatabaseSizes) SetFile(v int32) {
	o.File = &v
}

func (o DatabaseSizes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseSizes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DatabaseSizes) UnmarshalJSON(bytes []byte) (err error) {
	varDatabaseSizes := _DatabaseSizes{}

	if err = json.Unmarshal(bytes, &varDatabaseSizes); err == nil {
		*o = DatabaseSizes(varDatabaseSizes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "external")
		delete(additionalProperties, "file")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDatabaseSizes struct {
	value *DatabaseSizes
	isSet bool
}

func (v NullableDatabaseSizes) Get() *DatabaseSizes {
	return v.value
}

func (v *NullableDatabaseSizes) Set(val *DatabaseSizes) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseSizes) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseSizes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseSizes(val *DatabaseSizes) *NullableDatabaseSizes {
	return &NullableDatabaseSizes{value: val, isSet: true}
}

func (v NullableDatabaseSizes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseSizes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


