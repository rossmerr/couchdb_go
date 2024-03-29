/*
CouchDB API

*Note* This is not a definitive implementation of the CouchDB API, it's missing a lot of the endpoints for server/database managment and everything for attachments all COPY operations and probably a few other parts.  It also targets golang, as such the '#/definitions/Document' is intentionally left empty to generate a go `interface`, which you can then cast to a `map[string]interface{}`. 

API version: 4.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package couchdb_go

import (
	"encoding/json"
)

// checks if the Task type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Task{}

// Task struct for Task
type Task struct {
	// Processed changes
	ChangesDone *int32 `json:"changes_done,omitempty"`
	// Source database
	Database *string `json:"database,omitempty"`
	// Process ID
	Pid *string `json:"pid,omitempty"`
	// Current percentage progress
	Progress *int32 `json:"progress,omitempty"`
	// ask start time as unix timestamp
	StartedOn *int32 `json:"started_on,omitempty"`
	// Task status message
	Status *string `json:"status,omitempty"`
	// Task name
	Task *string `json:"task,omitempty"`
	// Total changes to process
	TotalChanges *int32 `json:"total_changes,omitempty"`
	// Operation Type
	Type *string `json:"type,omitempty"`
	// Unix timestamp of last operation update
	UpdatedOn *string `json:"updated_on,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Task Task

// NewTask instantiates a new Task object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTask() *Task {
	this := Task{}
	return &this
}

// NewTaskWithDefaults instantiates a new Task object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetChangesDone returns the ChangesDone field value if set, zero value otherwise.
func (o *Task) GetChangesDone() int32 {
	if o == nil || IsNil(o.ChangesDone) {
		var ret int32
		return ret
	}
	return *o.ChangesDone
}

// GetChangesDoneOk returns a tuple with the ChangesDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetChangesDoneOk() (*int32, bool) {
	if o == nil || IsNil(o.ChangesDone) {
		return nil, false
	}
	return o.ChangesDone, true
}

// HasChangesDone returns a boolean if a field has been set.
func (o *Task) HasChangesDone() bool {
	if o != nil && !IsNil(o.ChangesDone) {
		return true
	}

	return false
}

// SetChangesDone gets a reference to the given int32 and assigns it to the ChangesDone field.
func (o *Task) SetChangesDone(v int32) {
	o.ChangesDone = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *Task) GetDatabase() string {
	if o == nil || IsNil(o.Database) {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetDatabaseOk() (*string, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *Task) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *Task) SetDatabase(v string) {
	o.Database = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *Task) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *Task) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *Task) SetPid(v string) {
	o.Pid = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *Task) GetProgress() int32 {
	if o == nil || IsNil(o.Progress) {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetProgressOk() (*int32, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *Task) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *Task) SetProgress(v int32) {
	o.Progress = &v
}

// GetStartedOn returns the StartedOn field value if set, zero value otherwise.
func (o *Task) GetStartedOn() int32 {
	if o == nil || IsNil(o.StartedOn) {
		var ret int32
		return ret
	}
	return *o.StartedOn
}

// GetStartedOnOk returns a tuple with the StartedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetStartedOnOk() (*int32, bool) {
	if o == nil || IsNil(o.StartedOn) {
		return nil, false
	}
	return o.StartedOn, true
}

// HasStartedOn returns a boolean if a field has been set.
func (o *Task) HasStartedOn() bool {
	if o != nil && !IsNil(o.StartedOn) {
		return true
	}

	return false
}

// SetStartedOn gets a reference to the given int32 and assigns it to the StartedOn field.
func (o *Task) SetStartedOn(v int32) {
	o.StartedOn = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Task) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Task) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Task) SetStatus(v string) {
	o.Status = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *Task) GetTask() string {
	if o == nil || IsNil(o.Task) {
		var ret string
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTaskOk() (*string, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *Task) HasTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given string and assigns it to the Task field.
func (o *Task) SetTask(v string) {
	o.Task = &v
}

// GetTotalChanges returns the TotalChanges field value if set, zero value otherwise.
func (o *Task) GetTotalChanges() int32 {
	if o == nil || IsNil(o.TotalChanges) {
		var ret int32
		return ret
	}
	return *o.TotalChanges
}

// GetTotalChangesOk returns a tuple with the TotalChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTotalChangesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalChanges) {
		return nil, false
	}
	return o.TotalChanges, true
}

// HasTotalChanges returns a boolean if a field has been set.
func (o *Task) HasTotalChanges() bool {
	if o != nil && !IsNil(o.TotalChanges) {
		return true
	}

	return false
}

// SetTotalChanges gets a reference to the given int32 and assigns it to the TotalChanges field.
func (o *Task) SetTotalChanges(v int32) {
	o.TotalChanges = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Task) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Task) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Task) SetType(v string) {
	o.Type = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *Task) GetUpdatedOn() string {
	if o == nil || IsNil(o.UpdatedOn) {
		var ret string
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetUpdatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedOn) {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *Task) HasUpdatedOn() bool {
	if o != nil && !IsNil(o.UpdatedOn) {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given string and assigns it to the UpdatedOn field.
func (o *Task) SetUpdatedOn(v string) {
	o.UpdatedOn = &v
}

func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Task) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChangesDone) {
		toSerialize["changes_done"] = o.ChangesDone
	}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !IsNil(o.StartedOn) {
		toSerialize["started_on"] = o.StartedOn
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}
	if !IsNil(o.TotalChanges) {
		toSerialize["total_changes"] = o.TotalChanges
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedOn) {
		toSerialize["updated_on"] = o.UpdatedOn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Task) UnmarshalJSON(bytes []byte) (err error) {
	varTask := _Task{}

	if err = json.Unmarshal(bytes, &varTask); err == nil {
		*o = Task(varTask)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "changes_done")
		delete(additionalProperties, "database")
		delete(additionalProperties, "pid")
		delete(additionalProperties, "progress")
		delete(additionalProperties, "started_on")
		delete(additionalProperties, "status")
		delete(additionalProperties, "task")
		delete(additionalProperties, "total_changes")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updated_on")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTask struct {
	value *Task
	isSet bool
}

func (v NullableTask) Get() *Task {
	return v.value
}

func (v *NullableTask) Set(val *Task) {
	v.value = val
	v.isSet = true
}

func (v NullableTask) IsSet() bool {
	return v.isSet
}

func (v *NullableTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTask(val *Task) *NullableTask {
	return &NullableTask{value: val, isSet: true}
}

func (v NullableTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


