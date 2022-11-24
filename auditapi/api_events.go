/*
Audit Log API

The Audit Logs API lets you read audit log entries and track API calls or activities in the Bare Metal Cloud Portal.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#audit-log-api' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/audit/v1/)</b> 

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)


type EventsApi interface {

	/*
	EventsGet List event logs.

	Retrieves the event logs for given time period. All date & times are in UTC.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiEventsGetRequest
	*/
	EventsGet(ctx context.Context) ApiEventsGetRequest

	// EventsGetExecute executes the request
	//  @return []Event
	EventsGetExecute(r ApiEventsGetRequest) ([]Event, *http.Response, error)
}

// EventsApiService EventsApi service
type EventsApiService service

type ApiEventsGetRequest struct {
	ctx context.Context
	ApiService EventsApi
	from *time.Time
	to *time.Time
	limit *int32
	order *string
	username *string
	verb *string
	uri *string
}

// From the date and time (inclusive) to filter event log records by.
func (r ApiEventsGetRequest) From(from time.Time) ApiEventsGetRequest {
	r.from = &from
	return r
}

// To the date and time (inclusive) to filter event log records by.
func (r ApiEventsGetRequest) To(to time.Time) ApiEventsGetRequest {
	r.to = &to
	return r
}

// Limit the number of records returned.
func (r ApiEventsGetRequest) Limit(limit int32) ApiEventsGetRequest {
	r.limit = &limit
	return r
}

// Ordering of the event&#39;s time. SortBy can be introduced later on.
func (r ApiEventsGetRequest) Order(order string) ApiEventsGetRequest {
	r.order = &order
	return r
}

// The username that did the actions.
func (r ApiEventsGetRequest) Username(username string) ApiEventsGetRequest {
	r.username = &username
	return r
}

// The HTTP verb corresponding to the action.
func (r ApiEventsGetRequest) Verb(verb string) ApiEventsGetRequest {
	r.verb = &verb
	return r
}

// The request uri.
func (r ApiEventsGetRequest) Uri(uri string) ApiEventsGetRequest {
	r.uri = &uri
	return r
}

func (r ApiEventsGetRequest) Execute() ([]Event, *http.Response, error) {
	return r.ApiService.EventsGetExecute(r)
}

/*
EventsGet List event logs.

Retrieves the event logs for given time period. All date & times are in UTC.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEventsGetRequest
*/
func (a *EventsApiService) EventsGet(ctx context.Context) ApiEventsGetRequest {
	return ApiEventsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Event
func (a *EventsApiService) EventsGetExecute(r ApiEventsGetRequest) ([]Event, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Event
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsApiService.EventsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
	}
	if r.username != nil {
		localVarQueryParams.Add("username", parameterToString(*r.username, ""))
	}
	if r.verb != nil {
		localVarQueryParams.Add("verb", parameterToString(*r.verb, ""))
	}
	if r.uri != nil {
		localVarQueryParams.Add("uri", parameterToString(*r.uri, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
