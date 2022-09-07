/*
Network Storage API

Create, list, edit, and delete storage networks with the Network Storage API. Use storage networks to expand storage capacity on a private network. <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='' target='_blank'>here</a> </span> <br> <b>All URLs are relative to (https://api-dev.phoenixnap.com/network-storage/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkstorageapi

import (
	"encoding/json"
)

// NfsPermissions NFS specific permissions on a volume.
type NfsPermissions struct {
	// Read/Write access.
	ReadWrite []string `json:"readWrite,omitempty"`
	// Read only access.
	ReadOnly []string `json:"readOnly,omitempty"`
	// Root squash permission.
	RootSquash []string `json:"rootSquash,omitempty"`
	// No squash permission.
	NoSquash []string `json:"noSquash,omitempty"`
	// All squash permission.
	AllSquash []string `json:"allSquash,omitempty"`
}

// NewNfsPermissions instantiates a new NfsPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfsPermissions() *NfsPermissions {
	this := NfsPermissions{}
	return &this
}

// NewNfsPermissionsWithDefaults instantiates a new NfsPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfsPermissionsWithDefaults() *NfsPermissions {
	this := NfsPermissions{}
	return &this
}

// GetReadWrite returns the ReadWrite field value if set, zero value otherwise.
func (o *NfsPermissions) GetReadWrite() []string {
	if o == nil || o.ReadWrite == nil {
		var ret []string
		return ret
	}
	return o.ReadWrite
}

// GetReadWriteOk returns a tuple with the ReadWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsPermissions) GetReadWriteOk() ([]string, bool) {
	if o == nil || o.ReadWrite == nil {
		return nil, false
	}
	return o.ReadWrite, true
}

// HasReadWrite returns a boolean if a field has been set.
func (o *NfsPermissions) HasReadWrite() bool {
	if o != nil && o.ReadWrite != nil {
		return true
	}

	return false
}

// SetReadWrite gets a reference to the given []string and assigns it to the ReadWrite field.
func (o *NfsPermissions) SetReadWrite(v []string) {
	o.ReadWrite = v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *NfsPermissions) GetReadOnly() []string {
	if o == nil || o.ReadOnly == nil {
		var ret []string
		return ret
	}
	return o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsPermissions) GetReadOnlyOk() ([]string, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *NfsPermissions) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given []string and assigns it to the ReadOnly field.
func (o *NfsPermissions) SetReadOnly(v []string) {
	o.ReadOnly = v
}

// GetRootSquash returns the RootSquash field value if set, zero value otherwise.
func (o *NfsPermissions) GetRootSquash() []string {
	if o == nil || o.RootSquash == nil {
		var ret []string
		return ret
	}
	return o.RootSquash
}

// GetRootSquashOk returns a tuple with the RootSquash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsPermissions) GetRootSquashOk() ([]string, bool) {
	if o == nil || o.RootSquash == nil {
		return nil, false
	}
	return o.RootSquash, true
}

// HasRootSquash returns a boolean if a field has been set.
func (o *NfsPermissions) HasRootSquash() bool {
	if o != nil && o.RootSquash != nil {
		return true
	}

	return false
}

// SetRootSquash gets a reference to the given []string and assigns it to the RootSquash field.
func (o *NfsPermissions) SetRootSquash(v []string) {
	o.RootSquash = v
}

// GetNoSquash returns the NoSquash field value if set, zero value otherwise.
func (o *NfsPermissions) GetNoSquash() []string {
	if o == nil || o.NoSquash == nil {
		var ret []string
		return ret
	}
	return o.NoSquash
}

// GetNoSquashOk returns a tuple with the NoSquash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsPermissions) GetNoSquashOk() ([]string, bool) {
	if o == nil || o.NoSquash == nil {
		return nil, false
	}
	return o.NoSquash, true
}

// HasNoSquash returns a boolean if a field has been set.
func (o *NfsPermissions) HasNoSquash() bool {
	if o != nil && o.NoSquash != nil {
		return true
	}

	return false
}

// SetNoSquash gets a reference to the given []string and assigns it to the NoSquash field.
func (o *NfsPermissions) SetNoSquash(v []string) {
	o.NoSquash = v
}

// GetAllSquash returns the AllSquash field value if set, zero value otherwise.
func (o *NfsPermissions) GetAllSquash() []string {
	if o == nil || o.AllSquash == nil {
		var ret []string
		return ret
	}
	return o.AllSquash
}

// GetAllSquashOk returns a tuple with the AllSquash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsPermissions) GetAllSquashOk() ([]string, bool) {
	if o == nil || o.AllSquash == nil {
		return nil, false
	}
	return o.AllSquash, true
}

// HasAllSquash returns a boolean if a field has been set.
func (o *NfsPermissions) HasAllSquash() bool {
	if o != nil && o.AllSquash != nil {
		return true
	}

	return false
}

// SetAllSquash gets a reference to the given []string and assigns it to the AllSquash field.
func (o *NfsPermissions) SetAllSquash(v []string) {
	o.AllSquash = v
}

func (o NfsPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReadWrite != nil {
		toSerialize["readWrite"] = o.ReadWrite
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if o.RootSquash != nil {
		toSerialize["rootSquash"] = o.RootSquash
	}
	if o.NoSquash != nil {
		toSerialize["noSquash"] = o.NoSquash
	}
	if o.AllSquash != nil {
		toSerialize["allSquash"] = o.AllSquash
	}
	return json.Marshal(toSerialize)
}

type NullableNfsPermissions struct {
	value *NfsPermissions
	isSet bool
}

func (v NullableNfsPermissions) Get() *NfsPermissions {
	return v.value
}

func (v *NullableNfsPermissions) Set(val *NfsPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableNfsPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableNfsPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfsPermissions(val *NfsPermissions) *NullableNfsPermissions {
	return &NullableNfsPermissions{value: val, isSet: true}
}

func (v NullableNfsPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfsPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
