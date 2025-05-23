# \PublicNetworksAPI

All URIs are relative to *https://api.phoenixnap.com/networks/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicNetworksGet**](PublicNetworksAPI.md#PublicNetworksGet) | **Get** /public-networks | List Public Networks.
[**PublicNetworksNetworkIdDelete**](PublicNetworksAPI.md#PublicNetworksNetworkIdDelete) | **Delete** /public-networks/{publicNetworkId} | Delete a Public Network.
[**PublicNetworksNetworkIdGet**](PublicNetworksAPI.md#PublicNetworksNetworkIdGet) | **Get** /public-networks/{publicNetworkId} | Get a Public Network.
[**PublicNetworksNetworkIdIpBlocksIpBlockIdDelete**](PublicNetworksAPI.md#PublicNetworksNetworkIdIpBlocksIpBlockIdDelete) | **Delete** /public-networks/{publicNetworkId}/ip-blocks/{ipBlockId} | Removes the IP Block from the Public Network.
[**PublicNetworksNetworkIdIpBlocksPost**](PublicNetworksAPI.md#PublicNetworksNetworkIdIpBlocksPost) | **Post** /public-networks/{publicNetworkId}/ip-blocks | Adds an IP block to this public network.
[**PublicNetworksNetworkIdPatch**](PublicNetworksAPI.md#PublicNetworksNetworkIdPatch) | **Patch** /public-networks/{publicNetworkId} | Update Public Network&#39;s Details.
[**PublicNetworksPost**](PublicNetworksAPI.md#PublicNetworksPost) | **Post** /public-networks | Create a public network.



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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func main() {
	location := "PHX" // string | If present will filter the result by the given location of the Public Networks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicNetworksAPI.PublicNetworksGet(context.Background()).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksAPI.PublicNetworksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicNetworksGet`: []PublicNetwork
	fmt.Fprintf(os.Stdout, "Response from `PublicNetworksAPI.PublicNetworksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func main() {
	publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicNetworksAPI.PublicNetworksNetworkIdDelete(context.Background(), publicNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksAPI.PublicNetworksNetworkIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicNetworkId** | **string** | The Public Network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksNetworkIdDeleteRequest struct via the builder pattern


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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func main() {
	publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicNetworksAPI.PublicNetworksNetworkIdGet(context.Background(), publicNetworkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksAPI.PublicNetworksNetworkIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicNetworksNetworkIdGet`: PublicNetwork
	fmt.Fprintf(os.Stdout, "Response from `PublicNetworksAPI.PublicNetworksNetworkIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicNetworkId** | **string** | The Public Network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksNetworkIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> string PublicNetworksNetworkIdIpBlocksIpBlockIdDelete(ctx, publicNetworkId, ipBlockId).Force(force).Execute()

Removes the IP Block from the Public Network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func main() {
	publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.
	ipBlockId := "6047127fed34ecc3ba8402d2" // string | The IP Block identifier.
	force := true // bool | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicNetworksAPI.PublicNetworksNetworkIdIpBlocksIpBlockIdDelete(context.Background(), publicNetworkId, ipBlockId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksAPI.PublicNetworksNetworkIdIpBlocksIpBlockIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicNetworksNetworkIdIpBlocksIpBlockIdDelete`: string
	fmt.Fprintf(os.Stdout, "Response from `PublicNetworksAPI.PublicNetworksNetworkIdIpBlocksIpBlockIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicNetworkId** | **string** | The Public Network identifier. | 
**ipBlockId** | **string** | The IP Block identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksNetworkIdIpBlocksIpBlockIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. | [default to false]

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

> PublicNetworkIpBlock PublicNetworksNetworkIdIpBlocksPost(ctx, publicNetworkId).PublicNetworkIpBlockCreate(publicNetworkIpBlockCreate).Execute()

Adds an IP block to this public network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func main() {
	publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.
	publicNetworkIpBlockCreate := *openapiclient.NewPublicNetworkIpBlockCreate("60473a6115e34466c9f8f083") // PublicNetworkIpBlockCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicNetworksAPI.PublicNetworksNetworkIdIpBlocksPost(context.Background(), publicNetworkId).PublicNetworkIpBlockCreate(publicNetworkIpBlockCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksAPI.PublicNetworksNetworkIdIpBlocksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicNetworksNetworkIdIpBlocksPost`: PublicNetworkIpBlock
	fmt.Fprintf(os.Stdout, "Response from `PublicNetworksAPI.PublicNetworksNetworkIdIpBlocksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicNetworkId** | **string** | The Public Network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksNetworkIdIpBlocksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicNetworkIpBlockCreate** | [**PublicNetworkIpBlockCreate**](PublicNetworkIpBlockCreate.md) |  | 

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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func main() {
	publicNetworkId := "603f3b2cfcaf050643b89a4b" // string | The Public Network identifier.
	publicNetworkModify := *openapiclient.NewPublicNetworkModify() // PublicNetworkModify | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicNetworksAPI.PublicNetworksNetworkIdPatch(context.Background(), publicNetworkId).PublicNetworkModify(publicNetworkModify).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksAPI.PublicNetworksNetworkIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicNetworksNetworkIdPatch`: PublicNetwork
	fmt.Fprintf(os.Stdout, "Response from `PublicNetworksAPI.PublicNetworksNetworkIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicNetworkId** | **string** | The Public Network identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksNetworkIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func main() {
	publicNetworkCreate := *openapiclient.NewPublicNetworkCreate("Sample Network", "PHX") // PublicNetworkCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicNetworksAPI.PublicNetworksPost(context.Background()).PublicNetworkCreate(publicNetworkCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicNetworksAPI.PublicNetworksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicNetworksPost`: PublicNetwork
	fmt.Fprintf(os.Stdout, "Response from `PublicNetworksAPI.PublicNetworksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicNetworksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

