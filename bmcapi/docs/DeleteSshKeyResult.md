# DeleteSshKeyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Resource has been deleted. | 
**SshKeyId** | **string** | The unique identifier of the deleted resource. | 

## Methods

### NewDeleteSshKeyResult

`func NewDeleteSshKeyResult(result string, sshKeyId string, ) *DeleteSshKeyResult`

NewDeleteSshKeyResult instantiates a new DeleteSshKeyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSshKeyResultWithDefaults

`func NewDeleteSshKeyResultWithDefaults() *DeleteSshKeyResult`

NewDeleteSshKeyResultWithDefaults instantiates a new DeleteSshKeyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DeleteSshKeyResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeleteSshKeyResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeleteSshKeyResult) SetResult(v string)`

SetResult sets Result field to given value.


### GetSshKeyId

`func (o *DeleteSshKeyResult) GetSshKeyId() string`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *DeleteSshKeyResult) GetSshKeyIdOk() (*string, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *DeleteSshKeyResult) SetSshKeyId(v string)`

SetSshKeyId sets SshKeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


