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

// checks if the EsxiOsConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsxiOsConfiguration{}

// EsxiOsConfiguration Esxi OS configuration.
type EsxiOsConfiguration struct {
	DatastoreConfiguration *EsxiDatastoreConfiguration `json:"datastoreConfiguration,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _EsxiOsConfiguration EsxiOsConfiguration

// NewEsxiOsConfiguration instantiates a new EsxiOsConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsxiOsConfiguration() *EsxiOsConfiguration {
	this := EsxiOsConfiguration{}
	return &this
}

// NewEsxiOsConfigurationWithDefaults instantiates a new EsxiOsConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsxiOsConfigurationWithDefaults() *EsxiOsConfiguration {
	this := EsxiOsConfiguration{}
	return &this
}

// GetDatastoreConfiguration returns the DatastoreConfiguration field value if set, zero value otherwise.
func (o *EsxiOsConfiguration) GetDatastoreConfiguration() EsxiDatastoreConfiguration {
	if o == nil || IsNil(o.DatastoreConfiguration) {
		var ret EsxiDatastoreConfiguration
		return ret
	}
	return *o.DatastoreConfiguration
}

// GetDatastoreConfigurationOk returns a tuple with the DatastoreConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsxiOsConfiguration) GetDatastoreConfigurationOk() (*EsxiDatastoreConfiguration, bool) {
	if o == nil || IsNil(o.DatastoreConfiguration) {
		return nil, false
	}
	return o.DatastoreConfiguration, true
}

// HasDatastoreConfiguration returns a boolean if a field has been set.
func (o *EsxiOsConfiguration) HasDatastoreConfiguration() bool {
	if o != nil && !IsNil(o.DatastoreConfiguration) {
		return true
	}

	return false
}

// SetDatastoreConfiguration gets a reference to the given EsxiDatastoreConfiguration and assigns it to the DatastoreConfiguration field.
func (o *EsxiOsConfiguration) SetDatastoreConfiguration(v EsxiDatastoreConfiguration) {
	o.DatastoreConfiguration = &v
}

func (o EsxiOsConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsxiOsConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatastoreConfiguration) {
		toSerialize["datastoreConfiguration"] = o.DatastoreConfiguration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EsxiOsConfiguration) UnmarshalJSON(data []byte) (err error) {
	varEsxiOsConfiguration := _EsxiOsConfiguration{}

	err = json.Unmarshal(data, &varEsxiOsConfiguration)

	if err != nil {
		return err
	}

	*o = EsxiOsConfiguration(varEsxiOsConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "datastoreConfiguration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEsxiOsConfiguration struct {
	value *EsxiOsConfiguration
	isSet bool
}

func (v NullableEsxiOsConfiguration) Get() *EsxiOsConfiguration {
	return v.value
}

func (v *NullableEsxiOsConfiguration) Set(val *EsxiOsConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableEsxiOsConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableEsxiOsConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsxiOsConfiguration(val *EsxiOsConfiguration) *NullableEsxiOsConfiguration {
	return &NullableEsxiOsConfiguration{value: val, isSet: true}
}

func (v NullableEsxiOsConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsxiOsConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
