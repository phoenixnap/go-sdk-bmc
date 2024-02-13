# \ProductAvailabilityAPI

All URIs are relative to *https://api.phoenixnap.com/billing/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductAvailabilityGet**](ProductAvailabilityAPI.md#ProductAvailabilityGet) | **Get** /product-availability | List all Product availabilities.



## ProductAvailabilityGet

> []ProductAvailability ProductAvailabilityGet(ctx).ProductCategory(productCategory).ProductCode(productCode).ShowOnlyMinQuantityAvailable(showOnlyMinQuantityAvailable).Location(location).Solution(solution).MinQuantity(minQuantity).Execute()

List all Product availabilities.



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
	productCategory := []string{"ProductCategory_example"} // []string | Product category. Currently only SERVER category is supported. (optional)
	productCode := []string{"Inner_example"} // []string |  (optional)
	showOnlyMinQuantityAvailable := true // bool | Show only locations where product with requested quantity is available or all locations where product is offered. (optional) (default to true)
	location := []openapiclient.LocationEnum{openapiclient.LocationEnum("PHX")} // []LocationEnum |  (optional)
	solution := []string{"Solution_example"} // []string |  (optional)
	minQuantity := float32(2) // float32 | Minimal quantity of product needed. Minimum, maximum and default values might differ for different products. For servers, they are 1, 10 and 1 respectively. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAvailabilityAPI.ProductAvailabilityGet(context.Background()).ProductCategory(productCategory).ProductCode(productCode).ShowOnlyMinQuantityAvailable(showOnlyMinQuantityAvailable).Location(location).Solution(solution).MinQuantity(minQuantity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAvailabilityAPI.ProductAvailabilityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductAvailabilityGet`: []ProductAvailability
	fmt.Fprintf(os.Stdout, "Response from `ProductAvailabilityAPI.ProductAvailabilityGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAvailabilityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCategory** | **[]string** | Product category. Currently only SERVER category is supported. | 
 **productCode** | **[]string** |  | 
 **showOnlyMinQuantityAvailable** | **bool** | Show only locations where product with requested quantity is available or all locations where product is offered. | [default to true]
 **location** | [**[]LocationEnum**](LocationEnum.md) |  | 
 **solution** | **[]string** |  | 
 **minQuantity** | **float32** | Minimal quantity of product needed. Minimum, maximum and default values might differ for different products. For servers, they are 1, 10 and 1 respectively. | 

### Return type

[**[]ProductAvailability**](ProductAvailability.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

