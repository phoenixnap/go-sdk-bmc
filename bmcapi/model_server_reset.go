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

// ServerReset Reset bare metal server.
type ServerReset struct {
	// Whether or not to install SSH keys marked as default in addition to any SSH keys specified in this request.
	InstallDefaultSshKeys *bool `json:"installDefaultSshKeys,omitempty"`
	// A list of SSH keys that will be installed on the server.
	SshKeys *[]string `json:"sshKeys,omitempty"`
	// A list of SSH key IDs that will be installed on the server in addition to any SSH keys specified in this request.
	SshKeyIds       *[]string           `json:"sshKeyIds,omitempty"`
	OsConfiguration *OsConfigurationMap `json:"osConfiguration,omitempty"`
}

// NewServerReset instantiates a new ServerReset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerReset() *ServerReset {
	this := ServerReset{}
	var installDefaultSshKeys bool = true
	this.InstallDefaultSshKeys = &installDefaultSshKeys
	return &this
}

// NewServerResetWithDefaults instantiates a new ServerReset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerResetWithDefaults() *ServerReset {
	this := ServerReset{}
	var installDefaultSshKeys bool = true
	this.InstallDefaultSshKeys = &installDefaultSshKeys
	return &this
}

// GetInstallDefaultSshKeys returns the InstallDefaultSshKeys field value if set, zero value otherwise.
func (o *ServerReset) GetInstallDefaultSshKeys() bool {
	if o == nil || o.InstallDefaultSshKeys == nil {
		var ret bool
		return ret
	}
	return *o.InstallDefaultSshKeys
}

// GetInstallDefaultSshKeysOk returns a tuple with the InstallDefaultSshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerReset) GetInstallDefaultSshKeysOk() (*bool, bool) {
	if o == nil || o.InstallDefaultSshKeys == nil {
		return nil, false
	}
	return o.InstallDefaultSshKeys, true
}

// HasInstallDefaultSshKeys returns a boolean if a field has been set.
func (o *ServerReset) HasInstallDefaultSshKeys() bool {
	if o != nil && o.InstallDefaultSshKeys != nil {
		return true
	}

	return false
}

// SetInstallDefaultSshKeys gets a reference to the given bool and assigns it to the InstallDefaultSshKeys field.
func (o *ServerReset) SetInstallDefaultSshKeys(v bool) {
	o.InstallDefaultSshKeys = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *ServerReset) GetSshKeys() []string {
	if o == nil || o.SshKeys == nil {
		var ret []string
		return ret
	}
	return *o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerReset) GetSshKeysOk() (*[]string, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *ServerReset) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []string and assigns it to the SshKeys field.
func (o *ServerReset) SetSshKeys(v []string) {
	o.SshKeys = &v
}

// GetSshKeyIds returns the SshKeyIds field value if set, zero value otherwise.
func (o *ServerReset) GetSshKeyIds() []string {
	if o == nil || o.SshKeyIds == nil {
		var ret []string
		return ret
	}
	return *o.SshKeyIds
}

// GetSshKeyIdsOk returns a tuple with the SshKeyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerReset) GetSshKeyIdsOk() (*[]string, bool) {
	if o == nil || o.SshKeyIds == nil {
		return nil, false
	}
	return o.SshKeyIds, true
}

// HasSshKeyIds returns a boolean if a field has been set.
func (o *ServerReset) HasSshKeyIds() bool {
	if o != nil && o.SshKeyIds != nil {
		return true
	}

	return false
}

// SetSshKeyIds gets a reference to the given []string and assigns it to the SshKeyIds field.
func (o *ServerReset) SetSshKeyIds(v []string) {
	o.SshKeyIds = &v
}

// GetOsConfiguration returns the OsConfiguration field value if set, zero value otherwise.
func (o *ServerReset) GetOsConfiguration() OsConfigurationMap {
	if o == nil || o.OsConfiguration == nil {
		var ret OsConfigurationMap
		return ret
	}
	return *o.OsConfiguration
}

// GetOsConfigurationOk returns a tuple with the OsConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerReset) GetOsConfigurationOk() (*OsConfigurationMap, bool) {
	if o == nil || o.OsConfiguration == nil {
		return nil, false
	}
	return o.OsConfiguration, true
}

// HasOsConfiguration returns a boolean if a field has been set.
func (o *ServerReset) HasOsConfiguration() bool {
	if o != nil && o.OsConfiguration != nil {
		return true
	}

	return false
}

// SetOsConfiguration gets a reference to the given OsConfigurationMap and assigns it to the OsConfiguration field.
func (o *ServerReset) SetOsConfiguration(v OsConfigurationMap) {
	o.OsConfiguration = &v
}

func (o ServerReset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstallDefaultSshKeys != nil {
		toSerialize["installDefaultSshKeys"] = o.InstallDefaultSshKeys
	}
	if o.SshKeys != nil {
		toSerialize["sshKeys"] = o.SshKeys
	}
	if o.SshKeyIds != nil {
		toSerialize["sshKeyIds"] = o.SshKeyIds
	}
	if o.OsConfiguration != nil {
		toSerialize["osConfiguration"] = o.OsConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableServerReset struct {
	value *ServerReset
	isSet bool
}

func (v NullableServerReset) Get() *ServerReset {
	return v.value
}

func (v *NullableServerReset) Set(val *ServerReset) {
	v.value = val
	v.isSet = true
}

func (v NullableServerReset) IsSet() bool {
	return v.isSet
}

func (v *NullableServerReset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerReset(val *ServerReset) *NullableServerReset {
	return &NullableServerReset{value: val, isSet: true}
}

func (v NullableServerReset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerReset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}