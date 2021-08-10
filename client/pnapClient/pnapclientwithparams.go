package pnapClient

import (
	"github.com/phoenixnap/go-sdk-bmc/dto"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/phoenixnap/go-sdk-bmc/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"golang.org/x/oauth2/clientcredentials"
)

// PNAPClient is a Client that performs HTTP requests.
type PNAPClient struct {
	client *http.Client
	auth dto.Authentication
}

//NewPNAPClientWithDefaultConfig creates a new PNAPClient. Verification of configuration will be done prior to creation
//and error will be returned in case credentials or whole configuration file is missing
func NewPNAPClientWithDefaultConfig() (PNAPClient, error) {
	// Find home directory
	home, err := homedir.Dir()
	if err != nil {
		return PNAPClient{}, err
	}

	configPath := home + config.DefaultConfigPath
	confPathErr := Verify(configPath)
	if confPathErr != nil {
		return PNAPClient{}, confPathErr
	}


	config := load(configPath)
	httpClient := config.Client(context.Background())
	pnapClient := PNAPClient{httpClient, dto.Authentication{}}
	return pnapClient, err
}

// NewPNAPClient creates a new PNAPClient with forwarded credentials
func NewPNAPClient(auth dto.Authentication) PNAPClient {
	config := clientcredentials.Config{
		ClientID:     auth.ClientID,
		ClientSecret: auth.ClientSecret,
		TokenURL:     auth.TokenURL,
		Scopes:       []string{"bmc", "bmc.read"},
	}
	httpClient := config.Client(context.Background())
	pnapClient := PNAPClient{httpClient, auth}
	return pnapClient
}
//NewPNAPClientWithCustomConfig creates a new PNAPClient. Verification of configuration will be done prior to creation
//and error will be returned in case credentials or whole configuration file is missing
func NewPNAPClientWithCustomConfig(path string) (PNAPClient, error) {
	err := Verify(path)
	if err != nil {
		return PNAPClient{}, err
	}
	config := load(path)
	httpClient := config.Client(context.Background())
	pnapClient := PNAPClient{httpClient, dto.Authentication{}}
	return pnapClient, err
}

func load(configPath string) clientcredentials.Config {
	
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.ReadInConfig()
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

//Verify verifies existence of configuration file and credentials
func Verify(configPath string) error {
	

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// Checks whether the config file exists, by attempting to cast the error.
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("A config file is required in order to proceed.\n" +
				"Config file path is the home directory (" + configPath + "config.yaml)\n\n" +
				"The following shows a sample config file:\n\n" +
				"# =====================================================\n" +
				"# Sample yaml config file\n" +
				"# =====================================================\n\n" +
				"# Authentication\n" +
				"clientId: <enter your client id>\n" +
				"clientSecret: <enter your client secret>")
		}
		return err
	}
	clientID := viper.GetString("clientId")
	if clientID == "" {
		return fmt.Errorf("API client ID not found in configuration")
	}
	clientSecret := viper.GetString("clientSecret")
	if clientSecret == "" {
		return fmt.Errorf("API clientSecret not found in configuration")
	}
	return nil
}

// Get performs a Get request and check for auth errors
func (pnapClient PNAPClient) Get(resource string) (*http.Response, error) {

	response, err := pnapClient.client.Get(buildURI(resource, pnapClient.auth))

	return response, err
}

// Delete performs a Delete request and check for auth errors
func (pnapClient PNAPClient) Delete(resource string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", buildURI(resource, pnapClient.auth), nil)
	// replicating Get/Post error handling
	if err != nil {
		return nil, err
	}
	return pnapClient.client.Do(req)
}

// Post performs a Post request and check for auth errors
func (pnapClient PNAPClient) Post(resource string, body io.Reader) (*http.Response, error) {
	return pnapClient.client.Post(buildURI(resource, pnapClient.auth), "application/json", body)
}

// Put performs a Put request and check for auth errors
func (pnapClient PNAPClient) Put(resource string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("PUT", buildURI(resource, pnapClient.auth), body)
	if err != nil {
		return nil, err
	}
	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	return pnapClient.client.Do(req)
}

/* func buildURI(resource string) string {
	return config.Hostname + resource
} */
func buildURI(resource string, auth dto.Authentication) string {
	if auth.ApiHostName != "" && auth.PoweredBy != ""{
		return auth.ApiHostName + resource + "?_xPoweredBy=" + auth.PoweredBy
	}else if auth.ApiHostName != ""{
		return auth.ApiHostName + resource
	}
	return config.Hostname + resource
}
