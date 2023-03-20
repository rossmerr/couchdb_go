# Replicate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancel** | Pointer to **bool** | Cancels the replication | [optional] 
**Continuous** | Pointer to **bool** | Configure the replication to be continuous | [optional] 
**CreateTarget** | Pointer to **bool** | Creates the target database. Required administrator’s privileges on target server. | [optional] 
**CreateTargetParams** | Pointer to [**ReplicateCreateTargetParams**](ReplicateCreateTargetParams.md) |  | [optional] 
**DocIds** | Pointer to **[]string** | Array of document IDs to be synchronized | [optional] 
**Filter** | Pointer to **string** | The name of a filter function. | [optional] 
**SourceProxy** | Pointer to **string** | Address of a proxy server through which replication from the source should occur (protocol can be “http” or “socks5”) | [optional] 
**TargetProxy** | Pointer to **string** | Address of a proxy server through which replication to the target should occur (protocol can be “http” or “socks5”) | [optional] 
**Source** | Pointer to [**Request**](Request.md) |  | [optional] 
**Target** | Pointer to [**Request**](Request.md) |  | [optional] 

## Methods

### NewReplicate

`func NewReplicate() *Replicate`

NewReplicate instantiates a new Replicate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicateWithDefaults

`func NewReplicateWithDefaults() *Replicate`

NewReplicateWithDefaults instantiates a new Replicate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancel

`func (o *Replicate) GetCancel() bool`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *Replicate) GetCancelOk() (*bool, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *Replicate) SetCancel(v bool)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *Replicate) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### GetContinuous

`func (o *Replicate) GetContinuous() bool`

GetContinuous returns the Continuous field if non-nil, zero value otherwise.

### GetContinuousOk

`func (o *Replicate) GetContinuousOk() (*bool, bool)`

GetContinuousOk returns a tuple with the Continuous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuous

`func (o *Replicate) SetContinuous(v bool)`

SetContinuous sets Continuous field to given value.

### HasContinuous

`func (o *Replicate) HasContinuous() bool`

HasContinuous returns a boolean if a field has been set.

### GetCreateTarget

`func (o *Replicate) GetCreateTarget() bool`

GetCreateTarget returns the CreateTarget field if non-nil, zero value otherwise.

### GetCreateTargetOk

`func (o *Replicate) GetCreateTargetOk() (*bool, bool)`

GetCreateTargetOk returns a tuple with the CreateTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTarget

`func (o *Replicate) SetCreateTarget(v bool)`

SetCreateTarget sets CreateTarget field to given value.

### HasCreateTarget

`func (o *Replicate) HasCreateTarget() bool`

HasCreateTarget returns a boolean if a field has been set.

### GetCreateTargetParams

`func (o *Replicate) GetCreateTargetParams() ReplicateCreateTargetParams`

GetCreateTargetParams returns the CreateTargetParams field if non-nil, zero value otherwise.

### GetCreateTargetParamsOk

`func (o *Replicate) GetCreateTargetParamsOk() (*ReplicateCreateTargetParams, bool)`

GetCreateTargetParamsOk returns a tuple with the CreateTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTargetParams

`func (o *Replicate) SetCreateTargetParams(v ReplicateCreateTargetParams)`

SetCreateTargetParams sets CreateTargetParams field to given value.

### HasCreateTargetParams

`func (o *Replicate) HasCreateTargetParams() bool`

HasCreateTargetParams returns a boolean if a field has been set.

### GetDocIds

`func (o *Replicate) GetDocIds() []string`

GetDocIds returns the DocIds field if non-nil, zero value otherwise.

### GetDocIdsOk

`func (o *Replicate) GetDocIdsOk() (*[]string, bool)`

GetDocIdsOk returns a tuple with the DocIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocIds

`func (o *Replicate) SetDocIds(v []string)`

SetDocIds sets DocIds field to given value.

### HasDocIds

`func (o *Replicate) HasDocIds() bool`

HasDocIds returns a boolean if a field has been set.

### GetFilter

`func (o *Replicate) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Replicate) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Replicate) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Replicate) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSourceProxy

`func (o *Replicate) GetSourceProxy() string`

GetSourceProxy returns the SourceProxy field if non-nil, zero value otherwise.

### GetSourceProxyOk

`func (o *Replicate) GetSourceProxyOk() (*string, bool)`

GetSourceProxyOk returns a tuple with the SourceProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProxy

`func (o *Replicate) SetSourceProxy(v string)`

SetSourceProxy sets SourceProxy field to given value.

### HasSourceProxy

`func (o *Replicate) HasSourceProxy() bool`

HasSourceProxy returns a boolean if a field has been set.

### GetTargetProxy

`func (o *Replicate) GetTargetProxy() string`

GetTargetProxy returns the TargetProxy field if non-nil, zero value otherwise.

### GetTargetProxyOk

`func (o *Replicate) GetTargetProxyOk() (*string, bool)`

GetTargetProxyOk returns a tuple with the TargetProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProxy

`func (o *Replicate) SetTargetProxy(v string)`

SetTargetProxy sets TargetProxy field to given value.

### HasTargetProxy

`func (o *Replicate) HasTargetProxy() bool`

HasTargetProxy returns a boolean if a field has been set.

### GetSource

`func (o *Replicate) GetSource() Request`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Replicate) GetSourceOk() (*Request, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Replicate) SetSource(v Request)`

SetSource sets Source field to given value.

### HasSource

`func (o *Replicate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *Replicate) GetTarget() Request`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Replicate) GetTargetOk() (*Request, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Replicate) SetTarget(v Request)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Replicate) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


