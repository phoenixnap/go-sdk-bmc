# PricingPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | **string** | The SKU identifying this pricing plan. | 
**SkuDescription** | Pointer to **string** | Description of this pricing plan. | [optional] 
**Location** | **string** | The code identifying the location. | 
**PricingModel** | **string** | The pricing model. | 
**Price** | **float32** | The price per unit. | 
**PriceUnit** | [**PriceUnitEnum**](PriceUnitEnum.md) |  | 
**ApplicableDiscounts** | Pointer to [**ApplicableDiscounts**](ApplicableDiscounts.md) |  | [optional] 
**CorrelatedProductCode** | Pointer to **string** | Product code of the product this product is correlated with | [optional] 
**PackageQuantity** | Pointer to **float32** | Package size per month. | [optional] 
**PackageUnit** | Pointer to **string** | Package size unit. | [optional] 

## Methods

### NewPricingPlan

`func NewPricingPlan(sku string, location string, pricingModel string, price float32, priceUnit PriceUnitEnum, ) *PricingPlan`

NewPricingPlan instantiates a new PricingPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingPlanWithDefaults

`func NewPricingPlanWithDefaults() *PricingPlan`

NewPricingPlanWithDefaults instantiates a new PricingPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *PricingPlan) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *PricingPlan) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *PricingPlan) SetSku(v string)`

SetSku sets Sku field to given value.


### GetSkuDescription

`func (o *PricingPlan) GetSkuDescription() string`

GetSkuDescription returns the SkuDescription field if non-nil, zero value otherwise.

### GetSkuDescriptionOk

`func (o *PricingPlan) GetSkuDescriptionOk() (*string, bool)`

GetSkuDescriptionOk returns a tuple with the SkuDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuDescription

`func (o *PricingPlan) SetSkuDescription(v string)`

SetSkuDescription sets SkuDescription field to given value.

### HasSkuDescription

`func (o *PricingPlan) HasSkuDescription() bool`

HasSkuDescription returns a boolean if a field has been set.

### GetLocation

`func (o *PricingPlan) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PricingPlan) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PricingPlan) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetPricingModel

`func (o *PricingPlan) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *PricingPlan) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *PricingPlan) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetPrice

`func (o *PricingPlan) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PricingPlan) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PricingPlan) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetPriceUnit

`func (o *PricingPlan) GetPriceUnit() PriceUnitEnum`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *PricingPlan) GetPriceUnitOk() (*PriceUnitEnum, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *PricingPlan) SetPriceUnit(v PriceUnitEnum)`

SetPriceUnit sets PriceUnit field to given value.


### GetApplicableDiscounts

`func (o *PricingPlan) GetApplicableDiscounts() ApplicableDiscounts`

GetApplicableDiscounts returns the ApplicableDiscounts field if non-nil, zero value otherwise.

### GetApplicableDiscountsOk

`func (o *PricingPlan) GetApplicableDiscountsOk() (*ApplicableDiscounts, bool)`

GetApplicableDiscountsOk returns a tuple with the ApplicableDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableDiscounts

`func (o *PricingPlan) SetApplicableDiscounts(v ApplicableDiscounts)`

SetApplicableDiscounts sets ApplicableDiscounts field to given value.

### HasApplicableDiscounts

`func (o *PricingPlan) HasApplicableDiscounts() bool`

HasApplicableDiscounts returns a boolean if a field has been set.

### GetCorrelatedProductCode

`func (o *PricingPlan) GetCorrelatedProductCode() string`

GetCorrelatedProductCode returns the CorrelatedProductCode field if non-nil, zero value otherwise.

### GetCorrelatedProductCodeOk

`func (o *PricingPlan) GetCorrelatedProductCodeOk() (*string, bool)`

GetCorrelatedProductCodeOk returns a tuple with the CorrelatedProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedProductCode

`func (o *PricingPlan) SetCorrelatedProductCode(v string)`

SetCorrelatedProductCode sets CorrelatedProductCode field to given value.

### HasCorrelatedProductCode

`func (o *PricingPlan) HasCorrelatedProductCode() bool`

HasCorrelatedProductCode returns a boolean if a field has been set.

### GetPackageQuantity

`func (o *PricingPlan) GetPackageQuantity() float32`

GetPackageQuantity returns the PackageQuantity field if non-nil, zero value otherwise.

### GetPackageQuantityOk

`func (o *PricingPlan) GetPackageQuantityOk() (*float32, bool)`

GetPackageQuantityOk returns a tuple with the PackageQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQuantity

`func (o *PricingPlan) SetPackageQuantity(v float32)`

SetPackageQuantity sets PackageQuantity field to given value.

### HasPackageQuantity

`func (o *PricingPlan) HasPackageQuantity() bool`

HasPackageQuantity returns a boolean if a field has been set.

### GetPackageUnit

`func (o *PricingPlan) GetPackageUnit() string`

GetPackageUnit returns the PackageUnit field if non-nil, zero value otherwise.

### GetPackageUnitOk

`func (o *PricingPlan) GetPackageUnitOk() (*string, bool)`

GetPackageUnitOk returns a tuple with the PackageUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUnit

`func (o *PricingPlan) SetPackageUnit(v string)`

SetPackageUnit sets PackageUnit field to given value.

### HasPackageUnit

`func (o *PricingPlan) HasPackageUnit() bool`

HasPackageUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


