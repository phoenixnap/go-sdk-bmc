package tests

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/ipapi"
	"github.com/stretchr/testify/suite"
)

type IpApiTestSuite struct {
	suite.Suite
	Client *ipapi.APIClient
	Ctx    context.Context
}

func (suite *IpApiTestSuite) SetupSuite() {
	// Set configuration
	configuration := ipapi.NewConfiguration()
	configuration.Host = "127.0.0.1:1080"
	configuration.Scheme = "http"

	// Set the context background
	suite.Ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.Ctx = context.WithValue(suite.Ctx, "serverIndex", nil)

	// New ApiClient
	suite.Client = ipapi.NewAPIClient(configuration)
}

func (suite *IpApiTestSuite) TearDownTest() {
	TestUtilsImpl{}.reset_expectations()
}

func (suite *IpApiTestSuite) VerifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verify_expectation_matched_times(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *IpApiTestSuite) TestGetIpBlocks() {
	// Generate payload
	request, response := TestUtilsImpl{}.generate_payloads_from("ipapi/ip_blocks_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setup_expectation(request, response, 1)

	// Operation Execution
	result, _, _ := suite.Client.IPBlocksApi.IpBlocksGet(suite.Ctx).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.VerifyCalledOnce(expectationId)
}

func (suite *IpApiTestSuite) TestCreateIpBlock() {
	// Generate payload
	request, response := TestUtilsImpl{}.generate_payloads_from("ipapi/ip_blocks_post", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setup_expectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var ipBlockCreate ipapi.IpBlockCreate
	json.Unmarshal(body, &ipBlockCreate)

	// Operation Execution
	result, _, _ := suite.Client.IPBlocksApi.IpBlocksPost(suite.Ctx).IpBlockCreate(ipBlockCreate).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.VerifyCalledOnce(expectationId)
}

func (suite *IpApiTestSuite) TestGetIpBlocksById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generate_payloads_from("ipapi/ip_blocks_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setup_expectation(request, response, 1)

	// Extract the ipBlockId
	ipBlockId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.IPBlocksApi.IpBlocksIpBlockIdGet(suite.Ctx, ipBlockId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.VerifyCalledOnce(expectationId)
}

func (suite *IpApiTestSuite) TestDeleteIpBlocksById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generate_payloads_from("ipapi/ip_blocks_delete_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setup_expectation(request, response, 1)

	// Extract the ipBlockId
	ipBlockId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.IPBlocksApi.IpBlocksIpBlockIdDelete(suite.Ctx, ipBlockId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.VerifyCalledOnce(expectationId)
}

func (suite *IpApiTestSuite) TestPatchIpBlocksById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generate_payloads_from("ipapi/ip_blocks_patch_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setup_expectation(request, response, 1)

	// Extract the ipBlockId
	ipBlockId := request.PathParameters["id"][0]

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var ipBlockPatch ipapi.IpBlockPatch
	json.Unmarshal(body, &ipBlockPatch)

	// Operation Execution
	result, _, _ := suite.Client.IPBlocksApi.IpBlocksIpBlockIdPatch(suite.Ctx, ipBlockId).IpBlockPatch(ipBlockPatch).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.VerifyCalledOnce(expectationId)
}

func (suite *IpApiTestSuite) TestPutTagsIpBlocksById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generate_payloads_from("ipapi/ip_blocks_put_tags_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setup_expectation(request, response, 1)

	// Extract the ipBlockId
	ipBlockId := request.PathParameters["id"][0]

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var tagAssignmentRequest []ipapi.TagAssignmentRequest
	json.Unmarshal(body, &tagAssignmentRequest)

	// Operation Execution
	result, _, _ := suite.Client.IPBlocksApi.IpBlocksIpBlockIdTagsPut(suite.Ctx, ipBlockId).TagAssignmentRequest(tagAssignmentRequest).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.VerifyCalledOnce(expectationId)
}

func TestIpApiTestSuite(t *testing.T) {
	suite.Run(t, new(IpApiTestSuite))
}
