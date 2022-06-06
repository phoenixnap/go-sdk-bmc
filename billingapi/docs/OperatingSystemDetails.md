# OperatingSystemDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cores** | **int32** | Number of cores. | 
**CorrelationId** | **string** | Correlation is used to associate Operating System License charges and the Server on which it was installed. In this scenario, the correlation ID will be equal to the rated usage record ID representing the charge for the Server. | 

## Methods

### NewOperatingSystemDetails

`func NewOperatingSystemDetails(cores int32, correlationId string, ) *OperatingSystemDetails`

NewOperatingSystemDetails instantiates a new OperatingSystemDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatingSystemDetailsWithDefaults

`func NewOperatingSystemDetailsWithDefaults() *OperatingSystemDetails`

NewOperatingSystemDetailsWithDefaults instantiates a new OperatingSystemDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCores

`func (o *OperatingSystemDetails) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *OperatingSystemDetails) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *OperatingSystemDetails) SetCores(v int32)`

SetCores sets Cores field to given value.


### GetCorrelationId

`func (o *OperatingSystemDetails) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *OperatingSystemDetails) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *OperatingSystemDetails) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


