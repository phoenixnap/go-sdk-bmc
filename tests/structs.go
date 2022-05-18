package tests

import (

)

type Request struct {
	Method         string `json:"method"`
	Path           string `json:"path"`
	PathParameters []PathParameter `json:"pathParameters"`
	QueryStringParameters []QueryParameter `json:"queryStringParameters"`
}

type Response struct {
	StatusCode int `json:"statusCode"`
	Body       []interface{} `json:"body"`
}

type PathParameter struct {
	Id []string `json:"id"`
}

type QueryParameter struct {
	Name   string      `json:"name"`
	Values []string `json:"values"`
}

type Payload struct {
	Request Request `json:"request"`
	Response Response `json:"response"`
}

func(r *Request) setPathParams() {
	r.PathParameters = make([]PathParameter,0)
}