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

// checks if the PrivateNetworkServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateNetworkServer{}

// PrivateNetworkServer Server details linked to the Private Network.
type PrivateNetworkServer struct {
	// The server identifier.
	Id string `json:"id"`
	// List of private IPs associated to the server.
	Ips                  []string `json:"ips"`
	AdditionalProperties map[string]interface{}
}

type _PrivateNetworkServer PrivateNetworkServer

// NewPrivateNetworkServer instantiates a new PrivateNetworkServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateNetworkServer(id string, ips []string) *PrivateNetworkServer {
	this := PrivateNetworkServer{}
	this.Id = id
	this.Ips = ips
	return &this
}

// NewPrivateNetworkServerWithDefaults instantiates a new PrivateNetworkServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateNetworkServerWithDefaults() *PrivateNetworkServer {
	this := PrivateNetworkServer{}
	return &this
}

// GetId returns the Id field value
func (o *PrivateNetworkServer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkServer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrivateNetworkServer) SetId(v string) {
	o.Id = v
}

// GetIps returns the Ips field value
func (o *PrivateNetworkServer) GetIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
func (o *PrivateNetworkServer) GetIpsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ips, true
}

// SetIps sets field value
func (o *PrivateNetworkServer) SetIps(v []string) {
	o.Ips = v
}

func (o PrivateNetworkServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateNetworkServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["ips"] = o.Ips

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrivateNetworkServer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"ips",
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

	varPrivateNetworkServer := _PrivateNetworkServer{}

	err = json.Unmarshal(data, &varPrivateNetworkServer)

	if err != nil {
		return err
	}

	*o = PrivateNetworkServer(varPrivateNetworkServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "ips")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrivateNetworkServer struct {
	value *PrivateNetworkServer
	isSet bool
}

func (v NullablePrivateNetworkServer) Get() *PrivateNetworkServer {
	return v.value
}

func (v *NullablePrivateNetworkServer) Set(val *PrivateNetworkServer) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateNetworkServer) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateNetworkServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateNetworkServer(val *PrivateNetworkServer) *NullablePrivateNetworkServer {
	return &NullablePrivateNetworkServer{value: val, isSet: true}
}

func (v NullablePrivateNetworkServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateNetworkServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
