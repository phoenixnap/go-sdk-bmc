/*
Bare Metal Cloud API

Create, power off, power on, reset, reboot, or shut down your server with the Bare Metal Cloud API. Deprovision servers, get or edit SSH key details, and a lot more. Manage your infrastructure more efficiently using just a few simple api calls. <br/></br>**Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/how-to-deploy-bare-metal-cloud-server' target='_blank'>here</a>**</br></br>**All URLs are relative to (https://api.phoenixnap.com/bmc/v1/)**

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmcapi

import (
	"encoding/json"
)

// ServerIpBlock IP block ID assigned to server
type ServerIpBlock struct {
	// The IP block's ID.
	Id string `json:"id"`
	// The VLAN on which this IP block has been configured within the network switch.
	VlanId *int32 `json:"vlanId,omitempty"`
}

// NewServerIpBlock instantiates a new ServerIpBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerIpBlock(id string) *ServerIpBlock {
	this := ServerIpBlock{}
	this.Id = id
	return &this
}

// NewServerIpBlockWithDefaults instantiates a new ServerIpBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerIpBlockWithDefaults() *ServerIpBlock {
	this := ServerIpBlock{}
	return &this
}

// GetId returns the Id field value
func (o *ServerIpBlock) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerIpBlock) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerIpBlock) SetId(v string) {
	o.Id = v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *ServerIpBlock) GetVlanId() int32 {
	if o == nil || o.VlanId == nil {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerIpBlock) GetVlanIdOk() (*int32, bool) {
	if o == nil || o.VlanId == nil {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *ServerIpBlock) HasVlanId() bool {
	if o != nil && o.VlanId != nil {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *ServerIpBlock) SetVlanId(v int32) {
	o.VlanId = &v
}

func (o ServerIpBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.VlanId != nil {
		toSerialize["vlanId"] = o.VlanId
	}
	return json.Marshal(toSerialize)
}

type NullableServerIpBlock struct {
	value *ServerIpBlock
	isSet bool
}

func (v NullableServerIpBlock) Get() *ServerIpBlock {
	return v.value
}

func (v *NullableServerIpBlock) Set(val *ServerIpBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableServerIpBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableServerIpBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerIpBlock(val *ServerIpBlock) *NullableServerIpBlock {
	return &NullableServerIpBlock{value: val, isSet: true}
}

func (v NullableServerIpBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerIpBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}