package tests


type QueryStringParameter struct {
	name   string
	values []string
}
type ID struct {
	id []string
}

type UserInfo struct {
	accountId string
	clientId  string
	username  string
}

type ResponseBody struct {
	name      string
	timestamp string
	userInfo  UserInfo
}

type Response struct {
	statusCode int
	body       []ResponseBody
}
type Payload struct {
	request  Request
	response Response
}

type Json struct {
	limit  int
	reason string
}

type RequestBody struct {
	_type     string
	json      Json
	matchType string
}

type Request struct {
	method                string
	path                  string
	queryStringParameters []QueryStringParameter
	pathParameters        ID
	body                  RequestBody
}

type Opts struct {
	Opt map[string]string
}
