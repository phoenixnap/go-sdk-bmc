# ServerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | Hostname of server. | 
**Description** | Pointer to **string** | Description of server. | [optional] 
**Os** | **string** | The server’s OS ID used when the server was created. Currently this field should be set to either &#x60;ubuntu/bionic&#x60;, &#x60;ubuntu/focal&#x60;, &#x60;centos/centos7&#x60;,&#x60;centos/centos8&#x60;, &#x60;windows/srv2019std&#x60;, &#x60;windows/srv2019dc&#x60;, &#x60;esxi/esxi70u2&#x60;, &#x60;debian/bullseye&#x60; or &#x60;proxmox/bullseye&#x60;. | 
**Type** | **string** | Server type ID. Cannot be changed once a server is created. Currently this field should be set to either &#x60;s0.d1.small&#x60;, &#x60;s0.d1.medium&#x60;, &#x60;s1.c1.small&#x60;, &#x60;s1.c1.medium&#x60;, &#x60;s1.c2.medium&#x60;, &#x60;s1.c2.large&#x60;, &#x60;s1.e1.small&#x60;, &#x60;s1.e1.medium&#x60;, &#x60;s1.e1.large&#x60;, &#x60;s2.c1.small&#x60;, &#x60;s2.c1.medium&#x60;, &#x60;s2.c1.large&#x60;, &#x60;s2.c2.small&#x60;, &#x60;s2.c2.medium&#x60;, &#x60;s2.c2.large&#x60;, &#x60;d1.c1.small&#x60;, &#x60;d1.c2.small&#x60;, &#x60;d1.c3.small&#x60;, &#x60;d1.c4.small&#x60;, &#x60;d1.c1.medium&#x60;, &#x60;d1.c2.medium&#x60;, &#x60;d1.c3.medium&#x60;, &#x60;d1.c4.medium&#x60;, &#x60;d1.c1.large&#x60;, &#x60;d1.c2.large&#x60;, &#x60;d1.c3.large&#x60;, &#x60;d1.c4.large&#x60;, &#x60;d1.m1.medium&#x60;, &#x60;d1.m2.medium&#x60;, &#x60;d1.m3.medium&#x60;, &#x60;d1.m4.medium&#x60;, &#x60;d2.c1.medium&#x60;, &#x60;d2.c2.medium&#x60;, &#x60;d2.c3.medium&#x60;, &#x60;d2.c4.medium&#x60;, &#x60;d2.c5.medium&#x60;, &#x60;d2.c1.large&#x60;, &#x60;d2.c2.large&#x60;, &#x60;d2.c3.large&#x60;, &#x60;d2.c4.large&#x60;, &#x60;d2.c5.large&#x60;, &#x60;d2.m1.medium&#x60;, &#x60;d2.m1.large&#x60;, &#x60;d2.m2.medium&#x60;, &#x60;d2.m2.large&#x60;, &#x60;d2.m2.xlarge&#x60; or &#x60;d2.c4.storage.pliops1&#x60;. | 
**Location** | **string** | Server location ID. Cannot be changed once a server is created. Currently this field should be set to &#x60;PHX&#x60;, &#x60;ASH&#x60;, &#x60;SGP&#x60;, &#x60;NLD&#x60;, &#x60;CHI&#x60;, &#x60;SEA&#x60; or &#x60;AUS&#x60;. | 
**InstallDefaultSshKeys** | Pointer to **bool** | Whether or not to install SSH keys marked as default in addition to any SSH keys specified in this request. | [optional] [default to true]
**SshKeys** | Pointer to **[]string** | A list of SSH keys that will be installed on the server. | [optional] 
**SshKeyIds** | Pointer to **[]string** | A list of SSH key IDs that will be installed on the server in addition to any SSH keys specified in this request. | [optional] 
**ReservationId** | Pointer to **string** | Server reservation ID. | [optional] 
**PricingModel** | Pointer to **string** | Server pricing model. Currently this field should be set to &#x60;HOURLY&#x60;, &#x60;ONE_MONTH_RESERVATION&#x60;, &#x60;TWELVE_MONTHS_RESERVATION&#x60;, &#x60;TWENTY_FOUR_MONTHS_RESERVATION&#x60; or &#x60;THIRTY_SIX_MONTHS_RESERVATION&#x60;. | [optional] [default to "HOURLY"]
**NetworkType** | Pointer to **string** | The type of network configuration for this server. Currently this field should be set to &#x60;PUBLIC_AND_PRIVATE&#x60; or &#x60;PRIVATE_ONLY&#x60;. | [optional] [default to "PUBLIC_AND_PRIVATE"]
**OsConfiguration** | Pointer to [**OsConfiguration**](OsConfiguration.md) |  | [optional] 
**Tags** | Pointer to [**[]TagAssignmentRequest**](TagAssignmentRequest.md) | Tags to set to server, if any. | [optional] 
**NetworkConfiguration** | Pointer to [**NetworkConfiguration**](NetworkConfiguration.md) |  | [optional] 

## Methods

### NewServerCreate

`func NewServerCreate(hostname string, os string, type_ string, location string, ) *ServerCreate`

NewServerCreate instantiates a new ServerCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerCreateWithDefaults

`func NewServerCreateWithDefaults() *ServerCreate`

NewServerCreateWithDefaults instantiates a new ServerCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *ServerCreate) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerCreate) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerCreate) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetDescription

`func (o *ServerCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOs

`func (o *ServerCreate) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ServerCreate) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ServerCreate) SetOs(v string)`

SetOs sets Os field to given value.


### GetType

`func (o *ServerCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerCreate) SetType(v string)`

SetType sets Type field to given value.


### GetLocation

`func (o *ServerCreate) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ServerCreate) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ServerCreate) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetInstallDefaultSshKeys

`func (o *ServerCreate) GetInstallDefaultSshKeys() bool`

GetInstallDefaultSshKeys returns the InstallDefaultSshKeys field if non-nil, zero value otherwise.

### GetInstallDefaultSshKeysOk

`func (o *ServerCreate) GetInstallDefaultSshKeysOk() (*bool, bool)`

GetInstallDefaultSshKeysOk returns a tuple with the InstallDefaultSshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDefaultSshKeys

`func (o *ServerCreate) SetInstallDefaultSshKeys(v bool)`

SetInstallDefaultSshKeys sets InstallDefaultSshKeys field to given value.

### HasInstallDefaultSshKeys

`func (o *ServerCreate) HasInstallDefaultSshKeys() bool`

HasInstallDefaultSshKeys returns a boolean if a field has been set.

### GetSshKeys

`func (o *ServerCreate) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *ServerCreate) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *ServerCreate) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *ServerCreate) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetSshKeyIds

`func (o *ServerCreate) GetSshKeyIds() []string`

GetSshKeyIds returns the SshKeyIds field if non-nil, zero value otherwise.

### GetSshKeyIdsOk

`func (o *ServerCreate) GetSshKeyIdsOk() (*[]string, bool)`

GetSshKeyIdsOk returns a tuple with the SshKeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyIds

`func (o *ServerCreate) SetSshKeyIds(v []string)`

SetSshKeyIds sets SshKeyIds field to given value.

### HasSshKeyIds

`func (o *ServerCreate) HasSshKeyIds() bool`

HasSshKeyIds returns a boolean if a field has been set.

### GetReservationId

`func (o *ServerCreate) GetReservationId() string`

GetReservationId returns the ReservationId field if non-nil, zero value otherwise.

### GetReservationIdOk

`func (o *ServerCreate) GetReservationIdOk() (*string, bool)`

GetReservationIdOk returns a tuple with the ReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationId

`func (o *ServerCreate) SetReservationId(v string)`

SetReservationId sets ReservationId field to given value.

### HasReservationId

`func (o *ServerCreate) HasReservationId() bool`

HasReservationId returns a boolean if a field has been set.

### GetPricingModel

`func (o *ServerCreate) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *ServerCreate) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *ServerCreate) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.

### HasPricingModel

`func (o *ServerCreate) HasPricingModel() bool`

HasPricingModel returns a boolean if a field has been set.

### GetNetworkType

`func (o *ServerCreate) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *ServerCreate) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *ServerCreate) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *ServerCreate) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetOsConfiguration

`func (o *ServerCreate) GetOsConfiguration() OsConfiguration`

GetOsConfiguration returns the OsConfiguration field if non-nil, zero value otherwise.

### GetOsConfigurationOk

`func (o *ServerCreate) GetOsConfigurationOk() (*OsConfiguration, bool)`

GetOsConfigurationOk returns a tuple with the OsConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsConfiguration

`func (o *ServerCreate) SetOsConfiguration(v OsConfiguration)`

SetOsConfiguration sets OsConfiguration field to given value.

### HasOsConfiguration

`func (o *ServerCreate) HasOsConfiguration() bool`

HasOsConfiguration returns a boolean if a field has been set.

### GetTags

`func (o *ServerCreate) GetTags() []TagAssignmentRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerCreate) GetTagsOk() (*[]TagAssignmentRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerCreate) SetTags(v []TagAssignmentRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworkConfiguration

`func (o *ServerCreate) GetNetworkConfiguration() NetworkConfiguration`

GetNetworkConfiguration returns the NetworkConfiguration field if non-nil, zero value otherwise.

### GetNetworkConfigurationOk

`func (o *ServerCreate) GetNetworkConfigurationOk() (*NetworkConfiguration, bool)`

GetNetworkConfigurationOk returns a tuple with the NetworkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfiguration

`func (o *ServerCreate) SetNetworkConfiguration(v NetworkConfiguration)`

SetNetworkConfiguration sets NetworkConfiguration field to given value.

### HasNetworkConfiguration

`func (o *ServerCreate) HasNetworkConfiguration() bool`

HasNetworkConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


