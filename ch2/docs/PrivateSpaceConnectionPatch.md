# PrivateSpaceConnectionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**StaticRoutes** | Pointer to **[]string** | The CIDRs that the Cloudhub VPC will send traffic to via the VPN.  | [optional] 

## Methods

### NewPrivateSpaceConnectionPatch

`func NewPrivateSpaceConnectionPatch(id string, name string, ) *PrivateSpaceConnectionPatch`

NewPrivateSpaceConnectionPatch instantiates a new PrivateSpaceConnectionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceConnectionPatchWithDefaults

`func NewPrivateSpaceConnectionPatchWithDefaults() *PrivateSpaceConnectionPatch`

NewPrivateSpaceConnectionPatchWithDefaults instantiates a new PrivateSpaceConnectionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateSpaceConnectionPatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateSpaceConnectionPatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateSpaceConnectionPatch) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PrivateSpaceConnectionPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpaceConnectionPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpaceConnectionPatch) SetName(v string)`

SetName sets Name field to given value.


### GetStaticRoutes

`func (o *PrivateSpaceConnectionPatch) GetStaticRoutes() []string`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *PrivateSpaceConnectionPatch) GetStaticRoutesOk() (*[]string, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *PrivateSpaceConnectionPatch) SetStaticRoutes(v []string)`

SetStaticRoutes sets StaticRoutes field to given value.

### HasStaticRoutes

`func (o *PrivateSpaceConnectionPatch) HasStaticRoutes() bool`

HasStaticRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


