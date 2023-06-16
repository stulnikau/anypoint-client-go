# NodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Location** | **string** |  | 
**IsAvailableForDeployments** | **bool** |  | 
**Availability** | Pointer to **string** | If set to \&quot;AVAILABLE\&quot; the cluster will accept deployments | [optional] 

## Methods

### NewNodeInfo

`func NewNodeInfo(id string, location string, isAvailableForDeployments bool, ) *NodeInfo`

NewNodeInfo instantiates a new NodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeInfoWithDefaults

`func NewNodeInfoWithDefaults() *NodeInfo`

NewNodeInfoWithDefaults instantiates a new NodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeInfo) SetId(v string)`

SetId sets Id field to given value.


### GetLocation

`func (o *NodeInfo) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NodeInfo) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NodeInfo) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetIsAvailableForDeployments

`func (o *NodeInfo) GetIsAvailableForDeployments() bool`

GetIsAvailableForDeployments returns the IsAvailableForDeployments field if non-nil, zero value otherwise.

### GetIsAvailableForDeploymentsOk

`func (o *NodeInfo) GetIsAvailableForDeploymentsOk() (*bool, bool)`

GetIsAvailableForDeploymentsOk returns a tuple with the IsAvailableForDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailableForDeployments

`func (o *NodeInfo) SetIsAvailableForDeployments(v bool)`

SetIsAvailableForDeployments sets IsAvailableForDeployments field to given value.


### GetAvailability

`func (o *NodeInfo) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *NodeInfo) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *NodeInfo) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *NodeInfo) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


