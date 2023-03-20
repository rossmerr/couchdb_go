# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangesDone** | Pointer to **int32** | Processed changes | [optional] 
**Database** | Pointer to **string** | Source database | [optional] 
**Pid** | Pointer to **string** | Process ID | [optional] 
**Progress** | Pointer to **int32** | Current percentage progress | [optional] 
**StartedOn** | Pointer to **int32** | ask start time as unix timestamp | [optional] 
**Status** | Pointer to **string** | Task status message | [optional] 
**Task** | Pointer to **string** | Task name | [optional] 
**TotalChanges** | Pointer to **int32** | Total changes to process | [optional] 
**Type** | Pointer to **string** | Operation Type | [optional] 
**UpdatedOn** | Pointer to **string** | Unix timestamp of last operation update | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangesDone

`func (o *Task) GetChangesDone() int32`

GetChangesDone returns the ChangesDone field if non-nil, zero value otherwise.

### GetChangesDoneOk

`func (o *Task) GetChangesDoneOk() (*int32, bool)`

GetChangesDoneOk returns a tuple with the ChangesDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangesDone

`func (o *Task) SetChangesDone(v int32)`

SetChangesDone sets ChangesDone field to given value.

### HasChangesDone

`func (o *Task) HasChangesDone() bool`

HasChangesDone returns a boolean if a field has been set.

### GetDatabase

`func (o *Task) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *Task) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *Task) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *Task) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetPid

`func (o *Task) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *Task) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *Task) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *Task) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetProgress

`func (o *Task) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Task) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Task) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Task) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStartedOn

`func (o *Task) GetStartedOn() int32`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *Task) GetStartedOnOk() (*int32, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *Task) SetStartedOn(v int32)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *Task) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### GetStatus

`func (o *Task) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Task) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Task) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Task) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTask

`func (o *Task) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *Task) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *Task) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *Task) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTotalChanges

`func (o *Task) GetTotalChanges() int32`

GetTotalChanges returns the TotalChanges field if non-nil, zero value otherwise.

### GetTotalChangesOk

`func (o *Task) GetTotalChangesOk() (*int32, bool)`

GetTotalChangesOk returns a tuple with the TotalChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChanges

`func (o *Task) SetTotalChanges(v int32)`

SetTotalChanges sets TotalChanges field to given value.

### HasTotalChanges

`func (o *Task) HasTotalChanges() bool`

HasTotalChanges returns a boolean if a field has been set.

### GetType

`func (o *Task) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Task) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Task) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Task) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *Task) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *Task) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *Task) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *Task) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


