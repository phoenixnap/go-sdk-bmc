# PrivateNetworkServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The server identifier. | 
**Ips** | **[]string** | List of private IPs associated to the server. | 

## Methods

### NewPrivateNetworkServer

`func NewPrivateNetworkServer(id string, ips []string, ) *PrivateNetworkServer`

NewPrivateNetworkServer instantiates a new PrivateNetworkServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkServerWithDefaults

`func NewPrivateNetworkServerWithDefaults() *PrivateNetworkServer`

NewPrivateNetworkServerWithDefaults instantiates a new PrivateNetworkServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateNetworkServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateNetworkServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateNetworkServer) SetId(v string)`

SetId sets Id field to given value.


### GetIps

`func (o *PrivateNetworkServer) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *PrivateNetworkServer) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *PrivateNetworkServer) SetIps(v []string)`

SetIps sets Ips field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


