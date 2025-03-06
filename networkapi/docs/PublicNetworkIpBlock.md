# PublicNetworkIpBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The IP Block identifier. | 
**Cidr** | **string** | The CIDR notation of the IP block. | 
**UsedIpsCount** | **string** | The number of IPs used in the IP block. | 

## Methods

### NewPublicNetworkIpBlock

`func NewPublicNetworkIpBlock(id string, cidr string, usedIpsCount string, ) *PublicNetworkIpBlock`

NewPublicNetworkIpBlock instantiates a new PublicNetworkIpBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicNetworkIpBlockWithDefaults

`func NewPublicNetworkIpBlockWithDefaults() *PublicNetworkIpBlock`

NewPublicNetworkIpBlockWithDefaults instantiates a new PublicNetworkIpBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicNetworkIpBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicNetworkIpBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicNetworkIpBlock) SetId(v string)`

SetId sets Id field to given value.


### GetCidr

`func (o *PublicNetworkIpBlock) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *PublicNetworkIpBlock) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *PublicNetworkIpBlock) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetUsedIpsCount

`func (o *PublicNetworkIpBlock) GetUsedIpsCount() string`

GetUsedIpsCount returns the UsedIpsCount field if non-nil, zero value otherwise.

### GetUsedIpsCountOk

`func (o *PublicNetworkIpBlock) GetUsedIpsCountOk() (*string, bool)`

GetUsedIpsCountOk returns a tuple with the UsedIpsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedIpsCount

`func (o *PublicNetworkIpBlock) SetUsedIpsCount(v string)`

SetUsedIpsCount sets UsedIpsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


