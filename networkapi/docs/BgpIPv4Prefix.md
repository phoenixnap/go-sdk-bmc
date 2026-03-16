# BgpIPv4Prefix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4AllocationId** | **string** | IPv4 allocation ID. | 
**Cidr** | **string** | The IP block in CIDR format. | 
**Status** | **string** | The BGP IPv4 Prefix status. Can have one of the following values: &#x60;PENDING&#x60;, &#x60;BUSY&#x60;, &#x60;READY&#x60;, &#x60;ERROR&#x60; and &#x60;DELETING&#x60;. | 
**IsBringYourOwnIp** | **bool** | Identifies IP as a &#x60;bring your own&#x60; IP block. | 
**InUse** | **bool** | The Boolean value of the BGP IPv4 Prefix is in use. | 

## Methods

### NewBgpIPv4Prefix

`func NewBgpIPv4Prefix(ipv4AllocationId string, cidr string, status string, isBringYourOwnIp bool, inUse bool, ) *BgpIPv4Prefix`

NewBgpIPv4Prefix instantiates a new BgpIPv4Prefix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpIPv4PrefixWithDefaults

`func NewBgpIPv4PrefixWithDefaults() *BgpIPv4Prefix`

NewBgpIPv4PrefixWithDefaults instantiates a new BgpIPv4Prefix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4AllocationId

`func (o *BgpIPv4Prefix) GetIpv4AllocationId() string`

GetIpv4AllocationId returns the Ipv4AllocationId field if non-nil, zero value otherwise.

### GetIpv4AllocationIdOk

`func (o *BgpIPv4Prefix) GetIpv4AllocationIdOk() (*string, bool)`

GetIpv4AllocationIdOk returns a tuple with the Ipv4AllocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4AllocationId

`func (o *BgpIPv4Prefix) SetIpv4AllocationId(v string)`

SetIpv4AllocationId sets Ipv4AllocationId field to given value.


### GetCidr

`func (o *BgpIPv4Prefix) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *BgpIPv4Prefix) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *BgpIPv4Prefix) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetStatus

`func (o *BgpIPv4Prefix) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BgpIPv4Prefix) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BgpIPv4Prefix) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIsBringYourOwnIp

`func (o *BgpIPv4Prefix) GetIsBringYourOwnIp() bool`

GetIsBringYourOwnIp returns the IsBringYourOwnIp field if non-nil, zero value otherwise.

### GetIsBringYourOwnIpOk

`func (o *BgpIPv4Prefix) GetIsBringYourOwnIpOk() (*bool, bool)`

GetIsBringYourOwnIpOk returns a tuple with the IsBringYourOwnIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBringYourOwnIp

`func (o *BgpIPv4Prefix) SetIsBringYourOwnIp(v bool)`

SetIsBringYourOwnIp sets IsBringYourOwnIp field to given value.


### GetInUse

`func (o *BgpIPv4Prefix) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *BgpIPv4Prefix) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *BgpIPv4Prefix) SetInUse(v bool)`

SetInUse sets InUse field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


