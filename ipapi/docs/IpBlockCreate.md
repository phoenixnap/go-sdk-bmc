# IpBlockCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | IP Block location ID. Currently this field should be set to &#x60;PHX&#x60;, &#x60;ASH&#x60;, &#x60;SGP&#x60;, &#x60;NLD&#x60;, &#x60;CHI&#x60;, &#x60;SEA&#x60; or &#x60;AUS&#x60;. | 
**CidrBlockSize** | **string** | CIDR IP Block Size. V4 supported sizes: [&#x60;/31&#x60;, &#x60;/30&#x60;, &#x60;/29&#x60; or &#x60;/28&#x60;]. V6 supported sizes: [&#x60;/64&#x60;]. For a larger Block Size contact support. | 
**IpVersion** | Pointer to **string** | IP Version. This field should be set to &#x60;V4&#x60; or &#x60;V6&#x60; | [optional] [default to "V4"]
**Description** | Pointer to **string** | The description of the IP Block. | [optional] 
**Tags** | Pointer to [**[]TagAssignmentRequest**](TagAssignmentRequest.md) | Tags to set to the ip-block. To create a new tag or list all the existing tags that you can use, refer to [Tags API](https://developers.phoenixnap.com/docs/tags/1/overview). | [optional] 

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


### GetIpVersion

`func (o *IpBlockCreate) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *IpBlockCreate) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *IpBlockCreate) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *IpBlockCreate) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetDescription

`func (o *IpBlockCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpBlockCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpBlockCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpBlockCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *IpBlockCreate) GetTags() []TagAssignmentRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpBlockCreate) GetTagsOk() (*[]TagAssignmentRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpBlockCreate) SetTags(v []TagAssignmentRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpBlockCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


