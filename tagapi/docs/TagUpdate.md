# TagUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of the tag. Tag names are case-sensitive, and should be composed of a maximum of 100 characters including UTF-8 Unicode letters, numbers, and the following symbols: &#39;-&#39;, &#39;_&#39;. Regex: [A-zÀ-ú0-9_-]{1,100} | 
**Description** | Pointer to **string** | The description of the tag. | [optional] 
**IsBillingTag** | **bool** | Whether or not to show the tag as part of billing and invoices. | 

## Methods

### NewTagUpdate

`func NewTagUpdate(name string, isBillingTag bool, ) *TagUpdate`

NewTagUpdate instantiates a new TagUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagUpdateWithDefaults

`func NewTagUpdateWithDefaults() *TagUpdate`

NewTagUpdateWithDefaults instantiates a new TagUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TagUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsBillingTag

`func (o *TagUpdate) GetIsBillingTag() bool`

GetIsBillingTag returns the IsBillingTag field if non-nil, zero value otherwise.

### GetIsBillingTagOk

`func (o *TagUpdate) GetIsBillingTagOk() (*bool, bool)`

GetIsBillingTagOk returns a tuple with the IsBillingTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillingTag

`func (o *TagUpdate) SetIsBillingTag(v bool)`

SetIsBillingTag sets IsBillingTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


