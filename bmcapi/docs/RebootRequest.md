# RebootRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootType** | Pointer to **string** | Specifies whether to boot via &#x60;IPXE&#x60; (requires script) or &#x60;STANDARD&#x60; (default mechanism, incompatible with &#x60;ipxeUrl&#x60;). | [optional] [default to "STANDARD"]
**IpxeUrl** | Pointer to **NullableString** | The URL for the iPXE script, used only with &#x60;IPXE&#x60; boot type. If provided, it updates and replaces the existing stored URL; if not provided, the existing URL will be used. | [optional] 

## Methods

### NewRebootRequest

`func NewRebootRequest() *RebootRequest`

NewRebootRequest instantiates a new RebootRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootRequestWithDefaults

`func NewRebootRequestWithDefaults() *RebootRequest`

NewRebootRequestWithDefaults instantiates a new RebootRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootType

`func (o *RebootRequest) GetBootType() string`

GetBootType returns the BootType field if non-nil, zero value otherwise.

### GetBootTypeOk

`func (o *RebootRequest) GetBootTypeOk() (*string, bool)`

GetBootTypeOk returns a tuple with the BootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootType

`func (o *RebootRequest) SetBootType(v string)`

SetBootType sets BootType field to given value.

### HasBootType

`func (o *RebootRequest) HasBootType() bool`

HasBootType returns a boolean if a field has been set.

### GetIpxeUrl

`func (o *RebootRequest) GetIpxeUrl() string`

GetIpxeUrl returns the IpxeUrl field if non-nil, zero value otherwise.

### GetIpxeUrlOk

`func (o *RebootRequest) GetIpxeUrlOk() (*string, bool)`

GetIpxeUrlOk returns a tuple with the IpxeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpxeUrl

`func (o *RebootRequest) SetIpxeUrl(v string)`

SetIpxeUrl sets IpxeUrl field to given value.

### HasIpxeUrl

`func (o *RebootRequest) HasIpxeUrl() bool`

HasIpxeUrl returns a boolean if a field has been set.

### SetIpxeUrlNil

`func (o *RebootRequest) SetIpxeUrlNil(b bool)`

 SetIpxeUrlNil sets the value for IpxeUrl to be an explicit nil

### UnsetIpxeUrl
`func (o *RebootRequest) UnsetIpxeUrl()`

UnsetIpxeUrl ensures that no value is present for IpxeUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


