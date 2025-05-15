# ApplicableDiscountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A unique code associated with the discount. | 
**Type** | [**DiscountTypeEnum**](DiscountTypeEnum.md) |  | 
**Value** | **float32** | The value or amount of the discount. The interpretation of this value depends on the &#39;type&#39; of discount.  | 
**CouponCode** | Pointer to **string** | Coupon code which is the source of the discount. | [optional] 

## Methods

### NewApplicableDiscountDetails

`func NewApplicableDiscountDetails(code string, type_ DiscountTypeEnum, value float32, ) *ApplicableDiscountDetails`

NewApplicableDiscountDetails instantiates a new ApplicableDiscountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicableDiscountDetailsWithDefaults

`func NewApplicableDiscountDetailsWithDefaults() *ApplicableDiscountDetails`

NewApplicableDiscountDetailsWithDefaults instantiates a new ApplicableDiscountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ApplicableDiscountDetails) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApplicableDiscountDetails) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApplicableDiscountDetails) SetCode(v string)`

SetCode sets Code field to given value.


### GetType

`func (o *ApplicableDiscountDetails) GetType() DiscountTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicableDiscountDetails) GetTypeOk() (*DiscountTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicableDiscountDetails) SetType(v DiscountTypeEnum)`

SetType sets Type field to given value.


### GetValue

`func (o *ApplicableDiscountDetails) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApplicableDiscountDetails) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApplicableDiscountDetails) SetValue(v float32)`

SetValue sets Value field to given value.


### GetCouponCode

`func (o *ApplicableDiscountDetails) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *ApplicableDiscountDetails) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *ApplicableDiscountDetails) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *ApplicableDiscountDetails) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


