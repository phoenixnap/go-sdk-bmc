# PublicNetworkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicNetworks** | Pointer to [**[]ServerPublicNetwork**](ServerPublicNetwork.md) | The list of public networks this server is member of. When this field is part of request body, it&#39;ll be used to specify the public networks to assign to this server upon provisioning. | [optional] 

## Methods

### NewPublicNetworkConfiguration

`func NewPublicNetworkConfiguration() *PublicNetworkConfiguration`

NewPublicNetworkConfiguration instantiates a new PublicNetworkConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicNetworkConfigurationWithDefaults

`func NewPublicNetworkConfigurationWithDefaults() *PublicNetworkConfiguration`

NewPublicNetworkConfigurationWithDefaults instantiates a new PublicNetworkConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicNetworks

`func (o *PublicNetworkConfiguration) GetPublicNetworks() []ServerPublicNetwork`

GetPublicNetworks returns the PublicNetworks field if non-nil, zero value otherwise.

### GetPublicNetworksOk

`func (o *PublicNetworkConfiguration) GetPublicNetworksOk() (*[]ServerPublicNetwork, bool)`

GetPublicNetworksOk returns a tuple with the PublicNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworks

`func (o *PublicNetworkConfiguration) SetPublicNetworks(v []ServerPublicNetwork)`

SetPublicNetworks sets PublicNetworks field to given value.

### HasPublicNetworks

`func (o *PublicNetworkConfiguration) HasPublicNetworks() bool`

HasPublicNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


