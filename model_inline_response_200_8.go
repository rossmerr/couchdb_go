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

// checks if the InlineResponse2008 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineResponse2008{}

// InlineResponse2008 struct for InlineResponse2008
type InlineResponse2008 struct {
	Name *string `json:"name,omitempty"`
	SearchIndex *SearchIndex `json:"search_index,omitempty"`
}

// NewInlineResponse2008 instantiates a new InlineResponse2008 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2008() *InlineResponse2008 {
	this := InlineResponse2008{}
	return &this
}

// NewInlineResponse2008WithDefaults instantiates a new InlineResponse2008 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2008WithDefaults() *InlineResponse2008 {
	this := InlineResponse2008{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse2008) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse2008) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse2008) SetName(v string) {
	o.Name = &v
}

// GetSearchIndex returns the SearchIndex field value if set, zero value otherwise.
func (o *InlineResponse2008) GetSearchIndex() SearchIndex {
	if o == nil || IsNil(o.SearchIndex) {
		var ret SearchIndex
		return ret
	}
	return *o.SearchIndex
}

// GetSearchIndexOk returns a tuple with the SearchIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008) GetSearchIndexOk() (*SearchIndex, bool) {
	if o == nil || IsNil(o.SearchIndex) {
		return nil, false
	}
	return o.SearchIndex, true
}

// HasSearchIndex returns a boolean if a field has been set.
func (o *InlineResponse2008) HasSearchIndex() bool {
	if o != nil && !IsNil(o.SearchIndex) {
		return true
	}

	return false
}

// SetSearchIndex gets a reference to the given SearchIndex and assigns it to the SearchIndex field.
func (o *InlineResponse2008) SetSearchIndex(v SearchIndex) {
	o.SearchIndex = &v
}

func (o InlineResponse2008) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineResponse2008) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SearchIndex) {
		toSerialize["search_index"] = o.SearchIndex
	}
	return toSerialize, nil
}

type NullableInlineResponse2008 struct {
	value *InlineResponse2008
	isSet bool
}

func (v NullableInlineResponse2008) Get() *InlineResponse2008 {
	return v.value
}

func (v *NullableInlineResponse2008) Set(val *InlineResponse2008) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2008) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2008) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2008(val *InlineResponse2008) *NullableInlineResponse2008 {
	return &NullableInlineResponse2008{value: val, isSet: true}
}

func (v NullableInlineResponse2008) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2008) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

