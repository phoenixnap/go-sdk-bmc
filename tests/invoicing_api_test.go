package tests

import (
	"context"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/invoicingapi"
	"github.com/stretchr/testify/suite"
)

type InvoicingApiTestSuite struct {
	suite.Suite
	ctx       context.Context
	apiClient *invoicingapi.APIClient
}

func TestInvoicingApiTestSuite(t *testing.T) {
	suite.Run(t, new(InvoicingApiTestSuite))
}

func (suite *InvoicingApiTestSuite) SetupSuite() {
	// Remove any existing expectations
	TestUtilsImpl{}.resetExpectations()

	// Set configuration
	configuration := invoicingapi.NewConfiguration()
	configuration.Host = "127.0.0.1:1080"
	configuration.Scheme = "http"

	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)

	// New ApiClient
	suite.apiClient = invoicingapi.NewAPIClient(configuration)
}

func (suite *InvoicingApiTestSuite) TearDownTest() {
	TestUtilsImpl{}.resetExpectations()
}

func (suite *InvoicingApiTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

// func (suite *InvoicingApiTestSuite) TestGetLocations() {
// 	// generate payload
// 	request, response := TestUtilsImpl{}.generatePayloadsFrom("invoicingapi/locations_get", "./payloads")

// 	// extract expectation id
// 	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

// 	// fetch query parameters
// 	qpMap := TestUtilsImpl{}.generateQueryParams(request)

// 	location := fmt.Sprintf("%v", qpMap["location"])
// 	productCategory := fmt.Sprintf("%v", qpMap["productCategory"])

// 	// execution
// 	result, _, _ := suite.apiClient.LocationsAPI.
// 		GetLocations(suite.ctx).
// 		Location(invoicingapi.LocationEnum(location)).
// 		ProductCategory(invoicingapi.ProductCategoryEnum(productCategory)).
// 		Execute()

// 	// convert to json strings
// 	jsonResult, _ := json.Marshal(result)
// 	jsonResponseBody, _ := json.Marshal(response.Body)

// 	// asserts
// 	suite.Equal(jsonResult, jsonResponseBody)

// 	// verify
// 	suite.verifyCalledOnce(expectationId)
// }
