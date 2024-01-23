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

// checks if the OsConfigurationMap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsConfigurationMap{}

// OsConfigurationMap OS specific configuration properties.
type OsConfigurationMap struct {
	Windows *OsConfigurationWindows    `json:"windows,omitempty"`
	Esxi    *OsConfigurationMapEsxi    `json:"esxi,omitempty"`
	Proxmox *OsConfigurationMapProxmox `json:"proxmox,omitempty"`
}

// NewOsConfigurationMap instantiates a new OsConfigurationMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsConfigurationMap() *OsConfigurationMap {
	this := OsConfigurationMap{}
	return &this
}

// NewOsConfigurationMapWithDefaults instantiates a new OsConfigurationMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsConfigurationMapWithDefaults() *OsConfigurationMap {
	this := OsConfigurationMap{}
	return &this
}

// GetWindows returns the Windows field value if set, zero value otherwise.
func (o *OsConfigurationMap) GetWindows() OsConfigurationWindows {
	if o == nil || IsNil(o.Windows) {
		var ret OsConfigurationWindows
		return ret
	}
	return *o.Windows
}

// GetWindowsOk returns a tuple with the Windows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationMap) GetWindowsOk() (*OsConfigurationWindows, bool) {
	if o == nil || IsNil(o.Windows) {
		return nil, false
	}
	return o.Windows, true
}

// HasWindows returns a boolean if a field has been set.
func (o *OsConfigurationMap) HasWindows() bool {
	if o != nil && !IsNil(o.Windows) {
		return true
	}

	return false
}

// SetWindows gets a reference to the given OsConfigurationWindows and assigns it to the Windows field.
func (o *OsConfigurationMap) SetWindows(v OsConfigurationWindows) {
	o.Windows = &v
}

// GetEsxi returns the Esxi field value if set, zero value otherwise.
func (o *OsConfigurationMap) GetEsxi() OsConfigurationMapEsxi {
	if o == nil || IsNil(o.Esxi) {
		var ret OsConfigurationMapEsxi
		return ret
	}
	return *o.Esxi
}

// GetEsxiOk returns a tuple with the Esxi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationMap) GetEsxiOk() (*OsConfigurationMapEsxi, bool) {
	if o == nil || IsNil(o.Esxi) {
		return nil, false
	}
	return o.Esxi, true
}

// HasEsxi returns a boolean if a field has been set.
func (o *OsConfigurationMap) HasEsxi() bool {
	if o != nil && !IsNil(o.Esxi) {
		return true
	}

	return false
}

// SetEsxi gets a reference to the given OsConfigurationMapEsxi and assigns it to the Esxi field.
func (o *OsConfigurationMap) SetEsxi(v OsConfigurationMapEsxi) {
	o.Esxi = &v
}

// GetProxmox returns the Proxmox field value if set, zero value otherwise.
func (o *OsConfigurationMap) GetProxmox() OsConfigurationMapProxmox {
	if o == nil || IsNil(o.Proxmox) {
		var ret OsConfigurationMapProxmox
		return ret
	}
	return *o.Proxmox
}

// GetProxmoxOk returns a tuple with the Proxmox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationMap) GetProxmoxOk() (*OsConfigurationMapProxmox, bool) {
	if o == nil || IsNil(o.Proxmox) {
		return nil, false
	}
	return o.Proxmox, true
}

// HasProxmox returns a boolean if a field has been set.
func (o *OsConfigurationMap) HasProxmox() bool {
	if o != nil && !IsNil(o.Proxmox) {
		return true
	}

	return false
}

// SetProxmox gets a reference to the given OsConfigurationMapProxmox and assigns it to the Proxmox field.
func (o *OsConfigurationMap) SetProxmox(v OsConfigurationMapProxmox) {
	o.Proxmox = &v
}

func (o OsConfigurationMap) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsConfigurationMap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Windows) {
		toSerialize["windows"] = o.Windows
	}
	if !IsNil(o.Esxi) {
		toSerialize["esxi"] = o.Esxi
	}
	if !IsNil(o.Proxmox) {
		toSerialize["proxmox"] = o.Proxmox
	}
	return toSerialize, nil
}

type NullableOsConfigurationMap struct {
	value *OsConfigurationMap
	isSet bool
}

func (v NullableOsConfigurationMap) Get() *OsConfigurationMap {
	return v.value
}

func (v *NullableOsConfigurationMap) Set(val *OsConfigurationMap) {
	v.value = val
	v.isSet = true
}

func (v NullableOsConfigurationMap) IsSet() bool {
	return v.isSet
}

func (v *NullableOsConfigurationMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsConfigurationMap(val *OsConfigurationMap) *NullableOsConfigurationMap {
	return &NullableOsConfigurationMap{value: val, isSet: true}
}

func (v NullableOsConfigurationMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsConfigurationMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
