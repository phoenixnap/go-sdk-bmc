# CardPaymentMethodDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardType** | **string** | The Card Type. Supported Card Types include: VISA, MASTERCARD, DISCOVER, JCB &amp; AMEX. | 
**LastFourDigits** | **string** | The last four digits of the card number. | 

## Methods

### NewCardPaymentMethodDetails

`func NewCardPaymentMethodDetails(cardType string, lastFourDigits string, ) *CardPaymentMethodDetails`

NewCardPaymentMethodDetails instantiates a new CardPaymentMethodDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardPaymentMethodDetailsWithDefaults

`func NewCardPaymentMethodDetailsWithDefaults() *CardPaymentMethodDetails`

NewCardPaymentMethodDetailsWithDefaults instantiates a new CardPaymentMethodDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardType

`func (o *CardPaymentMethodDetails) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CardPaymentMethodDetails) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CardPaymentMethodDetails) SetCardType(v string)`

SetCardType sets CardType field to given value.


### GetLastFourDigits

`func (o *CardPaymentMethodDetails) GetLastFourDigits() string`

GetLastFourDigits returns the LastFourDigits field if non-nil, zero value otherwise.

### GetLastFourDigitsOk

`func (o *CardPaymentMethodDetails) GetLastFourDigitsOk() (*string, bool)`

GetLastFourDigitsOk returns a tuple with the LastFourDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFourDigits

`func (o *CardPaymentMethodDetails) SetLastFourDigits(v string)`

SetLastFourDigits sets LastFourDigits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


