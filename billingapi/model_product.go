/*
Billing API

Automate your infrastructure billing with the Bare Metal Cloud Billing API. Reserve your server instances to ensure guaranteed resource availability for 12, 24, and 36 months. Retrieve your server’s rated usage for a given period and enable or disable auto-renewals.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/phoenixnap-bare-metal-cloud-billing-models' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/billing/v1/)</b> 

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package billingapi

import (
	"encoding/json"
)

// Product Product details
type Product struct {
	// The code identifying the product. This code has significant across all locations.
	ProductCode string `json:"productCode"`
	// The product category.
	ProductCategory string `json:"productCategory"`
	// The pricing plans available for this product.
	Plans []PricingPlan `json:"plans,omitempty"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct(productCode string, productCategory string) *Product {
	this := Product{}
	this.ProductCode = productCode
	this.ProductCategory = productCategory
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetProductCode returns the ProductCode field value
func (o *Product) GetProductCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value
// and a boolean to check if the value has been set.
func (o *Product) GetProductCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductCode, true
}

// SetProductCode sets field value
func (o *Product) SetProductCode(v string) {
	o.ProductCode = v
}

// GetProductCategory returns the ProductCategory field value
func (o *Product) GetProductCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductCategory
}

// GetProductCategoryOk returns a tuple with the ProductCategory field value
// and a boolean to check if the value has been set.
func (o *Product) GetProductCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductCategory, true
}

// SetProductCategory sets field value
func (o *Product) SetProductCategory(v string) {
	o.ProductCategory = v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *Product) GetPlans() []PricingPlan {
	if o == nil || o.Plans == nil {
		var ret []PricingPlan
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetPlansOk() ([]PricingPlan, bool) {
	if o == nil || o.Plans == nil {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *Product) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []PricingPlan and assigns it to the Plans field.
func (o *Product) SetPlans(v []PricingPlan) {
	o.Plans = v
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["productCode"] = o.ProductCode
	}
	if true {
		toSerialize["productCategory"] = o.ProductCategory
	}
	if o.Plans != nil {
		toSerialize["plans"] = o.Plans
	}
	return json.Marshal(toSerialize)
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


