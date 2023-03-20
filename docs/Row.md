# Row

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Document ID | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**Doc** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRow

`func NewRow() *Row`

NewRow instantiates a new Row object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRowWithDefaults

`func NewRowWithDefaults() *Row`

NewRowWithDefaults instantiates a new Row object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Row) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Row) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Row) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Row) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *Row) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Row) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Row) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Row) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *Row) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Row) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Row) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *Row) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDoc

`func (o *Row) GetDoc() map[string]interface{}`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *Row) GetDocOk() (*map[string]interface{}, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *Row) SetDoc(v map[string]interface{})`

SetDoc sets Doc field to given value.

### HasDoc

`func (o *Row) HasDoc() bool`

HasDoc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


