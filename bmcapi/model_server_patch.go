/*
Bare Metal Cloud API

Create, power off, power on, reset, reboot, or shut down your server with the Bare Metal Cloud API.  Deprovision servers, get or edit SSH key details, assign public IPs, assign servers to networks and a lot more.  Manage your infrastructure more efficiently using just a few simple API calls.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/how-to-deploy-bare-metal-cloud-server' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/bmc/v1/)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmcapi

import (
	"encoding/json"
)

// checks if the ServerPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerPatch{}

// ServerPatch Patch bare metal server.
type ServerPatch struct {
	// Description of server.
	Description *string `json:"description,omitempty"`
	// Hostname of server
	Hostname *string `json:"hostname,omitempty"`
}

// NewServerPatch instantiates a new ServerPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerPatch() *ServerPatch {
	this := ServerPatch{}
	return &this
}

// NewServerPatchWithDefaults instantiates a new ServerPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerPatchWithDefaults() *ServerPatch {
	this := ServerPatch{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServerPatch) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerPatch) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServerPatch) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServerPatch) SetDescription(v string) {
	o.Description = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ServerPatch) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerPatch) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ServerPatch) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ServerPatch) SetHostname(v string) {
	o.Hostname = &v
}

func (o ServerPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	return toSerialize, nil
}

type NullableServerPatch struct {
	value *ServerPatch
	isSet bool
}

func (v NullableServerPatch) Get() *ServerPatch {
	return v.value
}

func (v *NullableServerPatch) Set(val *ServerPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableServerPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableServerPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerPatch(val *ServerPatch) *NullableServerPatch {
	return &NullableServerPatch{value: val, isSet: true}
}

func (v NullableServerPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
