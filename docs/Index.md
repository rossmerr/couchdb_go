# Index

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ddoc** | Pointer to **string** | Name of the design document in which the index will be created. By default, each index will be created in its own design document. Indexes can be grouped into design documents for efficiency. However, a change to one index in a design document will invalidate all other indexes in the same document (similar to views). Optional | [optional] 
**Name** | Pointer to **string** | Name of the index. If no name is provided, a name will be generated automatically. Optional | [optional] 
**Type** | Pointer to **string** | Defaults to json. Geospatial indexes will be supported in the future. Optional Text indexes are supported via a third party library Optional | [optional] [default to "json"]
**Partitioned** | Pointer to **bool** | Determines whether a JSON index is partitioned or global. The default value of partitioned is the partitioned property of the database. To create a global index on a partitioned database, specify false for the \&quot;partitioned\&quot; field. If you specify true for the \&quot;partitioned\&quot; field on an unpartitioned database, an error occurs. | [optional] 
**Index** | Pointer to [**IndexDefinitionsDefField**](IndexDefinitionsDefField.md) |  | [optional] 

## Methods

### NewIndex

`func NewIndex() *Index`

NewIndex instantiates a new Index object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexWithDefaults

`func NewIndexWithDefaults() *Index`

NewIndexWithDefaults instantiates a new Index object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDdoc

`func (o *Index) GetDdoc() string`

GetDdoc returns the Ddoc field if non-nil, zero value otherwise.

### GetDdocOk

`func (o *Index) GetDdocOk() (*string, bool)`

GetDdocOk returns a tuple with the Ddoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdoc

`func (o *Index) SetDdoc(v string)`

SetDdoc sets Ddoc field to given value.

### HasDdoc

`func (o *Index) HasDdoc() bool`

HasDdoc returns a boolean if a field has been set.

### GetName

`func (o *Index) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Index) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Index) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Index) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Index) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Index) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Index) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Index) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPartitioned

`func (o *Index) GetPartitioned() bool`

GetPartitioned returns the Partitioned field if non-nil, zero value otherwise.

### GetPartitionedOk

`func (o *Index) GetPartitionedOk() (*bool, bool)`

GetPartitionedOk returns a tuple with the Partitioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioned

`func (o *Index) SetPartitioned(v bool)`

SetPartitioned sets Partitioned field to given value.

### HasPartitioned

`func (o *Index) HasPartitioned() bool`

HasPartitioned returns a boolean if a field has been set.

### GetIndex

`func (o *Index) GetIndex() IndexDefinitionsDefField`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Index) GetIndexOk() (*IndexDefinitionsDefField, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Index) SetIndex(v IndexDefinitionsDefField)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Index) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


