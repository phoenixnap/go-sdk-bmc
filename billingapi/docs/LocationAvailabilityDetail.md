# LocationAvailabilityDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**LocationEnum**](LocationEnum.md) |  | 
**MinQuantityRequested** | **float32** | Requested quantity. | 
**MinQuantityAvailable** | **bool** | Is product available in specific location for requested quantity | 
**AvailableQuantity** | **float32** | Total available quantity of product in specific location. Max value is 10. | 
**Solutions** | **[]string** | Solutions supported in specific location for a product. | 

## Methods

### NewLocationAvailabilityDetail

`func NewLocationAvailabilityDetail(location LocationEnum, minQuantityRequested float32, minQuantityAvailable bool, availableQuantity float32, solutions []string, ) *LocationAvailabilityDetail`

NewLocationAvailabilityDetail instantiates a new LocationAvailabilityDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationAvailabilityDetailWithDefaults

`func NewLocationAvailabilityDetailWithDefaults() *LocationAvailabilityDetail`

NewLocationAvailabilityDetailWithDefaults instantiates a new LocationAvailabilityDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *LocationAvailabilityDetail) GetLocation() LocationEnum`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LocationAvailabilityDetail) GetLocationOk() (*LocationEnum, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LocationAvailabilityDetail) SetLocation(v LocationEnum)`

SetLocation sets Location field to given value.


### GetMinQuantityRequested

`func (o *LocationAvailabilityDetail) GetMinQuantityRequested() float32`

GetMinQuantityRequested returns the MinQuantityRequested field if non-nil, zero value otherwise.

### GetMinQuantityRequestedOk

`func (o *LocationAvailabilityDetail) GetMinQuantityRequestedOk() (*float32, bool)`

GetMinQuantityRequestedOk returns a tuple with the MinQuantityRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantityRequested

`func (o *LocationAvailabilityDetail) SetMinQuantityRequested(v float32)`

SetMinQuantityRequested sets MinQuantityRequested field to given value.


### GetMinQuantityAvailable

`func (o *LocationAvailabilityDetail) GetMinQuantityAvailable() bool`

GetMinQuantityAvailable returns the MinQuantityAvailable field if non-nil, zero value otherwise.

### GetMinQuantityAvailableOk

`func (o *LocationAvailabilityDetail) GetMinQuantityAvailableOk() (*bool, bool)`

GetMinQuantityAvailableOk returns a tuple with the MinQuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantityAvailable

`func (o *LocationAvailabilityDetail) SetMinQuantityAvailable(v bool)`

SetMinQuantityAvailable sets MinQuantityAvailable field to given value.


### GetAvailableQuantity

`func (o *LocationAvailabilityDetail) GetAvailableQuantity() float32`

GetAvailableQuantity returns the AvailableQuantity field if non-nil, zero value otherwise.

### GetAvailableQuantityOk

`func (o *LocationAvailabilityDetail) GetAvailableQuantityOk() (*float32, bool)`

GetAvailableQuantityOk returns a tuple with the AvailableQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableQuantity

`func (o *LocationAvailabilityDetail) SetAvailableQuantity(v float32)`

SetAvailableQuantity sets AvailableQuantity field to given value.


### GetSolutions

`func (o *LocationAvailabilityDetail) GetSolutions() []string`

GetSolutions returns the Solutions field if non-nil, zero value otherwise.

### GetSolutionsOk

`func (o *LocationAvailabilityDetail) GetSolutionsOk() (*[]string, bool)`

GetSolutionsOk returns a tuple with the Solutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutions

`func (o *LocationAvailabilityDetail) SetSolutions(v []string)`

SetSolutions sets Solutions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


