# SystemCreditDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedAmount** | **float32** | Amount applied. | 
**Type** | [**CreditTypeEnum**](CreditTypeEnum.md) |  | 
**Cause** | [**SystemCreditCauseEnum**](SystemCreditCauseEnum.md) |  | 

## Methods

### NewSystemCreditDetails

`func NewSystemCreditDetails(appliedAmount float32, type_ CreditTypeEnum, cause SystemCreditCauseEnum, ) *SystemCreditDetails`

NewSystemCreditDetails instantiates a new SystemCreditDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemCreditDetailsWithDefaults

`func NewSystemCreditDetailsWithDefaults() *SystemCreditDetails`

NewSystemCreditDetailsWithDefaults instantiates a new SystemCreditDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedAmount

`func (o *SystemCreditDetails) GetAppliedAmount() float32`

GetAppliedAmount returns the AppliedAmount field if non-nil, zero value otherwise.

### GetAppliedAmountOk

`func (o *SystemCreditDetails) GetAppliedAmountOk() (*float32, bool)`

GetAppliedAmountOk returns a tuple with the AppliedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAmount

`func (o *SystemCreditDetails) SetAppliedAmount(v float32)`

SetAppliedAmount sets AppliedAmount field to given value.


### GetType

`func (o *SystemCreditDetails) GetType() CreditTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemCreditDetails) GetTypeOk() (*CreditTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemCreditDetails) SetType(v CreditTypeEnum)`

SetType sets Type field to given value.


### GetCause

`func (o *SystemCreditDetails) GetCause() SystemCreditCauseEnum`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *SystemCreditDetails) GetCauseOk() (*SystemCreditCauseEnum, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *SystemCreditDetails) SetCause(v SystemCreditCauseEnum)`

SetCause sets Cause field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


