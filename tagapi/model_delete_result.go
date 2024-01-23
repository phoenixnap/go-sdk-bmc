/*
Tags API

Tags are case-sensitive key-value pairs that simplify resource management. The Tag Manager API allows you to create and manage such tags to later assign them to related resources in your Bare Metal Cloud (through the respective resource apis) in order to group and categorize them.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#server-tag-manager-api' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/tag-manager/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tagapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the DeleteResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteResult{}

// DeleteResult Result of a successful delete action.
type DeleteResult struct {
	// Tag deletion result message.
	Result string `json:"result"`
	// The unique identifier of the tag.
	TagId string `json:"tagId"`
}

type _DeleteResult DeleteResult

// NewDeleteResult instantiates a new DeleteResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteResult(result string, tagId string) *DeleteResult {
	this := DeleteResult{}
	this.Result = result
	this.TagId = tagId
	return &this
}

// NewDeleteResultWithDefaults instantiates a new DeleteResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteResultWithDefaults() *DeleteResult {
	this := DeleteResult{}
	return &this
}

// GetResult returns the Result field value
func (o *DeleteResult) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *DeleteResult) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *DeleteResult) SetResult(v string) {
	o.Result = v
}

// GetTagId returns the TagId field value
func (o *DeleteResult) GetTagId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TagId
}

// GetTagIdOk returns a tuple with the TagId field value
// and a boolean to check if the value has been set.
func (o *DeleteResult) GetTagIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagId, true
}

// SetTagId sets field value
func (o *DeleteResult) SetTagId(v string) {
	o.TagId = v
}

func (o DeleteResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["tagId"] = o.TagId
	return toSerialize, nil
}

func (o *DeleteResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
		"tagId",
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

	varDeleteResult := _DeleteResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteResult)

	if err != nil {
		return err
	}

	*o = DeleteResult(varDeleteResult)

	return err
}

type NullableDeleteResult struct {
	value *DeleteResult
	isSet bool
}

func (v NullableDeleteResult) Get() *DeleteResult {
	return v.value
}

func (v *NullableDeleteResult) Set(val *DeleteResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteResult(val *DeleteResult) *NullableDeleteResult {
	return &NullableDeleteResult{value: val, isSet: true}
}

func (v NullableDeleteResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
