/*
Network Storage API

Create, list, edit, and delete storage networks with the Network Storage API. Use storage networks to expand storage capacity on a private network. <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/bare-metal-cloud-storage' target='_blank'>here</a> </span> <br> <b>All URLs are relative to (https://api.phoenixnap.com/network-storage/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkstorageapi

import (
	"encoding/json"
)

// PermissionsUpdate Update permissions for a volume.
type PermissionsUpdate struct {
	Nfs *NfsPermissionsUpdate `json:"nfs,omitempty"`
}

// NewPermissionsUpdate instantiates a new PermissionsUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsUpdate() *PermissionsUpdate {
	this := PermissionsUpdate{}
	return &this
}

// NewPermissionsUpdateWithDefaults instantiates a new PermissionsUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsUpdateWithDefaults() *PermissionsUpdate {
	this := PermissionsUpdate{}
	return &this
}

// GetNfs returns the Nfs field value if set, zero value otherwise.
func (o *PermissionsUpdate) GetNfs() NfsPermissionsUpdate {
	if o == nil || o.Nfs == nil {
		var ret NfsPermissionsUpdate
		return ret
	}
	return *o.Nfs
}

// GetNfsOk returns a tuple with the Nfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsUpdate) GetNfsOk() (*NfsPermissionsUpdate, bool) {
	if o == nil || o.Nfs == nil {
		return nil, false
	}
	return o.Nfs, true
}

// HasNfs returns a boolean if a field has been set.
func (o *PermissionsUpdate) HasNfs() bool {
	if o != nil && o.Nfs != nil {
		return true
	}

	return false
}

// SetNfs gets a reference to the given NfsPermissionsUpdate and assigns it to the Nfs field.
func (o *PermissionsUpdate) SetNfs(v NfsPermissionsUpdate) {
	o.Nfs = &v
}

func (o PermissionsUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nfs != nil {
		toSerialize["nfs"] = o.Nfs
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionsUpdate struct {
	value *PermissionsUpdate
	isSet bool
}

func (v NullablePermissionsUpdate) Get() *PermissionsUpdate {
	return v.value
}

func (v *NullablePermissionsUpdate) Set(val *PermissionsUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsUpdate(val *PermissionsUpdate) *NullablePermissionsUpdate {
	return &NullablePermissionsUpdate{value: val, isSet: true}
}

func (v NullablePermissionsUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
