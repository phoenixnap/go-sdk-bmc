package tests

import (
	"os"
	"fmt"
	"testing"
	"context"
	"encoding/json"
	"github.com/phoenixnap/go-sdk-bmc/tagapi"
	"github.com/stretchr/testify/suite"
)

type TagApiTestSuite struct {
	suite.Suite
	ctx           context.Context
	configuration *tagapi.Configuration
}

func (suite *TagApiTestSuite) SetupTest() {
	// Set configuration
	suite.configuration = tagapi.NewConfiguration()
	suite.configuration.Host = "127.0.0.1:1080"
	suite.configuration.Scheme = "http"
	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)
	fmt.Println(">>> From SetupSuite")
}

// this function executes after all tests executed
func (suite *TagApiTestSuite) TearDownTest() {
	fmt.Println(">>> From TearDownSuite")
	TestUtilsImpl{}.reset_expectations()
}

func verify_called_once(suite *TagApiTestSuite, expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verify_expectation_matched_times(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *TagApiTestSuite) Test_get_tags() {

	// Generate payload
	request, response := TestUtilsImpl{}.generate_payloads_from("tagapi/tags_get", "./payloads")
	
	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setup_expectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generate_query_params(request)

	name:= fmt.Sprintf("%v", qpMap["name"])

	// New  ApiClient
    apiClient := tagapi.NewAPIClient(suite.configuration)

	// Operation Execution
    result, r, err := apiClient.TagsApi.TagsGet(suite.ctx).Name(name).Execute()//.From(from).To(to).Limit(limit).Order(order).Username(username).Verb(verb).Uri(uri).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.TagsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	verify_called_once(suite, expectationId)

}

func (suite *TagApiTestSuite) Test_create_tags () {
// Generate payload
request, response := TestUtilsImpl{}.generate_payloads_from("tagapi/tags_get", "./payloads")
	
// Extract the response expectation id
expectationId := TestUtilsImpl{}.setup_expectation(request, response, 1)

// New  ApiClient
apiClient := tagapi.NewAPIClient(suite.configuration)
byteData := TestUtilsImpl{}.extract_request_body(request)
var tagCreate tagapi.TagCreate
json.Unmarshal(byteData, &tagCreate)

// Operation Execution
result, r, err := apiClient.TagsApi.TagsPost(suite.ctx).TagCreate(tagCreate).Execute()
if err != nil {
	fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.TagsPost``: %v\n", err)
	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
}
}

func TestTagApiTestSuite(t *testing.T) {
	suite.Run(t, new(TagApiTestSuite))
}