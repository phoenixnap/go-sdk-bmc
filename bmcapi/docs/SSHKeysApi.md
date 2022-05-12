# \SSHKeysApi

All URIs are relative to *https://api.phoenixnap.com/bmc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SshKeysGet**](SSHKeysApi.md#SshKeysGet) | **Get** /ssh-keys | List SSH Keys.
[**SshKeysPost**](SSHKeysApi.md#SshKeysPost) | **Post** /ssh-keys | Create SSH Key.
[**SshKeysSshKeyIdDelete**](SSHKeysApi.md#SshKeysSshKeyIdDelete) | **Delete** /ssh-keys/{sshKeyId} | Delete SSH Key.
[**SshKeysSshKeyIdGet**](SSHKeysApi.md#SshKeysSshKeyIdGet) | **Get** /ssh-keys/{sshKeyId} | Get SSH Key.
[**SshKeysSshKeyIdPut**](SSHKeysApi.md#SshKeysSshKeyIdPut) | **Put** /ssh-keys/{sshKeyId} | Edit SSH Key.



## SshKeysGet

> []SshKey SshKeysGet(ctx).Execute()

List SSH Keys.



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
    resp, r, err := apiClient.SSHKeysApi.SshKeysGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysApi.SshKeysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshKeysGet`: []SshKey
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysApi.SshKeysGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysGetRequest struct via the builder pattern


### Return type

[**[]SshKey**](SshKey.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeysPost

> SshKey SshKeysPost(ctx).SshKeyCreate(sshKeyCreate).Execute()

Create SSH Key.



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
    sshKeyCreate := *openapiclient.NewSshKeyCreate(true, "sshkey-name-01", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDF9LdAFElNCi7JoWh6KUcchrJ2Gac1aqGRPpdZNowObpRtmiRCecAMb7bUgNAaNfcmwiQi7tos9TlnFgprIcfMWb8MSs3ABYHmBgqEEt3RWYf0fAc9CsIpJdMCUG28TPGTlRXCEUVNKgLMdcseAlJoGp1CgbHWIN65fB3he3kAZcfpPn5mapV0tsl2p+ZyuAGRYdn5dJv2RZDHUZBkOeUobwsij+weHCKAFmKQKtCP7ybgVHaQjAPrj8MGnk1jBbjDt5ws+Be+9JNjQJee9zCKbAOsIo3i+GcUIkrw5jxPU/RTGlWBcemPaKHdciSzGcjWboapzIy49qypQhZe1U75 user@my_ip") // SshKeyCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysApi.SshKeysPost(context.Background()).SshKeyCreate(sshKeyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysApi.SshKeysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshKeysPost`: SshKey
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysApi.SshKeysPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysPostRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----  
**sshKeyCreate** | [**SshKeyCreate**](SshKeyCreate.md) |  | 

### Return type

[**SshKey**](SshKey.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeysSshKeyIdDelete

> DeleteSshKeyResult SshKeysSshKeyIdDelete(ctx, sshKeyId).Execute()

Delete SSH Key.



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
    sshKeyId := "5fa54d1e91867c03a0a7b4a4" // string | The SSH Key's ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysApi.SshKeysSshKeyIdDelete(context.Background(), sshKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysApi.SshKeysSshKeyIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshKeysSshKeyIdDelete`: DeleteSshKeyResult
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysApi.SshKeysSshKeyIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**sshKeyId** | **string** | The SSH Key&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysSshKeyIdDeleteRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----  

### Return type

[**DeleteSshKeyResult**](DeleteSshKeyResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeysSshKeyIdGet

> SshKey SshKeysSshKeyIdGet(ctx, sshKeyId).Execute()

Get SSH Key.



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
    sshKeyId := "5fa54d1e91867c03a0a7b4a4" // string | The SSH Key's ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysApi.SshKeysSshKeyIdGet(context.Background(), sshKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysApi.SshKeysSshKeyIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshKeysSshKeyIdGet`: SshKey
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysApi.SshKeysSshKeyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**sshKeyId** | **string** | The SSH Key&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysSshKeyIdGetRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----  

### Return type

[**SshKey**](SshKey.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeysSshKeyIdPut

> SshKey SshKeysSshKeyIdPut(ctx, sshKeyId).SshKeyUpdate(sshKeyUpdate).Execute()

Edit SSH Key.



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
    sshKeyId := "5fa54d1e91867c03a0a7b4a4" // string | The SSH Key's ID.
    sshKeyUpdate := *openapiclient.NewSshKeyUpdate(true, "sshkey-name-01") // SshKeyUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysApi.SshKeysSshKeyIdPut(context.Background(), sshKeyId).SshKeyUpdate(sshKeyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysApi.SshKeysSshKeyIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshKeysSshKeyIdPut`: SshKey
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysApi.SshKeysSshKeyIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description | Notes
---- | ---- | ----------- | -----
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
**sshKeyId** | **string** | The SSH Key&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeysSshKeyIdPutRequest struct via the builder pattern

Name | Type | Description | Notes
---- | ---- | ----------- | -----  
 
**sshKeyUpdate** | [**SshKeyUpdate**](SshKeyUpdate.md) |  | 

### Return type

[**SshKey**](SshKey.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

