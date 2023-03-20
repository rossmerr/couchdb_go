# ReplicationHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocWriteFailures** | Pointer to **int32** | Number of document write failures | [optional] 
**DocsRead** | Pointer to **int32** | Number of documents read | [optional] 
**DocsWritten** | Pointer to **int32** | Number of documents written to target | [optional] 
**EndLastSeq** | Pointer to **int32** | Last sequence number in changes stream | [optional] 
**EndTime** | Pointer to **string** | Date/Time replication operation completed in RFC 2822 format | [optional] 
**MissingChecked** | Pointer to **int32** | Number of missing documents checked | [optional] 
**MissingFound** | Pointer to **int32** | Number of missing documents found | [optional] 
**RecordedSeq** | Pointer to **int32** | Last recorded sequence number | [optional] 
**SessionId** | Pointer to **string** | Session ID for this replication operation | [optional] 
**StartLastSeq** | Pointer to **int32** | First sequence number in changes stream | [optional] 
**StartTime** | Pointer to **string** | Date/Time replication operation started in RFC 2822 format | [optional] 

## Methods

### NewReplicationHistory

`func NewReplicationHistory() *ReplicationHistory`

NewReplicationHistory instantiates a new ReplicationHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationHistoryWithDefaults

`func NewReplicationHistoryWithDefaults() *ReplicationHistory`

NewReplicationHistoryWithDefaults instantiates a new ReplicationHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocWriteFailures

`func (o *ReplicationHistory) GetDocWriteFailures() int32`

GetDocWriteFailures returns the DocWriteFailures field if non-nil, zero value otherwise.

### GetDocWriteFailuresOk

`func (o *ReplicationHistory) GetDocWriteFailuresOk() (*int32, bool)`

GetDocWriteFailuresOk returns a tuple with the DocWriteFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocWriteFailures

`func (o *ReplicationHistory) SetDocWriteFailures(v int32)`

SetDocWriteFailures sets DocWriteFailures field to given value.

### HasDocWriteFailures

`func (o *ReplicationHistory) HasDocWriteFailures() bool`

HasDocWriteFailures returns a boolean if a field has been set.

### GetDocsRead

`func (o *ReplicationHistory) GetDocsRead() int32`

GetDocsRead returns the DocsRead field if non-nil, zero value otherwise.

### GetDocsReadOk

`func (o *ReplicationHistory) GetDocsReadOk() (*int32, bool)`

GetDocsReadOk returns a tuple with the DocsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsRead

`func (o *ReplicationHistory) SetDocsRead(v int32)`

SetDocsRead sets DocsRead field to given value.

### HasDocsRead

`func (o *ReplicationHistory) HasDocsRead() bool`

HasDocsRead returns a boolean if a field has been set.

### GetDocsWritten

`func (o *ReplicationHistory) GetDocsWritten() int32`

GetDocsWritten returns the DocsWritten field if non-nil, zero value otherwise.

### GetDocsWrittenOk

`func (o *ReplicationHistory) GetDocsWrittenOk() (*int32, bool)`

GetDocsWrittenOk returns a tuple with the DocsWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsWritten

`func (o *ReplicationHistory) SetDocsWritten(v int32)`

SetDocsWritten sets DocsWritten field to given value.

### HasDocsWritten

`func (o *ReplicationHistory) HasDocsWritten() bool`

HasDocsWritten returns a boolean if a field has been set.

### GetEndLastSeq

`func (o *ReplicationHistory) GetEndLastSeq() int32`

GetEndLastSeq returns the EndLastSeq field if non-nil, zero value otherwise.

### GetEndLastSeqOk

`func (o *ReplicationHistory) GetEndLastSeqOk() (*int32, bool)`

GetEndLastSeqOk returns a tuple with the EndLastSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLastSeq

`func (o *ReplicationHistory) SetEndLastSeq(v int32)`

SetEndLastSeq sets EndLastSeq field to given value.

### HasEndLastSeq

`func (o *ReplicationHistory) HasEndLastSeq() bool`

HasEndLastSeq returns a boolean if a field has been set.

### GetEndTime

`func (o *ReplicationHistory) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ReplicationHistory) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ReplicationHistory) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ReplicationHistory) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMissingChecked

`func (o *ReplicationHistory) GetMissingChecked() int32`

GetMissingChecked returns the MissingChecked field if non-nil, zero value otherwise.

### GetMissingCheckedOk

`func (o *ReplicationHistory) GetMissingCheckedOk() (*int32, bool)`

GetMissingCheckedOk returns a tuple with the MissingChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingChecked

`func (o *ReplicationHistory) SetMissingChecked(v int32)`

SetMissingChecked sets MissingChecked field to given value.

### HasMissingChecked

`func (o *ReplicationHistory) HasMissingChecked() bool`

HasMissingChecked returns a boolean if a field has been set.

### GetMissingFound

`func (o *ReplicationHistory) GetMissingFound() int32`

GetMissingFound returns the MissingFound field if non-nil, zero value otherwise.

### GetMissingFoundOk

`func (o *ReplicationHistory) GetMissingFoundOk() (*int32, bool)`

GetMissingFoundOk returns a tuple with the MissingFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingFound

`func (o *ReplicationHistory) SetMissingFound(v int32)`

SetMissingFound sets MissingFound field to given value.

### HasMissingFound

`func (o *ReplicationHistory) HasMissingFound() bool`

HasMissingFound returns a boolean if a field has been set.

### GetRecordedSeq

`func (o *ReplicationHistory) GetRecordedSeq() int32`

GetRecordedSeq returns the RecordedSeq field if non-nil, zero value otherwise.

### GetRecordedSeqOk

`func (o *ReplicationHistory) GetRecordedSeqOk() (*int32, bool)`

GetRecordedSeqOk returns a tuple with the RecordedSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedSeq

`func (o *ReplicationHistory) SetRecordedSeq(v int32)`

SetRecordedSeq sets RecordedSeq field to given value.

### HasRecordedSeq

`func (o *ReplicationHistory) HasRecordedSeq() bool`

HasRecordedSeq returns a boolean if a field has been set.

### GetSessionId

`func (o *ReplicationHistory) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ReplicationHistory) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ReplicationHistory) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ReplicationHistory) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStartLastSeq

`func (o *ReplicationHistory) GetStartLastSeq() int32`

GetStartLastSeq returns the StartLastSeq field if non-nil, zero value otherwise.

### GetStartLastSeqOk

`func (o *ReplicationHistory) GetStartLastSeqOk() (*int32, bool)`

GetStartLastSeqOk returns a tuple with the StartLastSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLastSeq

`func (o *ReplicationHistory) SetStartLastSeq(v int32)`

SetStartLastSeq sets StartLastSeq field to given value.

### HasStartLastSeq

`func (o *ReplicationHistory) HasStartLastSeq() bool`

HasStartLastSeq returns a boolean if a field has been set.

### GetStartTime

`func (o *ReplicationHistory) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ReplicationHistory) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ReplicationHistory) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ReplicationHistory) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


