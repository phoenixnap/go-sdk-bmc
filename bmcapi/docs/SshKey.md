# SshKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the SSH Key. | 
**Default** | **bool** | Keys marked as default are always included on server creation and reset unless toggled off in creation/reset request. | 
**Name** | **string** | Friendly SSH key name to represent an SSH key. | 
**Key** | **string** | SSH Key value. | 
**Fingerprint** | **string** | SSH key auto-generated SHA-256 fingerprint. | 
**CreatedOn** | **time.Time** | Date and time of creation. | 
**LastUpdatedOn** | **time.Time** | Date and time of last update. | 

## Methods

### NewSshKey

`func NewSshKey(id string, default_ bool, name string, key string, fingerprint string, createdOn time.Time, lastUpdatedOn time.Time, ) *SshKey`

NewSshKey instantiates a new SshKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeyWithDefaults

`func NewSshKeyWithDefaults() *SshKey`

NewSshKeyWithDefaults instantiates a new SshKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SshKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SshKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SshKey) SetId(v string)`

SetId sets Id field to given value.


### GetDefault

`func (o *SshKey) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SshKey) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SshKey) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetName

`func (o *SshKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SshKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SshKey) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *SshKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SshKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SshKey) SetKey(v string)`

SetKey sets Key field to given value.


### GetFingerprint

`func (o *SshKey) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *SshKey) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *SshKey) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetCreatedOn

`func (o *SshKey) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *SshKey) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *SshKey) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### GetLastUpdatedOn

`func (o *SshKey) GetLastUpdatedOn() time.Time`

GetLastUpdatedOn returns the LastUpdatedOn field if non-nil, zero value otherwise.

### GetLastUpdatedOnOk

`func (o *SshKey) GetLastUpdatedOnOk() (*time.Time, bool)`

GetLastUpdatedOnOk returns a tuple with the LastUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedOn

`func (o *SshKey) SetLastUpdatedOn(v time.Time)`

SetLastUpdatedOn sets LastUpdatedOn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


