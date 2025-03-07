# ApplicableDiscounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountedPrice** | Pointer to **float32** | The price of the product after applying a discount. | [optional] 
**DiscountDetails** | Pointer to [**[]ApplicableDiscountDetails**](ApplicableDiscountDetails.md) |  | [optional] 

## Methods

### NewApplicableDiscounts

`func NewApplicableDiscounts() *ApplicableDiscounts`

NewApplicableDiscounts instantiates a new ApplicableDiscounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicableDiscountsWithDefaults

`func NewApplicableDiscountsWithDefaults() *ApplicableDiscounts`

NewApplicableDiscountsWithDefaults instantiates a new ApplicableDiscounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountedPrice

`func (o *ApplicableDiscounts) GetDiscountedPrice() float32`

GetDiscountedPrice returns the DiscountedPrice field if non-nil, zero value otherwise.

### GetDiscountedPriceOk

`func (o *ApplicableDiscounts) GetDiscountedPriceOk() (*float32, bool)`

GetDiscountedPriceOk returns a tuple with the DiscountedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPrice

`func (o *ApplicableDiscounts) SetDiscountedPrice(v float32)`

SetDiscountedPrice sets DiscountedPrice field to given value.

### HasDiscountedPrice

`func (o *ApplicableDiscounts) HasDiscountedPrice() bool`

HasDiscountedPrice returns a boolean if a field has been set.

### GetDiscountDetails

`func (o *ApplicableDiscounts) GetDiscountDetails() []ApplicableDiscountDetails`

GetDiscountDetails returns the DiscountDetails field if non-nil, zero value otherwise.

### GetDiscountDetailsOk

`func (o *ApplicableDiscounts) GetDiscountDetailsOk() (*[]ApplicableDiscountDetails, bool)`

GetDiscountDetailsOk returns a tuple with the DiscountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDetails

`func (o *ApplicableDiscounts) SetDiscountDetails(v []ApplicableDiscountDetails)`

SetDiscountDetails sets DiscountDetails field to given value.

### HasDiscountDetails

`func (o *ApplicableDiscounts) HasDiscountDetails() bool`

HasDiscountDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


