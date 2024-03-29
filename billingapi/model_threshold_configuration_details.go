/*
Billing API

Automate your infrastructure billing with the Bare Metal Cloud Billing API. Reserve your server instances to ensure guaranteed resource availability for 12, 24, and 36 months. Retrieve your server’s rated usage for a given period and enable or disable auto-renewals.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/phoenixnap-bare-metal-cloud-billing-models' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/billing/v1/)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package billingapi

import (
	"encoding/json"
	"fmt"
)

// checks if the ThresholdConfigurationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThresholdConfigurationDetails{}

// ThresholdConfigurationDetails Threshold billing configuration.
type ThresholdConfigurationDetails struct {
	// Threshold billing amount.
	ThresholdAmount      float32 `json:"thresholdAmount"`
	AdditionalProperties map[string]interface{}
}

type _ThresholdConfigurationDetails ThresholdConfigurationDetails

// NewThresholdConfigurationDetails instantiates a new ThresholdConfigurationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdConfigurationDetails(thresholdAmount float32) *ThresholdConfigurationDetails {
	this := ThresholdConfigurationDetails{}
	this.ThresholdAmount = thresholdAmount
	return &this
}

// NewThresholdConfigurationDetailsWithDefaults instantiates a new ThresholdConfigurationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdConfigurationDetailsWithDefaults() *ThresholdConfigurationDetails {
	this := ThresholdConfigurationDetails{}
	return &this
}

// GetThresholdAmount returns the ThresholdAmount field value
func (o *ThresholdConfigurationDetails) GetThresholdAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ThresholdAmount
}

// GetThresholdAmountOk returns a tuple with the ThresholdAmount field value
// and a boolean to check if the value has been set.
func (o *ThresholdConfigurationDetails) GetThresholdAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThresholdAmount, true
}

// SetThresholdAmount sets field value
func (o *ThresholdConfigurationDetails) SetThresholdAmount(v float32) {
	o.ThresholdAmount = v
}

func (o ThresholdConfigurationDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThresholdConfigurationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["thresholdAmount"] = o.ThresholdAmount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThresholdConfigurationDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"thresholdAmount",
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

	varThresholdConfigurationDetails := _ThresholdConfigurationDetails{}

	err = json.Unmarshal(data, &varThresholdConfigurationDetails)

	if err != nil {
		return err
	}

	*o = ThresholdConfigurationDetails(varThresholdConfigurationDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "thresholdAmount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThresholdConfigurationDetails struct {
	value *ThresholdConfigurationDetails
	isSet bool
}

func (v NullableThresholdConfigurationDetails) Get() *ThresholdConfigurationDetails {
	return v.value
}

func (v *NullableThresholdConfigurationDetails) Set(val *ThresholdConfigurationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdConfigurationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdConfigurationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdConfigurationDetails(val *ThresholdConfigurationDetails) *NullableThresholdConfigurationDetails {
	return &NullableThresholdConfigurationDetails{value: val, isSet: true}
}

func (v NullableThresholdConfigurationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdConfigurationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
