# DeleteIpBlockResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | IP Block has been deleted. | [optional] 
**IpBlockId** | Pointer to **string** | The unique identifier of the IP Block. | [optional] 

## Methods

### NewDeleteIpBlockResult

`func NewDeleteIpBlockResult() *DeleteIpBlockResult`

NewDeleteIpBlockResult instantiates a new DeleteIpBlockResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteIpBlockResultWithDefaults

`func NewDeleteIpBlockResultWithDefaults() *DeleteIpBlockResult`

NewDeleteIpBlockResultWithDefaults instantiates a new DeleteIpBlockResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DeleteIpBlockResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeleteIpBlockResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeleteIpBlockResult) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DeleteIpBlockResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetIpBlockId

`func (o *DeleteIpBlockResult) GetIpBlockId() string`

GetIpBlockId returns the IpBlockId field if non-nil, zero value otherwise.

### GetIpBlockIdOk

`func (o *DeleteIpBlockResult) GetIpBlockIdOk() (*string, bool)`

GetIpBlockIdOk returns a tuple with the IpBlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlockId

`func (o *DeleteIpBlockResult) SetIpBlockId(v string)`

SetIpBlockId sets IpBlockId field to given value.

### HasIpBlockId

`func (o *DeleteIpBlockResult) HasIpBlockId() bool`

HasIpBlockId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


