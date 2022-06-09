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

// ReservationModelEnum The reservation model.
type ReservationModelEnum string

// List of ReservationModelEnum
const (
	ONE_MONTH_RESERVATION          ReservationModelEnum = "ONE_MONTH_RESERVATION"
	TWELVE_MONTHS_RESERVATION      ReservationModelEnum = "TWELVE_MONTHS_RESERVATION"
	TWENTY_FOUR_MONTHS_RESERVATION ReservationModelEnum = "TWENTY_FOUR_MONTHS_RESERVATION"
	THIRTY_SIX_MONTHS_RESERVATION  ReservationModelEnum = "THIRTY_SIX_MONTHS_RESERVATION"
	FREE_TIER                      ReservationModelEnum = "FREE_TIER"
)

// All allowed values of ReservationModelEnum enum
var AllowedReservationModelEnumEnumValues = []ReservationModelEnum{
	"ONE_MONTH_RESERVATION",
	"TWELVE_MONTHS_RESERVATION",
	"TWENTY_FOUR_MONTHS_RESERVATION",
	"THIRTY_SIX_MONTHS_RESERVATION",
	"FREE_TIER",
}

func (v *ReservationModelEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReservationModelEnum(value)
	for _, existing := range AllowedReservationModelEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReservationModelEnum", value)
}

// NewReservationModelEnumFromValue returns a pointer to a valid ReservationModelEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReservationModelEnumFromValue(v string) (*ReservationModelEnum, error) {
	ev := ReservationModelEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReservationModelEnum: valid values are %v", v, AllowedReservationModelEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReservationModelEnum) IsValid() bool {
	for _, existing := range AllowedReservationModelEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReservationModelEnum value
func (v ReservationModelEnum) Ptr() *ReservationModelEnum {
	return &v
}

type NullableReservationModelEnum struct {
	value *ReservationModelEnum
	isSet bool
}

func (v NullableReservationModelEnum) Get() *ReservationModelEnum {
	return v.value
}

func (v *NullableReservationModelEnum) Set(val *ReservationModelEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationModelEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationModelEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationModelEnum(val *ReservationModelEnum) *NullableReservationModelEnum {
	return &NullableReservationModelEnum{value: val, isSet: true}
}

func (v NullableReservationModelEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationModelEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
