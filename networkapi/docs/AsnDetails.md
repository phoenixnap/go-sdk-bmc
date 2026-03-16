# AsnDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | **int64** | The BGP Peer Group ASN. | 
**IsBringYourOwn** | **bool** | True if the BGP Peer Group ASN is a &#x60;bring your own&#x60; ASN. | 
**VerificationStatus** | **string** | The BGP Peer Group ASN verification status. Can have one of the following values: &#x60;PENDING_VERIFICATION&#x60;, &#x60;FAILED_VERIFICATION&#x60; and &#x60;VERIFIED&#x60;. | 
**VerificationReason** | Pointer to **string** | The BGP Peer Group ASN verification reason for the respective status. | [optional] 

## Methods

### NewAsnDetails

`func NewAsnDetails(asn int64, isBringYourOwn bool, verificationStatus string, ) *AsnDetails`

NewAsnDetails instantiates a new AsnDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsnDetailsWithDefaults

`func NewAsnDetailsWithDefaults() *AsnDetails`

NewAsnDetailsWithDefaults instantiates a new AsnDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *AsnDetails) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *AsnDetails) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *AsnDetails) SetAsn(v int64)`

SetAsn sets Asn field to given value.


### GetIsBringYourOwn

`func (o *AsnDetails) GetIsBringYourOwn() bool`

GetIsBringYourOwn returns the IsBringYourOwn field if non-nil, zero value otherwise.

### GetIsBringYourOwnOk

`func (o *AsnDetails) GetIsBringYourOwnOk() (*bool, bool)`

GetIsBringYourOwnOk returns a tuple with the IsBringYourOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBringYourOwn

`func (o *AsnDetails) SetIsBringYourOwn(v bool)`

SetIsBringYourOwn sets IsBringYourOwn field to given value.


### GetVerificationStatus

`func (o *AsnDetails) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *AsnDetails) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *AsnDetails) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.


### GetVerificationReason

`func (o *AsnDetails) GetVerificationReason() string`

GetVerificationReason returns the VerificationReason field if non-nil, zero value otherwise.

### GetVerificationReasonOk

`func (o *AsnDetails) GetVerificationReasonOk() (*string, bool)`

GetVerificationReasonOk returns a tuple with the VerificationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationReason

`func (o *AsnDetails) SetVerificationReason(v string)`

SetVerificationReason sets VerificationReason field to given value.

### HasVerificationReason

`func (o *AsnDetails) HasVerificationReason() bool`

HasVerificationReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


