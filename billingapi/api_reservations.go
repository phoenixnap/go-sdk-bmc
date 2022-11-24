/*
Billing API

Automate your infrastructure billing with the Bare Metal Cloud Billing API. Reserve your server instances to ensure guaranteed resource availability for 12, 24, and 36 months. Retrieve your server’s rated usage for a given period and enable or disable auto-renewals.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/phoenixnap-bare-metal-cloud-billing-models' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/billing/v1/)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package billingapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ReservationsApi interface {

	/*
		ReservationsGet List all Reservations.

		Retrieves all reservations associated with the authenticated account. All date & times are in UTC.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiReservationsGetRequest
	*/
	ReservationsGet(ctx context.Context) ApiReservationsGetRequest

	// ReservationsGetExecute executes the request
	//  @return []Reservation
	ReservationsGetExecute(r ApiReservationsGetRequest) ([]Reservation, *http.Response, error)

	/*
		ReservationsPost Create a reservation.

		Creates new package reservation for authenticated account.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiReservationsPostRequest
	*/
	ReservationsPost(ctx context.Context) ApiReservationsPostRequest

	// ReservationsPostExecute executes the request
	//  @return Reservation
	ReservationsPostExecute(r ApiReservationsPostRequest) (*Reservation, *http.Response, error)

	/*
		ReservationsReservationIdActionsAutoRenewDisablePost Disable auto-renewal for reservation by id.

		Disable auto-renewal for reservation by reservation id.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reservationId The reservation's ID.
		@return ApiReservationsReservationIdActionsAutoRenewDisablePostRequest
	*/
	ReservationsReservationIdActionsAutoRenewDisablePost(ctx context.Context, reservationId string) ApiReservationsReservationIdActionsAutoRenewDisablePostRequest

	// ReservationsReservationIdActionsAutoRenewDisablePostExecute executes the request
	//  @return Reservation
	ReservationsReservationIdActionsAutoRenewDisablePostExecute(r ApiReservationsReservationIdActionsAutoRenewDisablePostRequest) (*Reservation, *http.Response, error)

	/*
		ReservationsReservationIdActionsAutoRenewEnablePost Enable auto-renewal for unexpired reservation by reservation id.

		Enable auto-renewal for unexpired reservation by reservation id.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reservationId The reservation's ID.
		@return ApiReservationsReservationIdActionsAutoRenewEnablePostRequest
	*/
	ReservationsReservationIdActionsAutoRenewEnablePost(ctx context.Context, reservationId string) ApiReservationsReservationIdActionsAutoRenewEnablePostRequest

	// ReservationsReservationIdActionsAutoRenewEnablePostExecute executes the request
	//  @return Reservation
	ReservationsReservationIdActionsAutoRenewEnablePostExecute(r ApiReservationsReservationIdActionsAutoRenewEnablePostRequest) (*Reservation, *http.Response, error)

	/*
		ReservationsReservationIdActionsConvertPost Convert reservation pricing model by reservation ID.

		Convert reservation pricing model by reservation id.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reservationId The reservation's ID.
		@return ApiReservationsReservationIdActionsConvertPostRequest
	*/
	ReservationsReservationIdActionsConvertPost(ctx context.Context, reservationId string) ApiReservationsReservationIdActionsConvertPostRequest

	// ReservationsReservationIdActionsConvertPostExecute executes the request
	//  @return Reservation
	ReservationsReservationIdActionsConvertPostExecute(r ApiReservationsReservationIdActionsConvertPostRequest) (*Reservation, *http.Response, error)

	/*
		ReservationsReservationIdGet Get a reservation.

		Retrieves the reservations with the specified identifier. All date & times are in UTC.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param reservationId The reservation's ID.
		@return ApiReservationsReservationIdGetRequest
	*/
	ReservationsReservationIdGet(ctx context.Context, reservationId string) ApiReservationsReservationIdGetRequest

	// ReservationsReservationIdGetExecute executes the request
	//  @return Reservation
	ReservationsReservationIdGetExecute(r ApiReservationsReservationIdGetRequest) (*Reservation, *http.Response, error)
}

// ReservationsApiService ReservationsApi service
type ReservationsApiService service

type ApiReservationsGetRequest struct {
	ctx             context.Context
	ApiService      ReservationsApi
	productCategory *ProductCategoryEnum
}

// The product category
func (r ApiReservationsGetRequest) ProductCategory(productCategory ProductCategoryEnum) ApiReservationsGetRequest {
	r.productCategory = &productCategory
	return r
}

func (r ApiReservationsGetRequest) Execute() ([]Reservation, *http.Response, error) {
	return r.ApiService.ReservationsGetExecute(r)
}

/*
ReservationsGet List all Reservations.

Retrieves all reservations associated with the authenticated account. All date & times are in UTC.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiReservationsGetRequest
*/
func (a *ReservationsApiService) ReservationsGet(ctx context.Context) ApiReservationsGetRequest {
	return ApiReservationsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Reservation
func (a *ReservationsApiService) ReservationsGetExecute(r ApiReservationsGetRequest) ([]Reservation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Reservation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReservationsApiService.ReservationsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reservations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiReservationsPostRequest struct {
	ctx                context.Context
	ApiService         ReservationsApi
	reservationRequest *ReservationRequest
}

func (r ApiReservationsPostRequest) ReservationRequest(reservationRequest ReservationRequest) ApiReservationsPostRequest {
	r.reservationRequest = &reservationRequest
	return r
}

func (r ApiReservationsPostRequest) Execute() (*Reservation, *http.Response, error) {
	return r.ApiService.ReservationsPostExecute(r)
}

/*
ReservationsPost Create a reservation.

Creates new package reservation for authenticated account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiReservationsPostRequest
*/
func (a *ReservationsApiService) ReservationsPost(ctx context.Context) ApiReservationsPostRequest {
	return ApiReservationsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Reservation
func (a *ReservationsApiService) ReservationsPostExecute(r ApiReservationsPostRequest) (*Reservation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Reservation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReservationsApiService.ReservationsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reservations"

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.reservationRequest
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

type ApiReservationsReservationIdActionsAutoRenewDisablePostRequest struct {
	ctx                                context.Context
	ApiService                         ReservationsApi
	reservationId                      string
	reservationAutoRenewDisableRequest *ReservationAutoRenewDisableRequest
}

func (r ApiReservationsReservationIdActionsAutoRenewDisablePostRequest) ReservationAutoRenewDisableRequest(reservationAutoRenewDisableRequest ReservationAutoRenewDisableRequest) ApiReservationsReservationIdActionsAutoRenewDisablePostRequest {
	r.reservationAutoRenewDisableRequest = &reservationAutoRenewDisableRequest
	return r
}

func (r ApiReservationsReservationIdActionsAutoRenewDisablePostRequest) Execute() (*Reservation, *http.Response, error) {
	return r.ApiService.ReservationsReservationIdActionsAutoRenewDisablePostExecute(r)
}

/*
ReservationsReservationIdActionsAutoRenewDisablePost Disable auto-renewal for reservation by id.

Disable auto-renewal for reservation by reservation id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reservationId The reservation's ID.
	@return ApiReservationsReservationIdActionsAutoRenewDisablePostRequest
*/
func (a *ReservationsApiService) ReservationsReservationIdActionsAutoRenewDisablePost(ctx context.Context, reservationId string) ApiReservationsReservationIdActionsAutoRenewDisablePostRequest {
	return ApiReservationsReservationIdActionsAutoRenewDisablePostRequest{
		ApiService:    a,
		ctx:           ctx,
		reservationId: reservationId,
	}
}

// Execute executes the request
//
//	@return Reservation
func (a *ReservationsApiService) ReservationsReservationIdActionsAutoRenewDisablePostExecute(r ApiReservationsReservationIdActionsAutoRenewDisablePostRequest) (*Reservation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Reservation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReservationsApiService.ReservationsReservationIdActionsAutoRenewDisablePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reservations/{reservationId}/actions/auto-renew/disable"
	localVarPath = strings.Replace(localVarPath, "{"+"reservationId"+"}", url.PathEscape(parameterToString(r.reservationId, "")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.reservationAutoRenewDisableRequest
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
		if localVarHTTPResponse.StatusCode == 409 {
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

type ApiReservationsReservationIdActionsAutoRenewEnablePostRequest struct {
	ctx           context.Context
	ApiService    ReservationsApi
	reservationId string
}

func (r ApiReservationsReservationIdActionsAutoRenewEnablePostRequest) Execute() (*Reservation, *http.Response, error) {
	return r.ApiService.ReservationsReservationIdActionsAutoRenewEnablePostExecute(r)
}

/*
ReservationsReservationIdActionsAutoRenewEnablePost Enable auto-renewal for unexpired reservation by reservation id.

Enable auto-renewal for unexpired reservation by reservation id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reservationId The reservation's ID.
	@return ApiReservationsReservationIdActionsAutoRenewEnablePostRequest
*/
func (a *ReservationsApiService) ReservationsReservationIdActionsAutoRenewEnablePost(ctx context.Context, reservationId string) ApiReservationsReservationIdActionsAutoRenewEnablePostRequest {
	return ApiReservationsReservationIdActionsAutoRenewEnablePostRequest{
		ApiService:    a,
		ctx:           ctx,
		reservationId: reservationId,
	}
}

// Execute executes the request
//
//	@return Reservation
func (a *ReservationsApiService) ReservationsReservationIdActionsAutoRenewEnablePostExecute(r ApiReservationsReservationIdActionsAutoRenewEnablePostRequest) (*Reservation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Reservation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReservationsApiService.ReservationsReservationIdActionsAutoRenewEnablePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reservations/{reservationId}/actions/auto-renew/enable"
	localVarPath = strings.Replace(localVarPath, "{"+"reservationId"+"}", url.PathEscape(parameterToString(r.reservationId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 409 {
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

type ApiReservationsReservationIdActionsConvertPostRequest struct {
	ctx                context.Context
	ApiService         ReservationsApi
	reservationId      string
	reservationRequest *ReservationRequest
}

func (r ApiReservationsReservationIdActionsConvertPostRequest) ReservationRequest(reservationRequest ReservationRequest) ApiReservationsReservationIdActionsConvertPostRequest {
	r.reservationRequest = &reservationRequest
	return r
}

func (r ApiReservationsReservationIdActionsConvertPostRequest) Execute() (*Reservation, *http.Response, error) {
	return r.ApiService.ReservationsReservationIdActionsConvertPostExecute(r)
}

/*
ReservationsReservationIdActionsConvertPost Convert reservation pricing model by reservation ID.

Convert reservation pricing model by reservation id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reservationId The reservation's ID.
	@return ApiReservationsReservationIdActionsConvertPostRequest
*/
func (a *ReservationsApiService) ReservationsReservationIdActionsConvertPost(ctx context.Context, reservationId string) ApiReservationsReservationIdActionsConvertPostRequest {
	return ApiReservationsReservationIdActionsConvertPostRequest{
		ApiService:    a,
		ctx:           ctx,
		reservationId: reservationId,
	}
}

// Execute executes the request
//
//	@return Reservation
func (a *ReservationsApiService) ReservationsReservationIdActionsConvertPostExecute(r ApiReservationsReservationIdActionsConvertPostRequest) (*Reservation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Reservation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReservationsApiService.ReservationsReservationIdActionsConvertPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reservations/{reservationId}/actions/convert"
	localVarPath = strings.Replace(localVarPath, "{"+"reservationId"+"}", url.PathEscape(parameterToString(r.reservationId, "")), -1)

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.reservationRequest
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

type ApiReservationsReservationIdGetRequest struct {
	ctx           context.Context
	ApiService    ReservationsApi
	reservationId string
}

func (r ApiReservationsReservationIdGetRequest) Execute() (*Reservation, *http.Response, error) {
	return r.ApiService.ReservationsReservationIdGetExecute(r)
}

/*
ReservationsReservationIdGet Get a reservation.

Retrieves the reservations with the specified identifier. All date & times are in UTC.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reservationId The reservation's ID.
	@return ApiReservationsReservationIdGetRequest
*/
func (a *ReservationsApiService) ReservationsReservationIdGet(ctx context.Context, reservationId string) ApiReservationsReservationIdGetRequest {
	return ApiReservationsReservationIdGetRequest{
		ApiService:    a,
		ctx:           ctx,
		reservationId: reservationId,
	}
}

// Execute executes the request
//
//	@return Reservation
func (a *ReservationsApiService) ReservationsReservationIdGetExecute(r ApiReservationsReservationIdGetRequest) (*Reservation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Reservation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReservationsApiService.ReservationsReservationIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reservations/{reservationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"reservationId"+"}", url.PathEscape(parameterToString(r.reservationId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 409 {
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
