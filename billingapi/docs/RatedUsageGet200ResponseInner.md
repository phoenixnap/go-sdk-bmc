# RatedUsageGet200ResponseInner

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
**PriceModel** | **string** | The price model applied to this usage record. | 
**UnitPrice** | **float32** | The unit price. | 
**UnitPriceDescription** | **string** | User friendly description of the unit price. | 
**Quantity** | **float32** | The number of units being charged. | 
**Active** | **bool** | A flag indicating whether the rated usage record is still active. | 
**UsageSessionId** | **string** | The usage session ID is used to correlate rated usage records across periods of time. For example, a server used for over a month will generate multiple rated usage records. The entire usage session cost can be computed by aggregating the records having the same usage session ID. It is usual to have one rated usage record per month or invoice. | 
**CorrelationId** | **string** | Holds usage record id | 
**ReservationId** | Pointer to **string** | Reservation id associated with this rated usage record. | [optional] 
**Metadata** | [**PublicSubnetDetails**](PublicSubnetDetails.md) |  | 

## Methods

### NewRatedUsageGet200ResponseInner

`func NewRatedUsageGet200ResponseInner(id string, productCategory string, productCode string, location LocationEnum, startDateTime time.Time, endDateTime time.Time, cost int64, priceModel string, unitPrice float32, unitPriceDescription string, quantity float32, active bool, usageSessionId string, correlationId string, metadata PublicSubnetDetails, ) *RatedUsageGet200ResponseInner`

NewRatedUsageGet200ResponseInner instantiates a new RatedUsageGet200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatedUsageGet200ResponseInnerWithDefaults

`func NewRatedUsageGet200ResponseInnerWithDefaults() *RatedUsageGet200ResponseInner`

NewRatedUsageGet200ResponseInnerWithDefaults instantiates a new RatedUsageGet200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RatedUsageGet200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RatedUsageGet200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RatedUsageGet200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetProductCategory

`func (o *RatedUsageGet200ResponseInner) GetProductCategory() string`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *RatedUsageGet200ResponseInner) GetProductCategoryOk() (*string, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *RatedUsageGet200ResponseInner) SetProductCategory(v string)`

SetProductCategory sets ProductCategory field to given value.


### GetProductCode

`func (o *RatedUsageGet200ResponseInner) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *RatedUsageGet200ResponseInner) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *RatedUsageGet200ResponseInner) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetLocation

`func (o *RatedUsageGet200ResponseInner) GetLocation() LocationEnum`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *RatedUsageGet200ResponseInner) GetLocationOk() (*LocationEnum, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *RatedUsageGet200ResponseInner) SetLocation(v LocationEnum)`

SetLocation sets Location field to given value.


### GetYearMonth

`func (o *RatedUsageGet200ResponseInner) GetYearMonth() string`

GetYearMonth returns the YearMonth field if non-nil, zero value otherwise.

### GetYearMonthOk

`func (o *RatedUsageGet200ResponseInner) GetYearMonthOk() (*string, bool)`

GetYearMonthOk returns a tuple with the YearMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearMonth

`func (o *RatedUsageGet200ResponseInner) SetYearMonth(v string)`

SetYearMonth sets YearMonth field to given value.

### HasYearMonth

`func (o *RatedUsageGet200ResponseInner) HasYearMonth() bool`

HasYearMonth returns a boolean if a field has been set.

### GetStartDateTime

`func (o *RatedUsageGet200ResponseInner) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *RatedUsageGet200ResponseInner) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *RatedUsageGet200ResponseInner) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *RatedUsageGet200ResponseInner) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *RatedUsageGet200ResponseInner) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *RatedUsageGet200ResponseInner) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.


### GetCost

`func (o *RatedUsageGet200ResponseInner) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *RatedUsageGet200ResponseInner) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *RatedUsageGet200ResponseInner) SetCost(v int64)`

SetCost sets Cost field to given value.


### GetPriceModel

`func (o *RatedUsageGet200ResponseInner) GetPriceModel() string`

GetPriceModel returns the PriceModel field if non-nil, zero value otherwise.

### GetPriceModelOk

`func (o *RatedUsageGet200ResponseInner) GetPriceModelOk() (*string, bool)`

GetPriceModelOk returns a tuple with the PriceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModel

`func (o *RatedUsageGet200ResponseInner) SetPriceModel(v string)`

SetPriceModel sets PriceModel field to given value.


### GetUnitPrice

`func (o *RatedUsageGet200ResponseInner) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *RatedUsageGet200ResponseInner) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *RatedUsageGet200ResponseInner) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.


### GetUnitPriceDescription

`func (o *RatedUsageGet200ResponseInner) GetUnitPriceDescription() string`

GetUnitPriceDescription returns the UnitPriceDescription field if non-nil, zero value otherwise.

### GetUnitPriceDescriptionOk

`func (o *RatedUsageGet200ResponseInner) GetUnitPriceDescriptionOk() (*string, bool)`

GetUnitPriceDescriptionOk returns a tuple with the UnitPriceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceDescription

`func (o *RatedUsageGet200ResponseInner) SetUnitPriceDescription(v string)`

SetUnitPriceDescription sets UnitPriceDescription field to given value.


### GetQuantity

`func (o *RatedUsageGet200ResponseInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *RatedUsageGet200ResponseInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *RatedUsageGet200ResponseInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetActive

`func (o *RatedUsageGet200ResponseInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RatedUsageGet200ResponseInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RatedUsageGet200ResponseInner) SetActive(v bool)`

SetActive sets Active field to given value.


### GetUsageSessionId

`func (o *RatedUsageGet200ResponseInner) GetUsageSessionId() string`

GetUsageSessionId returns the UsageSessionId field if non-nil, zero value otherwise.

### GetUsageSessionIdOk

`func (o *RatedUsageGet200ResponseInner) GetUsageSessionIdOk() (*string, bool)`

GetUsageSessionIdOk returns a tuple with the UsageSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageSessionId

`func (o *RatedUsageGet200ResponseInner) SetUsageSessionId(v string)`

SetUsageSessionId sets UsageSessionId field to given value.


### GetCorrelationId

`func (o *RatedUsageGet200ResponseInner) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *RatedUsageGet200ResponseInner) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *RatedUsageGet200ResponseInner) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.


### GetReservationId

`func (o *RatedUsageGet200ResponseInner) GetReservationId() string`

GetReservationId returns the ReservationId field if non-nil, zero value otherwise.

### GetReservationIdOk

`func (o *RatedUsageGet200ResponseInner) GetReservationIdOk() (*string, bool)`

GetReservationIdOk returns a tuple with the ReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationId

`func (o *RatedUsageGet200ResponseInner) SetReservationId(v string)`

SetReservationId sets ReservationId field to given value.

### HasReservationId

`func (o *RatedUsageGet200ResponseInner) HasReservationId() bool`

HasReservationId returns a boolean if a field has been set.

### GetMetadata

`func (o *RatedUsageGet200ResponseInner) GetMetadata() PublicSubnetDetails`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RatedUsageGet200ResponseInner) GetMetadataOk() (*PublicSubnetDetails, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RatedUsageGet200ResponseInner) SetMetadata(v PublicSubnetDetails)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


