# TgwSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceShare** | [**ResourceShare**](ResourceShare.md) |  | 
**Region** | **string** |  | 
**SpaceName** | Pointer to **string** |  | [optional] 

## Methods

### NewTgwSpec

`func NewTgwSpec(resourceShare ResourceShare, region string, ) *TgwSpec`

NewTgwSpec instantiates a new TgwSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTgwSpecWithDefaults

`func NewTgwSpecWithDefaults() *TgwSpec`

NewTgwSpecWithDefaults instantiates a new TgwSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceShare

`func (o *TgwSpec) GetResourceShare() ResourceShare`

GetResourceShare returns the ResourceShare field if non-nil, zero value otherwise.

### GetResourceShareOk

`func (o *TgwSpec) GetResourceShareOk() (*ResourceShare, bool)`

GetResourceShareOk returns a tuple with the ResourceShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceShare

`func (o *TgwSpec) SetResourceShare(v ResourceShare)`

SetResourceShare sets ResourceShare field to given value.


### GetRegion

`func (o *TgwSpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TgwSpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TgwSpec) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSpaceName

`func (o *TgwSpec) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *TgwSpec) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *TgwSpec) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *TgwSpec) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


