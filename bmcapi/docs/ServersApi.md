# \ServersApi

All URIs are relative to *https://api.phoenixnap.com/bmc/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePrivateNetwork**](ServersApi.md#DeletePrivateNetwork) | **Delete** /servers/{serverId}/network-configuration/private-network-configuration/private-networks/{privateNetworkId} | Removes the server from private network.
[**ServersGet**](ServersApi.md#ServersGet) | **Get** /servers | List servers.
[**ServersPost**](ServersApi.md#ServersPost) | **Post** /servers | Create new server.
[**ServersServerIdActionsPowerOffPost**](ServersApi.md#ServersServerIdActionsPowerOffPost) | **Post** /servers/{serverId}/actions/power-off | Power off server.
[**ServersServerIdActionsPowerOnPost**](ServersApi.md#ServersServerIdActionsPowerOnPost) | **Post** /servers/{serverId}/actions/power-on | Power on server.
[**ServersServerIdActionsRebootPost**](ServersApi.md#ServersServerIdActionsRebootPost) | **Post** /servers/{serverId}/actions/reboot | Reboot server.
[**ServersServerIdActionsReservePost**](ServersApi.md#ServersServerIdActionsReservePost) | **Post** /servers/{serverId}/actions/reserve | Reserve server.
[**ServersServerIdActionsResetPost**](ServersApi.md#ServersServerIdActionsResetPost) | **Post** /servers/{serverId}/actions/reset | Reset server.
[**ServersServerIdActionsShutdownPost**](ServersApi.md#ServersServerIdActionsShutdownPost) | **Post** /servers/{serverId}/actions/shutdown | Shutdown server.
[**ServersServerIdDelete**](ServersApi.md#ServersServerIdDelete) | **Delete** /servers/{serverId} | Delete server.
[**ServersServerIdGet**](ServersApi.md#ServersServerIdGet) | **Get** /servers/{serverId} | Get server.
[**ServersServerIdPatch**](ServersApi.md#ServersServerIdPatch) | **Patch** /servers/{serverId} | Patch a Server.
[**ServersServerIdPrivateNetworksPost**](ServersApi.md#ServersServerIdPrivateNetworksPost) | **Post** /servers/{serverId}/network-configuration/private-network-configuration/private-networks | Adds the server to a private network.
[**ServersServerIdTagsPut**](ServersApi.md#ServersServerIdTagsPut) | **Put** /servers/{serverId}/tags | Set server tags.



## DeletePrivateNetwork

> string DeletePrivateNetwork(ctx, serverId, privateNetworkId).Execute()

Removes the server from private network.



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
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    privateNetworkId := "603f3b2cfcaf050643b89a4b" // string | The private network identifier.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.DeletePrivateNetwork(context.Background(), serverId, privateNetworkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.DeletePrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePrivateNetwork`: string
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.DeletePrivateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 
**privateNetworkId** | **string** | The private network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ServersGet

> []Server ServersGet(ctx).Tag(tag).Execute()

List servers.



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
    tag := []string{"Inner_example"} // []string | A list of queryparameters related to tags in the form of tagName.tagValue (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.ServersGet(context.Background()).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersGet`: []Server
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **[]string** | A list of queryparameters related to tags in the form of tagName.tagValue | 

### Return type

[**[]Server**](Server.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersPost

> Server ServersPost(ctx).ServerCreate(serverCreate).Execute()

Create new server.



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
    serverCreate := *openapiclient.NewServerCreate("my-server-1", "ubuntu/bionic", "s1.c1.small", "PHX") // ServerCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.ServersPost(context.Background()).ServerCreate(serverCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersPost`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverCreate** | [**ServerCreate**](ServerCreate.md) |  | 

### Return type

[**Server**](Server.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdActionsPowerOffPost

> ActionResult ServersServerIdActionsPowerOffPost(ctx, serverId).Execute()

Power off server.



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
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.ServersServerIdActionsPowerOffPost(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdActionsPowerOffPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdActionsPowerOffPost`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdActionsPowerOffPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsPowerOffPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResult**](ActionResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdActionsPowerOnPost

> ActionResult ServersServerIdActionsPowerOnPost(ctx, serverId).Execute()

Power on server.



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
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.ServersServerIdActionsPowerOnPost(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdActionsPowerOnPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdActionsPowerOnPost`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdActionsPowerOnPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsPowerOnPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResult**](ActionResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdActionsRebootPost

> ActionResult ServersServerIdActionsRebootPost(ctx, serverId).Execute()

Reboot server.



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
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.ServersServerIdActionsRebootPost(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdActionsRebootPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdActionsRebootPost`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdActionsRebootPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsRebootPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResult**](ActionResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdActionsReservePost

> Server ServersServerIdActionsReservePost(ctx, serverId).ServerReserve(serverReserve).Execute()

Reserve server.



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
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    serverReserve := *openapiclient.NewServerReserve("ONE_MONTH_RESERVATION") // ServerReserve |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.ServersServerIdActionsReservePost(context.Background(), serverId).ServerReserve(serverReserve).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdActionsReservePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdActionsReservePost`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdActionsReservePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsReservePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverReserve** | [**ServerReserve**](ServerReserve.md) |  | 

### Return type

[**Server**](Server.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdActionsResetPost

> ResetResult ServersServerIdActionsResetPost(ctx, serverId).ServerReset(serverReset).Execute()

Reset server.



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
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    serverReset := *openapiclient.NewServerReset() // ServerReset |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.ServersServerIdActionsResetPost(context.Background(), serverId).ServerReset(serverReset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdActionsResetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdActionsResetPost`: ResetResult
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdActionsResetPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverReset** | [**ServerReset**](ServerReset.md) |  | 

### Return type

[**ResetResult**](ResetResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdActionsShutdownPost

> ActionResult ServersServerIdActionsShutdownPost(ctx, serverId).Execute()

Shutdown server.



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
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.ServersServerIdActionsShutdownPost(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdActionsShutdownPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdActionsShutdownPost`: ActionResult
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdActionsShutdownPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsShutdownPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResult**](ActionResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdDelete

> DeleteResult ServersServerIdDelete(ctx, serverId).Execute()

Delete server.



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
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.ServersServerIdDelete(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdDelete`: DeleteResult
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdDeleteRequest struct via the builder pattern


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


## ServersServerIdGet

> Server ServersServerIdGet(ctx, serverId).Execute()

Get server.



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
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.ServersServerIdGet(context.Background(), serverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdGet`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Server**](Server.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdPatch

> Server ServersServerIdPatch(ctx, serverId).ServerPatch(serverPatch).Execute()

Patch a Server.



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
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    serverPatch := *openapiclient.NewServerPatch() // ServerPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.ServersServerIdPatch(context.Background(), serverId).ServerPatch(serverPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdPatch`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverPatch** | [**ServerPatch**](ServerPatch.md) |  | 

### Return type

[**Server**](Server.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdPrivateNetworksPost

> ServerPrivateNetwork ServersServerIdPrivateNetworksPost(ctx, serverId).ServerPrivateNetwork(serverPrivateNetwork).Execute()

Adds the server to a private network.



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
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    serverPrivateNetwork := *openapiclient.NewServerPrivateNetwork("603f3b2cfcaf050643b89a4b") // ServerPrivateNetwork |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.ServersServerIdPrivateNetworksPost(context.Background(), serverId).ServerPrivateNetwork(serverPrivateNetwork).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdPrivateNetworksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdPrivateNetworksPost`: ServerPrivateNetwork
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdPrivateNetworksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdPrivateNetworksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverPrivateNetwork** | [**ServerPrivateNetwork**](ServerPrivateNetwork.md) |  | 

### Return type

[**ServerPrivateNetwork**](ServerPrivateNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdTagsPut

> Server ServersServerIdTagsPut(ctx, serverId).TagAssignmentRequest(tagAssignmentRequest).Execute()

Set server tags.



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
    serverId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The server's ID.
    tagAssignmentRequest := []openapiclient.TagAssignmentRequest{*openapiclient.NewTagAssignmentRequest("Environment")} // []TagAssignmentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServersApi.ServersServerIdTagsPut(context.Background(), serverId).TagAssignmentRequest(tagAssignmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersApi.ServersServerIdTagsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServersServerIdTagsPut`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServersApi.ServersServerIdTagsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdTagsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagAssignmentRequest** | [**[]TagAssignmentRequest**](TagAssignmentRequest.md) |  | 

### Return type

[**Server**](Server.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
