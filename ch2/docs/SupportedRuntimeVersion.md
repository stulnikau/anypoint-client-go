# SupportedRuntimeVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseVersion** | **string** | Identifier of the Base Version of the Runtime. For example, in the case of a patched 3.8.0 Mule Runtime, the base version is 3.8.0. This is the version that is used in order to get feature validations by the Anypoint Management Center. This is the information shown to customers and the one that has to be used in order to make deployments. | 
**Tag** | **string** | Identifier of the actual version of the runtime that is used for deployments. | 
**ReleaseDate** | **float32** | Release date of the runtime tag that is used for deployments. | 
**MinimumTag** | **string** | Identifier of the oldest version of the runtime that can be used for deployments. | 
**Href** | Pointer to **string** | Href to get information about the particular asset. Only if supported. | [optional] 

## Methods

### NewSupportedRuntimeVersion

`func NewSupportedRuntimeVersion(baseVersion string, tag string, releaseDate float32, minimumTag string, ) *SupportedRuntimeVersion`

NewSupportedRuntimeVersion instantiates a new SupportedRuntimeVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedRuntimeVersionWithDefaults

`func NewSupportedRuntimeVersionWithDefaults() *SupportedRuntimeVersion`

NewSupportedRuntimeVersionWithDefaults instantiates a new SupportedRuntimeVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseVersion

`func (o *SupportedRuntimeVersion) GetBaseVersion() string`

GetBaseVersion returns the BaseVersion field if non-nil, zero value otherwise.

### GetBaseVersionOk

`func (o *SupportedRuntimeVersion) GetBaseVersionOk() (*string, bool)`

GetBaseVersionOk returns a tuple with the BaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVersion

`func (o *SupportedRuntimeVersion) SetBaseVersion(v string)`

SetBaseVersion sets BaseVersion field to given value.


### GetTag

`func (o *SupportedRuntimeVersion) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *SupportedRuntimeVersion) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *SupportedRuntimeVersion) SetTag(v string)`

SetTag sets Tag field to given value.


### GetReleaseDate

`func (o *SupportedRuntimeVersion) GetReleaseDate() float32`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SupportedRuntimeVersion) GetReleaseDateOk() (*float32, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SupportedRuntimeVersion) SetReleaseDate(v float32)`

SetReleaseDate sets ReleaseDate field to given value.


### GetMinimumTag

`func (o *SupportedRuntimeVersion) GetMinimumTag() string`

GetMinimumTag returns the MinimumTag field if non-nil, zero value otherwise.

### GetMinimumTagOk

`func (o *SupportedRuntimeVersion) GetMinimumTagOk() (*string, bool)`

GetMinimumTagOk returns a tuple with the MinimumTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTag

`func (o *SupportedRuntimeVersion) SetMinimumTag(v string)`

SetMinimumTag sets MinimumTag field to given value.


### GetHref

`func (o *SupportedRuntimeVersion) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SupportedRuntimeVersion) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SupportedRuntimeVersion) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SupportedRuntimeVersion) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


