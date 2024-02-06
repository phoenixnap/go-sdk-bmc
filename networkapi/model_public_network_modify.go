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
)

// checks if the PublicNetworkModify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicNetworkModify{}

// PublicNetworkModify Public Network Modifiable Details.
type PublicNetworkModify struct {
	// A friendly name given to the network. This name should be unique.
	Name *string `json:"name,omitempty"`
	// The description of this public network
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PublicNetworkModify PublicNetworkModify

// NewPublicNetworkModify instantiates a new PublicNetworkModify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicNetworkModify() *PublicNetworkModify {
	this := PublicNetworkModify{}
	return &this
}

// NewPublicNetworkModifyWithDefaults instantiates a new PublicNetworkModify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicNetworkModifyWithDefaults() *PublicNetworkModify {
	this := PublicNetworkModify{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PublicNetworkModify) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicNetworkModify) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PublicNetworkModify) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PublicNetworkModify) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PublicNetworkModify) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicNetworkModify) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PublicNetworkModify) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PublicNetworkModify) SetDescription(v string) {
	o.Description = &v
}

func (o PublicNetworkModify) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicNetworkModify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicNetworkModify) UnmarshalJSON(data []byte) (err error) {
	varPublicNetworkModify := _PublicNetworkModify{}

	err = json.Unmarshal(data, &varPublicNetworkModify)

	if err != nil {
		return err
	}

	*o = PublicNetworkModify(varPublicNetworkModify)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicNetworkModify struct {
	value *PublicNetworkModify
	isSet bool
}

func (v NullablePublicNetworkModify) Get() *PublicNetworkModify {
	return v.value
}

func (v *NullablePublicNetworkModify) Set(val *PublicNetworkModify) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicNetworkModify) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicNetworkModify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicNetworkModify(val *PublicNetworkModify) *NullablePublicNetworkModify {
	return &NullablePublicNetworkModify{value: val, isSet: true}
}

func (v NullablePublicNetworkModify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicNetworkModify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
