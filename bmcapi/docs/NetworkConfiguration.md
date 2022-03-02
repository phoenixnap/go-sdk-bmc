# NetworkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateNetworkConfiguration** | Pointer to [**PrivateNetworkConfiguration**](PrivateNetworkConfiguration.md) |  | [optional] 
**IpBlocksConfiguration** | Pointer to [**IpBlocksConfiguration**](IpBlocksConfiguration.md) |  | [optional] 

## Methods

### NewNetworkConfiguration

`func NewNetworkConfiguration() *NetworkConfiguration`

NewNetworkConfiguration instantiates a new NetworkConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConfigurationWithDefaults

`func NewNetworkConfigurationWithDefaults() *NetworkConfiguration`

NewNetworkConfigurationWithDefaults instantiates a new NetworkConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateNetworkConfiguration

`func (o *NetworkConfiguration) GetPrivateNetworkConfiguration() PrivateNetworkConfiguration`

GetPrivateNetworkConfiguration returns the PrivateNetworkConfiguration field if non-nil, zero value otherwise.

### GetPrivateNetworkConfigurationOk

`func (o *NetworkConfiguration) GetPrivateNetworkConfigurationOk() (*PrivateNetworkConfiguration, bool)`

GetPrivateNetworkConfigurationOk returns a tuple with the PrivateNetworkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkConfiguration

`func (o *NetworkConfiguration) SetPrivateNetworkConfiguration(v PrivateNetworkConfiguration)`

SetPrivateNetworkConfiguration sets PrivateNetworkConfiguration field to given value.

### HasPrivateNetworkConfiguration

`func (o *NetworkConfiguration) HasPrivateNetworkConfiguration() bool`

HasPrivateNetworkConfiguration returns a boolean if a field has been set.

### GetIpBlocksConfiguration

`func (o *NetworkConfiguration) GetIpBlocksConfiguration() IpBlocksConfiguration`

GetIpBlocksConfiguration returns the IpBlocksConfiguration field if non-nil, zero value otherwise.

### GetIpBlocksConfigurationOk

`func (o *NetworkConfiguration) GetIpBlocksConfigurationOk() (*IpBlocksConfiguration, bool)`

GetIpBlocksConfigurationOk returns a tuple with the IpBlocksConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlocksConfiguration

`func (o *NetworkConfiguration) SetIpBlocksConfiguration(v IpBlocksConfiguration)`

SetIpBlocksConfiguration sets IpBlocksConfiguration field to given value.

### HasIpBlocksConfiguration

`func (o *NetworkConfiguration) HasIpBlocksConfiguration() bool`

HasIpBlocksConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


