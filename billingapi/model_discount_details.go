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

// checks if the DiscountDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscountDetails{}

// DiscountDetails Represents the details of a discount applied to a product or charge.
type DiscountDetails struct {
	// A unique code associated with the discount.
	Code string           `json:"code"`
	Type DiscountTypeEnum `json:"type"`
	// The value or amount of the discount. The interpretation of this value depends on the 'type' of discount.
	Value                float32 `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _DiscountDetails DiscountDetails

// NewDiscountDetails instantiates a new DiscountDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscountDetails(code string, type_ DiscountTypeEnum, value float32) *DiscountDetails {
	this := DiscountDetails{}
	this.Code = code
	this.Type = type_
	this.Value = value
	return &this
}

// NewDiscountDetailsWithDefaults instantiates a new DiscountDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountDetailsWithDefaults() *DiscountDetails {
	this := DiscountDetails{}
	return &this
}

// GetCode returns the Code field value
func (o *DiscountDetails) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DiscountDetails) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DiscountDetails) SetCode(v string) {
	o.Code = v
}

// GetType returns the Type field value
func (o *DiscountDetails) GetType() DiscountTypeEnum {
	if o == nil {
		var ret DiscountTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DiscountDetails) GetTypeOk() (*DiscountTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DiscountDetails) SetType(v DiscountTypeEnum) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *DiscountDetails) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DiscountDetails) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *DiscountDetails) SetValue(v float32) {
	o.Value = v
}

func (o DiscountDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscountDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscountDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"type",
		"value",
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

	varDiscountDetails := _DiscountDetails{}

	err = json.Unmarshal(data, &varDiscountDetails)

	if err != nil {
		return err
	}

	*o = DiscountDetails(varDiscountDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscountDetails struct {
	value *DiscountDetails
	isSet bool
}

func (v NullableDiscountDetails) Get() *DiscountDetails {
	return v.value
}

func (v *NullableDiscountDetails) Set(val *DiscountDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscountDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscountDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscountDetails(val *DiscountDetails) *NullableDiscountDetails {
	return &NullableDiscountDetails{value: val, isSet: true}
}

func (v NullableDiscountDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscountDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
