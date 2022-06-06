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
	"time"
)

// SshKey SSH Key.
type SshKey struct {
	// The unique identifier of the SSH Key.
	Id string `json:"id"`
	// Keys marked as default are always included on server creation and reset unless toggled off in creation/reset request.
	Default bool `json:"default"`
	// Friendly SSH key name to represent an SSH key.
	Name string `json:"name"`
	// SSH Key value.
	Key string `json:"key"`
	// SSH key auto-generated SHA-256 fingerprint.
	Fingerprint string `json:"fingerprint"`
	// Date and time of creation.
	CreatedOn time.Time `json:"createdOn"`
	// Date and time of last update.
	LastUpdatedOn time.Time `json:"lastUpdatedOn"`
}

// NewSshKey instantiates a new SshKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshKey(id string, default_ bool, name string, key string, fingerprint string, createdOn time.Time, lastUpdatedOn time.Time) *SshKey {
	this := SshKey{}
	this.Id = id
	this.Default = default_
	this.Name = name
	this.Key = key
	this.Fingerprint = fingerprint
	this.CreatedOn = createdOn
	this.LastUpdatedOn = lastUpdatedOn
	return &this
}

// NewSshKeyWithDefaults instantiates a new SshKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshKeyWithDefaults() *SshKey {
	this := SshKey{}
	return &this
}

// GetId returns the Id field value
func (o *SshKey) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SshKey) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SshKey) SetId(v string) {
	o.Id = v
}

// GetDefault returns the Default field value
func (o *SshKey) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *SshKey) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *SshKey) SetDefault(v bool) {
	o.Default = v
}

// GetName returns the Name field value
func (o *SshKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SshKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SshKey) SetName(v string) {
	o.Name = v
}

// GetKey returns the Key field value
func (o *SshKey) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SshKey) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SshKey) SetKey(v string) {
	o.Key = v
}

// GetFingerprint returns the Fingerprint field value
func (o *SshKey) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *SshKey) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value
func (o *SshKey) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *SshKey) GetCreatedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *SshKey) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *SshKey) SetCreatedOn(v time.Time) {
	o.CreatedOn = v
}

// GetLastUpdatedOn returns the LastUpdatedOn field value
func (o *SshKey) GetLastUpdatedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedOn
}

// GetLastUpdatedOnOk returns a tuple with the LastUpdatedOn field value
// and a boolean to check if the value has been set.
func (o *SshKey) GetLastUpdatedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedOn, true
}

// SetLastUpdatedOn sets field value
func (o *SshKey) SetLastUpdatedOn(v time.Time) {
	o.LastUpdatedOn = v
}

func (o SshKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["default"] = o.Default
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if true {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if true {
		toSerialize["lastUpdatedOn"] = o.LastUpdatedOn
	}
	return json.Marshal(toSerialize)
}

type NullableSshKey struct {
	value *SshKey
	isSet bool
}

func (v NullableSshKey) Get() *SshKey {
	return v.value
}

func (v *NullableSshKey) Set(val *SshKey) {
	v.value = val
	v.isSet = true
}

func (v NullableSshKey) IsSet() bool {
	return v.isSet
}

func (v *NullableSshKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshKey(val *SshKey) *NullableSshKey {
	return &NullableSshKey{value: val, isSet: true}
}

func (v NullableSshKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
