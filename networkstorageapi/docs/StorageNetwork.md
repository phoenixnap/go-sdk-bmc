# StorageNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Storage network ID. | [optional] 
**Name** | Pointer to **string** | Storage network friendly name. | [optional] 
**Description** | Pointer to **string** | Storage network description. | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Location** | Pointer to **string** | Location of storage network. Currently this field should be set to &#x60;PHX&#x60; or &#x60;ASH&#x60;. | [optional] 
**NetworkId** | Pointer to **string** | Id of network the storage belongs to. | [optional] 
**Ips** | Pointer to **[]string** | IP of the storage network. | [optional] 
**CreatedOn** | Pointer to **time.Time** | Date and time when this storage network was created. | [optional] 
**DeleteRequestedOn** | Pointer to **time.Time** | Date and time of the initial request for storage network deletion. | [optional] 
**Volumes** | Pointer to [**[]Volume**](Volume.md) | Volume for a storage network. | [optional] 

## Methods

### NewStorageNetwork

`func NewStorageNetwork() *StorageNetwork`

NewStorageNetwork instantiates a new StorageNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetworkWithDefaults

`func NewStorageNetworkWithDefaults() *StorageNetwork`

NewStorageNetworkWithDefaults instantiates a new StorageNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageNetwork) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StorageNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StorageNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StorageNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *StorageNetwork) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageNetwork) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageNetwork) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageNetwork) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLocation

`func (o *StorageNetwork) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *StorageNetwork) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *StorageNetwork) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *StorageNetwork) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNetworkId

`func (o *StorageNetwork) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *StorageNetwork) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *StorageNetwork) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *StorageNetwork) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetIps

`func (o *StorageNetwork) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *StorageNetwork) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *StorageNetwork) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *StorageNetwork) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetCreatedOn

`func (o *StorageNetwork) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *StorageNetwork) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *StorageNetwork) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *StorageNetwork) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetDeleteRequestedOn

`func (o *StorageNetwork) GetDeleteRequestedOn() time.Time`

GetDeleteRequestedOn returns the DeleteRequestedOn field if non-nil, zero value otherwise.

### GetDeleteRequestedOnOk

`func (o *StorageNetwork) GetDeleteRequestedOnOk() (*time.Time, bool)`

GetDeleteRequestedOnOk returns a tuple with the DeleteRequestedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRequestedOn

`func (o *StorageNetwork) SetDeleteRequestedOn(v time.Time)`

SetDeleteRequestedOn sets DeleteRequestedOn field to given value.

### HasDeleteRequestedOn

`func (o *StorageNetwork) HasDeleteRequestedOn() bool`

HasDeleteRequestedOn returns a boolean if a field has been set.

### GetVolumes

`func (o *StorageNetwork) GetVolumes() []Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *StorageNetwork) GetVolumesOk() (*[]Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *StorageNetwork) SetVolumes(v []Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *StorageNetwork) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


