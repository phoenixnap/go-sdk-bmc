# Utilization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | [**Quantity**](Quantity.md) |  | 
**Percentage** | **float32** |  | 

## Methods

### NewUtilization

`func NewUtilization(quantity Quantity, percentage float32, ) *Utilization`

NewUtilization instantiates a new Utilization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilizationWithDefaults

`func NewUtilizationWithDefaults() *Utilization`

NewUtilizationWithDefaults instantiates a new Utilization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *Utilization) GetQuantity() Quantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Utilization) GetQuantityOk() (*Quantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Utilization) SetQuantity(v Quantity)`

SetQuantity sets Quantity field to given value.


### GetPercentage

`func (o *Utilization) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *Utilization) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *Utilization) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


