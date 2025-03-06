# CreditDetailsBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedAmount** | **float32** | Amount applied. | 
**Type** | [**CreditTypeEnum**](CreditTypeEnum.md) |  | 

## Methods

### NewCreditDetailsBase

`func NewCreditDetailsBase(appliedAmount float32, type_ CreditTypeEnum, ) *CreditDetailsBase`

NewCreditDetailsBase instantiates a new CreditDetailsBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditDetailsBaseWithDefaults

`func NewCreditDetailsBaseWithDefaults() *CreditDetailsBase`

NewCreditDetailsBaseWithDefaults instantiates a new CreditDetailsBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedAmount

`func (o *CreditDetailsBase) GetAppliedAmount() float32`

GetAppliedAmount returns the AppliedAmount field if non-nil, zero value otherwise.

### GetAppliedAmountOk

`func (o *CreditDetailsBase) GetAppliedAmountOk() (*float32, bool)`

GetAppliedAmountOk returns a tuple with the AppliedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAmount

`func (o *CreditDetailsBase) SetAppliedAmount(v float32)`

SetAppliedAmount sets AppliedAmount field to given value.


### GetType

`func (o *CreditDetailsBase) GetType() CreditTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditDetailsBase) GetTypeOk() (*CreditTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditDetailsBase) SetType(v CreditTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


