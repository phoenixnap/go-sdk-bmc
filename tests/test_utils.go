package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TestUtilsImpl struct {
}

func (t TestUtilsImpl) setup_expectation(requestToMock map[string]interface{}, responseToGet map[string]interface{}, times int) string {

	url := "http://localhost:1080/expectation"

	body := map[string]interface{}{
		"httpRequest":  requestToMock,
		"httpResponse": responseToGet,
		"times": struct {
			remainingTimes int
			unlimited      bool
		}{1, false},
	}

	client := &http.Client{}

	jsonResult, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonResult))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var payload interface{}
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(buffer, &payload)

	m := payload.(map[string]interface{})["id"].(string)

	return m
}

func (t TestUtilsImpl) verify_expectation_matched_times(expectationId string, timesIn int) *http.Response {

	url := "http://localhost:1080/verify"

	body := map[string]interface{}{
		"expectationId": struct {
			id string
		}{expectationId},
		"times": struct {
			atLeast int
			atMost  int
		}{timesIn, timesIn},
	}

	client := &http.Client{}

	json, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp
}

func (t TestUtilsImpl) reset_expectations() {
	url := "http://localhost:1080/reset"

	http.NewRequest(http.MethodPut, url, http.NoBody)

}

func (t TestUtilsImpl) generate_payloads_from(filename string, payloadsPath string) (interface{}, interface{}) {

	if payloadsPath == "" {
		payloadsPath = "./payloads"
	}

	var payload interface{}

	file, err := ioutil.ReadFile("#{payloadsPath}/#{filename}.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &payload)
	m := payload.(map[string]interface{})

	return m["request"], m["response"]

}

func (t TestUtilsImpl) generate_query_params(request map[string]interface{}) context.Context {
	type QueryStringParameter struct {
		name   string
		values []string
	}

	ctx := context.Background()

	qplist := request["queryStringParameters"].([]QueryStringParameter)

	for _, qp := range qplist {
		ctx = context.WithValue(ctx, qp.name, qp.values[0])
	}

	return ctx
}

func (t TestUtilsImpl) extract_id_from(request map[string]interface{}, symbol *string) (id string) {
	return request["pathParameters"].(map[string]interface{})["id"].([]string)[0]
}

func (t TestUtilsImpl) extract_request_body(request map[string]interface{}) map[string]interface{} {
	return request["body"].(map[string]interface{})["json"].(map[string]interface{})
}
