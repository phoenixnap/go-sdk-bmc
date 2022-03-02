/*
IP Addresses API

Public IP blocks are a set of contiguous IPs that allow you to access your servers or networks from the internet. Use the IP Addresses API to request and delete IP blocks. <br/></br>**Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/public-ip-management#bmc-public-ip-allocations-api' target='_blank'>here</a>**</br></br>**All URLs are relative to (https://api.phoenixnap.com/ips/v1/)**

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipapi

import (
	"encoding/json"
)

// IpBlockCreate IP Block Request.
type IpBlockCreate struct {
	// IP Block location ID. Currently this field should be set to `PHX`, `ASH`, `SGP`, `NLD`, `CHI` or `SEA`.
	Location string `json:"location"`
	// CIDR IP Block Size. Currently this field should be set to either `/31`, `/30`, `/29` or `/28`.
	CidrBlockSize string `json:"cidrBlockSize"`
}

// NewIpBlockCreate instantiates a new IpBlockCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpBlockCreate(location string, cidrBlockSize string) *IpBlockCreate {
	this := IpBlockCreate{}
	this.Location = location
	this.CidrBlockSize = cidrBlockSize
	return &this
}

// NewIpBlockCreateWithDefaults instantiates a new IpBlockCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpBlockCreateWithDefaults() *IpBlockCreate {
	this := IpBlockCreate{}
	return &this
}

// GetLocation returns the Location field value
func (o *IpBlockCreate) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *IpBlockCreate) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *IpBlockCreate) SetLocation(v string) {
	o.Location = v
}

// GetCidrBlockSize returns the CidrBlockSize field value
func (o *IpBlockCreate) GetCidrBlockSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CidrBlockSize
}

// GetCidrBlockSizeOk returns a tuple with the CidrBlockSize field value
// and a boolean to check if the value has been set.
func (o *IpBlockCreate) GetCidrBlockSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CidrBlockSize, true
}

// SetCidrBlockSize sets field value
func (o *IpBlockCreate) SetCidrBlockSize(v string) {
	o.CidrBlockSize = v
}

func (o IpBlockCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["cidrBlockSize"] = o.CidrBlockSize
	}
	return json.Marshal(toSerialize)
}

type NullableIpBlockCreate struct {
	value *IpBlockCreate
	isSet bool
}

func (v NullableIpBlockCreate) Get() *IpBlockCreate {
	return v.value
}

func (v *NullableIpBlockCreate) Set(val *IpBlockCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableIpBlockCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableIpBlockCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpBlockCreate(val *IpBlockCreate) *NullableIpBlockCreate {
	return &NullableIpBlockCreate{value: val, isSet: true}
}

func (v NullableIpBlockCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpBlockCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
