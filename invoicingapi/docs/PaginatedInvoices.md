# PaginatedInvoices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | Maximum number of items in the page (actual returned length can be less). | 
**Offset** | **int32** | The number of returned items skipped. | 
**Total** | **int32** | The total number of records available for retrieval. | 
**Results** | [**[]Invoice**](Invoice.md) |  | 

## Methods

### NewPaginatedInvoices

`func NewPaginatedInvoices(limit int32, offset int32, total int32, results []Invoice, ) *PaginatedInvoices`

NewPaginatedInvoices instantiates a new PaginatedInvoices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedInvoicesWithDefaults

`func NewPaginatedInvoicesWithDefaults() *PaginatedInvoices`

NewPaginatedInvoicesWithDefaults instantiates a new PaginatedInvoices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PaginatedInvoices) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginatedInvoices) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginatedInvoices) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *PaginatedInvoices) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PaginatedInvoices) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PaginatedInvoices) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *PaginatedInvoices) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginatedInvoices) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginatedInvoices) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *PaginatedInvoices) GetResults() []Invoice`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedInvoices) GetResultsOk() (*[]Invoice, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedInvoices) SetResults(v []Invoice)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


