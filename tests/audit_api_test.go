package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/phoenixnap/go-sdk-bmc/auditapi"
	"github.com/stretchr/testify/suite"
)

type AuditApiTestSuite struct {
	suite.Suite
	ctx           context.Context
	configuration *auditapi.Configuration
	apiClient     *auditapi.APIClient
}

func (suite *AuditApiTestSuite) SetupTest() {
	// Set configuration
	suite.configuration = auditapi.NewConfiguration()
	suite.configuration.Host = "127.0.0.1:1080"
	suite.configuration.Scheme = "http"
	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)
	fmt.Println(">>> From SetupSuite")
	suite.apiClient = auditapi.NewAPIClient(suite.configuration)
}

// this function executes after all tests executed
func (suite *AuditApiTestSuite) TearDownTest() {
	fmt.Println(">>> From TearDownSuite")
	TestUtilsImpl{}.resetExpectations()
}

func (suite *AuditApiTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *AuditApiTestSuite) Test_get_events_all_query_params() {

	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("auditapi/events_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	// Converting to int32
	limitstring, ok := qpMap["limit"].(string)
	if !ok {
		panic(ok)
	}

	from, _ := time.Parse(time.RFC3339, request.QueryStringParameters[0].Values[0]) 
	to, _ := time.Parse(time.RFC3339, request.QueryStringParameters[1].Values[0])   
	limit64, _ := strconv.ParseInt(limitstring, 10, 32)
	limit := int32(limit64)                          
	order := fmt.Sprintf("%v", qpMap["order"])       
	username := fmt.Sprintf("%v", qpMap["username"]) 
	verb := fmt.Sprintf("%v", qpMap["verb"])         
	uri := fmt.Sprintf("%v", qpMap["uri"])

	// Operation Execution
	result, r, err := suite.apiClient.EventsApi.EventsGet(suite.ctx).From(from).To(to).Limit(limit).Order(order).Username(username).Verb(verb).Uri(uri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func TestAuditApiTestSuite(t *testing.T) {
	suite.Run(t, new(AuditApiTestSuite))
}
