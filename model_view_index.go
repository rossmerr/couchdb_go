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

// checks if the ViewIndex type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViewIndex{}

// ViewIndex struct for ViewIndex
type ViewIndex struct {
	// Indicates whether a compaction routine is currently running on the view
	CompactRunning *bool `json:"compact_running,omitempty"`
	// Language for the defined views
	Language *string `json:"language,omitempty"`
	// The purge sequence that has been processed
	PurgeSeq *int32 `json:"purge_seq,omitempty"`
	// MD5 signature of the views for the design document
	Signature *string `json:"signature,omitempty"`
	Sizes *ViewIndexSizes `json:"sizes,omitempty"`
	// The update sequence of the corresponding database that has been indexed
	UpdateSeq *string `json:"update_seq,omitempty"`
	// Indicates if the view is currently being updated
	UpdaterRunning *bool `json:"updater_running,omitempty"`
	// Number of clients waiting on views from this design document
	WaitingClients *int32 `json:"waiting_clients,omitempty"`
	// Indicates if there are outstanding commits to the underlying database that need to processed
	WaitingCommit *bool `json:"waiting_commit,omitempty"`
}

// NewViewIndex instantiates a new ViewIndex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewIndex() *ViewIndex {
	this := ViewIndex{}
	return &this
}

// NewViewIndexWithDefaults instantiates a new ViewIndex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewIndexWithDefaults() *ViewIndex {
	this := ViewIndex{}
	return &this
}

// GetCompactRunning returns the CompactRunning field value if set, zero value otherwise.
func (o *ViewIndex) GetCompactRunning() bool {
	if o == nil || IsNil(o.CompactRunning) {
		var ret bool
		return ret
	}
	return *o.CompactRunning
}

// GetCompactRunningOk returns a tuple with the CompactRunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewIndex) GetCompactRunningOk() (*bool, bool) {
	if o == nil || IsNil(o.CompactRunning) {
		return nil, false
	}
	return o.CompactRunning, true
}

// HasCompactRunning returns a boolean if a field has been set.
func (o *ViewIndex) HasCompactRunning() bool {
	if o != nil && !IsNil(o.CompactRunning) {
		return true
	}

	return false
}

// SetCompactRunning gets a reference to the given bool and assigns it to the CompactRunning field.
func (o *ViewIndex) SetCompactRunning(v bool) {
	o.CompactRunning = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ViewIndex) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewIndex) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ViewIndex) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ViewIndex) SetLanguage(v string) {
	o.Language = &v
}

// GetPurgeSeq returns the PurgeSeq field value if set, zero value otherwise.
func (o *ViewIndex) GetPurgeSeq() int32 {
	if o == nil || IsNil(o.PurgeSeq) {
		var ret int32
		return ret
	}
	return *o.PurgeSeq
}

// GetPurgeSeqOk returns a tuple with the PurgeSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewIndex) GetPurgeSeqOk() (*int32, bool) {
	if o == nil || IsNil(o.PurgeSeq) {
		return nil, false
	}
	return o.PurgeSeq, true
}

// HasPurgeSeq returns a boolean if a field has been set.
func (o *ViewIndex) HasPurgeSeq() bool {
	if o != nil && !IsNil(o.PurgeSeq) {
		return true
	}

	return false
}

// SetPurgeSeq gets a reference to the given int32 and assigns it to the PurgeSeq field.
func (o *ViewIndex) SetPurgeSeq(v int32) {
	o.PurgeSeq = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *ViewIndex) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewIndex) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *ViewIndex) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *ViewIndex) SetSignature(v string) {
	o.Signature = &v
}

// GetSizes returns the Sizes field value if set, zero value otherwise.
func (o *ViewIndex) GetSizes() ViewIndexSizes {
	if o == nil || IsNil(o.Sizes) {
		var ret ViewIndexSizes
		return ret
	}
	return *o.Sizes
}

// GetSizesOk returns a tuple with the Sizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewIndex) GetSizesOk() (*ViewIndexSizes, bool) {
	if o == nil || IsNil(o.Sizes) {
		return nil, false
	}
	return o.Sizes, true
}

// HasSizes returns a boolean if a field has been set.
func (o *ViewIndex) HasSizes() bool {
	if o != nil && !IsNil(o.Sizes) {
		return true
	}

	return false
}

// SetSizes gets a reference to the given ViewIndexSizes and assigns it to the Sizes field.
func (o *ViewIndex) SetSizes(v ViewIndexSizes) {
	o.Sizes = &v
}

// GetUpdateSeq returns the UpdateSeq field value if set, zero value otherwise.
func (o *ViewIndex) GetUpdateSeq() string {
	if o == nil || IsNil(o.UpdateSeq) {
		var ret string
		return ret
	}
	return *o.UpdateSeq
}

// GetUpdateSeqOk returns a tuple with the UpdateSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewIndex) GetUpdateSeqOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateSeq) {
		return nil, false
	}
	return o.UpdateSeq, true
}

// HasUpdateSeq returns a boolean if a field has been set.
func (o *ViewIndex) HasUpdateSeq() bool {
	if o != nil && !IsNil(o.UpdateSeq) {
		return true
	}

	return false
}

// SetUpdateSeq gets a reference to the given string and assigns it to the UpdateSeq field.
func (o *ViewIndex) SetUpdateSeq(v string) {
	o.UpdateSeq = &v
}

// GetUpdaterRunning returns the UpdaterRunning field value if set, zero value otherwise.
func (o *ViewIndex) GetUpdaterRunning() bool {
	if o == nil || IsNil(o.UpdaterRunning) {
		var ret bool
		return ret
	}
	return *o.UpdaterRunning
}

// GetUpdaterRunningOk returns a tuple with the UpdaterRunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewIndex) GetUpdaterRunningOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdaterRunning) {
		return nil, false
	}
	return o.UpdaterRunning, true
}

// HasUpdaterRunning returns a boolean if a field has been set.
func (o *ViewIndex) HasUpdaterRunning() bool {
	if o != nil && !IsNil(o.UpdaterRunning) {
		return true
	}

	return false
}

// SetUpdaterRunning gets a reference to the given bool and assigns it to the UpdaterRunning field.
func (o *ViewIndex) SetUpdaterRunning(v bool) {
	o.UpdaterRunning = &v
}

// GetWaitingClients returns the WaitingClients field value if set, zero value otherwise.
func (o *ViewIndex) GetWaitingClients() int32 {
	if o == nil || IsNil(o.WaitingClients) {
		var ret int32
		return ret
	}
	return *o.WaitingClients
}

// GetWaitingClientsOk returns a tuple with the WaitingClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewIndex) GetWaitingClientsOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitingClients) {
		return nil, false
	}
	return o.WaitingClients, true
}

// HasWaitingClients returns a boolean if a field has been set.
func (o *ViewIndex) HasWaitingClients() bool {
	if o != nil && !IsNil(o.WaitingClients) {
		return true
	}

	return false
}

// SetWaitingClients gets a reference to the given int32 and assigns it to the WaitingClients field.
func (o *ViewIndex) SetWaitingClients(v int32) {
	o.WaitingClients = &v
}

// GetWaitingCommit returns the WaitingCommit field value if set, zero value otherwise.
func (o *ViewIndex) GetWaitingCommit() bool {
	if o == nil || IsNil(o.WaitingCommit) {
		var ret bool
		return ret
	}
	return *o.WaitingCommit
}

// GetWaitingCommitOk returns a tuple with the WaitingCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewIndex) GetWaitingCommitOk() (*bool, bool) {
	if o == nil || IsNil(o.WaitingCommit) {
		return nil, false
	}
	return o.WaitingCommit, true
}

// HasWaitingCommit returns a boolean if a field has been set.
func (o *ViewIndex) HasWaitingCommit() bool {
	if o != nil && !IsNil(o.WaitingCommit) {
		return true
	}

	return false
}

// SetWaitingCommit gets a reference to the given bool and assigns it to the WaitingCommit field.
func (o *ViewIndex) SetWaitingCommit(v bool) {
	o.WaitingCommit = &v
}

func (o ViewIndex) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViewIndex) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompactRunning) {
		toSerialize["compact_running"] = o.CompactRunning
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.PurgeSeq) {
		toSerialize["purge_seq"] = o.PurgeSeq
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !IsNil(o.Sizes) {
		toSerialize["sizes"] = o.Sizes
	}
	if !IsNil(o.UpdateSeq) {
		toSerialize["update_seq"] = o.UpdateSeq
	}
	if !IsNil(o.UpdaterRunning) {
		toSerialize["updater_running"] = o.UpdaterRunning
	}
	if !IsNil(o.WaitingClients) {
		toSerialize["waiting_clients"] = o.WaitingClients
	}
	if !IsNil(o.WaitingCommit) {
		toSerialize["waiting_commit"] = o.WaitingCommit
	}
	return toSerialize, nil
}

type NullableViewIndex struct {
	value *ViewIndex
	isSet bool
}

func (v NullableViewIndex) Get() *ViewIndex {
	return v.value
}

func (v *NullableViewIndex) Set(val *ViewIndex) {
	v.value = val
	v.isSet = true
}

func (v NullableViewIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableViewIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewIndex(val *ViewIndex) *NullableViewIndex {
	return &NullableViewIndex{value: val, isSet: true}
}

func (v NullableViewIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

