# TagCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of the tag. Tag names are case-sensitive, and should be composed of a maximum of 100 characters including UTF-8 Unicode letters, numbers, and the following symbols: &#39;-&#39;, &#39;_&#39;. Regex: [A-zÀ-ú0-9_-]{1,100} | 
**Description** | Pointer to **string** | The description of the tag. | [optional] 
**IsBillingTag** | **bool** | Whether or not to show the tag as part of billing and invoices. | 

## Methods

### NewTagCreate

`func NewTagCreate(name string, isBillingTag bool, ) *TagCreate`

NewTagCreate instantiates a new TagCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagCreateWithDefaults

`func NewTagCreateWithDefaults() *TagCreate`

NewTagCreateWithDefaults instantiates a new TagCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TagCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsBillingTag

`func (o *TagCreate) GetIsBillingTag() bool`

GetIsBillingTag returns the IsBillingTag field if non-nil, zero value otherwise.

### GetIsBillingTagOk

`func (o *TagCreate) GetIsBillingTagOk() (*bool, bool)`

GetIsBillingTagOk returns a tuple with the IsBillingTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillingTag

`func (o *TagCreate) SetIsBillingTag(v bool)`

SetIsBillingTag sets IsBillingTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


