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
	"time"
)

// checks if the Volume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Volume{}

// Volume Volume for a storage network.
type Volume struct {
	// Volume ID.
	Id *string `json:"id,omitempty"`
	// Volume friendly name.
	Name *string `json:"name,omitempty"`
	// Volume description.
	Description *string `json:"description,omitempty"`
	// Volume's full path. It is in form of `/{volumeId}/pathSuffix`'.
	Path *string `json:"path,omitempty"`
	// Last part of volume's path.
	PathSuffix *string `json:"pathSuffix,omitempty"`
	// Maximum capacity in GB.
	CapacityInGb *int32 `json:"capacityInGb,omitempty"`
	// Used capacity in GB, updated periodically.
	UsedCapacityInGb *int32 `json:"usedCapacityInGb,omitempty"`
	// File system protocol. Currently this field should be set to `NFS`.
	Protocol  *string    `json:"protocol,omitempty"`
	Status    *Status    `json:"status,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	// Date and time of the initial request for volume deletion.
	DeleteRequestedOn *time.Time   `json:"deleteRequestedOn,omitempty"`
	Permissions       *Permissions `json:"permissions,omitempty"`
	// The tags assigned if any.
	Tags                 []TagAssignment `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Volume Volume

// NewVolume instantiates a new Volume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolume() *Volume {
	this := Volume{}
	return &this
}

// NewVolumeWithDefaults instantiates a new Volume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeWithDefaults() *Volume {
	this := Volume{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Volume) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Volume) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Volume) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Volume) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Volume) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Volume) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Volume) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Volume) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Volume) SetDescription(v string) {
	o.Description = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Volume) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Volume) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Volume) SetPath(v string) {
	o.Path = &v
}

// GetPathSuffix returns the PathSuffix field value if set, zero value otherwise.
func (o *Volume) GetPathSuffix() string {
	if o == nil || IsNil(o.PathSuffix) {
		var ret string
		return ret
	}
	return *o.PathSuffix
}

// GetPathSuffixOk returns a tuple with the PathSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetPathSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.PathSuffix) {
		return nil, false
	}
	return o.PathSuffix, true
}

// HasPathSuffix returns a boolean if a field has been set.
func (o *Volume) HasPathSuffix() bool {
	if o != nil && !IsNil(o.PathSuffix) {
		return true
	}

	return false
}

// SetPathSuffix gets a reference to the given string and assigns it to the PathSuffix field.
func (o *Volume) SetPathSuffix(v string) {
	o.PathSuffix = &v
}

// GetCapacityInGb returns the CapacityInGb field value if set, zero value otherwise.
func (o *Volume) GetCapacityInGb() int32 {
	if o == nil || IsNil(o.CapacityInGb) {
		var ret int32
		return ret
	}
	return *o.CapacityInGb
}

// GetCapacityInGbOk returns a tuple with the CapacityInGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetCapacityInGbOk() (*int32, bool) {
	if o == nil || IsNil(o.CapacityInGb) {
		return nil, false
	}
	return o.CapacityInGb, true
}

// HasCapacityInGb returns a boolean if a field has been set.
func (o *Volume) HasCapacityInGb() bool {
	if o != nil && !IsNil(o.CapacityInGb) {
		return true
	}

	return false
}

// SetCapacityInGb gets a reference to the given int32 and assigns it to the CapacityInGb field.
func (o *Volume) SetCapacityInGb(v int32) {
	o.CapacityInGb = &v
}

// GetUsedCapacityInGb returns the UsedCapacityInGb field value if set, zero value otherwise.
func (o *Volume) GetUsedCapacityInGb() int32 {
	if o == nil || IsNil(o.UsedCapacityInGb) {
		var ret int32
		return ret
	}
	return *o.UsedCapacityInGb
}

// GetUsedCapacityInGbOk returns a tuple with the UsedCapacityInGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetUsedCapacityInGbOk() (*int32, bool) {
	if o == nil || IsNil(o.UsedCapacityInGb) {
		return nil, false
	}
	return o.UsedCapacityInGb, true
}

// HasUsedCapacityInGb returns a boolean if a field has been set.
func (o *Volume) HasUsedCapacityInGb() bool {
	if o != nil && !IsNil(o.UsedCapacityInGb) {
		return true
	}

	return false
}

// SetUsedCapacityInGb gets a reference to the given int32 and assigns it to the UsedCapacityInGb field.
func (o *Volume) SetUsedCapacityInGb(v int32) {
	o.UsedCapacityInGb = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *Volume) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *Volume) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *Volume) SetProtocol(v string) {
	o.Protocol = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Volume) GetStatus() Status {
	if o == nil || IsNil(o.Status) {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetStatusOk() (*Status, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Volume) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *Volume) SetStatus(v Status) {
	o.Status = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Volume) GetCreatedOn() time.Time {
	if o == nil || IsNil(o.CreatedOn) {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Volume) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Volume) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetDeleteRequestedOn returns the DeleteRequestedOn field value if set, zero value otherwise.
func (o *Volume) GetDeleteRequestedOn() time.Time {
	if o == nil || IsNil(o.DeleteRequestedOn) {
		var ret time.Time
		return ret
	}
	return *o.DeleteRequestedOn
}

// GetDeleteRequestedOnOk returns a tuple with the DeleteRequestedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetDeleteRequestedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeleteRequestedOn) {
		return nil, false
	}
	return o.DeleteRequestedOn, true
}

// HasDeleteRequestedOn returns a boolean if a field has been set.
func (o *Volume) HasDeleteRequestedOn() bool {
	if o != nil && !IsNil(o.DeleteRequestedOn) {
		return true
	}

	return false
}

// SetDeleteRequestedOn gets a reference to the given time.Time and assigns it to the DeleteRequestedOn field.
func (o *Volume) SetDeleteRequestedOn(v time.Time) {
	o.DeleteRequestedOn = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *Volume) GetPermissions() Permissions {
	if o == nil || IsNil(o.Permissions) {
		var ret Permissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetPermissionsOk() (*Permissions, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *Volume) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given Permissions and assigns it to the Permissions field.
func (o *Volume) SetPermissions(v Permissions) {
	o.Permissions = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Volume) GetTags() []TagAssignment {
	if o == nil || IsNil(o.Tags) {
		var ret []TagAssignment
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetTagsOk() ([]TagAssignment, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Volume) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagAssignment and assigns it to the Tags field.
func (o *Volume) SetTags(v []TagAssignment) {
	o.Tags = v
}

func (o Volume) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Volume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.PathSuffix) {
		toSerialize["pathSuffix"] = o.PathSuffix
	}
	if !IsNil(o.CapacityInGb) {
		toSerialize["capacityInGb"] = o.CapacityInGb
	}
	if !IsNil(o.UsedCapacityInGb) {
		toSerialize["usedCapacityInGb"] = o.UsedCapacityInGb
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CreatedOn) {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if !IsNil(o.DeleteRequestedOn) {
		toSerialize["deleteRequestedOn"] = o.DeleteRequestedOn
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Volume) UnmarshalJSON(data []byte) (err error) {
	varVolume := _Volume{}

	err = json.Unmarshal(data, &varVolume)

	if err != nil {
		return err
	}

	*o = Volume(varVolume)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "path")
		delete(additionalProperties, "pathSuffix")
		delete(additionalProperties, "capacityInGb")
		delete(additionalProperties, "usedCapacityInGb")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "status")
		delete(additionalProperties, "createdOn")
		delete(additionalProperties, "deleteRequestedOn")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVolume struct {
	value *Volume
	isSet bool
}

func (v NullableVolume) Get() *Volume {
	return v.value
}

func (v *NullableVolume) Set(val *Volume) {
	v.value = val
	v.isSet = true
}

func (v NullableVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolume(val *Volume) *NullableVolume {
	return &NullableVolume{value: val, isSet: true}
}

func (v NullableVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
