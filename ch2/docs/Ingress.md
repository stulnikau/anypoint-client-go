# Ingress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The last known status | [optional] 
**Message** | Pointer to **string** | Message describing the last known status | [optional] 
**LastSeenTimeStamp** | Pointer to **float32** | Last known timestamp of the status in EPOCH | [optional] 

## Methods

### NewIngress

`func NewIngress() *Ingress`

NewIngress instantiates a new Ingress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressWithDefaults

`func NewIngressWithDefaults() *Ingress`

NewIngressWithDefaults instantiates a new Ingress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Ingress) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ingress) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ingress) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Ingress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *Ingress) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Ingress) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Ingress) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Ingress) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLastSeenTimeStamp

`func (o *Ingress) GetLastSeenTimeStamp() float32`

GetLastSeenTimeStamp returns the LastSeenTimeStamp field if non-nil, zero value otherwise.

### GetLastSeenTimeStampOk

`func (o *Ingress) GetLastSeenTimeStampOk() (*float32, bool)`

GetLastSeenTimeStampOk returns a tuple with the LastSeenTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenTimeStamp

`func (o *Ingress) SetLastSeenTimeStamp(v float32)`

SetLastSeenTimeStamp sets LastSeenTimeStamp field to given value.

### HasLastSeenTimeStamp

`func (o *Ingress) HasLastSeenTimeStamp() bool`

HasLastSeenTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


