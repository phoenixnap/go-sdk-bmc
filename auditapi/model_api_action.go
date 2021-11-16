/*
Audit Log API

The API that receives and shows actions that an account has taken in the system. </br></br>**All URLs are relative to (https://api.phoenixnap.com/audit/v1/)**

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditapi

import (
	"encoding/json"
	"time"
)

// ApiAction struct for ApiAction
type ApiAction struct {
	// The name of the event.
	Name *string `json:"name,omitempty"`
	// The UTC time the event initiated.
	Timestamp time.Time `json:"timestamp"`
	UserInfo  UserInfo  `json:"userInfo"`
	Response  Response  `json:"response"`
	Request   Request   `json:"request"`
}

// NewApiAction instantiates a new ApiAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAction(timestamp time.Time, userInfo UserInfo, response Response, request Request) *ApiAction {
	this := ApiAction{}
	this.Timestamp = timestamp
	this.UserInfo = userInfo
	this.Response = response
	this.Request = request
	return &this
}

// NewApiActionWithDefaults instantiates a new ApiAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiActionWithDefaults() *ApiAction {
	this := ApiAction{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiAction) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAction) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiAction) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiAction) SetName(v string) {
	o.Name = &v
}

// GetTimestamp returns the Timestamp field value
func (o *ApiAction) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ApiAction) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ApiAction) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetUserInfo returns the UserInfo field value
func (o *ApiAction) GetUserInfo() UserInfo {
	if o == nil {
		var ret UserInfo
		return ret
	}

	return o.UserInfo
}

// GetUserInfoOk returns a tuple with the UserInfo field value
// and a boolean to check if the value has been set.
func (o *ApiAction) GetUserInfoOk() (*UserInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserInfo, true
}

// SetUserInfo sets field value
func (o *ApiAction) SetUserInfo(v UserInfo) {
	o.UserInfo = v
}

// GetResponse returns the Response field value
func (o *ApiAction) GetResponse() Response {
	if o == nil {
		var ret Response
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *ApiAction) GetResponseOk() (*Response, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *ApiAction) SetResponse(v Response) {
	o.Response = v
}

// GetRequest returns the Request field value
func (o *ApiAction) GetRequest() Request {
	if o == nil {
		var ret Request
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *ApiAction) GetRequestOk() (*Request, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *ApiAction) SetRequest(v Request) {
	o.Request = v
}

func (o ApiAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["userInfo"] = o.UserInfo
	}
	if true {
		toSerialize["response"] = o.Response
	}
	if true {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableApiAction struct {
	value *ApiAction
	isSet bool
}

func (v NullableApiAction) Get() *ApiAction {
	return v.value
}

func (v *NullableApiAction) Set(val *ApiAction) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAction) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAction(val *ApiAction) *NullableApiAction {
	return &NullableApiAction{value: val, isSet: true}
}

func (v NullableApiAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
