# TransactionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | The Invoice ID that this transaction pertains to. | 
**InvoiceNumber** | Pointer to **string** | A user-friendly reference number assigned to the invoice that this transaction pertains to. | [optional] 
**IsAutoCharge** | **bool** | Whether this transaction was triggered by an auto charge or not. | 

## Methods

### NewTransactionMetadata

`func NewTransactionMetadata(invoiceId string, isAutoCharge bool, ) *TransactionMetadata`

NewTransactionMetadata instantiates a new TransactionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionMetadataWithDefaults

`func NewTransactionMetadataWithDefaults() *TransactionMetadata`

NewTransactionMetadataWithDefaults instantiates a new TransactionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *TransactionMetadata) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *TransactionMetadata) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *TransactionMetadata) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetInvoiceNumber

`func (o *TransactionMetadata) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *TransactionMetadata) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *TransactionMetadata) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *TransactionMetadata) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetIsAutoCharge

`func (o *TransactionMetadata) GetIsAutoCharge() bool`

GetIsAutoCharge returns the IsAutoCharge field if non-nil, zero value otherwise.

### GetIsAutoChargeOk

`func (o *TransactionMetadata) GetIsAutoChargeOk() (*bool, bool)`

GetIsAutoChargeOk returns a tuple with the IsAutoCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoCharge

`func (o *TransactionMetadata) SetIsAutoCharge(v bool)`

SetIsAutoCharge sets IsAutoCharge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


