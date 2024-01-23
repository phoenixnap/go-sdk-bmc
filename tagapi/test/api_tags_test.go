/*
Tags API

Testing TagsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tagapi

import (
	"context"
	"testing"

	openapiclient "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_tagapi_TagsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TagsAPIService TagsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TagsAPI.TagsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService TagsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TagsAPI.TagsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService TagsTagIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tagId string

		resp, httpRes, err := apiClient.TagsAPI.TagsTagIdDelete(context.Background(), tagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService TagsTagIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tagId string

		resp, httpRes, err := apiClient.TagsAPI.TagsTagIdGet(context.Background(), tagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService TagsTagIdPatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tagId string

		resp, httpRes, err := apiClient.TagsAPI.TagsTagIdPatch(context.Background(), tagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
