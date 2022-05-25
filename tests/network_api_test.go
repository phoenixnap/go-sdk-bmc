package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	"github.com/stretchr/testify/suite"
)

type NetworkApiTestSuite struct {
	suite.Suite
	ctx context.Context
	configuration *networkapi.Configuration
	apiClient *networkapi.APIClient
}

// this function executes before each test
func (suite *NetworkApiTestSuite) SetupSuite() {
	// Set configuration
	suite.configuration = networkapi.NewConfiguration()
	suite.configuration.Host = "127.0.0.1:1080"
	suite.configuration.Scheme = "http"
	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)
	fmt.Println(">>> From SetupSuite")
	suite.apiClient = networkapi.NewAPIClient(suite.configuration)
}

// this function executes after each test
func (suite *NetworkApiTestSuite) TearDownTest() {
	fmt.Println(">>> From TearDownSuite")
	defer TestUtilsImpl{}.resetExpectations()
}

func (suite *NetworkApiTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *NetworkApiTestSuite) TestGetPrivateNetworks() {

}

func (suite *NetworkApiTestSuite) TestCreatePrivateNetwork() {

}

func (suite *NetworkApiTestSuite) TestGetPrivateNetworkById() {

}

func (suite *NetworkApiTestSuite) TestPutPrivateNetworkById() {

}

func (suite *NetworkApiTestSuite) TestDeletePrivateNetworkById() {

}

func (suite *NetworkApiTestSuite) TestGetPublicNetworks() {

}

func (suite *NetworkApiTestSuite) TestCreatePublicNetwork() {

}

func (suite *NetworkApiTestSuite) TestGetPublicNetworkById() {

}

func (suite *NetworkApiTestSuite) TestPatchPublicNetworkById() {

}

func (suite *NetworkApiTestSuite) TestDeletePublicNetworkById() {

}

func (suite *NetworkApiTestSuite) TestPostPublicNetworkIpBlockByPublicNetworkId() {

}

func (suite *NetworkApiTestSuite) TestDeletePublicNetworkIpBlocksByPublicNetworkIdAndIpBlockID() {

}

func TestNetworkApiTestSuite(t *testing.T) {
	suite.Run(t, new(NetworkApiTestSuite))
}