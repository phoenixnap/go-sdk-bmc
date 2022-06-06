# ProductsGet200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductCode** | **string** | The code identifying the product. This code has significant across all locations. | 
**ProductCategory** | **string** | The product category. | 
**Plans** | Pointer to [**[]PricingPlan**](PricingPlan.md) | The pricing plans available for this product. | [optional] 
**Metadata** | [**ServerProductMetadata**](ServerProductMetadata.md) |  | 

## Methods

### NewProductsGet200ResponseInner

`func NewProductsGet200ResponseInner(productCode string, productCategory string, metadata ServerProductMetadata, ) *ProductsGet200ResponseInner`

NewProductsGet200ResponseInner instantiates a new ProductsGet200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsGet200ResponseInnerWithDefaults

`func NewProductsGet200ResponseInnerWithDefaults() *ProductsGet200ResponseInner`

NewProductsGet200ResponseInnerWithDefaults instantiates a new ProductsGet200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductCode

`func (o *ProductsGet200ResponseInner) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *ProductsGet200ResponseInner) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *ProductsGet200ResponseInner) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetProductCategory

`func (o *ProductsGet200ResponseInner) GetProductCategory() string`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *ProductsGet200ResponseInner) GetProductCategoryOk() (*string, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *ProductsGet200ResponseInner) SetProductCategory(v string)`

SetProductCategory sets ProductCategory field to given value.


### GetPlans

`func (o *ProductsGet200ResponseInner) GetPlans() []PricingPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ProductsGet200ResponseInner) GetPlansOk() (*[]PricingPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ProductsGet200ResponseInner) SetPlans(v []PricingPlan)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ProductsGet200ResponseInner) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetMetadata

`func (o *ProductsGet200ResponseInner) GetMetadata() ServerProductMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProductsGet200ResponseInner) GetMetadataOk() (*ServerProductMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProductsGet200ResponseInner) SetMetadata(v ServerProductMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


