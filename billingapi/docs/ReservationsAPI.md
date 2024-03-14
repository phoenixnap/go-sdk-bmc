# \ReservationsAPI

All URIs are relative to *https://api.phoenixnap.com/billing/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReservationsGet**](ReservationsAPI.md#ReservationsGet) | **Get** /reservations | List all Reservations.
[**ReservationsPost**](ReservationsAPI.md#ReservationsPost) | **Post** /reservations | Create a reservation.
[**ReservationsReservationIdActionsAutoRenewDisablePost**](ReservationsAPI.md#ReservationsReservationIdActionsAutoRenewDisablePost) | **Post** /reservations/{id}/actions/auto-renew/disable | Disable auto-renewal for reservation by id.
[**ReservationsReservationIdActionsAutoRenewEnablePost**](ReservationsAPI.md#ReservationsReservationIdActionsAutoRenewEnablePost) | **Post** /reservations/{id}/actions/auto-renew/enable | Enable auto-renewal for unexpired reservation by reservation id.
[**ReservationsReservationIdActionsConvertPost**](ReservationsAPI.md#ReservationsReservationIdActionsConvertPost) | **Post** /reservations/{id}/actions/convert | Convert reservation pricing model by reservation ID.
[**ReservationsReservationIdGet**](ReservationsAPI.md#ReservationsReservationIdGet) | **Get** /reservations/{id} | Get a reservation.



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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi"
)

func main() {
	productCategory := openapiclient.ReservationProductCategoryEnum("server") // ReservationProductCategoryEnum | The product category (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReservationsAPI.ReservationsGet(context.Background()).ProductCategory(productCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReservationsAPI.ReservationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReservationsGet`: []Reservation
	fmt.Fprintf(os.Stdout, "Response from `ReservationsAPI.ReservationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReservationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCategory** | [**ReservationProductCategoryEnum**](ReservationProductCategoryEnum.md) | The product category | 

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
	openapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi"
)

func main() {
	reservationRequest := *openapiclient.NewReservationRequest("XXX-XXX-XXX") // ReservationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReservationsAPI.ReservationsPost(context.Background()).ReservationRequest(reservationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReservationsAPI.ReservationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReservationsPost`: Reservation
	fmt.Fprintf(os.Stdout, "Response from `ReservationsAPI.ReservationsPost`: %v\n", resp)
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

> Reservation ReservationsReservationIdActionsAutoRenewDisablePost(ctx, id).ReservationAutoRenewDisableRequest(reservationAutoRenewDisableRequest).Execute()

Disable auto-renewal for reservation by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi"
)

func main() {
	id := "d90bbea9-5725-47ce-879e-d3905bafac2a" // string | Resource id.
	reservationAutoRenewDisableRequest := *openapiclient.NewReservationAutoRenewDisableRequest() // ReservationAutoRenewDisableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReservationsAPI.ReservationsReservationIdActionsAutoRenewDisablePost(context.Background(), id).ReservationAutoRenewDisableRequest(reservationAutoRenewDisableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReservationsAPI.ReservationsReservationIdActionsAutoRenewDisablePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReservationsReservationIdActionsAutoRenewDisablePost`: Reservation
	fmt.Fprintf(os.Stdout, "Response from `ReservationsAPI.ReservationsReservationIdActionsAutoRenewDisablePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource id. | 

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

> Reservation ReservationsReservationIdActionsAutoRenewEnablePost(ctx, id).Execute()

Enable auto-renewal for unexpired reservation by reservation id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi"
)

func main() {
	id := "d90bbea9-5725-47ce-879e-d3905bafac2a" // string | Resource id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReservationsAPI.ReservationsReservationIdActionsAutoRenewEnablePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReservationsAPI.ReservationsReservationIdActionsAutoRenewEnablePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReservationsReservationIdActionsAutoRenewEnablePost`: Reservation
	fmt.Fprintf(os.Stdout, "Response from `ReservationsAPI.ReservationsReservationIdActionsAutoRenewEnablePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource id. | 

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

> Reservation ReservationsReservationIdActionsConvertPost(ctx, id).ReservationRequest(reservationRequest).Execute()

Convert reservation pricing model by reservation ID.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi"
)

func main() {
	id := "d90bbea9-5725-47ce-879e-d3905bafac2a" // string | Resource id.
	reservationRequest := *openapiclient.NewReservationRequest("XXX-XXX-XXX") // ReservationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReservationsAPI.ReservationsReservationIdActionsConvertPost(context.Background(), id).ReservationRequest(reservationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReservationsAPI.ReservationsReservationIdActionsConvertPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReservationsReservationIdActionsConvertPost`: Reservation
	fmt.Fprintf(os.Stdout, "Response from `ReservationsAPI.ReservationsReservationIdActionsConvertPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource id. | 

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

> Reservation ReservationsReservationIdGet(ctx, id).Execute()

Get a reservation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi"
)

func main() {
	id := "d90bbea9-5725-47ce-879e-d3905bafac2a" // string | Resource id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReservationsAPI.ReservationsReservationIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReservationsAPI.ReservationsReservationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReservationsReservationIdGet`: Reservation
	fmt.Fprintf(os.Stdout, "Response from `ReservationsAPI.ReservationsReservationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource id. | 

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

