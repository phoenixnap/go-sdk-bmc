# ReservationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Reservation identifier. | [optional] 
**Quantity** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 

## Methods

### NewReservationDetails

`func NewReservationDetails() *ReservationDetails`

NewReservationDetails instantiates a new ReservationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationDetailsWithDefaults

`func NewReservationDetailsWithDefaults() *ReservationDetails`

NewReservationDetailsWithDefaults instantiates a new ReservationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReservationDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReservationDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReservationDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReservationDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuantity

`func (o *ReservationDetails) GetQuantity() Quantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReservationDetails) GetQuantityOk() (*Quantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReservationDetails) SetQuantity(v Quantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReservationDetails) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


