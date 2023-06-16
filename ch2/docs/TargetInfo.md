# TargetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Version** | **string** | The version of the target. | 
**Type** | **string** | The type of the target. | 
**Runtimes** | [**[]SupportedRuntime**](SupportedRuntime.md) |  | 
**Status** | **string** |  | 
**Environments** | Pointer to **[]string** | If present determines if the target is restricted to a particular set of environments. If it is restricted to all environments then an empty list has to be provided. If not present then there is not target restriction over an environment. | [optional] 
**IsAvailableForDeployments** | **bool** | Determines if a deployment can be perfomed in a particular target. | 
**ReplicationStrategies** | Pointer to **map[string]interface{}** | Replication Strategies for this target per deployment type. | [optional] 
**Nodes** | [**[]NodeInfo**](NodeInfo.md) |  | 
**Defaults** | Pointer to **map[string]interface{}** | Default deployment settings for this target  | [optional] 
**EnhancedSecurity** | Pointer to **bool** | Flag for whether the target supports enhanced security features.  | [optional] 
**FeatureFlags** | Pointer to **map[string]interface{}** | Supported Feature flags for the target  | [optional] 
**Region** | **string** | Target Region | 

## Methods

### NewTargetInfo

`func NewTargetInfo(id string, name string, version string, type_ string, runtimes []SupportedRuntime, status string, isAvailableForDeployments bool, nodes []NodeInfo, region string, ) *TargetInfo`

NewTargetInfo instantiates a new TargetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetInfoWithDefaults

`func NewTargetInfoWithDefaults() *TargetInfo`

NewTargetInfoWithDefaults instantiates a new TargetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TargetInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TargetInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TargetInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TargetInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetInfo) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *TargetInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TargetInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TargetInfo) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetType

`func (o *TargetInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TargetInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TargetInfo) SetType(v string)`

SetType sets Type field to given value.


### GetRuntimes

`func (o *TargetInfo) GetRuntimes() []SupportedRuntime`

GetRuntimes returns the Runtimes field if non-nil, zero value otherwise.

### GetRuntimesOk

`func (o *TargetInfo) GetRuntimesOk() (*[]SupportedRuntime, bool)`

GetRuntimesOk returns a tuple with the Runtimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimes

`func (o *TargetInfo) SetRuntimes(v []SupportedRuntime)`

SetRuntimes sets Runtimes field to given value.


### GetStatus

`func (o *TargetInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TargetInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TargetInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEnvironments

`func (o *TargetInfo) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *TargetInfo) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *TargetInfo) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *TargetInfo) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetIsAvailableForDeployments

`func (o *TargetInfo) GetIsAvailableForDeployments() bool`

GetIsAvailableForDeployments returns the IsAvailableForDeployments field if non-nil, zero value otherwise.

### GetIsAvailableForDeploymentsOk

`func (o *TargetInfo) GetIsAvailableForDeploymentsOk() (*bool, bool)`

GetIsAvailableForDeploymentsOk returns a tuple with the IsAvailableForDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailableForDeployments

`func (o *TargetInfo) SetIsAvailableForDeployments(v bool)`

SetIsAvailableForDeployments sets IsAvailableForDeployments field to given value.


### GetReplicationStrategies

`func (o *TargetInfo) GetReplicationStrategies() map[string]interface{}`

GetReplicationStrategies returns the ReplicationStrategies field if non-nil, zero value otherwise.

### GetReplicationStrategiesOk

`func (o *TargetInfo) GetReplicationStrategiesOk() (*map[string]interface{}, bool)`

GetReplicationStrategiesOk returns a tuple with the ReplicationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStrategies

`func (o *TargetInfo) SetReplicationStrategies(v map[string]interface{})`

SetReplicationStrategies sets ReplicationStrategies field to given value.

### HasReplicationStrategies

`func (o *TargetInfo) HasReplicationStrategies() bool`

HasReplicationStrategies returns a boolean if a field has been set.

### GetNodes

`func (o *TargetInfo) GetNodes() []NodeInfo`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *TargetInfo) GetNodesOk() (*[]NodeInfo, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *TargetInfo) SetNodes(v []NodeInfo)`

SetNodes sets Nodes field to given value.


### GetDefaults

`func (o *TargetInfo) GetDefaults() map[string]interface{}`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *TargetInfo) GetDefaultsOk() (*map[string]interface{}, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *TargetInfo) SetDefaults(v map[string]interface{})`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *TargetInfo) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetEnhancedSecurity

`func (o *TargetInfo) GetEnhancedSecurity() bool`

GetEnhancedSecurity returns the EnhancedSecurity field if non-nil, zero value otherwise.

### GetEnhancedSecurityOk

`func (o *TargetInfo) GetEnhancedSecurityOk() (*bool, bool)`

GetEnhancedSecurityOk returns a tuple with the EnhancedSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedSecurity

`func (o *TargetInfo) SetEnhancedSecurity(v bool)`

SetEnhancedSecurity sets EnhancedSecurity field to given value.

### HasEnhancedSecurity

`func (o *TargetInfo) HasEnhancedSecurity() bool`

HasEnhancedSecurity returns a boolean if a field has been set.

### GetFeatureFlags

`func (o *TargetInfo) GetFeatureFlags() map[string]interface{}`

GetFeatureFlags returns the FeatureFlags field if non-nil, zero value otherwise.

### GetFeatureFlagsOk

`func (o *TargetInfo) GetFeatureFlagsOk() (*map[string]interface{}, bool)`

GetFeatureFlagsOk returns a tuple with the FeatureFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlags

`func (o *TargetInfo) SetFeatureFlags(v map[string]interface{})`

SetFeatureFlags sets FeatureFlags field to given value.

### HasFeatureFlags

`func (o *TargetInfo) HasFeatureFlags() bool`

HasFeatureFlags returns a boolean if a field has been set.

### GetRegion

`func (o *TargetInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TargetInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TargetInfo) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


