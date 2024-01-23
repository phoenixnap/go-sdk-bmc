/*
Network Storage API

Testing StorageNetworksAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package networkstorageapi

import (
	"context"
	"testing"

	openapiclient "github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_networkstorageapi_StorageNetworksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StorageNetworksAPIService StorageNetworksGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StorageNetworksAPI.StorageNetworksGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StorageNetworksAPIService StorageNetworksIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var storageId string

		httpRes, err := apiClient.StorageNetworksAPI.StorageNetworksIdDelete(context.Background(), storageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StorageNetworksAPIService StorageNetworksIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var storageId string

		resp, httpRes, err := apiClient.StorageNetworksAPI.StorageNetworksIdGet(context.Background(), storageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StorageNetworksAPIService StorageNetworksIdPatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var storageId string

		resp, httpRes, err := apiClient.StorageNetworksAPI.StorageNetworksIdPatch(context.Background(), storageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StorageNetworksAPIService StorageNetworksPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StorageNetworksAPI.StorageNetworksPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StorageNetworksAPIService StorageNetworksStorageNetworkIdVolumesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var storageId string

		resp, httpRes, err := apiClient.StorageNetworksAPI.StorageNetworksStorageNetworkIdVolumesGet(context.Background(), storageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StorageNetworksAPIService StorageNetworksStorageNetworkIdVolumesPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var storageId string

		resp, httpRes, err := apiClient.StorageNetworksAPI.StorageNetworksStorageNetworkIdVolumesPost(context.Background(), storageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StorageNetworksAPIService StorageNetworksStorageNetworkIdVolumesVolumeIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var storageId string
		var volumeId string

		httpRes, err := apiClient.StorageNetworksAPI.StorageNetworksStorageNetworkIdVolumesVolumeIdDelete(context.Background(), storageId, volumeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StorageNetworksAPIService StorageNetworksStorageNetworkIdVolumesVolumeIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var storageId string
		var volumeId string

		resp, httpRes, err := apiClient.StorageNetworksAPI.StorageNetworksStorageNetworkIdVolumesVolumeIdGet(context.Background(), storageId, volumeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StorageNetworksAPIService StorageNetworksStorageNetworkIdVolumesVolumeIdPatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var storageId string
		var volumeId string

		resp, httpRes, err := apiClient.StorageNetworksAPI.StorageNetworksStorageNetworkIdVolumesVolumeIdPatch(context.Background(), storageId, volumeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StorageNetworksAPIService StorageNetworksStorageNetworkIdVolumesVolumeIdTagsPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var storageId string
		var volumeId string

		resp, httpRes, err := apiClient.StorageNetworksAPI.StorageNetworksStorageNetworkIdVolumesVolumeIdTagsPut(context.Background(), storageId, volumeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
