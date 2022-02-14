/*
Networks API

Use the Networks API to create, list, edit, and delete private networks to best fit your business needs. Private networks allow your servers to communicate without connecting to the public internet, avoiding unnecessary egress data charges. </br></br>**Knowledge base articles to help you can be found <a href=' https://phoenixnap.com/kb/bmc-server-management-via-api#multi-private-backend-network-api' target='_blank'>here</a>**</br></br>**All URLs are relative to (https://api.phoenixnap.com/networks/v1/)**

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type PrivateNetworksApi interface {

	/*
		PrivateNetworksGet List Private Networks.

		List all Private Networks owned by account.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return ApiPrivateNetworksGetRequest
	*/
	PrivateNetworksGet(ctx _context.Context) ApiPrivateNetworksGetRequest

	// PrivateNetworksGetExecute executes the request
	//  @return []PrivateNetwork
	PrivateNetworksGetExecute(r ApiPrivateNetworksGetRequest) ([]PrivateNetwork, *_nethttp.Response, error)

	/*
		PrivateNetworksNetworkIdDelete Delete a Private Network.

		Delete Private Network.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param privateNetworkId The private network identifier.
		 @return ApiPrivateNetworksNetworkIdDeleteRequest
	*/
	PrivateNetworksNetworkIdDelete(ctx _context.Context, privateNetworkId string) ApiPrivateNetworksNetworkIdDeleteRequest

	// PrivateNetworksNetworkIdDeleteExecute executes the request
	PrivateNetworksNetworkIdDeleteExecute(r ApiPrivateNetworksNetworkIdDeleteRequest) (*_nethttp.Response, error)

	/*
		PrivateNetworksNetworkIdGet Get a Private Network.

		Retrieve Private Network Details.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param privateNetworkId The private network identifier.
		 @return ApiPrivateNetworksNetworkIdGetRequest
	*/
	PrivateNetworksNetworkIdGet(ctx _context.Context, privateNetworkId string) ApiPrivateNetworksNetworkIdGetRequest

	// PrivateNetworksNetworkIdGetExecute executes the request
	//  @return PrivateNetwork
	PrivateNetworksNetworkIdGetExecute(r ApiPrivateNetworksNetworkIdGetRequest) (PrivateNetwork, *_nethttp.Response, error)

	/*
		PrivateNetworksNetworkIdPut Update a Private Network.

		Update Private Network Details.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param privateNetworkId The private network identifier.
		 @return ApiPrivateNetworksNetworkIdPutRequest
	*/
	PrivateNetworksNetworkIdPut(ctx _context.Context, privateNetworkId string) ApiPrivateNetworksNetworkIdPutRequest

	// PrivateNetworksNetworkIdPutExecute executes the request
	//  @return PrivateNetwork
	PrivateNetworksNetworkIdPutExecute(r ApiPrivateNetworksNetworkIdPutRequest) (PrivateNetwork, *_nethttp.Response, error)

	/*
		PrivateNetworksPost Create a Private Network.

		Create a Private Network.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return ApiPrivateNetworksPostRequest
	*/
	PrivateNetworksPost(ctx _context.Context) ApiPrivateNetworksPostRequest

	// PrivateNetworksPostExecute executes the request
	//  @return PrivateNetwork
	PrivateNetworksPostExecute(r ApiPrivateNetworksPostRequest) (PrivateNetwork, *_nethttp.Response, error)
}

// PrivateNetworksApiService PrivateNetworksApi service
type PrivateNetworksApiService service

type ApiPrivateNetworksGetRequest struct {
	ctx        _context.Context
	ApiService PrivateNetworksApi
	location   *string
}

// If present will filter the result by the given location of the Private Networks.
func (r ApiPrivateNetworksGetRequest) Location(location string) ApiPrivateNetworksGetRequest {
	r.location = &location
	return r
}

func (r ApiPrivateNetworksGetRequest) Execute() ([]PrivateNetwork, *_nethttp.Response, error) {
	return r.ApiService.PrivateNetworksGetExecute(r)
}

/*
PrivateNetworksGet List Private Networks.

List all Private Networks owned by account.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPrivateNetworksGetRequest
*/
func (a *PrivateNetworksApiService) PrivateNetworksGet(ctx _context.Context) ApiPrivateNetworksGetRequest {
	return ApiPrivateNetworksGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []PrivateNetwork
func (a *PrivateNetworksApiService) PrivateNetworksGetExecute(r ApiPrivateNetworksGetRequest) ([]PrivateNetwork, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []PrivateNetwork
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateNetworksApiService.PrivateNetworksGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/private-networks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.location != nil {
		localVarQueryParams.Add("location", parameterToString(*r.location, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPrivateNetworksNetworkIdDeleteRequest struct {
	ctx              _context.Context
	ApiService       PrivateNetworksApi
	privateNetworkId string
}

func (r ApiPrivateNetworksNetworkIdDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PrivateNetworksNetworkIdDeleteExecute(r)
}

/*
PrivateNetworksNetworkIdDelete Delete a Private Network.

Delete Private Network.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param privateNetworkId The private network identifier.
 @return ApiPrivateNetworksNetworkIdDeleteRequest
*/
func (a *PrivateNetworksApiService) PrivateNetworksNetworkIdDelete(ctx _context.Context, privateNetworkId string) ApiPrivateNetworksNetworkIdDeleteRequest {
	return ApiPrivateNetworksNetworkIdDeleteRequest{
		ApiService:       a,
		ctx:              ctx,
		privateNetworkId: privateNetworkId,
	}
}

// Execute executes the request
func (a *PrivateNetworksApiService) PrivateNetworksNetworkIdDeleteExecute(r ApiPrivateNetworksNetworkIdDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateNetworksApiService.PrivateNetworksNetworkIdDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/private-networks/{privateNetworkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkId"+"}", _neturl.PathEscape(parameterToString(r.privateNetworkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPrivateNetworksNetworkIdGetRequest struct {
	ctx              _context.Context
	ApiService       PrivateNetworksApi
	privateNetworkId string
}

func (r ApiPrivateNetworksNetworkIdGetRequest) Execute() (PrivateNetwork, *_nethttp.Response, error) {
	return r.ApiService.PrivateNetworksNetworkIdGetExecute(r)
}

/*
PrivateNetworksNetworkIdGet Get a Private Network.

Retrieve Private Network Details.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param privateNetworkId The private network identifier.
 @return ApiPrivateNetworksNetworkIdGetRequest
*/
func (a *PrivateNetworksApiService) PrivateNetworksNetworkIdGet(ctx _context.Context, privateNetworkId string) ApiPrivateNetworksNetworkIdGetRequest {
	return ApiPrivateNetworksNetworkIdGetRequest{
		ApiService:       a,
		ctx:              ctx,
		privateNetworkId: privateNetworkId,
	}
}

// Execute executes the request
//  @return PrivateNetwork
func (a *PrivateNetworksApiService) PrivateNetworksNetworkIdGetExecute(r ApiPrivateNetworksNetworkIdGetRequest) (PrivateNetwork, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PrivateNetwork
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateNetworksApiService.PrivateNetworksNetworkIdGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/private-networks/{privateNetworkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkId"+"}", _neturl.PathEscape(parameterToString(r.privateNetworkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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
		if localVarHTTPResponse.StatusCode == 404 {
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPrivateNetworksNetworkIdPutRequest struct {
	ctx                  _context.Context
	ApiService           PrivateNetworksApi
	privateNetworkId     string
	privateNetworkModify *PrivateNetworkModify
}

func (r ApiPrivateNetworksNetworkIdPutRequest) PrivateNetworkModify(privateNetworkModify PrivateNetworkModify) ApiPrivateNetworksNetworkIdPutRequest {
	r.privateNetworkModify = &privateNetworkModify
	return r
}

func (r ApiPrivateNetworksNetworkIdPutRequest) Execute() (PrivateNetwork, *_nethttp.Response, error) {
	return r.ApiService.PrivateNetworksNetworkIdPutExecute(r)
}

/*
PrivateNetworksNetworkIdPut Update a Private Network.

Update Private Network Details.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param privateNetworkId The private network identifier.
 @return ApiPrivateNetworksNetworkIdPutRequest
*/
func (a *PrivateNetworksApiService) PrivateNetworksNetworkIdPut(ctx _context.Context, privateNetworkId string) ApiPrivateNetworksNetworkIdPutRequest {
	return ApiPrivateNetworksNetworkIdPutRequest{
		ApiService:       a,
		ctx:              ctx,
		privateNetworkId: privateNetworkId,
	}
}

// Execute executes the request
//  @return PrivateNetwork
func (a *PrivateNetworksApiService) PrivateNetworksNetworkIdPutExecute(r ApiPrivateNetworksNetworkIdPutRequest) (PrivateNetwork, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PrivateNetwork
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateNetworksApiService.PrivateNetworksNetworkIdPut")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/private-networks/{privateNetworkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkId"+"}", _neturl.PathEscape(parameterToString(r.privateNetworkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.privateNetworkModify
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		if localVarHTTPResponse.StatusCode == 404 {
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPrivateNetworksPostRequest struct {
	ctx                  _context.Context
	ApiService           PrivateNetworksApi
	privateNetworkCreate *PrivateNetworkCreate
}

func (r ApiPrivateNetworksPostRequest) PrivateNetworkCreate(privateNetworkCreate PrivateNetworkCreate) ApiPrivateNetworksPostRequest {
	r.privateNetworkCreate = &privateNetworkCreate
	return r
}

func (r ApiPrivateNetworksPostRequest) Execute() (PrivateNetwork, *_nethttp.Response, error) {
	return r.ApiService.PrivateNetworksPostExecute(r)
}

/*
PrivateNetworksPost Create a Private Network.

Create a Private Network.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPrivateNetworksPostRequest
*/
func (a *PrivateNetworksApiService) PrivateNetworksPost(ctx _context.Context) ApiPrivateNetworksPostRequest {
	return ApiPrivateNetworksPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return PrivateNetwork
func (a *PrivateNetworksApiService) PrivateNetworksPostExecute(r ApiPrivateNetworksPostRequest) (PrivateNetwork, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PrivateNetwork
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateNetworksApiService.PrivateNetworksPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/private-networks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.privateNetworkCreate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
