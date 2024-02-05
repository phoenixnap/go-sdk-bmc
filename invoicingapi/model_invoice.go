/*
Invoicing API

List, fetch and pay invoices with the Invoicing API.

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invoicingapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the Invoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invoice{}

// Invoice struct for Invoice
type Invoice struct {
	// The unique resource identifier of the Invoice.
	Id string `json:"id"`
	// A user-friendly reference number assigned to the invoice.
	Number string `json:"number"`
	// The currency of the invoice. Currently, this field should be set to `EUR` or `USD`.
	Currency string `json:"currency"`
	// The invoice amount.
	Amount float32 `json:"amount"`
	// The invoice outstanding amount.
	OutstandingAmount float32 `json:"outstandingAmount"`
	// The status of the invoice. Currently, this field should be set to `PAID`, `OVERDUE`, `PROCESSING_PAYMENT`, or `UNPAID`.
	Status string `json:"status"`
	// Date and time when the invoice was sent.
	SentOn time.Time `json:"sentOn"`
	// Date and time when the invoice payment is due.
	DueDate time.Time `json:"dueDate"`
}

type _Invoice Invoice

// NewInvoice instantiates a new Invoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoice(id string, number string, currency string, amount float32, outstandingAmount float32, status string, sentOn time.Time, dueDate time.Time) *Invoice {
	this := Invoice{}
	this.Id = id
	this.Number = number
	this.Currency = currency
	this.Amount = amount
	this.OutstandingAmount = outstandingAmount
	this.Status = status
	this.SentOn = sentOn
	this.DueDate = dueDate
	return &this
}

// NewInvoiceWithDefaults instantiates a new Invoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceWithDefaults() *Invoice {
	this := Invoice{}
	return &this
}

// GetId returns the Id field value
func (o *Invoice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Invoice) SetId(v string) {
	o.Id = v
}

// GetNumber returns the Number field value
func (o *Invoice) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *Invoice) SetNumber(v string) {
	o.Number = v
}

// GetCurrency returns the Currency field value
func (o *Invoice) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Invoice) SetCurrency(v string) {
	o.Currency = v
}

// GetAmount returns the Amount field value
func (o *Invoice) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Invoice) SetAmount(v float32) {
	o.Amount = v
}

// GetOutstandingAmount returns the OutstandingAmount field value
func (o *Invoice) GetOutstandingAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OutstandingAmount
}

// GetOutstandingAmountOk returns a tuple with the OutstandingAmount field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetOutstandingAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutstandingAmount, true
}

// SetOutstandingAmount sets field value
func (o *Invoice) SetOutstandingAmount(v float32) {
	o.OutstandingAmount = v
}

// GetStatus returns the Status field value
func (o *Invoice) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Invoice) SetStatus(v string) {
	o.Status = v
}

// GetSentOn returns the SentOn field value
func (o *Invoice) GetSentOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SentOn
}

// GetSentOnOk returns a tuple with the SentOn field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetSentOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SentOn, true
}

// SetSentOn sets field value
func (o *Invoice) SetSentOn(v time.Time) {
	o.SentOn = v
}

// GetDueDate returns the DueDate field value
func (o *Invoice) GetDueDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetDueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DueDate, true
}

// SetDueDate sets field value
func (o *Invoice) SetDueDate(v time.Time) {
	o.DueDate = v
}

func (o Invoice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["number"] = o.Number
	toSerialize["currency"] = o.Currency
	toSerialize["amount"] = o.Amount
	toSerialize["outstandingAmount"] = o.OutstandingAmount
	toSerialize["status"] = o.Status
	toSerialize["sentOn"] = o.SentOn
	toSerialize["dueDate"] = o.DueDate
	return toSerialize, nil
}

func (o *Invoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"number",
		"currency",
		"amount",
		"outstandingAmount",
		"status",
		"sentOn",
		"dueDate",
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

	varInvoice := _Invoice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInvoice)

	if err != nil {
		return err
	}

	*o = Invoice(varInvoice)

	return err
}

type NullableInvoice struct {
	value *Invoice
	isSet bool
}

func (v NullableInvoice) Get() *Invoice {
	return v.value
}

func (v *NullableInvoice) Set(val *Invoice) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoice(val *Invoice) *NullableInvoice {
	return &NullableInvoice{value: val, isSet: true}
}

func (v NullableInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}