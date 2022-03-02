/*
Bare Metal Cloud API

Create, power off, power on, reset, reboot, or shut down your server with the Bare Metal Cloud API. Deprovision servers, get or edit SSH key details, and a lot more. Manage your infrastructure more efficiently using just a few simple api calls.  <span id=\"pnap-api-knowledge-base-link\"> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/how-to-deploy-bare-metal-cloud-server' target='_blank'>here</a> </span>  **All URLs are relative to (https://api.phoenixnap.com/bmc/v1/)**

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmcapi

import (
	"encoding/json"
	"time"
)

// QuotaEditLimitRequestDetailsAllOf Pending quota modification requests yet to be reviewed.
type QuotaEditLimitRequestDetailsAllOf struct {
	// The point in time the request was submitted.
	RequestedOn time.Time `json:"requestedOn"`
}

// NewQuotaEditLimitRequestDetailsAllOf instantiates a new QuotaEditLimitRequestDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaEditLimitRequestDetailsAllOf(requestedOn time.Time) *QuotaEditLimitRequestDetailsAllOf {
	this := QuotaEditLimitRequestDetailsAllOf{}
	this.RequestedOn = requestedOn
	return &this
}

// NewQuotaEditLimitRequestDetailsAllOfWithDefaults instantiates a new QuotaEditLimitRequestDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaEditLimitRequestDetailsAllOfWithDefaults() *QuotaEditLimitRequestDetailsAllOf {
	this := QuotaEditLimitRequestDetailsAllOf{}
	return &this
}

// GetRequestedOn returns the RequestedOn field value
func (o *QuotaEditLimitRequestDetailsAllOf) GetRequestedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.RequestedOn
}

// GetRequestedOnOk returns a tuple with the RequestedOn field value
// and a boolean to check if the value has been set.
func (o *QuotaEditLimitRequestDetailsAllOf) GetRequestedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedOn, true
}

// SetRequestedOn sets field value
func (o *QuotaEditLimitRequestDetailsAllOf) SetRequestedOn(v time.Time) {
	o.RequestedOn = v
}

func (o QuotaEditLimitRequestDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requestedOn"] = o.RequestedOn
	}
	return json.Marshal(toSerialize)
}

type NullableQuotaEditLimitRequestDetailsAllOf struct {
	value *QuotaEditLimitRequestDetailsAllOf
	isSet bool
}

func (v NullableQuotaEditLimitRequestDetailsAllOf) Get() *QuotaEditLimitRequestDetailsAllOf {
	return v.value
}

func (v *NullableQuotaEditLimitRequestDetailsAllOf) Set(val *QuotaEditLimitRequestDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaEditLimitRequestDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaEditLimitRequestDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaEditLimitRequestDetailsAllOf(val *QuotaEditLimitRequestDetailsAllOf) *NullableQuotaEditLimitRequestDetailsAllOf {
	return &NullableQuotaEditLimitRequestDetailsAllOf{value: val, isSet: true}
}

func (v NullableQuotaEditLimitRequestDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaEditLimitRequestDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
