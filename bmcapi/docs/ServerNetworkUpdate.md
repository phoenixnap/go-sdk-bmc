# ServerNetworkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | Pointer to **[]string** | List of IPs to be associated to the server.&lt;br&gt; Valid IP formats include single IP addresses or IP ranges (IPv4 or IPv6). IPs must be within the network&#39;s range.&lt;br&gt; Setting the &#x60;force&#x60; query parameter to &#x60;true&#x60; allows you to:&lt;ul&gt; &lt;li&gt; Assign no specific IP addresses by designating an empty array of IPs. &lt;li&gt; Assign one or more IP addresses which are already configured on other resource(s) in network. &lt;li&gt; Assign IP addresses which are considered as reserved in network.&lt;/ul&gt; | [optional] 

## Methods

### NewServerNetworkUpdate

`func NewServerNetworkUpdate() *ServerNetworkUpdate`

NewServerNetworkUpdate instantiates a new ServerNetworkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerNetworkUpdateWithDefaults

`func NewServerNetworkUpdateWithDefaults() *ServerNetworkUpdate`

NewServerNetworkUpdateWithDefaults instantiates a new ServerNetworkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *ServerNetworkUpdate) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *ServerNetworkUpdate) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *ServerNetworkUpdate) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *ServerNetworkUpdate) HasIps() bool`

HasIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


