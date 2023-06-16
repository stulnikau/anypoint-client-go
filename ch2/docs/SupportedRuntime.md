# SupportedRuntime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Versions** | Pointer to [**[]SupportedRuntimeVersion**](SupportedRuntimeVersion.md) |  | [optional] 

## Methods

### NewSupportedRuntime

`func NewSupportedRuntime(type_ string, ) *SupportedRuntime`

NewSupportedRuntime instantiates a new SupportedRuntime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedRuntimeWithDefaults

`func NewSupportedRuntimeWithDefaults() *SupportedRuntime`

NewSupportedRuntimeWithDefaults instantiates a new SupportedRuntime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SupportedRuntime) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SupportedRuntime) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SupportedRuntime) SetType(v string)`

SetType sets Type field to given value.


### GetVersions

`func (o *SupportedRuntime) GetVersions() []SupportedRuntimeVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *SupportedRuntime) GetVersionsOk() (*[]SupportedRuntimeVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *SupportedRuntime) SetVersions(v []SupportedRuntimeVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *SupportedRuntime) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


