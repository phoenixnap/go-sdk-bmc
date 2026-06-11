# OsConfigurationIPXENativeVlanConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanId** | Pointer to **int32** | The VLAN ID of the network to be used as the native VLAN. The value must reference a public network with IP V4 block(s) or a public IP V4 block network to which the server is (or will be) attached. If omitted during provisioning, the native VLAN is matched to the configured/auto-purchased public IP V4 block. If no public IP block is available, a VLAN ID must be provided. The VLAN ID must belong to one of the public networks for any of the specified servers. During post-provisioning, if Native VLAN is omitted, the server will be configured with no native VLAN. If provided, the VLAN ID must be specified and must belong to any of the existing server public networks or IP block networks attached to the server.  | [optional] 
**StaticDhcpAddressV4** | Pointer to **string** | The static IP V4 address assigned to the server within the native VLAN. This address is set as the DHCP reservation and used for the iPXE boot process.  Value must be an available/unused IP V4 address within the native network usable IP range. If omitted, the first available IP in the native network will be automatically assigned. Therefore, at least one IP must be available within the native network.  | [optional] 
**Status** | Pointer to **string** | (Read-only) The status of the native VLAN configuration. | [optional] [readonly] 

## Methods

### NewOsConfigurationIPXENativeVlanConfiguration

`func NewOsConfigurationIPXENativeVlanConfiguration() *OsConfigurationIPXENativeVlanConfiguration`

NewOsConfigurationIPXENativeVlanConfiguration instantiates a new OsConfigurationIPXENativeVlanConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsConfigurationIPXENativeVlanConfigurationWithDefaults

`func NewOsConfigurationIPXENativeVlanConfigurationWithDefaults() *OsConfigurationIPXENativeVlanConfiguration`

NewOsConfigurationIPXENativeVlanConfigurationWithDefaults instantiates a new OsConfigurationIPXENativeVlanConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlanId

`func (o *OsConfigurationIPXENativeVlanConfiguration) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *OsConfigurationIPXENativeVlanConfiguration) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *OsConfigurationIPXENativeVlanConfiguration) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *OsConfigurationIPXENativeVlanConfiguration) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetStaticDhcpAddressV4

`func (o *OsConfigurationIPXENativeVlanConfiguration) GetStaticDhcpAddressV4() string`

GetStaticDhcpAddressV4 returns the StaticDhcpAddressV4 field if non-nil, zero value otherwise.

### GetStaticDhcpAddressV4Ok

`func (o *OsConfigurationIPXENativeVlanConfiguration) GetStaticDhcpAddressV4Ok() (*string, bool)`

GetStaticDhcpAddressV4Ok returns a tuple with the StaticDhcpAddressV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticDhcpAddressV4

`func (o *OsConfigurationIPXENativeVlanConfiguration) SetStaticDhcpAddressV4(v string)`

SetStaticDhcpAddressV4 sets StaticDhcpAddressV4 field to given value.

### HasStaticDhcpAddressV4

`func (o *OsConfigurationIPXENativeVlanConfiguration) HasStaticDhcpAddressV4() bool`

HasStaticDhcpAddressV4 returns a boolean if a field has been set.

### GetStatus

`func (o *OsConfigurationIPXENativeVlanConfiguration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OsConfigurationIPXENativeVlanConfiguration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OsConfigurationIPXENativeVlanConfiguration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OsConfigurationIPXENativeVlanConfiguration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


