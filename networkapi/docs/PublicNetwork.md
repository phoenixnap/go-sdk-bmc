# PublicNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The public network identifier. | 
**VlanId** | **int32** | The VLAN of this public network. | 
**Memberships** | [**[]PublicNetworkMembership**](PublicNetworkMembership.md) | A list of resources that are members in this public network. | 
**Name** | **string** | The friendly name of this public network. | 
**Location** | **string** | The location of this public network. Supported values are &#x60;PHX&#x60;, &#x60;ASH&#x60;, &#x60;SGP&#x60;, &#x60;NLD&#x60;, &#x60;CHI&#x60;, &#x60;SEA&#x60; and &#x60;AUS&#x60;. | 
**Description** | Pointer to **string** | The description of this public network. | [optional] 
**CreatedOn** | **time.Time** | Date and time when this public network was created. | 
**IpBlocks** | [**[]PublicNetworkIpBlock**](PublicNetworkIpBlock.md) | A list of IP Blocks that are associated with this public network. | 

## Methods

### NewPublicNetwork

`func NewPublicNetwork(id string, vlanId int32, memberships []PublicNetworkMembership, name string, location string, createdOn time.Time, ipBlocks []PublicNetworkIpBlock, ) *PublicNetwork`

NewPublicNetwork instantiates a new PublicNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicNetworkWithDefaults

`func NewPublicNetworkWithDefaults() *PublicNetwork`

NewPublicNetworkWithDefaults instantiates a new PublicNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetVlanId

`func (o *PublicNetwork) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *PublicNetwork) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *PublicNetwork) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.


### GetMemberships

`func (o *PublicNetwork) GetMemberships() []PublicNetworkMembership`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *PublicNetwork) GetMembershipsOk() (*[]PublicNetworkMembership, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *PublicNetwork) SetMemberships(v []PublicNetworkMembership)`

SetMemberships sets Memberships field to given value.


### GetName

`func (o *PublicNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicNetwork) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *PublicNetwork) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PublicNetwork) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PublicNetwork) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetDescription

`func (o *PublicNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublicNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedOn

`func (o *PublicNetwork) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *PublicNetwork) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *PublicNetwork) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### GetIpBlocks

`func (o *PublicNetwork) GetIpBlocks() []PublicNetworkIpBlock`

GetIpBlocks returns the IpBlocks field if non-nil, zero value otherwise.

### GetIpBlocksOk

`func (o *PublicNetwork) GetIpBlocksOk() (*[]PublicNetworkIpBlock, bool)`

GetIpBlocksOk returns a tuple with the IpBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlocks

`func (o *PublicNetwork) SetIpBlocks(v []PublicNetworkIpBlock)`

SetIpBlocks sets IpBlocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


