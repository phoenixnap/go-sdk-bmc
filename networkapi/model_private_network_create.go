/*
Networks API

Use the Networks API to create, list, edit, and delete private networks to best fit your business needs. Private networks allow your servers to communicate without connecting to the public internet, avoiding unnecessary egress data charges. </br></br>**Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#multi-private-backend-network-api' target='_blank'>here</a>**</br></br>**All URLs are relative to (https://api.phoenixnap.com/networks/v1/)**

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkapi

import (
	"encoding/json"
)

// PrivateNetworkCreate Details of Private Network to be created.
type PrivateNetworkCreate struct {
	// The friendly name of this private network. This name should be unique.
	Name string `json:"name"`
	// The description of this private network.
	Description *string `json:"description,omitempty"`
	// The location of this private network. Supported values are `PHX`, `ASH`, `SGP`, `NLD`, `CHI`, `SEA` and `AUS`.
	Location string `json:"location"`
	// Identifies network as the default private network for the specified location.
	LocationDefault *bool `json:"locationDefault,omitempty"`
	// IP range associated with this private network in CIDR notation.
	Cidr string `json:"cidr"`
}

// NewPrivateNetworkCreate instantiates a new PrivateNetworkCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateNetworkCreate(name string, location string, cidr string) *PrivateNetworkCreate {
	this := PrivateNetworkCreate{}
	this.Name = name
	this.Location = location
	var locationDefault bool = false
	this.LocationDefault = &locationDefault
	this.Cidr = cidr
	return &this
}

// NewPrivateNetworkCreateWithDefaults instantiates a new PrivateNetworkCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateNetworkCreateWithDefaults() *PrivateNetworkCreate {
	this := PrivateNetworkCreate{}
	var locationDefault bool = false
	this.LocationDefault = &locationDefault
	return &this
}

// GetName returns the Name field value
func (o *PrivateNetworkCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrivateNetworkCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PrivateNetworkCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetworkCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PrivateNetworkCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PrivateNetworkCreate) SetDescription(v string) {
	o.Description = &v
}

// GetLocation returns the Location field value
func (o *PrivateNetworkCreate) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkCreate) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *PrivateNetworkCreate) SetLocation(v string) {
	o.Location = v
}

// GetLocationDefault returns the LocationDefault field value if set, zero value otherwise.
func (o *PrivateNetworkCreate) GetLocationDefault() bool {
	if o == nil || o.LocationDefault == nil {
		var ret bool
		return ret
	}
	return *o.LocationDefault
}

// GetLocationDefaultOk returns a tuple with the LocationDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetworkCreate) GetLocationDefaultOk() (*bool, bool) {
	if o == nil || o.LocationDefault == nil {
		return nil, false
	}
	return o.LocationDefault, true
}

// HasLocationDefault returns a boolean if a field has been set.
func (o *PrivateNetworkCreate) HasLocationDefault() bool {
	if o != nil && o.LocationDefault != nil {
		return true
	}

	return false
}

// SetLocationDefault gets a reference to the given bool and assigns it to the LocationDefault field.
func (o *PrivateNetworkCreate) SetLocationDefault(v bool) {
	o.LocationDefault = &v
}

// GetCidr returns the Cidr field value
func (o *PrivateNetworkCreate) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkCreate) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *PrivateNetworkCreate) SetCidr(v string) {
	o.Cidr = v
}

func (o PrivateNetworkCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if o.LocationDefault != nil {
		toSerialize["locationDefault"] = o.LocationDefault
	}
	if true {
		toSerialize["cidr"] = o.Cidr
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateNetworkCreate struct {
	value *PrivateNetworkCreate
	isSet bool
}

func (v NullablePrivateNetworkCreate) Get() *PrivateNetworkCreate {
	return v.value
}

func (v *NullablePrivateNetworkCreate) Set(val *PrivateNetworkCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateNetworkCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateNetworkCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateNetworkCreate(val *PrivateNetworkCreate) *NullablePrivateNetworkCreate {
	return &NullablePrivateNetworkCreate{value: val, isSet: true}
}

func (v NullablePrivateNetworkCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateNetworkCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
