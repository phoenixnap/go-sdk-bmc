# SshKeyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | **bool** | Keys marked as default are always included on server creation and reset unless toggled off in creation/reset request. | 
**Name** | **string** | SSH key name that can represent the key as an alternative to its ID. | 

## Methods

### NewSshKeyUpdate

`func NewSshKeyUpdate(default_ bool, name string, ) *SshKeyUpdate`

NewSshKeyUpdate instantiates a new SshKeyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeyUpdateWithDefaults

`func NewSshKeyUpdateWithDefaults() *SshKeyUpdate`

NewSshKeyUpdateWithDefaults instantiates a new SshKeyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *SshKeyUpdate) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SshKeyUpdate) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SshKeyUpdate) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetName

`func (o *SshKeyUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshKeyUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshKeyUpdate) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


