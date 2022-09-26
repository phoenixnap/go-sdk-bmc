# OsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Windows** | Pointer to [**OsConfigurationWindows**](OsConfigurationWindows.md) |  | [optional] 
**RootPassword** | Pointer to **string** | Password set for user root on an ESXi server which will only be returned in response to provisioning a server. | [optional] [readonly] 
**ManagementUiUrl** | Pointer to **string** | The URL of the management UI which will only be returned in response to provisioning a server. | [optional] [readonly] 
**ManagementAccessAllowedIps** | Pointer to **[]string** | List of IPs allowed to access the Management UI. Supported in single IP, CIDR and range format. When undefined, Management UI is disabled. This will only be returned in response to provisioning a server. | [optional] 
**InstallOsToRam** | Pointer to **bool** | If true, OS will be installed to and booted from the server&#39;s RAM. On restart RAM OS will be lost and the server will not be reachable unless a custom bootable OS has been deployed. Only supported for ubuntu/focal. | [optional] [default to false]

## Methods

### NewOsConfiguration

`func NewOsConfiguration() *OsConfiguration`

NewOsConfiguration instantiates a new OsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsConfigurationWithDefaults

`func NewOsConfigurationWithDefaults() *OsConfiguration`

NewOsConfigurationWithDefaults instantiates a new OsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindows

`func (o *OsConfiguration) GetWindows() OsConfigurationWindows`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *OsConfiguration) GetWindowsOk() (*OsConfigurationWindows, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *OsConfiguration) SetWindows(v OsConfigurationWindows)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *OsConfiguration) HasWindows() bool`

HasWindows returns a boolean if a field has been set.

### GetRootPassword

`func (o *OsConfiguration) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *OsConfiguration) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *OsConfiguration) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *OsConfiguration) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetManagementUiUrl

`func (o *OsConfiguration) GetManagementUiUrl() string`

GetManagementUiUrl returns the ManagementUiUrl field if non-nil, zero value otherwise.

### GetManagementUiUrlOk

`func (o *OsConfiguration) GetManagementUiUrlOk() (*string, bool)`

GetManagementUiUrlOk returns a tuple with the ManagementUiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementUiUrl

`func (o *OsConfiguration) SetManagementUiUrl(v string)`

SetManagementUiUrl sets ManagementUiUrl field to given value.

### HasManagementUiUrl

`func (o *OsConfiguration) HasManagementUiUrl() bool`

HasManagementUiUrl returns a boolean if a field has been set.

### GetManagementAccessAllowedIps

`func (o *OsConfiguration) GetManagementAccessAllowedIps() []string`

GetManagementAccessAllowedIps returns the ManagementAccessAllowedIps field if non-nil, zero value otherwise.

### GetManagementAccessAllowedIpsOk

`func (o *OsConfiguration) GetManagementAccessAllowedIpsOk() (*[]string, bool)`

GetManagementAccessAllowedIpsOk returns a tuple with the ManagementAccessAllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAccessAllowedIps

`func (o *OsConfiguration) SetManagementAccessAllowedIps(v []string)`

SetManagementAccessAllowedIps sets ManagementAccessAllowedIps field to given value.

### HasManagementAccessAllowedIps

`func (o *OsConfiguration) HasManagementAccessAllowedIps() bool`

HasManagementAccessAllowedIps returns a boolean if a field has been set.

### GetInstallOsToRam

`func (o *OsConfiguration) GetInstallOsToRam() bool`

GetInstallOsToRam returns the InstallOsToRam field if non-nil, zero value otherwise.

### GetInstallOsToRamOk

`func (o *OsConfiguration) GetInstallOsToRamOk() (*bool, bool)`

GetInstallOsToRamOk returns a tuple with the InstallOsToRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallOsToRam

`func (o *OsConfiguration) SetInstallOsToRam(v bool)`

SetInstallOsToRam sets InstallOsToRam field to given value.

### HasInstallOsToRam

`func (o *OsConfiguration) HasInstallOsToRam() bool`

HasInstallOsToRam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


