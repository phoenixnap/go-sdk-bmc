# PaginatedTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | Maximum number of items in the page (actual returned length can be less). | 
**Offset** | **int32** | The number of returned items skipped. | 
**Total** | **int64** | The total number of records available for retrieval. | 
**Results** | [**[]Transaction**](Transaction.md) |  | 

## Methods

### NewPaginatedTransactions

`func NewPaginatedTransactions(limit int32, offset int32, total int64, results []Transaction, ) *PaginatedTransactions`

NewPaginatedTransactions instantiates a new PaginatedTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedTransactionsWithDefaults

`func NewPaginatedTransactionsWithDefaults() *PaginatedTransactions`

NewPaginatedTransactionsWithDefaults instantiates a new PaginatedTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PaginatedTransactions) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginatedTransactions) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginatedTransactions) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *PaginatedTransactions) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PaginatedTransactions) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PaginatedTransactions) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *PaginatedTransactions) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginatedTransactions) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginatedTransactions) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetResults

`func (o *PaginatedTransactions) GetResults() []Transaction`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedTransactions) GetResultsOk() (*[]Transaction, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedTransactions) SetResults(v []Transaction)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


