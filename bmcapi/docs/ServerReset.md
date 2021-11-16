# ServerReset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallDefaultSshKeys** | Pointer to **bool** | Whether or not to install SSH keys marked as default in addition to any SSH keys specified in this request. | [optional] [default to true]
**SshKeys** | Pointer to **[]string** | A list of SSH keys that will be installed on the server. | [optional] 
**SshKeyIds** | Pointer to **[]string** | A list of SSH key IDs that will be installed on the server in addition to any SSH keys specified in this request. | [optional] 
**OsConfiguration** | Pointer to [**OsConfigurationMap**](OsConfigurationMap.md) |  | [optional] 

## Methods

### NewServerReset

`func NewServerReset() *ServerReset`

NewServerReset instantiates a new ServerReset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerResetWithDefaults

`func NewServerResetWithDefaults() *ServerReset`

NewServerResetWithDefaults instantiates a new ServerReset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallDefaultSshKeys

`func (o *ServerReset) GetInstallDefaultSshKeys() bool`

GetInstallDefaultSshKeys returns the InstallDefaultSshKeys field if non-nil, zero value otherwise.

### GetInstallDefaultSshKeysOk

`func (o *ServerReset) GetInstallDefaultSshKeysOk() (*bool, bool)`

GetInstallDefaultSshKeysOk returns a tuple with the InstallDefaultSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDefaultSshKeys

`func (o *ServerReset) SetInstallDefaultSshKeys(v bool)`

SetInstallDefaultSshKeys sets InstallDefaultSshKeys field to given value.

### HasInstallDefaultSshKeys

`func (o *ServerReset) HasInstallDefaultSshKeys() bool`

HasInstallDefaultSshKeys returns a boolean if a field has been set.

### GetSshKeys

`func (o *ServerReset) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *ServerReset) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *ServerReset) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *ServerReset) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetSshKeyIds

`func (o *ServerReset) GetSshKeyIds() []string`

GetSshKeyIds returns the SshKeyIds field if non-nil, zero value otherwise.

### GetSshKeyIdsOk

`func (o *ServerReset) GetSshKeyIdsOk() (*[]string, bool)`

GetSshKeyIdsOk returns a tuple with the SshKeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyIds

`func (o *ServerReset) SetSshKeyIds(v []string)`

SetSshKeyIds sets SshKeyIds field to given value.

### HasSshKeyIds

`func (o *ServerReset) HasSshKeyIds() bool`

HasSshKeyIds returns a boolean if a field has been set.

### GetOsConfiguration

`func (o *ServerReset) GetOsConfiguration() OsConfigurationMap`

GetOsConfiguration returns the OsConfiguration field if non-nil, zero value otherwise.

### GetOsConfigurationOk

`func (o *ServerReset) GetOsConfigurationOk() (*OsConfigurationMap, bool)`

GetOsConfigurationOk returns a tuple with the OsConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsConfiguration

`func (o *ServerReset) SetOsConfiguration(v OsConfigurationMap)`

SetOsConfiguration sets OsConfiguration field to given value.

### HasOsConfiguration

`func (o *ServerReset) HasOsConfiguration() bool`

HasOsConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


