# DeleteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Server has been deleted. | 
**ServerId** | **string** | The unique identifier of the server. | 

## Methods

### NewDeleteResult

`func NewDeleteResult(result string, serverId string, ) *DeleteResult`

NewDeleteResult instantiates a new DeleteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteResultWithDefaults

`func NewDeleteResultWithDefaults() *DeleteResult`

NewDeleteResultWithDefaults instantiates a new DeleteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DeleteResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeleteResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeleteResult) SetResult(v string)`

SetResult sets Result field to given value.


### GetServerId

`func (o *DeleteResult) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DeleteResult) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DeleteResult) SetServerId(v string)`

SetServerId sets ServerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


