# \ClustersApi

All URIs are relative to *https://api.phoenixnap.com/solutions/rancher/v1beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersGet**](ClustersApi.md#ClustersGet) | **Get** /clusters | Cluster list.
[**ClustersIdDelete**](ClustersApi.md#ClustersIdDelete) | **Delete** /clusters/{id} | Delete a cluster.
[**ClustersIdGet**](ClustersApi.md#ClustersIdGet) | **Get** /clusters/{id} | Retrieve a Cluster
[**ClustersPost**](ClustersApi.md#ClustersPost) | **Post** /clusters | Create a Rancher Server Deployment.



## ClustersGet

> []Cluster ClustersGet(ctx).Execute()

Cluster list.



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
    resp, r, err := apiClient.ClustersApi.ClustersGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersGet`: []Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClustersGetRequest struct via the builder pattern



### Return type

[**[]Cluster**](Cluster.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersIdDelete

> DeleteResult ClustersIdDelete(ctx, id).Execute()

Delete a cluster.



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
    id := "123" // string | The Cluster identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ClustersIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersIdDelete`: DeleteResult
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**id** | **string** | The Cluster identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteResult**](DeleteResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersIdGet

> Cluster ClustersIdGet(ctx, id).Execute()

Retrieve a Cluster



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
    id := "123" // string | The Cluster identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ClustersIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersIdGet`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**id** | **string** | The Cluster identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClustersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Cluster**](Cluster.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClustersPost

> Cluster ClustersPost(ctx).Cluster(cluster).Execute()

Create a Rancher Server Deployment.



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
    cluster := *openapiclient.NewCluster("PHX") // Cluster |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ClustersPost(context.Background()).Cluster(cluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClustersPost`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClustersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | [**Cluster**](Cluster.md) |  | 


### Return type

[**Cluster**](Cluster.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

