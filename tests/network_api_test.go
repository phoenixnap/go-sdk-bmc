package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	"github.com/stretchr/testify/suite"
)

type NetworkApiTestSuite struct {
	suite.Suite
	ctx           context.Context
	configuration *networkapi.Configuration
	apiClient     *networkapi.APIClient
}

// this function executes before each test
func (suite *NetworkApiTestSuite) SetupSuite() {
	// Set configuration
	suite.configuration = networkapi.NewConfiguration()
	suite.configuration.Host = "127.0.0.1:1080"
	suite.configuration.Scheme = "http"
	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)
	fmt.Println(">>> From SetupSuite")
	suite.apiClient = networkapi.NewAPIClient(suite.configuration)
}

// this function executes after each test
func (suite *NetworkApiTestSuite) TearDownTest() {
	fmt.Println(">>> From TearDownSuite")
	defer TestUtilsImpl{}.resetExpectations()
}

func (suite *NetworkApiTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *NetworkApiTestSuite) TestGetPrivateNetworks() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/private_networks_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	loc := fmt.Sprintf("%v", qpMap["location"])

	// Operation Execution
	result, _, _ := suite.apiClient.PrivateNetworksApi.PrivateNetworksGet(suite.ctx).Location(loc).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestCreatePrivateNetwork() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/private_networks_post", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	body, _ := json.Marshal(request.Body.Json)
	var networkCreate networkapi.PrivateNetworkCreate
	json.Unmarshal(body, &networkCreate)

	// Operation Execution
	result, _, _ := suite.apiClient.PrivateNetworksApi.PrivateNetworksPost(suite.ctx).PrivateNetworkCreate(networkCreate).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestGetPrivateNetworkById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/private_networks_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	networkId := TestUtilsImpl{}.extractIdFrom(request)

	// Operation Execution
	result, _, _ := suite.apiClient.PrivateNetworksApi.PrivateNetworksNetworkIdGet(suite.ctx, networkId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestPutPrivateNetworkById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/private_networks_put_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	body, _ := json.Marshal(request.Body.Json)
	var privateNetworkModify networkapi.PrivateNetworkModify
	json.Unmarshal(body, &privateNetworkModify)

	privateNetworkId := TestUtilsImpl{}.extractIdFrom(request)

	// Operation Execution
	result, _, _ := suite.apiClient.PrivateNetworksApi.PrivateNetworksNetworkIdPut(suite.ctx, privateNetworkId).PrivateNetworkModify(privateNetworkModify).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

// func (suite *NetworkApiTestSuite) TestDeletePrivateNetworkById() {
// 	// Generate payload
// 	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/networks_delete_by_id", "./payloads")

// 	// Extract the response expectation id
// 	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

// 	privateNetworkId := TestUtilsImpl{}.extractIdFrom(request)

// 	// Operation Execution
// 	result, _ := suite.apiClient.PrivateNetworksApi.PrivateNetworksNetworkIdDelete(suite.ctx, privateNetworkId).Execute()

// 	// Convert the result and response body to json strings
// 	jsonResult, _ := json.Marshal(result)
// 	jsonResponseBody, _ := json.Marshal(response.Body)

// 	suite.Equal(jsonResult, jsonResponseBody)

// 	// Verify
// 	suite.verifyCalledOnce(expectationId)
// }

func (suite *NetworkApiTestSuite) TestGetPublicNetworks() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/public_networks_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	loc := fmt.Sprintf("%v", qpMap["location"])

	// Operation Execution
	result, _, _ := suite.apiClient.PublicNetworksApi.PublicNetworksGet(suite.ctx).Location(loc).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestCreatePublicNetwork() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/public_networks_post", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	body, _ := json.Marshal(request.Body.Json)
	var networkCreate networkapi.PublicNetworkCreate
	json.Unmarshal(body, &networkCreate)

	// Operation Execution
	result, _, _ := suite.apiClient.PublicNetworksApi.PublicNetworksPost(suite.ctx).PublicNetworkCreate(networkCreate).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestGetPublicNetworkById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/public_networks_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	networkId := TestUtilsImpl{}.extractIdFrom(request)

	// Operation Execution
	result, _, _ := suite.apiClient.PublicNetworksApi.PublicNetworksNetworkIdGet(suite.ctx, networkId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestPatchPublicNetworkById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/public_networks_patch_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	body, _ := json.Marshal(request.Body.Json)
	var publicNetworkModify networkapi.PublicNetworkModify
	json.Unmarshal(body, &publicNetworkModify)

	publicNetworkId := TestUtilsImpl{}.extractIdFrom(request)

	// Operation Execution
	result, _, _ := suite.apiClient.PublicNetworksApi.PublicNetworksNetworkIdPatch(suite.ctx, publicNetworkId).PublicNetworkModify(publicNetworkModify).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

// func (suite *NetworkApiTestSuite) TestDeletePublicNetworkById() {

// }

func (suite *NetworkApiTestSuite) TestPostPublicNetworkIpBlockByPublicNetworkId() {

}

// func (suite *NetworkApiTestSuite) TestDeletePublicNetworkIpBlocksByPublicNetworkIdAndIpBlockID() {

// }

func TestNetworkApiTestSuite(t *testing.T) {
	suite.Run(t, new(NetworkApiTestSuite))
}
