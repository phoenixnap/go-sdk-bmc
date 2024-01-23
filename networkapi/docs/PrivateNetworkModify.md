# PrivateNetworkModify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A friendly name given to the private network. This name should be unique. | 
**Description** | Pointer to **string** | A description of this private network | [optional] 
**LocationDefault** | **bool** | Identifies network as the default private network for the specified location. | 

## Methods

### NewPrivateNetworkModify

`func NewPrivateNetworkModify(name string, locationDefault bool, ) *PrivateNetworkModify`

NewPrivateNetworkModify instantiates a new PrivateNetworkModify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkModifyWithDefaults

`func NewPrivateNetworkModifyWithDefaults() *PrivateNetworkModify`

NewPrivateNetworkModifyWithDefaults instantiates a new PrivateNetworkModify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateNetworkModify) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateNetworkModify) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateNetworkModify) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PrivateNetworkModify) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateNetworkModify) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateNetworkModify) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrivateNetworkModify) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocationDefault

`func (o *PrivateNetworkModify) GetLocationDefault() bool`

GetLocationDefault returns the LocationDefault field if non-nil, zero value otherwise.

### GetLocationDefaultOk

`func (o *PrivateNetworkModify) GetLocationDefaultOk() (*bool, bool)`

GetLocationDefaultOk returns a tuple with the LocationDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationDefault

`func (o *PrivateNetworkModify) SetLocationDefault(v bool)`

SetLocationDefault sets LocationDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


