package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/phoenixnap/go-sdk-bmc/paymentsapi"
	"github.com/stretchr/testify/suite"
)

type PaymentsApiTestSuite struct {
	suite.Suite
	ctx       context.Context
	apiClient *paymentsapi.APIClient
}

func TestPaymentsApiTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentsApiTestSuite))
}

func (suite *PaymentsApiTestSuite) SetupSuite() {
	// Remove any existing expectations
	TestUtilsImpl{}.resetExpectations()

	// Set configuration
	configuration := paymentsapi.NewConfiguration()
	configuration.Host = "127.0.0.1:1080"
	configuration.Scheme = "http"

	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)

	// New ApiClient
	suite.apiClient = paymentsapi.NewAPIClient(configuration)
}

func (suite *PaymentsApiTestSuite) TearDownTest() {
	TestUtilsImpl{}.resetExpectations()
}

func (suite *PaymentsApiTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *PaymentsApiTestSuite) TestGetAllTransactions() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("paymentsapi/transactions_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	// Extracting into variables
	limit, _ := strconv.ParseInt(fmt.Sprintf("%v", qpMap["limit"]), 10, 32)
	offset, _ := strconv.ParseInt(fmt.Sprintf("%v", qpMap["offset"]), 10, 32)
	sortDirection := fmt.Sprintf("%v", qpMap["sortDirection"])
	sortField := fmt.Sprintf("%v", qpMap["sortField"])
	from, _ := time.Parse(time.RFC3339, fmt.Sprintf("%v", qpMap["from"]))
	to, _ := time.Parse(time.RFC3339, fmt.Sprintf("%v", qpMap["to"]))

	// Execution
	result, r, err := suite.apiClient.TransactionsAPI.TransactionsGet(context.Background()).Limit(int32(limit)).Offset(int32(offset)).SortDirection(sortDirection).SortField(sortField).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling 'TransactionsAPI.TransactionsGet': %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *PaymentsApiTestSuite) TestGetTransactionById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("paymentsapi/transactions_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extracting path parameter
	transactionId := request.PathParameters["id"][0]

	// Execution
	result, r, err := suite.apiClient.TransactionsAPI.TransactionIdGet(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling 'TransactionsAPI.TransactionsGet': %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}
