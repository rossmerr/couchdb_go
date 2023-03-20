# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Couchdb** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to [**ServerVendor**](ServerVendor.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewServer

`func NewServer() *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCouchdb

`func (o *Server) GetCouchdb() string`

GetCouchdb returns the Couchdb field if non-nil, zero value otherwise.

### GetCouchdbOk

`func (o *Server) GetCouchdbOk() (*string, bool)`

GetCouchdbOk returns a tuple with the Couchdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchdb

`func (o *Server) SetCouchdb(v string)`

SetCouchdb sets Couchdb field to given value.

### HasCouchdb

`func (o *Server) HasCouchdb() bool`

HasCouchdb returns a boolean if a field has been set.

### GetUuid

`func (o *Server) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Server) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Server) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Server) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendor

`func (o *Server) GetVendor() ServerVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Server) GetVendorOk() (*ServerVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Server) SetVendor(v ServerVendor)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Server) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *Server) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Server) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Server) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Server) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


