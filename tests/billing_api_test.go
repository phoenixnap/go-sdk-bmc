package tests

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"

	"fmt"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/suite"
)

type BillingApiTestSuite struct {
	suite.Suite
	Client *billingapi.APIClient
	Ctx    context.Context
}

func (suite *BillingApiTestSuite) SetupSuite() {
	// Remove any existing expectations
	TestUtilsImpl{}.resetExpectations()

	// Set configuration
	configuration := billingapi.NewConfiguration()
	configuration.Host = "127.0.0.1:1080"
	configuration.Scheme = "http"

	// Set the context background
	suite.Ctx = context.WithValue(context.Background(), "accessToken", "ACCESSTOKENSTRING")
	suite.Ctx = context.WithValue(suite.Ctx, "serverIndex", nil)

	// New ApiClient
	suite.Client = billingapi.NewAPIClient(configuration)
}

func (suite *BillingApiTestSuite) TearDownTest() {
	TestUtilsImpl{}.resetExpectations()
}

func (suite *BillingApiTestSuite) verifyCalledOnce(expectationId string) {
	// Result retrieved from server's verification
	// Verifying expectation matched exactly once.
	verifyResult := TestUtilsImpl{}.verifyExpectationMatchedTimes(expectationId, 1)

	// Asserts a successful result
	suite.Equal(202, verifyResult.StatusCode)
}

func (suite *BillingApiTestSuite) TestGetAccountBillingConfigurationsMe() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("billingapi/billing_configurations/account_billing_configurations_me_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Operation Execution
	result, _, _ := suite.Client.BillingConfigurationsAPI.AccountBillingConfigurationMeGet(context.Background()).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BillingApiTestSuite) TestGetProductAvailability() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("billingapi/products/product_availability_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	minQuantityStr, ok := qpMap["minQuantity"].(string)
	if !ok {
		panic(ok)
	}

	productCategory := []string{fmt.Sprintf("%v", qpMap["productCategory"])}
	productCode := []string{fmt.Sprintf("%v", qpMap["productCode"])}
	showOnlyMinQuantityAvailable, _ := strconv.ParseBool(fmt.Sprintf("%v", qpMap["showOnlyMinQuantityAvailable"]))
	location := []billingapi.LocationEnum{billingapi.LocationEnum(fmt.Sprintf("%v", qpMap["location"]))}
	solution := []string{fmt.Sprintf("%v", qpMap["solution"])}
	minQuantity64, _ := strconv.ParseFloat(minQuantityStr, 32)
	minQuantity := float32(minQuantity64)

	// Operation Execution
	result, _, _ := suite.Client.ProductsAPI.ProductAvailabilityGet(context.Background()).ProductCategory(productCategory).ProductCode(productCode).ShowOnlyMinQuantityAvailable(showOnlyMinQuantityAvailable).Location(location).Solution(solution).MinQuantity(minQuantity).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BillingApiTestSuite) TestGetProducts() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("billingapi/products/products_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	productCode := fmt.Sprintf("%v", qpMap["productCode"])
	productCategory := fmt.Sprintf("%v", qpMap["productCategory"])
	skuCode := fmt.Sprintf("%v", qpMap["skuCode"])
	location := fmt.Sprintf("%v", qpMap["location"])

	// Operation Execution
	result, _, _ := suite.Client.ProductsAPI.ProductsGet(context.Background()).ProductCode(productCode).ProductCategory(productCategory).SkuCode(skuCode).Location(location).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BillingApiTestSuite) TestGetRatedUsage() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("billingapi/rated_usage/rated_usage_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	// Converting to int32
	fromYearMonth := fmt.Sprintf("%v", qpMap["fromYearMonth"])
	toYearMonth := fmt.Sprintf("%v", qpMap["toYearMonth"])
	productCategory := billingapi.ProductCategoryEnum(fmt.Sprintf("%v", qpMap["productCategory"]))

	// Operation Execution
	result, _, _ := suite.Client.RatedUsageAPI.RatedUsageGet(context.Background()).FromYearMonth(fromYearMonth).ToYearMonth(toYearMonth).ProductCategory(productCategory).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BillingApiTestSuite) TestGetRatedUsageMonthToDate() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("billingapi/rated_usage/rated_usage_month_to_date_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	productCategory := billingapi.ProductCategoryEnum(fmt.Sprintf("%v", qpMap["productCategory"]))

	// Operation Execution
	result, _, _ := suite.Client.RatedUsageAPI.RatedUsageMonthToDateGet(context.Background()).ProductCategory(productCategory).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BillingApiTestSuite) TestGetReservations() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("billingapi/reservations/reservations_get", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Fetch a map of query parameters
	qpMap := TestUtilsImpl{}.generateQueryParams(request)

	productCategory := billingapi.ProductCategoryEnum(fmt.Sprintf("%v", qpMap["productCategory"]))

	// Operation Execution
	result, _, _ := suite.Client.ReservationsAPI.ReservationsGet(context.Background()).ProductCategory(productCategory).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BillingApiTestSuite) TestGetReservationsById() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("billingapi/reservations/reservations_get_by_id", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract ID
	reservationId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ReservationsAPI.ReservationsReservationIdGet(context.Background(), reservationId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BillingApiTestSuite) TestCreateReservation() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("billingapi/reservations/reservations_post", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var reservationRequest billingapi.ReservationRequest
	json.Unmarshal(body, &reservationRequest)

	// Operation Execution
	result, _, _ := suite.Client.ReservationsAPI.ReservationsPost(context.Background()).ReservationRequest(reservationRequest).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BillingApiTestSuite) TestConvertReservation() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("billingapi/reservations/reservations_actions_convert", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var reservationRequest billingapi.ReservationRequest
	json.Unmarshal(body, &reservationRequest)

	// Extract ID
	reservationId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ReservationsAPI.ReservationsReservationIdActionsConvertPost(context.Background(), reservationId).ReservationRequest(reservationRequest).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BillingApiTestSuite) TestEnableAutoRenewOnReservation() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("billingapi/reservations/reservations_actions_enable_auto_renew", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Extract ID
	reservationId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ReservationsAPI.ReservationsReservationIdActionsAutoRenewEnablePost(context.Background(), reservationId).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func (suite *BillingApiTestSuite) TestDisableAutoRenewOnReservation() {
	// Generate payload
	request, response := TestUtilsImpl{}.generatePayloadsFrom("billingapi/reservations/reservations_actions_disable_auto_renew", "./payloads")

	// Extract the response expectation id
	expectationId := TestUtilsImpl{}.setupExpectation(request, response, 1)

	// Prepare request body
	body, _ := json.Marshal(request.Body.Json)
	var reservationAutoRenewDisableRequest billingapi.ReservationAutoRenewDisableRequest
	json.Unmarshal(body, &reservationAutoRenewDisableRequest)

	// Extract ID
	reservationId := request.PathParameters["id"][0]

	// Operation Execution
	result, _, _ := suite.Client.ReservationsAPI.ReservationsReservationIdActionsAutoRenewDisablePost(context.Background(), reservationId).ReservationAutoRenewDisableRequest(reservationAutoRenewDisableRequest).Execute()

	// Convert the result and response body to json strings
	jsonResult, _ := json.Marshal(result)
	jsonResponseBody, _ := json.Marshal(response.Body)

	// Asserts
	suite.Equal(string(jsonResult), string(jsonResponseBody))

	// Verify
	suite.verifyCalledOnce(expectationId)
}

func TestBillingApiTestSuite(t *testing.T) {
	suite.Run(t, new(BillingApiTestSuite))
}
