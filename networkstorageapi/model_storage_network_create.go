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

// StorageNetworkCreate Create Storage Network.
type StorageNetworkCreate struct {
	// Storage network friendly name.
	Name string `json:"name"`
	// Storage network description.
	Description *string `json:"description,omitempty"`
	// Location of storage network. Currently this field should be set to `PHX` or `ASH`.
	Location string `json:"location"`
	// Volume to be created alongside storage. Currently only 1 volume is supported.
	Volumes []StorageNetworkVolumeCreate `json:"volumes"`
	// Custom Client VLAN that the Storage Network will be set to.
	ClientVlan *int32 `json:"clientVlan,omitempty"`
}

// NewStorageNetworkCreate instantiates a new StorageNetworkCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetworkCreate(name string, location string, volumes []StorageNetworkVolumeCreate) *StorageNetworkCreate {
	this := StorageNetworkCreate{}
	this.Name = name
	this.Location = location
	this.Volumes = volumes
	return &this
}

// NewStorageNetworkCreateWithDefaults instantiates a new StorageNetworkCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetworkCreateWithDefaults() *StorageNetworkCreate {
	this := StorageNetworkCreate{}
	return &this
}

// GetName returns the Name field value
func (o *StorageNetworkCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StorageNetworkCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StorageNetworkCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StorageNetworkCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetworkCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StorageNetworkCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StorageNetworkCreate) SetDescription(v string) {
	o.Description = &v
}

// GetLocation returns the Location field value
func (o *StorageNetworkCreate) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *StorageNetworkCreate) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *StorageNetworkCreate) SetLocation(v string) {
	o.Location = v
}

// GetVolumes returns the Volumes field value
func (o *StorageNetworkCreate) GetVolumes() []StorageNetworkVolumeCreate {
	if o == nil {
		var ret []StorageNetworkVolumeCreate
		return ret
	}

	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value
// and a boolean to check if the value has been set.
func (o *StorageNetworkCreate) GetVolumesOk() ([]StorageNetworkVolumeCreate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Volumes, true
}

// SetVolumes sets field value
func (o *StorageNetworkCreate) SetVolumes(v []StorageNetworkVolumeCreate) {
	o.Volumes = v
}

// GetClientVlan returns the ClientVlan field value if set, zero value otherwise.
func (o *StorageNetworkCreate) GetClientVlan() int32 {
	if o == nil || o.ClientVlan == nil {
		var ret int32
		return ret
	}
	return *o.ClientVlan
}

// GetClientVlanOk returns a tuple with the ClientVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetworkCreate) GetClientVlanOk() (*int32, bool) {
	if o == nil || o.ClientVlan == nil {
		return nil, false
	}
	return o.ClientVlan, true
}

// HasClientVlan returns a boolean if a field has been set.
func (o *StorageNetworkCreate) HasClientVlan() bool {
	if o != nil && o.ClientVlan != nil {
		return true
	}

	return false
}

// SetClientVlan gets a reference to the given int32 and assigns it to the ClientVlan field.
func (o *StorageNetworkCreate) SetClientVlan(v int32) {
	o.ClientVlan = &v
}

func (o StorageNetworkCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["volumes"] = o.Volumes
	}
	if o.ClientVlan != nil {
		toSerialize["clientVlan"] = o.ClientVlan
	}
	return json.Marshal(toSerialize)
}

type NullableStorageNetworkCreate struct {
	value *StorageNetworkCreate
	isSet bool
}

func (v NullableStorageNetworkCreate) Get() *StorageNetworkCreate {
	return v.value
}

func (v *NullableStorageNetworkCreate) Set(val *StorageNetworkCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetworkCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetworkCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetworkCreate(val *StorageNetworkCreate) *NullableStorageNetworkCreate {
	return &NullableStorageNetworkCreate{value: val, isSet: true}
}

func (v NullableStorageNetworkCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetworkCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
