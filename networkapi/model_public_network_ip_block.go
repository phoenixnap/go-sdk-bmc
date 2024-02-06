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

// checks if the PublicNetworkIpBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicNetworkIpBlock{}

// PublicNetworkIpBlock The assigned IP block to the Public Network.
type PublicNetworkIpBlock struct {
	// The IP Block identifier.
	Id                   string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _PublicNetworkIpBlock PublicNetworkIpBlock

// NewPublicNetworkIpBlock instantiates a new PublicNetworkIpBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicNetworkIpBlock(id string) *PublicNetworkIpBlock {
	this := PublicNetworkIpBlock{}
	this.Id = id
	return &this
}

// NewPublicNetworkIpBlockWithDefaults instantiates a new PublicNetworkIpBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicNetworkIpBlockWithDefaults() *PublicNetworkIpBlock {
	this := PublicNetworkIpBlock{}
	return &this
}

// GetId returns the Id field value
func (o *PublicNetworkIpBlock) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicNetworkIpBlock) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicNetworkIpBlock) SetId(v string) {
	o.Id = v
}

func (o PublicNetworkIpBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicNetworkIpBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicNetworkIpBlock) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varPublicNetworkIpBlock := _PublicNetworkIpBlock{}

	err = json.Unmarshal(data, &varPublicNetworkIpBlock)

	if err != nil {
		return err
	}

	*o = PublicNetworkIpBlock(varPublicNetworkIpBlock)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicNetworkIpBlock struct {
	value *PublicNetworkIpBlock
	isSet bool
}

func (v NullablePublicNetworkIpBlock) Get() *PublicNetworkIpBlock {
	return v.value
}

func (v *NullablePublicNetworkIpBlock) Set(val *PublicNetworkIpBlock) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicNetworkIpBlock) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicNetworkIpBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicNetworkIpBlock(val *PublicNetworkIpBlock) *NullablePublicNetworkIpBlock {
	return &NullablePublicNetworkIpBlock{value: val, isSet: true}
}

func (v NullablePublicNetworkIpBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicNetworkIpBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
