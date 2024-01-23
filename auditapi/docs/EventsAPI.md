# \EventsAPI

All URIs are relative to *https://api.phoenixnap.com/audit/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsGet**](EventsAPI.md#EventsGet) | **Get** /events | List event logs.



## EventsGet

> []Event EventsGet(ctx).From(from).To(to).Limit(limit).Order(order).Username(username).Verb(verb).Uri(uri).Execute()

List event logs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/phoenixnap/go-sdk-bmc/auditapi"
)

func main() {
	from := time.Now() // time.Time | From the date and time (inclusive) to filter event log records by. (optional)
	to := time.Now() // time.Time | To the date and time (inclusive) to filter event log records by. (optional)
	limit := int32(10) // int32 | Limit the number of records returned. (optional)
	order := "order_example" // string | Ordering of the event's time. SortBy can be introduced later on. (optional) (default to "ASC")
	username := "johnd@phoenixnap.com" // string | The username that did the actions. (optional)
	verb := "verb_example" // string | The HTTP verb corresponding to the action. (optional)
	uri := "/ams/v1/clients/12345" // string | The request uri. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsGet(context.Background()).From(from).To(to).Limit(limit).Order(order).Username(username).Verb(verb).Uri(uri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsGet`: []Event
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** | From the date and time (inclusive) to filter event log records by. | 
 **to** | **time.Time** | To the date and time (inclusive) to filter event log records by. | 
 **limit** | **int32** | Limit the number of records returned. | 
 **order** | **string** | Ordering of the event&#39;s time. SortBy can be introduced later on. | [default to &quot;ASC&quot;]
 **username** | **string** | The username that did the actions. | 
 **verb** | **string** | The HTTP verb corresponding to the action. | 
 **uri** | **string** | The request uri. | 

### Return type

[**[]Event**](Event.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

