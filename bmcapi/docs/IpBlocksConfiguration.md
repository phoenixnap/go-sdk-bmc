# IpBlocksConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationType** | Pointer to **string** | (Write-only) Determines the approach for configuring IP blocks for the server being provisioned. If PURCHASE_NEW is selected, the smallest supported range, depending on the operating system, is allocated to the server. | [optional] [default to "PURCHASE_NEW"]
**IpBlocks** | Pointer to [**[]ServerIpBlock**](ServerIpBlock.md) | Used to specify the previously purchased IP blocks to assign to this server upon provisioning. Used alongside the USER_DEFINED configurationType. | [optional] 

## Methods

### NewIpBlocksConfiguration

`func NewIpBlocksConfiguration() *IpBlocksConfiguration`

NewIpBlocksConfiguration instantiates a new IpBlocksConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlocksConfigurationWithDefaults

`func NewIpBlocksConfigurationWithDefaults() *IpBlocksConfiguration`

NewIpBlocksConfigurationWithDefaults instantiates a new IpBlocksConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationType

`func (o *IpBlocksConfiguration) GetConfigurationType() string`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *IpBlocksConfiguration) GetConfigurationTypeOk() (*string, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *IpBlocksConfiguration) SetConfigurationType(v string)`

SetConfigurationType sets ConfigurationType field to given value.

### HasConfigurationType

`func (o *IpBlocksConfiguration) HasConfigurationType() bool`

HasConfigurationType returns a boolean if a field has been set.

### GetIpBlocks

`func (o *IpBlocksConfiguration) GetIpBlocks() []ServerIpBlock`

GetIpBlocks returns the IpBlocks field if non-nil, zero value otherwise.

### GetIpBlocksOk

`func (o *IpBlocksConfiguration) GetIpBlocksOk() (*[]ServerIpBlock, bool)`

GetIpBlocksOk returns a tuple with the IpBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlocks

`func (o *IpBlocksConfiguration) SetIpBlocks(v []ServerIpBlock)`

SetIpBlocks sets IpBlocks field to given value.

### HasIpBlocks

`func (o *IpBlocksConfiguration) HasIpBlocks() bool`

HasIpBlocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


