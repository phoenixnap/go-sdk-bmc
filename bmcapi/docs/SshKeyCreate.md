# SshKeyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | **bool** | Keys marked as default are always included on server creation and reset unless toggled off in creation/reset request. | 
**Name** | **string** | Friendly SSH key name to represent an SSH key. | 
**Key** | **string** | SSH key actual key value. | 

## Methods

### NewSshKeyCreate

`func NewSshKeyCreate(default_ bool, name string, key string, ) *SshKeyCreate`

NewSshKeyCreate instantiates a new SshKeyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeyCreateWithDefaults

`func NewSshKeyCreateWithDefaults() *SshKeyCreate`

NewSshKeyCreateWithDefaults instantiates a new SshKeyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *SshKeyCreate) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SshKeyCreate) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SshKeyCreate) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetName

`func (o *SshKeyCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshKeyCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshKeyCreate) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *SshKeyCreate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SshKeyCreate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SshKeyCreate) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


