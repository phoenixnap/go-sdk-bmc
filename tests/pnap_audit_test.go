package tests

import (
	"fmt"
	"github.com/phoenixnap/go-sdk-bmc/auditapi"
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
	//"time"
	"os"
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

	//qpMap := TestUtilsImpl{}.generate_query_params(request)
	const layout = "2006-Jan-02"
	
	// from,_ := time.Parse(layout,fmt.Sprintf("%v", qpMap["from"])) // time.Time | From the date and time (inclusive) to filter event log records by. (optional)
    // to, _:= time.Parse(layout,fmt.Sprintf("%v", qpMap["to"])) // time.Time | To the date and time (inclusive) to filter event log records by. (optional)
    // limit,_ := qpMap["limit"].(int32) // int32 | Limit the number of records returned. (optional)
    // order:= fmt.Sprintf("%v", qpMap["order"]) // string | Ordering of the event's time. SortBy can be introduced later on. (optional) (default to "ASC")
    // username:= fmt.Sprintf("%v", qpMap["username"]) // string | The username that did the actions. (optional)
    // verb := fmt.Sprintf("%v", qpMap["verb"]) // string | The HTTP verb corresponding to the action. (optional)
    // uri:= fmt.Sprintf("%v", qpMap["to"])


	//ctx = context.WithValue(ctx, "accessToken", "fake")
	
	configuration := auditapi.NewConfiguration()
	configuration.Host = "127.0.0.1:1080"
	configuration.Scheme = "http"
	ctx := context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	ctx = context.WithValue(ctx, "serverIndex", nil)

    apiClient := auditapi.NewAPIClient(configuration)
	
    result, r, err := apiClient.EventsApi.EventsGet(ctx).Execute()//.From(from).To(to).Limit(limit).Order(order).Username(username).Verb(verb).Uri(uri).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }

	// //parsedTime, _ := time.Parse("MM/DD/YYYY", string(response["body"].(string)["timestamp"]))

	// //response["body"][0].(string)["timestamp"] = parsedTime.String()

	assert.Equal(t, result, response.Body)
	verify_called_once(t, expectationId)
	TestUtilsImpl{}.reset_expectations()
	return
}
