# StorageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkStorageId** | **string** | Network storage ID. | 
**NetworkStorageName** | **string** | Network storage name. | 
**VolumeId** | **string** | Volume ID. | 
**VolumeName** | **string** | Volume name. | 
**CapacityInGb** | **int64** | Capacity in GB. | 
**CreatedOn** | **time.Time** | Timestamp when the record was created. | 

## Methods

### NewStorageDetails

`func NewStorageDetails(networkStorageId string, networkStorageName string, volumeId string, volumeName string, capacityInGb int64, createdOn time.Time, ) *StorageDetails`

NewStorageDetails instantiates a new StorageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDetailsWithDefaults

`func NewStorageDetailsWithDefaults() *StorageDetails`

NewStorageDetailsWithDefaults instantiates a new StorageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkStorageId

`func (o *StorageDetails) GetNetworkStorageId() string`

GetNetworkStorageId returns the NetworkStorageId field if non-nil, zero value otherwise.

### GetNetworkStorageIdOk

`func (o *StorageDetails) GetNetworkStorageIdOk() (*string, bool)`

GetNetworkStorageIdOk returns a tuple with the NetworkStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStorageId

`func (o *StorageDetails) SetNetworkStorageId(v string)`

SetNetworkStorageId sets NetworkStorageId field to given value.


### GetNetworkStorageName

`func (o *StorageDetails) GetNetworkStorageName() string`

GetNetworkStorageName returns the NetworkStorageName field if non-nil, zero value otherwise.

### GetNetworkStorageNameOk

`func (o *StorageDetails) GetNetworkStorageNameOk() (*string, bool)`

GetNetworkStorageNameOk returns a tuple with the NetworkStorageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStorageName

`func (o *StorageDetails) SetNetworkStorageName(v string)`

SetNetworkStorageName sets NetworkStorageName field to given value.


### GetVolumeId

`func (o *StorageDetails) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *StorageDetails) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *StorageDetails) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.


### GetVolumeName

`func (o *StorageDetails) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *StorageDetails) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *StorageDetails) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.


### GetCapacityInGb

`func (o *StorageDetails) GetCapacityInGb() int64`

GetCapacityInGb returns the CapacityInGb field if non-nil, zero value otherwise.

### GetCapacityInGbOk

`func (o *StorageDetails) GetCapacityInGbOk() (*int64, bool)`

GetCapacityInGbOk returns a tuple with the CapacityInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInGb

`func (o *StorageDetails) SetCapacityInGb(v int64)`

SetCapacityInGb sets CapacityInGb field to given value.


### GetCreatedOn

`func (o *StorageDetails) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *StorageDetails) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *StorageDetails) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


