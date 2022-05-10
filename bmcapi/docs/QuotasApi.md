# \QuotasApi

All URIs are relative to *https://api.phoenixnap.com/bmc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuotasGet**](QuotasApi.md#QuotasGet) | **Get** /quotas | List quotas
[**QuotasQuotaIdActionsRequestEditPost**](QuotasApi.md#QuotasQuotaIdActionsRequestEditPost) | **Post** /quotas/{quotaId}/actions/request-edit | Request quota limit change.
[**QuotasQuotaIdGet**](QuotasApi.md#QuotasQuotaIdGet) | **Get** /quotas/{quotaId} | Get a quota.



## QuotasGet

> []Quota QuotasGet(ctx).Execute()

List quotas



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotasApi.QuotasGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotasApi.QuotasGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuotasGet`: []Quota
    fmt.Fprintf(os.Stdout, "Response from `QuotasApi.QuotasGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiQuotasGetRequest struct via the builder pattern


### Return type

[**[]Quota**](Quota.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotasQuotaIdActionsRequestEditPost

> QuotasQuotaIdActionsRequestEditPost(ctx, quotaId).QuotaEditLimitRequest(quotaEditLimitRequest).Execute()

Request quota limit change.



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
    quotaId := "bmc.servers.max_count" // string | The ID of the Quota.
    quotaEditLimitRequest := *openapiclient.NewQuotaEditLimitRequest(int32(10), "I need more servers for my cluster.") // QuotaEditLimitRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotasApi.QuotasQuotaIdActionsRequestEditPost(context.Background(), quotaId).QuotaEditLimitRequest(quotaEditLimitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotasApi.QuotasQuotaIdActionsRequestEditPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quotaId** | **string** | The ID of the Quota. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuotasQuotaIdActionsRequestEditPostRequest struct via the builder pattern




Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quotaEditLimitRequest** | [**QuotaEditLimitRequest**](QuotaEditLimitRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotasQuotaIdGet

> Quota QuotasQuotaIdGet(ctx, quotaId).Execute()

Get a quota.



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
    quotaId := "bmc.servers.max_count" // string | The ID of the Quota.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotasApi.QuotasQuotaIdGet(context.Background(), quotaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotasApi.QuotasQuotaIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuotasQuotaIdGet`: Quota
    fmt.Fprintf(os.Stdout, "Response from `QuotasApi.QuotasQuotaIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quotaId** | **string** | The ID of the Quota. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuotasQuotaIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Quota**](Quota.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

