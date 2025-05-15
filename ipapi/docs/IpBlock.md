# IpBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | IP Block identifier. | [optional] 
**Location** | Pointer to **string** | IP Block location ID. Currently this field should be set to &#x60;PHX&#x60;, &#x60;ASH&#x60;, &#x60;SGP&#x60;, &#x60;NLD&#x60;, &#x60;CHI&#x60;, &#x60;SEA&#x60; or &#x60;AUS&#x60;. | [optional] 
**CidrBlockSize** | Pointer to **string** | CIDR IP Block Size. Currently this field should be set to either &#x60;/31&#x60;, &#x60;/30&#x60;, &#x60;/29&#x60;, &#x60;/28&#x60;, &#x60;/27&#x60;, &#x60;/26&#x60;, &#x60;/25&#x60;, &#x60;/24&#x60;, &#x60;/23&#x60; or &#x60;/22&#x60;. | [optional] 
**Cidr** | Pointer to **string** | The IP Block in CIDR notation. | [optional] 
**IpVersion** | Pointer to **string** | The IP Version of the block. | [optional] 
**Status** | Pointer to **string** | The status of the IP Block. Can have one of the following values: &#x60;creating&#x60;, &#x60;subnetted&#x60;, &#x60;assigning&#x60; , &#x60;error assigning&#x60; , &#x60;assigned&#x60; , &#x60;unassigning&#x60; , &#x60;error unassigning&#x60; or &#x60;unassigned&#x60;. | [optional] 
**ParentIpBlockAllocationId** | Pointer to **string** | IP Block parent identifier. If present, this block is subnetted from the parent IP Block. | [optional] 
**AssignedResourceId** | Pointer to **string** | ID of the resource assigned to the IP Block. | [optional] 
**AssignedResourceType** | Pointer to **string** | Type of the resource assigned to the IP Block. | [optional] 
**Description** | Pointer to **string** | The description of the IP Block. | [optional] 
**Tags** | Pointer to [**[]TagAssignment**](TagAssignment.md) | The tags assigned if any. | [optional] 
**IsSystemManaged** | Pointer to **bool** | True if the IP block is a &#x60;system managed&#x60; block. | [optional] 
**IsBringYourOwn** | Pointer to **bool** | True if the IP block is a &#x60;bring your own&#x60; block. | [optional] 
**CreatedOn** | Pointer to **time.Time** | Date and time when the IP block was created. | [optional] 

## Methods

### NewIpBlock

`func NewIpBlock() *IpBlock`

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

### HasId

`func (o *IpBlock) HasId() bool`

HasId returns a boolean if a field has been set.

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

### HasLocation

`func (o *IpBlock) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

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

### HasCidrBlockSize

`func (o *IpBlock) HasCidrBlockSize() bool`

HasCidrBlockSize returns a boolean if a field has been set.

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

### HasCidr

`func (o *IpBlock) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetIpVersion

`func (o *IpBlock) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *IpBlock) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *IpBlock) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *IpBlock) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

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

### HasStatus

`func (o *IpBlock) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParentIpBlockAllocationId

`func (o *IpBlock) GetParentIpBlockAllocationId() string`

GetParentIpBlockAllocationId returns the ParentIpBlockAllocationId field if non-nil, zero value otherwise.

### GetParentIpBlockAllocationIdOk

`func (o *IpBlock) GetParentIpBlockAllocationIdOk() (*string, bool)`

GetParentIpBlockAllocationIdOk returns a tuple with the ParentIpBlockAllocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIpBlockAllocationId

`func (o *IpBlock) SetParentIpBlockAllocationId(v string)`

SetParentIpBlockAllocationId sets ParentIpBlockAllocationId field to given value.

### HasParentIpBlockAllocationId

`func (o *IpBlock) HasParentIpBlockAllocationId() bool`

HasParentIpBlockAllocationId returns a boolean if a field has been set.

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

### GetTags

`func (o *IpBlock) GetTags() []TagAssignment`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpBlock) GetTagsOk() (*[]TagAssignment, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpBlock) SetTags(v []TagAssignment)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsSystemManaged

`func (o *IpBlock) GetIsSystemManaged() bool`

GetIsSystemManaged returns the IsSystemManaged field if non-nil, zero value otherwise.

### GetIsSystemManagedOk

`func (o *IpBlock) GetIsSystemManagedOk() (*bool, bool)`

GetIsSystemManagedOk returns a tuple with the IsSystemManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemManaged

`func (o *IpBlock) SetIsSystemManaged(v bool)`

SetIsSystemManaged sets IsSystemManaged field to given value.

### HasIsSystemManaged

`func (o *IpBlock) HasIsSystemManaged() bool`

HasIsSystemManaged returns a boolean if a field has been set.

### GetIsBringYourOwn

`func (o *IpBlock) GetIsBringYourOwn() bool`

GetIsBringYourOwn returns the IsBringYourOwn field if non-nil, zero value otherwise.

### GetIsBringYourOwnOk

`func (o *IpBlock) GetIsBringYourOwnOk() (*bool, bool)`

GetIsBringYourOwnOk returns a tuple with the IsBringYourOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBringYourOwn

`func (o *IpBlock) SetIsBringYourOwn(v bool)`

SetIsBringYourOwn sets IsBringYourOwn field to given value.

### HasIsBringYourOwn

`func (o *IpBlock) HasIsBringYourOwn() bool`

HasIsBringYourOwn returns a boolean if a field has been set.

### GetCreatedOn

`func (o *IpBlock) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *IpBlock) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *IpBlock) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *IpBlock) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


