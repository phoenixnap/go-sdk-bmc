# Headers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAgent** | **string** | The UA String | 

## Methods

### NewHeaders

`func NewHeaders(userAgent string, ) *Headers`

NewHeaders instantiates a new Headers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeadersWithDefaults

`func NewHeadersWithDefaults() *Headers`

NewHeadersWithDefaults instantiates a new Headers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAgent

`func (o *Headers) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *Headers) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *Headers) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


