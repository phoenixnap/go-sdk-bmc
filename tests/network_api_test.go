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
	// Remove any existing expectations
	TestUtilsImpl{}.resetExpectations()

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
	result, _, _ := suite.apiClient.PrivateNetworksAPI.PrivateNetworksGet(suite.ctx).Location(loc).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestCreatePrivateNetwork() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/private_networks_post", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)
	force := fmt.Sprintf("%v", qpMap["force"]) == "true"

	body, _ := json.Marshal(request.Body.Json)
	var networkCreate networkapi.PrivateNetworkCreate
	json.Unmarshal(body, &networkCreate)

	// Operation Execution
	result, _, _ := suite.apiClient.PrivateNetworksAPI.PrivateNetworksPost(suite.ctx).PrivateNetworkCreate(networkCreate).Force(force).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestGetPrivateNetworkById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/private_networks_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	privateNetworkId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.apiClient.PrivateNetworksAPI.PrivateNetworksNetworkIdGet(suite.ctx, privateNetworkId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(string(jsonResult), string(jsonResponseBody))

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

	privateNetworkId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.apiClient.PrivateNetworksAPI.PrivateNetworksNetworkIdPut(suite.ctx, privateNetworkId).PrivateNetworkModify(privateNetworkModify).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestDeletePrivateNetworkById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/private_networks_delete_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	privateNetworkId := request.PathParameters["id"][0]

	// Operation Execution
	result, _ := suite.apiClient.PrivateNetworksAPI.PrivateNetworksNetworkIdDelete(suite.ctx, privateNetworkId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)

	suite.Empty(string(jsonResult))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestGetPublicNetworks() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/public_networks_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	loc := fmt.Sprintf("%v", qpMap["location"])

	// Operation Execution
	result, _, _ := suite.apiClient.PublicNetworksAPI.PublicNetworksGet(suite.ctx).Location(loc).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

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
	result, _, _ := suite.apiClient.PublicNetworksAPI.PublicNetworksPost(suite.ctx).PublicNetworkCreate(networkCreate).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestGetPublicNetworkById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/public_networks_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	publicNetworkId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.apiClient.PublicNetworksAPI.PublicNetworksNetworkIdGet(suite.ctx, publicNetworkId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(string(jsonResult), string(jsonResponseBody))

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

	publicNetworkId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.apiClient.PublicNetworksAPI.PublicNetworksNetworkIdPatch(suite.ctx, publicNetworkId).PublicNetworkModify(publicNetworkModify).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestDeletePublicNetworkById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/public_networks_delete_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	publicNetworkId := request.PathParameters["id"][0]

	// Operation Execution
	result, _ := suite.apiClient.PublicNetworksAPI.PublicNetworksNetworkIdDelete(suite.ctx, publicNetworkId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)

	suite.Empty(string(jsonResult))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestPostPublicNetworkIpBlockByPublicNetworkId() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/public_networks_post_ip_block", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	body, _ := json.Marshal(request.Body.Json)
	var publicNetworkIpBlockCreate networkapi.PublicNetworkIpBlockCreate
	json.Unmarshal(body, &publicNetworkIpBlockCreate)

	publicNetworkId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.apiClient.PublicNetworksAPI.PublicNetworksNetworkIdIpBlocksPost(suite.ctx, publicNetworkId).PublicNetworkIpBlockCreate(publicNetworkIpBlockCreate).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestDeletePublicNetworkIpBlocksByPublicNetworkIdAndIpBlockID() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/public_networks_delete_ip_block_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	publicNetworkId := request.PathParameters["id"][0]
	ipId := request.PathParameters["ipId"][0]

	// Operation Execution
	result, _, _ := suite.apiClient.PublicNetworksAPI.PublicNetworksNetworkIdIpBlocksIpBlockIdDelete(suite.ctx, publicNetworkId, ipId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestGetBgpPeerGroups() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/bgp_peer_groups_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	loc := fmt.Sprintf("%v", qpMap["location"])

	// Operation Execution
	result, _, _ := suite.apiClient.BGPPeerGroupsAPI.BgpPeerGroupsGet(suite.ctx).Location(loc).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestCreateBgpPeerGroups() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/bgp_peer_groups_post", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	body, _ := json.Marshal(request.Body.Json)
	var bgpPeerGroupCreate networkapi.BgpPeerGroupCreate
	json.Unmarshal(body, &bgpPeerGroupCreate)

	// Operation Execution
	result, _, _ := suite.apiClient.BGPPeerGroupsAPI.BgpPeerGroupsPost(suite.ctx).BgpPeerGroupCreate(bgpPeerGroupCreate).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestGetBgpPeerGroupsByID() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/bgp_peer_groups_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	bgpPeerGroupId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.apiClient.BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdGet(suite.ctx, bgpPeerGroupId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestDeleteBgpPeerGroupsByID() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/bgp_peer_groups_delete_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	bgpPeerGroupId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.apiClient.BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdDelete(suite.ctx, bgpPeerGroupId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *NetworkApiTestSuite) TestPatchBgpPeerGroupsByID() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("networkapi/bgp_peer_groups_patch_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	bgpPeerGroupId := request.PathParameters["id"][0]

	body, _ := json.Marshal(request.Body.Json)
	var bgpPeerGroupPatch networkapi.BgpPeerGroupPatch
	json.Unmarshal(body, &bgpPeerGroupPatch)

	// Operation Execution
	result, _, _ := suite.apiClient.BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdPatch(suite.ctx, bgpPeerGroupId).BgpPeerGroupPatch(bgpPeerGroupPatch).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func TestNetworkApiTestSuite(t *testing.T) {
	suite.Run(t, new(NetworkApiTestSuite))
}
