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

// OsConfigurationWindows Windows OS configuration properties.
type OsConfigurationWindows struct {
	// List of IPs allowed for RDP access to Windows OS. Supported in single IP, CIDR and range format. When undefined, RDP is disabled. To allow RDP access from any IP use 0.0.0.0/0. This will only be returned in response to provisioning a server.
	RdpAllowedIps *[]string `json:"rdpAllowedIps,omitempty"`
}

// NewOsConfigurationWindows instantiates a new OsConfigurationWindows object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsConfigurationWindows() *OsConfigurationWindows {
	this := OsConfigurationWindows{}
	return &this
}

// NewOsConfigurationWindowsWithDefaults instantiates a new OsConfigurationWindows object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsConfigurationWindowsWithDefaults() *OsConfigurationWindows {
	this := OsConfigurationWindows{}
	return &this
}

// GetRdpAllowedIps returns the RdpAllowedIps field value if set, zero value otherwise.
func (o *OsConfigurationWindows) GetRdpAllowedIps() []string {
	if o == nil || o.RdpAllowedIps == nil {
		var ret []string
		return ret
	}
	return *o.RdpAllowedIps
}

// GetRdpAllowedIpsOk returns a tuple with the RdpAllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationWindows) GetRdpAllowedIpsOk() (*[]string, bool) {
	if o == nil || o.RdpAllowedIps == nil {
		return nil, false
	}
	return o.RdpAllowedIps, true
}

// HasRdpAllowedIps returns a boolean if a field has been set.
func (o *OsConfigurationWindows) HasRdpAllowedIps() bool {
	if o != nil && o.RdpAllowedIps != nil {
		return true
	}

	return false
}

// SetRdpAllowedIps gets a reference to the given []string and assigns it to the RdpAllowedIps field.
func (o *OsConfigurationWindows) SetRdpAllowedIps(v []string) {
	o.RdpAllowedIps = &v
}

func (o OsConfigurationWindows) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RdpAllowedIps != nil {
		toSerialize["rdpAllowedIps"] = o.RdpAllowedIps
	}
	return json.Marshal(toSerialize)
}

type NullableOsConfigurationWindows struct {
	value *OsConfigurationWindows
	isSet bool
}

func (v NullableOsConfigurationWindows) Get() *OsConfigurationWindows {
	return v.value
}

func (v *NullableOsConfigurationWindows) Set(val *OsConfigurationWindows) {
	v.value = val
	v.isSet = true
}

func (v NullableOsConfigurationWindows) IsSet() bool {
	return v.isSet
}

func (v *NullableOsConfigurationWindows) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsConfigurationWindows(val *OsConfigurationWindows) *NullableOsConfigurationWindows {
	return &NullableOsConfigurationWindows{value: val, isSet: true}
}

func (v NullableOsConfigurationWindows) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsConfigurationWindows) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
