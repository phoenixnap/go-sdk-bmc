# DiscountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A unique code associated with the discount. | 
**Type** | **string** | The type of discount applied. | 
**Value** | **float32** | The value or amount of the discount. The interpretation of this value depends on the &#39;type&#39; of discount.  | 

## Methods

### NewDiscountDetails

`func NewDiscountDetails(code string, type_ string, value float32, ) *DiscountDetails`

NewDiscountDetails instantiates a new DiscountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountDetailsWithDefaults

`func NewDiscountDetailsWithDefaults() *DiscountDetails`

NewDiscountDetailsWithDefaults instantiates a new DiscountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DiscountDetails) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DiscountDetails) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DiscountDetails) SetCode(v string)`

SetCode sets Code field to given value.


### GetType

`func (o *DiscountDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscountDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscountDetails) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *DiscountDetails) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DiscountDetails) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DiscountDetails) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


