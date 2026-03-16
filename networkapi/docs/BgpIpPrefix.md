# BgpIpPrefix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAllocationId** | **string** | IP allocation ID. | 
**Cidr** | **string** | The IP block in CIDR format, dependent on IP version. | 
**IpVersion** | **string** | The IP block version. | 
**Status** | **string** | The BGP IP Prefix status. Can have one of the following values: &#x60;PENDING&#x60;, &#x60;BUSY&#x60;, &#x60;READY&#x60;, &#x60;ERROR&#x60; and &#x60;DELETING&#x60;. | 

## Methods

### NewBgpIpPrefix

`func NewBgpIpPrefix(ipAllocationId string, cidr string, ipVersion string, status string, ) *BgpIpPrefix`

NewBgpIpPrefix instantiates a new BgpIpPrefix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpIpPrefixWithDefaults

`func NewBgpIpPrefixWithDefaults() *BgpIpPrefix`

NewBgpIpPrefixWithDefaults instantiates a new BgpIpPrefix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAllocationId

`func (o *BgpIpPrefix) GetIpAllocationId() string`

GetIpAllocationId returns the IpAllocationId field if non-nil, zero value otherwise.

### GetIpAllocationIdOk

`func (o *BgpIpPrefix) GetIpAllocationIdOk() (*string, bool)`

GetIpAllocationIdOk returns a tuple with the IpAllocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllocationId

`func (o *BgpIpPrefix) SetIpAllocationId(v string)`

SetIpAllocationId sets IpAllocationId field to given value.


### GetCidr

`func (o *BgpIpPrefix) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *BgpIpPrefix) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *BgpIpPrefix) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetIpVersion

`func (o *BgpIpPrefix) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *BgpIpPrefix) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *BgpIpPrefix) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.


### GetStatus

`func (o *BgpIpPrefix) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BgpIpPrefix) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BgpIpPrefix) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


