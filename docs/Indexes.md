# Indexes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRows** | Pointer to **float32** |  | [optional] 
**Indexes** | Pointer to [**[]IndexDefinitions**](IndexDefinitions.md) |  | [optional] 

## Methods

### NewIndexes

`func NewIndexes() *Indexes`

NewIndexes instantiates a new Indexes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexesWithDefaults

`func NewIndexesWithDefaults() *Indexes`

NewIndexesWithDefaults instantiates a new Indexes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRows

`func (o *Indexes) GetTotalRows() float32`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *Indexes) GetTotalRowsOk() (*float32, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *Indexes) SetTotalRows(v float32)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *Indexes) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.

### GetIndexes

`func (o *Indexes) GetIndexes() []IndexDefinitions`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *Indexes) GetIndexesOk() (*[]IndexDefinitions, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *Indexes) SetIndexes(v []IndexDefinitions)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *Indexes) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


