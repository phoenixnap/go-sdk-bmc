# PrivateNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The private network identifier. | 
**Name** | **string** | The friendly name of this private network. | 
**Description** | Pointer to **string** | The description of this private network. | [optional] 
**VlanId** | **int32** | The VLAN of this private network. | 
**Type** | **string** | The type of the private network. | 
**Location** | **string** | The location of this private network. | 
**LocationDefault** | **bool** | Identifies network as the default private network for the specified location. | 
**Cidr** | Pointer to **string** | IP range associated with this private network in CIDR notation. | [optional] 
**Servers** | [**[]PrivateNetworkServer**](PrivateNetworkServer.md) |  | 
**Memberships** | [**[]NetworkMembership**](NetworkMembership.md) | A list of resources that are members of this private network. | 
**Status** | **string** | The status of the private network. Can have one of the following values: &#x60;BUSY&#x60;, &#x60;READY&#x60;, &#x60;DELETING&#x60; or &#x60;ERROR&#x60;. | 
**CreatedOn** | **time.Time** | Date and time when this private network was created. | 

## Methods

### NewPrivateNetwork

`func NewPrivateNetwork(id string, name string, vlanId int32, type_ string, location string, locationDefault bool, servers []PrivateNetworkServer, memberships []NetworkMembership, status string, createdOn time.Time, ) *PrivateNetwork`

NewPrivateNetwork instantiates a new PrivateNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkWithDefaults

`func NewPrivateNetworkWithDefaults() *PrivateNetwork`

NewPrivateNetworkWithDefaults instantiates a new PrivateNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PrivateNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateNetwork) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PrivateNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrivateNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVlanId

`func (o *PrivateNetwork) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *PrivateNetwork) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *PrivateNetwork) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.


### GetType

`func (o *PrivateNetwork) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrivateNetwork) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrivateNetwork) SetType(v string)`

SetType sets Type field to given value.


### GetLocation

`func (o *PrivateNetwork) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PrivateNetwork) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PrivateNetwork) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetLocationDefault

`func (o *PrivateNetwork) GetLocationDefault() bool`

GetLocationDefault returns the LocationDefault field if non-nil, zero value otherwise.

### GetLocationDefaultOk

`func (o *PrivateNetwork) GetLocationDefaultOk() (*bool, bool)`

GetLocationDefaultOk returns a tuple with the LocationDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationDefault

`func (o *PrivateNetwork) SetLocationDefault(v bool)`

SetLocationDefault sets LocationDefault field to given value.


### GetCidr

`func (o *PrivateNetwork) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *PrivateNetwork) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *PrivateNetwork) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *PrivateNetwork) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetServers

`func (o *PrivateNetwork) GetServers() []PrivateNetworkServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *PrivateNetwork) GetServersOk() (*[]PrivateNetworkServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *PrivateNetwork) SetServers(v []PrivateNetworkServer)`

SetServers sets Servers field to given value.


### GetMemberships

`func (o *PrivateNetwork) GetMemberships() []NetworkMembership`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *PrivateNetwork) GetMembershipsOk() (*[]NetworkMembership, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *PrivateNetwork) SetMemberships(v []NetworkMembership)`

SetMemberships sets Memberships field to given value.


### GetStatus

`func (o *PrivateNetwork) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivateNetwork) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivateNetwork) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedOn

`func (o *PrivateNetwork) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *PrivateNetwork) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *PrivateNetwork) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


