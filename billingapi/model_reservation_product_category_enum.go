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

// ReservationProductCategoryEnum The commitment product category.
type ReservationProductCategoryEnum string

// List of ReservationProductCategoryEnum
const (
	SERVER    ReservationProductCategoryEnum = "server"
	BANDWIDTH ReservationProductCategoryEnum = "bandwidth"
)

// All allowed values of ReservationProductCategoryEnum enum
var AllowedReservationProductCategoryEnumEnumValues = []ReservationProductCategoryEnum{
	"server",
	"bandwidth",
}

func (v *ReservationProductCategoryEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReservationProductCategoryEnum(value)
	for _, existing := range AllowedReservationProductCategoryEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReservationProductCategoryEnum", value)
}

// NewReservationProductCategoryEnumFromValue returns a pointer to a valid ReservationProductCategoryEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReservationProductCategoryEnumFromValue(v string) (*ReservationProductCategoryEnum, error) {
	ev := ReservationProductCategoryEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReservationProductCategoryEnum: valid values are %v", v, AllowedReservationProductCategoryEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReservationProductCategoryEnum) IsValid() bool {
	for _, existing := range AllowedReservationProductCategoryEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReservationProductCategoryEnum value
func (v ReservationProductCategoryEnum) Ptr() *ReservationProductCategoryEnum {
	return &v
}

type NullableReservationProductCategoryEnum struct {
	value *ReservationProductCategoryEnum
	isSet bool
}

func (v NullableReservationProductCategoryEnum) Get() *ReservationProductCategoryEnum {
	return v.value
}

func (v *NullableReservationProductCategoryEnum) Set(val *ReservationProductCategoryEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationProductCategoryEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationProductCategoryEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationProductCategoryEnum(val *ReservationProductCategoryEnum) *NullableReservationProductCategoryEnum {
	return &NullableReservationProductCategoryEnum{value: val, isSet: true}
}

func (v NullableReservationProductCategoryEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationProductCategoryEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}