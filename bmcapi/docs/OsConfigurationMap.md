# OsConfigurationMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Windows** | Pointer to [**OsConfigurationWindows**](OsConfigurationWindows.md) |  | [optional] 
**Esxi** | Pointer to [**OsConfigurationMapEsxi**](OsConfigurationMapEsxi.md) |  | [optional] 
**Proxmox** | Pointer to [**OsConfigurationMapProxmox**](OsConfigurationMapProxmox.md) |  | [optional] 

## Methods

### NewOsConfigurationMap

`func NewOsConfigurationMap() *OsConfigurationMap`

NewOsConfigurationMap instantiates a new OsConfigurationMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsConfigurationMapWithDefaults

`func NewOsConfigurationMapWithDefaults() *OsConfigurationMap`

NewOsConfigurationMapWithDefaults instantiates a new OsConfigurationMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindows

`func (o *OsConfigurationMap) GetWindows() OsConfigurationWindows`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *OsConfigurationMap) GetWindowsOk() (*OsConfigurationWindows, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *OsConfigurationMap) SetWindows(v OsConfigurationWindows)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *OsConfigurationMap) HasWindows() bool`

HasWindows returns a boolean if a field has been set.

### GetEsxi

`func (o *OsConfigurationMap) GetEsxi() OsConfigurationMapEsxi`

GetEsxi returns the Esxi field if non-nil, zero value otherwise.

### GetEsxiOk

`func (o *OsConfigurationMap) GetEsxiOk() (*OsConfigurationMapEsxi, bool)`

GetEsxiOk returns a tuple with the Esxi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsxi

`func (o *OsConfigurationMap) SetEsxi(v OsConfigurationMapEsxi)`

SetEsxi sets Esxi field to given value.

### HasEsxi

`func (o *OsConfigurationMap) HasEsxi() bool`

HasEsxi returns a boolean if a field has been set.

### GetProxmox

`func (o *OsConfigurationMap) GetProxmox() OsConfigurationMapProxmox`

GetProxmox returns the Proxmox field if non-nil, zero value otherwise.

### GetProxmoxOk

`func (o *OsConfigurationMap) GetProxmoxOk() (*OsConfigurationMapProxmox, bool)`

GetProxmoxOk returns a tuple with the Proxmox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxmox

`func (o *OsConfigurationMap) SetProxmox(v OsConfigurationMapProxmox)`

SetProxmox sets Proxmox field to given value.

### HasProxmox

`func (o *OsConfigurationMap) HasProxmox() bool`

HasProxmox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


