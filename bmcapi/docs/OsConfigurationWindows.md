# OsConfigurationWindows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RdpAllowedIps** | Pointer to **[]string** | List of IPs allowed for RDP access to Windows OS. Supported in single IP, CIDR and range format. When undefined, RDP is disabled. To allow RDP access from any IP use 0.0.0.0/0. This will only be returned in response to provisioning a server. | [optional] 
**BringYourOwnLicense** | Pointer to **bool** | Use a Bring Your Own (BYO) Windows license.  If true, the server is provisioned in trial mode, and you must activate your own license.  If false (default), the server includes a managed Windows license billed by the platform.  | [optional] [default to false]

## Methods

### NewOsConfigurationWindows

`func NewOsConfigurationWindows() *OsConfigurationWindows`

NewOsConfigurationWindows instantiates a new OsConfigurationWindows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsConfigurationWindowsWithDefaults

`func NewOsConfigurationWindowsWithDefaults() *OsConfigurationWindows`

NewOsConfigurationWindowsWithDefaults instantiates a new OsConfigurationWindows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRdpAllowedIps

`func (o *OsConfigurationWindows) GetRdpAllowedIps() []string`

GetRdpAllowedIps returns the RdpAllowedIps field if non-nil, zero value otherwise.

### GetRdpAllowedIpsOk

`func (o *OsConfigurationWindows) GetRdpAllowedIpsOk() (*[]string, bool)`

GetRdpAllowedIpsOk returns a tuple with the RdpAllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpAllowedIps

`func (o *OsConfigurationWindows) SetRdpAllowedIps(v []string)`

SetRdpAllowedIps sets RdpAllowedIps field to given value.

### HasRdpAllowedIps

`func (o *OsConfigurationWindows) HasRdpAllowedIps() bool`

HasRdpAllowedIps returns a boolean if a field has been set.

### GetBringYourOwnLicense

`func (o *OsConfigurationWindows) GetBringYourOwnLicense() bool`

GetBringYourOwnLicense returns the BringYourOwnLicense field if non-nil, zero value otherwise.

### GetBringYourOwnLicenseOk

`func (o *OsConfigurationWindows) GetBringYourOwnLicenseOk() (*bool, bool)`

GetBringYourOwnLicenseOk returns a tuple with the BringYourOwnLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBringYourOwnLicense

`func (o *OsConfigurationWindows) SetBringYourOwnLicense(v bool)`

SetBringYourOwnLicense sets BringYourOwnLicense field to given value.

### HasBringYourOwnLicense

`func (o *OsConfigurationWindows) HasBringYourOwnLicense() bool`

HasBringYourOwnLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


