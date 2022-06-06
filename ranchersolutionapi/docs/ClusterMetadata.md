# ClusterMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The Rancher Server URL. | [optional] 
**Username** | Pointer to **string** | The username to use to login to the Rancher Server. This field is returned only as a response to the create cluster request. Make sure to take note or you will not be able to access the server. | [optional] 
**Password** | Pointer to **string** | This is the password to be used to login to the Rancher Server. This field is returned only as a response to the create cluster request. Make sure to take note or you will not be able to access the server. | [optional] 

## Methods

### NewClusterMetadata

`func NewClusterMetadata() *ClusterMetadata`

NewClusterMetadata instantiates a new ClusterMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMetadataWithDefaults

`func NewClusterMetadataWithDefaults() *ClusterMetadata`

NewClusterMetadataWithDefaults instantiates a new ClusterMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ClusterMetadata) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ClusterMetadata) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ClusterMetadata) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ClusterMetadata) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *ClusterMetadata) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ClusterMetadata) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ClusterMetadata) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ClusterMetadata) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ClusterMetadata) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ClusterMetadata) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ClusterMetadata) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ClusterMetadata) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


