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

// SystemCreditCauseEnum The cause/reason of system credit being applied.
type SystemCreditCauseEnum string

// List of SystemCreditCauseEnum
const (
	SYSTEMCREDITCAUSEENUM_RESERVATION_UPGRADE SystemCreditCauseEnum = "RESERVATION_UPGRADE"
)

// All allowed values of SystemCreditCauseEnum enum
var AllowedSystemCreditCauseEnumEnumValues = []SystemCreditCauseEnum{
	"RESERVATION_UPGRADE",
}

func (v *SystemCreditCauseEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SystemCreditCauseEnum(value)
	for _, existing := range AllowedSystemCreditCauseEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SystemCreditCauseEnum", value)
}

// NewSystemCreditCauseEnumFromValue returns a pointer to a valid SystemCreditCauseEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSystemCreditCauseEnumFromValue(v string) (*SystemCreditCauseEnum, error) {
	ev := SystemCreditCauseEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SystemCreditCauseEnum: valid values are %v", v, AllowedSystemCreditCauseEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SystemCreditCauseEnum) IsValid() bool {
	for _, existing := range AllowedSystemCreditCauseEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SystemCreditCauseEnum value
func (v SystemCreditCauseEnum) Ptr() *SystemCreditCauseEnum {
	return &v
}

type NullableSystemCreditCauseEnum struct {
	value *SystemCreditCauseEnum
	isSet bool
}

func (v NullableSystemCreditCauseEnum) Get() *SystemCreditCauseEnum {
	return v.value
}

func (v *NullableSystemCreditCauseEnum) Set(val *SystemCreditCauseEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemCreditCauseEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemCreditCauseEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemCreditCauseEnum(val *SystemCreditCauseEnum) *NullableSystemCreditCauseEnum {
	return &NullableSystemCreditCauseEnum{value: val, isSet: true}
}

func (v NullableSystemCreditCauseEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemCreditCauseEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
