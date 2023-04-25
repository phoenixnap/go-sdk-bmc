# OsConfigurationNetrisSoftgate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostOs** | Pointer to **string** | (Read-only) Host OS on which the Netris Softgate is installed. | [optional] [readonly] 
**ControllerAddress** | Pointer to **string** | (Write-only) IP address or hostname through which to reach the Netris Controller. | [optional] 
**ControllerVersion** | Pointer to **string** | (Write-only) The version of the Netris Controller to connect to. | [optional] 
**ControllerAuthKey** | Pointer to **string** | (Write-only) The authentication key of the Netris Controller to connect to. Required for the softgate agent to be able to interact with the Netris Controller. | [optional] 

## Methods

### NewOsConfigurationNetrisSoftgate

`func NewOsConfigurationNetrisSoftgate() *OsConfigurationNetrisSoftgate`

NewOsConfigurationNetrisSoftgate instantiates a new OsConfigurationNetrisSoftgate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsConfigurationNetrisSoftgateWithDefaults

`func NewOsConfigurationNetrisSoftgateWithDefaults() *OsConfigurationNetrisSoftgate`

NewOsConfigurationNetrisSoftgateWithDefaults instantiates a new OsConfigurationNetrisSoftgate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostOs

`func (o *OsConfigurationNetrisSoftgate) GetHostOs() string`

GetHostOs returns the HostOs field if non-nil, zero value otherwise.

### GetHostOsOk

`func (o *OsConfigurationNetrisSoftgate) GetHostOsOk() (*string, bool)`

GetHostOsOk returns a tuple with the HostOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOs

`func (o *OsConfigurationNetrisSoftgate) SetHostOs(v string)`

SetHostOs sets HostOs field to given value.

### HasHostOs

`func (o *OsConfigurationNetrisSoftgate) HasHostOs() bool`

HasHostOs returns a boolean if a field has been set.

### GetControllerAddress

`func (o *OsConfigurationNetrisSoftgate) GetControllerAddress() string`

GetControllerAddress returns the ControllerAddress field if non-nil, zero value otherwise.

### GetControllerAddressOk

`func (o *OsConfigurationNetrisSoftgate) GetControllerAddressOk() (*string, bool)`

GetControllerAddressOk returns a tuple with the ControllerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerAddress

`func (o *OsConfigurationNetrisSoftgate) SetControllerAddress(v string)`

SetControllerAddress sets ControllerAddress field to given value.

### HasControllerAddress

`func (o *OsConfigurationNetrisSoftgate) HasControllerAddress() bool`

HasControllerAddress returns a boolean if a field has been set.

### GetControllerVersion

`func (o *OsConfigurationNetrisSoftgate) GetControllerVersion() string`

GetControllerVersion returns the ControllerVersion field if non-nil, zero value otherwise.

### GetControllerVersionOk

`func (o *OsConfigurationNetrisSoftgate) GetControllerVersionOk() (*string, bool)`

GetControllerVersionOk returns a tuple with the ControllerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVersion

`func (o *OsConfigurationNetrisSoftgate) SetControllerVersion(v string)`

SetControllerVersion sets ControllerVersion field to given value.

### HasControllerVersion

`func (o *OsConfigurationNetrisSoftgate) HasControllerVersion() bool`

HasControllerVersion returns a boolean if a field has been set.

### GetControllerAuthKey

`func (o *OsConfigurationNetrisSoftgate) GetControllerAuthKey() string`

GetControllerAuthKey returns the ControllerAuthKey field if non-nil, zero value otherwise.

### GetControllerAuthKeyOk

`func (o *OsConfigurationNetrisSoftgate) GetControllerAuthKeyOk() (*string, bool)`

GetControllerAuthKeyOk returns a tuple with the ControllerAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerAuthKey

`func (o *OsConfigurationNetrisSoftgate) SetControllerAuthKey(v string)`

SetControllerAuthKey sets ControllerAuthKey field to given value.

### HasControllerAuthKey

`func (o *OsConfigurationNetrisSoftgate) HasControllerAuthKey() bool`

HasControllerAuthKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


