# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Volume ID. | [optional] 
**Name** | Pointer to **string** | Volume friendly name. | [optional] 
**Description** | Pointer to **string** | Volume description. | [optional] 
**Path** | Pointer to **string** | Volume&#39;s full path. It is in form of &#x60;/{volumeId}/pathSuffix&#x60;&#39;. | [optional] 
**PathSuffix** | Pointer to **string** | Last part of volume&#39;s path. | [optional] 
**CapacityInGb** | Pointer to **int32** | Maximum capacity in GB. | [optional] 
**UsedCapacityInGb** | Pointer to **int32** | Used capacity in GB, updated periodically. | [optional] 
**Protocol** | Pointer to **string** | File system protocol. Currently this field should be set to &#x60;NFS&#x60;. | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**Permissions** | Pointer to [**Permissions**](Permissions.md) |  | [optional] 

## Methods

### NewVolume

`func NewVolume() *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Volume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Volume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Volume) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Volume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Volume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Volume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Volume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Volume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Volume) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Volume) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Volume) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Volume) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPath

`func (o *Volume) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Volume) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Volume) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Volume) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPathSuffix

`func (o *Volume) GetPathSuffix() string`

GetPathSuffix returns the PathSuffix field if non-nil, zero value otherwise.

### GetPathSuffixOk

`func (o *Volume) GetPathSuffixOk() (*string, bool)`

GetPathSuffixOk returns a tuple with the PathSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSuffix

`func (o *Volume) SetPathSuffix(v string)`

SetPathSuffix sets PathSuffix field to given value.

### HasPathSuffix

`func (o *Volume) HasPathSuffix() bool`

HasPathSuffix returns a boolean if a field has been set.

### GetCapacityInGb

`func (o *Volume) GetCapacityInGb() int32`

GetCapacityInGb returns the CapacityInGb field if non-nil, zero value otherwise.

### GetCapacityInGbOk

`func (o *Volume) GetCapacityInGbOk() (*int32, bool)`

GetCapacityInGbOk returns a tuple with the CapacityInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInGb

`func (o *Volume) SetCapacityInGb(v int32)`

SetCapacityInGb sets CapacityInGb field to given value.

### HasCapacityInGb

`func (o *Volume) HasCapacityInGb() bool`

HasCapacityInGb returns a boolean if a field has been set.

### GetUsedCapacityInGb

`func (o *Volume) GetUsedCapacityInGb() int32`

GetUsedCapacityInGb returns the UsedCapacityInGb field if non-nil, zero value otherwise.

### GetUsedCapacityInGbOk

`func (o *Volume) GetUsedCapacityInGbOk() (*int32, bool)`

GetUsedCapacityInGbOk returns a tuple with the UsedCapacityInGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCapacityInGb

`func (o *Volume) SetUsedCapacityInGb(v int32)`

SetUsedCapacityInGb sets UsedCapacityInGb field to given value.

### HasUsedCapacityInGb

`func (o *Volume) HasUsedCapacityInGb() bool`

HasUsedCapacityInGb returns a boolean if a field has been set.

### GetProtocol

`func (o *Volume) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Volume) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Volume) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Volume) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStatus

`func (o *Volume) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Volume) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Volume) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Volume) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Volume) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Volume) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Volume) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Volume) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetPermissions

`func (o *Volume) GetPermissions() Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Volume) GetPermissionsOk() (*Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Volume) SetPermissions(v Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Volume) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


