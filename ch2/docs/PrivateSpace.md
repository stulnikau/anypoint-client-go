# PrivateSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Private Space. | 
**Name** | **string** | The name of the Private Space. | 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** | The last reported infra status message. | [optional] 
**Provisioning** | Pointer to [**PrivateSpaceProvisioning**](PrivateSpaceProvisioning.md) |  | [optional] 
**Region** | Pointer to **string** | [Deprecated - will be removed when UI is updated] The region of the Private Space  | [optional] 
**OrganizationId** | **string** | The parent organization ID of the Private Space | 
**ManagedFirewallRules** | Pointer to [**[]FirewallRule**](FirewallRule.md) | [Deprecated - will be removed when UI is updated] The firewall rules for the Private Space  | [optional] 
**Environments** | [**Environments**](Environments.md) |  | 
**Network** | [**Network**](Network.md) |  | 
**FirewallRules** | [**[]FirewallRule**](FirewallRule.md) | The firewall rules for the Private Space  | 
**IngressConfiguration** | [**IngressConfiguration**](IngressConfiguration.md) |  | 
**EnableIAMRole** | Pointer to **bool** | If true, application deployed to this space will have the Private Space IAM role attached to the service account. | [optional] 
**GlobalSpaceStatus** | Pointer to [**GlobalSpaceStatus**](GlobalSpaceStatus.md) |  | [optional] 

## Methods

### NewPrivateSpace

`func NewPrivateSpace(id string, name string, organizationId string, environments Environments, network Network, firewallRules []FirewallRule, ingressConfiguration IngressConfiguration, ) *PrivateSpace`

NewPrivateSpace instantiates a new PrivateSpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpaceWithDefaults

`func NewPrivateSpaceWithDefaults() *PrivateSpace`

NewPrivateSpaceWithDefaults instantiates a new PrivateSpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateSpace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateSpace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateSpace) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PrivateSpace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpace) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *PrivateSpace) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivateSpace) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivateSpace) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrivateSpace) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *PrivateSpace) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *PrivateSpace) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *PrivateSpace) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *PrivateSpace) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetProvisioning

`func (o *PrivateSpace) GetProvisioning() PrivateSpaceProvisioning`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *PrivateSpace) GetProvisioningOk() (*PrivateSpaceProvisioning, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *PrivateSpace) SetProvisioning(v PrivateSpaceProvisioning)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *PrivateSpace) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetRegion

`func (o *PrivateSpace) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PrivateSpace) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PrivateSpace) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PrivateSpace) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetOrganizationId

`func (o *PrivateSpace) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PrivateSpace) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PrivateSpace) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetManagedFirewallRules

`func (o *PrivateSpace) GetManagedFirewallRules() []FirewallRule`

GetManagedFirewallRules returns the ManagedFirewallRules field if non-nil, zero value otherwise.

### GetManagedFirewallRulesOk

`func (o *PrivateSpace) GetManagedFirewallRulesOk() (*[]FirewallRule, bool)`

GetManagedFirewallRulesOk returns a tuple with the ManagedFirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedFirewallRules

`func (o *PrivateSpace) SetManagedFirewallRules(v []FirewallRule)`

SetManagedFirewallRules sets ManagedFirewallRules field to given value.

### HasManagedFirewallRules

`func (o *PrivateSpace) HasManagedFirewallRules() bool`

HasManagedFirewallRules returns a boolean if a field has been set.

### GetEnvironments

`func (o *PrivateSpace) GetEnvironments() Environments`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PrivateSpace) GetEnvironmentsOk() (*Environments, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PrivateSpace) SetEnvironments(v Environments)`

SetEnvironments sets Environments field to given value.


### GetNetwork

`func (o *PrivateSpace) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PrivateSpace) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PrivateSpace) SetNetwork(v Network)`

SetNetwork sets Network field to given value.


### GetFirewallRules

`func (o *PrivateSpace) GetFirewallRules() []FirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *PrivateSpace) GetFirewallRulesOk() (*[]FirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *PrivateSpace) SetFirewallRules(v []FirewallRule)`

SetFirewallRules sets FirewallRules field to given value.


### GetIngressConfiguration

`func (o *PrivateSpace) GetIngressConfiguration() IngressConfiguration`

GetIngressConfiguration returns the IngressConfiguration field if non-nil, zero value otherwise.

### GetIngressConfigurationOk

`func (o *PrivateSpace) GetIngressConfigurationOk() (*IngressConfiguration, bool)`

GetIngressConfigurationOk returns a tuple with the IngressConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressConfiguration

`func (o *PrivateSpace) SetIngressConfiguration(v IngressConfiguration)`

SetIngressConfiguration sets IngressConfiguration field to given value.


### GetEnableIAMRole

`func (o *PrivateSpace) GetEnableIAMRole() bool`

GetEnableIAMRole returns the EnableIAMRole field if non-nil, zero value otherwise.

### GetEnableIAMRoleOk

`func (o *PrivateSpace) GetEnableIAMRoleOk() (*bool, bool)`

GetEnableIAMRoleOk returns a tuple with the EnableIAMRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIAMRole

`func (o *PrivateSpace) SetEnableIAMRole(v bool)`

SetEnableIAMRole sets EnableIAMRole field to given value.

### HasEnableIAMRole

`func (o *PrivateSpace) HasEnableIAMRole() bool`

HasEnableIAMRole returns a boolean if a field has been set.

### GetGlobalSpaceStatus

`func (o *PrivateSpace) GetGlobalSpaceStatus() GlobalSpaceStatus`

GetGlobalSpaceStatus returns the GlobalSpaceStatus field if non-nil, zero value otherwise.

### GetGlobalSpaceStatusOk

`func (o *PrivateSpace) GetGlobalSpaceStatusOk() (*GlobalSpaceStatus, bool)`

GetGlobalSpaceStatusOk returns a tuple with the GlobalSpaceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSpaceStatus

`func (o *PrivateSpace) SetGlobalSpaceStatus(v GlobalSpaceStatus)`

SetGlobalSpaceStatus sets GlobalSpaceStatus field to given value.

### HasGlobalSpaceStatus

`func (o *PrivateSpace) HasGlobalSpaceStatus() bool`

HasGlobalSpaceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


