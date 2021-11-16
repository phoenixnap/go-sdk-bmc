# ServerPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of server. | [optional] 
**Hostname** | Pointer to **string** | Hostname of server | [optional] 

## Methods

### NewServerPatch

`func NewServerPatch() *ServerPatch`

NewServerPatch instantiates a new ServerPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPatchWithDefaults

`func NewServerPatchWithDefaults() *ServerPatch`

NewServerPatchWithDefaults instantiates a new ServerPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServerPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostname

`func (o *ServerPatch) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerPatch) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerPatch) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ServerPatch) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


