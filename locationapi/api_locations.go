/*
Locations API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package locationapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

type LocationsApi interface {

	/*
		GetLocations Get All Locations

		Retrieve the locations info.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetLocationsRequest
	*/
	GetLocations(ctx context.Context) ApiGetLocationsRequest

	// GetLocationsExecute executes the request
	//  @return []Location
	GetLocationsExecute(r ApiGetLocationsRequest) ([]Location, *http.Response, error)
}

// LocationsApiService LocationsApi service
type LocationsApiService service

type ApiGetLocationsRequest struct {
	ctx             context.Context
	ApiService      LocationsApi
	location        *LocationEnum
	productCategory *ProductCategoryEnum
}

// Location of interest
func (r ApiGetLocationsRequest) Location(location LocationEnum) ApiGetLocationsRequest {
	r.location = &location
	return r
}

// Product category of interest
func (r ApiGetLocationsRequest) ProductCategory(productCategory ProductCategoryEnum) ApiGetLocationsRequest {
	r.productCategory = &productCategory
	return r
}

func (r ApiGetLocationsRequest) Execute() ([]Location, *http.Response, error) {
	return r.ApiService.GetLocationsExecute(r)
}

/*
GetLocations Get All Locations

Retrieve the locations info.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetLocationsRequest
*/
func (a *LocationsApiService) GetLocations(ctx context.Context) ApiGetLocationsRequest {
	return ApiGetLocationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []Location
func (a *LocationsApiService) GetLocationsExecute(r ApiGetLocationsRequest) ([]Location, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Location
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsApiService.GetLocations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/locations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.location != nil {
		localVarQueryParams.Add("location", parameterToString(*r.location, ""))
	}
	if r.productCategory != nil {
		localVarQueryParams.Add("productCategory", parameterToString(*r.productCategory, ""))
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
