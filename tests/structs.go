package tests

type Request struct {
	Method                string           `json:"method"`
	Path                  string           `json:"path"`
	PathParameters        PathParameters   `json:"pathParameters,omitempty"`
	QueryStringParameters []QueryParameter `json:"queryStringParameters,omitempty"`
	Body                  *RequestBody     `json:"body,omitempty"`
}

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

type PathParameters map[string][]string

type QueryParameter struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type RequestBody struct {
	Type      string      `json:"type"`
	Json      interface{} `json:"json"`
	MatchType string      `json:"matchType"`
}

type Payload struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

type Times struct {
	RemainingTimes int  `json:"remainingTimes"`
	Unlimited      bool `json:"unlimited"`
}

type Body struct {
	HttpRequest  Request  `json:"httpRequest"`
	HttpResponse Response `json:"httpResponse"`
	Times        Times    `json:"times"`
}

type ResponseBody struct {
	HttpRequest  Request  `json:"httpRequest"`
	HttpResponse Response `json:"httpResponse"`
	Times        Times    `json:"times"`
	Id           string   `json:"id"`
}

type ExpectationId struct {
	Id string `json:"id"`
}
