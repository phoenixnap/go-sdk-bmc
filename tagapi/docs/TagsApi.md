# \TagsApi

All URIs are relative to *https://api.phoenixnap.com/tag-manager/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TagsGet**](TagsApi.md#TagsGet) | **Get** /tags | List tags.
[**TagsPost**](TagsApi.md#TagsPost) | **Post** /tags | Create a Tag.
[**TagsTagIdDelete**](TagsApi.md#TagsTagIdDelete) | **Delete** /tags/{tagId} | Delete a Tag.
[**TagsTagIdGet**](TagsApi.md#TagsTagIdGet) | **Get** /tags/{tagId} | Get a Tag.
[**TagsTagIdPatch**](TagsApi.md#TagsTagIdPatch) | **Patch** /tags/{tagId} | Modify a Tag.



## TagsGet

> []Tag TagsGet(ctx).Name(name).Execute()

List tags.



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
    name := "env" // string | Query a tag by its name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.TagsGet(context.Background()).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.TagsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TagsGet`: []Tag
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.TagsGet`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiTagsGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
|  **name** | **string** | Query a tag by its name. |  |

### Return type

[**[]Tag**](Tag.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagsPost

> Tag TagsPost(ctx).TagCreate(tagCreate).Execute()

Create a Tag.



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
    tagCreate := *openapiclient.NewTagCreate("Environment", true) // TagCreate | The body containing the tag details.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.TagsPost(context.Background()).TagCreate(tagCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.TagsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TagsPost`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.TagsPost`: %v\n", resp)
}
```

### Path Parameters

 |

### Other Parameters

Other parameters are passed through a pointer to a apiTagsPostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
|  **tagCreate** | [**TagCreate**](TagCreate.md) | The body containing the tag details. |  |

### Return type

[**Tag**](Tag.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagsTagIdDelete

> DeleteResult TagsTagIdDelete(ctx, tagId).Execute()

Delete a Tag.



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
    tagId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The tag's ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.TagsTagIdDelete(context.Background(), tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.TagsTagIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TagsTagIdDelete`: DeleteResult
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.TagsTagIdDelete`: %v\n", resp)
}
```

### Path Parameters


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **tagId** | **string** | The tag&#39;s ID. |  |

### Other Parameters

Other parameters are passed through a pointer to a apiTagsTagIdDeleteRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
|  |

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


## TagsTagIdGet

> Tag TagsTagIdGet(ctx, tagId).Execute()

Get a Tag.



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
    tagId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The tag's ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.TagsTagIdGet(context.Background(), tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.TagsTagIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TagsTagIdGet`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.TagsTagIdGet`: %v\n", resp)
}
```

### Path Parameters


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **tagId** | **string** | The tag&#39;s ID. |  |

### Other Parameters

Other parameters are passed through a pointer to a apiTagsTagIdGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
|  |

### Return type

[**Tag**](Tag.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagsTagIdPatch

> Tag TagsTagIdPatch(ctx, tagId).TagUpdate(tagUpdate).Execute()

Modify a Tag.



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
    tagId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The tag's ID.
    tagUpdate := *openapiclient.NewTagUpdate("Environment", true) // TagUpdate | The body containing the tag changes.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.TagsTagIdPatch(context.Background(), tagId).TagUpdate(tagUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.TagsTagIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TagsTagIdPatch`: Tag
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.TagsTagIdPatch`: %v\n", resp)
}
```

### Path Parameters


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **tagId** | **string** | The tag&#39;s ID. |  |

### Other Parameters

Other parameters are passed through a pointer to a apiTagsTagIdPatchRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
|  |
|  **tagUpdate** | [**TagUpdate**](TagUpdate.md) | The body containing the tag changes. |  |

### Return type

[**Tag**](Tag.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

