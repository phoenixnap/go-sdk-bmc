package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/auditapi"
	"github.com/stretchr/testify/suite"
)

type MiscTestSuite struct {
	suite.Suite
	ctx           context.Context
	configuration *auditapi.Configuration
	apiClient     *auditapi.APIClient
}

func (suite *MiscTestSuite) SetupSuite() {
	// Remove any existing expectations
	TestUtilsImpl{}.resetExpectations()

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
func (suite *MiscTestSuite) TearDownTest() {
	fmt.Println(">>> From TearDownSuite")
	TestUtilsImpl{}.resetExpectations()
}

func (suite *MiscTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *MiscTestSuite) TestResponseWithAdditionalProperties() {

	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("misc/events_get_extra_props", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Operation Execution
	result, r, err := suite.apiClient.EventsAPI.EventsGet(suite.ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// Convert the result and response body to json strings
	// If additional properties are reserved, both of these should be equal.
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func TestMiscTestSuite(t *testing.T) {
	suite.Run(t, new(MiscTestSuite))
}
