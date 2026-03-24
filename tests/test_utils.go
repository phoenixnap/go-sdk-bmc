package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type TestUtilsImpl struct {
}

// mock http client
type MyHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// mock application
type MyApplicationClient struct {
	HttpClient MyHttpClient
}

func NewApplicationClient(httpClient MyHttpClient) *MyApplicationClient {
	return &MyApplicationClient{
		HttpClient: httpClient,
	}
}

// sets up an expectation using the mockserver. the request must be matched *exactly*.
func (t TestUtilsImpl) setupExpectation(requestToMock Request, responseToGet Response, timesParam int) string {
	// setup mockserver request body
	body := Body{
		HttpRequest:  requestToMock,
		HttpResponse: responseToGet,
		Times: Times{
			RemainingTimes: timesParam,
			Unlimited:      false,
		},
	}

	client := &http.Client{}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	// setup the request itself
	req, err := http.NewRequest(http.MethodPut, "http://localhost:1080/expectation", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	// dispatch request and ensure successful return
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		log.Fatal(string(body))
	}

	defer resp.Body.Close()

	// read the contents of the response into a byte-buffer
	buffer, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert the byte-buffer to string
	bufferString := string(buffer)
	var result []map[string]interface{}

	// unmarshal result into a more user-friendly structure
	json.Unmarshal([]byte(bufferString), &result)

	return result[0]["id"].(string)
}

func (t TestUtilsImpl) verifyExpectationMatchedTimes(expectationId string, timesIn int) *http.Response {

	type Times struct {
		AtLeast int `json:"atMost"`
		AtMost  int `json:"atLeast"`
	}

	type Body struct {
		ExpectationId ExpectationId `json:"expectationId"`
		Times         Times         `json:"times"`
	}

	body := Body{
		ExpectationId: ExpectationId{Id: expectationId},
		Times: Times{
			AtLeast: timesIn,
			AtMost:  timesIn,
		},
	}

	client := &http.Client{}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPut, "http://localhost:1080/verify", bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return resp

}

var readyForTesting = false

func (t TestUtilsImpl) resetExpectations() {

	var req *http.Request
	var err error

	if !readyForTesting {
		req, err = http.NewRequest(http.MethodPut, "http://localhost:1080/mockserver/reset", http.NoBody)
		readyForTesting = true
	} else {
		req, err = http.NewRequest(http.MethodPut, "http://localhost:1080/mockserver/clear", http.NoBody)
	}

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}

func (t TestUtilsImpl) generatePayloadsFrom(filename string, payloadsPath string) (Request, Response) {

	if payloadsPath == "" {
		payloadsPath = "./payloads"
	}

	var payload Payload

	file, err := os.ReadFile(payloadsPath + "/" + filename + ".json")
	if err != nil {
		errMsg, _ := fmt.Printf("error when generating payloads: %s", err)
		log.Fatal(errMsg)
	}

	json.Unmarshal(file, &payload)

	return payload.Request, payload.Response
}

func (t TestUtilsImpl) generateQueryParams(request Request) map[string]interface{} {
	type QueryStringParameter struct {
		name   string
		values []string
	}

	qplist := request.QueryStringParameters

	elementMap := make(map[string]interface{})
	for i := 0; i < len(qplist); i++ {
		elementMap[qplist[i].Name] = qplist[i].Values[0]
	}

	return elementMap
}
