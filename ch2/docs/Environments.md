# Environments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of environments. Possible values are all, sandbox, production | 
**BusinessGroups** | **[]string** | The Ids of business groups. | 

## Methods

### NewEnvironments

`func NewEnvironments(type_ string, businessGroups []string, ) *Environments`

NewEnvironments instantiates a new Environments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentsWithDefaults

`func NewEnvironmentsWithDefaults() *Environments`

NewEnvironmentsWithDefaults instantiates a new Environments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Environments) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Environments) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Environments) SetType(v string)`

SetType sets Type field to given value.


### GetBusinessGroups

`func (o *Environments) GetBusinessGroups() []string`

GetBusinessGroups returns the BusinessGroups field if non-nil, zero value otherwise.

### GetBusinessGroupsOk

`func (o *Environments) GetBusinessGroupsOk() (*[]string, bool)`

GetBusinessGroupsOk returns a tuple with the BusinessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessGroups

`func (o *Environments) SetBusinessGroups(v []string)`

SetBusinessGroups sets BusinessGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


