# ProductCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductCategory** | [**ProductCategoryEnum**](ProductCategoryEnum.md) |  | 
**ProductCategoryDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewProductCategory

`func NewProductCategory(productCategory ProductCategoryEnum, ) *ProductCategory`

NewProductCategory instantiates a new ProductCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCategoryWithDefaults

`func NewProductCategoryWithDefaults() *ProductCategory`

NewProductCategoryWithDefaults instantiates a new ProductCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductCategory

`func (o *ProductCategory) GetProductCategory() ProductCategoryEnum`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *ProductCategory) GetProductCategoryOk() (*ProductCategoryEnum, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *ProductCategory) SetProductCategory(v ProductCategoryEnum)`

SetProductCategory sets ProductCategory field to given value.


### GetProductCategoryDescription

`func (o *ProductCategory) GetProductCategoryDescription() string`

GetProductCategoryDescription returns the ProductCategoryDescription field if non-nil, zero value otherwise.

### GetProductCategoryDescriptionOk

`func (o *ProductCategory) GetProductCategoryDescriptionOk() (*string, bool)`

GetProductCategoryDescriptionOk returns a tuple with the ProductCategoryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategoryDescription

`func (o *ProductCategory) SetProductCategoryDescription(v string)`

SetProductCategoryDescription sets ProductCategoryDescription field to given value.

### HasProductCategoryDescription

`func (o *ProductCategory) HasProductCategoryDescription() bool`

HasProductCategoryDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


