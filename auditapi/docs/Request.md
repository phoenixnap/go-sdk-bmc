# Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | [**Headers**](Headers.md) |  | 
**Uri** | **string** | The request URI. | 
**Verb** | **string** | The HTTP request verb. | 

## Methods

### NewRequest

`func NewRequest(headers Headers, uri string, verb string, ) *Request`

NewRequest instantiates a new Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWithDefaults

`func NewRequestWithDefaults() *Request`

NewRequestWithDefaults instantiates a new Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *Request) GetHeaders() Headers`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Request) GetHeadersOk() (*Headers, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Request) SetHeaders(v Headers)`

SetHeaders sets Headers field to given value.


### GetUri

`func (o *Request) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Request) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Request) SetUri(v string)`

SetUri sets Uri field to given value.


### GetVerb

`func (o *Request) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *Request) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *Request) SetVerb(v string)`

SetVerb sets Verb field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


