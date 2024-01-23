# \StorageNetworksApi

All URIs are relative to *https://api.phoenixnap.com/network-storage/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageNetworksGet**](StorageNetworksApi.md#StorageNetworksGet) | **Get** /storage-networks | List all storage networks.
[**StorageNetworksIdDelete**](StorageNetworksApi.md#StorageNetworksIdDelete) | **Delete** /storage-networks/{storageId} | Delete a storage network and its volume.
[**StorageNetworksIdGet**](StorageNetworksApi.md#StorageNetworksIdGet) | **Get** /storage-networks/{storageId} | Get storage network details.
[**StorageNetworksIdPatch**](StorageNetworksApi.md#StorageNetworksIdPatch) | **Patch** /storage-networks/{storageId} | Update storage network details.
[**StorageNetworksPost**](StorageNetworksApi.md#StorageNetworksPost) | **Post** /storage-networks | Create a storage network and volume.
[**StorageNetworksStorageNetworkIdVolumesGet**](StorageNetworksApi.md#StorageNetworksStorageNetworkIdVolumesGet) | **Get** /storage-networks/{storageId}/volumes | Display one or more volumes belonging to a storage network.
[**StorageNetworksStorageNetworkIdVolumesPost**](StorageNetworksApi.md#StorageNetworksStorageNetworkIdVolumesPost) | **Post** /storage-networks/{storageId}/volumes | Create a volume belonging to a storage network.
[**StorageNetworksStorageNetworkIdVolumesVolumeIdDelete**](StorageNetworksApi.md#StorageNetworksStorageNetworkIdVolumesVolumeIdDelete) | **Delete** /storage-networks/{storageId}/volumes/{volumeId} | Delete a Storage Network&#39;s Volume
[**StorageNetworksStorageNetworkIdVolumesVolumeIdGet**](StorageNetworksApi.md#StorageNetworksStorageNetworkIdVolumesVolumeIdGet) | **Get** /storage-networks/{storageId}/volumes/{volumeId} | Get a storage network&#39;s volume details.
[**StorageNetworksStorageNetworkIdVolumesVolumeIdPatch**](StorageNetworksApi.md#StorageNetworksStorageNetworkIdVolumesVolumeIdPatch) | **Patch** /storage-networks/{storageId}/volumes/{volumeId} | Update a storage network&#39;s volume details.
[**StorageNetworksStorageNetworkIdVolumesVolumeIdTagsPut**](StorageNetworksApi.md#StorageNetworksStorageNetworkIdVolumesVolumeIdTagsPut) | **Put** /storage-networks/{storageId}/volumes/{volumeId}/tags | Overwrites tags assigned for the volume.



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

> StorageNetworksIdDelete(ctx, storageId).Execute()

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
    storageId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of the storage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksIdDelete(context.Background(), storageId).Execute()
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
**storageId** | **string** | ID of the storage. | 

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

> StorageNetwork StorageNetworksIdGet(ctx, storageId).Execute()

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
    storageId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of the storage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksIdGet(context.Background(), storageId).Execute()
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
**storageId** | **string** | ID of the storage. | 

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

> StorageNetwork StorageNetworksIdPatch(ctx, storageId).StorageNetworkUpdate(storageNetworkUpdate).Execute()

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
    storageId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of the storage.
    storageNetworkUpdate := *openapiclient.NewStorageNetworkUpdate() // StorageNetworkUpdate | Storage network to be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksIdPatch(context.Background(), storageId).StorageNetworkUpdate(storageNetworkUpdate).Execute()
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
**storageId** | **string** | ID of the storage. | 

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
    storageNetworkCreate := *openapiclient.NewStorageNetworkCreate("My storage network", "PHX", []openapiclient.StorageNetworkVolumeCreate{*openapiclient.NewStorageNetworkVolumeCreate("My volume name", int32(2000))}) // StorageNetworkCreate | 

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

> []Volume StorageNetworksStorageNetworkIdVolumesGet(ctx, storageId).Tag(tag).Execute()

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
    storageId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of the storage.
    tag := []string{"Inner_example"} // []string | A list of query parameters related to tags in the form of tagName.tagValue (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesGet(context.Background(), storageId).Tag(tag).Execute()
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
**storageId** | **string** | ID of the storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageNetworksStorageNetworkIdVolumesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **[]string** | A list of query parameters related to tags in the form of tagName.tagValue | 

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


## StorageNetworksStorageNetworkIdVolumesPost

> Volume StorageNetworksStorageNetworkIdVolumesPost(ctx, storageId).VolumeCreate(volumeCreate).Execute()

Create a volume belonging to a storage network.



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
    storageId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of the storage.
    volumeCreate := *openapiclient.NewVolumeCreate("My volume name", int32(2000)) // VolumeCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesPost(context.Background(), storageId).VolumeCreate(volumeCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageNetworksStorageNetworkIdVolumesPost`: Volume
    fmt.Fprintf(os.Stdout, "Response from `StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **string** | ID of the storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageNetworksStorageNetworkIdVolumesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeCreate** | [**VolumeCreate**](VolumeCreate.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageNetworksStorageNetworkIdVolumesVolumeIdDelete

> StorageNetworksStorageNetworkIdVolumesVolumeIdDelete(ctx, storageId, volumeId).Execute()

Delete a Storage Network's Volume



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
    storageId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of the storage.
    volumeId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of volume.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdDelete(context.Background(), storageId, volumeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **string** | ID of the storage. | 
**volumeId** | **string** | ID of volume. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageNetworksStorageNetworkIdVolumesVolumeIdDeleteRequest struct via the builder pattern


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


## StorageNetworksStorageNetworkIdVolumesVolumeIdGet

> Volume StorageNetworksStorageNetworkIdVolumesVolumeIdGet(ctx, storageId, volumeId).Execute()

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
    storageId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of the storage.
    volumeId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of volume.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdGet(context.Background(), storageId, volumeId).Execute()
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
**storageId** | **string** | ID of the storage. | 
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


## StorageNetworksStorageNetworkIdVolumesVolumeIdPatch

> Volume StorageNetworksStorageNetworkIdVolumesVolumeIdPatch(ctx, storageId, volumeId).VolumeUpdate(volumeUpdate).Execute()

Update a storage network's volume details.



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
    storageId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of the storage.
    volumeId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of volume.
    volumeUpdate := *openapiclient.NewVolumeUpdate() // VolumeUpdate | Storage network volume to be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdPatch(context.Background(), storageId, volumeId).VolumeUpdate(volumeUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageNetworksStorageNetworkIdVolumesVolumeIdPatch`: Volume
    fmt.Fprintf(os.Stdout, "Response from `StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **string** | ID of the storage. | 
**volumeId** | **string** | ID of volume. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageNetworksStorageNetworkIdVolumesVolumeIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **volumeUpdate** | [**VolumeUpdate**](VolumeUpdate.md) | Storage network volume to be updated. | 

### Return type

[**Volume**](Volume.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageNetworksStorageNetworkIdVolumesVolumeIdTagsPut

> Volume StorageNetworksStorageNetworkIdVolumesVolumeIdTagsPut(ctx, storageId, volumeId).TagAssignmentRequest(tagAssignmentRequest).Execute()

Overwrites tags assigned for the volume.



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
    storageId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of the storage.
    volumeId := "50dc434c-9bba-427b-bcd6-0bdba45c4dd2" // string | ID of volume.
    tagAssignmentRequest := []openapiclient.TagAssignmentRequest{*openapiclient.NewTagAssignmentRequest("Environment")} // []TagAssignmentRequest | Tags to assign to the volume.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdTagsPut(context.Background(), storageId, volumeId).TagAssignmentRequest(tagAssignmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdTagsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageNetworksStorageNetworkIdVolumesVolumeIdTagsPut`: Volume
    fmt.Fprintf(os.Stdout, "Response from `StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdTagsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storageId** | **string** | ID of the storage. | 
**volumeId** | **string** | ID of volume. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageNetworksStorageNetworkIdVolumesVolumeIdTagsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagAssignmentRequest** | [**[]TagAssignmentRequest**](TagAssignmentRequest.md) | Tags to assign to the volume. | 

### Return type

[**Volume**](Volume.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

