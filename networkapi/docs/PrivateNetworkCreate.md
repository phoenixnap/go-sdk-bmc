# PrivateNetworkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The friendly name of this private network. This name should be unique. | 
**Description** | Pointer to **string** | The description of this private network. | [optional] 
**Location** | **string** | The location of this private network. Supported values are &#x60;PHX&#x60;, &#x60;ASH&#x60;, &#x60;SGP&#x60;, &#x60;NLD&#x60;, &#x60;CHI&#x60; and &#x60;SEA&#x60;. | 
**LocationDefault** | Pointer to **bool** | Identifies network as the default private network for the specified location. | [optional] [default to false]
**Cidr** | **string** | IP range associated with this private network in CIDR notation. | 

## Methods

### NewPrivateNetworkCreate

`func NewPrivateNetworkCreate(name string, location string, cidr string, ) *PrivateNetworkCreate`

NewPrivateNetworkCreate instantiates a new PrivateNetworkCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkCreateWithDefaults

`func NewPrivateNetworkCreateWithDefaults() *PrivateNetworkCreate`

NewPrivateNetworkCreateWithDefaults instantiates a new PrivateNetworkCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateNetworkCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateNetworkCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateNetworkCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PrivateNetworkCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateNetworkCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateNetworkCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrivateNetworkCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *PrivateNetworkCreate) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PrivateNetworkCreate) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PrivateNetworkCreate) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetLocationDefault

`func (o *PrivateNetworkCreate) GetLocationDefault() bool`

GetLocationDefault returns the LocationDefault field if non-nil, zero value otherwise.

### GetLocationDefaultOk

`func (o *PrivateNetworkCreate) GetLocationDefaultOk() (*bool, bool)`

GetLocationDefaultOk returns a tuple with the LocationDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationDefault

`func (o *PrivateNetworkCreate) SetLocationDefault(v bool)`

SetLocationDefault sets LocationDefault field to given value.

### HasLocationDefault

`func (o *PrivateNetworkCreate) HasLocationDefault() bool`

HasLocationDefault returns a boolean if a field has been set.

### GetCidr

`func (o *PrivateNetworkCreate) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *PrivateNetworkCreate) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *PrivateNetworkCreate) SetCidr(v string)`

SetCidr sets Cidr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


