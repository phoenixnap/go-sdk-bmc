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
)

// ServerRecordAllOf Server usage record.
type ServerRecordAllOf struct {
	Metadata ServerDetails `json:"metadata"`
}

// NewServerRecordAllOf instantiates a new ServerRecordAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerRecordAllOf(metadata ServerDetails) *ServerRecordAllOf {
	this := ServerRecordAllOf{}
	this.Metadata = metadata
	return &this
}

// NewServerRecordAllOfWithDefaults instantiates a new ServerRecordAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerRecordAllOfWithDefaults() *ServerRecordAllOf {
	this := ServerRecordAllOf{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *ServerRecordAllOf) GetMetadata() ServerDetails {
	if o == nil {
		var ret ServerDetails
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ServerRecordAllOf) GetMetadataOk() (*ServerDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ServerRecordAllOf) SetMetadata(v ServerDetails) {
	o.Metadata = v
}

func (o ServerRecordAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableServerRecordAllOf struct {
	value *ServerRecordAllOf
	isSet bool
}

func (v NullableServerRecordAllOf) Get() *ServerRecordAllOf {
	return v.value
}

func (v *NullableServerRecordAllOf) Set(val *ServerRecordAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServerRecordAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServerRecordAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerRecordAllOf(val *ServerRecordAllOf) *NullableServerRecordAllOf {
	return &NullableServerRecordAllOf{value: val, isSet: true}
}

func (v NullableServerRecordAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerRecordAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


