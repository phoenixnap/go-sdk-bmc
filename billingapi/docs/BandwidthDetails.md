# BandwidthDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngressGb** | **float32** | The amount of GB consumed in ingress (IN). | 
**EgressGb** | **float32** | The amount of GB consumed in egress (OUT). | 
**PackageQuantity** | Pointer to **float32** | Package size per month. | [optional] 
**PackageUnit** | Pointer to **string** | Package size unit. | [optional] 

## Methods

### NewBandwidthDetails

`func NewBandwidthDetails(ingressGb float32, egressGb float32, ) *BandwidthDetails`

NewBandwidthDetails instantiates a new BandwidthDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthDetailsWithDefaults

`func NewBandwidthDetailsWithDefaults() *BandwidthDetails`

NewBandwidthDetailsWithDefaults instantiates a new BandwidthDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngressGb

`func (o *BandwidthDetails) GetIngressGb() float32`

GetIngressGb returns the IngressGb field if non-nil, zero value otherwise.

### GetIngressGbOk

`func (o *BandwidthDetails) GetIngressGbOk() (*float32, bool)`

GetIngressGbOk returns a tuple with the IngressGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressGb

`func (o *BandwidthDetails) SetIngressGb(v float32)`

SetIngressGb sets IngressGb field to given value.


### GetEgressGb

`func (o *BandwidthDetails) GetEgressGb() float32`

GetEgressGb returns the EgressGb field if non-nil, zero value otherwise.

### GetEgressGbOk

`func (o *BandwidthDetails) GetEgressGbOk() (*float32, bool)`

GetEgressGbOk returns a tuple with the EgressGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressGb

`func (o *BandwidthDetails) SetEgressGb(v float32)`

SetEgressGb sets EgressGb field to given value.


### GetPackageQuantity

`func (o *BandwidthDetails) GetPackageQuantity() float32`

GetPackageQuantity returns the PackageQuantity field if non-nil, zero value otherwise.

### GetPackageQuantityOk

`func (o *BandwidthDetails) GetPackageQuantityOk() (*float32, bool)`

GetPackageQuantityOk returns a tuple with the PackageQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageQuantity

`func (o *BandwidthDetails) SetPackageQuantity(v float32)`

SetPackageQuantity sets PackageQuantity field to given value.

### HasPackageQuantity

`func (o *BandwidthDetails) HasPackageQuantity() bool`

HasPackageQuantity returns a boolean if a field has been set.

### GetPackageUnit

`func (o *BandwidthDetails) GetPackageUnit() string`

GetPackageUnit returns the PackageUnit field if non-nil, zero value otherwise.

### GetPackageUnitOk

`func (o *BandwidthDetails) GetPackageUnitOk() (*string, bool)`

GetPackageUnitOk returns a tuple with the PackageUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUnit

`func (o *BandwidthDetails) SetPackageUnit(v string)`

SetPackageUnit sets PackageUnit field to given value.

### HasPackageUnit

`func (o *BandwidthDetails) HasPackageUnit() bool`

HasPackageUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


