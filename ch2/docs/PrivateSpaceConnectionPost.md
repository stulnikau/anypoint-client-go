# PrivateSpaceConnectionPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Vpns** | [**[]Items**](Items.md) |  | 

## Methods

### NewPrivateSpaceConnectionPost

`func NewPrivateSpaceConnectionPost(name string, vpns []Items, ) *PrivateSpaceConnectionPost`

NewPrivateSpaceConnectionPost instantiates a new PrivateSpaceConnectionPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceConnectionPostWithDefaults

`func NewPrivateSpaceConnectionPostWithDefaults() *PrivateSpaceConnectionPost`

NewPrivateSpaceConnectionPostWithDefaults instantiates a new PrivateSpaceConnectionPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateSpaceConnectionPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpaceConnectionPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpaceConnectionPost) SetName(v string)`

SetName sets Name field to given value.


### GetVpns

`func (o *PrivateSpaceConnectionPost) GetVpns() []Items`

GetVpns returns the Vpns field if non-nil, zero value otherwise.

### GetVpnsOk

`func (o *PrivateSpaceConnectionPost) GetVpnsOk() (*[]Items, bool)`

GetVpnsOk returns a tuple with the Vpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpns

`func (o *PrivateSpaceConnectionPost) SetVpns(v []Items)`

SetVpns sets Vpns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


