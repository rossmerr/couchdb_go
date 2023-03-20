# Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **float32** | Maximum number of results returned. Default is 25. Optional | [optional] 
**Skip** | Pointer to **float32** | Skip the first ‘n’ results, where ‘n’ is the value specified. Optional | [optional] 
**Name** | Pointer to **string** | Name of the index. If no name is provided, a name will be generated automatically. Optional | [optional] 
**Sort** | Pointer to **map[string]interface{}** |  | [optional] 
**Field** | Pointer to **[]string** | Instruct a query to use a specific index. Specified either as \&quot;&lt;design_document&gt;\&quot; or [\&quot;&lt;design_document&gt;\&quot;, \&quot;&lt;index_name&gt;\&quot;]. Optional | [optional] 
**UseIndex** | Pointer to **[]string** | JSON array specifying which fields of each object should be returned. If it is omitted, the entire object is returned. More information provided in the section on filtering fields. Optional | [optional] 
**Conflitsc** | Pointer to **bool** | Include conflicted documents if true. Intended use is to easily find conflicted documents, without an index or view. Default is false. Optional | [optional] 
**Bookmark** | Pointer to **string** | A string that enables you to specify which page of results you require. Used for paging through result sets. Every query returns an opaque string under the bookmark key that can then be passed back in a query to get the next page of results. If any part of the selector query changes between requests, the results are undefined. Optional, default: null  | [optional] 
**Update** | Pointer to **bool** | Whether to update the index prior to returning the result. Default is true. Optional | [optional] 
**Stable** | Pointer to **bool** | Whether or not the view results should be returned from a “stable” set of shards. Optional | [optional] 
**Stale** | Pointer to **string** | Combination of update&#x3D;false and stable&#x3D;true options. Possible options: \&quot;ok\&quot;, false (default). Optional Note that this parameter is deprecated. Use stable and update instead. See Views Generation for more details.  | [optional] 
**ExecutionStats** | Pointer to **bool** | Include execution statistics in the query response. Optional, default: false  | [optional] 
**Selector** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewQuery

`func NewQuery() *Query`

NewQuery instantiates a new Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryWithDefaults

`func NewQueryWithDefaults() *Query`

NewQueryWithDefaults instantiates a new Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *Query) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Query) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Query) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Query) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSkip

`func (o *Query) GetSkip() float32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *Query) GetSkipOk() (*float32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *Query) SetSkip(v float32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *Query) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetName

`func (o *Query) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Query) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Query) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Query) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSort

`func (o *Query) GetSort() map[string]interface{}`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Query) GetSortOk() (*map[string]interface{}, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Query) SetSort(v map[string]interface{})`

SetSort sets Sort field to given value.

### HasSort

`func (o *Query) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetField

`func (o *Query) GetField() []string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Query) GetFieldOk() (*[]string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Query) SetField(v []string)`

SetField sets Field field to given value.

### HasField

`func (o *Query) HasField() bool`

HasField returns a boolean if a field has been set.

### GetUseIndex

`func (o *Query) GetUseIndex() []string`

GetUseIndex returns the UseIndex field if non-nil, zero value otherwise.

### GetUseIndexOk

`func (o *Query) GetUseIndexOk() (*[]string, bool)`

GetUseIndexOk returns a tuple with the UseIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIndex

`func (o *Query) SetUseIndex(v []string)`

SetUseIndex sets UseIndex field to given value.

### HasUseIndex

`func (o *Query) HasUseIndex() bool`

HasUseIndex returns a boolean if a field has been set.

### GetConflitsc

`func (o *Query) GetConflitsc() bool`

GetConflitsc returns the Conflitsc field if non-nil, zero value otherwise.

### GetConflitscOk

`func (o *Query) GetConflitscOk() (*bool, bool)`

GetConflitscOk returns a tuple with the Conflitsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflitsc

`func (o *Query) SetConflitsc(v bool)`

SetConflitsc sets Conflitsc field to given value.

### HasConflitsc

`func (o *Query) HasConflitsc() bool`

HasConflitsc returns a boolean if a field has been set.

### GetBookmark

`func (o *Query) GetBookmark() string`

GetBookmark returns the Bookmark field if non-nil, zero value otherwise.

### GetBookmarkOk

`func (o *Query) GetBookmarkOk() (*string, bool)`

GetBookmarkOk returns a tuple with the Bookmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmark

`func (o *Query) SetBookmark(v string)`

SetBookmark sets Bookmark field to given value.

### HasBookmark

`func (o *Query) HasBookmark() bool`

HasBookmark returns a boolean if a field has been set.

### GetUpdate

`func (o *Query) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Query) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Query) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Query) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetStable

`func (o *Query) GetStable() bool`

GetStable returns the Stable field if non-nil, zero value otherwise.

### GetStableOk

`func (o *Query) GetStableOk() (*bool, bool)`

GetStableOk returns a tuple with the Stable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStable

`func (o *Query) SetStable(v bool)`

SetStable sets Stable field to given value.

### HasStable

`func (o *Query) HasStable() bool`

HasStable returns a boolean if a field has been set.

### GetStale

`func (o *Query) GetStale() string`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *Query) GetStaleOk() (*string, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *Query) SetStale(v string)`

SetStale sets Stale field to given value.

### HasStale

`func (o *Query) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetExecutionStats

`func (o *Query) GetExecutionStats() bool`

GetExecutionStats returns the ExecutionStats field if non-nil, zero value otherwise.

### GetExecutionStatsOk

`func (o *Query) GetExecutionStatsOk() (*bool, bool)`

GetExecutionStatsOk returns a tuple with the ExecutionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStats

`func (o *Query) SetExecutionStats(v bool)`

SetExecutionStats sets ExecutionStats field to given value.

### HasExecutionStats

`func (o *Query) HasExecutionStats() bool`

HasExecutionStats returns a boolean if a field has been set.

### GetSelector

`func (o *Query) GetSelector() map[string]interface{}`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *Query) GetSelectorOk() (*map[string]interface{}, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *Query) SetSelector(v map[string]interface{})`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *Query) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


