# ResetResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Message describing the reset result. | 
**Password** | Pointer to **string** | Password set for user Admin on Windows server or user root on ESXi server. | [optional] 
**OsConfiguration** | Pointer to [**OsConfigurationMap**](OsConfigurationMap.md) |  | [optional] 

## Methods

### NewResetResult

`func NewResetResult(result string, ) *ResetResult`

NewResetResult instantiates a new ResetResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetResultWithDefaults

`func NewResetResultWithDefaults() *ResetResult`

NewResetResultWithDefaults instantiates a new ResetResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ResetResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ResetResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ResetResult) SetResult(v string)`

SetResult sets Result field to given value.


### GetPassword

`func (o *ResetResult) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ResetResult) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ResetResult) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ResetResult) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetOsConfiguration

`func (o *ResetResult) GetOsConfiguration() OsConfigurationMap`

GetOsConfiguration returns the OsConfiguration field if non-nil, zero value otherwise.

### GetOsConfigurationOk

`func (o *ResetResult) GetOsConfigurationOk() (*OsConfigurationMap, bool)`

GetOsConfigurationOk returns a tuple with the OsConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsConfiguration

`func (o *ResetResult) SetOsConfiguration(v OsConfigurationMap)`

SetOsConfiguration sets OsConfiguration field to given value.

### HasOsConfiguration

`func (o *ResetResult) HasOsConfiguration() bool`

HasOsConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


