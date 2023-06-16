# PrivateSpacePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Private Space. | 
**Environments** | Pointer to [**Environments2**](Environments2.md) |  | [optional] 
**Network** | Pointer to [**Network3**](Network3.md) |  | [optional] 
**ManagedFirewallRules** | Pointer to [**[]FirewallRule**](FirewallRule.md) | Deprecated | [optional] 
**FirewallRules** | Pointer to [**[]FirewallRule**](FirewallRule.md) | The firewall rules for the Private Space  | [optional] 

## Methods

### NewPrivateSpacePost

`func NewPrivateSpacePost(name string, ) *PrivateSpacePost`

NewPrivateSpacePost instantiates a new PrivateSpacePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSpacePostWithDefaults

`func NewPrivateSpacePostWithDefaults() *PrivateSpacePost`

NewPrivateSpacePostWithDefaults instantiates a new PrivateSpacePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateSpacePost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateSpacePost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateSpacePost) SetName(v string)`

SetName sets Name field to given value.


### GetEnvironments

`func (o *PrivateSpacePost) GetEnvironments() Environments2`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PrivateSpacePost) GetEnvironmentsOk() (*Environments2, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PrivateSpacePost) SetEnvironments(v Environments2)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *PrivateSpacePost) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetNetwork

`func (o *PrivateSpacePost) GetNetwork() Network3`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PrivateSpacePost) GetNetworkOk() (*Network3, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PrivateSpacePost) SetNetwork(v Network3)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *PrivateSpacePost) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetManagedFirewallRules

`func (o *PrivateSpacePost) GetManagedFirewallRules() []FirewallRule`

GetManagedFirewallRules returns the ManagedFirewallRules field if non-nil, zero value otherwise.

### GetManagedFirewallRulesOk

`func (o *PrivateSpacePost) GetManagedFirewallRulesOk() (*[]FirewallRule, bool)`

GetManagedFirewallRulesOk returns a tuple with the ManagedFirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedFirewallRules

`func (o *PrivateSpacePost) SetManagedFirewallRules(v []FirewallRule)`

SetManagedFirewallRules sets ManagedFirewallRules field to given value.

### HasManagedFirewallRules

`func (o *PrivateSpacePost) HasManagedFirewallRules() bool`

HasManagedFirewallRules returns a boolean if a field has been set.

### GetFirewallRules

`func (o *PrivateSpacePost) GetFirewallRules() []FirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *PrivateSpacePost) GetFirewallRulesOk() (*[]FirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *PrivateSpacePost) SetFirewallRules(v []FirewallRule)`

SetFirewallRules sets FirewallRules field to given value.

### HasFirewallRules

`func (o *PrivateSpacePost) HasFirewallRules() bool`

HasFirewallRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


