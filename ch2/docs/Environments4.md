# Environments4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of environments. Possible values are all, sandbox, production | 
**BusinessGroups** | **[]string** | The Ids of business groups. | 

## Methods

### NewEnvironments4

`func NewEnvironments4(type_ string, businessGroups []string, ) *Environments4`

NewEnvironments4 instantiates a new Environments4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironments4WithDefaults

`func NewEnvironments4WithDefaults() *Environments4`

NewEnvironments4WithDefaults instantiates a new Environments4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Environments4) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Environments4) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Environments4) SetType(v string)`

SetType sets Type field to given value.


### GetBusinessGroups

`func (o *Environments4) GetBusinessGroups() []string`

GetBusinessGroups returns the BusinessGroups field if non-nil, zero value otherwise.

### GetBusinessGroupsOk

`func (o *Environments4) GetBusinessGroupsOk() (*[]string, bool)`

GetBusinessGroupsOk returns a tuple with the BusinessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessGroups

`func (o *Environments4) SetBusinessGroups(v []string)`

SetBusinessGroups sets BusinessGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


