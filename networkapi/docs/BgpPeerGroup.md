# BgpPeerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the BGP Peer Group. | 
**Status** | **string** | The BGP Peer Group status. Can have one of the following values: &#x60;PENDING&#x60;, &#x60;ON_HOLD&#x60;, &#x60;BUSY&#x60;, &#x60;READY&#x60;, &#x60;ERROR&#x60;, &#x60;PENDING_DELETION&#x60; and &#x60;DELETING&#x60;. | 
**Location** | **string** | The BGP Peer Group location. Can have one of the following values: &#x60;PHX&#x60;, &#x60;ASH&#x60;, &#x60;SGP&#x60;, &#x60;NLD&#x60;, &#x60;CHI&#x60;, &#x60;SEA&#x60; and &#x60;AUS&#x60;. | 
**Ipv4Prefixes** | [**[]BgpIPv4Prefix**](BgpIPv4Prefix.md) | The List of the BGP Peer Group IPv4 prefixes. | 
**TargetAsnDetails** | [**AsnDetails**](AsnDetails.md) |  | 
**ActiveAsnDetails** | Pointer to [**AsnDetails**](AsnDetails.md) |  | [optional] 
**Password** | **string** | The BGP Peer Group password. | 
**AdvertisedRoutes** | **string** | The Advertised routes for the BGP Peer Group. Can have one of the following values: &#x60;DEFAULT&#x60; and &#x60;NONE&#x60;. | 
**RpkiRoaOriginAsn** | **int64** | The RPKI ROA Origin ASN of the BGP Peer Group based on location. | 
**EBgpMultiHop** | **int32** | The eBGP Multi-hop of the BGP Peer Group. | 
**PeeringLoopbacksV4** | **[]string** | The IPv4 Peering Loopback addresses of the BGP Peer Group. Valid IP formats are IPv4 addresses. | 
**KeepAliveTimerSeconds** | **int32** | The Keep Alive Timer in seconds of the BGP Peer Group. | 
**HoldTimerSeconds** | **int32** | The Hold Timer in seconds of the BGP Peer Group. | 
**CreatedOn** | Pointer to **string** | Date and time of creation. | [optional] 
**LastUpdatedOn** | Pointer to **string** | Date and time of last update. | [optional] 

## Methods

### NewBgpPeerGroup

`func NewBgpPeerGroup(id string, status string, location string, ipv4Prefixes []BgpIPv4Prefix, targetAsnDetails AsnDetails, password string, advertisedRoutes string, rpkiRoaOriginAsn int64, eBgpMultiHop int32, peeringLoopbacksV4 []string, keepAliveTimerSeconds int32, holdTimerSeconds int32, ) *BgpPeerGroup`

NewBgpPeerGroup instantiates a new BgpPeerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpPeerGroupWithDefaults

`func NewBgpPeerGroupWithDefaults() *BgpPeerGroup`

NewBgpPeerGroupWithDefaults instantiates a new BgpPeerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BgpPeerGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BgpPeerGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BgpPeerGroup) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *BgpPeerGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BgpPeerGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BgpPeerGroup) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLocation

`func (o *BgpPeerGroup) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BgpPeerGroup) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BgpPeerGroup) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetIpv4Prefixes

`func (o *BgpPeerGroup) GetIpv4Prefixes() []BgpIPv4Prefix`

GetIpv4Prefixes returns the Ipv4Prefixes field if non-nil, zero value otherwise.

### GetIpv4PrefixesOk

`func (o *BgpPeerGroup) GetIpv4PrefixesOk() (*[]BgpIPv4Prefix, bool)`

GetIpv4PrefixesOk returns a tuple with the Ipv4Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Prefixes

`func (o *BgpPeerGroup) SetIpv4Prefixes(v []BgpIPv4Prefix)`

SetIpv4Prefixes sets Ipv4Prefixes field to given value.


### GetTargetAsnDetails

`func (o *BgpPeerGroup) GetTargetAsnDetails() AsnDetails`

GetTargetAsnDetails returns the TargetAsnDetails field if non-nil, zero value otherwise.

### GetTargetAsnDetailsOk

`func (o *BgpPeerGroup) GetTargetAsnDetailsOk() (*AsnDetails, bool)`

GetTargetAsnDetailsOk returns a tuple with the TargetAsnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAsnDetails

`func (o *BgpPeerGroup) SetTargetAsnDetails(v AsnDetails)`

SetTargetAsnDetails sets TargetAsnDetails field to given value.


### GetActiveAsnDetails

`func (o *BgpPeerGroup) GetActiveAsnDetails() AsnDetails`

GetActiveAsnDetails returns the ActiveAsnDetails field if non-nil, zero value otherwise.

### GetActiveAsnDetailsOk

`func (o *BgpPeerGroup) GetActiveAsnDetailsOk() (*AsnDetails, bool)`

GetActiveAsnDetailsOk returns a tuple with the ActiveAsnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAsnDetails

`func (o *BgpPeerGroup) SetActiveAsnDetails(v AsnDetails)`

SetActiveAsnDetails sets ActiveAsnDetails field to given value.

### HasActiveAsnDetails

`func (o *BgpPeerGroup) HasActiveAsnDetails() bool`

HasActiveAsnDetails returns a boolean if a field has been set.

### GetPassword

`func (o *BgpPeerGroup) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BgpPeerGroup) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BgpPeerGroup) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAdvertisedRoutes

`func (o *BgpPeerGroup) GetAdvertisedRoutes() string`

GetAdvertisedRoutes returns the AdvertisedRoutes field if non-nil, zero value otherwise.

### GetAdvertisedRoutesOk

`func (o *BgpPeerGroup) GetAdvertisedRoutesOk() (*string, bool)`

GetAdvertisedRoutesOk returns a tuple with the AdvertisedRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisedRoutes

`func (o *BgpPeerGroup) SetAdvertisedRoutes(v string)`

SetAdvertisedRoutes sets AdvertisedRoutes field to given value.


### GetRpkiRoaOriginAsn

`func (o *BgpPeerGroup) GetRpkiRoaOriginAsn() int64`

GetRpkiRoaOriginAsn returns the RpkiRoaOriginAsn field if non-nil, zero value otherwise.

### GetRpkiRoaOriginAsnOk

`func (o *BgpPeerGroup) GetRpkiRoaOriginAsnOk() (*int64, bool)`

GetRpkiRoaOriginAsnOk returns a tuple with the RpkiRoaOriginAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpkiRoaOriginAsn

`func (o *BgpPeerGroup) SetRpkiRoaOriginAsn(v int64)`

SetRpkiRoaOriginAsn sets RpkiRoaOriginAsn field to given value.


### GetEBgpMultiHop

`func (o *BgpPeerGroup) GetEBgpMultiHop() int32`

GetEBgpMultiHop returns the EBgpMultiHop field if non-nil, zero value otherwise.

### GetEBgpMultiHopOk

`func (o *BgpPeerGroup) GetEBgpMultiHopOk() (*int32, bool)`

GetEBgpMultiHopOk returns a tuple with the EBgpMultiHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBgpMultiHop

`func (o *BgpPeerGroup) SetEBgpMultiHop(v int32)`

SetEBgpMultiHop sets EBgpMultiHop field to given value.


### GetPeeringLoopbacksV4

`func (o *BgpPeerGroup) GetPeeringLoopbacksV4() []string`

GetPeeringLoopbacksV4 returns the PeeringLoopbacksV4 field if non-nil, zero value otherwise.

### GetPeeringLoopbacksV4Ok

`func (o *BgpPeerGroup) GetPeeringLoopbacksV4Ok() (*[]string, bool)`

GetPeeringLoopbacksV4Ok returns a tuple with the PeeringLoopbacksV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeringLoopbacksV4

`func (o *BgpPeerGroup) SetPeeringLoopbacksV4(v []string)`

SetPeeringLoopbacksV4 sets PeeringLoopbacksV4 field to given value.


### GetKeepAliveTimerSeconds

`func (o *BgpPeerGroup) GetKeepAliveTimerSeconds() int32`

GetKeepAliveTimerSeconds returns the KeepAliveTimerSeconds field if non-nil, zero value otherwise.

### GetKeepAliveTimerSecondsOk

`func (o *BgpPeerGroup) GetKeepAliveTimerSecondsOk() (*int32, bool)`

GetKeepAliveTimerSecondsOk returns a tuple with the KeepAliveTimerSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAliveTimerSeconds

`func (o *BgpPeerGroup) SetKeepAliveTimerSeconds(v int32)`

SetKeepAliveTimerSeconds sets KeepAliveTimerSeconds field to given value.


### GetHoldTimerSeconds

`func (o *BgpPeerGroup) GetHoldTimerSeconds() int32`

GetHoldTimerSeconds returns the HoldTimerSeconds field if non-nil, zero value otherwise.

### GetHoldTimerSecondsOk

`func (o *BgpPeerGroup) GetHoldTimerSecondsOk() (*int32, bool)`

GetHoldTimerSecondsOk returns a tuple with the HoldTimerSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTimerSeconds

`func (o *BgpPeerGroup) SetHoldTimerSeconds(v int32)`

SetHoldTimerSeconds sets HoldTimerSeconds field to given value.


### GetCreatedOn

`func (o *BgpPeerGroup) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *BgpPeerGroup) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *BgpPeerGroup) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *BgpPeerGroup) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetLastUpdatedOn

`func (o *BgpPeerGroup) GetLastUpdatedOn() string`

GetLastUpdatedOn returns the LastUpdatedOn field if non-nil, zero value otherwise.

### GetLastUpdatedOnOk

`func (o *BgpPeerGroup) GetLastUpdatedOnOk() (*string, bool)`

GetLastUpdatedOnOk returns a tuple with the LastUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedOn

`func (o *BgpPeerGroup) SetLastUpdatedOn(v string)`

SetLastUpdatedOn sets LastUpdatedOn field to given value.

### HasLastUpdatedOn

`func (o *BgpPeerGroup) HasLastUpdatedOn() bool`

HasLastUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


