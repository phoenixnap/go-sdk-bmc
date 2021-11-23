# QuotaEditLimitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The new limit that is requested. | 
**Reason** | **string** | The reason for changing the limit. | 

## Methods

### NewQuotaEditLimitRequest

`func NewQuotaEditLimitRequest(limit int32, reason string, ) *QuotaEditLimitRequest`

NewQuotaEditLimitRequest instantiates a new QuotaEditLimitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaEditLimitRequestWithDefaults

`func NewQuotaEditLimitRequestWithDefaults() *QuotaEditLimitRequest`

NewQuotaEditLimitRequestWithDefaults instantiates a new QuotaEditLimitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *QuotaEditLimitRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QuotaEditLimitRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QuotaEditLimitRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetReason

`func (o *QuotaEditLimitRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *QuotaEditLimitRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *QuotaEditLimitRequest) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


