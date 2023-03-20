# DatabaseProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partitioned** | Pointer to **bool** | If present and true, this indicates that the database is partitioned. | [optional] 

## Methods

### NewDatabaseProps

`func NewDatabaseProps() *DatabaseProps`

NewDatabaseProps instantiates a new DatabaseProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabasePropsWithDefaults

`func NewDatabasePropsWithDefaults() *DatabaseProps`

NewDatabasePropsWithDefaults instantiates a new DatabaseProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartitioned

`func (o *DatabaseProps) GetPartitioned() bool`

GetPartitioned returns the Partitioned field if non-nil, zero value otherwise.

### GetPartitionedOk

`func (o *DatabaseProps) GetPartitionedOk() (*bool, bool)`

GetPartitionedOk returns a tuple with the Partitioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioned

`func (o *DatabaseProps) SetPartitioned(v bool)`

SetPartitioned sets Partitioned field to given value.

### HasPartitioned

`func (o *DatabaseProps) HasPartitioned() bool`

HasPartitioned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


