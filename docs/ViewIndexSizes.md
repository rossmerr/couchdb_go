# ViewIndexSizes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** | The size of live data inside the view, in bytes | [optional] 
**Disk** | Pointer to **int32** | Size in bytes of the view as stored on disk | [optional] 
**External** | Pointer to **int32** | The uncompressed size of view contents in bytes | [optional] 

## Methods

### NewViewIndexSizes

`func NewViewIndexSizes() *ViewIndexSizes`

NewViewIndexSizes instantiates a new ViewIndexSizes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewIndexSizesWithDefaults

`func NewViewIndexSizesWithDefaults() *ViewIndexSizes`

NewViewIndexSizesWithDefaults instantiates a new ViewIndexSizes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ViewIndexSizes) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ViewIndexSizes) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ViewIndexSizes) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *ViewIndexSizes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDisk

`func (o *ViewIndexSizes) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ViewIndexSizes) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ViewIndexSizes) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *ViewIndexSizes) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetExternal

`func (o *ViewIndexSizes) GetExternal() int32`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *ViewIndexSizes) GetExternalOk() (*int32, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *ViewIndexSizes) SetExternal(v int32)`

SetExternal sets External field to given value.

### HasExternal

`func (o *ViewIndexSizes) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


