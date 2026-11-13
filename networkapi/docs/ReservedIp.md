# ReservedIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | **[]string** | A list containing reserved IP addresses or inclusive IP ranges. | 
**AssignableWithForce** | **bool** | Indicates whether the reserved IP addresses can be assigned to a resource when the assignment request is made with &#x60;force&#x60; query parameter set to &#x60;true&#x60;. | 

## Methods

### NewReservedIp

`func NewReservedIp(ips []string, assignableWithForce bool, ) *ReservedIp`

NewReservedIp instantiates a new ReservedIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservedIpWithDefaults

`func NewReservedIpWithDefaults() *ReservedIp`

NewReservedIpWithDefaults instantiates a new ReservedIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *ReservedIp) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *ReservedIp) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *ReservedIp) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetAssignableWithForce

`func (o *ReservedIp) GetAssignableWithForce() bool`

GetAssignableWithForce returns the AssignableWithForce field if non-nil, zero value otherwise.

### GetAssignableWithForceOk

`func (o *ReservedIp) GetAssignableWithForceOk() (*bool, bool)`

GetAssignableWithForceOk returns a tuple with the AssignableWithForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableWithForce

`func (o *ReservedIp) SetAssignableWithForce(v bool)`

SetAssignableWithForce sets AssignableWithForce field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


