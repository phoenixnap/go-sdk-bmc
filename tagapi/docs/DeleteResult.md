# DeleteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Tag deletion result message. | 
**TagId** | **string** | The unique identifier of the tag. | 

## Methods

### NewDeleteResult

`func NewDeleteResult(result string, tagId string, ) *DeleteResult`

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


### GetTagId

`func (o *DeleteResult) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *DeleteResult) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *DeleteResult) SetTagId(v string)`

SetTagId sets TagId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


