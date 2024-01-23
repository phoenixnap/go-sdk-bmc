# \TransactionsAPI

All URIs are relative to *https://api.phoenixnap.com/payments/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionIdGet**](TransactionsAPI.md#TransactionIdGet) | **Get** /transactions/{transactionId} | Get Transaction.
[**TransactionsGet**](TransactionsAPI.md#TransactionsGet) | **Get** /transactions | Get Transactions.



## TransactionIdGet

> Transaction TransactionIdGet(ctx, transactionId).Execute()

Get Transaction.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/paymentsapi"
)

func main() {
	transactionId := "0a1b2c3d4f5g6h7i8j9k" // string | The transaction identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.TransactionIdGet(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.TransactionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionIdGet`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.TransactionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The transaction identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsGet

> PaginatedTransactions TransactionsGet(ctx).Limit(limit).Offset(offset).SortDirection(sortDirection).SortField(sortField).From(from).To(to).Execute()

Get Transactions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/paymentsapi"
)

func main() {
	limit := int32(25) // int32 | The limit of the number of results returned. The number of records returned may be smaller than the limit. (optional) (default to 100)
	offset := int32(5) // int32 | The number of items to skip in the results. (optional) (default to 0)
	sortDirection := "sortDirection_example" // string | Sort Given Field depending on the desired direction. Default sorting is descending. (optional) (default to "DESC")
	sortField := "sortField_example" // string | If a sortField is requested, pagination will be done after sorting. Default sorting is by date. (optional) (default to "date")
	from := time.Now() // time.Time | From the date and time (inclusive) to filter transactions by. (optional)
	to := time.Now() // time.Time | To the date and time (inclusive) to filter transactions by. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.TransactionsGet(context.Background()).Limit(limit).Offset(offset).SortDirection(sortDirection).SortField(sortField).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.TransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsGet`: PaginatedTransactions
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.TransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The limit of the number of results returned. The number of records returned may be smaller than the limit. | [default to 100]
 **offset** | **int32** | The number of items to skip in the results. | [default to 0]
 **sortDirection** | **string** | Sort Given Field depending on the desired direction. Default sorting is descending. | [default to &quot;DESC&quot;]
 **sortField** | **string** | If a sortField is requested, pagination will be done after sorting. Default sorting is by date. | [default to &quot;date&quot;]
 **from** | **time.Time** | From the date and time (inclusive) to filter transactions by. | 
 **to** | **time.Time** | To the date and time (inclusive) to filter transactions by. | 

### Return type

[**PaginatedTransactions**](PaginatedTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

