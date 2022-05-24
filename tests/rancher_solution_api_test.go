package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"github.com/stretchr/testify/suite"
)

type RancherSolutionApiTestSuite struct {
	suite.Suite
	ctx context.Context
	configuration *ranchersolutionapi.Configuration
	apiClient *ranchersolutionapi.APIClient
}

// this function executes before each test
func (suite *RancherSolutionApiTestSuite) SetupTest() {
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

func verifyCalledOnce(suite *RancherSolutionApiTestSuite, expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func TestRancherSolutionApiTestSuite(t *testing.T){
	suite.Run(t, new(RancherSolutionApiTestSuite))
}