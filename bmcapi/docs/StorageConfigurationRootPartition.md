# StorageConfigurationRootPartition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Raid** | Pointer to **string** | Software RAID configuration. The following RAID options are available: NO_RAID, RAID_0, RAID_1. | [optional] [default to "NO_RAID"]
**Size** | Pointer to **int32** | The size of the root partition in GB. -1 to use all available space. | [optional] [default to -1]

## Methods

### NewStorageConfigurationRootPartition

`func NewStorageConfigurationRootPartition() *StorageConfigurationRootPartition`

NewStorageConfigurationRootPartition instantiates a new StorageConfigurationRootPartition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageConfigurationRootPartitionWithDefaults

`func NewStorageConfigurationRootPartitionWithDefaults() *StorageConfigurationRootPartition`

NewStorageConfigurationRootPartitionWithDefaults instantiates a new StorageConfigurationRootPartition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaid

`func (o *StorageConfigurationRootPartition) GetRaid() string`

GetRaid returns the Raid field if non-nil, zero value otherwise.

### GetRaidOk

`func (o *StorageConfigurationRootPartition) GetRaidOk() (*string, bool)`

GetRaidOk returns a tuple with the Raid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaid

`func (o *StorageConfigurationRootPartition) SetRaid(v string)`

SetRaid sets Raid field to given value.

### HasRaid

`func (o *StorageConfigurationRootPartition) HasRaid() bool`

HasRaid returns a boolean if a field has been set.

### GetSize

`func (o *StorageConfigurationRootPartition) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageConfigurationRootPartition) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageConfigurationRootPartition) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageConfigurationRootPartition) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


