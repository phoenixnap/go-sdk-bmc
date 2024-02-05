# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Transaction ID. | 
**Status** | **string** | The Transaction status. Status may be: SUCCESS, PROCESSING, FAILED | 
**Details** | Pointer to **string** | Details about the transaction. Contains failure reason in case of failed transactions. | [optional] 
**Amount** | **float32** | The transaction amount. | 
**Currency** | **string** | The transaction currency. | 
**Date** | **time.Time** | Date and time when transaction was created. | 
**Metadata** | [**TransactionMetadata**](TransactionMetadata.md) |  | 
**CardPaymentMethodDetails** | [**CardPaymentMethodDetails**](CardPaymentMethodDetails.md) |  | 

## Methods

### NewTransaction

`func NewTransaction(id string, status string, amount float32, currency string, date time.Time, metadata TransactionMetadata, cardPaymentMethodDetails CardPaymentMethodDetails, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *Transaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDetails

`func (o *Transaction) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Transaction) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Transaction) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Transaction) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetAmount

`func (o *Transaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *Transaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Transaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Transaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDate

`func (o *Transaction) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Transaction) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Transaction) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetMetadata

`func (o *Transaction) GetMetadata() TransactionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Transaction) GetMetadataOk() (*TransactionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Transaction) SetMetadata(v TransactionMetadata)`

SetMetadata sets Metadata field to given value.


### GetCardPaymentMethodDetails

`func (o *Transaction) GetCardPaymentMethodDetails() CardPaymentMethodDetails`

GetCardPaymentMethodDetails returns the CardPaymentMethodDetails field if non-nil, zero value otherwise.

### GetCardPaymentMethodDetailsOk

`func (o *Transaction) GetCardPaymentMethodDetailsOk() (*CardPaymentMethodDetails, bool)`

GetCardPaymentMethodDetailsOk returns a tuple with the CardPaymentMethodDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPaymentMethodDetails

`func (o *Transaction) SetCardPaymentMethodDetails(v CardPaymentMethodDetails)`

SetCardPaymentMethodDetails sets CardPaymentMethodDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


