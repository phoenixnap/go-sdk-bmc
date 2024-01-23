/*
Locations API

Testing LocationsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package locationapi

import (
	"context"
	"testing"

	openapiclient "github.com/phoenixnap/go-sdk-bmc/locationapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_locationapi_LocationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LocationsAPIService GetLocations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.LocationsAPI.GetLocations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
