# ErrorBulkGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error type. Available if response code is 4xx | [optional] 
**Reason** | Pointer to **string** | Error description. Available if response code is 4xx | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Rev** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorBulkGetResponse

`func NewErrorBulkGetResponse() *ErrorBulkGetResponse`

NewErrorBulkGetResponse instantiates a new ErrorBulkGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorBulkGetResponseWithDefaults

`func NewErrorBulkGetResponseWithDefaults() *ErrorBulkGetResponse`

NewErrorBulkGetResponseWithDefaults instantiates a new ErrorBulkGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ErrorBulkGetResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorBulkGetResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorBulkGetResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ErrorBulkGetResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetReason

`func (o *ErrorBulkGetResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorBulkGetResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorBulkGetResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorBulkGetResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetId

`func (o *ErrorBulkGetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorBulkGetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorBulkGetResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ErrorBulkGetResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRev

`func (o *ErrorBulkGetResponse) GetRev() string`

GetRev returns the Rev field if non-nil, zero value otherwise.

### GetRevOk

`func (o *ErrorBulkGetResponse) GetRevOk() (*string, bool)`

GetRevOk returns a tuple with the Rev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRev

`func (o *ErrorBulkGetResponse) SetRev(v string)`

SetRev sets Rev field to given value.

### HasRev

`func (o *ErrorBulkGetResponse) HasRev() bool`

HasRev returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


