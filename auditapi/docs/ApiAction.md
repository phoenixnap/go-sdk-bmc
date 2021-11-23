# ApiAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the event. | [optional] 
**Timestamp** | **time.Time** | The UTC time the event initiated. | 
**UserInfo** | [**UserInfo**](UserInfo.md) |  | 
**Response** | [**Response**](Response.md) |  | 
**Request** | [**Request**](Request.md) |  | 

## Methods

### NewApiAction

`func NewApiAction(timestamp time.Time, userInfo UserInfo, response Response, request Request, ) *ApiAction`

NewApiAction instantiates a new ApiAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiActionWithDefaults

`func NewApiActionWithDefaults() *ApiAction`

NewApiActionWithDefaults instantiates a new ApiAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimestamp

`func (o *ApiAction) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiAction) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiAction) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetUserInfo

`func (o *ApiAction) GetUserInfo() UserInfo`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *ApiAction) GetUserInfoOk() (*UserInfo, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *ApiAction) SetUserInfo(v UserInfo)`

SetUserInfo sets UserInfo field to given value.


### GetResponse

`func (o *ApiAction) GetResponse() Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ApiAction) GetResponseOk() (*Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ApiAction) SetResponse(v Response)`

SetResponse sets Response field to given value.


### GetRequest

`func (o *ApiAction) GetRequest() Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ApiAction) GetRequestOk() (*Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ApiAction) SetRequest(v Request)`

SetRequest sets Request field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


