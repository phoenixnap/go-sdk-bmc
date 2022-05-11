package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TestUtilsImpl struct {
}

func (t TestUtilsImpl) setup_expectation(requestToMock Request, responseToGet Response, times int) *http.Response {

	type Times struct {
		remainingTimes int
		unlimited      bool
	}

	type Body struct {
		httpRequest  Request
		httpResponse Response
		times        Times
	}

	url := "http://localhost:1080/expectation"

	body := Body{
		httpRequest:  requestToMock,
		httpResponse: responseToGet,
		times: Times{
			remainingTimes: times,
			unlimited:      false,
		},
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

func (t TestUtilsImpl) verify_expectation_matched_times(expectationId string, times int) *http.Response {

	url := "http://localhost:1080/verify"

	type ExpectationId struct {
		id string
	}

	type Times struct {
		atLeast int
		atMost  int
	}

	type Body struct {
		expectationId ExpectationId
		times         Times
	}

	body := Body{
		expectationId: ExpectationId{
			id: "",
		},
		times: Times{
			atLeast: times,
			atMost:  times,
		},
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

func (t TestUtilsImpl) generate_payloads_from(filename string, payloadsPath string) (request Request, response Response) {

	if payloadsPath == "" {
		payloadsPath = "./payloads"
	}

	var payload Payload

	file, err := ioutil.ReadFile("#{payloadsPath}/#{filename}.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &payload)

	return payload.request, payload.response

}

func (t TestUtilsImpl) generate_query_params(request Request) Opts {

	opts := Opts{}

	opts.Opt = make(map[string]string)

	qplist := request.queryStringParameters
	for _, qp := range qplist {
		values := qp.values[0]
		opts.Opt[qp.name] = values
	}

	return opts
}

func (t TestUtilsImpl) extract_id_from(request *Request, symbol *string) (id string) {
	return request.pathParameters.id[0]
}

func (t TestUtilsImpl) extract_request_body(request *Request) (json Json) {
	return request.body.json
}
