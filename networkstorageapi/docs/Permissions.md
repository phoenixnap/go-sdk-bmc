# Permissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nfs** | Pointer to [**NfsPermissions**](NfsPermissions.md) |  | [optional] 

## Methods

### NewPermissions

`func NewPermissions() *Permissions`

NewPermissions instantiates a new Permissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsWithDefaults

`func NewPermissionsWithDefaults() *Permissions`

NewPermissionsWithDefaults instantiates a new Permissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfs

`func (o *Permissions) GetNfs() NfsPermissions`

GetNfs returns the Nfs field if non-nil, zero value otherwise.

### GetNfsOk

`func (o *Permissions) GetNfsOk() (*NfsPermissions, bool)`

GetNfsOk returns a tuple with the Nfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfs

`func (o *Permissions) SetNfs(v NfsPermissions)`

SetNfs sets Nfs field to given value.

### HasNfs

`func (o *Permissions) HasNfs() bool`

HasNfs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


