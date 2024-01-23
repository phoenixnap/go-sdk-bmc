package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/locationapi"
	"github.com/stretchr/testify/suite"
)

type LocationApiTestSuite struct {
	suite.Suite
	ctx       context.Context
	apiClient *locationapi.APIClient
}

func TestLocationApiTestSuite(t *testing.T) {
	suite.Run(t, new(LocationApiTestSuite))
}

func (suite *LocationApiTestSuite) SetupSuite() {
	// Remove any existing expectations
	TestUtilsImpl{}.resetExpectations()

	// Set configuration
	configuration := locationapi.NewConfiguration()
	configuration.Host = "127.0.0.1:1080"
	configuration.Scheme = "http"

	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)

	// New ApiClient
	suite.apiClient = locationapi.NewAPIClient(configuration)
}

func (suite *LocationApiTestSuite) TearDownTest() {
	TestUtilsImpl{}.resetExpectations()
}

func (suite *LocationApiTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *LocationApiTestSuite) TestGetLocations() {
	// generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("locationapi/locations_get", "./payloads")

	// extract expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// fetch query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	location := fmt.Sprintf("%v", qpMap["location"])
	productCategory := fmt.Sprintf("%v", qpMap["productCategory"])

	// execution
	result, _, _ := suite.apiClient.LocationsAPI.
		GetLocations(suite.ctx).
		Location(locationapi.LocationEnum(location)).
		ProductCategory(locationapi.ProductCategoryEnum(productCategory)).
		Execute()

	// convert to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// verify
	suite.verifyCalledOnce(expectationId)
}
