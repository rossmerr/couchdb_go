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

// checks if the IndexDefinitions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexDefinitions{}

// IndexDefinitions struct for IndexDefinitions
type IndexDefinitions struct {
	Ddoc *string `json:"ddoc,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Def *IndexDefinitionsDef `json:"def,omitempty"`
}

// NewIndexDefinitions instantiates a new IndexDefinitions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexDefinitions() *IndexDefinitions {
	this := IndexDefinitions{}
	return &this
}

// NewIndexDefinitionsWithDefaults instantiates a new IndexDefinitions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexDefinitionsWithDefaults() *IndexDefinitions {
	this := IndexDefinitions{}
	return &this
}

// GetDdoc returns the Ddoc field value if set, zero value otherwise.
func (o *IndexDefinitions) GetDdoc() string {
	if o == nil || IsNil(o.Ddoc) {
		var ret string
		return ret
	}
	return *o.Ddoc
}

// GetDdocOk returns a tuple with the Ddoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexDefinitions) GetDdocOk() (*string, bool) {
	if o == nil || IsNil(o.Ddoc) {
		return nil, false
	}
	return o.Ddoc, true
}

// HasDdoc returns a boolean if a field has been set.
func (o *IndexDefinitions) HasDdoc() bool {
	if o != nil && !IsNil(o.Ddoc) {
		return true
	}

	return false
}

// SetDdoc gets a reference to the given string and assigns it to the Ddoc field.
func (o *IndexDefinitions) SetDdoc(v string) {
	o.Ddoc = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IndexDefinitions) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexDefinitions) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IndexDefinitions) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IndexDefinitions) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IndexDefinitions) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexDefinitions) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IndexDefinitions) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IndexDefinitions) SetType(v string) {
	o.Type = &v
}

// GetDef returns the Def field value if set, zero value otherwise.
func (o *IndexDefinitions) GetDef() IndexDefinitionsDef {
	if o == nil || IsNil(o.Def) {
		var ret IndexDefinitionsDef
		return ret
	}
	return *o.Def
}

// GetDefOk returns a tuple with the Def field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexDefinitions) GetDefOk() (*IndexDefinitionsDef, bool) {
	if o == nil || IsNil(o.Def) {
		return nil, false
	}
	return o.Def, true
}

// HasDef returns a boolean if a field has been set.
func (o *IndexDefinitions) HasDef() bool {
	if o != nil && !IsNil(o.Def) {
		return true
	}

	return false
}

// SetDef gets a reference to the given IndexDefinitionsDef and assigns it to the Def field.
func (o *IndexDefinitions) SetDef(v IndexDefinitionsDef) {
	o.Def = &v
}

func (o IndexDefinitions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexDefinitions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ddoc) {
		toSerialize["ddoc"] = o.Ddoc
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Def) {
		toSerialize["def"] = o.Def
	}
	return toSerialize, nil
}

type NullableIndexDefinitions struct {
	value *IndexDefinitions
	isSet bool
}

func (v NullableIndexDefinitions) Get() *IndexDefinitions {
	return v.value
}

func (v *NullableIndexDefinitions) Set(val *IndexDefinitions) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexDefinitions) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexDefinitions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexDefinitions(val *IndexDefinitions) *NullableIndexDefinitions {
	return &NullableIndexDefinitions{value: val, isSet: true}
}

func (v NullableIndexDefinitions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexDefinitions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

