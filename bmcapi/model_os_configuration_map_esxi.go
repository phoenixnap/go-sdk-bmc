/*
Bare Metal Cloud API

Bare Metal Cloud API allows you to create, power off, power on, reset, reboot, or shut down your server. Also deprovision servers, manage SSH key details, and a lot more. Manage your infrastructure more efficiently using just a few simple API calls. </br></br>**All URLs are relative to (https://api.phoenixnap.com/bmc/v0/)**

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmcapi

import (
	"encoding/json"
)

// OsConfigurationMapEsxi VMWare ESXi configuration properties.
type OsConfigurationMapEsxi struct {
	// Password set for user root on an ESXi server which will only be returned in response to provisioning a server.
	RootPassword *string `json:"rootPassword,omitempty"`
	// The URL of the management UI which will only be returned in response to provisioning a server.
	ManagementUiUrl *string `json:"managementUiUrl,omitempty"`
	// List of IPs allowed to access the Management UI. Supported in single IP, CIDR and range format. When undefined, Management UI is disabled. This will only be returned in response to provisioning a server.
	ManagementAccessAllowedIps *[]string `json:"managementAccessAllowedIps,omitempty"`
}

// NewOsConfigurationMapEsxi instantiates a new OsConfigurationMapEsxi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsConfigurationMapEsxi() *OsConfigurationMapEsxi {
	this := OsConfigurationMapEsxi{}
	return &this
}

// NewOsConfigurationMapEsxiWithDefaults instantiates a new OsConfigurationMapEsxi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsConfigurationMapEsxiWithDefaults() *OsConfigurationMapEsxi {
	this := OsConfigurationMapEsxi{}
	return &this
}

// GetRootPassword returns the RootPassword field value if set, zero value otherwise.
func (o *OsConfigurationMapEsxi) GetRootPassword() string {
	if o == nil || o.RootPassword == nil {
		var ret string
		return ret
	}
	return *o.RootPassword
}

// GetRootPasswordOk returns a tuple with the RootPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationMapEsxi) GetRootPasswordOk() (*string, bool) {
	if o == nil || o.RootPassword == nil {
		return nil, false
	}
	return o.RootPassword, true
}

// HasRootPassword returns a boolean if a field has been set.
func (o *OsConfigurationMapEsxi) HasRootPassword() bool {
	if o != nil && o.RootPassword != nil {
		return true
	}

	return false
}

// SetRootPassword gets a reference to the given string and assigns it to the RootPassword field.
func (o *OsConfigurationMapEsxi) SetRootPassword(v string) {
	o.RootPassword = &v
}

// GetManagementUiUrl returns the ManagementUiUrl field value if set, zero value otherwise.
func (o *OsConfigurationMapEsxi) GetManagementUiUrl() string {
	if o == nil || o.ManagementUiUrl == nil {
		var ret string
		return ret
	}
	return *o.ManagementUiUrl
}

// GetManagementUiUrlOk returns a tuple with the ManagementUiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationMapEsxi) GetManagementUiUrlOk() (*string, bool) {
	if o == nil || o.ManagementUiUrl == nil {
		return nil, false
	}
	return o.ManagementUiUrl, true
}

// HasManagementUiUrl returns a boolean if a field has been set.
func (o *OsConfigurationMapEsxi) HasManagementUiUrl() bool {
	if o != nil && o.ManagementUiUrl != nil {
		return true
	}

	return false
}

// SetManagementUiUrl gets a reference to the given string and assigns it to the ManagementUiUrl field.
func (o *OsConfigurationMapEsxi) SetManagementUiUrl(v string) {
	o.ManagementUiUrl = &v
}

// GetManagementAccessAllowedIps returns the ManagementAccessAllowedIps field value if set, zero value otherwise.
func (o *OsConfigurationMapEsxi) GetManagementAccessAllowedIps() []string {
	if o == nil || o.ManagementAccessAllowedIps == nil {
		var ret []string
		return ret
	}
	return *o.ManagementAccessAllowedIps
}

// GetManagementAccessAllowedIpsOk returns a tuple with the ManagementAccessAllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationMapEsxi) GetManagementAccessAllowedIpsOk() (*[]string, bool) {
	if o == nil || o.ManagementAccessAllowedIps == nil {
		return nil, false
	}
	return o.ManagementAccessAllowedIps, true
}

// HasManagementAccessAllowedIps returns a boolean if a field has been set.
func (o *OsConfigurationMapEsxi) HasManagementAccessAllowedIps() bool {
	if o != nil && o.ManagementAccessAllowedIps != nil {
		return true
	}

	return false
}

// SetManagementAccessAllowedIps gets a reference to the given []string and assigns it to the ManagementAccessAllowedIps field.
func (o *OsConfigurationMapEsxi) SetManagementAccessAllowedIps(v []string) {
	o.ManagementAccessAllowedIps = &v
}

func (o OsConfigurationMapEsxi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RootPassword != nil {
		toSerialize["rootPassword"] = o.RootPassword
	}
	if o.ManagementUiUrl != nil {
		toSerialize["managementUiUrl"] = o.ManagementUiUrl
	}
	if o.ManagementAccessAllowedIps != nil {
		toSerialize["managementAccessAllowedIps"] = o.ManagementAccessAllowedIps
	}
	return json.Marshal(toSerialize)
}

type NullableOsConfigurationMapEsxi struct {
	value *OsConfigurationMapEsxi
	isSet bool
}

func (v NullableOsConfigurationMapEsxi) Get() *OsConfigurationMapEsxi {
	return v.value
}

func (v *NullableOsConfigurationMapEsxi) Set(val *OsConfigurationMapEsxi) {
	v.value = val
	v.isSet = true
}

func (v NullableOsConfigurationMapEsxi) IsSet() bool {
	return v.isSet
}

func (v *NullableOsConfigurationMapEsxi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsConfigurationMapEsxi(val *OsConfigurationMapEsxi) *NullableOsConfigurationMapEsxi {
	return &NullableOsConfigurationMapEsxi{value: val, isSet: true}
}

func (v NullableOsConfigurationMapEsxi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsConfigurationMapEsxi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
