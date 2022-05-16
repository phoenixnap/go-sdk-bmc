/*
Networks API

Create, list, edit and delete public/private networks with the Network API. Use public networks to place multiple  servers on the same network or VLAN. Assign new servers with IP addresses from the same CIDR range. Use private  networks to avoid unnecessary egress data charges. Model your networks according to your business needs.<br> <br> <span class='pnap-api-knowledge-base-link'> Helpful knowledge base articles are available for  <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#multi-private-backend-network-api' target='_blank'>multi-private backend networks</a> and <a href='https://phoenixnap.com/kb/bmc-server-management-via-api#ftoc-heading-15' target='_blank'>public networks</a>. </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/networks/v1/)</b> 

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type PrivateNetworksApi interface {

	/*
	PrivateNetworksGet List Private Networks.

	List all Private Networks owned by account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPrivateNetworksGetRequest
	*/
	PrivateNetworksGet(ctx context.Context) ApiPrivateNetworksGetRequest

	// PrivateNetworksGetExecute executes the request
	//  @return []PrivateNetwork
	PrivateNetworksGetExecute(r ApiPrivateNetworksGetRequest) ([]PrivateNetwork, *http.Response, error)

	/*
	PrivateNetworksNetworkIdDelete Delete a Private Network.

	Delete Private Network.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param privateNetworkId The private network identifier.
	@return ApiPrivateNetworksNetworkIdDeleteRequest
	*/
	PrivateNetworksNetworkIdDelete(ctx context.Context, privateNetworkId string) ApiPrivateNetworksNetworkIdDeleteRequest

	// PrivateNetworksNetworkIdDeleteExecute executes the request
	PrivateNetworksNetworkIdDeleteExecute(r ApiPrivateNetworksNetworkIdDeleteRequest) (*http.Response, error)

	/*
	PrivateNetworksNetworkIdGet Get a Private Network.

	Retrieve Private Network Details.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param privateNetworkId The private network identifier.
	@return ApiPrivateNetworksNetworkIdGetRequest
	*/
	PrivateNetworksNetworkIdGet(ctx context.Context, privateNetworkId string) ApiPrivateNetworksNetworkIdGetRequest

	// PrivateNetworksNetworkIdGetExecute executes the request
	//  @return PrivateNetwork
	PrivateNetworksNetworkIdGetExecute(r ApiPrivateNetworksNetworkIdGetRequest) (*PrivateNetwork, *http.Response, error)

	/*
	PrivateNetworksNetworkIdPut Update a Private Network.

	Update Private Network Details.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param privateNetworkId The private network identifier.
	@return ApiPrivateNetworksNetworkIdPutRequest
	*/
	PrivateNetworksNetworkIdPut(ctx context.Context, privateNetworkId string) ApiPrivateNetworksNetworkIdPutRequest

	// PrivateNetworksNetworkIdPutExecute executes the request
	//  @return PrivateNetwork
	PrivateNetworksNetworkIdPutExecute(r ApiPrivateNetworksNetworkIdPutRequest) (*PrivateNetwork, *http.Response, error)

	/*
	PrivateNetworksPost Create a Private Network.

	Create a Private Network.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPrivateNetworksPostRequest
	*/
	PrivateNetworksPost(ctx context.Context) ApiPrivateNetworksPostRequest

	// PrivateNetworksPostExecute executes the request
	//  @return PrivateNetwork
	PrivateNetworksPostExecute(r ApiPrivateNetworksPostRequest) (*PrivateNetwork, *http.Response, error)
}

// PrivateNetworksApiService PrivateNetworksApi service
type PrivateNetworksApiService service

type ApiPrivateNetworksGetRequest struct {
	ctx context.Context
	ApiService PrivateNetworksApi
	location *string
}

// If present will filter the result by the given location of the Private Networks.
func (r ApiPrivateNetworksGetRequest) Location(location string) ApiPrivateNetworksGetRequest {
	r.location = &location
	return r
}

func (r ApiPrivateNetworksGetRequest) Execute() ([]PrivateNetwork, *http.Response, error) {
	return r.ApiService.PrivateNetworksGetExecute(r)
}

/*
PrivateNetworksGet List Private Networks.

List all Private Networks owned by account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPrivateNetworksGetRequest
*/
func (a *PrivateNetworksApiService) PrivateNetworksGet(ctx context.Context) ApiPrivateNetworksGetRequest {
	return ApiPrivateNetworksGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []PrivateNetwork
func (a *PrivateNetworksApiService) PrivateNetworksGetExecute(r ApiPrivateNetworksGetRequest) ([]PrivateNetwork, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PrivateNetwork
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateNetworksApiService.PrivateNetworksGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/private-networks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
    
    localVarHeaderParams["X-Powered-By"] = "PNAP-go-sdk-bmc/" + SdkVersion;
    localVarHeaderParams["User-Agent"] = "PNAP-go-sdk-bmc/" + SdkVersion;

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

type ApiPrivateNetworksNetworkIdDeleteRequest struct {
	ctx context.Context
	ApiService PrivateNetworksApi
	privateNetworkId string
}

func (r ApiPrivateNetworksNetworkIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.PrivateNetworksNetworkIdDeleteExecute(r)
}

/*
PrivateNetworksNetworkIdDelete Delete a Private Network.

Delete Private Network.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param privateNetworkId The private network identifier.
 @return ApiPrivateNetworksNetworkIdDeleteRequest
*/
func (a *PrivateNetworksApiService) PrivateNetworksNetworkIdDelete(ctx context.Context, privateNetworkId string) ApiPrivateNetworksNetworkIdDeleteRequest {
	return ApiPrivateNetworksNetworkIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		privateNetworkId: privateNetworkId,
	}
}

// Execute executes the request
func (a *PrivateNetworksApiService) PrivateNetworksNetworkIdDeleteExecute(r ApiPrivateNetworksNetworkIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateNetworksApiService.PrivateNetworksNetworkIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/private-networks/{privateNetworkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkId"+"}", url.PathEscape(parameterToString(r.privateNetworkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}
    
    localVarHeaderParams["X-Powered-By"] = "PNAP-go-sdk-bmc/" + SdkVersion;
    localVarHeaderParams["User-Agent"] = "PNAP-go-sdk-bmc/" + SdkVersion;

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
	ctx context.Context
	ApiService PrivateNetworksApi
	privateNetworkId string
}

func (r ApiPrivateNetworksNetworkIdGetRequest) Execute() (*PrivateNetwork, *http.Response, error) {
	return r.ApiService.PrivateNetworksNetworkIdGetExecute(r)
}

/*
PrivateNetworksNetworkIdGet Get a Private Network.

Retrieve Private Network Details.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param privateNetworkId The private network identifier.
 @return ApiPrivateNetworksNetworkIdGetRequest
*/
func (a *PrivateNetworksApiService) PrivateNetworksNetworkIdGet(ctx context.Context, privateNetworkId string) ApiPrivateNetworksNetworkIdGetRequest {
	return ApiPrivateNetworksNetworkIdGetRequest{
		ApiService: a,
		ctx: ctx,
		privateNetworkId: privateNetworkId,
	}
}

// Execute executes the request
//  @return PrivateNetwork
func (a *PrivateNetworksApiService) PrivateNetworksNetworkIdGetExecute(r ApiPrivateNetworksNetworkIdGetRequest) (*PrivateNetwork, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PrivateNetwork
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateNetworksApiService.PrivateNetworksNetworkIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/private-networks/{privateNetworkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkId"+"}", url.PathEscape(parameterToString(r.privateNetworkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}
    
    localVarHeaderParams["X-Powered-By"] = "PNAP-go-sdk-bmc/" + SdkVersion;
    localVarHeaderParams["User-Agent"] = "PNAP-go-sdk-bmc/" + SdkVersion;

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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPrivateNetworksNetworkIdPutRequest struct {
	ctx context.Context
	ApiService PrivateNetworksApi
	privateNetworkId string
	privateNetworkModify *PrivateNetworkModify
}

func (r ApiPrivateNetworksNetworkIdPutRequest) PrivateNetworkModify(privateNetworkModify PrivateNetworkModify) ApiPrivateNetworksNetworkIdPutRequest {
	r.privateNetworkModify = &privateNetworkModify
	return r
}

func (r ApiPrivateNetworksNetworkIdPutRequest) Execute() (*PrivateNetwork, *http.Response, error) {
	return r.ApiService.PrivateNetworksNetworkIdPutExecute(r)
}

/*
PrivateNetworksNetworkIdPut Update a Private Network.

Update Private Network Details.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param privateNetworkId The private network identifier.
 @return ApiPrivateNetworksNetworkIdPutRequest
*/
func (a *PrivateNetworksApiService) PrivateNetworksNetworkIdPut(ctx context.Context, privateNetworkId string) ApiPrivateNetworksNetworkIdPutRequest {
	return ApiPrivateNetworksNetworkIdPutRequest{
		ApiService: a,
		ctx: ctx,
		privateNetworkId: privateNetworkId,
	}
}

// Execute executes the request
//  @return PrivateNetwork
func (a *PrivateNetworksApiService) PrivateNetworksNetworkIdPutExecute(r ApiPrivateNetworksNetworkIdPutRequest) (*PrivateNetwork, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PrivateNetwork
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateNetworksApiService.PrivateNetworksNetworkIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/private-networks/{privateNetworkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"privateNetworkId"+"}", url.PathEscape(parameterToString(r.privateNetworkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}
    
    localVarHeaderParams["X-Powered-By"] = "PNAP-go-sdk-bmc/" + SdkVersion;
    localVarHeaderParams["User-Agent"] = "PNAP-go-sdk-bmc/" + SdkVersion;

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.privateNetworkModify
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPrivateNetworksPostRequest struct {
	ctx context.Context
	ApiService PrivateNetworksApi
	privateNetworkCreate *PrivateNetworkCreate
}

func (r ApiPrivateNetworksPostRequest) PrivateNetworkCreate(privateNetworkCreate PrivateNetworkCreate) ApiPrivateNetworksPostRequest {
	r.privateNetworkCreate = &privateNetworkCreate
	return r
}

func (r ApiPrivateNetworksPostRequest) Execute() (*PrivateNetwork, *http.Response, error) {
	return r.ApiService.PrivateNetworksPostExecute(r)
}

/*
PrivateNetworksPost Create a Private Network.

Create a Private Network.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPrivateNetworksPostRequest
*/
func (a *PrivateNetworksApiService) PrivateNetworksPost(ctx context.Context) ApiPrivateNetworksPostRequest {
	return ApiPrivateNetworksPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PrivateNetwork
func (a *PrivateNetworksApiService) PrivateNetworksPostExecute(r ApiPrivateNetworksPostRequest) (*PrivateNetwork, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PrivateNetwork
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateNetworksApiService.PrivateNetworksPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/private-networks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}
    
    localVarHeaderParams["X-Powered-By"] = "PNAP-go-sdk-bmc/" + SdkVersion;
    localVarHeaderParams["User-Agent"] = "PNAP-go-sdk-bmc/" + SdkVersion;

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.privateNetworkCreate
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
