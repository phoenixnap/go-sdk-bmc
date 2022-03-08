# IpBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | IP Block identifier. | 
**Location** | **string** | IP Block location ID. Currently this field should be set to &#x60;PHX&#x60;, &#x60;ASH&#x60;, &#x60;SGP&#x60;, &#x60;NLD&#x60;, &#x60;CHI&#x60;, &#x60;SEA&#x60; or &#x60;AUS&#x60;. | 
**CidrBlockSize** | **string** | CIDR IP Block Size. Currently this field should be set to either &#x60;/31&#x60;, &#x60;/30&#x60;, &#x60;/29&#x60; or &#x60;/28&#x60;. | 
**Cidr** | **string** | The IP Block in CIDR notation. | 
**Status** | **string** | The status of the IP Block. | 
**AssignedResourceId** | Pointer to **string** | ID of the resource assigned to the IP Block. | [optional] 
**AssignedResourceType** | Pointer to **string** | Type of the resource assigned to the IP Block. | [optional] 
**Description** | Pointer to **string** | The description of the IP Block. | [optional] 

## Methods

### NewIpBlock

`func NewIpBlock(id string, location string, cidrBlockSize string, cidr string, status string, ) *IpBlock`

NewIpBlock instantiates a new IpBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockWithDefaults

`func NewIpBlockWithDefaults() *IpBlock`

NewIpBlockWithDefaults instantiates a new IpBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IpBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpBlock) SetId(v string)`

SetId sets Id field to given value.


### GetLocation

`func (o *IpBlock) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IpBlock) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IpBlock) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetCidrBlockSize

`func (o *IpBlock) GetCidrBlockSize() string`

GetCidrBlockSize returns the CidrBlockSize field if non-nil, zero value otherwise.

### GetCidrBlockSizeOk

`func (o *IpBlock) GetCidrBlockSizeOk() (*string, bool)`

GetCidrBlockSizeOk returns a tuple with the CidrBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlockSize

`func (o *IpBlock) SetCidrBlockSize(v string)`

SetCidrBlockSize sets CidrBlockSize field to given value.


### GetCidr

`func (o *IpBlock) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *IpBlock) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *IpBlock) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetStatus

`func (o *IpBlock) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IpBlock) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IpBlock) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAssignedResourceId

`func (o *IpBlock) GetAssignedResourceId() string`

GetAssignedResourceId returns the AssignedResourceId field if non-nil, zero value otherwise.

### GetAssignedResourceIdOk

`func (o *IpBlock) GetAssignedResourceIdOk() (*string, bool)`

GetAssignedResourceIdOk returns a tuple with the AssignedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedResourceId

`func (o *IpBlock) SetAssignedResourceId(v string)`

SetAssignedResourceId sets AssignedResourceId field to given value.

### HasAssignedResourceId

`func (o *IpBlock) HasAssignedResourceId() bool`

HasAssignedResourceId returns a boolean if a field has been set.

### GetAssignedResourceType

`func (o *IpBlock) GetAssignedResourceType() string`

GetAssignedResourceType returns the AssignedResourceType field if non-nil, zero value otherwise.

### GetAssignedResourceTypeOk

`func (o *IpBlock) GetAssignedResourceTypeOk() (*string, bool)`

GetAssignedResourceTypeOk returns a tuple with the AssignedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedResourceType

`func (o *IpBlock) SetAssignedResourceType(v string)`

SetAssignedResourceType sets AssignedResourceType field to given value.

### HasAssignedResourceType

`func (o *IpBlock) HasAssignedResourceType() bool`

HasAssignedResourceType returns a boolean if a field has been set.

### GetDescription

`func (o *IpBlock) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpBlock) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpBlock) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpBlock) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


