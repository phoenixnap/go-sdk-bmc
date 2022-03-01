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

// RelinquishIpBlock Object used to determine whether to relinquish ownership of the IP block upon server deletion.
type RelinquishIpBlock struct {
	// Determines whether the IP blocks assigned to the server should be deleted or not.
	DeleteIpBlocks *bool `json:"deleteIpBlocks,omitempty"`
}

// NewRelinquishIpBlock instantiates a new RelinquishIpBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelinquishIpBlock() *RelinquishIpBlock {
	this := RelinquishIpBlock{}
	var deleteIpBlocks bool = false
	this.DeleteIpBlocks = &deleteIpBlocks
	return &this
}

// NewRelinquishIpBlockWithDefaults instantiates a new RelinquishIpBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelinquishIpBlockWithDefaults() *RelinquishIpBlock {
	this := RelinquishIpBlock{}
	var deleteIpBlocks bool = false
	this.DeleteIpBlocks = &deleteIpBlocks
	return &this
}

// GetDeleteIpBlocks returns the DeleteIpBlocks field value if set, zero value otherwise.
func (o *RelinquishIpBlock) GetDeleteIpBlocks() bool {
	if o == nil || o.DeleteIpBlocks == nil {
		var ret bool
		return ret
	}
	return *o.DeleteIpBlocks
}

// GetDeleteIpBlocksOk returns a tuple with the DeleteIpBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelinquishIpBlock) GetDeleteIpBlocksOk() (*bool, bool) {
	if o == nil || o.DeleteIpBlocks == nil {
		return nil, false
	}
	return o.DeleteIpBlocks, true
}

// HasDeleteIpBlocks returns a boolean if a field has been set.
func (o *RelinquishIpBlock) HasDeleteIpBlocks() bool {
	if o != nil && o.DeleteIpBlocks != nil {
		return true
	}

	return false
}

// SetDeleteIpBlocks gets a reference to the given bool and assigns it to the DeleteIpBlocks field.
func (o *RelinquishIpBlock) SetDeleteIpBlocks(v bool) {
	o.DeleteIpBlocks = &v
}

func (o RelinquishIpBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteIpBlocks != nil {
		toSerialize["deleteIpBlocks"] = o.DeleteIpBlocks
	}
	return json.Marshal(toSerialize)
}

type NullableRelinquishIpBlock struct {
	value *RelinquishIpBlock
	isSet bool
}

func (v NullableRelinquishIpBlock) Get() *RelinquishIpBlock {
	return v.value
}

func (v *NullableRelinquishIpBlock) Set(val *RelinquishIpBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableRelinquishIpBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableRelinquishIpBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelinquishIpBlock(val *RelinquishIpBlock) *NullableRelinquishIpBlock {
	return &NullableRelinquishIpBlock{value: val, isSet: true}
}

func (v NullableRelinquishIpBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelinquishIpBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}