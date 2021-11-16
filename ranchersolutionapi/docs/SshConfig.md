# SshConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallDefaultKeys** | Pointer to **bool** | Define whether public keys marked as default should be installed on this node. These are public keys that were already recorded on this system. Use &lt;a href&#x3D;\&quot;https://developers.phoenixnap.com/docs/bmc/1/routes/ssh-keys/get\&quot; target&#x3D;\&quot;_blank\&quot;&gt;GET /ssh-keys&lt;/a&gt; to retrieve a list of possible values. | [optional] [default to true]
**Keys** | Pointer to **[]string** | List of public SSH keys. | [optional] 
**KeyIds** | Pointer to **[]string** | List of public SSH key identifiers. These are public keys that were already recorded on this system. Use &lt;a href&#x3D;\&quot;https://developers.phoenixnap.com/docs/bmc/1/routes/ssh-keys/get\&quot; target&#x3D;\&quot;_blank\&quot;&gt;GET /ssh-keys&lt;/a&gt; to retrieve a list of possible values. | [optional] 

## Methods

### NewSshConfig

`func NewSshConfig() *SshConfig`

NewSshConfig instantiates a new SshConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshConfigWithDefaults

`func NewSshConfigWithDefaults() *SshConfig`

NewSshConfigWithDefaults instantiates a new SshConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallDefaultKeys

`func (o *SshConfig) GetInstallDefaultKeys() bool`

GetInstallDefaultKeys returns the InstallDefaultKeys field if non-nil, zero value otherwise.

### GetInstallDefaultKeysOk

`func (o *SshConfig) GetInstallDefaultKeysOk() (*bool, bool)`

GetInstallDefaultKeysOk returns a tuple with the InstallDefaultKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDefaultKeys

`func (o *SshConfig) SetInstallDefaultKeys(v bool)`

SetInstallDefaultKeys sets InstallDefaultKeys field to given value.

### HasInstallDefaultKeys

`func (o *SshConfig) HasInstallDefaultKeys() bool`

HasInstallDefaultKeys returns a boolean if a field has been set.

### GetKeys

`func (o *SshConfig) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *SshConfig) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *SshConfig) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *SshConfig) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetKeyIds

`func (o *SshConfig) GetKeyIds() []string`

GetKeyIds returns the KeyIds field if non-nil, zero value otherwise.

### GetKeyIdsOk

`func (o *SshConfig) GetKeyIdsOk() (*[]string, bool)`

GetKeyIdsOk returns a tuple with the KeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyIds

`func (o *SshConfig) SetKeyIds(v []string)`

SetKeyIds sets KeyIds field to given value.

### HasKeyIds

`func (o *SshConfig) HasKeyIds() bool`

HasKeyIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


