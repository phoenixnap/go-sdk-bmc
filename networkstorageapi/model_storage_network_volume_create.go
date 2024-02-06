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
	"fmt"
)

// checks if the StorageNetworkVolumeCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetworkVolumeCreate{}

// StorageNetworkVolumeCreate Create Volume.
type StorageNetworkVolumeCreate struct {
	// Volume friendly name.
	Name string `json:"name"`
	// Volume description.
	Description *string `json:"description,omitempty"`
	// Last part of volume's path.
	PathSuffix *string `json:"pathSuffix,omitempty"`
	// Capacity of Volume in GB. Currently only whole numbers and multiples of 1000GB are supported.
	CapacityInGb int32 `json:"capacityInGb"`
	// Tags to set to the resource. To create a new tag or list all the existing tags that you can use, refer to [Tags API](https://developers.phoenixnap.com/docs/tags/1/overview).
	Tags                 []TagAssignmentRequest `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetworkVolumeCreate StorageNetworkVolumeCreate

// NewStorageNetworkVolumeCreate instantiates a new StorageNetworkVolumeCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetworkVolumeCreate(name string, capacityInGb int32) *StorageNetworkVolumeCreate {
	this := StorageNetworkVolumeCreate{}
	this.Name = name
	this.CapacityInGb = capacityInGb
	return &this
}

// NewStorageNetworkVolumeCreateWithDefaults instantiates a new StorageNetworkVolumeCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetworkVolumeCreateWithDefaults() *StorageNetworkVolumeCreate {
	this := StorageNetworkVolumeCreate{}
	return &this
}

// GetName returns the Name field value
func (o *StorageNetworkVolumeCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StorageNetworkVolumeCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StorageNetworkVolumeCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StorageNetworkVolumeCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetworkVolumeCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StorageNetworkVolumeCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StorageNetworkVolumeCreate) SetDescription(v string) {
	o.Description = &v
}

// GetPathSuffix returns the PathSuffix field value if set, zero value otherwise.
func (o *StorageNetworkVolumeCreate) GetPathSuffix() string {
	if o == nil || IsNil(o.PathSuffix) {
		var ret string
		return ret
	}
	return *o.PathSuffix
}

// GetPathSuffixOk returns a tuple with the PathSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetworkVolumeCreate) GetPathSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.PathSuffix) {
		return nil, false
	}
	return o.PathSuffix, true
}

// HasPathSuffix returns a boolean if a field has been set.
func (o *StorageNetworkVolumeCreate) HasPathSuffix() bool {
	if o != nil && !IsNil(o.PathSuffix) {
		return true
	}

	return false
}

// SetPathSuffix gets a reference to the given string and assigns it to the PathSuffix field.
func (o *StorageNetworkVolumeCreate) SetPathSuffix(v string) {
	o.PathSuffix = &v
}

// GetCapacityInGb returns the CapacityInGb field value
func (o *StorageNetworkVolumeCreate) GetCapacityInGb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CapacityInGb
}

// GetCapacityInGbOk returns a tuple with the CapacityInGb field value
// and a boolean to check if the value has been set.
func (o *StorageNetworkVolumeCreate) GetCapacityInGbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityInGb, true
}

// SetCapacityInGb sets field value
func (o *StorageNetworkVolumeCreate) SetCapacityInGb(v int32) {
	o.CapacityInGb = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *StorageNetworkVolumeCreate) GetTags() []TagAssignmentRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []TagAssignmentRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetworkVolumeCreate) GetTagsOk() ([]TagAssignmentRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *StorageNetworkVolumeCreate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagAssignmentRequest and assigns it to the Tags field.
func (o *StorageNetworkVolumeCreate) SetTags(v []TagAssignmentRequest) {
	o.Tags = v
}

func (o StorageNetworkVolumeCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetworkVolumeCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.PathSuffix) {
		toSerialize["pathSuffix"] = o.PathSuffix
	}
	toSerialize["capacityInGb"] = o.CapacityInGb
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetworkVolumeCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"capacityInGb",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varStorageNetworkVolumeCreate := _StorageNetworkVolumeCreate{}

	err = json.Unmarshal(data, &varStorageNetworkVolumeCreate)

	if err != nil {
		return err
	}

	*o = StorageNetworkVolumeCreate(varStorageNetworkVolumeCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "pathSuffix")
		delete(additionalProperties, "capacityInGb")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetworkVolumeCreate struct {
	value *StorageNetworkVolumeCreate
	isSet bool
}

func (v NullableStorageNetworkVolumeCreate) Get() *StorageNetworkVolumeCreate {
	return v.value
}

func (v *NullableStorageNetworkVolumeCreate) Set(val *StorageNetworkVolumeCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetworkVolumeCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetworkVolumeCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetworkVolumeCreate(val *StorageNetworkVolumeCreate) *NullableStorageNetworkVolumeCreate {
	return &NullableStorageNetworkVolumeCreate{value: val, isSet: true}
}

func (v NullableStorageNetworkVolumeCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetworkVolumeCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
