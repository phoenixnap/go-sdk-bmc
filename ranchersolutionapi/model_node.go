/*
Rancher Solution API

Simplify enterprise-grade Kubernetes cluster operations and management with Rancher on Bare Metal Cloud. Deploy Kubernetes clusters using a few API calls.<br> <br> <span class=\"pnap-api-knowledge-base-link\"> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/rancher-bmc-integration-kubernetes' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/solutions/rancher/v1beta)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ranchersolutionapi

import (
	"encoding/json"
)

// Node Node details.
type Node struct {
	// The server identifier.
	ServerId *string `json:"serverId,omitempty"`
}

// NewNode instantiates a new Node object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNode() *Node {
	this := Node{}
	return &this
}

// NewNodeWithDefaults instantiates a new Node object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeWithDefaults() *Node {
	this := Node{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *Node) GetServerId() string {
	if o == nil || o.ServerId == nil {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetServerIdOk() (*string, bool) {
	if o == nil || o.ServerId == nil {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *Node) HasServerId() bool {
	if o != nil && o.ServerId != nil {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *Node) SetServerId(v string) {
	o.ServerId = &v
}

func (o Node) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServerId != nil {
		toSerialize["serverId"] = o.ServerId
	}
	return json.Marshal(toSerialize)
}

type NullableNode struct {
	value *Node
	isSet bool
}

func (v NullableNode) Get() *Node {
	return v.value
}

func (v *NullableNode) Set(val *Node) {
	v.value = val
	v.isSet = true
}

func (v NullableNode) IsSet() bool {
	return v.isSet
}

func (v *NullableNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNode(val *Node) *NullableNode {
	return &NullableNode{value: val, isSet: true}
}

func (v NullableNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
