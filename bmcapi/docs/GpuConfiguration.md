# GpuConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LongName** | Pointer to **string** | The long name of the GPU. | [optional] 
**Count** | Pointer to **int32** | The number of GPUs. | [optional] 

## Methods

### NewGpuConfiguration

`func NewGpuConfiguration() *GpuConfiguration`

NewGpuConfiguration instantiates a new GpuConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuConfigurationWithDefaults

`func NewGpuConfigurationWithDefaults() *GpuConfiguration`

NewGpuConfigurationWithDefaults instantiates a new GpuConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLongName

`func (o *GpuConfiguration) GetLongName() string`

GetLongName returns the LongName field if non-nil, zero value otherwise.

### GetLongNameOk

`func (o *GpuConfiguration) GetLongNameOk() (*string, bool)`

GetLongNameOk returns a tuple with the LongName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongName

`func (o *GpuConfiguration) SetLongName(v string)`

SetLongName sets LongName field to given value.

### HasLongName

`func (o *GpuConfiguration) HasLongName() bool`

HasLongName returns a boolean if a field has been set.

### GetCount

`func (o *GpuConfiguration) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GpuConfiguration) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GpuConfiguration) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GpuConfiguration) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


