# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique resource identifier of the Invoice. | 
**Number** | **string** | A user-friendly reference number assigned to the invoice. | 
**Currency** | **string** | The currency of the invoice. Currently, this field should be set to &#x60;EUR&#x60; or &#x60;USD&#x60;. | 
**Amount** | **float32** | The invoice amount. | 
**OutstandingAmount** | **float32** | The invoice outstanding amount. | 
**Status** | **string** | The status of the invoice. Currently, this field should be set to &#x60;PAID&#x60;, &#x60;OVERDUE&#x60;, &#x60;PROCESSING_PAYMENT&#x60;, or &#x60;UNPAID&#x60;. | 
**SentOn** | **time.Time** | Date and time when the invoice was sent. | 
**DueDate** | **time.Time** | Date and time when the invoice payment is due. | 

## Methods

### NewInvoice

`func NewInvoice(id string, number string, currency string, amount float32, outstandingAmount float32, status string, sentOn time.Time, dueDate time.Time, ) *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.


### GetNumber

`func (o *Invoice) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Invoice) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Invoice) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetCurrency

`func (o *Invoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAmount

`func (o *Invoice) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Invoice) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Invoice) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetOutstandingAmount

`func (o *Invoice) GetOutstandingAmount() float32`

GetOutstandingAmount returns the OutstandingAmount field if non-nil, zero value otherwise.

### GetOutstandingAmountOk

`func (o *Invoice) GetOutstandingAmountOk() (*float32, bool)`

GetOutstandingAmountOk returns a tuple with the OutstandingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingAmount

`func (o *Invoice) SetOutstandingAmount(v float32)`

SetOutstandingAmount sets OutstandingAmount field to given value.


### GetStatus

`func (o *Invoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSentOn

`func (o *Invoice) GetSentOn() time.Time`

GetSentOn returns the SentOn field if non-nil, zero value otherwise.

### GetSentOnOk

`func (o *Invoice) GetSentOnOk() (*time.Time, bool)`

GetSentOnOk returns a tuple with the SentOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentOn

`func (o *Invoice) SetSentOn(v time.Time)`

SetSentOn sets SentOn field to given value.


### GetDueDate

`func (o *Invoice) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Invoice) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Invoice) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


