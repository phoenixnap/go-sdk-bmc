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

// checks if the GpuConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GpuConfiguration{}

// GpuConfiguration The GPU configuration.
type GpuConfiguration struct {
	// The long name of the GPU.
	LongName *string `json:"longName,omitempty"`
	// The number of GPUs.
	Count                *int32 `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GpuConfiguration GpuConfiguration

// NewGpuConfiguration instantiates a new GpuConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGpuConfiguration() *GpuConfiguration {
	this := GpuConfiguration{}
	return &this
}

// NewGpuConfigurationWithDefaults instantiates a new GpuConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGpuConfigurationWithDefaults() *GpuConfiguration {
	this := GpuConfiguration{}
	return &this
}

// GetLongName returns the LongName field value if set, zero value otherwise.
func (o *GpuConfiguration) GetLongName() string {
	if o == nil || IsNil(o.LongName) {
		var ret string
		return ret
	}
	return *o.LongName
}

// GetLongNameOk returns a tuple with the LongName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpuConfiguration) GetLongNameOk() (*string, bool) {
	if o == nil || IsNil(o.LongName) {
		return nil, false
	}
	return o.LongName, true
}

// HasLongName returns a boolean if a field has been set.
func (o *GpuConfiguration) HasLongName() bool {
	if o != nil && !IsNil(o.LongName) {
		return true
	}

	return false
}

// SetLongName gets a reference to the given string and assigns it to the LongName field.
func (o *GpuConfiguration) SetLongName(v string) {
	o.LongName = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GpuConfiguration) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpuConfiguration) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GpuConfiguration) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GpuConfiguration) SetCount(v int32) {
	o.Count = &v
}

func (o GpuConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GpuConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LongName) {
		toSerialize["longName"] = o.LongName
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GpuConfiguration) UnmarshalJSON(data []byte) (err error) {
	varGpuConfiguration := _GpuConfiguration{}

	err = json.Unmarshal(data, &varGpuConfiguration)

	if err != nil {
		return err
	}

	*o = GpuConfiguration(varGpuConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "longName")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGpuConfiguration struct {
	value *GpuConfiguration
	isSet bool
}

func (v NullableGpuConfiguration) Get() *GpuConfiguration {
	return v.value
}

func (v *NullableGpuConfiguration) Set(val *GpuConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableGpuConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableGpuConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGpuConfiguration(val *GpuConfiguration) *NullableGpuConfiguration {
	return &NullableGpuConfiguration{value: val, isSet: true}
}

func (v NullableGpuConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGpuConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
