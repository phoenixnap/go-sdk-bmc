/*
Rancher Solution API

Simplify enterprise-grade Kubernetes cluster operations and management with Rancher on Bare Metal Cloud. Deploy Kubernetes clusters using a few API calls.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/rancher-bmc-integration-kubernetes' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/solutions/rancher/v1beta)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ranchersolutionapi

import (
	"encoding/json"
	"fmt"
)

// checks if the WorkloadClusterConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkloadClusterConfig{}

// WorkloadClusterConfig (Write-only) Workload cluster configuration parameters.
type WorkloadClusterConfig struct {
	// The name of the workload cluster. This field is autogenerated if not provided.
	Name *string `json:"name,omitempty"`
	// Number of configured servers. Currently only server counts of 1 and 3 are possible.
	ServerCount *int32 `json:"serverCount,omitempty"`
	// Node server type. Cannot be changed once the cluster is created. Currently this field should be set to either `s0.d1.small`, `s0.d1.medium`, `s1.c1.small`, `s1.c1.medium`, `s1.c2.medium`, `s1.c2.large`, `s1.e1.small`, `s1.e1.medium`, `s1.e1.large`, `s2.c1.small`, `s2.c1.medium`, `s2.c1.large`, `s2.c2.small`, `s2.c2.medium`, `s2.c2.large`, `d1.c1.small`, `d1.c2.small`, `d1.c3.small`, `d1.c4.small`, `d1.c1.medium`, `d1.c2.medium`, `d1.c3.medium`, `d1.c4.medium`, `d1.c1.large`, `d1.c2.large`, `d1.c3.large`, `d1.c4.large`, `d1.m1.medium`, `d1.m2.medium`, `d1.m3.medium`, `d1.m4.medium`, `d2.c1.medium`, `d2.c2.medium`, `d2.c3.medium`, `d2.c4.medium`, `d2.c5.medium`, `d2.c1.large`, `d2.c2.large`, `d2.c3.large`, `d2.c4.large`, `d2.c5.large`, `d2.m1.xlarge`, `d2.m2.xxlarge`, `d2.m3.xlarge`, `d2.m4.xlarge`, `d2.m5.xlarge`, `d2.c4.db1.pliops1`, `d3.m4.xlarge`, `d3.m5.xlarge`, `d3.m6.xlarge`, `a1.c5.large`, `d3.s5.xlarge`, `d3.m4.xxlarge`, `d3.m5.xxlarge`, `d3.m6.xxlarge`, `s3.c3.medium`, `s3.c3.large`, `d3.c4.medium`, `d3.c5.medium`, `d3.c6.medium`, `d3.c1.large`, `d3.c2.large`, `d3.c3.large`, `d3.m1.xlarge`, `d3.m2.xlarge`, `d3.m3.xlarge`, `d3.g2.c1.xlarge`, `d3.g2.c2.xlarge` or `d3.g2.c3.xlarge`.
	ServerType string `json:"serverType"`
	// Workload cluster location. Cannot be changed once cluster is created. Currently this field should be set to `PHX`, `ASH`, `SGP`, `NLD`, `CHI`, `SEA` or `AUS`.
	Location             string `json:"location"`
	AdditionalProperties map[string]interface{}
}

type _WorkloadClusterConfig WorkloadClusterConfig

// NewWorkloadClusterConfig instantiates a new WorkloadClusterConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadClusterConfig(serverType string, location string) *WorkloadClusterConfig {
	this := WorkloadClusterConfig{}
	var serverCount int32 = 1
	this.ServerCount = &serverCount
	this.ServerType = serverType
	this.Location = location
	return &this
}

// NewWorkloadClusterConfigWithDefaults instantiates a new WorkloadClusterConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadClusterConfigWithDefaults() *WorkloadClusterConfig {
	this := WorkloadClusterConfig{}
	var serverCount int32 = 1
	this.ServerCount = &serverCount
	var serverType string = "s0.d1.small"
	this.ServerType = serverType
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkloadClusterConfig) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadClusterConfig) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkloadClusterConfig) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkloadClusterConfig) SetName(v string) {
	o.Name = &v
}

// GetServerCount returns the ServerCount field value if set, zero value otherwise.
func (o *WorkloadClusterConfig) GetServerCount() int32 {
	if o == nil || IsNil(o.ServerCount) {
		var ret int32
		return ret
	}
	return *o.ServerCount
}

// GetServerCountOk returns a tuple with the ServerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadClusterConfig) GetServerCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerCount) {
		return nil, false
	}
	return o.ServerCount, true
}

// HasServerCount returns a boolean if a field has been set.
func (o *WorkloadClusterConfig) HasServerCount() bool {
	if o != nil && !IsNil(o.ServerCount) {
		return true
	}

	return false
}

// SetServerCount gets a reference to the given int32 and assigns it to the ServerCount field.
func (o *WorkloadClusterConfig) SetServerCount(v int32) {
	o.ServerCount = &v
}

// GetServerType returns the ServerType field value
func (o *WorkloadClusterConfig) GetServerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value
// and a boolean to check if the value has been set.
func (o *WorkloadClusterConfig) GetServerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerType, true
}

// SetServerType sets field value
func (o *WorkloadClusterConfig) SetServerType(v string) {
	o.ServerType = v
}

// GetLocation returns the Location field value
func (o *WorkloadClusterConfig) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *WorkloadClusterConfig) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *WorkloadClusterConfig) SetLocation(v string) {
	o.Location = v
}

func (o WorkloadClusterConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkloadClusterConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ServerCount) {
		toSerialize["serverCount"] = o.ServerCount
	}
	toSerialize["serverType"] = o.ServerType
	toSerialize["location"] = o.Location

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkloadClusterConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serverType",
		"location",
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

	varWorkloadClusterConfig := _WorkloadClusterConfig{}

	err = json.Unmarshal(data, &varWorkloadClusterConfig)

	if err != nil {
		return err
	}

	*o = WorkloadClusterConfig(varWorkloadClusterConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "serverCount")
		delete(additionalProperties, "serverType")
		delete(additionalProperties, "location")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkloadClusterConfig struct {
	value *WorkloadClusterConfig
	isSet bool
}

func (v NullableWorkloadClusterConfig) Get() *WorkloadClusterConfig {
	return v.value
}

func (v *NullableWorkloadClusterConfig) Set(val *WorkloadClusterConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadClusterConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadClusterConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadClusterConfig(val *WorkloadClusterConfig) *NullableWorkloadClusterConfig {
	return &NullableWorkloadClusterConfig{value: val, isSet: true}
}

func (v NullableWorkloadClusterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadClusterConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
