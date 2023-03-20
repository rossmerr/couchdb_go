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

// checks if the Server type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Server{}

// Server struct for Server
type Server struct {
	Couchdb *string `json:"couchdb,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Vendor *ServerVendor `json:"vendor,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Server Server

// NewServer instantiates a new Server object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServer() *Server {
	this := Server{}
	return &this
}

// NewServerWithDefaults instantiates a new Server object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerWithDefaults() *Server {
	this := Server{}
	return &this
}

// GetCouchdb returns the Couchdb field value if set, zero value otherwise.
func (o *Server) GetCouchdb() string {
	if o == nil || IsNil(o.Couchdb) {
		var ret string
		return ret
	}
	return *o.Couchdb
}

// GetCouchdbOk returns a tuple with the Couchdb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetCouchdbOk() (*string, bool) {
	if o == nil || IsNil(o.Couchdb) {
		return nil, false
	}
	return o.Couchdb, true
}

// HasCouchdb returns a boolean if a field has been set.
func (o *Server) HasCouchdb() bool {
	if o != nil && !IsNil(o.Couchdb) {
		return true
	}

	return false
}

// SetCouchdb gets a reference to the given string and assigns it to the Couchdb field.
func (o *Server) SetCouchdb(v string) {
	o.Couchdb = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Server) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Server) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Server) SetUuid(v string) {
	o.Uuid = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *Server) GetVendor() ServerVendor {
	if o == nil || IsNil(o.Vendor) {
		var ret ServerVendor
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetVendorOk() (*ServerVendor, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *Server) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given ServerVendor and assigns it to the Vendor field.
func (o *Server) SetVendor(v ServerVendor) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Server) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Server) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Server) SetVersion(v string) {
	o.Version = &v
}

func (o Server) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Server) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Couchdb) {
		toSerialize["couchdb"] = o.Couchdb
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Server) UnmarshalJSON(bytes []byte) (err error) {
	varServer := _Server{}

	if err = json.Unmarshal(bytes, &varServer); err == nil {
		*o = Server(varServer)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "couchdb")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServer struct {
	value *Server
	isSet bool
}

func (v NullableServer) Get() *Server {
	return v.value
}

func (v *NullableServer) Set(val *Server) {
	v.value = val
	v.isSet = true
}

func (v NullableServer) IsSet() bool {
	return v.isSet
}

func (v *NullableServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServer(val *Server) *NullableServer {
	return &NullableServer{value: val, isSet: true}
}

func (v NullableServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


