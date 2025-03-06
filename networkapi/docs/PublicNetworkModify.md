# PublicNetworkModify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A friendly name given to the network. This name should be unique. | [optional] 
**Description** | Pointer to **string** | The description of this public network | [optional] 
**RaEnabled** | Pointer to **bool** | Boolean indicating whether Router Advertisement is enabled. Only applicable for Network with IPv6 Blocks. | [optional] 

## Methods

### NewPublicNetworkModify

`func NewPublicNetworkModify() *PublicNetworkModify`

NewPublicNetworkModify instantiates a new PublicNetworkModify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicNetworkModifyWithDefaults

`func NewPublicNetworkModifyWithDefaults() *PublicNetworkModify`

NewPublicNetworkModifyWithDefaults instantiates a new PublicNetworkModify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PublicNetworkModify) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicNetworkModify) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicNetworkModify) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicNetworkModify) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PublicNetworkModify) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicNetworkModify) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicNetworkModify) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublicNetworkModify) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRaEnabled

`func (o *PublicNetworkModify) GetRaEnabled() bool`

GetRaEnabled returns the RaEnabled field if non-nil, zero value otherwise.

### GetRaEnabledOk

`func (o *PublicNetworkModify) GetRaEnabledOk() (*bool, bool)`

GetRaEnabledOk returns a tuple with the RaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaEnabled

`func (o *PublicNetworkModify) SetRaEnabled(v bool)`

SetRaEnabled sets RaEnabled field to given value.

### HasRaEnabled

`func (o *PublicNetworkModify) HasRaEnabled() bool`

HasRaEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


