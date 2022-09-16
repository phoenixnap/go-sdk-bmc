# OperatingSystemRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the rated usage record. | 
**ProductCategory** | **string** | The category of the product associated with this usage record. | 
**ProductCode** | **string** | The code identifying the product associated to this usage record. | 
**Location** | [**LocationEnum**](LocationEnum.md) |  | 
**YearMonth** | Pointer to **string** | Year and month of the usage record. | [optional] 
**StartDateTime** | **time.Time** | The point in time (in UTC) when usage has started. | 
**EndDateTime** | **time.Time** | The point in time (in UTC) until usage has been rated. | 
**Cost** | **int64** | The rated usage in cents. | 
**CostDescription** | Pointer to **string** | The rated usage cost description. | [optional] 
**PriceModel** | **string** | The price model applied to this usage record. | 
**UnitPrice** | **float32** | The unit price. | 
**UnitPriceDescription** | **string** | User friendly description of the unit price. | 
**Quantity** | **float32** | The number of units being charged. | 
**Active** | **bool** | A flag indicating whether the rated usage record is still active. | 
**UsageSessionId** | **string** | The usage session ID is used to correlate rated usage records across periods of time. For example, a server used for over a month will generate multiple rated usage records. The entire usage session cost can be computed by aggregating the records having the same usage session ID. It is usual to have one rated usage record per month or invoice. | 
**CorrelationId** | **string** | Holds usage record id | 
**ReservationId** | Pointer to **string** | Reservation id associated with this rated usage record. | [optional] 
**Metadata** | [**OperatingSystemDetails**](OperatingSystemDetails.md) |  | 

## Methods

### NewOperatingSystemRecord

`func NewOperatingSystemRecord(id string, productCategory string, productCode string, location LocationEnum, startDateTime time.Time, endDateTime time.Time, cost int64, priceModel string, unitPrice float32, unitPriceDescription string, quantity float32, active bool, usageSessionId string, correlationId string, metadata OperatingSystemDetails, ) *OperatingSystemRecord`

NewOperatingSystemRecord instantiates a new OperatingSystemRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatingSystemRecordWithDefaults

`func NewOperatingSystemRecordWithDefaults() *OperatingSystemRecord`

NewOperatingSystemRecordWithDefaults instantiates a new OperatingSystemRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OperatingSystemRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperatingSystemRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperatingSystemRecord) SetId(v string)`

SetId sets Id field to given value.


### GetProductCategory

`func (o *OperatingSystemRecord) GetProductCategory() string`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *OperatingSystemRecord) GetProductCategoryOk() (*string, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *OperatingSystemRecord) SetProductCategory(v string)`

SetProductCategory sets ProductCategory field to given value.


### GetProductCode

`func (o *OperatingSystemRecord) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *OperatingSystemRecord) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *OperatingSystemRecord) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetLocation

`func (o *OperatingSystemRecord) GetLocation() LocationEnum`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OperatingSystemRecord) GetLocationOk() (*LocationEnum, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OperatingSystemRecord) SetLocation(v LocationEnum)`

SetLocation sets Location field to given value.


### GetYearMonth

`func (o *OperatingSystemRecord) GetYearMonth() string`

GetYearMonth returns the YearMonth field if non-nil, zero value otherwise.

### GetYearMonthOk

`func (o *OperatingSystemRecord) GetYearMonthOk() (*string, bool)`

GetYearMonthOk returns a tuple with the YearMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearMonth

`func (o *OperatingSystemRecord) SetYearMonth(v string)`

SetYearMonth sets YearMonth field to given value.

### HasYearMonth

`func (o *OperatingSystemRecord) HasYearMonth() bool`

HasYearMonth returns a boolean if a field has been set.

### GetStartDateTime

`func (o *OperatingSystemRecord) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *OperatingSystemRecord) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *OperatingSystemRecord) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *OperatingSystemRecord) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *OperatingSystemRecord) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *OperatingSystemRecord) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.


### GetCost

`func (o *OperatingSystemRecord) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OperatingSystemRecord) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OperatingSystemRecord) SetCost(v int64)`

SetCost sets Cost field to given value.


### GetCostDescription

`func (o *OperatingSystemRecord) GetCostDescription() string`

GetCostDescription returns the CostDescription field if non-nil, zero value otherwise.

### GetCostDescriptionOk

`func (o *OperatingSystemRecord) GetCostDescriptionOk() (*string, bool)`

GetCostDescriptionOk returns a tuple with the CostDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostDescription

`func (o *OperatingSystemRecord) SetCostDescription(v string)`

SetCostDescription sets CostDescription field to given value.

### HasCostDescription

`func (o *OperatingSystemRecord) HasCostDescription() bool`

HasCostDescription returns a boolean if a field has been set.

### GetPriceModel

`func (o *OperatingSystemRecord) GetPriceModel() string`

GetPriceModel returns the PriceModel field if non-nil, zero value otherwise.

### GetPriceModelOk

`func (o *OperatingSystemRecord) GetPriceModelOk() (*string, bool)`

GetPriceModelOk returns a tuple with the PriceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModel

`func (o *OperatingSystemRecord) SetPriceModel(v string)`

SetPriceModel sets PriceModel field to given value.


### GetUnitPrice

`func (o *OperatingSystemRecord) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *OperatingSystemRecord) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *OperatingSystemRecord) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.


### GetUnitPriceDescription

`func (o *OperatingSystemRecord) GetUnitPriceDescription() string`

GetUnitPriceDescription returns the UnitPriceDescription field if non-nil, zero value otherwise.

### GetUnitPriceDescriptionOk

`func (o *OperatingSystemRecord) GetUnitPriceDescriptionOk() (*string, bool)`

GetUnitPriceDescriptionOk returns a tuple with the UnitPriceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceDescription

`func (o *OperatingSystemRecord) SetUnitPriceDescription(v string)`

SetUnitPriceDescription sets UnitPriceDescription field to given value.


### GetQuantity

`func (o *OperatingSystemRecord) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OperatingSystemRecord) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OperatingSystemRecord) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetActive

`func (o *OperatingSystemRecord) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OperatingSystemRecord) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OperatingSystemRecord) SetActive(v bool)`

SetActive sets Active field to given value.


### GetUsageSessionId

`func (o *OperatingSystemRecord) GetUsageSessionId() string`

GetUsageSessionId returns the UsageSessionId field if non-nil, zero value otherwise.

### GetUsageSessionIdOk

`func (o *OperatingSystemRecord) GetUsageSessionIdOk() (*string, bool)`

GetUsageSessionIdOk returns a tuple with the UsageSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageSessionId

`func (o *OperatingSystemRecord) SetUsageSessionId(v string)`

SetUsageSessionId sets UsageSessionId field to given value.


### GetCorrelationId

`func (o *OperatingSystemRecord) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *OperatingSystemRecord) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *OperatingSystemRecord) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.


### GetReservationId

`func (o *OperatingSystemRecord) GetReservationId() string`

GetReservationId returns the ReservationId field if non-nil, zero value otherwise.

### GetReservationIdOk

`func (o *OperatingSystemRecord) GetReservationIdOk() (*string, bool)`

GetReservationIdOk returns a tuple with the ReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationId

`func (o *OperatingSystemRecord) SetReservationId(v string)`

SetReservationId sets ReservationId field to given value.

### HasReservationId

`func (o *OperatingSystemRecord) HasReservationId() bool`

HasReservationId returns a boolean if a field has been set.

### GetMetadata

`func (o *OperatingSystemRecord) GetMetadata() OperatingSystemDetails`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OperatingSystemRecord) GetMetadataOk() (*OperatingSystemDetails, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OperatingSystemRecord) SetMetadata(v OperatingSystemDetails)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


