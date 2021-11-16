# Tag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the tag. | 
**Name** | **string** | The name of the tag. | 
**Values** | Pointer to **[]string** | The optional values of the tag. | [optional] 
**Description** | Pointer to **string** | The description of the tag. | [optional] 
**IsBillingTag** | **bool** | Whether or not to show the tag as part of billing and invoices. | 
**ResourceAssignments** | Pointer to [**[]ResourceAssignment**](ResourceAssignment.md) | The tag&#39;s assigned resources. | [optional] 

## Methods

### NewTag

`func NewTag(id string, name string, isBillingTag bool, ) *Tag`

NewTag instantiates a new Tag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithDefaults

`func NewTagWithDefaults() *Tag`

NewTagWithDefaults instantiates a new Tag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tag) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Tag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tag) SetName(v string)`

SetName sets Name field to given value.


### GetValues

`func (o *Tag) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Tag) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Tag) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *Tag) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetDescription

`func (o *Tag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsBillingTag

`func (o *Tag) GetIsBillingTag() bool`

GetIsBillingTag returns the IsBillingTag field if non-nil, zero value otherwise.

### GetIsBillingTagOk

`func (o *Tag) GetIsBillingTagOk() (*bool, bool)`

GetIsBillingTagOk returns a tuple with the IsBillingTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillingTag

`func (o *Tag) SetIsBillingTag(v bool)`

SetIsBillingTag sets IsBillingTag field to given value.


### GetResourceAssignments

`func (o *Tag) GetResourceAssignments() []ResourceAssignment`

GetResourceAssignments returns the ResourceAssignments field if non-nil, zero value otherwise.

### GetResourceAssignmentsOk

`func (o *Tag) GetResourceAssignmentsOk() (*[]ResourceAssignment, bool)`

GetResourceAssignmentsOk returns a tuple with the ResourceAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAssignments

`func (o *Tag) SetResourceAssignments(v []ResourceAssignment)`

SetResourceAssignments sets ResourceAssignments field to given value.

### HasResourceAssignments

`func (o *Tag) HasResourceAssignments() bool`

HasResourceAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


