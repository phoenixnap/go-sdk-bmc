# VolumeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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


