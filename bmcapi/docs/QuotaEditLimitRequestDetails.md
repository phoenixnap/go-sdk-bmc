# QuotaEditLimitRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The new limit that is requested. | 
**Reason** | **string** | The reason for changing the limit. | 
**RequestedOn** | **time.Time** | The point in time the request was submitted. | 

## Methods

### NewQuotaEditLimitRequestDetails

`func NewQuotaEditLimitRequestDetails(limit int32, reason string, requestedOn time.Time, ) *QuotaEditLimitRequestDetails`

NewQuotaEditLimitRequestDetails instantiates a new QuotaEditLimitRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaEditLimitRequestDetailsWithDefaults

`func NewQuotaEditLimitRequestDetailsWithDefaults() *QuotaEditLimitRequestDetails`

NewQuotaEditLimitRequestDetailsWithDefaults instantiates a new QuotaEditLimitRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *QuotaEditLimitRequestDetails) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QuotaEditLimitRequestDetails) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QuotaEditLimitRequestDetails) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetReason

`func (o *QuotaEditLimitRequestDetails) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *QuotaEditLimitRequestDetails) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *QuotaEditLimitRequestDetails) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRequestedOn

`func (o *QuotaEditLimitRequestDetails) GetRequestedOn() time.Time`

GetRequestedOn returns the RequestedOn field if non-nil, zero value otherwise.

### GetRequestedOnOk

`func (o *QuotaEditLimitRequestDetails) GetRequestedOnOk() (*time.Time, bool)`

GetRequestedOnOk returns a tuple with the RequestedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedOn

`func (o *QuotaEditLimitRequestDetails) SetRequestedOn(v time.Time)`

SetRequestedOn sets RequestedOn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


