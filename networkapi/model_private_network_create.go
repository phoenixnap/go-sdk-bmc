/*
Networks API

Create, list, edit and delete public/private networks with the Network API. Use public networks to place multiple  servers on the same network or VLAN. Assign new servers with IP addresses from the same CIDR range. Use private  networks to avoid unnecessary egress data charges. Model your networks according to your business needs.<br> <br> <span class='pnap-api-knowledge-base-link'> Helpful knowledge base articles are available for  <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#multi-private-backend-network-api' target='_blank'>multi-private backend networks</a>,  <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#ftoc-heading-15' target='_blank'>public networks</a> and <a href='https://phoenixnap.com/kb/border-gateway-protocol-bmc' target='_blank'>border gateway protocol peer groups</a>. </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/networks/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkapi

import (
	"encoding/json"
	"fmt"
)

// checks if the PrivateNetworkCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateNetworkCreate{}

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
	// The VLAN that will be assigned to this network.
	VlanId *int32 `json:"vlanId,omitempty"`
	// IP range associated with this private network in CIDR notation.<br> Setting the `force` query parameter to `true` allows you to skip assigning a specific IP range to network.
	Cidr                 *string `json:"cidr,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrivateNetworkCreate PrivateNetworkCreate

// NewPrivateNetworkCreate instantiates a new PrivateNetworkCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateNetworkCreate(name string, location string) *PrivateNetworkCreate {
	this := PrivateNetworkCreate{}
	this.Name = name
	this.Location = location
	var locationDefault bool = false
	this.LocationDefault = &locationDefault
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
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetworkCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PrivateNetworkCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
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
	if o == nil || IsNil(o.LocationDefault) {
		var ret bool
		return ret
	}
	return *o.LocationDefault
}

// GetLocationDefaultOk returns a tuple with the LocationDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetworkCreate) GetLocationDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.LocationDefault) {
		return nil, false
	}
	return o.LocationDefault, true
}

// HasLocationDefault returns a boolean if a field has been set.
func (o *PrivateNetworkCreate) HasLocationDefault() bool {
	if o != nil && !IsNil(o.LocationDefault) {
		return true
	}

	return false
}

// SetLocationDefault gets a reference to the given bool and assigns it to the LocationDefault field.
func (o *PrivateNetworkCreate) SetLocationDefault(v bool) {
	o.LocationDefault = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *PrivateNetworkCreate) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetworkCreate) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *PrivateNetworkCreate) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *PrivateNetworkCreate) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *PrivateNetworkCreate) GetCidr() string {
	if o == nil || IsNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetworkCreate) GetCidrOk() (*string, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *PrivateNetworkCreate) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *PrivateNetworkCreate) SetCidr(v string) {
	o.Cidr = &v
}

func (o PrivateNetworkCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateNetworkCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["location"] = o.Location
	if !IsNil(o.LocationDefault) {
		toSerialize["locationDefault"] = o.LocationDefault
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrivateNetworkCreate) UnmarshalJSON(data []byte) (err error) {
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

	varPrivateNetworkCreate := _PrivateNetworkCreate{}

	err = json.Unmarshal(data, &varPrivateNetworkCreate)

	if err != nil {
		return err
	}

	*o = PrivateNetworkCreate(varPrivateNetworkCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "location")
		delete(additionalProperties, "locationDefault")
		delete(additionalProperties, "vlanId")
		delete(additionalProperties, "cidr")
		o.AdditionalProperties = additionalProperties
	}

	return err
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
