# \LocationsAPI

All URIs are relative to *https://api.phoenixnap.com/location-api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocations**](LocationsAPI.md#GetLocations) | **Get** /locations | Get All Locations



## GetLocations

> []Location GetLocations(ctx).Location(location).ProductCategory(productCategory).Execute()

Get All Locations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/locationapi"
)

func main() {
	location := openapiclient.LocationEnum("PHX") // LocationEnum | Location of interest (optional)
	productCategory := openapiclient.ProductCategoryEnum("SERVER") // ProductCategoryEnum | Product category of interest (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsAPI.GetLocations(context.Background()).Location(location).ProductCategory(productCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.GetLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocations`: []Location
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.GetLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | [**LocationEnum**](LocationEnum.md) | Location of interest | 
 **productCategory** | [**ProductCategoryEnum**](ProductCategoryEnum.md) | Product category of interest | 

### Return type

[**[]Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

