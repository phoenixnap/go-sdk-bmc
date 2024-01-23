package tests

import (
	"context"
	"encoding/json"
	"testing"

	"fmt"
	"strconv"

	"github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/stretchr/testify/suite"
)

type BmcApiTestSuite struct {
	suite.Suite
	Client *bmcapi.APIClient
	Ctx    context.Context
}

func (suite *BmcApiTestSuite) SetupSuite() {
	// Remove any existing expectations
	TestUtilsImpl{}.resetExpectations()

	// Set configuration
	configuration := bmcapi.NewConfiguration()
	configuration.Host = "127.0.0.1:1080"
	configuration.Scheme = "http"

	// Set the context background
	suite.Ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.Ctx = context.WithValue(suite.Ctx, "serverIndex", nil)

	// New ApiClient
	suite.Client = bmcapi.NewAPIClient(configuration)
}

func (suite *BmcApiTestSuite) TearDownTest() {
	TestUtilsImpl{}.resetExpectations()
}

func (suite *BmcApiTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *BmcApiTestSuite) TestGetQuotas() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/quotas/quotas_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Operation Execution
	result, _, _ := suite.Client.QuotasAPI.QuotasGet(suite.Ctx).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestQuotasCreateEditRequest() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/quotas/quotas_action_request_edit", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract the quotaId
	quotaId := request.PathParameters["id"][0]

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var quotaEditLimitRequest bmcapi.QuotaEditLimitRequest
	json.Unmarshal(body, &quotaEditLimitRequest)

	// Operation Execution
	suite.Client.QuotasAPI.QuotasQuotaIdActionsRequestEditPost(suite.Ctx, quotaId).QuotaEditLimitRequest(quotaEditLimitRequest).Execute()

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestGetQuotaById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/quotas/quotas_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract the quotaId
	quotaId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.QuotasAPI.QuotasQuotaIdGet(suite.Ctx, quotaId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestGetSshKeys() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/sshkeys/sshkeys_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Operation Execution
	result, _, _ := suite.Client.SSHKeysAPI.SshKeysGet(suite.Ctx).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestCreateSshKey() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/sshkeys/sshkeys_post", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var sshKeyCreate bmcapi.SshKeyCreate
	json.Unmarshal(body, &sshKeyCreate)

	// Operation Execution
	result, _, _ := suite.Client.SSHKeysAPI.SshKeysPost(suite.Ctx).SshKeyCreate(sshKeyCreate).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestGetSshKeyById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/sshkeys/sshkeys_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract the quotaId
	sshKeyId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.SSHKeysAPI.SshKeysSshKeyIdGet(suite.Ctx, sshKeyId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestDeleteSshKeyById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/sshkeys/sshkeys_delete_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract the quotaId
	sshKeyId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.SSHKeysAPI.SshKeysSshKeyIdDelete(suite.Ctx, sshKeyId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestPutSshKeyById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/sshkeys/sshkeys_put_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var sshKeyUpdate bmcapi.SshKeyUpdate
	json.Unmarshal(body, &sshKeyUpdate)

	// Extract the quotaId
	sshKeyId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.SSHKeysAPI.SshKeysSshKeyIdPut(suite.Ctx, sshKeyId).SshKeyUpdate(sshKeyUpdate).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestDeletePrivateNetworkById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_delete_private_network_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]
	networkId := request.PathParameters["networkId"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.DeletePrivateNetwork(suite.Ctx, serverId, networkId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestGetServers() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)
	tag := fmt.Sprintf("%v", qpMap["tag"])

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersGet(suite.Ctx).Tag([]string{tag}).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestCreateServer() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_post", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var serverCreate bmcapi.ServerCreate
	json.Unmarshal(body, &serverCreate)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)
	force, _ := strconv.ParseBool(qpMap["force"].(string))

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersPost(suite.Ctx).ServerCreate(serverCreate).Force(force).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestProvisionServer() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_action_provision", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var serverProvision bmcapi.ServerProvision
	json.Unmarshal(body, &serverProvision)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdActionsProvisionPost(suite.Ctx, serverId).ServerProvision(serverProvision).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))
	suite.Equal(string("bla"), string("not working"))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestDeprovisionServer() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_action_deprovision", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var relinquishIpBlock bmcapi.RelinquishIpBlock
	json.Unmarshal(body, &relinquishIpBlock)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdActionsDeprovisionPost(suite.Ctx, serverId).RelinquishIpBlock(relinquishIpBlock).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestPowerOffServer() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_action_power_off", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdActionsPowerOffPost(suite.Ctx, serverId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestPowerOnServer() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_action_power_on", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdActionsPowerOnPost(suite.Ctx, serverId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestRebootServer() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_action_reboot", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdActionsRebootPost(suite.Ctx, serverId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestReserveServer() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_action_reserve", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var serverReserve bmcapi.ServerReserve
	json.Unmarshal(body, &serverReserve)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdActionsReservePost(suite.Ctx, serverId).ServerReserve(serverReserve).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestResetServer() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_action_reset", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var serverReset bmcapi.ServerReset
	json.Unmarshal(body, &serverReset)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdActionsResetPost(suite.Ctx, serverId).ServerReset(serverReset).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestShutdownServer() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_action_shutdown", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdActionsShutdownPost(suite.Ctx, serverId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestDeleteServerById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_delete_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdDelete(suite.Ctx, serverId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestGetServerById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdGet(suite.Ctx, serverId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestDeleteIpBlockOnServerById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_delete_ip_block_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var relinquishIpBlock bmcapi.RelinquishIpBlock
	json.Unmarshal(body, &relinquishIpBlock)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]
	ipBlockId := request.PathParameters["ipId"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdIpBlocksIpBlockIdDelete(suite.Ctx, serverId, ipBlockId).RelinquishIpBlock(relinquishIpBlock).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestCreateIpBlockOnServerById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_post_ip_block", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var serverIpBlock bmcapi.ServerIpBlock
	json.Unmarshal(body, &serverIpBlock)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdIpBlocksPost(suite.Ctx, serverId).ServerIpBlock(serverIpBlock).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestPatchServerById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_patch_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var serverPatch bmcapi.ServerPatch
	json.Unmarshal(body, &serverPatch)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdPatch(suite.Ctx, serverId).ServerPatch(serverPatch).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestCreatePrivateNetworkOnServerById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_post_private_networks", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var serverPrivateNetwork bmcapi.ServerPrivateNetwork
	json.Unmarshal(body, &serverPrivateNetwork)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)
	force, _ := strconv.ParseBool(qpMap["force"].(string))

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdPrivateNetworksPost(suite.Ctx, serverId).ServerPrivateNetwork(serverPrivateNetwork).Force(force).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestDeletePublicNetworkOnServerById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_delete_public_network_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]
	publicNetworkId := request.PathParameters["networkId"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdPublicNetworksDelete(suite.Ctx, serverId, publicNetworkId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestCreatePublicNetworkOnServerById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_post_public_networks", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var serverPublicNetwork bmcapi.ServerPublicNetwork
	json.Unmarshal(body, &serverPublicNetwork)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)
	force, _ := strconv.ParseBool(qpMap["force"].(string))

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdPublicNetworksPost(suite.Ctx, serverId).ServerPublicNetwork(serverPublicNetwork).Force(force).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BmcApiTestSuite) TestUpdateTagsOnServerById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("bmcapi/servers/servers_put_tags_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var tagAssignmentRequest []bmcapi.TagAssignmentRequest
	json.Unmarshal(body, &tagAssignmentRequest)

	// Extract the quotaId
	serverId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ServersAPI.ServersServerIdTagsPut(suite.Ctx, serverId).TagAssignmentRequest(tagAssignmentRequest).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func TestBmcApiTestSuite(t *testing.T) {
	suite.Run(t, new(BmcApiTestSuite))
}
