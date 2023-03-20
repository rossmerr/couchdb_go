# Members

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **[]string** | List of CouchDB user names | [optional] 
**Roles** | Pointer to **[]string** | List of users roles | [optional] 

## Methods

### NewMembers

`func NewMembers() *Members`

NewMembers instantiates a new Members object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembersWithDefaults

`func NewMembersWithDefaults() *Members`

NewMembersWithDefaults instantiates a new Members object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Members) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Members) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Members) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *Members) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *Members) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Members) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Members) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Members) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


