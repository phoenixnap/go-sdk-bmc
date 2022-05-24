package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/tagapi"
	"github.com/stretchr/testify/suite"
)

type TagApiTestSuite struct {
	suite.Suite
	ctx           context.Context
	configuration *tagapi.Configuration
	apiClient     *tagapi.APIClient
}

// this function executes before each test
func (suite *TagApiTestSuite) SetupTest() {
	// Set configuration
	suite.configuration = tagapi.NewConfiguration()
	suite.configuration.Host = "127.0.0.1:1080"
	suite.configuration.Scheme = "http"
	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)
	fmt.Println(">>> From SetupSuite")
	suite.apiClient = tagapi.NewAPIClient(suite.configuration)
}

// this function executes after each test
func (suite *TagApiTestSuite) TearDownTest() {
	fmt.Println(">>> From TearDownSuite")
	defer TestUtilsImpl{}.resetExpectations()
}

func verifyCalledOnce(suite *TagApiTestSuite, expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *TagApiTestSuite) TestGetTags() {

	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("tagapi/tags_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	name := fmt.Sprintf("%v", qpMap["name"])

	// Operation Execution
	result, r, err := suite.apiClient.TagsApi.TagsGet(suite.ctx).Name(name).Execute()
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
	verifyCalledOnce(suite, expectationId)

}

func (suite *TagApiTestSuite) TestCreateTags() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("tagapi/tags_post", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	byteData := TestUtilsImpl{}.extractRequestBody(request)
	var tagCreate tagapi.TagCreate
	json.Unmarshal(byteData, &tagCreate)

	// Operation Execution
	result, r, err := suite.apiClient.TagsApi.TagsPost(suite.ctx).TagCreate(tagCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.TagsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	verifyCalledOnce(suite, expectationId)
}

func (suite *TagApiTestSuite) TestGetTagById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("tagapi/tags_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	tagId := TestUtilsImpl{}.extractIdFrom(request)

	// Operation Execution
	result, r, err := suite.apiClient.TagsApi.TagsTagIdGet(suite.ctx, tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.TagsTagIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	verifyCalledOnce(suite, expectationId)
}

func (suite *TagApiTestSuite) TestPatchTagById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("tagapi/tags_patch_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	byteData := TestUtilsImpl{}.extractRequestBody(request)
	var tagUpdate tagapi.TagUpdate
	json.Unmarshal(byteData, &tagUpdate)

	tagId := TestUtilsImpl{}.extractIdFrom(request)

	// Operation Execution
	result, r, err := suite.apiClient.TagsApi.TagsTagIdPatch(suite.ctx, tagId).TagUpdate(tagUpdate).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.TagsTagIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	verifyCalledOnce(suite, expectationId)
}

func (suite *TagApiTestSuite) TestDeleteTagById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("tagapi/tags_delete_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	tagId := TestUtilsImpl{}.extractIdFrom(request)

	// Operation Execution
	result, r, err := suite.apiClient.TagsApi.TagsTagIdDelete(suite.ctx, tagId).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.TagsTagIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	suite.Equal(jsonResult, jsonResponseBody)

	// Verify
	verifyCalledOnce(suite, expectationId)

}

func TestTagApiTestSuite(t *testing.T) {
	suite.Run(t, new(TagApiTestSuite))
}
