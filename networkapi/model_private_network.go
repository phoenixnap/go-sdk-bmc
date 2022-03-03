/*
Networks API

Use the Networks API to create, list, edit, and delete private networks to best fit your business needs. Private networks allow your servers to communicate without connecting to the public internet, avoiding unnecessary egress data charges.<br> <br> <span class=\"pnap-api-knowledge-base-link\"> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#multi-private-backend-network-api' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/networks/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkapi

import (
	"encoding/json"
)

// PrivateNetwork Private Network details.
type PrivateNetwork struct {
	// The private network identifier.
	Id string `json:"id"`
	// The friendly name of this private network.
	Name string `json:"name"`
	// The description of this private network.
	Description *string `json:"description,omitempty"`
	// The VLAN of this private network.
	VlanId int32 `json:"vlanId"`
	// The type of the private network.
	Type string `json:"type"`
	// The location of this private network.
	Location string `json:"location"`
	// Identifies network as the default private network for the specified location.
	LocationDefault bool `json:"locationDefault"`
	// IP range associated with this private network in CIDR notation.
	Cidr    string                 `json:"cidr"`
	Servers []PrivateNetworkServer `json:"servers"`
}

// NewPrivateNetwork instantiates a new PrivateNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateNetwork(id string, name string, vlanId int32, type_ string, location string, locationDefault bool, cidr string, servers []PrivateNetworkServer) *PrivateNetwork {
	this := PrivateNetwork{}
	this.Id = id
	this.Name = name
	this.VlanId = vlanId
	this.Type = type_
	this.Location = location
	this.LocationDefault = locationDefault
	this.Cidr = cidr
	this.Servers = servers
	return &this
}

// NewPrivateNetworkWithDefaults instantiates a new PrivateNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateNetworkWithDefaults() *PrivateNetwork {
	this := PrivateNetwork{}
	return &this
}

// GetId returns the Id field value
func (o *PrivateNetwork) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrivateNetwork) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrivateNetwork) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PrivateNetwork) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrivateNetwork) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrivateNetwork) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PrivateNetwork) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetwork) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PrivateNetwork) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PrivateNetwork) SetDescription(v string) {
	o.Description = &v
}

// GetVlanId returns the VlanId field value
func (o *PrivateNetwork) GetVlanId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value
// and a boolean to check if the value has been set.
func (o *PrivateNetwork) GetVlanIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VlanId, true
}

// SetVlanId sets field value
func (o *PrivateNetwork) SetVlanId(v int32) {
	o.VlanId = v
}

// GetType returns the Type field value
func (o *PrivateNetwork) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PrivateNetwork) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PrivateNetwork) SetType(v string) {
	o.Type = v
}

// GetLocation returns the Location field value
func (o *PrivateNetwork) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *PrivateNetwork) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *PrivateNetwork) SetLocation(v string) {
	o.Location = v
}

// GetLocationDefault returns the LocationDefault field value
func (o *PrivateNetwork) GetLocationDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LocationDefault
}

// GetLocationDefaultOk returns a tuple with the LocationDefault field value
// and a boolean to check if the value has been set.
func (o *PrivateNetwork) GetLocationDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationDefault, true
}

// SetLocationDefault sets field value
func (o *PrivateNetwork) SetLocationDefault(v bool) {
	o.LocationDefault = v
}

// GetCidr returns the Cidr field value
func (o *PrivateNetwork) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *PrivateNetwork) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *PrivateNetwork) SetCidr(v string) {
	o.Cidr = v
}

// GetServers returns the Servers field value
func (o *PrivateNetwork) GetServers() []PrivateNetworkServer {
	if o == nil {
		var ret []PrivateNetworkServer
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
func (o *PrivateNetwork) GetServersOk() (*[]PrivateNetworkServer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Servers, true
}

// SetServers sets field value
func (o *PrivateNetwork) SetServers(v []PrivateNetworkServer) {
	o.Servers = v
}

func (o PrivateNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["vlanId"] = o.VlanId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["locationDefault"] = o.LocationDefault
	}
	if true {
		toSerialize["cidr"] = o.Cidr
	}
	if true {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateNetwork struct {
	value *PrivateNetwork
	isSet bool
}

func (v NullablePrivateNetwork) Get() *PrivateNetwork {
	return v.value
}

func (v *NullablePrivateNetwork) Set(val *PrivateNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateNetwork(val *PrivateNetwork) *NullablePrivateNetwork {
	return &NullablePrivateNetwork{value: val, isSet: true}
}

func (v NullablePrivateNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
