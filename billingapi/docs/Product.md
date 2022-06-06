# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductCode** | **string** | The code identifying the product. This code has significant across all locations. | 
**ProductCategory** | **string** | The product category. | 
**Plans** | Pointer to [**[]PricingPlan**](PricingPlan.md) | The pricing plans available for this product. | [optional] 

## Methods

### NewProduct

`func NewProduct(productCode string, productCategory string, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductCode

`func (o *Product) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *Product) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *Product) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetProductCategory

`func (o *Product) GetProductCategory() string`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *Product) GetProductCategoryOk() (*string, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *Product) SetProductCategory(v string)`

SetProductCategory sets ProductCategory field to given value.


### GetPlans

`func (o *Product) GetPlans() []PricingPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *Product) GetPlansOk() (*[]PricingPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *Product) SetPlans(v []PricingPlan)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *Product) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


