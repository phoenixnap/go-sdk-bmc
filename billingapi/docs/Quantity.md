# Quantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | **float32** | Quantity size. | 
**Unit** | [**QuantityUnitEnum**](QuantityUnitEnum.md) |  | 

## Methods

### NewQuantity

`func NewQuantity(quantity float32, unit QuantityUnitEnum, ) *Quantity`

NewQuantity instantiates a new Quantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuantityWithDefaults

`func NewQuantityWithDefaults() *Quantity`

NewQuantityWithDefaults instantiates a new Quantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *Quantity) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Quantity) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Quantity) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnit

`func (o *Quantity) GetUnit() QuantityUnitEnum`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Quantity) GetUnitOk() (*QuantityUnitEnum, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Quantity) SetUnit(v QuantityUnitEnum)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


