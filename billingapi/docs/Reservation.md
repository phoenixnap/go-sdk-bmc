# Reservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The reservation identifier. | 
**ProductCode** | **string** | The code identifying the product. This code has significant across all locations. | 
**ProductCategory** | [**ReservationProductCategoryEnum**](ReservationProductCategoryEnum.md) |  | 
**Location** | [**LocationEnum**](LocationEnum.md) |  | 
**ReservationModel** | [**ReservationModelEnum**](ReservationModelEnum.md) |  | 
**InitialInvoiceModel** | Pointer to [**ReservationInvoicingModelEnum**](ReservationInvoicingModelEnum.md) |  | [optional] 
**StartDateTime** | **time.Time** | The point in time (in UTC) when the reservation starts. | 
**EndDateTime** | Pointer to **time.Time** | The point in time (in UTC) when the reservation end. | [optional] 
**LastRenewalDateTime** | Pointer to **time.Time** | The point in time (in UTC) when the reservation was renewed last. | [optional] 
**NextRenewalDateTime** | Pointer to **time.Time** | The point in time (in UTC) when the reservation will be renewed if auto renew is set to true. | [optional] 
**AutoRenew** | **bool** | A flag indicating whether the reservation will auto-renew (default is true). | 
**Sku** | **string** | The sku that will be applied to this reservation. It is useful to find out the price by querying the /product endpoint. | 
**Price** | **float32** | Reservation price. | 
**PriceUnit** | [**PriceUnitEnum**](PriceUnitEnum.md) |  | 
**AssignedResourceId** | Pointer to **string** | The resource ID currently being assigned to Reservation. | [optional] 
**NextBillingDate** | Pointer to **string** | Next billing date for Reservation. | [optional] 

## Methods

### NewReservation

`func NewReservation(id string, productCode string, productCategory ReservationProductCategoryEnum, location LocationEnum, reservationModel ReservationModelEnum, startDateTime time.Time, autoRenew bool, sku string, price float32, priceUnit PriceUnitEnum, ) *Reservation`

NewReservation instantiates a new Reservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationWithDefaults

`func NewReservationWithDefaults() *Reservation`

NewReservationWithDefaults instantiates a new Reservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Reservation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Reservation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Reservation) SetId(v string)`

SetId sets Id field to given value.


### GetProductCode

`func (o *Reservation) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *Reservation) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *Reservation) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetProductCategory

`func (o *Reservation) GetProductCategory() ReservationProductCategoryEnum`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *Reservation) GetProductCategoryOk() (*ReservationProductCategoryEnum, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *Reservation) SetProductCategory(v ReservationProductCategoryEnum)`

SetProductCategory sets ProductCategory field to given value.


### GetLocation

`func (o *Reservation) GetLocation() LocationEnum`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Reservation) GetLocationOk() (*LocationEnum, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Reservation) SetLocation(v LocationEnum)`

SetLocation sets Location field to given value.


### GetReservationModel

`func (o *Reservation) GetReservationModel() ReservationModelEnum`

GetReservationModel returns the ReservationModel field if non-nil, zero value otherwise.

### GetReservationModelOk

`func (o *Reservation) GetReservationModelOk() (*ReservationModelEnum, bool)`

GetReservationModelOk returns a tuple with the ReservationModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationModel

`func (o *Reservation) SetReservationModel(v ReservationModelEnum)`

SetReservationModel sets ReservationModel field to given value.


### GetInitialInvoiceModel

`func (o *Reservation) GetInitialInvoiceModel() ReservationInvoicingModelEnum`

GetInitialInvoiceModel returns the InitialInvoiceModel field if non-nil, zero value otherwise.

### GetInitialInvoiceModelOk

`func (o *Reservation) GetInitialInvoiceModelOk() (*ReservationInvoicingModelEnum, bool)`

GetInitialInvoiceModelOk returns a tuple with the InitialInvoiceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialInvoiceModel

`func (o *Reservation) SetInitialInvoiceModel(v ReservationInvoicingModelEnum)`

SetInitialInvoiceModel sets InitialInvoiceModel field to given value.

### HasInitialInvoiceModel

`func (o *Reservation) HasInitialInvoiceModel() bool`

HasInitialInvoiceModel returns a boolean if a field has been set.

### GetStartDateTime

`func (o *Reservation) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *Reservation) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *Reservation) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.


### GetEndDateTime

`func (o *Reservation) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *Reservation) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *Reservation) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *Reservation) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetLastRenewalDateTime

`func (o *Reservation) GetLastRenewalDateTime() time.Time`

GetLastRenewalDateTime returns the LastRenewalDateTime field if non-nil, zero value otherwise.

### GetLastRenewalDateTimeOk

`func (o *Reservation) GetLastRenewalDateTimeOk() (*time.Time, bool)`

GetLastRenewalDateTimeOk returns a tuple with the LastRenewalDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRenewalDateTime

`func (o *Reservation) SetLastRenewalDateTime(v time.Time)`

SetLastRenewalDateTime sets LastRenewalDateTime field to given value.

### HasLastRenewalDateTime

`func (o *Reservation) HasLastRenewalDateTime() bool`

HasLastRenewalDateTime returns a boolean if a field has been set.

### GetNextRenewalDateTime

`func (o *Reservation) GetNextRenewalDateTime() time.Time`

GetNextRenewalDateTime returns the NextRenewalDateTime field if non-nil, zero value otherwise.

### GetNextRenewalDateTimeOk

`func (o *Reservation) GetNextRenewalDateTimeOk() (*time.Time, bool)`

GetNextRenewalDateTimeOk returns a tuple with the NextRenewalDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRenewalDateTime

`func (o *Reservation) SetNextRenewalDateTime(v time.Time)`

SetNextRenewalDateTime sets NextRenewalDateTime field to given value.

### HasNextRenewalDateTime

`func (o *Reservation) HasNextRenewalDateTime() bool`

HasNextRenewalDateTime returns a boolean if a field has been set.

### GetAutoRenew

`func (o *Reservation) GetAutoRenew() bool`

GetAutoRenew returns the AutoRenew field if non-nil, zero value otherwise.

### GetAutoRenewOk

`func (o *Reservation) GetAutoRenewOk() (*bool, bool)`

GetAutoRenewOk returns a tuple with the AutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenew

`func (o *Reservation) SetAutoRenew(v bool)`

SetAutoRenew sets AutoRenew field to given value.


### GetSku

`func (o *Reservation) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Reservation) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Reservation) SetSku(v string)`

SetSku sets Sku field to given value.


### GetPrice

`func (o *Reservation) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Reservation) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Reservation) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetPriceUnit

`func (o *Reservation) GetPriceUnit() PriceUnitEnum`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *Reservation) GetPriceUnitOk() (*PriceUnitEnum, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *Reservation) SetPriceUnit(v PriceUnitEnum)`

SetPriceUnit sets PriceUnit field to given value.


### GetAssignedResourceId

`func (o *Reservation) GetAssignedResourceId() string`

GetAssignedResourceId returns the AssignedResourceId field if non-nil, zero value otherwise.

### GetAssignedResourceIdOk

`func (o *Reservation) GetAssignedResourceIdOk() (*string, bool)`

GetAssignedResourceIdOk returns a tuple with the AssignedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedResourceId

`func (o *Reservation) SetAssignedResourceId(v string)`

SetAssignedResourceId sets AssignedResourceId field to given value.

### HasAssignedResourceId

`func (o *Reservation) HasAssignedResourceId() bool`

HasAssignedResourceId returns a boolean if a field has been set.

### GetNextBillingDate

`func (o *Reservation) GetNextBillingDate() string`

GetNextBillingDate returns the NextBillingDate field if non-nil, zero value otherwise.

### GetNextBillingDateOk

`func (o *Reservation) GetNextBillingDateOk() (*string, bool)`

GetNextBillingDateOk returns a tuple with the NextBillingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBillingDate

`func (o *Reservation) SetNextBillingDate(v string)`

SetNextBillingDate sets NextBillingDate field to given value.

### HasNextBillingDate

`func (o *Reservation) HasNextBillingDate() bool`

HasNextBillingDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


