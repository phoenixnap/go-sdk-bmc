package tests

import (
	"context"
	"github.com/phoenixnap/go-sdk-bmc/auditapi"
	"testing"
	"fmt"
	"time"
	)
// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'


var configuration auditapi.Configuration = auditapi.Configuration{
	Host: "127.0.0.1:1080",
	Scheme: "http",
}

var apiClient = auditapi.NewAPIClient(&configuration)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
	  return
	  }
	  if len(message) == 0 {
		  message = fmt.Sprintf("%v != %v", a, b)
	  }
	  t.Fatal(message)
  }

func verify_called_once(t *testing.T, expectationId string)(){
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verify_expectation_matched_times(expectationId, 1)

	// Asserts a successful result
	assertEqual(t, 202, verifyResult.StatusCode, "")
}

func pop(m map[string]string, key string) (string, bool) {
    v, ok := m[key]
    if ok {
        delete(m, key)
    }
    return v, ok
}

func test_get_events_all_query_params(t *testing.T) (){
	// Setting up expectation
	request, response := TestUtilsImpl{}.generate_payloads_from("auditapi/events_get", "./payloads")
	expectationId := TestUtilsImpl{}.setup_expectation(request, response, 1) 

	// Creating new instance
	apiInstance := auditapi.EventsApiService{}

	opts := TestUtilsImpl{}.generate_query_params(request)

	result := apiInstance.EventsGet(opts)
	
	parsedTime, _ := time.Parse("MM/DD/YYYY", response.body[0].timestamp)

	response.body[0].timestamp = parsedTime.String()

	assertEqual(t, response.body[0], result, "")

	verify_called_once(&testing.T ,expectationId)
	
}