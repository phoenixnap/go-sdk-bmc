package tests

type Request struct {
	Method                string           `json:"method"`
	Path                  string           `json:"path"`
	PathParameters        PathParameters   `json:"pathParameters.omitempty"`
	QueryStringParameters []QueryParameter `json:"queryStringParameters.omitempty"`
	Body                  *ResponseBody    `json:"body.omitempty"`
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

func NewPayload() Payload {
	return Payload{
		Request:  NewRequest(),
		Response: NewResponse(),
	}
}

func NewRequest() Request {
	return Request{
		Method:                "",
		Path:                  "",
		PathParameters:        PathParameters{},
		QueryStringParameters: make([]QueryParameter, 0),
	}
}

func NewResponse() Response {
	return Response{
		StatusCode: 0,
		Body:       make([]interface{}, 0),
	}
}
