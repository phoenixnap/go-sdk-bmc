# GpuConfigurationMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **float32** | GPU card count. | [optional] 
**Name** | Pointer to **string** | GPU name. | [optional] 

## Methods

### NewGpuConfigurationMetadata

`func NewGpuConfigurationMetadata() *GpuConfigurationMetadata`

NewGpuConfigurationMetadata instantiates a new GpuConfigurationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuConfigurationMetadataWithDefaults

`func NewGpuConfigurationMetadataWithDefaults() *GpuConfigurationMetadata`

NewGpuConfigurationMetadataWithDefaults instantiates a new GpuConfigurationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GpuConfigurationMetadata) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GpuConfigurationMetadata) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GpuConfigurationMetadata) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GpuConfigurationMetadata) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetName

`func (o *GpuConfigurationMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GpuConfigurationMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GpuConfigurationMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GpuConfigurationMetadata) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


