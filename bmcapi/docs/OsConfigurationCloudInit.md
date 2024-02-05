# OsConfigurationCloudInit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserData** | Pointer to **string** | (Write-only) User data for the &lt;a href&#x3D;&#39;https://cloudinit.readthedocs.io/en/latest/&#39; target&#x3D;&#39;_blank&#39;&gt;cloud-init&lt;/a&gt; configuration in base64 encoding. NoCloud format is supported. Follow the &lt;a href&#x3D;&#39;https://phoenixnap.com/kb/bmc-cloud-init&#39; target&#x3D;&#39;_blank&#39;&gt;instructions&lt;/a&gt; on how to provision a server using cloud-init. Only ubuntu/bionic, ubuntu/focal and ubuntu/jammy are supported. User data will not be stored and cannot be retrieved once you deploy the server. Copy and save it for future reference. | [optional] 

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


