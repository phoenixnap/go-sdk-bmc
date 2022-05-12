# \IPBlocksApi

All URIs are relative to *https://api.phoenixnap.com/ips/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IpBlocksGet**](IPBlocksApi.md#IpBlocksGet) | **Get** /ip-blocks | List IP Blocks.
[**IpBlocksIpBlockIdDelete**](IPBlocksApi.md#IpBlocksIpBlockIdDelete) | **Delete** /ip-blocks/{ipBlockId} | Delete IP Block.
[**IpBlocksIpBlockIdGet**](IPBlocksApi.md#IpBlocksIpBlockIdGet) | **Get** /ip-blocks/{ipBlockId} | Get IP Block.
[**IpBlocksIpBlockIdPatch**](IPBlocksApi.md#IpBlocksIpBlockIdPatch) | **Patch** /ip-blocks/{ipBlockId} | Update IP block.
[**IpBlocksIpBlockIdTagsPut**](IPBlocksApi.md#IpBlocksIpBlockIdTagsPut) | **Put** /ip-blocks/{ipBlockId}/tags | Overwrite tags assigned for IP Block.
[**IpBlocksPost**](IPBlocksApi.md#IpBlocksPost) | **Post** /ip-blocks | Create an IP Block.



## IpBlocksGet

> []IpBlock IpBlocksGet(ctx).Tag(tag).Execute()

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
    tag := []string{"Inner_example"} // []string | List of tags, in the form tagName.tagValue, to filter by. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPBlocksApi.IpBlocksGet(context.Background()).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpBlocksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpBlocksGet`: []IpBlock
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpBlocksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIpBlocksGetRequest struct via the builder pattern

f
Name | Type | Description | Notes
---- | ---- | ----------- | ----- 
fN
 **tag** | **[]string** | List of tags, in the form tagName.tagValue, to filter by. | 

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPBlocksApi.IpBlocksIpBlockIdDelete(context.Background(), ipBlockId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpBlocksIpBlockIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpBlocksIpBlockIdDelete`: DeleteIpBlockResult
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpBlocksIpBlockIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**ipBlockId** | **string** | The IP Block identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpBlocksIpBlockIdDeleteRequest struct via the builder pattern

f
Name | Type | Description | Notes
---- | ---- | ----------- | ----- 
fN


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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPBlocksApi.IpBlocksIpBlockIdGet(context.Background(), ipBlockId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpBlocksIpBlockIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpBlocksIpBlockIdGet`: IpBlock
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpBlocksIpBlockIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**ipBlockId** | **string** | The IP Block identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpBlocksIpBlockIdGetRequest struct via the builder pattern

f
Name | Type | Description | Notes
---- | ---- | ----------- | ----- 
fN


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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPBlocksApi.IpBlocksIpBlockIdPatch(context.Background(), ipBlockId).IpBlockPatch(ipBlockPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpBlocksIpBlockIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpBlocksIpBlockIdPatch`: IpBlock
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpBlocksIpBlockIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**ipBlockId** | **string** | The IP Block identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpBlocksIpBlockIdPatchRequest struct via the builder pattern

f
Name | Type | Description | Notes
---- | ---- | ----------- | ----- 
fN

fN
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


## IpBlocksIpBlockIdTagsPut

> IpBlock IpBlocksIpBlockIdTagsPut(ctx, ipBlockId).TagAssignmentRequest(tagAssignmentRequest).Execute()

Overwrite tags assigned for IP Block.



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
    tagAssignmentRequest := []openapiclient.TagAssignmentRequest{*openapiclient.NewTagAssignmentRequest("Environment")} // []TagAssignmentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPBlocksApi.IpBlocksIpBlockIdTagsPut(context.Background(), ipBlockId).TagAssignmentRequest(tagAssignmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpBlocksIpBlockIdTagsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpBlocksIpBlockIdTagsPut`: IpBlock
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpBlocksIpBlockIdTagsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**ipBlockId** | **string** | The IP Block identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIpBlocksIpBlockIdTagsPutRequest struct via the builder pattern

f
Name | Type | Description | Notes
---- | ---- | ----------- | ----- 
fN

fN
 **tagAssignmentRequest** | [**[]TagAssignmentRequest**](TagAssignmentRequest.md) |  | 

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPBlocksApi.IpBlocksPost(context.Background()).IpBlockCreate(ipBlockCreate).Execute()
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

f
Name | Type | Description | Notes
---- | ---- | ----------- | ----- 
fN
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

