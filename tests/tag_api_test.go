package tests

import (
	"testing"
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
	request, response := TestUtilsImpl{}.generate_payloads_from("tagsapi/events_get", "./payloads")
	request.setPathParams(0)

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setup_expectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generate_query_params(request)
}