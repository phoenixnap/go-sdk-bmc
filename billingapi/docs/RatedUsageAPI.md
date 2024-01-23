# \RatedUsageAPI

All URIs are relative to *https://api.phoenixnap.com/billing/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RatedUsageGet**](RatedUsageAPI.md#RatedUsageGet) | **Get** /rated-usage | List the rated usage.
[**RatedUsageMonthToDateGet**](RatedUsageAPI.md#RatedUsageMonthToDateGet) | **Get** /rated-usage/month-to-date | List the rated usage records for the current calendar month.



## RatedUsageGet

> []RatedUsageGet200ResponseInner RatedUsageGet(ctx).FromYearMonth(fromYearMonth).ToYearMonth(toYearMonth).ProductCategory(productCategory).Execute()

List the rated usage.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi"
)

func main() {
	fromYearMonth := "2020-03" // string | From year month (inclusive) to filter rated usage records by.
	toYearMonth := "2020-04" // string | To year month (inclusive) to filter rated usage records by.
	productCategory := openapiclient.ProductCategoryEnum("SERVER") // ProductCategoryEnum | The product category (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatedUsageAPI.RatedUsageGet(context.Background()).FromYearMonth(fromYearMonth).ToYearMonth(toYearMonth).ProductCategory(productCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatedUsageAPI.RatedUsageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RatedUsageGet`: []RatedUsageGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RatedUsageAPI.RatedUsageGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRatedUsageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromYearMonth** | **string** | From year month (inclusive) to filter rated usage records by. | 
 **toYearMonth** | **string** | To year month (inclusive) to filter rated usage records by. | 
 **productCategory** | [**ProductCategoryEnum**](ProductCategoryEnum.md) | The product category | 

### Return type

[**[]RatedUsageGet200ResponseInner**](RatedUsageGet200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RatedUsageMonthToDateGet

> []RatedUsageGet200ResponseInner RatedUsageMonthToDateGet(ctx).ProductCategory(productCategory).Execute()

List the rated usage records for the current calendar month.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi"
)

func main() {
	productCategory := openapiclient.ProductCategoryEnum("SERVER") // ProductCategoryEnum | The product category (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatedUsageAPI.RatedUsageMonthToDateGet(context.Background()).ProductCategory(productCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatedUsageAPI.RatedUsageMonthToDateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RatedUsageMonthToDateGet`: []RatedUsageGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RatedUsageAPI.RatedUsageMonthToDateGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRatedUsageMonthToDateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCategory** | [**ProductCategoryEnum**](ProductCategoryEnum.md) | The product category | 

### Return type

[**[]RatedUsageGet200ResponseInner**](RatedUsageGet200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

