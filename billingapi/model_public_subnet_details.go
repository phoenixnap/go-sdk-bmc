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
	"fmt"
)

// checks if the PublicSubnetDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicSubnetDetails{}

// PublicSubnetDetails Details of public subnets.
type PublicSubnetDetails struct {
	// Public Subnet identifier as returned by the BMC API.
	Id *string `json:"id,omitempty"`
	// Classless Inter-Domain Routing
	Cidr string `json:"cidr"`
	// CIDR size
	Size                 string `json:"size"`
	AdditionalProperties map[string]interface{}
}

type _PublicSubnetDetails PublicSubnetDetails

// NewPublicSubnetDetails instantiates a new PublicSubnetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSubnetDetails(cidr string, size string) *PublicSubnetDetails {
	this := PublicSubnetDetails{}
	this.Cidr = cidr
	this.Size = size
	return &this
}

// NewPublicSubnetDetailsWithDefaults instantiates a new PublicSubnetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSubnetDetailsWithDefaults() *PublicSubnetDetails {
	this := PublicSubnetDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublicSubnetDetails) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSubnetDetails) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PublicSubnetDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublicSubnetDetails) SetId(v string) {
	o.Id = &v
}

// GetCidr returns the Cidr field value
func (o *PublicSubnetDetails) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetDetails) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *PublicSubnetDetails) SetCidr(v string) {
	o.Cidr = v
}

// GetSize returns the Size field value
func (o *PublicSubnetDetails) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetDetails) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *PublicSubnetDetails) SetSize(v string) {
	o.Size = v
}

func (o PublicSubnetDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicSubnetDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["cidr"] = o.Cidr
	toSerialize["size"] = o.Size

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicSubnetDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cidr",
		"size",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPublicSubnetDetails := _PublicSubnetDetails{}

	err = json.Unmarshal(data, &varPublicSubnetDetails)

	if err != nil {
		return err
	}

	*o = PublicSubnetDetails(varPublicSubnetDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "cidr")
		delete(additionalProperties, "size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicSubnetDetails struct {
	value *PublicSubnetDetails
	isSet bool
}

func (v NullablePublicSubnetDetails) Get() *PublicSubnetDetails {
	return v.value
}

func (v *NullablePublicSubnetDetails) Set(val *PublicSubnetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSubnetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSubnetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSubnetDetails(val *PublicSubnetDetails) *NullablePublicSubnetDetails {
	return &NullablePublicSubnetDetails{value: val, isSet: true}
}

func (v NullablePublicSubnetDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSubnetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
