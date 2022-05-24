package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"github.com/stretchr/testify/suite"
)

type RancherSolutionApiTestSuite struct {
	suite.Suite
	ctx           context.Context
	configuration *ranchersolutionapi.Configuration
	apiClient     *ranchersolutionapi.APIClient
}

// this function executes before each test
func (suite *RancherSolutionApiTestSuite) SetupSuite() {
	// Set configuration
	suite.configuration = ranchersolutionapi.NewConfiguration()
	suite.configuration.Host = "127.0.0.1:1080"
	suite.configuration.Scheme = "http"
	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)
	fmt.Println(">>> From SetupSuite")
	suite.apiClient = ranchersolutionapi.NewAPIClient(suite.configuration)
}

// this function executes after each test
func (suite *RancherSolutionApiTestSuite) TearDownTest() {
	fmt.Println(">>> From TearDownSuite")
	defer TestUtilsImpl{}.resetExpectations()
}

func (suite *RancherSolutionApiTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *RancherSolutionApiTestSuite) TestGetClusters() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("rancherapi/clusters_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Operation Execution
	result, _, _ := suite.apiClient.ClustersApi.ClustersGet(suite.ctx).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *RancherSolutionApiTestSuite) TestCreateClusters() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("rancherapi/clusters_post", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	byteData := TestUtilsImpl{}.extractRequestBody(request)
	var clusterCreate ranchersolutionapi.Cluster
	json.Unmarshal(byteData, &clusterCreate)

	// Operation Execution
	result, _, _ := suite.apiClient.ClustersApi.ClustersPost(suite.ctx).Cluster(clusterCreate).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)

}

func (suite *RancherSolutionApiTestSuite) TestGetClusterById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("rancherapi/clusters_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	clusterId := TestUtilsImpl{}.extractIdFrom(request)

	// Operation Execution
	result, _, _ := suite.apiClient.ClustersApi.ClustersIdGet(suite.ctx, clusterId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *RancherSolutionApiTestSuite) TestDeleteClusterById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("rancherapi/clusters_delete_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	clusterId := TestUtilsImpl{}.extractIdFrom(request)

	// Operation Execution
	result, _, _ := suite.apiClient.ClustersApi.ClustersIdDelete(suite.ctx, clusterId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func TestRancherSolutionApiTestSuite(t *testing.T) {
	suite.Run(t, new(RancherSolutionApiTestSuite))
}
