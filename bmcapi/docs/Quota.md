# Quota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the Quota. | 
**Name** | **string** | The name of the Quota. | 
**Description** | **string** | The Quota description. | 
**Status** | **string** | The status of the quota resource usage. | 
**Limit** | **int32** | The limit set for the quota. | 
**Unit** | **string** | An enum field describing what the limit is measured in. | 
**Used** | **int32** | The quota used expressed as a number. | 
**QuotaEditLimitRequestDetails** | [**[]QuotaEditLimitRequestDetails**](QuotaEditLimitRequestDetails.md) |  | [readonly] 

## Methods

### NewQuota

`func NewQuota(id string, name string, description string, status string, limit int32, unit string, used int32, quotaEditLimitRequestDetails []QuotaEditLimitRequestDetails, ) *Quota`

NewQuota instantiates a new Quota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaWithDefaults

`func NewQuotaWithDefaults() *Quota`

NewQuotaWithDefaults instantiates a new Quota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Quota) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Quota) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Quota) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Quota) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Quota) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Quota) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Quota) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Quota) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Quota) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *Quota) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Quota) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Quota) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLimit

`func (o *Quota) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Quota) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Quota) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetUnit

`func (o *Quota) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Quota) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Quota) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetUsed

`func (o *Quota) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Quota) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Quota) SetUsed(v int32)`

SetUsed sets Used field to given value.


### GetQuotaEditLimitRequestDetails

`func (o *Quota) GetQuotaEditLimitRequestDetails() []QuotaEditLimitRequestDetails`

GetQuotaEditLimitRequestDetails returns the QuotaEditLimitRequestDetails field if non-nil, zero value otherwise.

### GetQuotaEditLimitRequestDetailsOk

`func (o *Quota) GetQuotaEditLimitRequestDetailsOk() (*[]QuotaEditLimitRequestDetails, bool)`

GetQuotaEditLimitRequestDetailsOk returns a tuple with the QuotaEditLimitRequestDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaEditLimitRequestDetails

`func (o *Quota) SetQuotaEditLimitRequestDetails(v []QuotaEditLimitRequestDetails)`

SetQuotaEditLimitRequestDetails sets QuotaEditLimitRequestDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


