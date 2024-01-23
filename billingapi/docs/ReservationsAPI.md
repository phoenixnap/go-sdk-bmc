# \ReservationsApi

All URIs are relative to *https://api.phoenixnap.com/billing/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReservationsGet**](ReservationsApi.md#ReservationsGet) | **Get** /reservations | List all Reservations.
[**ReservationsPost**](ReservationsApi.md#ReservationsPost) | **Post** /reservations | Create a reservation.
[**ReservationsReservationIdActionsAutoRenewDisablePost**](ReservationsApi.md#ReservationsReservationIdActionsAutoRenewDisablePost) | **Post** /reservations/{reservationId}/actions/auto-renew/disable | Disable auto-renewal for reservation by id.
[**ReservationsReservationIdActionsAutoRenewEnablePost**](ReservationsApi.md#ReservationsReservationIdActionsAutoRenewEnablePost) | **Post** /reservations/{reservationId}/actions/auto-renew/enable | Enable auto-renewal for unexpired reservation by reservation id.
[**ReservationsReservationIdActionsConvertPost**](ReservationsApi.md#ReservationsReservationIdActionsConvertPost) | **Post** /reservations/{reservationId}/actions/convert | Convert reservation pricing model by reservation ID.
[**ReservationsReservationIdGet**](ReservationsApi.md#ReservationsReservationIdGet) | **Get** /reservations/{reservationId} | Get a reservation.



## ReservationsGet

> []Reservation ReservationsGet(ctx).ProductCategory(productCategory).Execute()

List all Reservations.



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
    productCategory := openapiclient.ProductCategoryEnum("SERVER") // ProductCategoryEnum | The product category (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsGet(context.Background()).ProductCategory(productCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsGet`: []Reservation
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReservationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCategory** | [**ProductCategoryEnum**](ProductCategoryEnum.md) | The product category | 

### Return type

[**[]Reservation**](Reservation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsPost

> Reservation ReservationsPost(ctx).ReservationRequest(reservationRequest).Execute()

Create a reservation.



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
    reservationRequest := *openapiclient.NewReservationRequest("XXX-XXX-XXX") // ReservationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsPost(context.Background()).ReservationRequest(reservationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsPost`: Reservation
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReservationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reservationRequest** | [**ReservationRequest**](ReservationRequest.md) |  | 

### Return type

[**Reservation**](Reservation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsReservationIdActionsAutoRenewDisablePost

> Reservation ReservationsReservationIdActionsAutoRenewDisablePost(ctx, reservationId).ReservationAutoRenewDisableRequest(reservationAutoRenewDisableRequest).Execute()

Disable auto-renewal for reservation by id.



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
    reservationId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The reservation's ID.
    reservationAutoRenewDisableRequest := *openapiclient.NewReservationAutoRenewDisableRequest() // ReservationAutoRenewDisableRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsReservationIdActionsAutoRenewDisablePost(context.Background(), reservationId).ReservationAutoRenewDisableRequest(reservationAutoRenewDisableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsReservationIdActionsAutoRenewDisablePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsReservationIdActionsAutoRenewDisablePost`: Reservation
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsReservationIdActionsAutoRenewDisablePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string** | The reservation&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReservationsReservationIdActionsAutoRenewDisablePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reservationAutoRenewDisableRequest** | [**ReservationAutoRenewDisableRequest**](ReservationAutoRenewDisableRequest.md) |  | 

### Return type

[**Reservation**](Reservation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsReservationIdActionsAutoRenewEnablePost

> Reservation ReservationsReservationIdActionsAutoRenewEnablePost(ctx, reservationId).Execute()

Enable auto-renewal for unexpired reservation by reservation id.



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
    reservationId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The reservation's ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsReservationIdActionsAutoRenewEnablePost(context.Background(), reservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsReservationIdActionsAutoRenewEnablePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsReservationIdActionsAutoRenewEnablePost`: Reservation
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsReservationIdActionsAutoRenewEnablePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string** | The reservation&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReservationsReservationIdActionsAutoRenewEnablePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Reservation**](Reservation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsReservationIdActionsConvertPost

> Reservation ReservationsReservationIdActionsConvertPost(ctx, reservationId).ReservationRequest(reservationRequest).Execute()

Convert reservation pricing model by reservation ID.



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
    reservationId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The reservation's ID.
    reservationRequest := *openapiclient.NewReservationRequest("XXX-XXX-XXX") // ReservationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsReservationIdActionsConvertPost(context.Background(), reservationId).ReservationRequest(reservationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsReservationIdActionsConvertPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsReservationIdActionsConvertPost`: Reservation
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsReservationIdActionsConvertPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string** | The reservation&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReservationsReservationIdActionsConvertPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reservationRequest** | [**ReservationRequest**](ReservationRequest.md) |  | 

### Return type

[**Reservation**](Reservation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsReservationIdGet

> Reservation ReservationsReservationIdGet(ctx, reservationId).Execute()

Get a reservation.



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
    reservationId := "e6afba51-7de8-4080-83ab-0f915570659c" // string | The reservation's ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsReservationIdGet(context.Background(), reservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsReservationIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsReservationIdGet`: Reservation
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsReservationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationId** | **string** | The reservation&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReservationsReservationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Reservation**](Reservation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

