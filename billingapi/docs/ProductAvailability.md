# ProductAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductCode** | **string** | Product code. | 
**ProductCategory** | **string** | Product category. | 
**LocationAvailabilityDetails** | [**[]LocationAvailabilityDetail**](LocationAvailabilityDetail.md) |  | 

## Methods

### NewProductAvailability

`func NewProductAvailability(productCode string, productCategory string, locationAvailabilityDetails []LocationAvailabilityDetail, ) *ProductAvailability`

NewProductAvailability instantiates a new ProductAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAvailabilityWithDefaults

`func NewProductAvailabilityWithDefaults() *ProductAvailability`

NewProductAvailabilityWithDefaults instantiates a new ProductAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductCode

`func (o *ProductAvailability) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *ProductAvailability) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *ProductAvailability) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetProductCategory

`func (o *ProductAvailability) GetProductCategory() string`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *ProductAvailability) GetProductCategoryOk() (*string, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *ProductAvailability) SetProductCategory(v string)`

SetProductCategory sets ProductCategory field to given value.


### GetLocationAvailabilityDetails

`func (o *ProductAvailability) GetLocationAvailabilityDetails() []LocationAvailabilityDetail`

GetLocationAvailabilityDetails returns the LocationAvailabilityDetails field if non-nil, zero value otherwise.

### GetLocationAvailabilityDetailsOk

`func (o *ProductAvailability) GetLocationAvailabilityDetailsOk() (*[]LocationAvailabilityDetail, bool)`

GetLocationAvailabilityDetailsOk returns a tuple with the LocationAvailabilityDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationAvailabilityDetails

`func (o *ProductAvailability) SetLocationAvailabilityDetails(v []LocationAvailabilityDetail)`

SetLocationAvailabilityDetails sets LocationAvailabilityDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


