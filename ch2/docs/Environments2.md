# Environments2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of environments. Possible values are all, sandbox, production | 
**BusinessGroups** | **[]string** | The Ids of business groups. | 

## Methods

### NewEnvironments2

`func NewEnvironments2(type_ string, businessGroups []string, ) *Environments2`

NewEnvironments2 instantiates a new Environments2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironments2WithDefaults

`func NewEnvironments2WithDefaults() *Environments2`

NewEnvironments2WithDefaults instantiates a new Environments2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Environments2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Environments2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Environments2) SetType(v string)`

SetType sets Type field to given value.


### GetBusinessGroups

`func (o *Environments2) GetBusinessGroups() []string`

GetBusinessGroups returns the BusinessGroups field if non-nil, zero value otherwise.

### GetBusinessGroupsOk

`func (o *Environments2) GetBusinessGroupsOk() (*[]string, bool)`

GetBusinessGroupsOk returns a tuple with the BusinessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessGroups

`func (o *Environments2) SetBusinessGroups(v []string)`

SetBusinessGroups sets BusinessGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


