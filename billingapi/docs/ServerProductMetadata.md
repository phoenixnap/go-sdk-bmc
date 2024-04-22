# ServerProductMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RamInGb** | **float32** | RAM in GB. | 
**Cpu** | **string** | CPU name. | 
**CpuCount** | **float32** | Number of CPUs. | 
**CoresPerCpu** | **float32** | The number of physical cores present on each CPU. | 
**CpuFrequency** | **float32** | CPU frequency in GHz. | 
**Network** | **string** | Server network. | 
**Storage** | **string** | Server storage. | 
**GpuConfigurations** | Pointer to [**[]GpuConfigurationMetadata**](GpuConfigurationMetadata.md) | GPU configurations | [optional] 

## Methods

### NewServerProductMetadata

`func NewServerProductMetadata(ramInGb float32, cpu string, cpuCount float32, coresPerCpu float32, cpuFrequency float32, network string, storage string, ) *ServerProductMetadata`

NewServerProductMetadata instantiates a new ServerProductMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerProductMetadataWithDefaults

`func NewServerProductMetadataWithDefaults() *ServerProductMetadata`

NewServerProductMetadataWithDefaults instantiates a new ServerProductMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRamInGb

`func (o *ServerProductMetadata) GetRamInGb() float32`

GetRamInGb returns the RamInGb field if non-nil, zero value otherwise.

### GetRamInGbOk

`func (o *ServerProductMetadata) GetRamInGbOk() (*float32, bool)`

GetRamInGbOk returns a tuple with the RamInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamInGb

`func (o *ServerProductMetadata) SetRamInGb(v float32)`

SetRamInGb sets RamInGb field to given value.


### GetCpu

`func (o *ServerProductMetadata) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ServerProductMetadata) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ServerProductMetadata) SetCpu(v string)`

SetCpu sets Cpu field to given value.


### GetCpuCount

`func (o *ServerProductMetadata) GetCpuCount() float32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *ServerProductMetadata) GetCpuCountOk() (*float32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *ServerProductMetadata) SetCpuCount(v float32)`

SetCpuCount sets CpuCount field to given value.


### GetCoresPerCpu

`func (o *ServerProductMetadata) GetCoresPerCpu() float32`

GetCoresPerCpu returns the CoresPerCpu field if non-nil, zero value otherwise.

### GetCoresPerCpuOk

`func (o *ServerProductMetadata) GetCoresPerCpuOk() (*float32, bool)`

GetCoresPerCpuOk returns a tuple with the CoresPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerCpu

`func (o *ServerProductMetadata) SetCoresPerCpu(v float32)`

SetCoresPerCpu sets CoresPerCpu field to given value.


### GetCpuFrequency

`func (o *ServerProductMetadata) GetCpuFrequency() float32`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *ServerProductMetadata) GetCpuFrequencyOk() (*float32, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *ServerProductMetadata) SetCpuFrequency(v float32)`

SetCpuFrequency sets CpuFrequency field to given value.


### GetNetwork

`func (o *ServerProductMetadata) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ServerProductMetadata) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ServerProductMetadata) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetStorage

`func (o *ServerProductMetadata) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ServerProductMetadata) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ServerProductMetadata) SetStorage(v string)`

SetStorage sets Storage field to given value.


### GetGpuConfigurations

`func (o *ServerProductMetadata) GetGpuConfigurations() []GpuConfigurationMetadata`

GetGpuConfigurations returns the GpuConfigurations field if non-nil, zero value otherwise.

### GetGpuConfigurationsOk

`func (o *ServerProductMetadata) GetGpuConfigurationsOk() (*[]GpuConfigurationMetadata, bool)`

GetGpuConfigurationsOk returns a tuple with the GpuConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuConfigurations

`func (o *ServerProductMetadata) SetGpuConfigurations(v []GpuConfigurationMetadata)`

SetGpuConfigurations sets GpuConfigurations field to given value.

### HasGpuConfigurations

`func (o *ServerProductMetadata) HasGpuConfigurations() bool`

HasGpuConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


