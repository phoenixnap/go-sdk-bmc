/*
Bare Metal Cloud API

Create, power off, power on, reset, reboot, or shut down your server with the Bare Metal Cloud API. Deprovision servers, get or edit SSH key details, and a lot more. Manage your infrastructure more efficiently using just a few simple api calls.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/how-to-deploy-bare-metal-cloud-server' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/bmc/v1/)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmcapi

import (
	"encoding/json"
)

// TagAssignmentRequest Tag request to assign to resource.
type TagAssignmentRequest struct {
	// The name of the tag. Tag names are case-sensitive, and should be composed of a maximum of 100 characters including UTF-8 Unicode letters, numbers, and the following symbols: '-', '_'. Regex: [A-zÀ-ú0-9_-]{1,100}.
	Name string `json:"name"`
	// The value of the tag assigned to the resource. Tag values are case-sensitive, and should be composed of a maximum of 100 characters including UTF-8 Unicode letters, numbers, and the following symbols: '-', '_'. Regex: [A-zÀ-ú0-9_-]{1,100}.
	Value *string `json:"value,omitempty"`
}

// NewTagAssignmentRequest instantiates a new TagAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagAssignmentRequest(name string) *TagAssignmentRequest {
	this := TagAssignmentRequest{}
	this.Name = name
	return &this
}

// NewTagAssignmentRequestWithDefaults instantiates a new TagAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagAssignmentRequestWithDefaults() *TagAssignmentRequest {
	this := TagAssignmentRequest{}
	return &this
}

// GetName returns the Name field value
func (o *TagAssignmentRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagAssignmentRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagAssignmentRequest) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TagAssignmentRequest) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagAssignmentRequest) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TagAssignmentRequest) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TagAssignmentRequest) SetValue(v string) {
	o.Value = &v
}

func (o TagAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableTagAssignmentRequest struct {
	value *TagAssignmentRequest
	isSet bool
}

func (v NullableTagAssignmentRequest) Get() *TagAssignmentRequest {
	return v.value
}

func (v *NullableTagAssignmentRequest) Set(val *TagAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTagAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTagAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagAssignmentRequest(val *TagAssignmentRequest) *NullableTagAssignmentRequest {
	return &NullableTagAssignmentRequest{value: val, isSet: true}
}

func (v NullableTagAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
