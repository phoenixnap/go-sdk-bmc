# IpBlockCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | IP Block location ID. Currently this field should be set to &#x60;PHX&#x60;, &#x60;ASH&#x60;, &#x60;SGP&#x60;, &#x60;NLD&#x60;, &#x60;CHI&#x60; or &#x60;SEA&#x60;. | 
**CidrBlockSize** | **string** | CIDR IP Block Size. Currently this field should be set to either &#x60;/31&#x60;, &#x60;/30&#x60;, &#x60;/29&#x60; or &#x60;/28&#x60;. | 

## Methods

### NewIpBlockCreate

`func NewIpBlockCreate(location string, cidrBlockSize string, ) *IpBlockCreate`

NewIpBlockCreate instantiates a new IpBlockCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockCreateWithDefaults

`func NewIpBlockCreateWithDefaults() *IpBlockCreate`

NewIpBlockCreateWithDefaults instantiates a new IpBlockCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *IpBlockCreate) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IpBlockCreate) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IpBlockCreate) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetCidrBlockSize

`func (o *IpBlockCreate) GetCidrBlockSize() string`

GetCidrBlockSize returns the CidrBlockSize field if non-nil, zero value otherwise.

### GetCidrBlockSizeOk

`func (o *IpBlockCreate) GetCidrBlockSizeOk() (*string, bool)`

GetCidrBlockSizeOk returns a tuple with the CidrBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlockSize

`func (o *IpBlockCreate) SetCidrBlockSize(v string)`

SetCidrBlockSize sets CidrBlockSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


