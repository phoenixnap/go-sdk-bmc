/*
Audit Log API

The Audit Logs API lets you read audit log entries and track API calls or activities in the Bare Metal Cloud Portal.  <span id=\"pnap-api-knowledge-base-link\"> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#audit-log-api' target='_blank'>here</a> </span>  **All URLs are relative to (https://api.phoenixnap.com/audit/v1/)**

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditapi

import (
	"encoding/json"
	"time"
)

// Event The event log.
type Event struct {
	// The name of the event.
	Name *string `json:"name,omitempty"`
	// The UTC time the event initiated.
	Timestamp time.Time `json:"timestamp"`
	UserInfo  UserInfo  `json:"userInfo"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(timestamp time.Time, userInfo UserInfo) *Event {
	this := Event{}
	this.Timestamp = timestamp
	this.UserInfo = userInfo
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Event) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Event) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Event) SetName(v string) {
	o.Name = &v
}

// GetTimestamp returns the Timestamp field value
func (o *Event) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Event) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Event) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetUserInfo returns the UserInfo field value
func (o *Event) GetUserInfo() UserInfo {
	if o == nil {
		var ret UserInfo
		return ret
	}

	return o.UserInfo
}

// GetUserInfoOk returns a tuple with the UserInfo field value
// and a boolean to check if the value has been set.
func (o *Event) GetUserInfoOk() (*UserInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserInfo, true
}

// SetUserInfo sets field value
func (o *Event) SetUserInfo(v UserInfo) {
	o.UserInfo = v
}

func (o Event) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
