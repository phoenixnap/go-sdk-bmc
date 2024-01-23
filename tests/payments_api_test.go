package tests

import (
	"context"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/paymentsapi"
	"github.com/stretchr/testify/suite"
)

type PaymentsApiTestSuite struct {
	suite.Suite
	ctx       context.Context
	apiClient *paymentsapi.APIClient
}

func TestPaymentsApiTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentsApiTestSuite))
}

func (suite *PaymentsApiTestSuite) SetupSuite() {
	// Remove any existing expectations
	TestUtilsImpl{}.resetExpectations()

	// Set configuration
	configuration := paymentsapi.NewConfiguration()
	configuration.Host = "127.0.0.1:1080"
	configuration.Scheme = "http"

	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)

	// New ApiClient
	suite.apiClient = paymentsapi.NewAPIClient(configuration)
}

func (suite *PaymentsApiTestSuite) TearDownTest() {
	TestUtilsImpl{}.resetExpectations()
}

func (suite *PaymentsApiTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

// func (suite *PaymentsApiTestSuite) TestGetLocations() {
// 	// generate payload
// 	request, response := TestUtilsImpl{}.generatePayloadsFrom("paymentsapi/locations_get", "./payloads")

// 	// extract expectation id
// 	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

// 	// fetch query parameters
// 	qpMap := TestUtilsImpl{}.generateQueryParams(request)

// 	location := fmt.Sprintf("%v", qpMap["location"])
// 	productCategory := fmt.Sprintf("%v", qpMap["productCategory"])

// 	// execution
// 	result, _, _ := suite.apiClient.LocationsAPI.
// 		GetLocations(suite.ctx).
// 		Location(paymentsapi.LocationEnum(location)).
// 		ProductCategory(paymentsapi.ProductCategoryEnum(productCategory)).
// 		Execute()

// 	// convert to json strings
// 	jsonResult, _ := json.Marshal(result)
// 	jsonResponseBody, _ := json.Marshal(response.Body)

// 	// asserts
// 	suite.Equal(jsonResult, jsonResponseBody)

// 	// verify
// 	suite.verifyCalledOnce(expectationId)
// }
