# Database

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbName** | Pointer to **string** | The name of the database. | [optional] 
**Cluster** | Pointer to [**DatabaseCluster**](DatabaseCluster.md) |  | [optional] 
**CompactRunning** | Pointer to **bool** | Set to true if the database compaction routine is operating on this database. | [optional] 
**DiskFormatVersion** | Pointer to **int32** | The version of the physical format used for the data when it is stored on disk. | [optional] 
**DocCount** | Pointer to **int32** | A count of the documents in the specified database. | [optional] 
**DocDelCount** | Pointer to **int32** | Number of deleted documents | [optional] 
**InstanceStartTime** | Pointer to **string** | Always \&quot;0\&quot;. (Returned for legacy reasons.) | [optional] 
**PurgeSeq** | Pointer to **string** | An opaque string that describes the purge state of the database. Do not rely on this string for counting the number of purge operations. | [optional] 
**Sizes** | Pointer to [**DatabaseSizes**](DatabaseSizes.md) |  | [optional] 
**UpdateSeq** | Pointer to **string** | Always \&quot;0\&quot;. (Returned for legacy reasons.) | [optional] 
**Props** | Pointer to [**DatabaseProps**](DatabaseProps.md) |  | [optional] 

## Methods

### NewDatabase

`func NewDatabase() *Database`

NewDatabase instantiates a new Database object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseWithDefaults

`func NewDatabaseWithDefaults() *Database`

NewDatabaseWithDefaults instantiates a new Database object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbName

`func (o *Database) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *Database) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *Database) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *Database) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetCluster

`func (o *Database) GetCluster() DatabaseCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Database) GetClusterOk() (*DatabaseCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Database) SetCluster(v DatabaseCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Database) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCompactRunning

`func (o *Database) GetCompactRunning() bool`

GetCompactRunning returns the CompactRunning field if non-nil, zero value otherwise.

### GetCompactRunningOk

`func (o *Database) GetCompactRunningOk() (*bool, bool)`

GetCompactRunningOk returns a tuple with the CompactRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactRunning

`func (o *Database) SetCompactRunning(v bool)`

SetCompactRunning sets CompactRunning field to given value.

### HasCompactRunning

`func (o *Database) HasCompactRunning() bool`

HasCompactRunning returns a boolean if a field has been set.

### GetDiskFormatVersion

`func (o *Database) GetDiskFormatVersion() int32`

GetDiskFormatVersion returns the DiskFormatVersion field if non-nil, zero value otherwise.

### GetDiskFormatVersionOk

`func (o *Database) GetDiskFormatVersionOk() (*int32, bool)`

GetDiskFormatVersionOk returns a tuple with the DiskFormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFormatVersion

`func (o *Database) SetDiskFormatVersion(v int32)`

SetDiskFormatVersion sets DiskFormatVersion field to given value.

### HasDiskFormatVersion

`func (o *Database) HasDiskFormatVersion() bool`

HasDiskFormatVersion returns a boolean if a field has been set.

### GetDocCount

`func (o *Database) GetDocCount() int32`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *Database) GetDocCountOk() (*int32, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *Database) SetDocCount(v int32)`

SetDocCount sets DocCount field to given value.

### HasDocCount

`func (o *Database) HasDocCount() bool`

HasDocCount returns a boolean if a field has been set.

### GetDocDelCount

`func (o *Database) GetDocDelCount() int32`

GetDocDelCount returns the DocDelCount field if non-nil, zero value otherwise.

### GetDocDelCountOk

`func (o *Database) GetDocDelCountOk() (*int32, bool)`

GetDocDelCountOk returns a tuple with the DocDelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocDelCount

`func (o *Database) SetDocDelCount(v int32)`

SetDocDelCount sets DocDelCount field to given value.

### HasDocDelCount

`func (o *Database) HasDocDelCount() bool`

HasDocDelCount returns a boolean if a field has been set.

### GetInstanceStartTime

`func (o *Database) GetInstanceStartTime() string`

GetInstanceStartTime returns the InstanceStartTime field if non-nil, zero value otherwise.

### GetInstanceStartTimeOk

`func (o *Database) GetInstanceStartTimeOk() (*string, bool)`

GetInstanceStartTimeOk returns a tuple with the InstanceStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStartTime

`func (o *Database) SetInstanceStartTime(v string)`

SetInstanceStartTime sets InstanceStartTime field to given value.

### HasInstanceStartTime

`func (o *Database) HasInstanceStartTime() bool`

HasInstanceStartTime returns a boolean if a field has been set.

### GetPurgeSeq

`func (o *Database) GetPurgeSeq() string`

GetPurgeSeq returns the PurgeSeq field if non-nil, zero value otherwise.

### GetPurgeSeqOk

`func (o *Database) GetPurgeSeqOk() (*string, bool)`

GetPurgeSeqOk returns a tuple with the PurgeSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeSeq

`func (o *Database) SetPurgeSeq(v string)`

SetPurgeSeq sets PurgeSeq field to given value.

### HasPurgeSeq

`func (o *Database) HasPurgeSeq() bool`

HasPurgeSeq returns a boolean if a field has been set.

### GetSizes

`func (o *Database) GetSizes() DatabaseSizes`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *Database) GetSizesOk() (*DatabaseSizes, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *Database) SetSizes(v DatabaseSizes)`

SetSizes sets Sizes field to given value.

### HasSizes

`func (o *Database) HasSizes() bool`

HasSizes returns a boolean if a field has been set.

### GetUpdateSeq

`func (o *Database) GetUpdateSeq() string`

GetUpdateSeq returns the UpdateSeq field if non-nil, zero value otherwise.

### GetUpdateSeqOk

`func (o *Database) GetUpdateSeqOk() (*string, bool)`

GetUpdateSeqOk returns a tuple with the UpdateSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSeq

`func (o *Database) SetUpdateSeq(v string)`

SetUpdateSeq sets UpdateSeq field to given value.

### HasUpdateSeq

`func (o *Database) HasUpdateSeq() bool`

HasUpdateSeq returns a boolean if a field has been set.

### GetProps

`func (o *Database) GetProps() DatabaseProps`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *Database) GetPropsOk() (*DatabaseProps, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *Database) SetProps(v DatabaseProps)`

SetProps sets Props field to given value.

### HasProps

`func (o *Database) HasProps() bool`

HasProps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


