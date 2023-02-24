# StorageRecord

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
**Metadata** | [**StorageDetails**](StorageDetails.md) |  | 

## Methods

### NewStorageRecord

`func NewStorageRecord(id string, productCategory string, productCode string, location LocationEnum, startDateTime time.Time, endDateTime time.Time, cost int64, priceModel string, unitPrice float32, unitPriceDescription string, quantity float32, active bool, usageSessionId string, correlationId string, metadata StorageDetails, ) *StorageRecord`

NewStorageRecord instantiates a new StorageRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageRecordWithDefaults

`func NewStorageRecordWithDefaults() *StorageRecord`

NewStorageRecordWithDefaults instantiates a new StorageRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageRecord) SetId(v string)`

SetId sets Id field to given value.


### GetProductCategory

`func (o *StorageRecord) GetProductCategory() string`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *StorageRecord) GetProductCategoryOk() (*string, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *StorageRecord) SetProductCategory(v string)`

SetProductCategory sets ProductCategory field to given value.


### GetProductCode

`func (o *StorageRecord) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *StorageRecord) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *StorageRecord) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetLocation

`func (o *StorageRecord) GetLocation() LocationEnum`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *StorageRecord) GetLocationOk() (*LocationEnum, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *StorageRecord) SetLocation(v LocationEnum)`

SetLocation sets Location field to given value.


### GetYearMonth

`func (o *StorageRecord) GetYearMonth() string`

GetYearMonth returns the YearMonth field if non-nil, zero value otherwise.

### GetYearMonthOk

`func (o *StorageRecord) GetYearMonthOk() (*string, bool)`

GetYearMonthOk returns a tuple with the YearMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearMonth

`func (o *StorageRecord) SetYearMonth(v string)`

SetYearMonth sets YearMonth field to given value.

### HasYearMonth

`func (o *StorageRecord) HasYearMonth() bool`

HasYearMonth returns a boolean if a field has been set.

### GetStartDateTime

`func (o *StorageRecord) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *StorageRecord) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *StorageRecord) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *StorageRecord) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *StorageRecord) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *StorageRecord) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.


### GetCost

`func (o *StorageRecord) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *StorageRecord) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *StorageRecord) SetCost(v int64)`

SetCost sets Cost field to given value.


### GetCostBeforeDiscount

`func (o *StorageRecord) GetCostBeforeDiscount() int64`

GetCostBeforeDiscount returns the CostBeforeDiscount field if non-nil, zero value otherwise.

### GetCostBeforeDiscountOk

`func (o *StorageRecord) GetCostBeforeDiscountOk() (*int64, bool)`

GetCostBeforeDiscountOk returns a tuple with the CostBeforeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBeforeDiscount

`func (o *StorageRecord) SetCostBeforeDiscount(v int64)`

SetCostBeforeDiscount sets CostBeforeDiscount field to given value.

### HasCostBeforeDiscount

`func (o *StorageRecord) HasCostBeforeDiscount() bool`

HasCostBeforeDiscount returns a boolean if a field has been set.

### GetCostDescription

`func (o *StorageRecord) GetCostDescription() string`

GetCostDescription returns the CostDescription field if non-nil, zero value otherwise.

### GetCostDescriptionOk

`func (o *StorageRecord) GetCostDescriptionOk() (*string, bool)`

GetCostDescriptionOk returns a tuple with the CostDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostDescription

`func (o *StorageRecord) SetCostDescription(v string)`

SetCostDescription sets CostDescription field to given value.

### HasCostDescription

`func (o *StorageRecord) HasCostDescription() bool`

HasCostDescription returns a boolean if a field has been set.

### GetPriceModel

`func (o *StorageRecord) GetPriceModel() string`

GetPriceModel returns the PriceModel field if non-nil, zero value otherwise.

### GetPriceModelOk

`func (o *StorageRecord) GetPriceModelOk() (*string, bool)`

GetPriceModelOk returns a tuple with the PriceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModel

`func (o *StorageRecord) SetPriceModel(v string)`

SetPriceModel sets PriceModel field to given value.


### GetUnitPrice

`func (o *StorageRecord) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *StorageRecord) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *StorageRecord) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.


### GetUnitPriceDescription

`func (o *StorageRecord) GetUnitPriceDescription() string`

GetUnitPriceDescription returns the UnitPriceDescription field if non-nil, zero value otherwise.

### GetUnitPriceDescriptionOk

`func (o *StorageRecord) GetUnitPriceDescriptionOk() (*string, bool)`

GetUnitPriceDescriptionOk returns a tuple with the UnitPriceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceDescription

`func (o *StorageRecord) SetUnitPriceDescription(v string)`

SetUnitPriceDescription sets UnitPriceDescription field to given value.


### GetQuantity

`func (o *StorageRecord) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *StorageRecord) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *StorageRecord) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetActive

`func (o *StorageRecord) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StorageRecord) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StorageRecord) SetActive(v bool)`

SetActive sets Active field to given value.


### GetUsageSessionId

`func (o *StorageRecord) GetUsageSessionId() string`

GetUsageSessionId returns the UsageSessionId field if non-nil, zero value otherwise.

### GetUsageSessionIdOk

`func (o *StorageRecord) GetUsageSessionIdOk() (*string, bool)`

GetUsageSessionIdOk returns a tuple with the UsageSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageSessionId

`func (o *StorageRecord) SetUsageSessionId(v string)`

SetUsageSessionId sets UsageSessionId field to given value.


### GetCorrelationId

`func (o *StorageRecord) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *StorageRecord) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *StorageRecord) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.


### GetReservationId

`func (o *StorageRecord) GetReservationId() string`

GetReservationId returns the ReservationId field if non-nil, zero value otherwise.

### GetReservationIdOk

`func (o *StorageRecord) GetReservationIdOk() (*string, bool)`

GetReservationIdOk returns a tuple with the ReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationId

`func (o *StorageRecord) SetReservationId(v string)`

SetReservationId sets ReservationId field to given value.

### HasReservationId

`func (o *StorageRecord) HasReservationId() bool`

HasReservationId returns a boolean if a field has been set.

### GetDiscountDetails

`func (o *StorageRecord) GetDiscountDetails() DiscountDetails`

GetDiscountDetails returns the DiscountDetails field if non-nil, zero value otherwise.

### GetDiscountDetailsOk

`func (o *StorageRecord) GetDiscountDetailsOk() (*DiscountDetails, bool)`

GetDiscountDetailsOk returns a tuple with the DiscountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDetails

`func (o *StorageRecord) SetDiscountDetails(v DiscountDetails)`

SetDiscountDetails sets DiscountDetails field to given value.

### HasDiscountDetails

`func (o *StorageRecord) HasDiscountDetails() bool`

HasDiscountDetails returns a boolean if a field has been set.

### GetMetadata

`func (o *StorageRecord) GetMetadata() StorageDetails`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StorageRecord) GetMetadataOk() (*StorageDetails, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StorageRecord) SetMetadata(v StorageDetails)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


