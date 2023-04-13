# OsConfigurationNetrisController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostOs** | Pointer to **string** | (Read-only) Host OS on which the Netris Controller is installed. | [optional] [readonly] 
**NetrisWebConsoleUrl** | Pointer to **string** | (Read-only) The URL for the Netris Controller web console. Will only be returned in response to provisioning a server. | [optional] [readonly] 
**NetrisUserPassword** | Pointer to **string** | (Read-only) Auto-generated password set for user &#39;netris&#39; in the web console.&lt;br&gt;  The password is not stored and therefore will only be returned in response to provisioning a server. Copy and save it for future reference. | [optional] [readonly] 

## Methods

### NewOsConfigurationNetrisController

`func NewOsConfigurationNetrisController() *OsConfigurationNetrisController`

NewOsConfigurationNetrisController instantiates a new OsConfigurationNetrisController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsConfigurationNetrisControllerWithDefaults

`func NewOsConfigurationNetrisControllerWithDefaults() *OsConfigurationNetrisController`

NewOsConfigurationNetrisControllerWithDefaults instantiates a new OsConfigurationNetrisController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostOs

`func (o *OsConfigurationNetrisController) GetHostOs() string`

GetHostOs returns the HostOs field if non-nil, zero value otherwise.

### GetHostOsOk

`func (o *OsConfigurationNetrisController) GetHostOsOk() (*string, bool)`

GetHostOsOk returns a tuple with the HostOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOs

`func (o *OsConfigurationNetrisController) SetHostOs(v string)`

SetHostOs sets HostOs field to given value.

### HasHostOs

`func (o *OsConfigurationNetrisController) HasHostOs() bool`

HasHostOs returns a boolean if a field has been set.

### GetNetrisWebConsoleUrl

`func (o *OsConfigurationNetrisController) GetNetrisWebConsoleUrl() string`

GetNetrisWebConsoleUrl returns the NetrisWebConsoleUrl field if non-nil, zero value otherwise.

### GetNetrisWebConsoleUrlOk

`func (o *OsConfigurationNetrisController) GetNetrisWebConsoleUrlOk() (*string, bool)`

GetNetrisWebConsoleUrlOk returns a tuple with the NetrisWebConsoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetrisWebConsoleUrl

`func (o *OsConfigurationNetrisController) SetNetrisWebConsoleUrl(v string)`

SetNetrisWebConsoleUrl sets NetrisWebConsoleUrl field to given value.

### HasNetrisWebConsoleUrl

`func (o *OsConfigurationNetrisController) HasNetrisWebConsoleUrl() bool`

HasNetrisWebConsoleUrl returns a boolean if a field has been set.

### GetNetrisUserPassword

`func (o *OsConfigurationNetrisController) GetNetrisUserPassword() string`

GetNetrisUserPassword returns the NetrisUserPassword field if non-nil, zero value otherwise.

### GetNetrisUserPasswordOk

`func (o *OsConfigurationNetrisController) GetNetrisUserPasswordOk() (*string, bool)`

GetNetrisUserPasswordOk returns a tuple with the NetrisUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetrisUserPassword

`func (o *OsConfigurationNetrisController) SetNetrisUserPassword(v string)`

SetNetrisUserPassword sets NetrisUserPassword field to given value.

### HasNetrisUserPassword

`func (o *OsConfigurationNetrisController) HasNetrisUserPassword() bool`

HasNetrisUserPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


