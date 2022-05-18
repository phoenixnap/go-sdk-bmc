package tests

import (
	"fmt"
	"github.com/phoenixnap/go-sdk-bmc/auditapi"
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"encoding/json"
	"strconv"
)

// func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
// 	if a == b {
// 		return
// 	}
// 	if len(message) == 0 {
// 		message = fmt.Sprintf("%v != %v", a, b)
// 	}
// 	t.Fatal(message)
// }

func verify_called_once(t *testing.T, expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verify_expectation_matched_times(expectationId, 1)

	// Asserts a successful result
	assert.Equal(t, 202, verifyResult.StatusCode)
}

func pop(m map[string]string, key string) (string, bool) {
	v, ok := m[key]
	if ok {
		delete(m, key)
	}
	return v, ok
}

func Test_get_events_all_query_params(t *testing.T) {

	// Setting up expectation
	request, response := TestUtilsImpl{}.generate_payloads_from("auditapi/events_get", "./payloads")
	request.setPathParams()
	expectationId := TestUtilsImpl{}.setup_expectation(request, response, 1)

	qpMap := TestUtilsImpl{}.generate_query_params(request)
	const layout = "2006-Jan-02"

	limitstring, ok := qpMap["limit"].(string)
	if !ok {
        panic(ok)
    }

	// from,_ := time.Parse(layout,fmt.Sprintf("%v", qpMap["from"])) // time.Time | From the date and time (inclusive) to filter event log records by. (optional)
    to, _:= time.Parse(layout,fmt.Sprintf("%v", qpMap["to"])) // time.Time | To the date and time (inclusive) to filter event log records by. (optional)
    limit64, _ := strconv.ParseInt(limitstring, 10, 32)
	limit := int32(limit64) // int32 | Limit the number of records returned. (optional)
    order:= fmt.Sprintf("%v", qpMap["order"]) // string | Ordering of the event's time. SortBy can be introduced later on. (optional) (default to "ASC")
    username:= fmt.Sprintf("%v", qpMap["username"]) // string | The username that did the actions. (optional)
    verb := fmt.Sprintf("%v", qpMap["verb"]) // string | The HTTP verb corresponding to the action. (optional)
    uri:= fmt.Sprintf("%v", qpMap["uri"])
	
	configuration := auditapi.NewConfiguration()
	configuration.Host = "127.0.0.1:1080"
	configuration.Scheme = "http"
	ctx := context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	ctx = context.WithValue(ctx, "serverIndex", nil)

    apiClient := auditapi.NewAPIClient(configuration)
	
    result, r, err := apiClient.EventsApi.EventsGet(ctx).Limit(limit).Order(order).Username(username).Verb(verb).Uri(uri).Execute()//.From(from).To(to).Limit(limit).Order(order).Username(username).Verb(verb).Uri(uri).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	assert.Equal(t, jsonResult, jsonResponseBody)
	verify_called_once(t, expectationId)

	TestUtilsImpl{}.reset_expectations()
	return
}
