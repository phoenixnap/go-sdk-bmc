/*
Audit Log API

The Audit Logs API lets you read audit log entries and track API calls or activities in the Bare Metal Cloud Portal.<br> <br> <span class=\"pnap-api-knowledge-base-link\"> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#audit-log-api' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/audit/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditapi

import (
	"encoding/json"
)

// ApiActionAllOf struct for ApiActionAllOf
type ApiActionAllOf struct {
	Response Response `json:"response"`
	Request  Request  `json:"request"`
}

// NewApiActionAllOf instantiates a new ApiActionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiActionAllOf(response Response, request Request) *ApiActionAllOf {
	this := ApiActionAllOf{}
	this.Response = response
	this.Request = request
	return &this
}

// NewApiActionAllOfWithDefaults instantiates a new ApiActionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiActionAllOfWithDefaults() *ApiActionAllOf {
	this := ApiActionAllOf{}
	return &this
}

// GetResponse returns the Response field value
func (o *ApiActionAllOf) GetResponse() Response {
	if o == nil {
		var ret Response
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *ApiActionAllOf) GetResponseOk() (*Response, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *ApiActionAllOf) SetResponse(v Response) {
	o.Response = v
}

// GetRequest returns the Request field value
func (o *ApiActionAllOf) GetRequest() Request {
	if o == nil {
		var ret Request
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *ApiActionAllOf) GetRequestOk() (*Request, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *ApiActionAllOf) SetRequest(v Request) {
	o.Request = v
}

func (o ApiActionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["response"] = o.Response
	}
	if true {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableApiActionAllOf struct {
	value *ApiActionAllOf
	isSet bool
}

func (v NullableApiActionAllOf) Get() *ApiActionAllOf {
	return v.value
}

func (v *NullableApiActionAllOf) Set(val *ApiActionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiActionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiActionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiActionAllOf(val *ApiActionAllOf) *NullableApiActionAllOf {
	return &NullableApiActionAllOf{value: val, isSet: true}
}

func (v NullableApiActionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiActionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
