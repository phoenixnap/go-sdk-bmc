package tests

import (
	"os"
	"fmt"
	"testing"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/phoenixnap/go-sdk-bmc/tagapi"
)

func tearDown() {
	TestUtilsImpl{}.reset_expectations()
}

func verify_called_once(t *testing.T, expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verify_expectation_matched_times(expectationId, 1)

	// Asserts a successful result
	assert.Equal(t, 202, verifyResult.StatusCode)
}

func Test_get_tags(t *testing.T) {
	// Generate payload
	request, response := TestUtilsImpl{}.generate_payloads_from("tagapi/tags_get", "./payloads")
	request.setPathParams(0)

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setup_expectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generate_query_params(request)

	name:= fmt.Sprintf("%v", qpMap["name"])

	configuration := tagapi.NewConfiguration()
	configuration.Host = "127.0.0.1:1080"
	configuration.Scheme = "http"

	// Set the context background
	ctx := context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	ctx = context.WithValue(ctx, "serverIndex", nil)

	// New  ApiClient
    apiClient := tagapi.NewAPIClient(configuration)

	// Operation Execution
    result, r, err := apiClient.TagsApi.TagsGet(ctx).Name(name).Execute()//.From(from).To(to).Limit(limit).Order(order).Username(username).Verb(verb).Uri(uri).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	assert.Equal(t, jsonResult, jsonResponseBody)

	// Verify
	verify_called_once(t, expectationId)

	tearDown()
}