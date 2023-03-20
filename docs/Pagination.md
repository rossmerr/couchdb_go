# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int32** | Offset where the document list started. | [optional] 
**TotalRows** | Pointer to **int32** | Number of documents in the database/view. | [optional] 
**Rows** | Pointer to [**[]Row**](Row.md) | Array of view row objects. By default the information returned contains only the document ID and revision. | [optional] 
**UpdateSeq** | Pointer to **map[string]interface{}** | Current update sequence for the database. | [optional] 

## Methods

### NewPagination

`func NewPagination() *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *Pagination) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Pagination) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Pagination) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Pagination) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotalRows

`func (o *Pagination) GetTotalRows() int32`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *Pagination) GetTotalRowsOk() (*int32, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *Pagination) SetTotalRows(v int32)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *Pagination) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.

### GetRows

`func (o *Pagination) GetRows() []Row`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *Pagination) GetRowsOk() (*[]Row, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *Pagination) SetRows(v []Row)`

SetRows sets Rows field to given value.

### HasRows

`func (o *Pagination) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetUpdateSeq

`func (o *Pagination) GetUpdateSeq() map[string]interface{}`

GetUpdateSeq returns the UpdateSeq field if non-nil, zero value otherwise.

### GetUpdateSeqOk

`func (o *Pagination) GetUpdateSeqOk() (*map[string]interface{}, bool)`

GetUpdateSeqOk returns a tuple with the UpdateSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSeq

`func (o *Pagination) SetUpdateSeq(v map[string]interface{})`

SetUpdateSeq sets UpdateSeq field to given value.

### HasUpdateSeq

`func (o *Pagination) HasUpdateSeq() bool`

HasUpdateSeq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


