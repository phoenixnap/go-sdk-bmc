/*
Networks API

Use the Networks API to create, list, edit, and delete private networks to best fit your business needs. Private networks allow your servers to communicate without connecting to the public internet, avoiding unnecessary egress data charges.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#multi-private-backend-network-api' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/networks/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkapi

import (
	"encoding/json"
)

// PrivateNetworkModify Object including details to be modified in the Private Network.
type PrivateNetworkModify struct {
	// A friendly name given to the private network. This name should be unique.
	Name string `json:"name"`
	// A description of this private network
	Description *string `json:"description,omitempty"`
	// Identifies network as the default private network for the specified location.
	LocationDefault bool `json:"locationDefault"`
}

// NewPrivateNetworkModify instantiates a new PrivateNetworkModify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateNetworkModify(name string, locationDefault bool) *PrivateNetworkModify {
	this := PrivateNetworkModify{}
	this.Name = name
	this.LocationDefault = locationDefault
	return &this
}

// NewPrivateNetworkModifyWithDefaults instantiates a new PrivateNetworkModify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateNetworkModifyWithDefaults() *PrivateNetworkModify {
	this := PrivateNetworkModify{}
	return &this
}

// GetName returns the Name field value
func (o *PrivateNetworkModify) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkModify) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrivateNetworkModify) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PrivateNetworkModify) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetworkModify) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PrivateNetworkModify) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PrivateNetworkModify) SetDescription(v string) {
	o.Description = &v
}

// GetLocationDefault returns the LocationDefault field value
func (o *PrivateNetworkModify) GetLocationDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LocationDefault
}

// GetLocationDefaultOk returns a tuple with the LocationDefault field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkModify) GetLocationDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationDefault, true
}

// SetLocationDefault sets field value
func (o *PrivateNetworkModify) SetLocationDefault(v bool) {
	o.LocationDefault = v
}

func (o PrivateNetworkModify) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["locationDefault"] = o.LocationDefault
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateNetworkModify struct {
	value *PrivateNetworkModify
	isSet bool
}

func (v NullablePrivateNetworkModify) Get() *PrivateNetworkModify {
	return v.value
}

func (v *NullablePrivateNetworkModify) Set(val *PrivateNetworkModify) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateNetworkModify) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateNetworkModify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateNetworkModify(val *PrivateNetworkModify) *NullablePrivateNetworkModify {
	return &NullablePrivateNetworkModify{value: val, isSet: true}
}

func (v NullablePrivateNetworkModify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateNetworkModify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
