# \PrivateNetworksApi

All URIs are relative to *https://api.phoenixnap.com/networks/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrivateNetworksGet**](PrivateNetworksApi.md#PrivateNetworksGet) | **Get** /private-networks | List Private Networks.
[**PrivateNetworksNetworkIdDelete**](PrivateNetworksApi.md#PrivateNetworksNetworkIdDelete) | **Delete** /private-networks/{privateNetworkId} | Delete a Private Network.
[**PrivateNetworksNetworkIdGet**](PrivateNetworksApi.md#PrivateNetworksNetworkIdGet) | **Get** /private-networks/{privateNetworkId} | Get a Private Network.
[**PrivateNetworksNetworkIdPut**](PrivateNetworksApi.md#PrivateNetworksNetworkIdPut) | **Put** /private-networks/{privateNetworkId} | Update a Private Network.
[**PrivateNetworksPost**](PrivateNetworksApi.md#PrivateNetworksPost) | **Post** /private-networks | Create a Private Network.



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
    openapiclient "./openapi"
)

func main() {
    location := "PHX" // string | If present will filter the result by the given location of the Private Networks. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateNetworksApi.PrivateNetworksGet(context.Background()).Location(location).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksApi.PrivateNetworksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateNetworksGet`: []PrivateNetwork
    fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksApi.PrivateNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateNetworksGetRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----
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
    openapiclient "./openapi"
)

func main() {
    privateNetworkId := "603f3b2cfcaf050643b89a4b" // string | The private network identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateNetworksApi.PrivateNetworksNetworkIdDelete(context.Background(), privateNetworkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksApi.PrivateNetworksNetworkIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**privateNetworkId** | **string** | The private network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateNetworksNetworkIdDeleteRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----


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
    openapiclient "./openapi"
)

func main() {
    privateNetworkId := "603f3b2cfcaf050643b89a4b" // string | The private network identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateNetworksApi.PrivateNetworksNetworkIdGet(context.Background(), privateNetworkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksApi.PrivateNetworksNetworkIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateNetworksNetworkIdGet`: PrivateNetwork
    fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksApi.PrivateNetworksNetworkIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**privateNetworkId** | **string** | The private network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateNetworksNetworkIdGetRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----


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
    openapiclient "./openapi"
)

func main() {
    privateNetworkId := "603f3b2cfcaf050643b89a4b" // string | The private network identifier.
    privateNetworkModify := *openapiclient.NewPrivateNetworkModify("Sample network", true) // PrivateNetworkModify |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateNetworksApi.PrivateNetworksNetworkIdPut(context.Background(), privateNetworkId).PrivateNetworkModify(privateNetworkModify).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksApi.PrivateNetworksNetworkIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateNetworksNetworkIdPut`: PrivateNetwork
    fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksApi.PrivateNetworksNetworkIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**privateNetworkId** | **string** | The private network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateNetworksNetworkIdPutRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----

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

> PrivateNetwork PrivateNetworksPost(ctx).PrivateNetworkCreate(privateNetworkCreate).Execute()

Create a Private Network.



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
    privateNetworkCreate := *openapiclient.NewPrivateNetworkCreate("Sample Network", "PHX", "10.0.0.0/24") // PrivateNetworkCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateNetworksApi.PrivateNetworksPost(context.Background()).PrivateNetworkCreate(privateNetworkCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksApi.PrivateNetworksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateNetworksPost`: PrivateNetwork
    fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksApi.PrivateNetworksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateNetworksPostRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----
 **privateNetworkCreate** | [**PrivateNetworkCreate**](PrivateNetworkCreate.md) |  | 

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

