# \PrivateNetworksAPI

All URIs are relative to *https://api.phoenixnap.com/networks/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrivateNetworksGet**](PrivateNetworksAPI.md#PrivateNetworksGet) | **Get** /private-networks | List Private Networks.
[**PrivateNetworksNetworkIdDelete**](PrivateNetworksAPI.md#PrivateNetworksNetworkIdDelete) | **Delete** /private-networks/{privateNetworkId} | Delete a Private Network.
[**PrivateNetworksNetworkIdGet**](PrivateNetworksAPI.md#PrivateNetworksNetworkIdGet) | **Get** /private-networks/{privateNetworkId} | Get a Private Network.
[**PrivateNetworksNetworkIdPut**](PrivateNetworksAPI.md#PrivateNetworksNetworkIdPut) | **Put** /private-networks/{privateNetworkId} | Update a Private Network.
[**PrivateNetworksPost**](PrivateNetworksAPI.md#PrivateNetworksPost) | **Post** /private-networks | Create a Private Network.



## PrivateNetworksGet

> []PrivateNetwork PrivateNetworksGet(ctx).Location(location).Execute()

List Private Networks.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func main() {
	location := "PHX" // string | If present will filter the result by the given location of the Private Networks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateNetworksAPI.PrivateNetworksGet(context.Background()).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksAPI.PrivateNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateNetworksGet`: []PrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksAPI.PrivateNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateNetworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **string** | If present will filter the result by the given location of the Private Networks. | 

### Return type

[**[]PrivateNetwork**](PrivateNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateNetworksNetworkIdDelete

> PrivateNetworksNetworkIdDelete(ctx, privateNetworkId).Execute()

Delete a Private Network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func main() {
	privateNetworkId := "603f3b2cfcaf050643b89a4b" // string | The private network identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PrivateNetworksAPI.PrivateNetworksNetworkIdDelete(context.Background(), privateNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksAPI.PrivateNetworksNetworkIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNetworkId** | **string** | The private network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateNetworksNetworkIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateNetworksNetworkIdGet

> PrivateNetwork PrivateNetworksNetworkIdGet(ctx, privateNetworkId).Execute()

Get a Private Network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func main() {
	privateNetworkId := "603f3b2cfcaf050643b89a4b" // string | The private network identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateNetworksAPI.PrivateNetworksNetworkIdGet(context.Background(), privateNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksAPI.PrivateNetworksNetworkIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateNetworksNetworkIdGet`: PrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksAPI.PrivateNetworksNetworkIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNetworkId** | **string** | The private network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateNetworksNetworkIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateNetwork**](PrivateNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateNetworksNetworkIdPut

> PrivateNetwork PrivateNetworksNetworkIdPut(ctx, privateNetworkId).PrivateNetworkModify(privateNetworkModify).Execute()

Update a Private Network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func main() {
	privateNetworkId := "603f3b2cfcaf050643b89a4b" // string | The private network identifier.
	privateNetworkModify := *openapiclient.NewPrivateNetworkModify("Sample network", true) // PrivateNetworkModify | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateNetworksAPI.PrivateNetworksNetworkIdPut(context.Background(), privateNetworkId).PrivateNetworkModify(privateNetworkModify).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksAPI.PrivateNetworksNetworkIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateNetworksNetworkIdPut`: PrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksAPI.PrivateNetworksNetworkIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNetworkId** | **string** | The private network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateNetworksNetworkIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateNetworkModify** | [**PrivateNetworkModify**](PrivateNetworkModify.md) |  | 

### Return type

[**PrivateNetwork**](PrivateNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateNetworksPost

> PrivateNetwork PrivateNetworksPost(ctx).PrivateNetworkCreate(privateNetworkCreate).Force(force).Execute()

Create a Private Network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func main() {
	privateNetworkCreate := *openapiclient.NewPrivateNetworkCreate("Sample Network", "PHX") // PrivateNetworkCreate | 
	force := true // bool | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateNetworksAPI.PrivateNetworksPost(context.Background()).PrivateNetworkCreate(privateNetworkCreate).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksAPI.PrivateNetworksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivateNetworksPost`: PrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksAPI.PrivateNetworksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateNetworksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **privateNetworkCreate** | [**PrivateNetworkCreate**](PrivateNetworkCreate.md) |  | 
 **force** | **bool** | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. | [default to false]

### Return type

[**PrivateNetwork**](PrivateNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

