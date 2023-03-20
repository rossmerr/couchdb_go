# ResultsDocs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to [**ErrorBulkGetResponse**](ErrorBulkGetResponse.md) |  | [optional] 

## Methods

### NewResultsDocs

`func NewResultsDocs() *ResultsDocs`

NewResultsDocs instantiates a new ResultsDocs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultsDocsWithDefaults

`func NewResultsDocsWithDefaults() *ResultsDocs`

NewResultsDocsWithDefaults instantiates a new ResultsDocs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *ResultsDocs) GetOk() map[string]interface{}`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *ResultsDocs) GetOkOk() (*map[string]interface{}, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *ResultsDocs) SetOk(v map[string]interface{})`

SetOk sets Ok field to given value.

### HasOk

`func (o *ResultsDocs) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetError

`func (o *ResultsDocs) GetError() ErrorBulkGetResponse`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResultsDocs) GetErrorOk() (*ErrorBulkGetResponse, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResultsDocs) SetError(v ErrorBulkGetResponse)`

SetError sets Error field to given value.

### HasError

`func (o *ResultsDocs) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


