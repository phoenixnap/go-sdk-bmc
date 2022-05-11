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

// ServerReserve Bare metal server reservation.
type ServerReserve struct {
	// Server pricing model. Currently this field should be set to `ONE_MONTH_RESERVATION`, `TWELVE_MONTHS_RESERVATION`, `TWENTY_FOUR_MONTHS_RESERVATION` or `THIRTY_SIX_MONTHS_RESERVATION`.
	PricingModel string `json:"pricingModel"`
}

// NewServerReserve instantiates a new ServerReserve object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerReserve(pricingModel string) *ServerReserve {
	this := ServerReserve{}
	this.PricingModel = pricingModel
	return &this
}

// NewServerReserveWithDefaults instantiates a new ServerReserve object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerReserveWithDefaults() *ServerReserve {
	this := ServerReserve{}
	return &this
}

// GetPricingModel returns the PricingModel field value
func (o *ServerReserve) GetPricingModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PricingModel
}

// GetPricingModelOk returns a tuple with the PricingModel field value
// and a boolean to check if the value has been set.
func (o *ServerReserve) GetPricingModelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PricingModel, true
}

// SetPricingModel sets field value
func (o *ServerReserve) SetPricingModel(v string) {
	o.PricingModel = v
}

func (o ServerReserve) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pricingModel"] = o.PricingModel
	}
	return json.Marshal(toSerialize)
}

type NullableServerReserve struct {
	value *ServerReserve
	isSet bool
}

func (v NullableServerReserve) Get() *ServerReserve {
	return v.value
}

func (v *NullableServerReserve) Set(val *ServerReserve) {
	v.value = val
	v.isSet = true
}

func (v NullableServerReserve) IsSet() bool {
	return v.isSet
}

func (v *NullableServerReserve) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerReserve(val *ServerReserve) *NullableServerReserve {
	return &NullableServerReserve{value: val, isSet: true}
}

func (v NullableServerReserve) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerReserve) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


