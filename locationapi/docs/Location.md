# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**LocationEnum**](LocationEnum.md) |  | 
**LocationDescription** | Pointer to **string** |  | [optional] 
**ProductCategories** | Pointer to [**[]ProductCategory**](ProductCategory.md) |  | [optional] 

## Methods

### NewLocation

`func NewLocation(location LocationEnum, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *Location) GetLocation() LocationEnum`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Location) GetLocationOk() (*LocationEnum, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Location) SetLocation(v LocationEnum)`

SetLocation sets Location field to given value.


### GetLocationDescription

`func (o *Location) GetLocationDescription() string`

GetLocationDescription returns the LocationDescription field if non-nil, zero value otherwise.

### GetLocationDescriptionOk

`func (o *Location) GetLocationDescriptionOk() (*string, bool)`

GetLocationDescriptionOk returns a tuple with the LocationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationDescription

`func (o *Location) SetLocationDescription(v string)`

SetLocationDescription sets LocationDescription field to given value.

### HasLocationDescription

`func (o *Location) HasLocationDescription() bool`

HasLocationDescription returns a boolean if a field has been set.

### GetProductCategories

`func (o *Location) GetProductCategories() []ProductCategory`

GetProductCategories returns the ProductCategories field if non-nil, zero value otherwise.

### GetProductCategoriesOk

`func (o *Location) GetProductCategoriesOk() (*[]ProductCategory, bool)`

GetProductCategoriesOk returns a tuple with the ProductCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategories

`func (o *Location) SetProductCategories(v []ProductCategory)`

SetProductCategories sets ProductCategories field to given value.

### HasProductCategories

`func (o *Location) HasProductCategories() bool`

HasProductCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


