package tests

import (
	"context"
	"fmt"

	"github.com/phoenixnap/go-sdk-bmc/auditapi"
	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2/clientcredentials"
)

type SmokeTestSuite struct {
	suite.Suite
	ctx           context.Context
	configuration *auditapi.Configuration
	apiClient     *auditapi.APIClient
}

func (suite *SmokeTestSuite) SetupSuite() {
	clientId := "..."
	clientSecret := "..."
	apiUrl := "https://api-dev.phoenixnap.com/audit/v1"
	tokenUrl := "https://auth-dev.phoenixnap.com/auth/realms/BMC/protocol/openid-connect/token"

	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     tokenUrl,
		Scopes:       []string{"bmc", "bmc.read"},
	}

	configuration := auditapi.NewConfiguration()

	configuration.Servers = auditapi.ServerConfigurations{
		{
			URL: apiUrl,
		},
	}

	configuration.HTTPClient = config.Client(context.Background())

	suite.configuration = configuration

	fmt.Println(">>> From SetupSuite")
	suite.apiClient = auditapi.NewAPIClient(configuration)
}

func (suite *SmokeTestSuite) TestSmoke() {
	result, _, err := suite.apiClient.EventsAPI.EventsGet(suite.ctx).Execute()
	fmt.Printf("ERROR: %v\n", err)
	fmt.Printf("RESULT: %v\n", result[0])
}

// func TestSmokeTestSuite(t *testing.T) {
// 	suite.Run(t, new(SmokeTestSuite))
// }
