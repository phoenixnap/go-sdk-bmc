# \PublicNetworksApi

All URIs are relative to *https://api.phoenixnap.com/networks/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicNetworksGet**](PublicNetworksApi.md#PublicNetworksGet) | **Get** /public-networks | List Public Networks.
[**PublicNetworksNetworkIdDelete**](PublicNetworksApi.md#PublicNetworksNetworkIdDelete) | **Delete** /public-networks/{publicNetworkId} | Delete a Public Network.
[**PublicNetworksNetworkIdGet**](PublicNetworksApi.md#PublicNetworksNetworkIdGet) | **Get** /public-networks/{publicNetworkId} | Get a Public Network.
[**PublicNetworksNetworkIdIpBlocksIpBlockIdDelete**](PublicNetworksApi.md#PublicNetworksNetworkIdIpBlocksIpBlockIdDelete) | **Delete** /public-networks/{publicNetworkId}/ip-blocks/{ipBlockId} | Removes the IP Block from the Public Network.
[**PublicNetworksNetworkIdIpBlocksPost**](PublicNetworksApi.md#PublicNetworksNetworkIdIpBlocksPost) | **Post** /public-networks/{publicNetworkId}/ip-blocks | Adds an IP block to this public network.
[**PublicNetworksNetworkIdPatch**](PublicNetworksApi.md#PublicNetworksNetworkIdPatch) | **Patch** /public-networks/{publicNetworkId} | Update Public Network&#39;s Details.
[**PublicNetworksPost**](PublicNetworksApi.md#PublicNetworksPost) | **Post** /public-networks | Create a public network.



## PublicNetworksGet

> []PublicNetwork PublicNetworksGet(ctx).Location(location).Execute()

List Public Networks.



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
    location := "PHX" // string | If present will filter the result by the given location of the Public Networks. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicNetworksApi.PublicNetworksGet(context.Background()).Location(location).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksApi.PublicNetworksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicNetworksGet`: []PublicNetwork
    fmt.Fprintf(os.Stdout, "Response from `PublicNetworksApi.PublicNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksGetRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----  
**location** | **string** | If present will filter the result by the given location of the Public Networks. | 

### Return type

[**[]PublicNetwork**](PublicNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicNetworksNetworkIdDelete

> PublicNetworksNetworkIdDelete(ctx, publicNetworkId).Execute()

Delete a Public Network.



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
    publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicNetworksApi.PublicNetworksNetworkIdDelete(context.Background(), publicNetworkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksApi.PublicNetworksNetworkIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicNetworkId** | **string** | The Public Network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksNetworkIdDeleteRequest struct via the builder pattern

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


## PublicNetworksNetworkIdGet

> PublicNetwork PublicNetworksNetworkIdGet(ctx, publicNetworkId).Execute()

Get a Public Network.



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
    publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicNetworksApi.PublicNetworksNetworkIdGet(context.Background(), publicNetworkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksApi.PublicNetworksNetworkIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicNetworksNetworkIdGet`: PublicNetwork
    fmt.Fprintf(os.Stdout, "Response from `PublicNetworksApi.PublicNetworksNetworkIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicNetworkId** | **string** | The Public Network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksNetworkIdGetRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----  

### Return type

[**PublicNetwork**](PublicNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicNetworksNetworkIdIpBlocksIpBlockIdDelete

> string PublicNetworksNetworkIdIpBlocksIpBlockIdDelete(ctx, publicNetworkId, ipBlockId).Execute()

Removes the IP Block from the Public Network.



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
    publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.
    ipBlockId := "6047127fed34ecc3ba8402d2" // string | The IP Block identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicNetworksApi.PublicNetworksNetworkIdIpBlocksIpBlockIdDelete(context.Background(), publicNetworkId, ipBlockId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksApi.PublicNetworksNetworkIdIpBlocksIpBlockIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicNetworksNetworkIdIpBlocksIpBlockIdDelete`: string
    fmt.Fprintf(os.Stdout, "Response from `PublicNetworksApi.PublicNetworksNetworkIdIpBlocksIpBlockIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicNetworkId** | **string** | The Public Network identifier. |  |
**ipBlockId** | **string** | The IP Block identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksNetworkIdIpBlocksIpBlockIdDeleteRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----  
 

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicNetworksNetworkIdIpBlocksPost

> PublicNetworkIpBlock PublicNetworksNetworkIdIpBlocksPost(ctx, publicNetworkId).PublicNetworkIpBlock(publicNetworkIpBlock).Execute()

Adds an IP block to this public network.



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
    publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.
    publicNetworkIpBlock := *openapiclient.NewPublicNetworkIpBlock("60473a6115e34466c9f8f083") // PublicNetworkIpBlock |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicNetworksApi.PublicNetworksNetworkIdIpBlocksPost(context.Background(), publicNetworkId).PublicNetworkIpBlock(publicNetworkIpBlock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksApi.PublicNetworksNetworkIdIpBlocksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicNetworksNetworkIdIpBlocksPost`: PublicNetworkIpBlock
    fmt.Fprintf(os.Stdout, "Response from `PublicNetworksApi.PublicNetworksNetworkIdIpBlocksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicNetworkId** | **string** | The Public Network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksNetworkIdIpBlocksPostRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----  
 
**publicNetworkIpBlock** | [**PublicNetworkIpBlock**](PublicNetworkIpBlock.md) |  | 

### Return type

[**PublicNetworkIpBlock**](PublicNetworkIpBlock.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicNetworksNetworkIdPatch

> PublicNetwork PublicNetworksNetworkIdPatch(ctx, publicNetworkId).PublicNetworkModify(publicNetworkModify).Execute()

Update Public Network's Details.



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
    publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.
    publicNetworkModify := *openapiclient.NewPublicNetworkModify() // PublicNetworkModify |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicNetworksApi.PublicNetworksNetworkIdPatch(context.Background(), publicNetworkId).PublicNetworkModify(publicNetworkModify).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksApi.PublicNetworksNetworkIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicNetworksNetworkIdPatch`: PublicNetwork
    fmt.Fprintf(os.Stdout, "Response from `PublicNetworksApi.PublicNetworksNetworkIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicNetworkId** | **string** | The Public Network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksNetworkIdPatchRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----  
 
**publicNetworkModify** | [**PublicNetworkModify**](PublicNetworkModify.md) |  | 

### Return type

[**PublicNetwork**](PublicNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicNetworksPost

> PublicNetwork PublicNetworksPost(ctx).PublicNetworkCreate(publicNetworkCreate).Execute()

Create a public network.



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
    publicNetworkCreate := *openapiclient.NewPublicNetworkCreate("Sample Network", "PHX") // PublicNetworkCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicNetworksApi.PublicNetworksPost(context.Background()).PublicNetworkCreate(publicNetworkCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksApi.PublicNetworksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicNetworksPost`: PublicNetwork
    fmt.Fprintf(os.Stdout, "Response from `PublicNetworksApi.PublicNetworksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksPostRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----  
**publicNetworkCreate** | [**PublicNetworkCreate**](PublicNetworkCreate.md) |  | 

### Return type

[**PublicNetwork**](PublicNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

