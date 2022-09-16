# NfsPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadWrite** | Pointer to **[]string** | Read/Write access. | [optional] 
**ReadOnly** | Pointer to **[]string** | Read only access. | [optional] 
**RootSquash** | Pointer to **[]string** | Root squash permission. | [optional] 
**NoSquash** | Pointer to **[]string** | No squash permission. | [optional] 
**AllSquash** | Pointer to **[]string** | All squash permission. | [optional] 

## Methods

### NewNfsPermissions

`func NewNfsPermissions() *NfsPermissions`

NewNfsPermissions instantiates a new NfsPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsPermissionsWithDefaults

`func NewNfsPermissionsWithDefaults() *NfsPermissions`

NewNfsPermissionsWithDefaults instantiates a new NfsPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadWrite

`func (o *NfsPermissions) GetReadWrite() []string`

GetReadWrite returns the ReadWrite field if non-nil, zero value otherwise.

### GetReadWriteOk

`func (o *NfsPermissions) GetReadWriteOk() (*[]string, bool)`

GetReadWriteOk returns a tuple with the ReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWrite

`func (o *NfsPermissions) SetReadWrite(v []string)`

SetReadWrite sets ReadWrite field to given value.

### HasReadWrite

`func (o *NfsPermissions) HasReadWrite() bool`

HasReadWrite returns a boolean if a field has been set.

### GetReadOnly

`func (o *NfsPermissions) GetReadOnly() []string`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *NfsPermissions) GetReadOnlyOk() (*[]string, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *NfsPermissions) SetReadOnly(v []string)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *NfsPermissions) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRootSquash

`func (o *NfsPermissions) GetRootSquash() []string`

GetRootSquash returns the RootSquash field if non-nil, zero value otherwise.

### GetRootSquashOk

`func (o *NfsPermissions) GetRootSquashOk() (*[]string, bool)`

GetRootSquashOk returns a tuple with the RootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSquash

`func (o *NfsPermissions) SetRootSquash(v []string)`

SetRootSquash sets RootSquash field to given value.

### HasRootSquash

`func (o *NfsPermissions) HasRootSquash() bool`

HasRootSquash returns a boolean if a field has been set.

### GetNoSquash

`func (o *NfsPermissions) GetNoSquash() []string`

GetNoSquash returns the NoSquash field if non-nil, zero value otherwise.

### GetNoSquashOk

`func (o *NfsPermissions) GetNoSquashOk() (*[]string, bool)`

GetNoSquashOk returns a tuple with the NoSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSquash

`func (o *NfsPermissions) SetNoSquash(v []string)`

SetNoSquash sets NoSquash field to given value.

### HasNoSquash

`func (o *NfsPermissions) HasNoSquash() bool`

HasNoSquash returns a boolean if a field has been set.

### GetAllSquash

`func (o *NfsPermissions) GetAllSquash() []string`

GetAllSquash returns the AllSquash field if non-nil, zero value otherwise.

### GetAllSquashOk

`func (o *NfsPermissions) GetAllSquashOk() (*[]string, bool)`

GetAllSquashOk returns a tuple with the AllSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSquash

`func (o *NfsPermissions) SetAllSquash(v []string)`

SetAllSquash sets AllSquash field to given value.

### HasAllSquash

`func (o *NfsPermissions) HasAllSquash() bool`

HasAllSquash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


