# PublicNetworkMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The resource identifier. | 
**ResourceType** | **string** | The resource&#39;s type. | 
**Ips** | **[]string** | List of public IPs associated to the resource. | 

## Methods

### NewPublicNetworkMembership

`func NewPublicNetworkMembership(resourceId string, resourceType string, ips []string, ) *PublicNetworkMembership`

NewPublicNetworkMembership instantiates a new PublicNetworkMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicNetworkMembershipWithDefaults

`func NewPublicNetworkMembershipWithDefaults() *PublicNetworkMembership`

NewPublicNetworkMembershipWithDefaults instantiates a new PublicNetworkMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *PublicNetworkMembership) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *PublicNetworkMembership) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *PublicNetworkMembership) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *PublicNetworkMembership) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *PublicNetworkMembership) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *PublicNetworkMembership) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetIps

`func (o *PublicNetworkMembership) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *PublicNetworkMembership) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *PublicNetworkMembership) SetIps(v []string)`

SetIps sets Ips field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


