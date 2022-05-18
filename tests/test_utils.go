package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"fmt"
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

func (t TestUtilsImpl) setup_expectation(requestToMock Request, responseToGet Response, timesParam int) string {
	servAddr := "localhost:1080"
	tcpAddr, _ := net.ResolveTCPAddr("tcp", servAddr)

	type Times struct{
		RemainingTimes int `json:"remainingTimes"`
		Unlimited bool `json:"unlimited"`
	}

	type Body struct {
		HttpRequest Request `json:"httpRequest"`
		HttpResponse Response `json:"httpResponse"`
		Times Times `json:"times"`
	}

	type ResponseBody struct {
		HttpRequest Request `json:"httpRequest"`
		HttpResponse Response `json:"httpResponse"`
		Times Times `json:"times"`
		Id string `json:"id"`
	}

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	body := Body{
		HttpRequest: requestToMock,
		HttpResponse: responseToGet,
		Times: Times{
			RemainingTimes: timesParam,
			Unlimited: false,
		},
	}

	client := &http.Client{}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	

	req, err := http.NewRequest(http.MethodPut, "http://localhost:1080/expectation", bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	conn.Close()
	sb:=string(buffer)
	var result []map[string]interface{}

	json.Unmarshal([]byte(sb), &result)

	return result[0]["id"].(string)
}

func (t TestUtilsImpl) verify_expectation_matched_times(expectationId string, timesIn int) *http.Response {
	servAddr := "localhost:1080"
	tcpAddr, _ := net.ResolveTCPAddr("tcp", servAddr)
	
	type ExpectationId struct{
		Id string `json:"id"`
	}
	type Times struct {
		AtLeast int `json:"atMost"`
		AtMost int	`json:"atLeast"`
	}
	type Body struct{
		ExpectationId ExpectationId `json:"expectationId"`
		Times Times `json:"times"`
	}
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	body := Body{
		ExpectationId: ExpectationId{Id:expectationId},
		Times: Times{
			AtLeast: timesIn,
			AtMost: timesIn,
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
	conn.Close()
	defer resp.Body.Close()

	return resp
}

func (t TestUtilsImpl) reset_expectations() {
	url := "http://localhost:1080/mockserver/reset"

	req, err :=http.NewRequest(http.MethodPut, url, http.NoBody)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(resp)
}

func (t TestUtilsImpl) generate_payloads_from(filename string, payloadsPath string) (Request, Response) {

	if payloadsPath == "" {
		payloadsPath = "./payloads"
	}

	var payload Payload

	file, err := ioutil.ReadFile(payloadsPath + "/" + filename + ".json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &payload)

	return payload.Request, payload.Response
}

func (t TestUtilsImpl) generate_query_params(request Request) map[string]interface{} {
	type QueryStringParameter struct {
		name   string
		values []string
	}

	qplist := request.QueryStringParameters

	
	elementMap := make(map[string]interface{})
	for i := 0; i < len(qplist); i ++ {
		elementMap[qplist[i].Name] = qplist[i].Values[0]
	}

	return elementMap
}

func (t TestUtilsImpl) extract_id_from(request map[string]interface{}, symbol *string) (id string) {
	return request["pathParameters"].(map[string]interface{})["id"].([]string)[0]
}

func (t TestUtilsImpl) extract_request_body(request map[string]interface{}) map[string]interface{} {
	return request["body"].(map[string]interface{})["json"].(map[string]interface{})
}
