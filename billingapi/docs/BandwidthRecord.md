# BandwidthRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the rated usage record. | 
**ProductCategory** | [**RatedUsageProductCategoryEnum**](RatedUsageProductCategoryEnum.md) |  | 
**ProductCode** | **string** | The code identifying the product associated to this usage record. | 
**Location** | [**LocationEnum**](LocationEnum.md) |  | 
**YearMonth** | Pointer to **string** | Year and month of the usage record. | [optional] 
**StartDateTime** | **time.Time** | The point in time (in UTC) when usage has started. | 
**EndDateTime** | **time.Time** | The point in time (in UTC) until usage has been rated. | 
**Cost** | **int64** | The rated usage in cents. | 
**CostBeforeDiscount** | Pointer to **int64** | The cost in cents before discount. | [optional] 
**CostDescription** | Pointer to **string** | The rated usage cost description. | [optional] 
**PriceModel** | **string** | The price model applied to this usage record. | 
**UnitPrice** | **float32** | The unit price. | 
**UnitPriceDescription** | **string** | User friendly description of the unit price. | 
**Quantity** | **float32** | The number of units being charged. | 
**Active** | **bool** | A flag indicating whether the rated usage record is still active. | 
**UsageSessionId** | **string** | The usage session ID is used to correlate rated usage records across periods of time. For example, a server used for over a month will generate multiple rated usage records. The entire usage session cost can be computed by aggregating the records having the same usage session ID. It is usual to have one rated usage record per month or invoice. | 
**CorrelationId** | **string** | Holds usage record id | 
**ReservationId** | Pointer to **string** | Reservation id associated with this rated usage record. | [optional] 
**DiscountDetails** | Pointer to [**DiscountDetails**](DiscountDetails.md) |  | [optional] 
**CreditDetails** | Pointer to [**[]CreditDetails**](CreditDetails.md) |  | [optional] 
**Metadata** | [**BandwidthDetails**](BandwidthDetails.md) |  | 

## Methods

### NewBandwidthRecord

`func NewBandwidthRecord(id string, productCategory RatedUsageProductCategoryEnum, productCode string, location LocationEnum, startDateTime time.Time, endDateTime time.Time, cost int64, priceModel string, unitPrice float32, unitPriceDescription string, quantity float32, active bool, usageSessionId string, correlationId string, metadata BandwidthDetails, ) *BandwidthRecord`

NewBandwidthRecord instantiates a new BandwidthRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthRecordWithDefaults

`func NewBandwidthRecordWithDefaults() *BandwidthRecord`

NewBandwidthRecordWithDefaults instantiates a new BandwidthRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BandwidthRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BandwidthRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BandwidthRecord) SetId(v string)`

SetId sets Id field to given value.


### GetProductCategory

`func (o *BandwidthRecord) GetProductCategory() RatedUsageProductCategoryEnum`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *BandwidthRecord) GetProductCategoryOk() (*RatedUsageProductCategoryEnum, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *BandwidthRecord) SetProductCategory(v RatedUsageProductCategoryEnum)`

SetProductCategory sets ProductCategory field to given value.


### GetProductCode

`func (o *BandwidthRecord) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *BandwidthRecord) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *BandwidthRecord) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetLocation

`func (o *BandwidthRecord) GetLocation() LocationEnum`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BandwidthRecord) GetLocationOk() (*LocationEnum, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BandwidthRecord) SetLocation(v LocationEnum)`

SetLocation sets Location field to given value.


### GetYearMonth

`func (o *BandwidthRecord) GetYearMonth() string`

GetYearMonth returns the YearMonth field if non-nil, zero value otherwise.

### GetYearMonthOk

`func (o *BandwidthRecord) GetYearMonthOk() (*string, bool)`

GetYearMonthOk returns a tuple with the YearMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearMonth

`func (o *BandwidthRecord) SetYearMonth(v string)`

SetYearMonth sets YearMonth field to given value.

### HasYearMonth

`func (o *BandwidthRecord) HasYearMonth() bool`

HasYearMonth returns a boolean if a field has been set.

### GetStartDateTime

`func (o *BandwidthRecord) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *BandwidthRecord) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *BandwidthRecord) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *BandwidthRecord) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *BandwidthRecord) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *BandwidthRecord) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.


### GetCost

`func (o *BandwidthRecord) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BandwidthRecord) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BandwidthRecord) SetCost(v int64)`

SetCost sets Cost field to given value.


### GetCostBeforeDiscount

`func (o *BandwidthRecord) GetCostBeforeDiscount() int64`

GetCostBeforeDiscount returns the CostBeforeDiscount field if non-nil, zero value otherwise.

### GetCostBeforeDiscountOk

`func (o *BandwidthRecord) GetCostBeforeDiscountOk() (*int64, bool)`

GetCostBeforeDiscountOk returns a tuple with the CostBeforeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBeforeDiscount

`func (o *BandwidthRecord) SetCostBeforeDiscount(v int64)`

SetCostBeforeDiscount sets CostBeforeDiscount field to given value.

### HasCostBeforeDiscount

`func (o *BandwidthRecord) HasCostBeforeDiscount() bool`

HasCostBeforeDiscount returns a boolean if a field has been set.

### GetCostDescription

`func (o *BandwidthRecord) GetCostDescription() string`

GetCostDescription returns the CostDescription field if non-nil, zero value otherwise.

### GetCostDescriptionOk

`func (o *BandwidthRecord) GetCostDescriptionOk() (*string, bool)`

GetCostDescriptionOk returns a tuple with the CostDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostDescription

`func (o *BandwidthRecord) SetCostDescription(v string)`

SetCostDescription sets CostDescription field to given value.

### HasCostDescription

`func (o *BandwidthRecord) HasCostDescription() bool`

HasCostDescription returns a boolean if a field has been set.

### GetPriceModel

`func (o *BandwidthRecord) GetPriceModel() string`

GetPriceModel returns the PriceModel field if non-nil, zero value otherwise.

### GetPriceModelOk

`func (o *BandwidthRecord) GetPriceModelOk() (*string, bool)`

GetPriceModelOk returns a tuple with the PriceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModel

`func (o *BandwidthRecord) SetPriceModel(v string)`

SetPriceModel sets PriceModel field to given value.


### GetUnitPrice

`func (o *BandwidthRecord) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *BandwidthRecord) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *BandwidthRecord) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.


### GetUnitPriceDescription

`func (o *BandwidthRecord) GetUnitPriceDescription() string`

GetUnitPriceDescription returns the UnitPriceDescription field if non-nil, zero value otherwise.

### GetUnitPriceDescriptionOk

`func (o *BandwidthRecord) GetUnitPriceDescriptionOk() (*string, bool)`

GetUnitPriceDescriptionOk returns a tuple with the UnitPriceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceDescription

`func (o *BandwidthRecord) SetUnitPriceDescription(v string)`

SetUnitPriceDescription sets UnitPriceDescription field to given value.


### GetQuantity

`func (o *BandwidthRecord) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BandwidthRecord) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BandwidthRecord) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetActive

`func (o *BandwidthRecord) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BandwidthRecord) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BandwidthRecord) SetActive(v bool)`

SetActive sets Active field to given value.


### GetUsageSessionId

`func (o *BandwidthRecord) GetUsageSessionId() string`

GetUsageSessionId returns the UsageSessionId field if non-nil, zero value otherwise.

### GetUsageSessionIdOk

`func (o *BandwidthRecord) GetUsageSessionIdOk() (*string, bool)`

GetUsageSessionIdOk returns a tuple with the UsageSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageSessionId

`func (o *BandwidthRecord) SetUsageSessionId(v string)`

SetUsageSessionId sets UsageSessionId field to given value.


### GetCorrelationId

`func (o *BandwidthRecord) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *BandwidthRecord) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *BandwidthRecord) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.


### GetReservationId

`func (o *BandwidthRecord) GetReservationId() string`

GetReservationId returns the ReservationId field if non-nil, zero value otherwise.

### GetReservationIdOk

`func (o *BandwidthRecord) GetReservationIdOk() (*string, bool)`

GetReservationIdOk returns a tuple with the ReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationId

`func (o *BandwidthRecord) SetReservationId(v string)`

SetReservationId sets ReservationId field to given value.

### HasReservationId

`func (o *BandwidthRecord) HasReservationId() bool`

HasReservationId returns a boolean if a field has been set.

### GetDiscountDetails

`func (o *BandwidthRecord) GetDiscountDetails() DiscountDetails`

GetDiscountDetails returns the DiscountDetails field if non-nil, zero value otherwise.

### GetDiscountDetailsOk

`func (o *BandwidthRecord) GetDiscountDetailsOk() (*DiscountDetails, bool)`

GetDiscountDetailsOk returns a tuple with the DiscountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDetails

`func (o *BandwidthRecord) SetDiscountDetails(v DiscountDetails)`

SetDiscountDetails sets DiscountDetails field to given value.

### HasDiscountDetails

`func (o *BandwidthRecord) HasDiscountDetails() bool`

HasDiscountDetails returns a boolean if a field has been set.

### GetCreditDetails

`func (o *BandwidthRecord) GetCreditDetails() []CreditDetails`

GetCreditDetails returns the CreditDetails field if non-nil, zero value otherwise.

### GetCreditDetailsOk

`func (o *BandwidthRecord) GetCreditDetailsOk() (*[]CreditDetails, bool)`

GetCreditDetailsOk returns a tuple with the CreditDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditDetails

`func (o *BandwidthRecord) SetCreditDetails(v []CreditDetails)`

SetCreditDetails sets CreditDetails field to given value.

### HasCreditDetails

`func (o *BandwidthRecord) HasCreditDetails() bool`

HasCreditDetails returns a boolean if a field has been set.

### GetMetadata

`func (o *BandwidthRecord) GetMetadata() BandwidthDetails`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BandwidthRecord) GetMetadataOk() (*BandwidthDetails, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BandwidthRecord) SetMetadata(v BandwidthDetails)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


