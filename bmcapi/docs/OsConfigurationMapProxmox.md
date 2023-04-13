# OsConfigurationMapProxmox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootPassword** | Pointer to **string** | (Read-only) Password set for user root on a Proxmox server which will only be returned in response to provisioning a server. | [optional] [readonly] 
**ManagementUiUrl** | Pointer to **string** | (Read-only) The URL of the management UI which will only be returned in response to provisioning a server. | [optional] [readonly] 
**ManagementAccessAllowedIps** | Pointer to **[]string** | List of IPs allowed to access the Management UI. Supported in single IP, CIDR and range format. When undefined, Management UI is disabled. This will only be returned in response to provisioning a server. | [optional] 

## Methods

### NewOsConfigurationMapProxmox

`func NewOsConfigurationMapProxmox() *OsConfigurationMapProxmox`

NewOsConfigurationMapProxmox instantiates a new OsConfigurationMapProxmox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsConfigurationMapProxmoxWithDefaults

`func NewOsConfigurationMapProxmoxWithDefaults() *OsConfigurationMapProxmox`

NewOsConfigurationMapProxmoxWithDefaults instantiates a new OsConfigurationMapProxmox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootPassword

`func (o *OsConfigurationMapProxmox) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *OsConfigurationMapProxmox) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *OsConfigurationMapProxmox) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *OsConfigurationMapProxmox) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetManagementUiUrl

`func (o *OsConfigurationMapProxmox) GetManagementUiUrl() string`

GetManagementUiUrl returns the ManagementUiUrl field if non-nil, zero value otherwise.

### GetManagementUiUrlOk

`func (o *OsConfigurationMapProxmox) GetManagementUiUrlOk() (*string, bool)`

GetManagementUiUrlOk returns a tuple with the ManagementUiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementUiUrl

`func (o *OsConfigurationMapProxmox) SetManagementUiUrl(v string)`

SetManagementUiUrl sets ManagementUiUrl field to given value.

### HasManagementUiUrl

`func (o *OsConfigurationMapProxmox) HasManagementUiUrl() bool`

HasManagementUiUrl returns a boolean if a field has been set.

### GetManagementAccessAllowedIps

`func (o *OsConfigurationMapProxmox) GetManagementAccessAllowedIps() []string`

GetManagementAccessAllowedIps returns the ManagementAccessAllowedIps field if non-nil, zero value otherwise.

### GetManagementAccessAllowedIpsOk

`func (o *OsConfigurationMapProxmox) GetManagementAccessAllowedIpsOk() (*[]string, bool)`

GetManagementAccessAllowedIpsOk returns a tuple with the ManagementAccessAllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAccessAllowedIps

`func (o *OsConfigurationMapProxmox) SetManagementAccessAllowedIps(v []string)`

SetManagementAccessAllowedIps sets ManagementAccessAllowedIps field to given value.

### HasManagementAccessAllowedIps

`func (o *OsConfigurationMapProxmox) HasManagementAccessAllowedIps() bool`

HasManagementAccessAllowedIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


