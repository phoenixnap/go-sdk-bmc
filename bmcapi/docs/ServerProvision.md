# ServerProvision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | Hostname of server. | 
**Description** | Pointer to **string** | Description of server. | [optional] 
**Os** | **string** | The server’s OS ID used when the server was created. Currently this field should be set to either &#x60;ubuntu/bionic&#x60;, &#x60;ubuntu/focal&#x60;, &#x60;ubuntu/jammy&#x60;, &#x60;ubuntu/jammy+pytorch&#x60;, &#x60;ubuntu/noble&#x60;, &#x60;centos/centos7&#x60;, &#x60;centos/centos8&#x60;, &#x60;windows/srv2019std&#x60;, &#x60;windows/srv2019dc&#x60;, &#x60;windows/srv2022std&#x60;, &#x60;windows/srv2022dc&#x60;, &#x60;esxi/esxi70&#x60;, &#x60;esxi/esxi80&#x60;, &#x60;almalinux/almalinux8&#x60;, &#x60;almalinux/almalinux9&#x60;, &#x60;rockylinux/rockylinux8&#x60;, &#x60;rockylinux/rockylinux9&#x60;, &#x60;virtuozzo/virtuozzo7&#x60;, &#x60;oraclelinux/oraclelinux9&#x60;, &#x60;debian/bullseye&#x60;, &#x60;debian/bookworm&#x60;, &#x60;proxmox/bullseye&#x60;, &#x60;proxmox/proxmox8&#x60;, &#x60;netris/controller&#x60;, &#x60;netris/softgate_1g&#x60;, &#x60;netris/softgate_10g&#x60; or &#x60;netris/softgate_25g&#x60;. | 
**InstallDefaultSshKeys** | Pointer to **bool** | Whether or not to install SSH keys marked as default in addition to any SSH keys specified in this request. | [optional] [default to true]
**SshKeys** | Pointer to **[]string** | A list of SSH keys that will be installed on the server. | [optional] 
**SshKeyIds** | Pointer to **[]string** | A list of SSH key IDs that will be installed on the server in addition to any SSH keys specified in this request. | [optional] 
**NetworkType** | Pointer to **string** | The type of network configuration for this server.&lt;br&gt; Currently this field should be set to &#x60;PUBLIC_AND_PRIVATE&#x60;, &#x60;PRIVATE_ONLY&#x60;, &#x60;PUBLIC_ONLY&#x60; or &#x60;USER_DEFINED&#x60;.&lt;br&gt; Setting the &#x60;force&#x60; query parameter to &#x60;true&#x60; allows you to configure network configuration type as &#x60;NONE&#x60;. | [optional] [default to "PUBLIC_AND_PRIVATE"]
**OsConfiguration** | Pointer to [**OsConfiguration**](OsConfiguration.md) |  | [optional] 
**Tags** | Pointer to [**[]TagAssignmentRequest**](TagAssignmentRequest.md) | Tags to set to the server. To create a new tag or list all the existing tags that you can use, refer to [Tags API](https://developers.phoenixnap.com/docs/tags/1/overview). | [optional] 
**NetworkConfiguration** | Pointer to [**NetworkConfiguration**](NetworkConfiguration.md) |  | [optional] 
**StorageConfiguration** | Pointer to [**StorageConfiguration**](StorageConfiguration.md) |  | [optional] 

## Methods

### NewServerProvision

`func NewServerProvision(hostname string, os string, ) *ServerProvision`

NewServerProvision instantiates a new ServerProvision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerProvisionWithDefaults

`func NewServerProvisionWithDefaults() *ServerProvision`

NewServerProvisionWithDefaults instantiates a new ServerProvision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *ServerProvision) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerProvision) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerProvision) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetDescription

`func (o *ServerProvision) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerProvision) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerProvision) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerProvision) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOs

`func (o *ServerProvision) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ServerProvision) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ServerProvision) SetOs(v string)`

SetOs sets Os field to given value.


### GetInstallDefaultSshKeys

`func (o *ServerProvision) GetInstallDefaultSshKeys() bool`

GetInstallDefaultSshKeys returns the InstallDefaultSshKeys field if non-nil, zero value otherwise.

### GetInstallDefaultSshKeysOk

`func (o *ServerProvision) GetInstallDefaultSshKeysOk() (*bool, bool)`

GetInstallDefaultSshKeysOk returns a tuple with the InstallDefaultSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDefaultSshKeys

`func (o *ServerProvision) SetInstallDefaultSshKeys(v bool)`

SetInstallDefaultSshKeys sets InstallDefaultSshKeys field to given value.

### HasInstallDefaultSshKeys

`func (o *ServerProvision) HasInstallDefaultSshKeys() bool`

HasInstallDefaultSshKeys returns a boolean if a field has been set.

### GetSshKeys

`func (o *ServerProvision) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *ServerProvision) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *ServerProvision) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *ServerProvision) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetSshKeyIds

`func (o *ServerProvision) GetSshKeyIds() []string`

GetSshKeyIds returns the SshKeyIds field if non-nil, zero value otherwise.

### GetSshKeyIdsOk

`func (o *ServerProvision) GetSshKeyIdsOk() (*[]string, bool)`

GetSshKeyIdsOk returns a tuple with the SshKeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyIds

`func (o *ServerProvision) SetSshKeyIds(v []string)`

SetSshKeyIds sets SshKeyIds field to given value.

### HasSshKeyIds

`func (o *ServerProvision) HasSshKeyIds() bool`

HasSshKeyIds returns a boolean if a field has been set.

### GetNetworkType

`func (o *ServerProvision) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *ServerProvision) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *ServerProvision) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *ServerProvision) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetOsConfiguration

`func (o *ServerProvision) GetOsConfiguration() OsConfiguration`

GetOsConfiguration returns the OsConfiguration field if non-nil, zero value otherwise.

### GetOsConfigurationOk

`func (o *ServerProvision) GetOsConfigurationOk() (*OsConfiguration, bool)`

GetOsConfigurationOk returns a tuple with the OsConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsConfiguration

`func (o *ServerProvision) SetOsConfiguration(v OsConfiguration)`

SetOsConfiguration sets OsConfiguration field to given value.

### HasOsConfiguration

`func (o *ServerProvision) HasOsConfiguration() bool`

HasOsConfiguration returns a boolean if a field has been set.

### GetTags

`func (o *ServerProvision) GetTags() []TagAssignmentRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerProvision) GetTagsOk() (*[]TagAssignmentRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerProvision) SetTags(v []TagAssignmentRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerProvision) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworkConfiguration

`func (o *ServerProvision) GetNetworkConfiguration() NetworkConfiguration`

GetNetworkConfiguration returns the NetworkConfiguration field if non-nil, zero value otherwise.

### GetNetworkConfigurationOk

`func (o *ServerProvision) GetNetworkConfigurationOk() (*NetworkConfiguration, bool)`

GetNetworkConfigurationOk returns a tuple with the NetworkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfiguration

`func (o *ServerProvision) SetNetworkConfiguration(v NetworkConfiguration)`

SetNetworkConfiguration sets NetworkConfiguration field to given value.

### HasNetworkConfiguration

`func (o *ServerProvision) HasNetworkConfiguration() bool`

HasNetworkConfiguration returns a boolean if a field has been set.

### GetStorageConfiguration

`func (o *ServerProvision) GetStorageConfiguration() StorageConfiguration`

GetStorageConfiguration returns the StorageConfiguration field if non-nil, zero value otherwise.

### GetStorageConfigurationOk

`func (o *ServerProvision) GetStorageConfigurationOk() (*StorageConfiguration, bool)`

GetStorageConfigurationOk returns a tuple with the StorageConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfiguration

`func (o *ServerProvision) SetStorageConfiguration(v StorageConfiguration)`

SetStorageConfiguration sets StorageConfiguration field to given value.

### HasStorageConfiguration

`func (o *ServerProvision) HasStorageConfiguration() bool`

HasStorageConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


