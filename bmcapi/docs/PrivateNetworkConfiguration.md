# PrivateNetworkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayAddress** | Pointer to **string** | The address of the gateway assigned / to assign to the server. When used as part of request body, it has to match one of the IP addresses used in the existing assigned private networks for the relevant location. Deprecated in favour of a common gateway address across all networks available under NetworkConfiguration. | [optional] 
**ConfigurationType** | Pointer to **string** | Determines the approach for configuring private network(s) for the server being provisioned. Currently this field should be set to &#x60;USE_OR_CREATE_DEFAULT&#x60; or &#x60;USER_DEFINED&#x60;. | [optional] [default to "USE_OR_CREATE_DEFAULT"]
**PrivateNetworks** | Pointer to [**[]ServerPrivateNetwork**](ServerPrivateNetwork.md) | The list of private networks this server is member of. When this field is part of request body, it&#39;ll be used to specify the private networks to assign to this server upon provisioning. Used alongside the &#x60;USER_DEFINED&#x60; configurationType. | [optional] 

## Methods

### NewPrivateNetworkConfiguration

`func NewPrivateNetworkConfiguration() *PrivateNetworkConfiguration`

NewPrivateNetworkConfiguration instantiates a new PrivateNetworkConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkConfigurationWithDefaults

`func NewPrivateNetworkConfigurationWithDefaults() *PrivateNetworkConfiguration`

NewPrivateNetworkConfigurationWithDefaults instantiates a new PrivateNetworkConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayAddress

`func (o *PrivateNetworkConfiguration) GetGatewayAddress() string`

GetGatewayAddress returns the GatewayAddress field if non-nil, zero value otherwise.

### GetGatewayAddressOk

`func (o *PrivateNetworkConfiguration) GetGatewayAddressOk() (*string, bool)`

GetGatewayAddressOk returns a tuple with the GatewayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddress

`func (o *PrivateNetworkConfiguration) SetGatewayAddress(v string)`

SetGatewayAddress sets GatewayAddress field to given value.

### HasGatewayAddress

`func (o *PrivateNetworkConfiguration) HasGatewayAddress() bool`

HasGatewayAddress returns a boolean if a field has been set.

### GetConfigurationType

`func (o *PrivateNetworkConfiguration) GetConfigurationType() string`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *PrivateNetworkConfiguration) GetConfigurationTypeOk() (*string, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *PrivateNetworkConfiguration) SetConfigurationType(v string)`

SetConfigurationType sets ConfigurationType field to given value.

### HasConfigurationType

`func (o *PrivateNetworkConfiguration) HasConfigurationType() bool`

HasConfigurationType returns a boolean if a field has been set.

### GetPrivateNetworks

`func (o *PrivateNetworkConfiguration) GetPrivateNetworks() []ServerPrivateNetwork`

GetPrivateNetworks returns the PrivateNetworks field if non-nil, zero value otherwise.

### GetPrivateNetworksOk

`func (o *PrivateNetworkConfiguration) GetPrivateNetworksOk() (*[]ServerPrivateNetwork, bool)`

GetPrivateNetworksOk returns a tuple with the PrivateNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworks

`func (o *PrivateNetworkConfiguration) SetPrivateNetworks(v []ServerPrivateNetwork)`

SetPrivateNetworks sets PrivateNetworks field to given value.

### HasPrivateNetworks

`func (o *PrivateNetworkConfiguration) HasPrivateNetworks() bool`

HasPrivateNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


