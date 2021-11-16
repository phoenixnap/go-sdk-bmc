/*
Rancher Solution API

Create and manage Rancher clusters. </br></br>**All URLs are relative to (https://api.phoenixnap.com/solutions/rancher/v0)**

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ranchersolutionapi

import (
	"encoding/json"
)

// SshConfig Configuration defining which public SSH keys are pre-installed as authorized on the server. Any manual configuration done on the server after installation is not reflected by this configuration.
type SshConfig struct {
	// Define whether public keys marked as default should be installed on this node. These are public keys that were already recorded on this system. Use <a href=\"https://developers.phoenixnap.com/docs/bmc/1/routes/ssh-keys/get\" target=\"_blank\">GET /ssh-keys</a> to retrieve a list of possible values.
	InstallDefaultKeys *bool `json:"installDefaultKeys,omitempty"`
	// List of public SSH keys.
	Keys *[]string `json:"keys,omitempty"`
	// List of public SSH key identifiers. These are public keys that were already recorded on this system. Use <a href=\"https://developers.phoenixnap.com/docs/bmc/1/routes/ssh-keys/get\" target=\"_blank\">GET /ssh-keys</a> to retrieve a list of possible values.
	KeyIds *[]string `json:"keyIds,omitempty"`
}

// NewSshConfig instantiates a new SshConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshConfig() *SshConfig {
	this := SshConfig{}
	var installDefaultKeys bool = true
	this.InstallDefaultKeys = &installDefaultKeys
	return &this
}

// NewSshConfigWithDefaults instantiates a new SshConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshConfigWithDefaults() *SshConfig {
	this := SshConfig{}
	var installDefaultKeys bool = true
	this.InstallDefaultKeys = &installDefaultKeys
	return &this
}

// GetInstallDefaultKeys returns the InstallDefaultKeys field value if set, zero value otherwise.
func (o *SshConfig) GetInstallDefaultKeys() bool {
	if o == nil || o.InstallDefaultKeys == nil {
		var ret bool
		return ret
	}
	return *o.InstallDefaultKeys
}

// GetInstallDefaultKeysOk returns a tuple with the InstallDefaultKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshConfig) GetInstallDefaultKeysOk() (*bool, bool) {
	if o == nil || o.InstallDefaultKeys == nil {
		return nil, false
	}
	return o.InstallDefaultKeys, true
}

// HasInstallDefaultKeys returns a boolean if a field has been set.
func (o *SshConfig) HasInstallDefaultKeys() bool {
	if o != nil && o.InstallDefaultKeys != nil {
		return true
	}

	return false
}

// SetInstallDefaultKeys gets a reference to the given bool and assigns it to the InstallDefaultKeys field.
func (o *SshConfig) SetInstallDefaultKeys(v bool) {
	o.InstallDefaultKeys = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *SshConfig) GetKeys() []string {
	if o == nil || o.Keys == nil {
		var ret []string
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshConfig) GetKeysOk() (*[]string, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *SshConfig) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *SshConfig) SetKeys(v []string) {
	o.Keys = &v
}

// GetKeyIds returns the KeyIds field value if set, zero value otherwise.
func (o *SshConfig) GetKeyIds() []string {
	if o == nil || o.KeyIds == nil {
		var ret []string
		return ret
	}
	return *o.KeyIds
}

// GetKeyIdsOk returns a tuple with the KeyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshConfig) GetKeyIdsOk() (*[]string, bool) {
	if o == nil || o.KeyIds == nil {
		return nil, false
	}
	return o.KeyIds, true
}

// HasKeyIds returns a boolean if a field has been set.
func (o *SshConfig) HasKeyIds() bool {
	if o != nil && o.KeyIds != nil {
		return true
	}

	return false
}

// SetKeyIds gets a reference to the given []string and assigns it to the KeyIds field.
func (o *SshConfig) SetKeyIds(v []string) {
	o.KeyIds = &v
}

func (o SshConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstallDefaultKeys != nil {
		toSerialize["installDefaultKeys"] = o.InstallDefaultKeys
	}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	if o.KeyIds != nil {
		toSerialize["keyIds"] = o.KeyIds
	}
	return json.Marshal(toSerialize)
}

type NullableSshConfig struct {
	value *SshConfig
	isSet bool
}

func (v NullableSshConfig) Get() *SshConfig {
	return v.value
}

func (v *NullableSshConfig) Set(val *SshConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSshConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSshConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshConfig(val *SshConfig) *NullableSshConfig {
	return &NullableSshConfig{value: val, isSet: true}
}

func (v NullableSshConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
