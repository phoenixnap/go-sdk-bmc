# StorageNetworkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Storage network friendly name. | 
**Description** | Pointer to **string** | Storage network description. | [optional] 
**Location** | **string** | Location of storage network. Currently this field should be set to &#x60;PHX&#x60; or &#x60;ASH&#x60;. | 
**Volumes** | [**[]VolumeCreate**](VolumeCreate.md) | Volume to be created alongside storage. Currently only 1 volume is supported. | 

## Methods

### NewStorageNetworkCreate

`func NewStorageNetworkCreate(name string, location string, volumes []VolumeCreate, ) *StorageNetworkCreate`

NewStorageNetworkCreate instantiates a new StorageNetworkCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetworkCreateWithDefaults

`func NewStorageNetworkCreateWithDefaults() *StorageNetworkCreate`

NewStorageNetworkCreateWithDefaults instantiates a new StorageNetworkCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageNetworkCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetworkCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetworkCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *StorageNetworkCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageNetworkCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageNetworkCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageNetworkCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *StorageNetworkCreate) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *StorageNetworkCreate) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *StorageNetworkCreate) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetVolumes

`func (o *StorageNetworkCreate) GetVolumes() []VolumeCreate`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *StorageNetworkCreate) GetVolumesOk() (*[]VolumeCreate, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *StorageNetworkCreate) SetVolumes(v []VolumeCreate)`

SetVolumes sets Volumes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


