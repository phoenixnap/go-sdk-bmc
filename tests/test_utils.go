package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type TestUtilsImpl struct {
}

type MyHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type MyApplicationClient struct {
	HttpClient MyHttpClient
}

func NewApplicationClient(httpClient MyHttpClient) *MyApplicationClient {
	return &MyApplicationClient{
		HttpClient: httpClient,
	}
}

func (t TestUtilsImpl) setupExpectation(requestToMock Request, responseToGet Response, timesParam int) string {

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

	req, err := http.NewRequest(http.MethodPut, "http://localhost:1080/expectation", bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatal(string(body))
	}

	defer resp.Body.Close()

	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bufferString := string(buffer)
	var result []map[string]interface{}

	json.Unmarshal([]byte(bufferString), &result)

	return result[0]["id"].(string)
}

func (t TestUtilsImpl) verifyExpectationMatchedTimes(expectationId string, timesIn int) *http.Response {

	type Times struct {
		AtLeast int `json:"atMost"`
		AtMost  int `json:"atLeast"`
	}
	type ResponseBody struct {
		ExpectationId ExpectationId `json:"expectationId"`
		Times         Times         `json:"times"`
	}

	body := ResponseBody{
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

func (t TestUtilsImpl) resetExpectations() {

	req, err := http.NewRequest(http.MethodPut, "http://localhost:1080/mockserver/reset", http.NoBody)

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

	file, err := ioutil.ReadFile(payloadsPath + "/" + filename + ".json")
	if err != nil {
		log.Fatal(err)
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

func (t TestUtilsImpl) extractRequestBody(request Request) ([]byte) {
	byteData, _ := json.Marshal(request.Body.Json)

	return byteData
}