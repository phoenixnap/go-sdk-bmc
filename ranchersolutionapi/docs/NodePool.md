# NodePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the node pool. | [optional] 
**NodeCount** | Pointer to **int32** | Number of configured nodes, currently only node counts of 1 and 3 are possible. | [optional] 
**ServerType** | Pointer to **string** | Node server type. Cannot be changed once a server is created. Currently this field should be set to either &#x60;s0.d1.small&#x60;, &#x60;s0.d1.medium&#x60;, &#x60;s1.c1.small&#x60;, &#x60;s1.c1.medium&#x60;, &#x60;s1.c2.medium&#x60;, &#x60;s1.c2.large&#x60;, &#x60;s2.c1.small&#x60;, &#x60;s2.c1.medium&#x60;, &#x60;s2.c1.large&#x60;, &#x60;s2.c2.small&#x60;, &#x60;s2.c2.medium&#x60;, &#x60;s2.c2.large&#x60;, &#x60;s1.e1.small&#x60;, &#x60;s1.e1.medium&#x60;, &#x60;s1.e1.large&#x60;. | [optional] [default to "s0.d1.small"]
**SshConfig** | Pointer to [**SshConfig**](SshConfig.md) |  | [optional] 
**Nodes** | Pointer to [**[]Node**](Node.md) | (Read-only) The nodes associated with this node pool. | [optional] [readonly] 

## Methods

### NewNodePool

`func NewNodePool() *NodePool`

NewNodePool instantiates a new NodePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolWithDefaults

`func NewNodePoolWithDefaults() *NodePool`

NewNodePoolWithDefaults instantiates a new NodePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodePool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodePool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *NodePool) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *NodePool) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *NodePool) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *NodePool) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetServerType

`func (o *NodePool) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *NodePool) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *NodePool) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *NodePool) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetSshConfig

`func (o *NodePool) GetSshConfig() SshConfig`

GetSshConfig returns the SshConfig field if non-nil, zero value otherwise.

### GetSshConfigOk

`func (o *NodePool) GetSshConfigOk() (*SshConfig, bool)`

GetSshConfigOk returns a tuple with the SshConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshConfig

`func (o *NodePool) SetSshConfig(v SshConfig)`

SetSshConfig sets SshConfig field to given value.

### HasSshConfig

`func (o *NodePool) HasSshConfig() bool`

HasSshConfig returns a boolean if a field has been set.

### SetSshConfigNil

`func (o *NodePool) SetSshConfigNil(b bool)`

 SetSshConfigNil sets the value for SshConfig to be an explicit nil

### UnsetSshConfig
`func (o *NodePool) UnsetSshConfig()`

UnsetSshConfig ensures that no value is present for SshConfig, not even an explicit nil
### GetNodes

`func (o *NodePool) GetNodes() []Node`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *NodePool) GetNodesOk() (*[]Node, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *NodePool) SetNodes(v []Node)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *NodePool) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


