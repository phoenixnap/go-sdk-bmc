# \StorageNetworksApi

All URIs are relative to *https://api-staging.phoenixnap.com/network-storage/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageNetworksGet**](StorageNetworksApi.md#StorageNetworksGet) | **Get** /storage-networks | List all storage networks.
[**StorageNetworksIdDelete**](StorageNetworksApi.md#StorageNetworksIdDelete) | **Delete** /storage-networks/{storageNetworkId} | Delete a storage network and its volume.
[**StorageNetworksIdGet**](StorageNetworksApi.md#StorageNetworksIdGet) | **Get** /storage-networks/{storageNetworkId} | Get storage network details.
[**StorageNetworksIdPatch**](StorageNetworksApi.md#StorageNetworksIdPatch) | **Patch** /storage-networks/{storageNetworkId} | Update storage network details.
[**StorageNetworksPost**](StorageNetworksApi.md#StorageNetworksPost) | **Post** /storage-networks | Create a storage network and volume.
[**StorageNetworksStorageNetworkIdVolumesGet**](StorageNetworksApi.md#StorageNetworksStorageNetworkIdVolumesGet) | **Get** /storage-networks/{storageNetworkId}/volumes | Display one or more volumes belonging to a storage network.
[**StorageNetworksStorageNetworkIdVolumesVolumeIdGet**](StorageNetworksApi.md#StorageNetworksStorageNetworkIdVolumesVolumeIdGet) | **Get** /storage-networks/{storageNetworkId}/volumes/{volumeId} | Get a storage network&#39;s volume details.



## StorageNetworksGet

> []StorageNetwork StorageNetworksGet(ctx).Location(location).Execute()

List all storage networks.



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
    location := "PHX" // string | If present will filter the result by the given location. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksGet(context.Background()).Location(location).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageNetworksApi.StorageNetworksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageNetworksGet`: []StorageNetwork
    fmt.Fprintf(os.Stdout, "Response from `StorageNetworksApi.StorageNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageNetworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **string** | If present will filter the result by the given location. | 

### Return type

[**[]StorageNetwork**](StorageNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageNetworksIdDelete

> StorageNetworksIdDelete(ctx, storageNetworkId).Execute()

Delete a storage network and its volume.



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
    storageNetworkId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of storage network.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksIdDelete(context.Background(), storageNetworkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageNetworksApi.StorageNetworksIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageNetworkId** | **string** | ID of storage network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageNetworksIdDeleteRequest struct via the builder pattern


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


## StorageNetworksIdGet

> StorageNetwork StorageNetworksIdGet(ctx, storageNetworkId).Execute()

Get storage network details.



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
    storageNetworkId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of storage network.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksIdGet(context.Background(), storageNetworkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageNetworksApi.StorageNetworksIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageNetworksIdGet`: StorageNetwork
    fmt.Fprintf(os.Stdout, "Response from `StorageNetworksApi.StorageNetworksIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageNetworkId** | **string** | ID of storage network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageNetworksIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StorageNetwork**](StorageNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageNetworksIdPatch

> StorageNetwork StorageNetworksIdPatch(ctx, storageNetworkId).StorageNetworkUpdate(storageNetworkUpdate).Execute()

Update storage network details.



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
    storageNetworkId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of storage network.
    storageNetworkUpdate := *openapiclient.NewStorageNetworkUpdate() // StorageNetworkUpdate | Storage network to be updated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksIdPatch(context.Background(), storageNetworkId).StorageNetworkUpdate(storageNetworkUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageNetworksApi.StorageNetworksIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageNetworksIdPatch`: StorageNetwork
    fmt.Fprintf(os.Stdout, "Response from `StorageNetworksApi.StorageNetworksIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageNetworkId** | **string** | ID of storage network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageNetworksIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageNetworkUpdate** | [**StorageNetworkUpdate**](StorageNetworkUpdate.md) | Storage network to be updated. | 

### Return type

[**StorageNetwork**](StorageNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageNetworksPost

> StorageNetwork StorageNetworksPost(ctx).StorageNetworkCreate(storageNetworkCreate).Execute()

Create a storage network and volume.



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
    storageNetworkCreate := *openapiclient.NewStorageNetworkCreate("My storage network", "PHX", []openapiclient.VolumeCreate{*openapiclient.NewVolumeCreate("My volume name", int32(2000))}) // StorageNetworkCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksPost(context.Background()).StorageNetworkCreate(storageNetworkCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageNetworksApi.StorageNetworksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageNetworksPost`: StorageNetwork
    fmt.Fprintf(os.Stdout, "Response from `StorageNetworksApi.StorageNetworksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageNetworksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storageNetworkCreate** | [**StorageNetworkCreate**](StorageNetworkCreate.md) |  | 

### Return type

[**StorageNetwork**](StorageNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageNetworksStorageNetworkIdVolumesGet

> []Volume StorageNetworksStorageNetworkIdVolumesGet(ctx, storageNetworkId).Execute()

Display one or more volumes belonging to a storage network.



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
    storageNetworkId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of storage network.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesGet(context.Background(), storageNetworkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageNetworksStorageNetworkIdVolumesGet`: []Volume
    fmt.Fprintf(os.Stdout, "Response from `StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageNetworkId** | **string** | ID of storage network. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageNetworksStorageNetworkIdVolumesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Volume**](Volume.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageNetworksStorageNetworkIdVolumesVolumeIdGet

> Volume StorageNetworksStorageNetworkIdVolumesVolumeIdGet(ctx, storageNetworkId, volumeId).Execute()

Get a storage network's volume details.



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
    storageNetworkId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of storage network.
    volumeId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of volume.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdGet(context.Background(), storageNetworkId, volumeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageNetworksStorageNetworkIdVolumesVolumeIdGet`: Volume
    fmt.Fprintf(os.Stdout, "Response from `StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageNetworkId** | **string** | ID of storage network. | 
**volumeId** | **string** | ID of volume. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageNetworksStorageNetworkIdVolumesVolumeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Volume**](Volume.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

