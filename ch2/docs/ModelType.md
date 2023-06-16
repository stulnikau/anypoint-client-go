# ModelType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Human readable error description.  | 
**Errors** | Pointer to [**[]TypeErrorsInner**](TypeErrorsInner.md) | Error list use to describe the set of validations that has failed. Only required in case HTTP error code is not enough to describe the failure case.  | [optional] 

## Methods

### NewModelType

`func NewModelType(message string, ) *ModelType`

NewModelType instantiates a new ModelType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTypeWithDefaults

`func NewModelTypeWithDefaults() *ModelType`

NewModelTypeWithDefaults instantiates a new ModelType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ModelType) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelType) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelType) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *ModelType) GetErrors() []TypeErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ModelType) GetErrorsOk() (*[]TypeErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ModelType) SetErrors(v []TypeErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ModelType) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


