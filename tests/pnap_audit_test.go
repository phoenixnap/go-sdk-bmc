package tests

import (
	"net/http"
	"github.com/phoenixnap/go-sdk-bmc/auditapi"
	"testing"
	"fmt"
	)
// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'


var configuration = auditapi.Configuration{
	Host: "127.0.0.1:1080",
	Scheme: "http",
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
	  return
	  }
	  if len(message) == 0 {
		  message = fmt.Sprintf("%v != %v", a, b)
	  }
	  t.Fatal(message)
  }

// var api_client = auditapi.APIClient.

func verify_called_once(t *testing.T, expectationId string)(){
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verify_expectation_matched_times(expectationId, 1)

	assertEqual(t, 202, verifyResult.StatusCode, "")
}