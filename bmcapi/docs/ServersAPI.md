# \ServersAPI

All URIs are relative to *https://api.phoenixnap.com/bmc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePrivateNetwork**](ServersAPI.md#DeletePrivateNetwork) | **Delete** /servers/{serverId}/network-configuration/private-network-configuration/private-networks/{privateNetworkId} | Removes the server from private network.
[**ServersGet**](ServersAPI.md#ServersGet) | **Get** /servers | List servers.
[**ServersPost**](ServersAPI.md#ServersPost) | **Post** /servers | Create new server.
[**ServersServerIdActionsDeprovisionPost**](ServersAPI.md#ServersServerIdActionsDeprovisionPost) | **Post** /servers/{serverId}/actions/deprovision | Deprovision a server.
[**ServersServerIdActionsPowerOffPost**](ServersAPI.md#ServersServerIdActionsPowerOffPost) | **Post** /servers/{serverId}/actions/power-off | Power off server.
[**ServersServerIdActionsPowerOnPost**](ServersAPI.md#ServersServerIdActionsPowerOnPost) | **Post** /servers/{serverId}/actions/power-on | Power on server.
[**ServersServerIdActionsProvisionPost**](ServersAPI.md#ServersServerIdActionsProvisionPost) | **Post** /servers/{serverId}/actions/provision | Provision server.
[**ServersServerIdActionsRebootPost**](ServersAPI.md#ServersServerIdActionsRebootPost) | **Post** /servers/{serverId}/actions/reboot | Reboot server.
[**ServersServerIdActionsReservePost**](ServersAPI.md#ServersServerIdActionsReservePost) | **Post** /servers/{serverId}/actions/reserve | Reserve server.
[**ServersServerIdActionsResetPost**](ServersAPI.md#ServersServerIdActionsResetPost) | **Post** /servers/{serverId}/actions/reset | Reset server.
[**ServersServerIdActionsShutdownPost**](ServersAPI.md#ServersServerIdActionsShutdownPost) | **Post** /servers/{serverId}/actions/shutdown | Shutdown server.
[**ServersServerIdDelete**](ServersAPI.md#ServersServerIdDelete) | **Delete** /servers/{serverId} | Delete server.
[**ServersServerIdGet**](ServersAPI.md#ServersServerIdGet) | **Get** /servers/{serverId} | Get server.
[**ServersServerIdIpBlocksIpBlockIdDelete**](ServersAPI.md#ServersServerIdIpBlocksIpBlockIdDelete) | **Delete** /servers/{serverId}/network-configuration/ip-block-configurations/ip-blocks/{ipBlockId} | Unassign IP Block from Server.
[**ServersServerIdIpBlocksPost**](ServersAPI.md#ServersServerIdIpBlocksPost) | **Post** /servers/{serverId}/network-configuration/ip-block-configurations/ip-blocks | Assign IP Block to Server.
[**ServersServerIdPatch**](ServersAPI.md#ServersServerIdPatch) | **Patch** /servers/{serverId} | Patch a Server.
[**ServersServerIdPrivateNetworksPatch**](ServersAPI.md#ServersServerIdPrivateNetworksPatch) | **Patch** /servers/{serverId}/network-configuration/private-network-configuration/private-networks/{privateNetworkId} | Updates the server&#39;s private network&#39;s IP addresses
[**ServersServerIdPrivateNetworksPost**](ServersAPI.md#ServersServerIdPrivateNetworksPost) | **Post** /servers/{serverId}/network-configuration/private-network-configuration/private-networks | Adds the server to a private network.
[**ServersServerIdPublicNetworksDelete**](ServersAPI.md#ServersServerIdPublicNetworksDelete) | **Delete** /servers/{serverId}/network-configuration/public-network-configuration/public-networks/{publicNetworkId} | Removes the server from the Public Network
[**ServersServerIdPublicNetworksPatch**](ServersAPI.md#ServersServerIdPublicNetworksPatch) | **Patch** /servers/{serverId}/network-configuration/public-network-configuration/public-networks/{publicNetworkId} | Updates the server&#39;s public network&#39;s IP addresses.
[**ServersServerIdPublicNetworksPost**](ServersAPI.md#ServersServerIdPublicNetworksPost) | **Post** /servers/{serverId}/network-configuration/public-network-configuration/public-networks | Adds the server to a Public Network.
[**ServersServerIdTagsPut**](ServersAPI.md#ServersServerIdTagsPut) | **Put** /servers/{serverId}/tags | Overwrite tags assigned for Server.



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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	privateNetworkId := "603f3b2cfcaf050643b89a4b" // string | The private network identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.DeletePrivateNetwork(context.Background(), serverId, privateNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.DeletePrivateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePrivateNetwork`: string
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.DeletePrivateNetwork`: %v\n", resp)
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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	tag := []string{"Inner_example"} // []string | A list of query parameters related to tags in the form of tagName.tagValue (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersGet(context.Background()).Tag(tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersGet`: []Server
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **[]string** | A list of query parameters related to tags in the form of tagName.tagValue | 

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

> Server ServersPost(ctx).ServerCreate(serverCreate).Force(force).Execute()

Create new server.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverCreate := *openapiclient.NewServerCreate("my-server-1", "ubuntu/bionic", "s1.c1.small", "PHX") // ServerCreate | 
	force := true // bool | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersPost(context.Background()).ServerCreate(serverCreate).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersPost`: Server
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverCreate** | [**ServerCreate**](ServerCreate.md) |  | 
 **force** | **bool** | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. | [default to false]

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


## ServersServerIdActionsDeprovisionPost

> string ServersServerIdActionsDeprovisionPost(ctx, serverId).RelinquishIpBlock(relinquishIpBlock).Execute()

Deprovision a server.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	relinquishIpBlock := *openapiclient.NewRelinquishIpBlock() // RelinquishIpBlock | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdActionsDeprovisionPost(context.Background(), serverId).RelinquishIpBlock(relinquishIpBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdActionsDeprovisionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdActionsDeprovisionPost`: string
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdActionsDeprovisionPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsDeprovisionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relinquishIpBlock** | [**RelinquishIpBlock**](RelinquishIpBlock.md) |  | 

### Return type

**string**

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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdActionsPowerOffPost(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdActionsPowerOffPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdActionsPowerOffPost`: ActionResult
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdActionsPowerOffPost`: %v\n", resp)
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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdActionsPowerOnPost(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdActionsPowerOnPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdActionsPowerOnPost`: ActionResult
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdActionsPowerOnPost`: %v\n", resp)
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


## ServersServerIdActionsProvisionPost

> Server ServersServerIdActionsProvisionPost(ctx, serverId).ServerProvision(serverProvision).Force(force).Execute()

Provision server.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	serverProvision := *openapiclient.NewServerProvision("my-server-1", "ubuntu/bionic") // ServerProvision | 
	force := true // bool | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdActionsProvisionPost(context.Background(), serverId).ServerProvision(serverProvision).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdActionsProvisionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdActionsProvisionPost`: Server
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdActionsProvisionPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdActionsProvisionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverProvision** | [**ServerProvision**](ServerProvision.md) |  | 
 **force** | **bool** | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. | [default to false]

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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdActionsRebootPost(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdActionsRebootPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdActionsRebootPost`: ActionResult
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdActionsRebootPost`: %v\n", resp)
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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	serverReserve := *openapiclient.NewServerReserve("ONE_MONTH_RESERVATION") // ServerReserve | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdActionsReservePost(context.Background(), serverId).ServerReserve(serverReserve).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdActionsReservePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdActionsReservePost`: Server
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdActionsReservePost`: %v\n", resp)
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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	serverReset := *openapiclient.NewServerReset() // ServerReset | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdActionsResetPost(context.Background(), serverId).ServerReset(serverReset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdActionsResetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdActionsResetPost`: ResetResult
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdActionsResetPost`: %v\n", resp)
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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdActionsShutdownPost(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdActionsShutdownPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdActionsShutdownPost`: ActionResult
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdActionsShutdownPost`: %v\n", resp)
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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdDelete(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdDelete`: DeleteResult
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdDelete`: %v\n", resp)
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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdGet(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdGet`: Server
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdGet`: %v\n", resp)
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


## ServersServerIdIpBlocksIpBlockIdDelete

> string ServersServerIdIpBlocksIpBlockIdDelete(ctx, serverId, ipBlockId).RelinquishIpBlock(relinquishIpBlock).Execute()

Unassign IP Block from Server.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	ipBlockId := "6047127fed34ecc3ba8402d2" // string | The IP Block identifier.
	relinquishIpBlock := *openapiclient.NewRelinquishIpBlock() // RelinquishIpBlock | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdIpBlocksIpBlockIdDelete(context.Background(), serverId, ipBlockId).RelinquishIpBlock(relinquishIpBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdIpBlocksIpBlockIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdIpBlocksIpBlockIdDelete`: string
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdIpBlocksIpBlockIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 
**ipBlockId** | **string** | The IP Block identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdIpBlocksIpBlockIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **relinquishIpBlock** | [**RelinquishIpBlock**](RelinquishIpBlock.md) |  | 

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdIpBlocksPost

> ServerIpBlock ServersServerIdIpBlocksPost(ctx, serverId).ServerIpBlock(serverIpBlock).Execute()

Assign IP Block to Server.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	serverIpBlock := *openapiclient.NewServerIpBlock("60473a6115e34466c9f8f083") // ServerIpBlock | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdIpBlocksPost(context.Background(), serverId).ServerIpBlock(serverIpBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdIpBlocksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdIpBlocksPost`: ServerIpBlock
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdIpBlocksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdIpBlocksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverIpBlock** | [**ServerIpBlock**](ServerIpBlock.md) |  | 

### Return type

[**ServerIpBlock**](ServerIpBlock.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	serverPatch := *openapiclient.NewServerPatch() // ServerPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdPatch(context.Background(), serverId).ServerPatch(serverPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdPatch`: Server
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdPatch`: %v\n", resp)
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


## ServersServerIdPrivateNetworksPatch

> ServerPrivateNetwork ServersServerIdPrivateNetworksPatch(ctx, serverId, privateNetworkId).ServerNetworkUpdate(serverNetworkUpdate).Force(force).Execute()

Updates the server's private network's IP addresses



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	privateNetworkId := "603f3b2cfcaf050643b89a4b" // string | The private network identifier.
	serverNetworkUpdate := *openapiclient.NewServerNetworkUpdate() // ServerNetworkUpdate | 
	force := true // bool | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdPrivateNetworksPatch(context.Background(), serverId, privateNetworkId).ServerNetworkUpdate(serverNetworkUpdate).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdPrivateNetworksPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdPrivateNetworksPatch`: ServerPrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdPrivateNetworksPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 
**privateNetworkId** | **string** | The private network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdPrivateNetworksPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serverNetworkUpdate** | [**ServerNetworkUpdate**](ServerNetworkUpdate.md) |  | 
 **force** | **bool** | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. | [default to false]

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


## ServersServerIdPrivateNetworksPost

> ServerPrivateNetwork ServersServerIdPrivateNetworksPost(ctx, serverId).ServerPrivateNetwork(serverPrivateNetwork).Force(force).Execute()

Adds the server to a private network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	serverPrivateNetwork := *openapiclient.NewServerPrivateNetwork("603f3b2cfcaf050643b89a4b") // ServerPrivateNetwork | 
	force := true // bool | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdPrivateNetworksPost(context.Background(), serverId).ServerPrivateNetwork(serverPrivateNetwork).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdPrivateNetworksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdPrivateNetworksPost`: ServerPrivateNetwork
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdPrivateNetworksPost`: %v\n", resp)
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
 **force** | **bool** | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. | [default to false]

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


## ServersServerIdPublicNetworksDelete

> string ServersServerIdPublicNetworksDelete(ctx, serverId, publicNetworkId).Execute()

Removes the server from the Public Network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdPublicNetworksDelete(context.Background(), serverId, publicNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdPublicNetworksDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdPublicNetworksDelete`: string
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdPublicNetworksDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 
**publicNetworkId** | **string** | The Public Network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdPublicNetworksDeleteRequest struct via the builder pattern


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


## ServersServerIdPublicNetworksPatch

> ServerPublicNetwork ServersServerIdPublicNetworksPatch(ctx, serverId, publicNetworkId).ServerNetworkUpdate(serverNetworkUpdate).Force(force).Execute()

Updates the server's public network's IP addresses.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.
	serverNetworkUpdate := *openapiclient.NewServerNetworkUpdate() // ServerNetworkUpdate | 
	force := true // bool | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdPublicNetworksPatch(context.Background(), serverId, publicNetworkId).ServerNetworkUpdate(serverNetworkUpdate).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdPublicNetworksPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdPublicNetworksPatch`: ServerPublicNetwork
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdPublicNetworksPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 
**publicNetworkId** | **string** | The Public Network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdPublicNetworksPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serverNetworkUpdate** | [**ServerNetworkUpdate**](ServerNetworkUpdate.md) |  | 
 **force** | **bool** | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. | [default to false]

### Return type

[**ServerPublicNetwork**](ServerPublicNetwork.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServersServerIdPublicNetworksPost

> ServerPublicNetwork ServersServerIdPublicNetworksPost(ctx, serverId).ServerPublicNetwork(serverPublicNetwork).Force(force).Execute()

Adds the server to a Public Network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	serverPublicNetwork := *openapiclient.NewServerPublicNetwork("60473c2509268bc77fd06d29") // ServerPublicNetwork | 
	force := true // bool | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdPublicNetworksPost(context.Background(), serverId).ServerPublicNetwork(serverPublicNetwork).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdPublicNetworksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdPublicNetworksPost`: ServerPublicNetwork
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdPublicNetworksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The server&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServersServerIdPublicNetworksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverPublicNetwork** | [**ServerPublicNetwork**](ServerPublicNetwork.md) |  | 
 **force** | **bool** | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. | [default to false]

### Return type

[**ServerPublicNetwork**](ServerPublicNetwork.md)

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

Overwrite tags assigned for Server.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func main() {
	serverId := "60473a6115e34466c9f8f083" // string | The server's ID.
	tagAssignmentRequest := []openapiclient.TagAssignmentRequest{*openapiclient.NewTagAssignmentRequest("Environment")} // []TagAssignmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ServersServerIdTagsPut(context.Background(), serverId).TagAssignmentRequest(tagAssignmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ServersServerIdTagsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServersServerIdTagsPut`: Server
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ServersServerIdTagsPut`: %v\n", resp)
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

