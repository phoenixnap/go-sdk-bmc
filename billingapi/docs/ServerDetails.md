# ServerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The server identifier as returned by the BMC API. | 
**Hostname** | **string** | The server hostname. | 

## Methods

### NewServerDetails

`func NewServerDetails(id string, hostname string, ) *ServerDetails`

NewServerDetails instantiates a new ServerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDetailsWithDefaults

`func NewServerDetailsWithDefaults() *ServerDetails`

NewServerDetailsWithDefaults instantiates a new ServerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerDetails) SetId(v string)`

SetId sets Id field to given value.


### GetHostname

`func (o *ServerDetails) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerDetails) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerDetails) SetHostname(v string)`

SetHostname sets Hostname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


