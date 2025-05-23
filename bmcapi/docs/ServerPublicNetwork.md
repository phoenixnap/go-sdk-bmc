# ServerPublicNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The network identifier. | 
**Ips** | Pointer to **[]string** | Configurable/configured IPs on the server.&lt;br&gt; At least 1 IP address is required. Valid IP format is single IP addresses. All IPs must be within the network&#39;s range.&lt;br&gt; Setting the &#x60;computeSlaacIp&#x60; field to &#x60;true&#x60; allows you to provide an empty array of IPs.&lt;br&gt; Additionally, setting the &#x60;force&#x60; query parameter to &#x60;true&#x60; allows you to:&lt;ul&gt; &lt;li&gt; Assign no specific IP addresses by designating an empty array of IPs. Note that at least one IP is required for the gateway address to be selected from this network. &lt;li&gt; Assign one or more IP addresses which are already configured on other resource(s) in network.&lt;/ul&gt; | [optional] 
**StatusDescription** | Pointer to **string** | (Read-only) The status of the assignment to the network. | [optional] [readonly] 
**ComputeSlaacIp** | Pointer to **bool** | (Write-only) Requests Stateless Address Autoconfiguration (SLAAC). Applicable for Network which contains IPv6 block(s). | [optional] 

## Methods

### NewServerPublicNetwork

`func NewServerPublicNetwork(id string, ) *ServerPublicNetwork`

NewServerPublicNetwork instantiates a new ServerPublicNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPublicNetworkWithDefaults

`func NewServerPublicNetworkWithDefaults() *ServerPublicNetwork`

NewServerPublicNetworkWithDefaults instantiates a new ServerPublicNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerPublicNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerPublicNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerPublicNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetIps

`func (o *ServerPublicNetwork) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *ServerPublicNetwork) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *ServerPublicNetwork) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *ServerPublicNetwork) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetStatusDescription

`func (o *ServerPublicNetwork) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *ServerPublicNetwork) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *ServerPublicNetwork) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *ServerPublicNetwork) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetComputeSlaacIp

`func (o *ServerPublicNetwork) GetComputeSlaacIp() bool`

GetComputeSlaacIp returns the ComputeSlaacIp field if non-nil, zero value otherwise.

### GetComputeSlaacIpOk

`func (o *ServerPublicNetwork) GetComputeSlaacIpOk() (*bool, bool)`

GetComputeSlaacIpOk returns a tuple with the ComputeSlaacIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeSlaacIp

`func (o *ServerPublicNetwork) SetComputeSlaacIp(v bool)`

SetComputeSlaacIp sets ComputeSlaacIp field to given value.

### HasComputeSlaacIp

`func (o *ServerPublicNetwork) HasComputeSlaacIp() bool`

HasComputeSlaacIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


