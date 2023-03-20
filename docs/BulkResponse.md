# BulkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | Pointer to **bool** | Operation status | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Rev** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** | Error type. Available if response code is 4xx | [optional] 
**Reason** | Pointer to **string** | Error description. Available if response code is 4xx | [optional] 

## Methods

### NewBulkResponse

`func NewBulkResponse() *BulkResponse`

NewBulkResponse instantiates a new BulkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResponseWithDefaults

`func NewBulkResponseWithDefaults() *BulkResponse`

NewBulkResponseWithDefaults instantiates a new BulkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *BulkResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *BulkResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *BulkResponse) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *BulkResponse) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetId

`func (o *BulkResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRev

`func (o *BulkResponse) GetRev() string`

GetRev returns the Rev field if non-nil, zero value otherwise.

### GetRevOk

`func (o *BulkResponse) GetRevOk() (*string, bool)`

GetRevOk returns a tuple with the Rev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRev

`func (o *BulkResponse) SetRev(v string)`

SetRev sets Rev field to given value.

### HasRev

`func (o *BulkResponse) HasRev() bool`

HasRev returns a boolean if a field has been set.

### GetError

`func (o *BulkResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BulkResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BulkResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BulkResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetReason

`func (o *BulkResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BulkResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BulkResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BulkResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


