# \ProductsAPI

All URIs are relative to *https://api.phoenixnap.com/billing/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductsGet**](ProductsAPI.md#ProductsGet) | **Get** /products | List all Products.



## ProductsGet

> []ProductsGet200ResponseInner ProductsGet(ctx).ProductCode(productCode).ProductCategory(productCategory).SkuCode(skuCode).Location(location).Execute()

List all Products.



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
	productCode := "s1.c1.small" // string |  (optional)
	productCategory := "SERVER" // string |  (optional)
	skuCode := "xxx-xxx-xxx" // string |  (optional)
	location := "PHX" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsGet(context.Background()).ProductCode(productCode).ProductCategory(productCategory).SkuCode(skuCode).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsGet`: []ProductsGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCode** | **string** |  | 
 **productCategory** | **string** |  | 
 **skuCode** | **string** |  | 
 **location** | **string** |  | 

### Return type

[**[]ProductsGet200ResponseInner**](ProductsGet200ResponseInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

