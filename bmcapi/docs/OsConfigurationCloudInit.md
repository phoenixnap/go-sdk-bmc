# OsConfigurationCloudInit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserData** | Pointer to **string** | User data for cloud init configuration in base64 encoding. NoCloud format is supported. | [optional] 

## Methods

### NewOsConfigurationCloudInit

`func NewOsConfigurationCloudInit() *OsConfigurationCloudInit`

NewOsConfigurationCloudInit instantiates a new OsConfigurationCloudInit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsConfigurationCloudInitWithDefaults

`func NewOsConfigurationCloudInitWithDefaults() *OsConfigurationCloudInit`

NewOsConfigurationCloudInitWithDefaults instantiates a new OsConfigurationCloudInit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserData

`func (o *OsConfigurationCloudInit) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *OsConfigurationCloudInit) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *OsConfigurationCloudInit) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *OsConfigurationCloudInit) HasUserData() bool`

HasUserData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


