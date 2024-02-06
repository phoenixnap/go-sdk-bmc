/*
Payments API

Payments API are currently designed to fetch Transactions only.

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paymentsapi

import (
	"encoding/json"
	"fmt"
)

// checks if the CardPaymentMethodDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardPaymentMethodDetails{}

// CardPaymentMethodDetails Card payment details of a transaction.
type CardPaymentMethodDetails struct {
	// The Card Type. Supported Card Types include: VISA, MASTERCARD, DISCOVER, JCB & AMEX.
	CardType string `json:"cardType"`
	// The last four digits of the card number.
	LastFourDigits       string `json:"lastFourDigits"`
	AdditionalProperties map[string]interface{}
}

type _CardPaymentMethodDetails CardPaymentMethodDetails

// NewCardPaymentMethodDetails instantiates a new CardPaymentMethodDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardPaymentMethodDetails(cardType string, lastFourDigits string) *CardPaymentMethodDetails {
	this := CardPaymentMethodDetails{}
	this.CardType = cardType
	this.LastFourDigits = lastFourDigits
	return &this
}

// NewCardPaymentMethodDetailsWithDefaults instantiates a new CardPaymentMethodDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardPaymentMethodDetailsWithDefaults() *CardPaymentMethodDetails {
	this := CardPaymentMethodDetails{}
	return &this
}

// GetCardType returns the CardType field value
func (o *CardPaymentMethodDetails) GetCardType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value
// and a boolean to check if the value has been set.
func (o *CardPaymentMethodDetails) GetCardTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardType, true
}

// SetCardType sets field value
func (o *CardPaymentMethodDetails) SetCardType(v string) {
	o.CardType = v
}

// GetLastFourDigits returns the LastFourDigits field value
func (o *CardPaymentMethodDetails) GetLastFourDigits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastFourDigits
}

// GetLastFourDigitsOk returns a tuple with the LastFourDigits field value
// and a boolean to check if the value has been set.
func (o *CardPaymentMethodDetails) GetLastFourDigitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastFourDigits, true
}

// SetLastFourDigits sets field value
func (o *CardPaymentMethodDetails) SetLastFourDigits(v string) {
	o.LastFourDigits = v
}

func (o CardPaymentMethodDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardPaymentMethodDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cardType"] = o.CardType
	toSerialize["lastFourDigits"] = o.LastFourDigits

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CardPaymentMethodDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cardType",
		"lastFourDigits",
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

	varCardPaymentMethodDetails := _CardPaymentMethodDetails{}

	err = json.Unmarshal(data, &varCardPaymentMethodDetails)

	if err != nil {
		return err
	}

	*o = CardPaymentMethodDetails(varCardPaymentMethodDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cardType")
		delete(additionalProperties, "lastFourDigits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCardPaymentMethodDetails struct {
	value *CardPaymentMethodDetails
	isSet bool
}

func (v NullableCardPaymentMethodDetails) Get() *CardPaymentMethodDetails {
	return v.value
}

func (v *NullableCardPaymentMethodDetails) Set(val *CardPaymentMethodDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCardPaymentMethodDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCardPaymentMethodDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardPaymentMethodDetails(val *CardPaymentMethodDetails) *NullableCardPaymentMethodDetails {
	return &NullableCardPaymentMethodDetails{value: val, isSet: true}
}

func (v NullableCardPaymentMethodDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardPaymentMethodDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
