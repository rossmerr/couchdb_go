# ViewIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompactRunning** | Pointer to **bool** | Indicates whether a compaction routine is currently running on the view | [optional] 
**Language** | Pointer to **string** | Language for the defined views | [optional] 
**PurgeSeq** | Pointer to **int32** | The purge sequence that has been processed | [optional] 
**Signature** | Pointer to **string** | MD5 signature of the views for the design document | [optional] 
**Sizes** | Pointer to [**ViewIndexSizes**](ViewIndexSizes.md) |  | [optional] 
**UpdateSeq** | Pointer to **string** | The update sequence of the corresponding database that has been indexed | [optional] 
**UpdaterRunning** | Pointer to **bool** | Indicates if the view is currently being updated | [optional] 
**WaitingClients** | Pointer to **int32** | Number of clients waiting on views from this design document | [optional] 
**WaitingCommit** | Pointer to **bool** | Indicates if there are outstanding commits to the underlying database that need to processed | [optional] 

## Methods

### NewViewIndex

`func NewViewIndex() *ViewIndex`

NewViewIndex instantiates a new ViewIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewIndexWithDefaults

`func NewViewIndexWithDefaults() *ViewIndex`

NewViewIndexWithDefaults instantiates a new ViewIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompactRunning

`func (o *ViewIndex) GetCompactRunning() bool`

GetCompactRunning returns the CompactRunning field if non-nil, zero value otherwise.

### GetCompactRunningOk

`func (o *ViewIndex) GetCompactRunningOk() (*bool, bool)`

GetCompactRunningOk returns a tuple with the CompactRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactRunning

`func (o *ViewIndex) SetCompactRunning(v bool)`

SetCompactRunning sets CompactRunning field to given value.

### HasCompactRunning

`func (o *ViewIndex) HasCompactRunning() bool`

HasCompactRunning returns a boolean if a field has been set.

### GetLanguage

`func (o *ViewIndex) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ViewIndex) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ViewIndex) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ViewIndex) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPurgeSeq

`func (o *ViewIndex) GetPurgeSeq() int32`

GetPurgeSeq returns the PurgeSeq field if non-nil, zero value otherwise.

### GetPurgeSeqOk

`func (o *ViewIndex) GetPurgeSeqOk() (*int32, bool)`

GetPurgeSeqOk returns a tuple with the PurgeSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeSeq

`func (o *ViewIndex) SetPurgeSeq(v int32)`

SetPurgeSeq sets PurgeSeq field to given value.

### HasPurgeSeq

`func (o *ViewIndex) HasPurgeSeq() bool`

HasPurgeSeq returns a boolean if a field has been set.

### GetSignature

`func (o *ViewIndex) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ViewIndex) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ViewIndex) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ViewIndex) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSizes

`func (o *ViewIndex) GetSizes() ViewIndexSizes`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *ViewIndex) GetSizesOk() (*ViewIndexSizes, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *ViewIndex) SetSizes(v ViewIndexSizes)`

SetSizes sets Sizes field to given value.

### HasSizes

`func (o *ViewIndex) HasSizes() bool`

HasSizes returns a boolean if a field has been set.

### GetUpdateSeq

`func (o *ViewIndex) GetUpdateSeq() string`

GetUpdateSeq returns the UpdateSeq field if non-nil, zero value otherwise.

### GetUpdateSeqOk

`func (o *ViewIndex) GetUpdateSeqOk() (*string, bool)`

GetUpdateSeqOk returns a tuple with the UpdateSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSeq

`func (o *ViewIndex) SetUpdateSeq(v string)`

SetUpdateSeq sets UpdateSeq field to given value.

### HasUpdateSeq

`func (o *ViewIndex) HasUpdateSeq() bool`

HasUpdateSeq returns a boolean if a field has been set.

### GetUpdaterRunning

`func (o *ViewIndex) GetUpdaterRunning() bool`

GetUpdaterRunning returns the UpdaterRunning field if non-nil, zero value otherwise.

### GetUpdaterRunningOk

`func (o *ViewIndex) GetUpdaterRunningOk() (*bool, bool)`

GetUpdaterRunningOk returns a tuple with the UpdaterRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterRunning

`func (o *ViewIndex) SetUpdaterRunning(v bool)`

SetUpdaterRunning sets UpdaterRunning field to given value.

### HasUpdaterRunning

`func (o *ViewIndex) HasUpdaterRunning() bool`

HasUpdaterRunning returns a boolean if a field has been set.

### GetWaitingClients

`func (o *ViewIndex) GetWaitingClients() int32`

GetWaitingClients returns the WaitingClients field if non-nil, zero value otherwise.

### GetWaitingClientsOk

`func (o *ViewIndex) GetWaitingClientsOk() (*int32, bool)`

GetWaitingClientsOk returns a tuple with the WaitingClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingClients

`func (o *ViewIndex) SetWaitingClients(v int32)`

SetWaitingClients sets WaitingClients field to given value.

### HasWaitingClients

`func (o *ViewIndex) HasWaitingClients() bool`

HasWaitingClients returns a boolean if a field has been set.

### GetWaitingCommit

`func (o *ViewIndex) GetWaitingCommit() bool`

GetWaitingCommit returns the WaitingCommit field if non-nil, zero value otherwise.

### GetWaitingCommitOk

`func (o *ViewIndex) GetWaitingCommitOk() (*bool, bool)`

GetWaitingCommitOk returns a tuple with the WaitingCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingCommit

`func (o *ViewIndex) SetWaitingCommit(v bool)`

SetWaitingCommit sets WaitingCommit field to given value.

### HasWaitingCommit

`func (o *ViewIndex) HasWaitingCommit() bool`

HasWaitingCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


