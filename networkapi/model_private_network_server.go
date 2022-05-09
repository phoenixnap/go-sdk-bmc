/*
Networks API

Create, list, edit and delete public/private networks with the Network API. Use public networks to place multiple  servers on the same network or VLAN. Assign new servers with IP addresses from the same CIDR range. Use private  networks to avoid unnecessary egress data charges. Model your networks according to your business needs.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#multi-private-backend-network-api' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/networks/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkapi

import (
	"encoding/json"
)

// PrivateNetworkServer Server details linked to the Private Network.
type PrivateNetworkServer struct {
	// The server identifier.
	Id string `json:"id"`
	// List of private IPs associated to the server.
	Ips []string `json:"ips"`
}

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
func (o *PrivateNetworkServer) GetIpsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ips, true
}

// SetIps sets field value
func (o *PrivateNetworkServer) SetIps(v []string) {
	o.Ips = v
}

func (o PrivateNetworkServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["ips"] = o.Ips
	}
	return json.Marshal(toSerialize)
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
