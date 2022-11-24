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
	"time"
)

// StorageRecord struct for StorageRecord
type StorageRecord struct {
	// The unique identifier of the rated usage record.
	Id string `json:"id"`
	// The category of the product associated with this usage record.
	ProductCategory string `json:"productCategory"`
	// The code identifying the product associated to this usage record.
	ProductCode string `json:"productCode"`
	Location LocationEnum `json:"location"`
	// Year and month of the usage record.
	YearMonth *string `json:"yearMonth,omitempty"`
	// The point in time (in UTC) when usage has started.
	StartDateTime time.Time `json:"startDateTime"`
	// The point in time (in UTC) until usage has been rated.
	EndDateTime time.Time `json:"endDateTime"`
	// The rated usage in cents.
	Cost int64 `json:"cost"`
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
	ReservationId *string `json:"reservationId,omitempty"`
	Metadata StorageDetails `json:"metadata"`
}

// NewStorageRecord instantiates a new StorageRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageRecord(id string, productCategory string, productCode string, location LocationEnum, startDateTime time.Time, endDateTime time.Time, cost int64, priceModel string, unitPrice float32, unitPriceDescription string, quantity float32, active bool, usageSessionId string, correlationId string, metadata StorageDetails) *StorageRecord {
	this := StorageRecord{}
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

// NewStorageRecordWithDefaults instantiates a new StorageRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageRecordWithDefaults() *StorageRecord {
	this := StorageRecord{}
	return &this
}

// GetId returns the Id field value
func (o *StorageRecord) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StorageRecord) SetId(v string) {
	o.Id = v
}

// GetProductCategory returns the ProductCategory field value
func (o *StorageRecord) GetProductCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductCategory
}

// GetProductCategoryOk returns a tuple with the ProductCategory field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetProductCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductCategory, true
}

// SetProductCategory sets field value
func (o *StorageRecord) SetProductCategory(v string) {
	o.ProductCategory = v
}

// GetProductCode returns the ProductCode field value
func (o *StorageRecord) GetProductCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetProductCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductCode, true
}

// SetProductCode sets field value
func (o *StorageRecord) SetProductCode(v string) {
	o.ProductCode = v
}

// GetLocation returns the Location field value
func (o *StorageRecord) GetLocation() LocationEnum {
	if o == nil {
		var ret LocationEnum
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetLocationOk() (*LocationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *StorageRecord) SetLocation(v LocationEnum) {
	o.Location = v
}

// GetYearMonth returns the YearMonth field value if set, zero value otherwise.
func (o *StorageRecord) GetYearMonth() string {
	if o == nil || o.YearMonth == nil {
		var ret string
		return ret
	}
	return *o.YearMonth
}

// GetYearMonthOk returns a tuple with the YearMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetYearMonthOk() (*string, bool) {
	if o == nil || o.YearMonth == nil {
		return nil, false
	}
	return o.YearMonth, true
}

// HasYearMonth returns a boolean if a field has been set.
func (o *StorageRecord) HasYearMonth() bool {
	if o != nil && o.YearMonth != nil {
		return true
	}

	return false
}

// SetYearMonth gets a reference to the given string and assigns it to the YearMonth field.
func (o *StorageRecord) SetYearMonth(v string) {
	o.YearMonth = &v
}

// GetStartDateTime returns the StartDateTime field value
func (o *StorageRecord) GetStartDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDateTime, true
}

// SetStartDateTime sets field value
func (o *StorageRecord) SetStartDateTime(v time.Time) {
	o.StartDateTime = v
}

// GetEndDateTime returns the EndDateTime field value
func (o *StorageRecord) GetEndDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDateTime, true
}

// SetEndDateTime sets field value
func (o *StorageRecord) SetEndDateTime(v time.Time) {
	o.EndDateTime = v
}

// GetCost returns the Cost field value
func (o *StorageRecord) GetCost() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetCostOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value
func (o *StorageRecord) SetCost(v int64) {
	o.Cost = v
}

// GetCostDescription returns the CostDescription field value if set, zero value otherwise.
func (o *StorageRecord) GetCostDescription() string {
	if o == nil || o.CostDescription == nil {
		var ret string
		return ret
	}
	return *o.CostDescription
}

// GetCostDescriptionOk returns a tuple with the CostDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetCostDescriptionOk() (*string, bool) {
	if o == nil || o.CostDescription == nil {
		return nil, false
	}
	return o.CostDescription, true
}

// HasCostDescription returns a boolean if a field has been set.
func (o *StorageRecord) HasCostDescription() bool {
	if o != nil && o.CostDescription != nil {
		return true
	}

	return false
}

// SetCostDescription gets a reference to the given string and assigns it to the CostDescription field.
func (o *StorageRecord) SetCostDescription(v string) {
	o.CostDescription = &v
}

// GetPriceModel returns the PriceModel field value
func (o *StorageRecord) GetPriceModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PriceModel
}

// GetPriceModelOk returns a tuple with the PriceModel field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetPriceModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceModel, true
}

// SetPriceModel sets field value
func (o *StorageRecord) SetPriceModel(v string) {
	o.PriceModel = v
}

// GetUnitPrice returns the UnitPrice field value
func (o *StorageRecord) GetUnitPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetUnitPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPrice, true
}

// SetUnitPrice sets field value
func (o *StorageRecord) SetUnitPrice(v float32) {
	o.UnitPrice = v
}

// GetUnitPriceDescription returns the UnitPriceDescription field value
func (o *StorageRecord) GetUnitPriceDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitPriceDescription
}

// GetUnitPriceDescriptionOk returns a tuple with the UnitPriceDescription field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetUnitPriceDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPriceDescription, true
}

// SetUnitPriceDescription sets field value
func (o *StorageRecord) SetUnitPriceDescription(v string) {
	o.UnitPriceDescription = v
}

// GetQuantity returns the Quantity field value
func (o *StorageRecord) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *StorageRecord) SetQuantity(v float32) {
	o.Quantity = v
}

// GetActive returns the Active field value
func (o *StorageRecord) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *StorageRecord) SetActive(v bool) {
	o.Active = v
}

// GetUsageSessionId returns the UsageSessionId field value
func (o *StorageRecord) GetUsageSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsageSessionId
}

// GetUsageSessionIdOk returns a tuple with the UsageSessionId field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetUsageSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageSessionId, true
}

// SetUsageSessionId sets field value
func (o *StorageRecord) SetUsageSessionId(v string) {
	o.UsageSessionId = v
}

// GetCorrelationId returns the CorrelationId field value
func (o *StorageRecord) GetCorrelationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetCorrelationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CorrelationId, true
}

// SetCorrelationId sets field value
func (o *StorageRecord) SetCorrelationId(v string) {
	o.CorrelationId = v
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *StorageRecord) GetReservationId() string {
	if o == nil || o.ReservationId == nil {
		var ret string
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetReservationIdOk() (*string, bool) {
	if o == nil || o.ReservationId == nil {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *StorageRecord) HasReservationId() bool {
	if o != nil && o.ReservationId != nil {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given string and assigns it to the ReservationId field.
func (o *StorageRecord) SetReservationId(v string) {
	o.ReservationId = &v
}

// GetMetadata returns the Metadata field value
func (o *StorageRecord) GetMetadata() StorageDetails {
	if o == nil {
		var ret StorageDetails
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *StorageRecord) GetMetadataOk() (*StorageDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *StorageRecord) SetMetadata(v StorageDetails) {
	o.Metadata = v
}

func (o StorageRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["productCategory"] = o.ProductCategory
	}
	if true {
		toSerialize["productCode"] = o.ProductCode
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if o.YearMonth != nil {
		toSerialize["yearMonth"] = o.YearMonth
	}
	if true {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if true {
		toSerialize["endDateTime"] = o.EndDateTime
	}
	if true {
		toSerialize["cost"] = o.Cost
	}
	if o.CostDescription != nil {
		toSerialize["costDescription"] = o.CostDescription
	}
	if true {
		toSerialize["priceModel"] = o.PriceModel
	}
	if true {
		toSerialize["unitPrice"] = o.UnitPrice
	}
	if true {
		toSerialize["unitPriceDescription"] = o.UnitPriceDescription
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["usageSessionId"] = o.UsageSessionId
	}
	if true {
		toSerialize["correlationId"] = o.CorrelationId
	}
	if o.ReservationId != nil {
		toSerialize["reservationId"] = o.ReservationId
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableStorageRecord struct {
	value *StorageRecord
	isSet bool
}

func (v NullableStorageRecord) Get() *StorageRecord {
	return v.value
}

func (v *NullableStorageRecord) Set(val *StorageRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageRecord(val *StorageRecord) *NullableStorageRecord {
	return &NullableStorageRecord{value: val, isSet: true}
}

func (v NullableStorageRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


