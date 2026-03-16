# BgpPeerGroupPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **int64** | The BGP Peer Group ASN. | [optional] 
**Password** | Pointer to **string** | The BGP Peer Group password. | [optional] 
**AdvertisedRoutes** | Pointer to **string** | The Advertised routes for the BGP Peer Group. Can have one of the following values: &#x60;DEFAULT&#x60; and &#x60;NONE&#x60;. | [optional] 

## Methods

### NewBgpPeerGroupPatch

`func NewBgpPeerGroupPatch() *BgpPeerGroupPatch`

NewBgpPeerGroupPatch instantiates a new BgpPeerGroupPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpPeerGroupPatchWithDefaults

`func NewBgpPeerGroupPatchWithDefaults() *BgpPeerGroupPatch`

NewBgpPeerGroupPatchWithDefaults instantiates a new BgpPeerGroupPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *BgpPeerGroupPatch) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *BgpPeerGroupPatch) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *BgpPeerGroupPatch) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *BgpPeerGroupPatch) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetPassword

`func (o *BgpPeerGroupPatch) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BgpPeerGroupPatch) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BgpPeerGroupPatch) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BgpPeerGroupPatch) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAdvertisedRoutes

`func (o *BgpPeerGroupPatch) GetAdvertisedRoutes() string`

GetAdvertisedRoutes returns the AdvertisedRoutes field if non-nil, zero value otherwise.

### GetAdvertisedRoutesOk

`func (o *BgpPeerGroupPatch) GetAdvertisedRoutesOk() (*string, bool)`

GetAdvertisedRoutesOk returns a tuple with the AdvertisedRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisedRoutes

`func (o *BgpPeerGroupPatch) SetAdvertisedRoutes(v string)`

SetAdvertisedRoutes sets AdvertisedRoutes field to given value.

### HasAdvertisedRoutes

`func (o *BgpPeerGroupPatch) HasAdvertisedRoutes() bool`

HasAdvertisedRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


