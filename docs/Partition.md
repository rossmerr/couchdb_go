# Partition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbName** | Pointer to **string** |  | [optional] 
**DocCount** | Pointer to **int32** |  | [optional] 
**DocDelCount** | Pointer to **int32** |  | [optional] 
**Partition** | Pointer to **string** |  | [optional] 
**Sizes** | Pointer to [**PartitionSizes**](PartitionSizes.md) |  | [optional] 

## Methods

### NewPartition

`func NewPartition() *Partition`

NewPartition instantiates a new Partition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionWithDefaults

`func NewPartitionWithDefaults() *Partition`

NewPartitionWithDefaults instantiates a new Partition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbName

`func (o *Partition) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *Partition) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *Partition) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *Partition) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDocCount

`func (o *Partition) GetDocCount() int32`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *Partition) GetDocCountOk() (*int32, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *Partition) SetDocCount(v int32)`

SetDocCount sets DocCount field to given value.

### HasDocCount

`func (o *Partition) HasDocCount() bool`

HasDocCount returns a boolean if a field has been set.

### GetDocDelCount

`func (o *Partition) GetDocDelCount() int32`

GetDocDelCount returns the DocDelCount field if non-nil, zero value otherwise.

### GetDocDelCountOk

`func (o *Partition) GetDocDelCountOk() (*int32, bool)`

GetDocDelCountOk returns a tuple with the DocDelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocDelCount

`func (o *Partition) SetDocDelCount(v int32)`

SetDocDelCount sets DocDelCount field to given value.

### HasDocDelCount

`func (o *Partition) HasDocDelCount() bool`

HasDocDelCount returns a boolean if a field has been set.

### GetPartition

`func (o *Partition) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *Partition) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *Partition) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *Partition) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetSizes

`func (o *Partition) GetSizes() PartitionSizes`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *Partition) GetSizesOk() (*PartitionSizes, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *Partition) SetSizes(v PartitionSizes)`

SetSizes sets Sizes field to given value.

### HasSizes

`func (o *Partition) HasSizes() bool`

HasSizes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


