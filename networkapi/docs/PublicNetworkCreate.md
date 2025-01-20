# PublicNetworkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The friendly name of this public network. This name should be unique. | 
**Description** | Pointer to **string** | The description of this public network. | [optional] 
**Location** | **string** | The location of this public network. Supported values are &#x60;PHX&#x60;, &#x60;ASH&#x60;, &#x60;SGP&#x60;, &#x60;NLD&#x60;, &#x60;CHI&#x60;, &#x60;SEA&#x60; and &#x60;AUS&#x60;. | 
**VlanId** | Pointer to **int32** | The VLAN that will be assigned to this network. | [optional] 
**IpBlocks** | Pointer to [**[]PublicNetworkIpBlockCreate**](PublicNetworkIpBlockCreate.md) | A list of IP Blocks that will be associated with this public network. | [optional] 

## Methods

### NewPublicNetworkCreate

`func NewPublicNetworkCreate(name string, location string, ) *PublicNetworkCreate`

NewPublicNetworkCreate instantiates a new PublicNetworkCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicNetworkCreateWithDefaults

`func NewPublicNetworkCreateWithDefaults() *PublicNetworkCreate`

NewPublicNetworkCreateWithDefaults instantiates a new PublicNetworkCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PublicNetworkCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicNetworkCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicNetworkCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PublicNetworkCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicNetworkCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicNetworkCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublicNetworkCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *PublicNetworkCreate) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PublicNetworkCreate) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PublicNetworkCreate) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetVlanId

`func (o *PublicNetworkCreate) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *PublicNetworkCreate) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *PublicNetworkCreate) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *PublicNetworkCreate) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetIpBlocks

`func (o *PublicNetworkCreate) GetIpBlocks() []PublicNetworkIpBlockCreate`

GetIpBlocks returns the IpBlocks field if non-nil, zero value otherwise.

### GetIpBlocksOk

`func (o *PublicNetworkCreate) GetIpBlocksOk() (*[]PublicNetworkIpBlockCreate, bool)`

GetIpBlocksOk returns a tuple with the IpBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlocks

`func (o *PublicNetworkCreate) SetIpBlocks(v []PublicNetworkIpBlockCreate)`

SetIpBlocks sets IpBlocks field to given value.

### HasIpBlocks

`func (o *PublicNetworkCreate) HasIpBlocks() bool`

HasIpBlocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


