/*
Rancher Solution API

Simplify enterprise-grade Kubernetes cluster operations and management with Rancher on Bare Metal Cloud. Deploy Kubernetes clusters using a few API calls.  <span id=\"pnap-api-knowledge-base-link\"> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/rancher-bmc-integration-kubernetes' target='_blank'>here</a> </span>  **All URLs are relative to (https://api.phoenixnap.com/solutions/rancher/v1beta)**

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ranchersolutionapi

import (
	"encoding/json"
)

// NodePool Node Pool Configuration. A node pool contains the name and configuration for a cluster's node pool. Node pools are set of nodes with a common configuration and specification.
type NodePool struct {
	// The name of the node pool.
	Name *string `json:"name,omitempty"`
	// Number of configured nodes, currently only node counts of 1 and 3 are possible.
	NodeCount *int32 `json:"nodeCount,omitempty"`
	// Node server type.
	ServerType *string    `json:"serverType,omitempty"`
	SshConfig  *SshConfig `json:"sshConfig,omitempty"`
	// (Read-only) The nodes associated with this node pool.
	Nodes *[]Node `json:"nodes,omitempty"`
}

// NewNodePool instantiates a new NodePool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodePool() *NodePool {
	this := NodePool{}
	var serverType string = "s0.d1.small"
	this.ServerType = &serverType
	return &this
}

// NewNodePoolWithDefaults instantiates a new NodePool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodePoolWithDefaults() *NodePool {
	this := NodePool{}
	var serverType string = "s0.d1.small"
	this.ServerType = &serverType
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NodePool) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePool) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NodePool) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NodePool) SetName(v string) {
	o.Name = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *NodePool) GetNodeCount() int32 {
	if o == nil || o.NodeCount == nil {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePool) GetNodeCountOk() (*int32, bool) {
	if o == nil || o.NodeCount == nil {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *NodePool) HasNodeCount() bool {
	if o != nil && o.NodeCount != nil {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *NodePool) SetNodeCount(v int32) {
	o.NodeCount = &v
}

// GetServerType returns the ServerType field value if set, zero value otherwise.
func (o *NodePool) GetServerType() string {
	if o == nil || o.ServerType == nil {
		var ret string
		return ret
	}
	return *o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePool) GetServerTypeOk() (*string, bool) {
	if o == nil || o.ServerType == nil {
		return nil, false
	}
	return o.ServerType, true
}

// HasServerType returns a boolean if a field has been set.
func (o *NodePool) HasServerType() bool {
	if o != nil && o.ServerType != nil {
		return true
	}

	return false
}

// SetServerType gets a reference to the given string and assigns it to the ServerType field.
func (o *NodePool) SetServerType(v string) {
	o.ServerType = &v
}

// GetSshConfig returns the SshConfig field value if set, zero value otherwise.
func (o *NodePool) GetSshConfig() SshConfig {
	if o == nil || o.SshConfig == nil {
		var ret SshConfig
		return ret
	}
	return *o.SshConfig
}

// GetSshConfigOk returns a tuple with the SshConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePool) GetSshConfigOk() (*SshConfig, bool) {
	if o == nil || o.SshConfig == nil {
		return nil, false
	}
	return o.SshConfig, true
}

// HasSshConfig returns a boolean if a field has been set.
func (o *NodePool) HasSshConfig() bool {
	if o != nil && o.SshConfig != nil {
		return true
	}

	return false
}

// SetSshConfig gets a reference to the given SshConfig and assigns it to the SshConfig field.
func (o *NodePool) SetSshConfig(v SshConfig) {
	o.SshConfig = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *NodePool) GetNodes() []Node {
	if o == nil || o.Nodes == nil {
		var ret []Node
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePool) GetNodesOk() (*[]Node, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *NodePool) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []Node and assigns it to the Nodes field.
func (o *NodePool) SetNodes(v []Node) {
	o.Nodes = &v
}

func (o NodePool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NodeCount != nil {
		toSerialize["nodeCount"] = o.NodeCount
	}
	if o.ServerType != nil {
		toSerialize["serverType"] = o.ServerType
	}
	if o.SshConfig != nil {
		toSerialize["sshConfig"] = o.SshConfig
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	return json.Marshal(toSerialize)
}

type NullableNodePool struct {
	value *NodePool
	isSet bool
}

func (v NullableNodePool) Get() *NodePool {
	return v.value
}

func (v *NullableNodePool) Set(val *NodePool) {
	v.value = val
	v.isSet = true
}

func (v NullableNodePool) IsSet() bool {
	return v.isSet
}

func (v *NullableNodePool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodePool(val *NodePool) *NullableNodePool {
	return &NullableNodePool{value: val, isSet: true}
}

func (v NullableNodePool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodePool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
