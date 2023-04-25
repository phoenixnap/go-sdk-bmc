# ServerIpBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The IP block&#39;s ID. | 
**VlanId** | Pointer to **int32** | (Read-only) The VLAN on which this IP block has been configured within the network switch. | [optional] [readonly] 

## Methods

### NewServerIpBlock

`func NewServerIpBlock(id string, ) *ServerIpBlock`

NewServerIpBlock instantiates a new ServerIpBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerIpBlockWithDefaults

`func NewServerIpBlockWithDefaults() *ServerIpBlock`

NewServerIpBlockWithDefaults instantiates a new ServerIpBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerIpBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerIpBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerIpBlock) SetId(v string)`

SetId sets Id field to given value.


### GetVlanId

`func (o *ServerIpBlock) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ServerIpBlock) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ServerIpBlock) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ServerIpBlock) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


