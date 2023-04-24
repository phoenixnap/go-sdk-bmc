# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the server. | 
**Status** | **string** | The status of the server. | 
**Hostname** | **string** | Hostname of server. | 
**Description** | Pointer to **string** | Description of server. | [optional] 
**Os** | **string** | The server’s OS ID used when the server was created. Currently this field should be set to either &#x60;ubuntu/bionic&#x60;, &#x60;ubuntu/focal&#x60;, &#x60;ubuntu/jammy&#x60;, &#x60;centos/centos7&#x60;, &#x60;centos/centos8&#x60;, &#x60;windows/srv2019std&#x60;, &#x60;windows/srv2019dc&#x60;, &#x60;esxi/esxi70&#x60;, &#x60;esxi/esxi80&#x60;, &#x60;debian/bullseye&#x60;, &#x60;proxmox/bullseye&#x60;, &#x60;netris/controller&#x60;, &#x60;netris/softgate_1g&#x60; or &#x60;netris/softgate_10g&#x60;. | 
**Type** | **string** | Server type ID. Cannot be changed once a server is created. Currently this field should be set to either &#x60;s0.d1.small&#x60;, &#x60;s0.d1.medium&#x60;, &#x60;s1.c1.small&#x60;, &#x60;s1.c1.medium&#x60;, &#x60;s1.c2.medium&#x60;, &#x60;s1.c2.large&#x60;, &#x60;s1.e1.small&#x60;, &#x60;s1.e1.medium&#x60;, &#x60;s1.e1.large&#x60;, &#x60;s2.c1.small&#x60;, &#x60;s2.c1.medium&#x60;, &#x60;s2.c1.large&#x60;, &#x60;s2.c2.small&#x60;, &#x60;s2.c2.medium&#x60;, &#x60;s2.c2.large&#x60;, &#x60;d1.c1.small&#x60;, &#x60;d1.c2.small&#x60;, &#x60;d1.c3.small&#x60;, &#x60;d1.c4.small&#x60;, &#x60;d1.c1.medium&#x60;, &#x60;d1.c2.medium&#x60;, &#x60;d1.c3.medium&#x60;, &#x60;d1.c4.medium&#x60;, &#x60;d1.c1.large&#x60;, &#x60;d1.c2.large&#x60;, &#x60;d1.c3.large&#x60;, &#x60;d1.c4.large&#x60;, &#x60;d1.m1.medium&#x60;, &#x60;d1.m2.medium&#x60;, &#x60;d1.m3.medium&#x60;, &#x60;d1.m4.medium&#x60;, &#x60;d2.c1.medium&#x60;, &#x60;d2.c2.medium&#x60;, &#x60;d2.c3.medium&#x60;, &#x60;d2.c4.medium&#x60;, &#x60;d2.c5.medium&#x60;, &#x60;d2.c1.large&#x60;, &#x60;d2.c2.large&#x60;, &#x60;d2.c3.large&#x60;, &#x60;d2.c4.large&#x60;, &#x60;d2.c5.large&#x60;, &#x60;d2.m1.medium&#x60;, &#x60;d2.m1.large&#x60;, &#x60;d2.m2.medium&#x60;, &#x60;d2.m2.large&#x60;, &#x60;d2.m2.xlarge&#x60;, &#x60;d2.c4.storage.pliops1&#x60;, &#x60;d3.m4.xlarge&#x60;, &#x60;d3.m5.xlarge&#x60; or &#x60;d3.m6.xlarge&#x60;. | 
**Location** | **string** | Server location ID. Cannot be changed once a server is created. Currently this field should be set to &#x60;PHX&#x60;, &#x60;ASH&#x60;, &#x60;SGP&#x60;, &#x60;NLD&#x60;, &#x60;CHI&#x60;, &#x60;SEA&#x60; or &#x60;AUS&#x60;. | 
**Cpu** | **string** | A description of the machine CPU. | 
**CpuCount** | **int32** | The number of CPUs available in the system. | 
**CoresPerCpu** | **int32** | The number of physical cores present on each CPU. | 
**CpuFrequency** | **float32** | The CPU frequency in GHz. | 
**Ram** | **string** | A description of the machine RAM. | 
**Storage** | **string** | A description of the machine storage. | 
**PrivateIpAddresses** | **[]string** | Private IP addresses assigned to server. | 
**PublicIpAddresses** | Pointer to **[]string** | Public IP addresses assigned to server. | [optional] 
**ReservationId** | Pointer to **string** | The reservation reference id if any. | [optional] 
**PricingModel** | **string** | The pricing model this server is being billed. Currently this field should be set to &#x60;HOURLY&#x60;, &#x60;ONE_MONTH_RESERVATION&#x60;, &#x60;TWELVE_MONTHS_RESERVATION&#x60;, &#x60;TWENTY_FOUR_MONTHS_RESERVATION&#x60; or &#x60;THIRTY_SIX_MONTHS_RESERVATION&#x60;. | [default to "HOURLY"]
**Password** | Pointer to **string** | Auto-generated password set for user &#x60;Admin&#x60; on Windows server, user &#x60;root&#x60; on ESXi servers, user &#x60;root&#x60; on Proxmox server and user &#x60;netris&#x60; on Netris servers.&lt;br&gt; The password is not stored and therefore will only be returned in response to provisioning a server. Copy and save it for future reference. | [optional] 
**NetworkType** | Pointer to **string** | The type of network configuration for this server. Currently this field should be set to &#x60;PUBLIC_AND_PRIVATE&#x60;, &#x60;PRIVATE_ONLY&#x60;, &#x60;PUBLIC_ONLY&#x60; or &#x60;NONE&#x60;. | [optional] [default to "PUBLIC_AND_PRIVATE"]
**ClusterId** | Pointer to **string** | The cluster reference id if any. | [optional] 
**Tags** | Pointer to [**[]TagAssignment**](TagAssignment.md) | The tags assigned if any. | [optional] 
**ProvisionedOn** | Pointer to **time.Time** | Date and time when server was provisioned. | [optional] 
**OsConfiguration** | Pointer to [**OsConfiguration**](OsConfiguration.md) |  | [optional] 
**NetworkConfiguration** | [**NetworkConfiguration**](NetworkConfiguration.md) |  | 

## Methods

### NewServer

`func NewServer(id string, status string, hostname string, os string, type_ string, location string, cpu string, cpuCount int32, coresPerCpu int32, cpuFrequency float32, ram string, storage string, privateIpAddresses []string, pricingModel string, networkConfiguration NetworkConfiguration, ) *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Server) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Server) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Server) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *Server) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Server) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Server) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetHostname

`func (o *Server) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Server) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Server) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetDescription

`func (o *Server) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Server) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Server) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Server) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOs

`func (o *Server) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Server) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Server) SetOs(v string)`

SetOs sets Os field to given value.


### GetType

`func (o *Server) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Server) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Server) SetType(v string)`

SetType sets Type field to given value.


### GetLocation

`func (o *Server) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Server) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Server) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetCpu

`func (o *Server) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Server) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Server) SetCpu(v string)`

SetCpu sets Cpu field to given value.


### GetCpuCount

`func (o *Server) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *Server) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *Server) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.


### GetCoresPerCpu

`func (o *Server) GetCoresPerCpu() int32`

GetCoresPerCpu returns the CoresPerCpu field if non-nil, zero value otherwise.

### GetCoresPerCpuOk

`func (o *Server) GetCoresPerCpuOk() (*int32, bool)`

GetCoresPerCpuOk returns a tuple with the CoresPerCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerCpu

`func (o *Server) SetCoresPerCpu(v int32)`

SetCoresPerCpu sets CoresPerCpu field to given value.


### GetCpuFrequency

`func (o *Server) GetCpuFrequency() float32`

GetCpuFrequency returns the CpuFrequency field if non-nil, zero value otherwise.

### GetCpuFrequencyOk

`func (o *Server) GetCpuFrequencyOk() (*float32, bool)`

GetCpuFrequencyOk returns a tuple with the CpuFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFrequency

`func (o *Server) SetCpuFrequency(v float32)`

SetCpuFrequency sets CpuFrequency field to given value.


### GetRam

`func (o *Server) GetRam() string`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *Server) GetRamOk() (*string, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *Server) SetRam(v string)`

SetRam sets Ram field to given value.


### GetStorage

`func (o *Server) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Server) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Server) SetStorage(v string)`

SetStorage sets Storage field to given value.


### GetPrivateIpAddresses

`func (o *Server) GetPrivateIpAddresses() []string`

GetPrivateIpAddresses returns the PrivateIpAddresses field if non-nil, zero value otherwise.

### GetPrivateIpAddressesOk

`func (o *Server) GetPrivateIpAddressesOk() (*[]string, bool)`

GetPrivateIpAddressesOk returns a tuple with the PrivateIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpAddresses

`func (o *Server) SetPrivateIpAddresses(v []string)`

SetPrivateIpAddresses sets PrivateIpAddresses field to given value.


### GetPublicIpAddresses

`func (o *Server) GetPublicIpAddresses() []string`

GetPublicIpAddresses returns the PublicIpAddresses field if non-nil, zero value otherwise.

### GetPublicIpAddressesOk

`func (o *Server) GetPublicIpAddressesOk() (*[]string, bool)`

GetPublicIpAddressesOk returns a tuple with the PublicIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddresses

`func (o *Server) SetPublicIpAddresses(v []string)`

SetPublicIpAddresses sets PublicIpAddresses field to given value.

### HasPublicIpAddresses

`func (o *Server) HasPublicIpAddresses() bool`

HasPublicIpAddresses returns a boolean if a field has been set.

### GetReservationId

`func (o *Server) GetReservationId() string`

GetReservationId returns the ReservationId field if non-nil, zero value otherwise.

### GetReservationIdOk

`func (o *Server) GetReservationIdOk() (*string, bool)`

GetReservationIdOk returns a tuple with the ReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationId

`func (o *Server) SetReservationId(v string)`

SetReservationId sets ReservationId field to given value.

### HasReservationId

`func (o *Server) HasReservationId() bool`

HasReservationId returns a boolean if a field has been set.

### GetPricingModel

`func (o *Server) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *Server) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *Server) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetPassword

`func (o *Server) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Server) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Server) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Server) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetNetworkType

`func (o *Server) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *Server) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *Server) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *Server) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetClusterId

`func (o *Server) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Server) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Server) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *Server) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetTags

`func (o *Server) GetTags() []TagAssignment`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Server) GetTagsOk() (*[]TagAssignment, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Server) SetTags(v []TagAssignment)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Server) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProvisionedOn

`func (o *Server) GetProvisionedOn() time.Time`

GetProvisionedOn returns the ProvisionedOn field if non-nil, zero value otherwise.

### GetProvisionedOnOk

`func (o *Server) GetProvisionedOnOk() (*time.Time, bool)`

GetProvisionedOnOk returns a tuple with the ProvisionedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedOn

`func (o *Server) SetProvisionedOn(v time.Time)`

SetProvisionedOn sets ProvisionedOn field to given value.

### HasProvisionedOn

`func (o *Server) HasProvisionedOn() bool`

HasProvisionedOn returns a boolean if a field has been set.

### GetOsConfiguration

`func (o *Server) GetOsConfiguration() OsConfiguration`

GetOsConfiguration returns the OsConfiguration field if non-nil, zero value otherwise.

### GetOsConfigurationOk

`func (o *Server) GetOsConfigurationOk() (*OsConfiguration, bool)`

GetOsConfigurationOk returns a tuple with the OsConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsConfiguration

`func (o *Server) SetOsConfiguration(v OsConfiguration)`

SetOsConfiguration sets OsConfiguration field to given value.

### HasOsConfiguration

`func (o *Server) HasOsConfiguration() bool`

HasOsConfiguration returns a boolean if a field has been set.

### GetNetworkConfiguration

`func (o *Server) GetNetworkConfiguration() NetworkConfiguration`

GetNetworkConfiguration returns the NetworkConfiguration field if non-nil, zero value otherwise.

### GetNetworkConfigurationOk

`func (o *Server) GetNetworkConfigurationOk() (*NetworkConfiguration, bool)`

GetNetworkConfigurationOk returns a tuple with the NetworkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfiguration

`func (o *Server) SetNetworkConfiguration(v NetworkConfiguration)`

SetNetworkConfiguration sets NetworkConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


