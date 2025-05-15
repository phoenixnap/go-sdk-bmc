# ServerPrivateNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The network identifier. | 
**Ips** | Pointer to **[]string** | IPs to configure/configured on the server.&lt;br&gt; Valid IP formats are single IPv4 addresses or IPv4 ranges. IPs must be within the network&#39;s range. Should be null or empty list if DHCP is true. &lt;br&gt; If field is undefined and DHCP is false, next available IP in network will be automatically allocated.&lt;br&gt; If the network contains a membership of type &#39;storage&#39;, the first twelve IPs are already reserved by BMC and not usable.&lt;br&gt; Setting the &#x60;force&#x60; query parameter to &#x60;true&#x60; allows you to:&lt;ul&gt; &lt;li&gt; Assign no specific IP addresses by designating an empty array of IPs. Note that at least one IP is required for the gateway address to be selected from this network. &lt;li&gt; Assign one or more IP addresses which are already configured on other resource(s) in network. &lt;li&gt; Assign IP addresses which are considered as reserved in network.&lt;/ul&gt; | [optional] 
**Dhcp** | Pointer to **bool** | Determines whether DHCP is enabled for this server.&lt;br&gt; The following restrictions apply when enabling DHCP:&lt;ul&gt; &lt;li&gt; DHCP support is limited to servers configured exclusively with private networks (PRIVATE_ONLY). &lt;li&gt; DHCP value needs to be consistent across all server-configured private networks.  &lt;li&gt; The server does not support manual gateway address configuration. &lt;li&gt; Private IP addresses for network cannot be specified.&lt;/ul&gt; Note: Not supported on Proxmox OS. | [optional] [default to false]
**StatusDescription** | Pointer to **string** | (Read-only) The status of the network. | [optional] [readonly] 

## Methods

### NewServerPrivateNetwork

`func NewServerPrivateNetwork(id string, ) *ServerPrivateNetwork`

NewServerPrivateNetwork instantiates a new ServerPrivateNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPrivateNetworkWithDefaults

`func NewServerPrivateNetworkWithDefaults() *ServerPrivateNetwork`

NewServerPrivateNetworkWithDefaults instantiates a new ServerPrivateNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerPrivateNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerPrivateNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerPrivateNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetIps

`func (o *ServerPrivateNetwork) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *ServerPrivateNetwork) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *ServerPrivateNetwork) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *ServerPrivateNetwork) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetDhcp

`func (o *ServerPrivateNetwork) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *ServerPrivateNetwork) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *ServerPrivateNetwork) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *ServerPrivateNetwork) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetStatusDescription

`func (o *ServerPrivateNetwork) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *ServerPrivateNetwork) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *ServerPrivateNetwork) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *ServerPrivateNetwork) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


