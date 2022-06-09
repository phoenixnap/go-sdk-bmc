/*
Bare Metal Cloud API

Create, power off, power on, reset, reboot, or shut down your server with the Bare Metal Cloud API.  Deprovision servers, get or edit SSH key details, assign public IPs, assign servers to networks and a lot more.  Manage your infrastructure more efficiently using just a few simple API calls.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/how-to-deploy-bare-metal-cloud-server' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/bmc/v1/)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmcapi

import (
	"encoding/json"
)

// NetworkConfiguration Entire network details of bare metal server.
type NetworkConfiguration struct {
	// The address of the gateway assigned / to assign to the server. When used as part of request body, IP address has to be part of private/public network assigned to this server.
	GatewayAddress              *string                      `json:"gatewayAddress,omitempty"`
	PrivateNetworkConfiguration *PrivateNetworkConfiguration `json:"privateNetworkConfiguration,omitempty"`
	IpBlocksConfiguration       *IpBlocksConfiguration       `json:"ipBlocksConfiguration,omitempty"`
	PublicNetworkConfiguration  *PublicNetworkConfiguration  `json:"publicNetworkConfiguration,omitempty"`
}

// NewNetworkConfiguration instantiates a new NetworkConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkConfiguration() *NetworkConfiguration {
	this := NetworkConfiguration{}
	return &this
}

// NewNetworkConfigurationWithDefaults instantiates a new NetworkConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkConfigurationWithDefaults() *NetworkConfiguration {
	this := NetworkConfiguration{}
	return &this
}

// GetGatewayAddress returns the GatewayAddress field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetGatewayAddress() string {
	if o == nil || o.GatewayAddress == nil {
		var ret string
		return ret
	}
	return *o.GatewayAddress
}

// GetGatewayAddressOk returns a tuple with the GatewayAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetGatewayAddressOk() (*string, bool) {
	if o == nil || o.GatewayAddress == nil {
		return nil, false
	}
	return o.GatewayAddress, true
}

// HasGatewayAddress returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasGatewayAddress() bool {
	if o != nil && o.GatewayAddress != nil {
		return true
	}

	return false
}

// SetGatewayAddress gets a reference to the given string and assigns it to the GatewayAddress field.
func (o *NetworkConfiguration) SetGatewayAddress(v string) {
	o.GatewayAddress = &v
}

// GetPrivateNetworkConfiguration returns the PrivateNetworkConfiguration field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetPrivateNetworkConfiguration() PrivateNetworkConfiguration {
	if o == nil || o.PrivateNetworkConfiguration == nil {
		var ret PrivateNetworkConfiguration
		return ret
	}
	return *o.PrivateNetworkConfiguration
}

// GetPrivateNetworkConfigurationOk returns a tuple with the PrivateNetworkConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetPrivateNetworkConfigurationOk() (*PrivateNetworkConfiguration, bool) {
	if o == nil || o.PrivateNetworkConfiguration == nil {
		return nil, false
	}
	return o.PrivateNetworkConfiguration, true
}

// HasPrivateNetworkConfiguration returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasPrivateNetworkConfiguration() bool {
	if o != nil && o.PrivateNetworkConfiguration != nil {
		return true
	}

	return false
}

// SetPrivateNetworkConfiguration gets a reference to the given PrivateNetworkConfiguration and assigns it to the PrivateNetworkConfiguration field.
func (o *NetworkConfiguration) SetPrivateNetworkConfiguration(v PrivateNetworkConfiguration) {
	o.PrivateNetworkConfiguration = &v
}

// GetIpBlocksConfiguration returns the IpBlocksConfiguration field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetIpBlocksConfiguration() IpBlocksConfiguration {
	if o == nil || o.IpBlocksConfiguration == nil {
		var ret IpBlocksConfiguration
		return ret
	}
	return *o.IpBlocksConfiguration
}

// GetIpBlocksConfigurationOk returns a tuple with the IpBlocksConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetIpBlocksConfigurationOk() (*IpBlocksConfiguration, bool) {
	if o == nil || o.IpBlocksConfiguration == nil {
		return nil, false
	}
	return o.IpBlocksConfiguration, true
}

// HasIpBlocksConfiguration returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasIpBlocksConfiguration() bool {
	if o != nil && o.IpBlocksConfiguration != nil {
		return true
	}

	return false
}

// SetIpBlocksConfiguration gets a reference to the given IpBlocksConfiguration and assigns it to the IpBlocksConfiguration field.
func (o *NetworkConfiguration) SetIpBlocksConfiguration(v IpBlocksConfiguration) {
	o.IpBlocksConfiguration = &v
}

// GetPublicNetworkConfiguration returns the PublicNetworkConfiguration field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetPublicNetworkConfiguration() PublicNetworkConfiguration {
	if o == nil || o.PublicNetworkConfiguration == nil {
		var ret PublicNetworkConfiguration
		return ret
	}
	return *o.PublicNetworkConfiguration
}

// GetPublicNetworkConfigurationOk returns a tuple with the PublicNetworkConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetPublicNetworkConfigurationOk() (*PublicNetworkConfiguration, bool) {
	if o == nil || o.PublicNetworkConfiguration == nil {
		return nil, false
	}
	return o.PublicNetworkConfiguration, true
}

// HasPublicNetworkConfiguration returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasPublicNetworkConfiguration() bool {
	if o != nil && o.PublicNetworkConfiguration != nil {
		return true
	}

	return false
}

// SetPublicNetworkConfiguration gets a reference to the given PublicNetworkConfiguration and assigns it to the PublicNetworkConfiguration field.
func (o *NetworkConfiguration) SetPublicNetworkConfiguration(v PublicNetworkConfiguration) {
	o.PublicNetworkConfiguration = &v
}

func (o NetworkConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayAddress != nil {
		toSerialize["gatewayAddress"] = o.GatewayAddress
	}
	if o.PrivateNetworkConfiguration != nil {
		toSerialize["privateNetworkConfiguration"] = o.PrivateNetworkConfiguration
	}
	if o.IpBlocksConfiguration != nil {
		toSerialize["ipBlocksConfiguration"] = o.IpBlocksConfiguration
	}
	if o.PublicNetworkConfiguration != nil {
		toSerialize["publicNetworkConfiguration"] = o.PublicNetworkConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkConfiguration struct {
	value *NetworkConfiguration
	isSet bool
}

func (v NullableNetworkConfiguration) Get() *NetworkConfiguration {
	return v.value
}

func (v *NullableNetworkConfiguration) Set(val *NetworkConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkConfiguration(val *NetworkConfiguration) *NullableNetworkConfiguration {
	return &NullableNetworkConfiguration{value: val, isSet: true}
}

func (v NullableNetworkConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
