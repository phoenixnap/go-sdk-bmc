package tests

import (
	"context"
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
