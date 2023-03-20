# DatabaseCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N** | Pointer to **int32** | Replicas. The number of copies of every document. | [optional] 
**Q** | Pointer to **int32** | Shards. The number of range partitions. | [optional] 
**R** | Pointer to **int32** | Read quorum. The number of consistent copies of a document that need to be read before a successful reply. | [optional] 
**W** | Pointer to **int32** | Write quorum. The number of copies of a document that need to be written before a successful reply. | [optional] 

## Methods

### NewDatabaseCluster

`func NewDatabaseCluster() *DatabaseCluster`

NewDatabaseCluster instantiates a new DatabaseCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseClusterWithDefaults

`func NewDatabaseClusterWithDefaults() *DatabaseCluster`

NewDatabaseClusterWithDefaults instantiates a new DatabaseCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN

`func (o *DatabaseCluster) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *DatabaseCluster) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *DatabaseCluster) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *DatabaseCluster) HasN() bool`

HasN returns a boolean if a field has been set.

### GetQ

`func (o *DatabaseCluster) GetQ() int32`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *DatabaseCluster) GetQOk() (*int32, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *DatabaseCluster) SetQ(v int32)`

SetQ sets Q field to given value.

### HasQ

`func (o *DatabaseCluster) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetR

`func (o *DatabaseCluster) GetR() int32`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *DatabaseCluster) GetROk() (*int32, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *DatabaseCluster) SetR(v int32)`

SetR sets R field to given value.

### HasR

`func (o *DatabaseCluster) HasR() bool`

HasR returns a boolean if a field has been set.

### GetW

`func (o *DatabaseCluster) GetW() int32`

GetW returns the W field if non-nil, zero value otherwise.

### GetWOk

`func (o *DatabaseCluster) GetWOk() (*int32, bool)`

GetWOk returns a tuple with the W field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW

`func (o *DatabaseCluster) SetW(v int32)`

SetW sets W field to given value.

### HasW

`func (o *DatabaseCluster) HasW() bool`

HasW returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


