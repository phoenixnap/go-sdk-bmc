# \BGPPeerGroupsAPI

All URIs are relative to *https://api.phoenixnap.com/networks/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BgpPeerGroupsGet**](BGPPeerGroupsAPI.md#BgpPeerGroupsGet) | **Get** /bgp-peer-groups | List BGP Peer Groups.
[**BgpPeerGroupsPeerGroupIdDelete**](BGPPeerGroupsAPI.md#BgpPeerGroupsPeerGroupIdDelete) | **Delete** /bgp-peer-groups/{bgpPeerGroupId} | Delete a BGP Peer Group.
[**BgpPeerGroupsPeerGroupIdGet**](BGPPeerGroupsAPI.md#BgpPeerGroupsPeerGroupIdGet) | **Get** /bgp-peer-groups/{bgpPeerGroupId} | Get a BGP Peer Group.
[**BgpPeerGroupsPeerGroupIdPatch**](BGPPeerGroupsAPI.md#BgpPeerGroupsPeerGroupIdPatch) | **Patch** /bgp-peer-groups/{bgpPeerGroupId} | Modify a BGP Peer Group.
[**BgpPeerGroupsPost**](BGPPeerGroupsAPI.md#BgpPeerGroupsPost) | **Post** /bgp-peer-groups | Create a BGP Peer Group.



## BgpPeerGroupsGet

> []BgpPeerGroup BgpPeerGroupsGet(ctx).Location(location).Execute()

List BGP Peer Groups.



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
	location := "PHX" // string | If present will filter the result by the given location of the BGP Peer Group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BGPPeerGroupsAPI.BgpPeerGroupsGet(context.Background()).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPPeerGroupsAPI.BgpPeerGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BgpPeerGroupsGet`: []BgpPeerGroup
	fmt.Fprintf(os.Stdout, "Response from `BGPPeerGroupsAPI.BgpPeerGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBgpPeerGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **string** | If present will filter the result by the given location of the BGP Peer Group. | 

### Return type

[**[]BgpPeerGroup**](BgpPeerGroup.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BgpPeerGroupsPeerGroupIdDelete

> BgpPeerGroup BgpPeerGroupsPeerGroupIdDelete(ctx, bgpPeerGroupId).Execute()

Delete a BGP Peer Group.



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
	bgpPeerGroupId := "603f3b2cfcaf050643b89a4b" // string | The BGP peer group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdDelete(context.Background(), bgpPeerGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BgpPeerGroupsPeerGroupIdDelete`: BgpPeerGroup
	fmt.Fprintf(os.Stdout, "Response from `BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bgpPeerGroupId** | **string** | The BGP peer group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBgpPeerGroupsPeerGroupIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BgpPeerGroup**](BgpPeerGroup.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BgpPeerGroupsPeerGroupIdGet

> BgpPeerGroup BgpPeerGroupsPeerGroupIdGet(ctx, bgpPeerGroupId).Execute()

Get a BGP Peer Group.



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
	bgpPeerGroupId := "603f3b2cfcaf050643b89a4b" // string | The BGP peer group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdGet(context.Background(), bgpPeerGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BgpPeerGroupsPeerGroupIdGet`: BgpPeerGroup
	fmt.Fprintf(os.Stdout, "Response from `BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bgpPeerGroupId** | **string** | The BGP peer group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBgpPeerGroupsPeerGroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BgpPeerGroup**](BgpPeerGroup.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BgpPeerGroupsPeerGroupIdPatch

> BgpPeerGroup BgpPeerGroupsPeerGroupIdPatch(ctx, bgpPeerGroupId).BgpPeerGroupPatch(bgpPeerGroupPatch).Execute()

Modify a BGP Peer Group.



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
	bgpPeerGroupId := "603f3b2cfcaf050643b89a4b" // string | The BGP peer group ID.
	bgpPeerGroupPatch := *openapiclient.NewBgpPeerGroupPatch() // BgpPeerGroupPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdPatch(context.Background(), bgpPeerGroupId).BgpPeerGroupPatch(bgpPeerGroupPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BgpPeerGroupsPeerGroupIdPatch`: BgpPeerGroup
	fmt.Fprintf(os.Stdout, "Response from `BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bgpPeerGroupId** | **string** | The BGP peer group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBgpPeerGroupsPeerGroupIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bgpPeerGroupPatch** | [**BgpPeerGroupPatch**](BgpPeerGroupPatch.md) |  | 

### Return type

[**BgpPeerGroup**](BgpPeerGroup.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BgpPeerGroupsPost

> BgpPeerGroup BgpPeerGroupsPost(ctx).BgpPeerGroupCreate(bgpPeerGroupCreate).Execute()

Create a BGP Peer Group.



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
	bgpPeerGroupCreate := *openapiclient.NewBgpPeerGroupCreate("ASH", int64(65401), "DEFAULT") // BgpPeerGroupCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BGPPeerGroupsAPI.BgpPeerGroupsPost(context.Background()).BgpPeerGroupCreate(bgpPeerGroupCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPPeerGroupsAPI.BgpPeerGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BgpPeerGroupsPost`: BgpPeerGroup
	fmt.Fprintf(os.Stdout, "Response from `BGPPeerGroupsAPI.BgpPeerGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBgpPeerGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bgpPeerGroupCreate** | [**BgpPeerGroupCreate**](BgpPeerGroupCreate.md) |  | 

### Return type

[**BgpPeerGroup**](BgpPeerGroup.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

