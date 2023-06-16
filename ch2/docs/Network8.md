# Network8

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The last known status | [optional] 
**Message** | Pointer to **string** | Message describing the last known status | [optional] 
**LastSeenTimeStamp** | Pointer to **float32** | Last known timestamp of the status in EPOCH | [optional] 

## Methods

### NewNetwork8

`func NewNetwork8() *Network8`

NewNetwork8 instantiates a new Network8 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetwork8WithDefaults

`func NewNetwork8WithDefaults() *Network8`

NewNetwork8WithDefaults instantiates a new Network8 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Network8) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Network8) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Network8) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Network8) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *Network8) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Network8) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Network8) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Network8) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLastSeenTimeStamp

`func (o *Network8) GetLastSeenTimeStamp() float32`

GetLastSeenTimeStamp returns the LastSeenTimeStamp field if non-nil, zero value otherwise.

### GetLastSeenTimeStampOk

`func (o *Network8) GetLastSeenTimeStampOk() (*float32, bool)`

GetLastSeenTimeStampOk returns a tuple with the LastSeenTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTimeStamp

`func (o *Network8) SetLastSeenTimeStamp(v float32)`

SetLastSeenTimeStamp sets LastSeenTimeStamp field to given value.

### HasLastSeenTimeStamp

`func (o *Network8) HasLastSeenTimeStamp() bool`

HasLastSeenTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


