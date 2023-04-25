# NfsPermissionsCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadWrite** | Pointer to **[]string** | Read/Write access. | [optional] 
**ReadOnly** | Pointer to **[]string** | Read only access. | [optional] 
**RootSquash** | Pointer to **[]string** | Root squash permission. | [optional] 
**NoSquash** | Pointer to **[]string** | No squash permission. | [optional] 
**AllSquash** | Pointer to **[]string** | All squash permission. | [optional] 

## Methods

### NewNfsPermissionsCreate

`func NewNfsPermissionsCreate() *NfsPermissionsCreate`

NewNfsPermissionsCreate instantiates a new NfsPermissionsCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsPermissionsCreateWithDefaults

`func NewNfsPermissionsCreateWithDefaults() *NfsPermissionsCreate`

NewNfsPermissionsCreateWithDefaults instantiates a new NfsPermissionsCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadWrite

`func (o *NfsPermissionsCreate) GetReadWrite() []string`

GetReadWrite returns the ReadWrite field if non-nil, zero value otherwise.

### GetReadWriteOk

`func (o *NfsPermissionsCreate) GetReadWriteOk() (*[]string, bool)`

GetReadWriteOk returns a tuple with the ReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWrite

`func (o *NfsPermissionsCreate) SetReadWrite(v []string)`

SetReadWrite sets ReadWrite field to given value.

### HasReadWrite

`func (o *NfsPermissionsCreate) HasReadWrite() bool`

HasReadWrite returns a boolean if a field has been set.

### GetReadOnly

`func (o *NfsPermissionsCreate) GetReadOnly() []string`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *NfsPermissionsCreate) GetReadOnlyOk() (*[]string, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *NfsPermissionsCreate) SetReadOnly(v []string)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *NfsPermissionsCreate) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRootSquash

`func (o *NfsPermissionsCreate) GetRootSquash() []string`

GetRootSquash returns the RootSquash field if non-nil, zero value otherwise.

### GetRootSquashOk

`func (o *NfsPermissionsCreate) GetRootSquashOk() (*[]string, bool)`

GetRootSquashOk returns a tuple with the RootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSquash

`func (o *NfsPermissionsCreate) SetRootSquash(v []string)`

SetRootSquash sets RootSquash field to given value.

### HasRootSquash

`func (o *NfsPermissionsCreate) HasRootSquash() bool`

HasRootSquash returns a boolean if a field has been set.

### GetNoSquash

`func (o *NfsPermissionsCreate) GetNoSquash() []string`

GetNoSquash returns the NoSquash field if non-nil, zero value otherwise.

### GetNoSquashOk

`func (o *NfsPermissionsCreate) GetNoSquashOk() (*[]string, bool)`

GetNoSquashOk returns a tuple with the NoSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSquash

`func (o *NfsPermissionsCreate) SetNoSquash(v []string)`

SetNoSquash sets NoSquash field to given value.

### HasNoSquash

`func (o *NfsPermissionsCreate) HasNoSquash() bool`

HasNoSquash returns a boolean if a field has been set.

### GetAllSquash

`func (o *NfsPermissionsCreate) GetAllSquash() []string`

GetAllSquash returns the AllSquash field if non-nil, zero value otherwise.

### GetAllSquashOk

`func (o *NfsPermissionsCreate) GetAllSquashOk() (*[]string, bool)`

GetAllSquashOk returns a tuple with the AllSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSquash

`func (o *NfsPermissionsCreate) SetAllSquash(v []string)`

SetAllSquash sets AllSquash field to given value.

### HasAllSquash

`func (o *NfsPermissionsCreate) HasAllSquash() bool`

HasAllSquash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


