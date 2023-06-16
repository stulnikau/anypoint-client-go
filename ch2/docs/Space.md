# Space

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The last known status | [optional] 
**Message** | Pointer to **string** | Message describing the last known status | [optional] 
**LastSeenTimeStamp** | Pointer to **float32** | Last known timestamp of the status in EPOCH | [optional] 

## Methods

### NewSpace

`func NewSpace() *Space`

NewSpace instantiates a new Space object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceWithDefaults

`func NewSpaceWithDefaults() *Space`

NewSpaceWithDefaults instantiates a new Space object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Space) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Space) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Space) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Space) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *Space) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Space) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Space) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Space) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLastSeenTimeStamp

`func (o *Space) GetLastSeenTimeStamp() float32`

GetLastSeenTimeStamp returns the LastSeenTimeStamp field if non-nil, zero value otherwise.

### GetLastSeenTimeStampOk

`func (o *Space) GetLastSeenTimeStampOk() (*float32, bool)`

GetLastSeenTimeStampOk returns a tuple with the LastSeenTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTimeStamp

`func (o *Space) SetLastSeenTimeStamp(v float32)`

SetLastSeenTimeStamp sets LastSeenTimeStamp field to given value.

### HasLastSeenTimeStamp

`func (o *Space) HasLastSeenTimeStamp() bool`

HasLastSeenTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


