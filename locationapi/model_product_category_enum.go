/*
Locations API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package locationapi

import (
	"encoding/json"
	"fmt"
)

// ProductCategoryEnum The product category.
type ProductCategoryEnum string

// List of ProductCategoryEnum
const (
	SERVER           ProductCategoryEnum = "SERVER"
	BANDWIDTH        ProductCategoryEnum = "BANDWIDTH"
	OPERATING_SYSTEM ProductCategoryEnum = "OPERATING_SYSTEM"
	PUBLIC_IP        ProductCategoryEnum = "PUBLIC_IP"
	STORAGE          ProductCategoryEnum = "STORAGE"
)

// All allowed values of ProductCategoryEnum enum
var AllowedProductCategoryEnumEnumValues = []ProductCategoryEnum{
	"SERVER",
	"BANDWIDTH",
	"OPERATING_SYSTEM",
	"PUBLIC_IP",
	"STORAGE",
}

func (v *ProductCategoryEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductCategoryEnum(value)
	for _, existing := range AllowedProductCategoryEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductCategoryEnum", value)
}

// NewProductCategoryEnumFromValue returns a pointer to a valid ProductCategoryEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductCategoryEnumFromValue(v string) (*ProductCategoryEnum, error) {
	ev := ProductCategoryEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductCategoryEnum: valid values are %v", v, AllowedProductCategoryEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductCategoryEnum) IsValid() bool {
	for _, existing := range AllowedProductCategoryEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductCategoryEnum value
func (v ProductCategoryEnum) Ptr() *ProductCategoryEnum {
	return &v
}

type NullableProductCategoryEnum struct {
	value *ProductCategoryEnum
	isSet bool
}

func (v NullableProductCategoryEnum) Get() *ProductCategoryEnum {
	return v.value
}

func (v *NullableProductCategoryEnum) Set(val *ProductCategoryEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableProductCategoryEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableProductCategoryEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductCategoryEnum(val *ProductCategoryEnum) *NullableProductCategoryEnum {
	return &NullableProductCategoryEnum{value: val, isSet: true}
}

func (v NullableProductCategoryEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductCategoryEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
