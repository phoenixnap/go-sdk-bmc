# RancherClusterConfigCertificates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCertificate** | Pointer to **string** | The SSL CA certificate to be used for rancher admin. | [optional] 
**Certificate** | Pointer to **string** | The SSL certificate to be used for rancher admin. | [optional] 
**CertificateKey** | Pointer to **string** | The SSL certificate key to be used for rancher admin. | [optional] 

## Methods

### NewRancherClusterConfigCertificates

`func NewRancherClusterConfigCertificates() *RancherClusterConfigCertificates`

NewRancherClusterConfigCertificates instantiates a new RancherClusterConfigCertificates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRancherClusterConfigCertificatesWithDefaults

`func NewRancherClusterConfigCertificatesWithDefaults() *RancherClusterConfigCertificates`

NewRancherClusterConfigCertificatesWithDefaults instantiates a new RancherClusterConfigCertificates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCertificate

`func (o *RancherClusterConfigCertificates) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *RancherClusterConfigCertificates) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *RancherClusterConfigCertificates) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *RancherClusterConfigCertificates) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCertificate

`func (o *RancherClusterConfigCertificates) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *RancherClusterConfigCertificates) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *RancherClusterConfigCertificates) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *RancherClusterConfigCertificates) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCertificateKey

`func (o *RancherClusterConfigCertificates) GetCertificateKey() string`

GetCertificateKey returns the CertificateKey field if non-nil, zero value otherwise.

### GetCertificateKeyOk

`func (o *RancherClusterConfigCertificates) GetCertificateKeyOk() (*string, bool)`

GetCertificateKeyOk returns a tuple with the CertificateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateKey

`func (o *RancherClusterConfigCertificates) SetCertificateKey(v string)`

SetCertificateKey sets CertificateKey field to given value.

### HasCertificateKey

`func (o *RancherClusterConfigCertificates) HasCertificateKey() bool`

HasCertificateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


