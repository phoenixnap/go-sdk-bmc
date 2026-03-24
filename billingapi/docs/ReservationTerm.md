# ReservationTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LengthInMonths** | **int32** | Term&#39;s length, expressed in months. | 
**ReservationModel** | [**ReservationModelEnum**](ReservationModelEnum.md) |  | 

## Methods

### NewReservationTerm

`func NewReservationTerm(lengthInMonths int32, reservationModel ReservationModelEnum, ) *ReservationTerm`

NewReservationTerm instantiates a new ReservationTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationTermWithDefaults

`func NewReservationTermWithDefaults() *ReservationTerm`

NewReservationTermWithDefaults instantiates a new ReservationTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLengthInMonths

`func (o *ReservationTerm) GetLengthInMonths() int32`

GetLengthInMonths returns the LengthInMonths field if non-nil, zero value otherwise.

### GetLengthInMonthsOk

`func (o *ReservationTerm) GetLengthInMonthsOk() (*int32, bool)`

GetLengthInMonthsOk returns a tuple with the LengthInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthInMonths

`func (o *ReservationTerm) SetLengthInMonths(v int32)`

SetLengthInMonths sets LengthInMonths field to given value.


### GetReservationModel

`func (o *ReservationTerm) GetReservationModel() ReservationModelEnum`

GetReservationModel returns the ReservationModel field if non-nil, zero value otherwise.

### GetReservationModelOk

`func (o *ReservationTerm) GetReservationModelOk() (*ReservationModelEnum, bool)`

GetReservationModelOk returns a tuple with the ReservationModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationModel

`func (o *ReservationTerm) SetReservationModel(v ReservationModelEnum)`

SetReservationModel sets ReservationModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


