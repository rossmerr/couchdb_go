# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**BindAddress** | Pointer to **string** | The IP address to which to bind the current node. The special value 0.0.0.0 may be specified to bind to all interfaces on the host. (enable_cluster and enable_single_node only) | [optional] 
**Username** | Pointer to **string** | The username of the server-level administrator to create. (enable_cluster and enable_single_node only), or the remote server’s administrator username (add_node) | [optional] 
**Password** | Pointer to **string** | The password for the server-level administrator to create. (enable_cluster and enable_single_node only), or the remote server’s administrator username (add_node) | [optional] 
**Port** | Pointer to **int32** | The TCP port to which to bind this node (enable_cluster and enable_single_node only) or the TCP port to which to bind a remote node (add_node only). | [optional] 
**NodeCount** | Pointer to **int32** | The total number of nodes to be joined into the cluster, including this one. Used to determine the value of the cluster’s n, up to a maximum of 3. (enable_cluster only) | [optional] 
**RemoteNode** | Pointer to **string** | The IP address of the remote node to setup as part of this cluster’s list of nodes. (enable_cluster only) | [optional] 
**RemoteCurrentUser** | Pointer to **string** | The username of the server-level administrator authorized on the remote node. (enable_cluster only) | [optional] 
**RemoteCurrentPassword** | Pointer to **string** | he password of the server-level administrator authorized on the remote node. (enable_cluster only) | [optional] 
**Host** | Pointer to **string** | The remote node IP of the node to add to the cluster. (add_node only) | [optional] 
**EnsureDbsExist** | Pointer to **[]string** | List of system databases to ensure exist on the node/cluster. Defaults to [\&quot;_users\&quot;,\&quot;_replicator\&quot;]. | [optional] 

## Methods

### NewCluster

`func NewCluster() *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Cluster) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Cluster) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Cluster) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Cluster) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBindAddress

`func (o *Cluster) GetBindAddress() string`

GetBindAddress returns the BindAddress field if non-nil, zero value otherwise.

### GetBindAddressOk

`func (o *Cluster) GetBindAddressOk() (*string, bool)`

GetBindAddressOk returns a tuple with the BindAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindAddress

`func (o *Cluster) SetBindAddress(v string)`

SetBindAddress sets BindAddress field to given value.

### HasBindAddress

`func (o *Cluster) HasBindAddress() bool`

HasBindAddress returns a boolean if a field has been set.

### GetUsername

`func (o *Cluster) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Cluster) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Cluster) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Cluster) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *Cluster) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Cluster) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Cluster) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Cluster) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *Cluster) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Cluster) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Cluster) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Cluster) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetNodeCount

`func (o *Cluster) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *Cluster) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *Cluster) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *Cluster) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetRemoteNode

`func (o *Cluster) GetRemoteNode() string`

GetRemoteNode returns the RemoteNode field if non-nil, zero value otherwise.

### GetRemoteNodeOk

`func (o *Cluster) GetRemoteNodeOk() (*string, bool)`

GetRemoteNodeOk returns a tuple with the RemoteNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNode

`func (o *Cluster) SetRemoteNode(v string)`

SetRemoteNode sets RemoteNode field to given value.

### HasRemoteNode

`func (o *Cluster) HasRemoteNode() bool`

HasRemoteNode returns a boolean if a field has been set.

### GetRemoteCurrentUser

`func (o *Cluster) GetRemoteCurrentUser() string`

GetRemoteCurrentUser returns the RemoteCurrentUser field if non-nil, zero value otherwise.

### GetRemoteCurrentUserOk

`func (o *Cluster) GetRemoteCurrentUserOk() (*string, bool)`

GetRemoteCurrentUserOk returns a tuple with the RemoteCurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCurrentUser

`func (o *Cluster) SetRemoteCurrentUser(v string)`

SetRemoteCurrentUser sets RemoteCurrentUser field to given value.

### HasRemoteCurrentUser

`func (o *Cluster) HasRemoteCurrentUser() bool`

HasRemoteCurrentUser returns a boolean if a field has been set.

### GetRemoteCurrentPassword

`func (o *Cluster) GetRemoteCurrentPassword() string`

GetRemoteCurrentPassword returns the RemoteCurrentPassword field if non-nil, zero value otherwise.

### GetRemoteCurrentPasswordOk

`func (o *Cluster) GetRemoteCurrentPasswordOk() (*string, bool)`

GetRemoteCurrentPasswordOk returns a tuple with the RemoteCurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCurrentPassword

`func (o *Cluster) SetRemoteCurrentPassword(v string)`

SetRemoteCurrentPassword sets RemoteCurrentPassword field to given value.

### HasRemoteCurrentPassword

`func (o *Cluster) HasRemoteCurrentPassword() bool`

HasRemoteCurrentPassword returns a boolean if a field has been set.

### GetHost

`func (o *Cluster) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Cluster) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Cluster) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Cluster) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetEnsureDbsExist

`func (o *Cluster) GetEnsureDbsExist() []string`

GetEnsureDbsExist returns the EnsureDbsExist field if non-nil, zero value otherwise.

### GetEnsureDbsExistOk

`func (o *Cluster) GetEnsureDbsExistOk() (*[]string, bool)`

GetEnsureDbsExistOk returns a tuple with the EnsureDbsExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsureDbsExist

`func (o *Cluster) SetEnsureDbsExist(v []string)`

SetEnsureDbsExist sets EnsureDbsExist field to given value.

### HasEnsureDbsExist

`func (o *Cluster) HasEnsureDbsExist() bool`

HasEnsureDbsExist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


