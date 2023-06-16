# IAMRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | **[]string** | IAM roles associated with the space | 
**OrganizationId** | **string** | The parent organization ID of the space | 
**SpaceId** | **string** | id of the space. | 
**Message** | Pointer to **string** | IAM roles associated with the space | [optional] 

## Methods

### NewIAMRole

`func NewIAMRole(roles []string, organizationId string, spaceId string, ) *IAMRole`

NewIAMRole instantiates a new IAMRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIAMRoleWithDefaults

`func NewIAMRoleWithDefaults() *IAMRole`

NewIAMRoleWithDefaults instantiates a new IAMRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *IAMRole) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IAMRole) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IAMRole) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetOrganizationId

`func (o *IAMRole) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *IAMRole) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *IAMRole) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetSpaceId

`func (o *IAMRole) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IAMRole) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IAMRole) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetMessage

`func (o *IAMRole) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IAMRole) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IAMRole) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IAMRole) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


