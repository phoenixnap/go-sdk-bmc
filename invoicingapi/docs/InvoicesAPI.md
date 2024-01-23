# \InvoicesAPI

All URIs are relative to *https://api.phoenixnap.com/invoicing/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoicesGet**](InvoicesAPI.md#InvoicesGet) | **Get** /invoices | List invoices.
[**InvoicesInvoiceIdGeneratePdfPost**](InvoicesAPI.md#InvoicesInvoiceIdGeneratePdfPost) | **Post** /invoices/{invoiceId}/actions/generate-pdf | Generate invoice details as PDF.
[**InvoicesInvoiceIdGet**](InvoicesAPI.md#InvoicesInvoiceIdGet) | **Get** /invoices/{invoiceId} | Get invoice details.
[**InvoicesInvoiceIdPayPost**](InvoicesAPI.md#InvoicesInvoiceIdPayPost) | **Post** /invoices/{invoiceId}/actions/pay | Pay an invoice.



## InvoicesGet

> PaginatedInvoices InvoicesGet(ctx).Number(number).Status(status).SentOnFrom(sentOnFrom).SentOnTo(sentOnTo).Limit(limit).Offset(offset).SortField(sortField).SortDirection(sortDirection).Execute()

List invoices.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/invoicingapi"
)

func main() {
	number := "13218-1180326" // string | A user-friendly reference number assigned to the invoice. (optional)
	status := "status_example" // string | Payment status of the invoice. (optional)
	sentOnFrom := time.Now() // time.Time | Minimum value to filter invoices by sent on date. (optional)
	sentOnTo := time.Now() // time.Time | Maximum value to filter invoices by sent on date. (optional)
	limit := int32(25) // int32 | The limit of the number of results returned. The number of records returned may be smaller than the limit. (optional) (default to 100)
	offset := int32(5) // int32 | The number of items to skip in the results. (optional) (default to 0)
	sortField := "sortField_example" // string | If a sortField is requested, pagination will be done after sorting. Default sorting is by number. (optional) (default to "sentOn")
	sortDirection := "sortDirection_example" // string | Sort Given Field depending on the desired direction. Default sorting is descending. (optional) (default to "DESC")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesGet(context.Background()).Number(number).Status(status).SentOnFrom(sentOnFrom).SentOnTo(sentOnTo).Limit(limit).Offset(offset).SortField(sortField).SortDirection(sortDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesGet`: PaginatedInvoices
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **number** | **string** | A user-friendly reference number assigned to the invoice. | 
 **status** | **string** | Payment status of the invoice. | 
 **sentOnFrom** | **time.Time** | Minimum value to filter invoices by sent on date. | 
 **sentOnTo** | **time.Time** | Maximum value to filter invoices by sent on date. | 
 **limit** | **int32** | The limit of the number of results returned. The number of records returned may be smaller than the limit. | [default to 100]
 **offset** | **int32** | The number of items to skip in the results. | [default to 0]
 **sortField** | **string** | If a sortField is requested, pagination will be done after sorting. Default sorting is by number. | [default to &quot;sentOn&quot;]
 **sortDirection** | **string** | Sort Given Field depending on the desired direction. Default sorting is descending. | [default to &quot;DESC&quot;]

### Return type

[**PaginatedInvoices**](PaginatedInvoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesInvoiceIdGeneratePdfPost

> *os.File InvoicesInvoiceIdGeneratePdfPost(ctx, invoiceId).Execute()

Generate invoice details as PDF.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/invoicingapi"
)

func main() {
	invoiceId := "5fa54d1e91867c03a0a7b4a4" // string | The unique resource identifier of the Invoice.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesInvoiceIdGeneratePdfPost(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesInvoiceIdGeneratePdfPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesInvoiceIdGeneratePdfPost`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesInvoiceIdGeneratePdfPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The unique resource identifier of the Invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesInvoiceIdGeneratePdfPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesInvoiceIdGet

> Invoice InvoicesInvoiceIdGet(ctx, invoiceId).Execute()

Get invoice details.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/invoicingapi"
)

func main() {
	invoiceId := "5fa54d1e91867c03a0a7b4a4" // string | The unique resource identifier of the Invoice.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesInvoiceIdGet(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesInvoiceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesInvoiceIdGet`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesInvoiceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The unique resource identifier of the Invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesInvoiceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesInvoiceIdPayPost

> map[string]interface{} InvoicesInvoiceIdPayPost(ctx, invoiceId).Body(body).Execute()

Pay an invoice.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/invoicingapi"
)

func main() {
	invoiceId := "5fa54d1e91867c03a0a7b4a4" // string | The unique resource identifier of the Invoice.
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesInvoiceIdPayPost(context.Background(), invoiceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesInvoiceIdPayPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesInvoiceIdPayPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesInvoiceIdPayPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | The unique resource identifier of the Invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesInvoiceIdPayPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

