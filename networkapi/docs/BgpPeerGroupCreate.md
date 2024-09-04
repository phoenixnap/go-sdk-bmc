# BgpPeerGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | The BGP Peer Group location. Can have one of the following values: &#x60;PHX&#x60;, &#x60;ASH&#x60;, &#x60;SGP&#x60;, &#x60;NLD&#x60;, &#x60;CHI&#x60;, &#x60;SEA&#x60; and &#x60;AUS&#x60;. | 
**Asn** | **int64** | The BGP Peer Group ASN. | [default to 65401]
**Password** | Pointer to **string** | The BGP Peer Group password. | [optional] 
**AdvertisedRoutes** | **string** | The Advertised routes for the BGP Peer Group. Can have one of the following values: &#x60;DEFAULT&#x60; and &#x60;NONE&#x60;. | [default to "NONE"]

## Methods

### NewBgpPeerGroupCreate

`func NewBgpPeerGroupCreate(location string, asn int64, advertisedRoutes string, ) *BgpPeerGroupCreate`

NewBgpPeerGroupCreate instantiates a new BgpPeerGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpPeerGroupCreateWithDefaults

`func NewBgpPeerGroupCreateWithDefaults() *BgpPeerGroupCreate`

NewBgpPeerGroupCreateWithDefaults instantiates a new BgpPeerGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *BgpPeerGroupCreate) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BgpPeerGroupCreate) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BgpPeerGroupCreate) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetAsn

`func (o *BgpPeerGroupCreate) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *BgpPeerGroupCreate) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *BgpPeerGroupCreate) SetAsn(v int64)`

SetAsn sets Asn field to given value.


### GetPassword

`func (o *BgpPeerGroupCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BgpPeerGroupCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BgpPeerGroupCreate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BgpPeerGroupCreate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAdvertisedRoutes

`func (o *BgpPeerGroupCreate) GetAdvertisedRoutes() string`

GetAdvertisedRoutes returns the AdvertisedRoutes field if non-nil, zero value otherwise.

### GetAdvertisedRoutesOk

`func (o *BgpPeerGroupCreate) GetAdvertisedRoutesOk() (*string, bool)`

GetAdvertisedRoutesOk returns a tuple with the AdvertisedRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisedRoutes

`func (o *BgpPeerGroupCreate) SetAdvertisedRoutes(v string)`

SetAdvertisedRoutes sets AdvertisedRoutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


