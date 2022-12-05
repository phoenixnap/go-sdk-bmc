package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/stretchr/testify/suite"
)

type NetworkStorageApiTestSuite struct {
	suite.Suite
	ctx       context.Context
	apiClient *networkstorageapi.APIClient
}

func TestNetworkStorageApiTestSuite(t *testing.T) {
	suite.Run(t, new(NetworkStorageApiTestSuite))
}

func (suite *NetworkStorageApiTestSuite) SetupSuite() {
	// Remove any existing expectations
	TestUtilsImpl{}.resetExpectations()

	// Set configuration
	configuration := networkstorageapi.NewConfiguration()
	configuration.Host = "127.0.0.1:1080"
	configuration.Scheme = "http"

	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)

	// New ApiClient
	suite.apiClient = networkstorageapi.NewAPIClient(configuration)
}

func (suite *NetworkStorageApiTestSuite) TearDownTest() {
	TestUtilsImpl{}.resetExpectations()
}

func (suite *NetworkStorageApiTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *NetworkStorageApiTestSuite) TestGetStorageNetworks() {

	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkstorageapi/networkstorage_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	location := fmt.Sprintf("%v", qpMap["location"])

	// Operation Execution
	result, _, _ := suite.apiClient.StorageNetworksApi.StorageNetworksGet(suite.ctx).Location(location).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkStorageApiTestSuite) TestGetStorageNetworkById() {

	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkstorageapi/networkstorage_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	storageNetworkId := request.PathParameters["storageNetworkId"][0]

	// Operation Execution
	result, _, _ := suite.apiClient.StorageNetworksApi.StorageNetworksIdGet(suite.ctx, storageNetworkId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkStorageApiTestSuite) TestPostStorageNetworks() {

	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkstorageapi/networkstorage_post", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	body, _ := json.Marshal(request.Body.Json)
	var storageNetworkCreate networkstorageapi.StorageNetworkCreate
	json.Unmarshal(body, &storageNetworkCreate)

	// Operation Execution
	result, _, _ := suite.apiClient.StorageNetworksApi.StorageNetworksPost(suite.ctx).StorageNetworkCreate(storageNetworkCreate).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkStorageApiTestSuite) TestPatchStorageNetworkById() {

	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkstorageapi/networkstorage_patch_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	body, _ := json.Marshal(request.Body.Json)
	var storageNetworkUpdate networkstorageapi.StorageNetworkUpdate
	json.Unmarshal(body, &storageNetworkUpdate)

	storageNetworkId := request.PathParameters["storageNetworkId"][0]

	// Operation Execution
	result, _, _ := suite.apiClient.StorageNetworksApi.StorageNetworksIdPatch(suite.ctx, storageNetworkId).StorageNetworkUpdate(storageNetworkUpdate).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkStorageApiTestSuite) TestDeleteStorageNetworkById() {

	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkstorageapi/networkstorage_delete_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	storageNetworkId := request.PathParameters["storageNetworkId"][0]

	// Operation Execution
	result, _ := suite.apiClient.StorageNetworksApi.StorageNetworksIdDelete(suite.ctx, storageNetworkId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)

	// Asserts
	suite.Empty(string(jsonResult))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkStorageApiTestSuite) TestGetVolumesByStorageNetworkId() {

	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkstorageapi/networkstorage_get_volumes", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	storageNetworkId := request.PathParameters["storageNetworkId"][0]

	// Operation Execution
	result, _, _ := suite.apiClient.StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesGet(suite.ctx, storageNetworkId).Execute()

	// fmt.Print(resp.StatusCode)
	// panic(error)

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkStorageApiTestSuite) TestGetVolumeByStorageNetworkIdAndVolumeId() {

	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkstorageapi/networkstorage_get_volume_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	storageNetworkId := request.PathParameters["storageNetworkId"][0]
	volumeId := request.PathParameters["volumeId"][0]

	// Operation Execution
	result, _, _ := suite.apiClient.StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdGet(suite.ctx, storageNetworkId, volumeId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkStorageApiTestSuite) TestPatchVolumeByStorageNetworkIdAndVolumeId() {

	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkstorageapi/networkstorage_patch_volume_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	storageNetworkId := request.PathParameters["storageNetworkId"][0]
	volumeId := request.PathParameters["volumeId"][0]

	body, _ := json.Marshal(request.Body.Json)
	var volumeUpdate networkstorageapi.VolumeUpdate
	json.Unmarshal(body, &volumeUpdate)

	// Operation Execution
	result, _, _ := suite.apiClient.StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdPatch(suite.ctx, storageNetworkId, volumeId).VolumeUpdate(volumeUpdate).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}
