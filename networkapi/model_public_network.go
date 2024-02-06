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
	"time"
)

// checks if the PublicNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicNetwork{}

// PublicNetwork Public Network details.
type PublicNetwork struct {
	// The public network identifier.
	Id string `json:"id"`
	// The VLAN of this public network.
	VlanId int32 `json:"vlanId"`
	// A list of resources that are members of this public network.
	Memberships []NetworkMembership `json:"memberships"`
	// The friendly name of this public network.
	Name string `json:"name"`
	// The location of this public network. Supported values are `PHX`, `ASH`, `SGP`, `NLD`, `CHI`, `SEA` and `AUS`.
	Location string `json:"location"`
	// The description of this public network.
	Description *string `json:"description,omitempty"`
	// The status of the public network. Can have one of the following values: `BUSY`, `READY`, `DELETING` or `ERROR`.
	Status string `json:"status"`
	// Date and time when this public network was created.
	CreatedOn time.Time `json:"createdOn"`
	// A list of IP Blocks that are associated with this public network.
	IpBlocks             []PublicNetworkIpBlock `json:"ipBlocks"`
	AdditionalProperties map[string]interface{}
}

type _PublicNetwork PublicNetwork

// NewPublicNetwork instantiates a new PublicNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicNetwork(id string, vlanId int32, memberships []NetworkMembership, name string, location string, status string, createdOn time.Time, ipBlocks []PublicNetworkIpBlock) *PublicNetwork {
	this := PublicNetwork{}
	this.Id = id
	this.VlanId = vlanId
	this.Memberships = memberships
	this.Name = name
	this.Location = location
	this.Status = status
	this.CreatedOn = createdOn
	this.IpBlocks = ipBlocks
	return &this
}

// NewPublicNetworkWithDefaults instantiates a new PublicNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicNetworkWithDefaults() *PublicNetwork {
	this := PublicNetwork{}
	return &this
}

// GetId returns the Id field value
func (o *PublicNetwork) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicNetwork) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicNetwork) SetId(v string) {
	o.Id = v
}

// GetVlanId returns the VlanId field value
func (o *PublicNetwork) GetVlanId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value
// and a boolean to check if the value has been set.
func (o *PublicNetwork) GetVlanIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VlanId, true
}

// SetVlanId sets field value
func (o *PublicNetwork) SetVlanId(v int32) {
	o.VlanId = v
}

// GetMemberships returns the Memberships field value
func (o *PublicNetwork) GetMemberships() []NetworkMembership {
	if o == nil {
		var ret []NetworkMembership
		return ret
	}

	return o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value
// and a boolean to check if the value has been set.
func (o *PublicNetwork) GetMembershipsOk() ([]NetworkMembership, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memberships, true
}

// SetMemberships sets field value
func (o *PublicNetwork) SetMemberships(v []NetworkMembership) {
	o.Memberships = v
}

// GetName returns the Name field value
func (o *PublicNetwork) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublicNetwork) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PublicNetwork) SetName(v string) {
	o.Name = v
}

// GetLocation returns the Location field value
func (o *PublicNetwork) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *PublicNetwork) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *PublicNetwork) SetLocation(v string) {
	o.Location = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PublicNetwork) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicNetwork) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PublicNetwork) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PublicNetwork) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *PublicNetwork) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PublicNetwork) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PublicNetwork) SetStatus(v string) {
	o.Status = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *PublicNetwork) GetCreatedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *PublicNetwork) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *PublicNetwork) SetCreatedOn(v time.Time) {
	o.CreatedOn = v
}

// GetIpBlocks returns the IpBlocks field value
func (o *PublicNetwork) GetIpBlocks() []PublicNetworkIpBlock {
	if o == nil {
		var ret []PublicNetworkIpBlock
		return ret
	}

	return o.IpBlocks
}

// GetIpBlocksOk returns a tuple with the IpBlocks field value
// and a boolean to check if the value has been set.
func (o *PublicNetwork) GetIpBlocksOk() ([]PublicNetworkIpBlock, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpBlocks, true
}

// SetIpBlocks sets field value
func (o *PublicNetwork) SetIpBlocks(v []PublicNetworkIpBlock) {
	o.IpBlocks = v
}

func (o PublicNetwork) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["vlanId"] = o.VlanId
	toSerialize["memberships"] = o.Memberships
	toSerialize["name"] = o.Name
	toSerialize["location"] = o.Location
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["status"] = o.Status
	toSerialize["createdOn"] = o.CreatedOn
	toSerialize["ipBlocks"] = o.IpBlocks

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicNetwork) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"vlanId",
		"memberships",
		"name",
		"location",
		"status",
		"createdOn",
		"ipBlocks",
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

	varPublicNetwork := _PublicNetwork{}

	err = json.Unmarshal(data, &varPublicNetwork)

	if err != nil {
		return err
	}

	*o = PublicNetwork(varPublicNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "vlanId")
		delete(additionalProperties, "memberships")
		delete(additionalProperties, "name")
		delete(additionalProperties, "location")
		delete(additionalProperties, "description")
		delete(additionalProperties, "status")
		delete(additionalProperties, "createdOn")
		delete(additionalProperties, "ipBlocks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicNetwork struct {
	value *PublicNetwork
	isSet bool
}

func (v NullablePublicNetwork) Get() *PublicNetwork {
	return v.value
}

func (v *NullablePublicNetwork) Set(val *PublicNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicNetwork(val *PublicNetwork) *NullablePublicNetwork {
	return &NullablePublicNetwork{value: val, isSet: true}
}

func (v NullablePublicNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
