/*
Rancher Solution API

Simplify enterprise-grade Kubernetes cluster operations and management with Rancher on Bare Metal Cloud. Deploy Kubernetes clusters using a few API calls.  <span id=\"pnap-api-knowledge-base-link\"> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/rancher-bmc-integration-kubernetes' target='_blank'>here</a> </span>  **All URLs are relative to (https://api.phoenixnap.com/solutions/rancher/v1beta)**

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ranchersolutionapi

import (
	"encoding/json"
)

// RancherServerMetadata (Read Only) Connection parameters to use to connect to the Rancher Server Administrative GUI.
type RancherServerMetadata struct {
	// The Rancher Server URL.
	Url *string `json:"url,omitempty"`
	// The username to use to login to the Rancher Server. This field is returned only as a response to the create cluster request. Make sure to take note or you will not be able to access the server.
	Username *string `json:"username,omitempty"`
	// This is the password to be used to login to the Rancher Server. This field is returned only as a response to the create cluster request. Make sure to take note or you will not be able to access the server.
	Password *string `json:"password,omitempty"`
}

// NewRancherServerMetadata instantiates a new RancherServerMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRancherServerMetadata() *RancherServerMetadata {
	this := RancherServerMetadata{}
	return &this
}

// NewRancherServerMetadataWithDefaults instantiates a new RancherServerMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRancherServerMetadataWithDefaults() *RancherServerMetadata {
	this := RancherServerMetadata{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RancherServerMetadata) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RancherServerMetadata) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RancherServerMetadata) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RancherServerMetadata) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *RancherServerMetadata) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RancherServerMetadata) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *RancherServerMetadata) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *RancherServerMetadata) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *RancherServerMetadata) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RancherServerMetadata) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *RancherServerMetadata) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *RancherServerMetadata) SetPassword(v string) {
	o.Password = &v
}

func (o RancherServerMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableRancherServerMetadata struct {
	value *RancherServerMetadata
	isSet bool
}

func (v NullableRancherServerMetadata) Get() *RancherServerMetadata {
	return v.value
}

func (v *NullableRancherServerMetadata) Set(val *RancherServerMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableRancherServerMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableRancherServerMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRancherServerMetadata(val *RancherServerMetadata) *NullableRancherServerMetadata {
	return &NullableRancherServerMetadata{value: val, isSet: true}
}

func (v NullableRancherServerMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRancherServerMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
