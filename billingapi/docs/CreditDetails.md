# CreditDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedAmount** | **float32** | Amount applied. | 
**Type** | [**CreditTypeEnum**](CreditTypeEnum.md) |  | 
**CouponId** | **string** | Coupon id which is the source of the promo credit. | 
**CouponCode** | **string** | Coupon code which is the source of the promo credit. | 
**Cause** | [**SystemCreditCauseEnum**](SystemCreditCauseEnum.md) |  | 

## Methods

### NewCreditDetails

`func NewCreditDetails(appliedAmount float32, type_ CreditTypeEnum, couponId string, couponCode string, cause SystemCreditCauseEnum, ) *CreditDetails`

NewCreditDetails instantiates a new CreditDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditDetailsWithDefaults

`func NewCreditDetailsWithDefaults() *CreditDetails`

NewCreditDetailsWithDefaults instantiates a new CreditDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedAmount

`func (o *CreditDetails) GetAppliedAmount() float32`

GetAppliedAmount returns the AppliedAmount field if non-nil, zero value otherwise.

### GetAppliedAmountOk

`func (o *CreditDetails) GetAppliedAmountOk() (*float32, bool)`

GetAppliedAmountOk returns a tuple with the AppliedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAmount

`func (o *CreditDetails) SetAppliedAmount(v float32)`

SetAppliedAmount sets AppliedAmount field to given value.


### GetType

`func (o *CreditDetails) GetType() CreditTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditDetails) GetTypeOk() (*CreditTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditDetails) SetType(v CreditTypeEnum)`

SetType sets Type field to given value.


### GetCouponId

`func (o *CreditDetails) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *CreditDetails) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *CreditDetails) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.


### GetCouponCode

`func (o *CreditDetails) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *CreditDetails) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *CreditDetails) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.


### GetCause

`func (o *CreditDetails) GetCause() SystemCreditCauseEnum`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *CreditDetails) GetCauseOk() (*SystemCreditCauseEnum, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *CreditDetails) SetCause(v SystemCreditCauseEnum)`

SetCause sets Cause field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


