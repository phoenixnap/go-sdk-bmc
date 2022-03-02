# TagAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the tag. | 
**Name** | **string** | The name of the tag. | 
**Value** | Pointer to **string** | The value of the tag assigned to the resource. | [optional] 
**IsBillingTag** | **bool** | Whether or not to show the tag as part of billing and invoices | 
**CreatedBy** | Pointer to **string** | Who the tag was created by. | [optional] 

## Methods

### NewTagAssignment

`func NewTagAssignment(id string, name string, isBillingTag bool, ) *TagAssignment`

NewTagAssignment instantiates a new TagAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagAssignmentWithDefaults

`func NewTagAssignmentWithDefaults() *TagAssignment`

NewTagAssignmentWithDefaults instantiates a new TagAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagAssignment) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TagAssignment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagAssignment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagAssignment) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *TagAssignment) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TagAssignment) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TagAssignment) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TagAssignment) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetIsBillingTag

`func (o *TagAssignment) GetIsBillingTag() bool`

GetIsBillingTag returns the IsBillingTag field if non-nil, zero value otherwise.

### GetIsBillingTagOk

`func (o *TagAssignment) GetIsBillingTagOk() (*bool, bool)`

GetIsBillingTagOk returns a tuple with the IsBillingTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillingTag

`func (o *TagAssignment) SetIsBillingTag(v bool)`

SetIsBillingTag sets IsBillingTag field to given value.


### GetCreatedBy

`func (o *TagAssignment) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TagAssignment) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TagAssignment) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TagAssignment) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


