# NetworkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayAddress** | Pointer to **string** | The address of the gateway assigned / to assign to the server. When used as part of request body, IP address has to be part of private/public network assigned to this server. | [optional] 
**PrivateNetworkConfiguration** | Pointer to [**PrivateNetworkConfiguration**](PrivateNetworkConfiguration.md) |  | [optional] 
**IpBlocksConfiguration** | Pointer to [**IpBlocksConfiguration**](IpBlocksConfiguration.md) |  | [optional] 
**PublicNetworkConfiguration** | Pointer to [**PublicNetworkConfiguration**](PublicNetworkConfiguration.md) |  | [optional] 

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

### GetGatewayAddress

`func (o *NetworkConfiguration) GetGatewayAddress() string`

GetGatewayAddress returns the GatewayAddress field if non-nil, zero value otherwise.

### GetGatewayAddressOk

`func (o *NetworkConfiguration) GetGatewayAddressOk() (*string, bool)`

GetGatewayAddressOk returns a tuple with the GatewayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddress

`func (o *NetworkConfiguration) SetGatewayAddress(v string)`

SetGatewayAddress sets GatewayAddress field to given value.

### HasGatewayAddress

`func (o *NetworkConfiguration) HasGatewayAddress() bool`

HasGatewayAddress returns a boolean if a field has been set.

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

### GetPublicNetworkConfiguration

`func (o *NetworkConfiguration) GetPublicNetworkConfiguration() PublicNetworkConfiguration`

GetPublicNetworkConfiguration returns the PublicNetworkConfiguration field if non-nil, zero value otherwise.

### GetPublicNetworkConfigurationOk

`func (o *NetworkConfiguration) GetPublicNetworkConfigurationOk() (*PublicNetworkConfiguration, bool)`

GetPublicNetworkConfigurationOk returns a tuple with the PublicNetworkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkConfiguration

`func (o *NetworkConfiguration) SetPublicNetworkConfiguration(v PublicNetworkConfiguration)`

SetPublicNetworkConfiguration sets PublicNetworkConfiguration field to given value.

### HasPublicNetworkConfiguration

`func (o *NetworkConfiguration) HasPublicNetworkConfiguration() bool`

HasPublicNetworkConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


