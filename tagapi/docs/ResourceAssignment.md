# ResourceAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | **string** | The resource name which is automatically generated by the tags-api. It is a unique resource identifier made up of the API name (e.g. bmc, ips), the resource type and the resource ID. This is not to be confused with custom names that are defined for particular resources, such as the server name or network name. | 
**Value** | Pointer to **string** | The value of the tag assigned to the resource. | [optional] 

## Methods

### NewResourceAssignment

`func NewResourceAssignment(resourceName string, ) *ResourceAssignment`

NewResourceAssignment instantiates a new ResourceAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAssignmentWithDefaults

`func NewResourceAssignmentWithDefaults() *ResourceAssignment`

NewResourceAssignmentWithDefaults instantiates a new ResourceAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceName

`func (o *ResourceAssignment) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceAssignment) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceAssignment) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetValue

`func (o *ResourceAssignment) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResourceAssignment) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResourceAssignment) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResourceAssignment) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


