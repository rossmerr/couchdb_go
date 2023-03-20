# DesignDoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to **string** | Defines Query Server to process design document functions | [optional] 
**Options** | Pointer to **map[string]interface{}** | View’s default options | [optional] 
**Filters** | Pointer to **map[string]interface{}** | Filter functions definition | [optional] 
**Updates** | Pointer to **map[string]interface{}** | Update functions definition | [optional] 
**ValidateDocUpdate** | Pointer to **string** | Validate document update function source | [optional] 
**Views** | Pointer to **map[string]interface{}** | View functions definition. | [optional] 
**Indexes** | Pointer to **map[string]interface{}** |  | [optional] 
**Autoupdate** | Pointer to **bool** | Indicates whether to automatically build indexes defined in this design document. Default is true. | [optional] 

## Methods

### NewDesignDoc

`func NewDesignDoc() *DesignDoc`

NewDesignDoc instantiates a new DesignDoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesignDocWithDefaults

`func NewDesignDocWithDefaults() *DesignDoc`

NewDesignDocWithDefaults instantiates a new DesignDoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *DesignDoc) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *DesignDoc) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *DesignDoc) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *DesignDoc) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetOptions

`func (o *DesignDoc) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DesignDoc) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DesignDoc) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DesignDoc) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *DesignDoc) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DesignDoc) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DesignDoc) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DesignDoc) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetUpdates

`func (o *DesignDoc) GetUpdates() map[string]interface{}`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *DesignDoc) GetUpdatesOk() (*map[string]interface{}, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *DesignDoc) SetUpdates(v map[string]interface{})`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *DesignDoc) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetValidateDocUpdate

`func (o *DesignDoc) GetValidateDocUpdate() string`

GetValidateDocUpdate returns the ValidateDocUpdate field if non-nil, zero value otherwise.

### GetValidateDocUpdateOk

`func (o *DesignDoc) GetValidateDocUpdateOk() (*string, bool)`

GetValidateDocUpdateOk returns a tuple with the ValidateDocUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateDocUpdate

`func (o *DesignDoc) SetValidateDocUpdate(v string)`

SetValidateDocUpdate sets ValidateDocUpdate field to given value.

### HasValidateDocUpdate

`func (o *DesignDoc) HasValidateDocUpdate() bool`

HasValidateDocUpdate returns a boolean if a field has been set.

### GetViews

`func (o *DesignDoc) GetViews() map[string]interface{}`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *DesignDoc) GetViewsOk() (*map[string]interface{}, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *DesignDoc) SetViews(v map[string]interface{})`

SetViews sets Views field to given value.

### HasViews

`func (o *DesignDoc) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetIndexes

`func (o *DesignDoc) GetIndexes() map[string]interface{}`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *DesignDoc) GetIndexesOk() (*map[string]interface{}, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *DesignDoc) SetIndexes(v map[string]interface{})`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *DesignDoc) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.

### GetAutoupdate

`func (o *DesignDoc) GetAutoupdate() bool`

GetAutoupdate returns the Autoupdate field if non-nil, zero value otherwise.

### GetAutoupdateOk

`func (o *DesignDoc) GetAutoupdateOk() (*bool, bool)`

GetAutoupdateOk returns a tuple with the Autoupdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoupdate

`func (o *DesignDoc) SetAutoupdate(v bool)`

SetAutoupdate sets Autoupdate field to given value.

### HasAutoupdate

`func (o *DesignDoc) HasAutoupdate() bool`

HasAutoupdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


