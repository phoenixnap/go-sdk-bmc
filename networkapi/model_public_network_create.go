/*
Networks API

Create, list, edit and delete public/private networks with the Network API. Use public networks to place multiple  servers on the same network or VLAN. Assign new servers with IP addresses from the same CIDR range. Use private  networks to avoid unnecessary egress data charges. Model your networks according to your business needs.<br> <br> <span class='pnap-api-knowledge-base-link'> Helpful knowledge base articles are available for  <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#multi-private-backend-network-api' target='_blank'>multi-private backend networks</a> and <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#ftoc-heading-15' target='_blank'>public networks</a>. </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/networks/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkapi

import (
	"encoding/json"
	"fmt"
)

// checks if the PublicNetworkCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicNetworkCreate{}

// PublicNetworkCreate Details of Public Network to be created.
type PublicNetworkCreate struct {
	// The friendly name of this public network. This name should be unique.
	Name string `json:"name"`
	// The description of this public network.
	Description *string `json:"description,omitempty"`
	// The location of this public network. Supported values are `PHX`, `ASH`, `SGP`, `NLD`, `CHI`, `SEA` and `AUS`.
	Location string `json:"location"`
	// The VLAN that will be assigned to this network.
	VlanId *int32 `json:"vlanId,omitempty"`
	// A list of IP Blocks that will be associated with this public network.
	IpBlocks             []PublicNetworkIpBlock `json:"ipBlocks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PublicNetworkCreate PublicNetworkCreate

// NewPublicNetworkCreate instantiates a new PublicNetworkCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicNetworkCreate(name string, location string) *PublicNetworkCreate {
	this := PublicNetworkCreate{}
	this.Name = name
	this.Location = location
	return &this
}

// NewPublicNetworkCreateWithDefaults instantiates a new PublicNetworkCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicNetworkCreateWithDefaults() *PublicNetworkCreate {
	this := PublicNetworkCreate{}
	return &this
}

// GetName returns the Name field value
func (o *PublicNetworkCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublicNetworkCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PublicNetworkCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PublicNetworkCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicNetworkCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PublicNetworkCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PublicNetworkCreate) SetDescription(v string) {
	o.Description = &v
}

// GetLocation returns the Location field value
func (o *PublicNetworkCreate) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *PublicNetworkCreate) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *PublicNetworkCreate) SetLocation(v string) {
	o.Location = v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *PublicNetworkCreate) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicNetworkCreate) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *PublicNetworkCreate) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *PublicNetworkCreate) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetIpBlocks returns the IpBlocks field value if set, zero value otherwise.
func (o *PublicNetworkCreate) GetIpBlocks() []PublicNetworkIpBlock {
	if o == nil || IsNil(o.IpBlocks) {
		var ret []PublicNetworkIpBlock
		return ret
	}
	return o.IpBlocks
}

// GetIpBlocksOk returns a tuple with the IpBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicNetworkCreate) GetIpBlocksOk() ([]PublicNetworkIpBlock, bool) {
	if o == nil || IsNil(o.IpBlocks) {
		return nil, false
	}
	return o.IpBlocks, true
}

// HasIpBlocks returns a boolean if a field has been set.
func (o *PublicNetworkCreate) HasIpBlocks() bool {
	if o != nil && !IsNil(o.IpBlocks) {
		return true
	}

	return false
}

// SetIpBlocks gets a reference to the given []PublicNetworkIpBlock and assigns it to the IpBlocks field.
func (o *PublicNetworkCreate) SetIpBlocks(v []PublicNetworkIpBlock) {
	o.IpBlocks = v
}

func (o PublicNetworkCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicNetworkCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["location"] = o.Location
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !IsNil(o.IpBlocks) {
		toSerialize["ipBlocks"] = o.IpBlocks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicNetworkCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"location",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPublicNetworkCreate := _PublicNetworkCreate{}

	err = json.Unmarshal(data, &varPublicNetworkCreate)

	if err != nil {
		return err
	}

	*o = PublicNetworkCreate(varPublicNetworkCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "location")
		delete(additionalProperties, "vlanId")
		delete(additionalProperties, "ipBlocks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicNetworkCreate struct {
	value *PublicNetworkCreate
	isSet bool
}

func (v NullablePublicNetworkCreate) Get() *PublicNetworkCreate {
	return v.value
}

func (v *NullablePublicNetworkCreate) Set(val *PublicNetworkCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicNetworkCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicNetworkCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicNetworkCreate(val *PublicNetworkCreate) *NullablePublicNetworkCreate {
	return &NullablePublicNetworkCreate{value: val, isSet: true}
}

func (v NullablePublicNetworkCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicNetworkCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
