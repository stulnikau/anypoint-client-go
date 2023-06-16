# Validity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotBefore** | **string** |  | 
**NotAfter** | **string** |  | 

## Methods

### NewValidity

`func NewValidity(notBefore string, notAfter string, ) *Validity`

NewValidity instantiates a new Validity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidityWithDefaults

`func NewValidityWithDefaults() *Validity`

NewValidityWithDefaults instantiates a new Validity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotBefore

`func (o *Validity) GetNotBefore() string`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *Validity) GetNotBeforeOk() (*string, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *Validity) SetNotBefore(v string)`

SetNotBefore sets NotBefore field to given value.


### GetNotAfter

`func (o *Validity) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *Validity) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *Validity) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


