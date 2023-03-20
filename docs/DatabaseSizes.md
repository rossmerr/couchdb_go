# DatabaseSizes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** | The size of live data inside the database, in bytes. | [optional] 
**External** | Pointer to **int32** | The uncompressed size of database contents in bytes. sizes.file (number) – The size of the database file on disk in bytes. Views indexes are not included in the calculation. | [optional] 
**File** | Pointer to **int32** | An opaque string that describes the state of the database. Do not rely on this string for counting the number of updates. | [optional] 

## Methods

### NewDatabaseSizes

`func NewDatabaseSizes() *DatabaseSizes`

NewDatabaseSizes instantiates a new DatabaseSizes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSizesWithDefaults

`func NewDatabaseSizesWithDefaults() *DatabaseSizes`

NewDatabaseSizesWithDefaults instantiates a new DatabaseSizes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *DatabaseSizes) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DatabaseSizes) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DatabaseSizes) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *DatabaseSizes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetExternal

`func (o *DatabaseSizes) GetExternal() int32`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *DatabaseSizes) GetExternalOk() (*int32, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *DatabaseSizes) SetExternal(v int32)`

SetExternal sets External field to given value.

### HasExternal

`func (o *DatabaseSizes) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetFile

`func (o *DatabaseSizes) GetFile() int32`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *DatabaseSizes) GetFileOk() (*int32, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *DatabaseSizes) SetFile(v int32)`

SetFile sets File field to given value.

### HasFile

`func (o *DatabaseSizes) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


