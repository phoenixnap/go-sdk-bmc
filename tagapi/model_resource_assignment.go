/*
Tags API

Tags are case-sensitive key-value pairs that simplify resource management. The Tag Manager API allows you to create and manage such tags to later assign them to related resources in your Bare Metal Cloud (through the respective resource apis) in order to group and categorize them.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#server-tag-manager-api' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/tag-manager/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tagapi

import (
	"encoding/json"
	"fmt"
)

// checks if the ResourceAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAssignment{}

// ResourceAssignment Resource assigned to a tag.
type ResourceAssignment struct {
	// The resource name which is automatically generated by the tags-api. It is a unique resource identifier made up of the API name (e.g. bmc, ips), the resource type and the resource ID. This is not to be confused with custom names that are defined for particular resources, such as the server name or network name.
	ResourceName string `json:"resourceName"`
	// The value of the tag assigned to the resource.
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceAssignment ResourceAssignment

// NewResourceAssignment instantiates a new ResourceAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAssignment(resourceName string) *ResourceAssignment {
	this := ResourceAssignment{}
	this.ResourceName = resourceName
	return &this
}

// NewResourceAssignmentWithDefaults instantiates a new ResourceAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAssignmentWithDefaults() *ResourceAssignment {
	this := ResourceAssignment{}
	return &this
}

// GetResourceName returns the ResourceName field value
func (o *ResourceAssignment) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *ResourceAssignment) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *ResourceAssignment) SetResourceName(v string) {
	o.ResourceName = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ResourceAssignment) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAssignment) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ResourceAssignment) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ResourceAssignment) SetValue(v string) {
	o.Value = &v
}

func (o ResourceAssignment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceName"] = o.ResourceName
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceAssignment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceName",
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

	varResourceAssignment := _ResourceAssignment{}

	err = json.Unmarshal(data, &varResourceAssignment)

	if err != nil {
		return err
	}

	*o = ResourceAssignment(varResourceAssignment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceName")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceAssignment struct {
	value *ResourceAssignment
	isSet bool
}

func (v NullableResourceAssignment) Get() *ResourceAssignment {
	return v.value
}

func (v *NullableResourceAssignment) Set(val *ResourceAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAssignment(val *ResourceAssignment) *NullableResourceAssignment {
	return &NullableResourceAssignment{value: val, isSet: true}
}

func (v NullableResourceAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
