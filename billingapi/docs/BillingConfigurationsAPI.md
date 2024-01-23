# \BillingConfigurationsAPI

All URIs are relative to *https://api.phoenixnap.com/billing/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountBillingConfigurationMeGet**](BillingConfigurationsAPI.md#AccountBillingConfigurationMeGet) | **Get** /account-billing-configurations/me | Retrieves billing configuration associated with the authenticated account.



## AccountBillingConfigurationMeGet

> ConfigurationDetails AccountBillingConfigurationMeGet(ctx).Execute()

Retrieves billing configuration associated with the authenticated account.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingConfigurationsAPI.AccountBillingConfigurationMeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingConfigurationsAPI.AccountBillingConfigurationMeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountBillingConfigurationMeGet`: ConfigurationDetails
	fmt.Fprintf(os.Stdout, "Response from `BillingConfigurationsAPI.AccountBillingConfigurationMeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountBillingConfigurationMeGetRequest struct via the builder pattern


### Return type

[**ConfigurationDetails**](ConfigurationDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

