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
	"fmt"
)

// checks if the ResetResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResetResult{}

// ResetResult Result of a successful reset action.
type ResetResult struct {
	// Message describing the reset result.
	Result string `json:"result"`
	// Password set for user Admin on Windows server or user root on ESXi server.
	Password             *string             `json:"password,omitempty"`
	OsConfiguration      *OsConfigurationMap `json:"osConfiguration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResetResult ResetResult

// NewResetResult instantiates a new ResetResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetResult(result string) *ResetResult {
	this := ResetResult{}
	this.Result = result
	return &this
}

// NewResetResultWithDefaults instantiates a new ResetResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetResultWithDefaults() *ResetResult {
	this := ResetResult{}
	return &this
}

// GetResult returns the Result field value
func (o *ResetResult) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ResetResult) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *ResetResult) SetResult(v string) {
	o.Result = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ResetResult) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetResult) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ResetResult) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ResetResult) SetPassword(v string) {
	o.Password = &v
}

// GetOsConfiguration returns the OsConfiguration field value if set, zero value otherwise.
func (o *ResetResult) GetOsConfiguration() OsConfigurationMap {
	if o == nil || IsNil(o.OsConfiguration) {
		var ret OsConfigurationMap
		return ret
	}
	return *o.OsConfiguration
}

// GetOsConfigurationOk returns a tuple with the OsConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetResult) GetOsConfigurationOk() (*OsConfigurationMap, bool) {
	if o == nil || IsNil(o.OsConfiguration) {
		return nil, false
	}
	return o.OsConfiguration, true
}

// HasOsConfiguration returns a boolean if a field has been set.
func (o *ResetResult) HasOsConfiguration() bool {
	if o != nil && !IsNil(o.OsConfiguration) {
		return true
	}

	return false
}

// SetOsConfiguration gets a reference to the given OsConfigurationMap and assigns it to the OsConfiguration field.
func (o *ResetResult) SetOsConfiguration(v OsConfigurationMap) {
	o.OsConfiguration = &v
}

func (o ResetResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.OsConfiguration) {
		toSerialize["osConfiguration"] = o.OsConfiguration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResetResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
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

	varResetResult := _ResetResult{}

	err = json.Unmarshal(data, &varResetResult)

	if err != nil {
		return err
	}

	*o = ResetResult(varResetResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		delete(additionalProperties, "password")
		delete(additionalProperties, "osConfiguration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResetResult struct {
	value *ResetResult
	isSet bool
}

func (v NullableResetResult) Get() *ResetResult {
	return v.value
}

func (v *NullableResetResult) Set(val *ResetResult) {
	v.value = val
	v.isSet = true
}

func (v NullableResetResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResetResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetResult(val *ResetResult) *NullableResetResult {
	return &NullableResetResult{value: val, isSet: true}
}

func (v NullableResetResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
