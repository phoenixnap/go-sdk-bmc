package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/phoenixnap/go-sdk-bmc/invoicingapi"
	"github.com/stretchr/testify/suite"
)

type InvoicingApiTestSuite struct {
	suite.Suite
	ctx       context.Context
	apiClient *invoicingapi.APIClient
}

func TestInvoicingApiTestSuite(t *testing.T) {
	suite.Run(t, new(InvoicingApiTestSuite))
}

func (suite *InvoicingApiTestSuite) SetupSuite() {
	// Remove any existing expectations
	TestUtilsImpl{}.resetExpectations()

	// Set configuration
	configuration := invoicingapi.NewConfiguration()
	configuration.Host = "127.0.0.1:1080"
	configuration.Scheme = "http"

	// Set the context background
	suite.ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.ctx = context.WithValue(suite.ctx, "serverIndex", nil)

	// New ApiClient
	suite.apiClient = invoicingapi.NewAPIClient(configuration)
}

func (suite *InvoicingApiTestSuite) TearDownTest() {
	TestUtilsImpl{}.resetExpectations()
}

func (suite *InvoicingApiTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *InvoicingApiTestSuite) TestGetInvoices() {
	// generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("invoicingapi/invoices_get", "./payloads")

	// extract expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// fetch query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	// extracting into variables
	number := fmt.Sprintf("%v", qpMap["number"])
	status := fmt.Sprintf("%v", qpMap["status"])
	sentOnFrom, _ := time.Parse(time.RFC3339, fmt.Sprintf("%v", qpMap["sentOnFrom"]))
	sentOnTo, _ := time.Parse(time.RFC3339, fmt.Sprintf("%v", qpMap["sentOnTo"]))
	limit, _ := strconv.ParseInt(fmt.Sprintf("%v", qpMap["limit"]), 10, 32)
	offset, _ := strconv.ParseInt(fmt.Sprintf("%v", qpMap["offset"]), 10, 32)
	sortField := fmt.Sprintf("%v", qpMap["sortField"])
	sortDirection := fmt.Sprintf("%v", qpMap["sortDirection"])

	// execution
	result, r, err := suite.apiClient.InvoicesAPI.InvoicesGet(context.Background()).Number(number).Status(status).SentOnFrom(sentOnFrom).SentOnTo(sentOnTo).Limit(int32(limit)).Offset(int32(offset)).SortField(sortField).SortDirection(sortDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling 'InvoicesAPI.InvoicesGet': %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *InvoicingApiTestSuite) TestGetInvoiceById() {
	// generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("invoicingapi/invoices_get_by_id", "./payloads")

	// extract expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// extracting path parameter
	invoiceId := request.PathParameters["id"][0]

	// execution
	result, r, err := suite.apiClient.InvoicesAPI.InvoicesInvoiceIdGet(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling 'InvoicesAPI.InvoicesInvoiceIdGet': %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *InvoicingApiTestSuite) TestInvoicePay() {
	// generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("invoicingapi/invoices_pay_post", "./payloads")

	// extract expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// extracting path parameter
	invoiceId := request.PathParameters["id"][0]

	// extracting request body
	body, _ := json.Marshal(request.Body.Json)
	var invoicePaymentRequest map[string]interface{}
	json.Unmarshal(body, &invoicePaymentRequest)

	// execution
	result, r, err := suite.apiClient.InvoicesAPI.InvoicesInvoiceIdPayPost(context.Background(), invoiceId).Body(invoicePaymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling 'InvoicesAPI.InvoicesInvoiceIdPayPost': %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *InvoicingApiTestSuite) TestInvoicesGeneratePdf() {
	// generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("invoicingapi/invoices_generate_pdf_post", "./payloads")

	// extract expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extracting path parameter
	invoiceId := request.PathParameters["id"][0]

	// execution
	_, r, err := suite.apiClient.InvoicesAPI.InvoicesInvoiceIdGeneratePdfPost(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling 'InvoicesAPI.InvoiceIdGeneratePdfPost': %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// we cannot assert that the response is equivalent, since this endpoint is meant to return a binary file
	// in this case, a PDF file. to mock this, we'd likely need a binary string that represents a PDF file, and
	// compare the response to a manually-constructed file that represents said binary string.
	// TODO: discuss whether we want to do this?

	// verify
	suite.verifyCalledOnce(expectationId)
}
