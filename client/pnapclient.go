package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/PNAP/bmc-api-sdk/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"golang.org/x/oauth2/clientcredentials"
)

// PNAPClient is a Client that performs HTTP requests.
type PNAPClient struct {
	client *http.Client
}

// Create creates a new PNAPClient
func Create() PNAPClient {
	config := loadConfiguration()
	httpClient := config.Client(context.Background())
	pnapClient := PNAPClient{httpClient}
	return pnapClient
}

// NewPNAPClient creates a new PNAPClient with forwarded credentials
func NewPNAPClient(clientID string, clientSecret string) PNAPClient {
	config := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     config.TokenURL,
		Scopes:       []string{"bmc", "bmc.read"},
	}
	httpClient := config.Client(context.Background())
	pnapClient := PNAPClient{httpClient}
	return pnapClient
}

func loadConfiguration() clientcredentials.Config {
	// Find home directory
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(home)
	configPath := home + config.DefaultConfigPath
	fmt.Println(configPath)
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// Checks whether the config file exists, by attempting to cast the error.
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("A config file is required in order to proceed.\n" +
				"Config file path is the home directory (" + configPath + "config.yaml)\n\n" +
				"The following shows a sample config file:\n\n" +
				"# =====================================================\n" +
				"# Sample yaml config file\n" +
				"# =====================================================\n\n" +
				"# Authentication\n" +
				"clientId: <enter your client id>\n" +
				"clientSecret: <enter your client secret>")
		} else {
			fmt.Println("Error reading config file:", err)
		}
		os.Exit(1)
	}
	clientID := viper.GetString("clientId")
	clientSecret := viper.GetString("clientSecret")
	tokenURL := config.TokenURL

	config := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		Scopes:       []string{"bmc", "bmc.read"},
	}
	return config
}

// Get performs a Get request and check for auth errors
func (pnapClient PNAPClient) Get(resource string) (*http.Response, error) {

	response, err := pnapClient.client.Get(buildURI(resource))

	// if e, isUrlError := err.(*url.Error); isUrlError && strings.Contains(err.Error(), "oauth2: cannot fetch token") {
	// 	//Timeout If there is an error it must have happened while resolving token
	// 	// ErrorURLs frome the actual request should be represented in the body
	// 	return response, ctlerrors.Error{Msg: "Failed to resolve provided credentials", Cause: e}
	// }

	return response, err
}

// Delete performs a Delete request and check for auth errors
func (pnapClient PNAPClient) Delete(resource string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", buildURI(resource), nil)
	// replicating Get/Post error handling
	if err != nil {
		return nil, err
	}
	return pnapClient.client.Do(req)
}

// Post performs a Post request and check for auth errors
func (pnapClient PNAPClient) Post(resource string, body io.Reader) (*http.Response, error) {
	return pnapClient.client.Post(buildURI(resource), "application/json", body)
}

func buildURI(resource string) string {
	return config.Hostname + resource
}
