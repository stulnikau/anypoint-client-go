# PrivateSpacePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | [**Environments4**](Environments4.md) |  | 
**Network** | [**Network5**](Network5.md) |  | 
**ManagedFirewallRules** | Pointer to [**[]FirewallRule**](FirewallRule.md) | Deprecated | [optional] 
**FirewallRules** | [**[]FirewallRule**](FirewallRule.md) | The firewall rules for the Private Space  | 
**LogForwarding** | [**LogForwarding**](LogForwarding.md) |  | 
**IngressConfiguration** | [**IngressConfiguration6**](IngressConfiguration6.md) |  | 
**EnableIAMRole** | Pointer to **bool** | If true, application deployed to this space will have the Private Space IAM role attached to the service account. | [optional] 

## Methods

### NewPrivateSpacePatch

`func NewPrivateSpacePatch(environments Environments4, network Network5, firewallRules []FirewallRule, logForwarding LogForwarding, ingressConfiguration IngressConfiguration6, ) *PrivateSpacePatch`

NewPrivateSpacePatch instantiates a new PrivateSpacePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpacePatchWithDefaults

`func NewPrivateSpacePatchWithDefaults() *PrivateSpacePatch`

NewPrivateSpacePatchWithDefaults instantiates a new PrivateSpacePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *PrivateSpacePatch) GetEnvironments() Environments4`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PrivateSpacePatch) GetEnvironmentsOk() (*Environments4, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PrivateSpacePatch) SetEnvironments(v Environments4)`

SetEnvironments sets Environments field to given value.


### GetNetwork

`func (o *PrivateSpacePatch) GetNetwork() Network5`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PrivateSpacePatch) GetNetworkOk() (*Network5, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PrivateSpacePatch) SetNetwork(v Network5)`

SetNetwork sets Network field to given value.


### GetManagedFirewallRules

`func (o *PrivateSpacePatch) GetManagedFirewallRules() []FirewallRule`

GetManagedFirewallRules returns the ManagedFirewallRules field if non-nil, zero value otherwise.

### GetManagedFirewallRulesOk

`func (o *PrivateSpacePatch) GetManagedFirewallRulesOk() (*[]FirewallRule, bool)`

GetManagedFirewallRulesOk returns a tuple with the ManagedFirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedFirewallRules

`func (o *PrivateSpacePatch) SetManagedFirewallRules(v []FirewallRule)`

SetManagedFirewallRules sets ManagedFirewallRules field to given value.

### HasManagedFirewallRules

`func (o *PrivateSpacePatch) HasManagedFirewallRules() bool`

HasManagedFirewallRules returns a boolean if a field has been set.

### GetFirewallRules

`func (o *PrivateSpacePatch) GetFirewallRules() []FirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *PrivateSpacePatch) GetFirewallRulesOk() (*[]FirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *PrivateSpacePatch) SetFirewallRules(v []FirewallRule)`

SetFirewallRules sets FirewallRules field to given value.


### GetLogForwarding

`func (o *PrivateSpacePatch) GetLogForwarding() LogForwarding`

GetLogForwarding returns the LogForwarding field if non-nil, zero value otherwise.

### GetLogForwardingOk

`func (o *PrivateSpacePatch) GetLogForwardingOk() (*LogForwarding, bool)`

GetLogForwardingOk returns a tuple with the LogForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogForwarding

`func (o *PrivateSpacePatch) SetLogForwarding(v LogForwarding)`

SetLogForwarding sets LogForwarding field to given value.


### GetIngressConfiguration

`func (o *PrivateSpacePatch) GetIngressConfiguration() IngressConfiguration6`

GetIngressConfiguration returns the IngressConfiguration field if non-nil, zero value otherwise.

### GetIngressConfigurationOk

`func (o *PrivateSpacePatch) GetIngressConfigurationOk() (*IngressConfiguration6, bool)`

GetIngressConfigurationOk returns a tuple with the IngressConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressConfiguration

`func (o *PrivateSpacePatch) SetIngressConfiguration(v IngressConfiguration6)`

SetIngressConfiguration sets IngressConfiguration field to given value.


### GetEnableIAMRole

`func (o *PrivateSpacePatch) GetEnableIAMRole() bool`

GetEnableIAMRole returns the EnableIAMRole field if non-nil, zero value otherwise.

### GetEnableIAMRoleOk

`func (o *PrivateSpacePatch) GetEnableIAMRoleOk() (*bool, bool)`

GetEnableIAMRoleOk returns a tuple with the EnableIAMRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIAMRole

`func (o *PrivateSpacePatch) SetEnableIAMRole(v bool)`

SetEnableIAMRole sets EnableIAMRole field to given value.

### HasEnableIAMRole

`func (o *PrivateSpacePatch) HasEnableIAMRole() bool`

HasEnableIAMRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


