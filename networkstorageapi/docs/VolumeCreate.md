# VolumeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Volume friendly name. | 
**Description** | Pointer to **string** | Volume description. | [optional] 
**PathSuffix** | Pointer to **string** | Last part of volume&#39;s path. | [optional] 
**CapacityInGb** | **int32** | Capacity of Volume in GB. Currently only whole numbers and multiples of 1000GB are supported. | 
**Permissions** | Pointer to [**PermissionsCreate**](PermissionsCreate.md) |  | [optional] 

## Methods

### NewVolumeCreate

`func NewVolumeCreate(name string, capacityInGb int32, ) *VolumeCreate`

NewVolumeCreate instantiates a new VolumeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeCreateWithDefaults

`func NewVolumeCreateWithDefaults() *VolumeCreate`

NewVolumeCreateWithDefaults instantiates a new VolumeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VolumeCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VolumeCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPathSuffix

`func (o *VolumeCreate) GetPathSuffix() string`

GetPathSuffix returns the PathSuffix field if non-nil, zero value otherwise.

### GetPathSuffixOk

`func (o *VolumeCreate) GetPathSuffixOk() (*string, bool)`

GetPathSuffixOk returns a tuple with the PathSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSuffix

`func (o *VolumeCreate) SetPathSuffix(v string)`

SetPathSuffix sets PathSuffix field to given value.

### HasPathSuffix

`func (o *VolumeCreate) HasPathSuffix() bool`

HasPathSuffix returns a boolean if a field has been set.

### GetCapacityInGb

`func (o *VolumeCreate) GetCapacityInGb() int32`

GetCapacityInGb returns the CapacityInGb field if non-nil, zero value otherwise.

### GetCapacityInGbOk

`func (o *VolumeCreate) GetCapacityInGbOk() (*int32, bool)`

GetCapacityInGbOk returns a tuple with the CapacityInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInGb

`func (o *VolumeCreate) SetCapacityInGb(v int32)`

SetCapacityInGb sets CapacityInGb field to given value.


### GetPermissions

`func (o *VolumeCreate) GetPermissions() PermissionsCreate`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *VolumeCreate) GetPermissionsOk() (*PermissionsCreate, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *VolumeCreate) SetPermissions(v PermissionsCreate)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *VolumeCreate) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


