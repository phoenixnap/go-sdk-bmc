# \ProductsApi

All URIs are relative to *https://api.phoenixnap.com/billing/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductAvailabilityGet**](ProductsApi.md#ProductAvailabilityGet) | **Get** /product-availability | List all Product availabilities.
[**ProductsGet**](ProductsApi.md#ProductsGet) | **Get** /products | List all Products.



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
    openapiclient "./openapi"
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
    resp, r, err := apiClient.ProductsApi.ProductAvailabilityGet(context.Background()).ProductCategory(productCategory).ProductCode(productCode).ShowOnlyMinQuantityAvailable(showOnlyMinQuantityAvailable).Location(location).Solution(solution).MinQuantity(minQuantity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ProductAvailabilityGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductAvailabilityGet`: []ProductAvailability
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ProductAvailabilityGet`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    productCode := "d1.tiny" // string |  (optional)
    productCategory := "server" // string |  (optional)
    skuCode := "xxx-xxx-xxx" // string |  (optional)
    location := "PHX" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.ProductsGet(context.Background()).ProductCode(productCode).ProductCategory(productCategory).SkuCode(skuCode).Location(location).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ProductsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductsGet`: []ProductsGet200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ProductsGet`: %v\n", resp)
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

