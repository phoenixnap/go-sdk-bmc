# OsConfigurationIPXE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL of the iPXE boot script used to start the server. | 
**NativeVlanConfiguration** | Pointer to [**OsConfigurationIPXENativeVlanConfiguration**](OsConfigurationIPXENativeVlanConfiguration.md) |  | [optional] 

## Methods

### NewOsConfigurationIPXE

`func NewOsConfigurationIPXE(url string, ) *OsConfigurationIPXE`

NewOsConfigurationIPXE instantiates a new OsConfigurationIPXE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsConfigurationIPXEWithDefaults

`func NewOsConfigurationIPXEWithDefaults() *OsConfigurationIPXE`

NewOsConfigurationIPXEWithDefaults instantiates a new OsConfigurationIPXE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OsConfigurationIPXE) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OsConfigurationIPXE) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OsConfigurationIPXE) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNativeVlanConfiguration

`func (o *OsConfigurationIPXE) GetNativeVlanConfiguration() OsConfigurationIPXENativeVlanConfiguration`

GetNativeVlanConfiguration returns the NativeVlanConfiguration field if non-nil, zero value otherwise.

### GetNativeVlanConfigurationOk

`func (o *OsConfigurationIPXE) GetNativeVlanConfigurationOk() (*OsConfigurationIPXENativeVlanConfiguration, bool)`

GetNativeVlanConfigurationOk returns a tuple with the NativeVlanConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVlanConfiguration

`func (o *OsConfigurationIPXE) SetNativeVlanConfiguration(v OsConfigurationIPXENativeVlanConfiguration)`

SetNativeVlanConfiguration sets NativeVlanConfiguration field to given value.

### HasNativeVlanConfiguration

`func (o *OsConfigurationIPXE) HasNativeVlanConfiguration() bool`

HasNativeVlanConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


