package tests

type Request struct {
	Method                string           `json:"method"`
	Path                  string           `json:"path"`
	PathParameters        []PathParameter  `json:"pathParameters"`
	QueryStringParameters []QueryParameter `json:"queryStringParameters"`
}

type Response struct {
	StatusCode int           `json:"statusCode"`
	Body       []interface{} `json:"body"`
}

type PathParameter struct {
	Id []string `json:"id"`
}

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

func (r *Request) setPathParams(size int) {
	r.PathParameters = make([]PathParameter, size)
}
