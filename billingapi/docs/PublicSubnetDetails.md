# PublicSubnetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Public Subnet identifier as returned by the BMC API. | [optional] 
**Cidr** | **string** | Classless Inter-Domain Routing | 
**Size** | **string** | CIDR size | 

## Methods

### NewPublicSubnetDetails

`func NewPublicSubnetDetails(cidr string, size string, ) *PublicSubnetDetails`

NewPublicSubnetDetails instantiates a new PublicSubnetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSubnetDetailsWithDefaults

`func NewPublicSubnetDetailsWithDefaults() *PublicSubnetDetails`

NewPublicSubnetDetailsWithDefaults instantiates a new PublicSubnetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicSubnetDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicSubnetDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicSubnetDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicSubnetDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCidr

`func (o *PublicSubnetDetails) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *PublicSubnetDetails) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *PublicSubnetDetails) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetSize

`func (o *PublicSubnetDetails) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PublicSubnetDetails) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PublicSubnetDetails) SetSize(v string)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


