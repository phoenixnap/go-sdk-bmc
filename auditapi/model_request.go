/*
Audit Log API

The Audit Logs API lets you read audit log entries and track API calls or activities in the Bare Metal Cloud Portal.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#audit-log-api' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/audit/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditapi

import (
	"encoding/json"
)

// Request The request sent by the user / application.
type Request struct {
	Headers Headers `json:"headers"`
	// The request URI.
	Uri string `json:"uri"`
	// The HTTP request verb.
	Verb string `json:"verb"`
}

// NewRequest instantiates a new Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequest(headers Headers, uri string, verb string) *Request {
	this := Request{}
	this.Headers = headers
	this.Uri = uri
	this.Verb = verb
	return &this
}

// NewRequestWithDefaults instantiates a new Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestWithDefaults() *Request {
	this := Request{}
	return &this
}

// GetHeaders returns the Headers field value
func (o *Request) GetHeaders() Headers {
	if o == nil {
		var ret Headers
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *Request) GetHeadersOk() (*Headers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headers, true
}

// SetHeaders sets field value
func (o *Request) SetHeaders(v Headers) {
	o.Headers = v
}

// GetUri returns the Uri field value
func (o *Request) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *Request) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *Request) SetUri(v string) {
	o.Uri = v
}

// GetVerb returns the Verb field value
func (o *Request) GetVerb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verb
}

// GetVerbOk returns a tuple with the Verb field value
// and a boolean to check if the value has been set.
func (o *Request) GetVerbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verb, true
}

// SetVerb sets field value
func (o *Request) SetVerb(v string) {
	o.Verb = v
}

func (o Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["headers"] = o.Headers
	}
	if true {
		toSerialize["uri"] = o.Uri
	}
	if true {
		toSerialize["verb"] = o.Verb
	}
	return json.Marshal(toSerialize)
}

type NullableRequest struct {
	value *Request
	isSet bool
}

func (v NullableRequest) Get() *Request {
	return v.value
}

func (v *NullableRequest) Set(val *Request) {
	v.value = val
	v.isSet = true
}

func (v NullableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequest(val *Request) *NullableRequest {
	return &NullableRequest{value: val, isSet: true}
}

func (v NullableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
