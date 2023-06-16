# PrivateSpaceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Vpns** | [**[]Items**](Items.md) |  | 

## Methods

### NewPrivateSpaceConnection

`func NewPrivateSpaceConnection(id string, name string, vpns []Items, ) *PrivateSpaceConnection`

NewPrivateSpaceConnection instantiates a new PrivateSpaceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceConnectionWithDefaults

`func NewPrivateSpaceConnectionWithDefaults() *PrivateSpaceConnection`

NewPrivateSpaceConnectionWithDefaults instantiates a new PrivateSpaceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateSpaceConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateSpaceConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateSpaceConnection) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PrivateSpaceConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpaceConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpaceConnection) SetName(v string)`

SetName sets Name field to given value.


### GetVpns

`func (o *PrivateSpaceConnection) GetVpns() []Items`

GetVpns returns the Vpns field if non-nil, zero value otherwise.

### GetVpnsOk

`func (o *PrivateSpaceConnection) GetVpnsOk() (*[]Items, bool)`

GetVpnsOk returns a tuple with the Vpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpns

`func (o *PrivateSpaceConnection) SetVpns(v []Items)`

SetVpns sets Vpns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


