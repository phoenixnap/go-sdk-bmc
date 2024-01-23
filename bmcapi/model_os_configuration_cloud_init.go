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

// OsConfigurationCloudInit Cloud-init configuration details.
type OsConfigurationCloudInit struct {
	// (Write-only) User data for the <a href='https://cloudinit.readthedocs.io/en/latest/' target='_blank'>cloud-init</a> configuration in base64 encoding. NoCloud format is supported. Follow the <a href='https://phoenixnap.com/kb/bmc-cloud-init' target='_blank'>instructions</a> on how to provision a server using cloud-init. Only ubuntu/bionic, ubuntu/focal and ubuntu/jammy are supported. User data will not be stored and cannot be retrieved once you deploy the server. Copy and save it for future reference.
	UserData *string `json:"userData,omitempty"`
}

// NewOsConfigurationCloudInit instantiates a new OsConfigurationCloudInit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsConfigurationCloudInit() *OsConfigurationCloudInit {
	this := OsConfigurationCloudInit{}
	return &this
}

// NewOsConfigurationCloudInitWithDefaults instantiates a new OsConfigurationCloudInit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsConfigurationCloudInitWithDefaults() *OsConfigurationCloudInit {
	this := OsConfigurationCloudInit{}
	return &this
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *OsConfigurationCloudInit) GetUserData() string {
	if o == nil || o.UserData == nil {
		var ret string
		return ret
	}
	return *o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationCloudInit) GetUserDataOk() (*string, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *OsConfigurationCloudInit) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given string and assigns it to the UserData field.
func (o *OsConfigurationCloudInit) SetUserData(v string) {
	o.UserData = &v
}

func (o OsConfigurationCloudInit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserData != nil {
		toSerialize["userData"] = o.UserData
	}
	return json.Marshal(toSerialize)
}

type NullableOsConfigurationCloudInit struct {
	value *OsConfigurationCloudInit
	isSet bool
}

func (v NullableOsConfigurationCloudInit) Get() *OsConfigurationCloudInit {
	return v.value
}

func (v *NullableOsConfigurationCloudInit) Set(val *OsConfigurationCloudInit) {
	v.value = val
	v.isSet = true
}

func (v NullableOsConfigurationCloudInit) IsSet() bool {
	return v.isSet
}

func (v *NullableOsConfigurationCloudInit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsConfigurationCloudInit(val *OsConfigurationCloudInit) *NullableOsConfigurationCloudInit {
	return &NullableOsConfigurationCloudInit{value: val, isSet: true}
}

func (v NullableOsConfigurationCloudInit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsConfigurationCloudInit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
