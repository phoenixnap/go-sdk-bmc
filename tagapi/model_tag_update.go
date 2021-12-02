/*
Tags API

The Tags Manager API. </br></br>**All URLs are relative to (https://api.phoenixnap.com/tag-manager/v1/)**

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tagapi

import (
	"encoding/json"
)

// TagUpdate Tag update model.
type TagUpdate struct {
	// The unique name of the tag. Tag names are case-sensitive, and should be composed of a maximum of 100 characters including UTF-8 Unicode letters, numbers, and the following symbols: '-', '_'. Regex: [A-zÀ-ú0-9_-]{1,100}
	Name string `json:"name"`
	// The description of the tag.
	Description *string `json:"description,omitempty"`
	// Whether or not to show the tag as part of billing and invoices.
	IsBillingTag bool `json:"isBillingTag"`
}

// NewTagUpdate instantiates a new TagUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagUpdate(name string, isBillingTag bool) *TagUpdate {
	this := TagUpdate{}
	this.Name = name
	this.IsBillingTag = isBillingTag
	return &this
}

// NewTagUpdateWithDefaults instantiates a new TagUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagUpdateWithDefaults() *TagUpdate {
	this := TagUpdate{}
	return &this
}

// GetName returns the Name field value
func (o *TagUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagUpdate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TagUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TagUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TagUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetIsBillingTag returns the IsBillingTag field value
func (o *TagUpdate) GetIsBillingTag() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBillingTag
}

// GetIsBillingTagOk returns a tuple with the IsBillingTag field value
// and a boolean to check if the value has been set.
func (o *TagUpdate) GetIsBillingTagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBillingTag, true
}

// SetIsBillingTag sets field value
func (o *TagUpdate) SetIsBillingTag(v bool) {
	o.IsBillingTag = v
}

func (o TagUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["isBillingTag"] = o.IsBillingTag
	}
	return json.Marshal(toSerialize)
}

type NullableTagUpdate struct {
	value *TagUpdate
	isSet bool
}

func (v NullableTagUpdate) Get() *TagUpdate {
	return v.value
}

func (v *NullableTagUpdate) Set(val *TagUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableTagUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableTagUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagUpdate(val *TagUpdate) *NullableTagUpdate {
	return &NullableTagUpdate{value: val, isSet: true}
}

func (v NullableTagUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}