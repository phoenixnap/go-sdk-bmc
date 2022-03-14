# \IPBlocksApi

All URIs are relative to *https://api.phoenixnap.com/ips/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IpBlocksGet**](IPBlocksApi.md#IpBlocksGet) | **Get** /ip-blocks | List IP Blocks.
[**IpBlocksIpBlockIdDelete**](IPBlocksApi.md#IpBlocksIpBlockIdDelete) | **Delete** /ip-blocks/{ipBlockId} | Delete IP Block.
[**IpBlocksIpBlockIdGet**](IPBlocksApi.md#IpBlocksIpBlockIdGet) | **Get** /ip-blocks/{ipBlockId} | Get IP Block.
[**IpBlocksIpBlockIdPatch**](IPBlocksApi.md#IpBlocksIpBlockIdPatch) | **Patch** /ip-blocks/{ipBlockId} | Update IP block.
[**IpBlocksPost**](IPBlocksApi.md#IpBlocksPost) | **Post** /ip-blocks | Create an IP Block.



## IpBlocksGet

> []IpBlock IpBlocksGet(ctx).Execute()

List IP Blocks.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPBlocksApi.IpBlocksGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpBlocksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpBlocksGet`: []IpBlock
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpBlocksGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIpBlocksGetRequest struct via the builder pattern


### Return type

[**[]IpBlock**](IpBlock.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpBlocksIpBlockIdDelete

> DeleteIpBlockResult IpBlocksIpBlockIdDelete(ctx, ipBlockId).Execute()

Delete IP Block.



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
    ipBlockId := "6047127fed34ecc3ba8402d2" // string | The IP Block identifier.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPBlocksApi.IpBlocksIpBlockIdDelete(context.Background(), ipBlockId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpBlocksIpBlockIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpBlocksIpBlockIdDelete`: DeleteIpBlockResult
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpBlocksIpBlockIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipBlockId** | **string** | The IP Block identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpBlocksIpBlockIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteIpBlockResult**](DeleteIpBlockResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpBlocksIpBlockIdGet

> IpBlock IpBlocksIpBlockIdGet(ctx, ipBlockId).Execute()

Get IP Block.



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
    ipBlockId := "6047127fed34ecc3ba8402d2" // string | The IP Block identifier.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPBlocksApi.IpBlocksIpBlockIdGet(context.Background(), ipBlockId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpBlocksIpBlockIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpBlocksIpBlockIdGet`: IpBlock
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpBlocksIpBlockIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipBlockId** | **string** | The IP Block identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpBlocksIpBlockIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpBlocksIpBlockIdPatch

> IpBlock IpBlocksIpBlockIdPatch(ctx, ipBlockId).IpBlockPatch(ipBlockPatch).Execute()

Update IP block.



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
    ipBlockId := "6047127fed34ecc3ba8402d2" // string | The IP Block identifier.
    ipBlockPatch := *openapiclient.NewIpBlockPatch() // IpBlockPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPBlocksApi.IpBlocksIpBlockIdPatch(context.Background(), ipBlockId).IpBlockPatch(ipBlockPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpBlocksIpBlockIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpBlocksIpBlockIdPatch`: IpBlock
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpBlocksIpBlockIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipBlockId** | **string** | The IP Block identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpBlocksIpBlockIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipBlockPatch** | [**IpBlockPatch**](IpBlockPatch.md) |  | 

### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IpBlocksPost

> IpBlock IpBlocksPost(ctx).IpBlockCreate(ipBlockCreate).Execute()

Create an IP Block.



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
    ipBlockCreate := *openapiclient.NewIpBlockCreate("PHX", "/30") // IpBlockCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPBlocksApi.IpBlocksPost(context.Background()).IpBlockCreate(ipBlockCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpBlocksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpBlocksPost`: IpBlock
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpBlocksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIpBlocksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipBlockCreate** | [**IpBlockCreate**](IpBlockCreate.md) |  | 

### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

