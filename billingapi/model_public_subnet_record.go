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
	"time"
)

// checks if the PublicSubnetRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicSubnetRecord{}

// PublicSubnetRecord struct for PublicSubnetRecord
type PublicSubnetRecord struct {
	// The unique identifier of the rated usage record.
	Id              string                        `json:"id"`
	ProductCategory RatedUsageProductCategoryEnum `json:"productCategory"`
	// The code identifying the product associated to this usage record.
	ProductCode string       `json:"productCode"`
	Location    LocationEnum `json:"location"`
	// Year and month of the usage record.
	YearMonth *string `json:"yearMonth,omitempty"`
	// The point in time (in UTC) when usage has started.
	StartDateTime time.Time `json:"startDateTime"`
	// The point in time (in UTC) until usage has been rated.
	EndDateTime time.Time `json:"endDateTime"`
	// The rated usage in cents.
	Cost int64 `json:"cost"`
	// The cost in cents before discount.
	CostBeforeDiscount *int64 `json:"costBeforeDiscount,omitempty"`
	// The rated usage cost description.
	CostDescription *string `json:"costDescription,omitempty"`
	// The price model applied to this usage record.
	PriceModel string `json:"priceModel"`
	// The unit price.
	UnitPrice float32 `json:"unitPrice"`
	// User friendly description of the unit price.
	UnitPriceDescription string `json:"unitPriceDescription"`
	// The number of units being charged.
	Quantity float32 `json:"quantity"`
	// A flag indicating whether the rated usage record is still active.
	Active bool `json:"active"`
	// The usage session ID is used to correlate rated usage records across periods of time. For example, a server used for over a month will generate multiple rated usage records. The entire usage session cost can be computed by aggregating the records having the same usage session ID. It is usual to have one rated usage record per month or invoice.
	UsageSessionId string `json:"usageSessionId"`
	// Holds usage record id
	CorrelationId string `json:"correlationId"`
	// Reservation id associated with this rated usage record.
	ReservationId        *string                    `json:"reservationId,omitempty"`
	DiscountDetails      *ApplicableDiscountDetails `json:"discountDetails,omitempty"`
	CreditDetails        []CreditDetails            `json:"creditDetails,omitempty"`
	Metadata             PublicSubnetDetails        `json:"metadata"`
	AdditionalProperties map[string]interface{}
}

type _PublicSubnetRecord PublicSubnetRecord

// NewPublicSubnetRecord instantiates a new PublicSubnetRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSubnetRecord(id string, productCategory RatedUsageProductCategoryEnum, productCode string, location LocationEnum, startDateTime time.Time, endDateTime time.Time, cost int64, priceModel string, unitPrice float32, unitPriceDescription string, quantity float32, active bool, usageSessionId string, correlationId string, metadata PublicSubnetDetails) *PublicSubnetRecord {
	this := PublicSubnetRecord{}
	this.Id = id
	this.ProductCategory = productCategory
	this.ProductCode = productCode
	this.Location = location
	this.StartDateTime = startDateTime
	this.EndDateTime = endDateTime
	this.Cost = cost
	this.PriceModel = priceModel
	this.UnitPrice = unitPrice
	this.UnitPriceDescription = unitPriceDescription
	this.Quantity = quantity
	this.Active = active
	this.UsageSessionId = usageSessionId
	this.CorrelationId = correlationId
	this.Metadata = metadata
	return &this
}

// NewPublicSubnetRecordWithDefaults instantiates a new PublicSubnetRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSubnetRecordWithDefaults() *PublicSubnetRecord {
	this := PublicSubnetRecord{}
	return &this
}

// GetId returns the Id field value
func (o *PublicSubnetRecord) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicSubnetRecord) SetId(v string) {
	o.Id = v
}

// GetProductCategory returns the ProductCategory field value
func (o *PublicSubnetRecord) GetProductCategory() RatedUsageProductCategoryEnum {
	if o == nil {
		var ret RatedUsageProductCategoryEnum
		return ret
	}

	return o.ProductCategory
}

// GetProductCategoryOk returns a tuple with the ProductCategory field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetProductCategoryOk() (*RatedUsageProductCategoryEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductCategory, true
}

// SetProductCategory sets field value
func (o *PublicSubnetRecord) SetProductCategory(v RatedUsageProductCategoryEnum) {
	o.ProductCategory = v
}

// GetProductCode returns the ProductCode field value
func (o *PublicSubnetRecord) GetProductCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetProductCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductCode, true
}

// SetProductCode sets field value
func (o *PublicSubnetRecord) SetProductCode(v string) {
	o.ProductCode = v
}

// GetLocation returns the Location field value
func (o *PublicSubnetRecord) GetLocation() LocationEnum {
	if o == nil {
		var ret LocationEnum
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetLocationOk() (*LocationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *PublicSubnetRecord) SetLocation(v LocationEnum) {
	o.Location = v
}

// GetYearMonth returns the YearMonth field value if set, zero value otherwise.
func (o *PublicSubnetRecord) GetYearMonth() string {
	if o == nil || IsNil(o.YearMonth) {
		var ret string
		return ret
	}
	return *o.YearMonth
}

// GetYearMonthOk returns a tuple with the YearMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetYearMonthOk() (*string, bool) {
	if o == nil || IsNil(o.YearMonth) {
		return nil, false
	}
	return o.YearMonth, true
}

// HasYearMonth returns a boolean if a field has been set.
func (o *PublicSubnetRecord) HasYearMonth() bool {
	if o != nil && !IsNil(o.YearMonth) {
		return true
	}

	return false
}

// SetYearMonth gets a reference to the given string and assigns it to the YearMonth field.
func (o *PublicSubnetRecord) SetYearMonth(v string) {
	o.YearMonth = &v
}

// GetStartDateTime returns the StartDateTime field value
func (o *PublicSubnetRecord) GetStartDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDateTime, true
}

// SetStartDateTime sets field value
func (o *PublicSubnetRecord) SetStartDateTime(v time.Time) {
	o.StartDateTime = v
}

// GetEndDateTime returns the EndDateTime field value
func (o *PublicSubnetRecord) GetEndDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDateTime, true
}

// SetEndDateTime sets field value
func (o *PublicSubnetRecord) SetEndDateTime(v time.Time) {
	o.EndDateTime = v
}

// GetCost returns the Cost field value
func (o *PublicSubnetRecord) GetCost() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetCostOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value
func (o *PublicSubnetRecord) SetCost(v int64) {
	o.Cost = v
}

// GetCostBeforeDiscount returns the CostBeforeDiscount field value if set, zero value otherwise.
func (o *PublicSubnetRecord) GetCostBeforeDiscount() int64 {
	if o == nil || IsNil(o.CostBeforeDiscount) {
		var ret int64
		return ret
	}
	return *o.CostBeforeDiscount
}

// GetCostBeforeDiscountOk returns a tuple with the CostBeforeDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetCostBeforeDiscountOk() (*int64, bool) {
	if o == nil || IsNil(o.CostBeforeDiscount) {
		return nil, false
	}
	return o.CostBeforeDiscount, true
}

// HasCostBeforeDiscount returns a boolean if a field has been set.
func (o *PublicSubnetRecord) HasCostBeforeDiscount() bool {
	if o != nil && !IsNil(o.CostBeforeDiscount) {
		return true
	}

	return false
}

// SetCostBeforeDiscount gets a reference to the given int64 and assigns it to the CostBeforeDiscount field.
func (o *PublicSubnetRecord) SetCostBeforeDiscount(v int64) {
	o.CostBeforeDiscount = &v
}

// GetCostDescription returns the CostDescription field value if set, zero value otherwise.
func (o *PublicSubnetRecord) GetCostDescription() string {
	if o == nil || IsNil(o.CostDescription) {
		var ret string
		return ret
	}
	return *o.CostDescription
}

// GetCostDescriptionOk returns a tuple with the CostDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetCostDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.CostDescription) {
		return nil, false
	}
	return o.CostDescription, true
}

// HasCostDescription returns a boolean if a field has been set.
func (o *PublicSubnetRecord) HasCostDescription() bool {
	if o != nil && !IsNil(o.CostDescription) {
		return true
	}

	return false
}

// SetCostDescription gets a reference to the given string and assigns it to the CostDescription field.
func (o *PublicSubnetRecord) SetCostDescription(v string) {
	o.CostDescription = &v
}

// GetPriceModel returns the PriceModel field value
func (o *PublicSubnetRecord) GetPriceModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PriceModel
}

// GetPriceModelOk returns a tuple with the PriceModel field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetPriceModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceModel, true
}

// SetPriceModel sets field value
func (o *PublicSubnetRecord) SetPriceModel(v string) {
	o.PriceModel = v
}

// GetUnitPrice returns the UnitPrice field value
func (o *PublicSubnetRecord) GetUnitPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetUnitPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPrice, true
}

// SetUnitPrice sets field value
func (o *PublicSubnetRecord) SetUnitPrice(v float32) {
	o.UnitPrice = v
}

// GetUnitPriceDescription returns the UnitPriceDescription field value
func (o *PublicSubnetRecord) GetUnitPriceDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitPriceDescription
}

// GetUnitPriceDescriptionOk returns a tuple with the UnitPriceDescription field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetUnitPriceDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPriceDescription, true
}

// SetUnitPriceDescription sets field value
func (o *PublicSubnetRecord) SetUnitPriceDescription(v string) {
	o.UnitPriceDescription = v
}

// GetQuantity returns the Quantity field value
func (o *PublicSubnetRecord) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *PublicSubnetRecord) SetQuantity(v float32) {
	o.Quantity = v
}

// GetActive returns the Active field value
func (o *PublicSubnetRecord) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *PublicSubnetRecord) SetActive(v bool) {
	o.Active = v
}

// GetUsageSessionId returns the UsageSessionId field value
func (o *PublicSubnetRecord) GetUsageSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsageSessionId
}

// GetUsageSessionIdOk returns a tuple with the UsageSessionId field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetUsageSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageSessionId, true
}

// SetUsageSessionId sets field value
func (o *PublicSubnetRecord) SetUsageSessionId(v string) {
	o.UsageSessionId = v
}

// GetCorrelationId returns the CorrelationId field value
func (o *PublicSubnetRecord) GetCorrelationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetCorrelationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CorrelationId, true
}

// SetCorrelationId sets field value
func (o *PublicSubnetRecord) SetCorrelationId(v string) {
	o.CorrelationId = v
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *PublicSubnetRecord) GetReservationId() string {
	if o == nil || IsNil(o.ReservationId) {
		var ret string
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetReservationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *PublicSubnetRecord) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given string and assigns it to the ReservationId field.
func (o *PublicSubnetRecord) SetReservationId(v string) {
	o.ReservationId = &v
}

// GetDiscountDetails returns the DiscountDetails field value if set, zero value otherwise.
func (o *PublicSubnetRecord) GetDiscountDetails() ApplicableDiscountDetails {
	if o == nil || IsNil(o.DiscountDetails) {
		var ret ApplicableDiscountDetails
		return ret
	}
	return *o.DiscountDetails
}

// GetDiscountDetailsOk returns a tuple with the DiscountDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetDiscountDetailsOk() (*ApplicableDiscountDetails, bool) {
	if o == nil || IsNil(o.DiscountDetails) {
		return nil, false
	}
	return o.DiscountDetails, true
}

// HasDiscountDetails returns a boolean if a field has been set.
func (o *PublicSubnetRecord) HasDiscountDetails() bool {
	if o != nil && !IsNil(o.DiscountDetails) {
		return true
	}

	return false
}

// SetDiscountDetails gets a reference to the given ApplicableDiscountDetails and assigns it to the DiscountDetails field.
func (o *PublicSubnetRecord) SetDiscountDetails(v ApplicableDiscountDetails) {
	o.DiscountDetails = &v
}

// GetCreditDetails returns the CreditDetails field value if set, zero value otherwise.
func (o *PublicSubnetRecord) GetCreditDetails() []CreditDetails {
	if o == nil || IsNil(o.CreditDetails) {
		var ret []CreditDetails
		return ret
	}
	return o.CreditDetails
}

// GetCreditDetailsOk returns a tuple with the CreditDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetCreditDetailsOk() ([]CreditDetails, bool) {
	if o == nil || IsNil(o.CreditDetails) {
		return nil, false
	}
	return o.CreditDetails, true
}

// HasCreditDetails returns a boolean if a field has been set.
func (o *PublicSubnetRecord) HasCreditDetails() bool {
	if o != nil && !IsNil(o.CreditDetails) {
		return true
	}

	return false
}

// SetCreditDetails gets a reference to the given []CreditDetails and assigns it to the CreditDetails field.
func (o *PublicSubnetRecord) SetCreditDetails(v []CreditDetails) {
	o.CreditDetails = v
}

// GetMetadata returns the Metadata field value
func (o *PublicSubnetRecord) GetMetadata() PublicSubnetDetails {
	if o == nil {
		var ret PublicSubnetDetails
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *PublicSubnetRecord) GetMetadataOk() (*PublicSubnetDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *PublicSubnetRecord) SetMetadata(v PublicSubnetDetails) {
	o.Metadata = v
}

func (o PublicSubnetRecord) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicSubnetRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["productCategory"] = o.ProductCategory
	toSerialize["productCode"] = o.ProductCode
	toSerialize["location"] = o.Location
	if !IsNil(o.YearMonth) {
		toSerialize["yearMonth"] = o.YearMonth
	}
	toSerialize["startDateTime"] = o.StartDateTime
	toSerialize["endDateTime"] = o.EndDateTime
	toSerialize["cost"] = o.Cost
	if !IsNil(o.CostBeforeDiscount) {
		toSerialize["costBeforeDiscount"] = o.CostBeforeDiscount
	}
	if !IsNil(o.CostDescription) {
		toSerialize["costDescription"] = o.CostDescription
	}
	toSerialize["priceModel"] = o.PriceModel
	toSerialize["unitPrice"] = o.UnitPrice
	toSerialize["unitPriceDescription"] = o.UnitPriceDescription
	toSerialize["quantity"] = o.Quantity
	toSerialize["active"] = o.Active
	toSerialize["usageSessionId"] = o.UsageSessionId
	toSerialize["correlationId"] = o.CorrelationId
	if !IsNil(o.ReservationId) {
		toSerialize["reservationId"] = o.ReservationId
	}
	if !IsNil(o.DiscountDetails) {
		toSerialize["discountDetails"] = o.DiscountDetails
	}
	if !IsNil(o.CreditDetails) {
		toSerialize["creditDetails"] = o.CreditDetails
	}
	toSerialize["metadata"] = o.Metadata

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicSubnetRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"productCategory",
		"productCode",
		"location",
		"startDateTime",
		"endDateTime",
		"cost",
		"priceModel",
		"unitPrice",
		"unitPriceDescription",
		"quantity",
		"active",
		"usageSessionId",
		"correlationId",
		"metadata",
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

	varPublicSubnetRecord := _PublicSubnetRecord{}

	err = json.Unmarshal(data, &varPublicSubnetRecord)

	if err != nil {
		return err
	}

	*o = PublicSubnetRecord(varPublicSubnetRecord)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "productCategory")
		delete(additionalProperties, "productCode")
		delete(additionalProperties, "location")
		delete(additionalProperties, "yearMonth")
		delete(additionalProperties, "startDateTime")
		delete(additionalProperties, "endDateTime")
		delete(additionalProperties, "cost")
		delete(additionalProperties, "costBeforeDiscount")
		delete(additionalProperties, "costDescription")
		delete(additionalProperties, "priceModel")
		delete(additionalProperties, "unitPrice")
		delete(additionalProperties, "unitPriceDescription")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "active")
		delete(additionalProperties, "usageSessionId")
		delete(additionalProperties, "correlationId")
		delete(additionalProperties, "reservationId")
		delete(additionalProperties, "discountDetails")
		delete(additionalProperties, "creditDetails")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicSubnetRecord struct {
	value *PublicSubnetRecord
	isSet bool
}

func (v NullablePublicSubnetRecord) Get() *PublicSubnetRecord {
	return v.value
}

func (v *NullablePublicSubnetRecord) Set(val *PublicSubnetRecord) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSubnetRecord) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSubnetRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSubnetRecord(val *PublicSubnetRecord) *NullablePublicSubnetRecord {
	return &NullablePublicSubnetRecord{value: val, isSet: true}
}

func (v NullablePublicSubnetRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSubnetRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
