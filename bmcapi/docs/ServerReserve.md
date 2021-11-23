# ServerReserve

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PricingModel** | **string** | Server pricing model. Currently this field should be set to &#x60;ONE_MONTH_RESERVATION&#x60;, &#x60;TWELVE_MONTHS_RESERVATION&#x60;, &#x60;TWENTY_FOUR_MONTHS_RESERVATION&#x60; or &#x60;THIRTY_SIX_MONTHS_RESERVATION&#x60;. | 

## Methods

### NewServerReserve

`func NewServerReserve(pricingModel string, ) *ServerReserve`

NewServerReserve instantiates a new ServerReserve object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerReserveWithDefaults

`func NewServerReserveWithDefaults() *ServerReserve`

NewServerReserveWithDefaults instantiates a new ServerReserve object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPricingModel

`func (o *ServerReserve) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *ServerReserve) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *ServerReserve) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


