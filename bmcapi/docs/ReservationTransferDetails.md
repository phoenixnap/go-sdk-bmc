# ReservationTransferDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetServerId** | **string** | ID of target server to transfer reservation to. | 

## Methods

### NewReservationTransferDetails

`func NewReservationTransferDetails(targetServerId string, ) *ReservationTransferDetails`

NewReservationTransferDetails instantiates a new ReservationTransferDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationTransferDetailsWithDefaults

`func NewReservationTransferDetailsWithDefaults() *ReservationTransferDetails`

NewReservationTransferDetailsWithDefaults instantiates a new ReservationTransferDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetServerId

`func (o *ReservationTransferDetails) GetTargetServerId() string`

GetTargetServerId returns the TargetServerId field if non-nil, zero value otherwise.

### GetTargetServerIdOk

`func (o *ReservationTransferDetails) GetTargetServerIdOk() (*string, bool)`

GetTargetServerIdOk returns a tuple with the TargetServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetServerId

`func (o *ReservationTransferDetails) SetTargetServerId(v string)`

SetTargetServerId sets TargetServerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


