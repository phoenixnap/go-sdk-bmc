# VolumeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Volume friendly name. | [optional] 
**Description** | Pointer to **string** | Volume description. | [optional] 
**CapacityInGb** | Pointer to **int32** | Capacity of Volume in GB. Currently only whole numbers and multiples of 1000GB are supported. | [optional] 

## Methods

### NewVolumeUpdate

`func NewVolumeUpdate() *VolumeUpdate`

NewVolumeUpdate instantiates a new VolumeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeUpdateWithDefaults

`func NewVolumeUpdateWithDefaults() *VolumeUpdate`

NewVolumeUpdateWithDefaults instantiates a new VolumeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VolumeUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VolumeUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCapacityInGb

`func (o *VolumeUpdate) GetCapacityInGb() int32`

GetCapacityInGb returns the CapacityInGb field if non-nil, zero value otherwise.

### GetCapacityInGbOk

`func (o *VolumeUpdate) GetCapacityInGbOk() (*int32, bool)`

GetCapacityInGbOk returns a tuple with the CapacityInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInGb

`func (o *VolumeUpdate) SetCapacityInGb(v int32)`

SetCapacityInGb sets CapacityInGb field to given value.

### HasCapacityInGb

`func (o *VolumeUpdate) HasCapacityInGb() bool`

HasCapacityInGb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


