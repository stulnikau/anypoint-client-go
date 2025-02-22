/*
CH2 Management API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PrivateSpacePost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateSpacePost{}

// PrivateSpacePost struct for PrivateSpacePost
type PrivateSpacePost struct {
	// The name of the Private Space.
	Name string `json:"name"`
	Environments *Environments2 `json:"environments,omitempty"`
	Network *Network3 `json:"network,omitempty"`
	// Deprecated
	ManagedFirewallRules []FirewallRule `json:"managedFirewallRules,omitempty"`
	// The firewall rules for the Private Space 
	FirewallRules []FirewallRule `json:"firewallRules,omitempty"`
}

// NewPrivateSpacePost instantiates a new PrivateSpacePost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateSpacePost(name string) *PrivateSpacePost {
	this := PrivateSpacePost{}
	this.Name = name
	return &this
}

// NewPrivateSpacePostWithDefaults instantiates a new PrivateSpacePost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateSpacePostWithDefaults() *PrivateSpacePost {
	this := PrivateSpacePost{}
	return &this
}

// GetName returns the Name field value
func (o *PrivateSpacePost) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrivateSpacePost) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrivateSpacePost) SetName(v string) {
	o.Name = v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *PrivateSpacePost) GetEnvironments() Environments2 {
	if o == nil || IsNil(o.Environments) {
		var ret Environments2
		return ret
	}
	return *o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSpacePost) GetEnvironmentsOk() (*Environments2, bool) {
	if o == nil || IsNil(o.Environments) {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *PrivateSpacePost) HasEnvironments() bool {
	if o != nil && !IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given Environments2 and assigns it to the Environments field.
func (o *PrivateSpacePost) SetEnvironments(v Environments2) {
	o.Environments = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *PrivateSpacePost) GetNetwork() Network3 {
	if o == nil || IsNil(o.Network) {
		var ret Network3
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSpacePost) GetNetworkOk() (*Network3, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *PrivateSpacePost) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given Network3 and assigns it to the Network field.
func (o *PrivateSpacePost) SetNetwork(v Network3) {
	o.Network = &v
}

// GetManagedFirewallRules returns the ManagedFirewallRules field value if set, zero value otherwise.
func (o *PrivateSpacePost) GetManagedFirewallRules() []FirewallRule {
	if o == nil || IsNil(o.ManagedFirewallRules) {
		var ret []FirewallRule
		return ret
	}
	return o.ManagedFirewallRules
}

// GetManagedFirewallRulesOk returns a tuple with the ManagedFirewallRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSpacePost) GetManagedFirewallRulesOk() ([]FirewallRule, bool) {
	if o == nil || IsNil(o.ManagedFirewallRules) {
		return nil, false
	}
	return o.ManagedFirewallRules, true
}

// HasManagedFirewallRules returns a boolean if a field has been set.
func (o *PrivateSpacePost) HasManagedFirewallRules() bool {
	if o != nil && !IsNil(o.ManagedFirewallRules) {
		return true
	}

	return false
}

// SetManagedFirewallRules gets a reference to the given []FirewallRule and assigns it to the ManagedFirewallRules field.
func (o *PrivateSpacePost) SetManagedFirewallRules(v []FirewallRule) {
	o.ManagedFirewallRules = v
}

// GetFirewallRules returns the FirewallRules field value if set, zero value otherwise.
func (o *PrivateSpacePost) GetFirewallRules() []FirewallRule {
	if o == nil || IsNil(o.FirewallRules) {
		var ret []FirewallRule
		return ret
	}
	return o.FirewallRules
}

// GetFirewallRulesOk returns a tuple with the FirewallRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSpacePost) GetFirewallRulesOk() ([]FirewallRule, bool) {
	if o == nil || IsNil(o.FirewallRules) {
		return nil, false
	}
	return o.FirewallRules, true
}

// HasFirewallRules returns a boolean if a field has been set.
func (o *PrivateSpacePost) HasFirewallRules() bool {
	if o != nil && !IsNil(o.FirewallRules) {
		return true
	}

	return false
}

// SetFirewallRules gets a reference to the given []FirewallRule and assigns it to the FirewallRules field.
func (o *PrivateSpacePost) SetFirewallRules(v []FirewallRule) {
	o.FirewallRules = v
}

func (o PrivateSpacePost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateSpacePost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Environments) {
		toSerialize["environments"] = o.Environments
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.ManagedFirewallRules) {
		toSerialize["managedFirewallRules"] = o.ManagedFirewallRules
	}
	if !IsNil(o.FirewallRules) {
		toSerialize["firewallRules"] = o.FirewallRules
	}
	return toSerialize, nil
}

type NullablePrivateSpacePost struct {
	value *PrivateSpacePost
	isSet bool
}

func (v NullablePrivateSpacePost) Get() *PrivateSpacePost {
	return v.value
}

func (v *NullablePrivateSpacePost) Set(val *PrivateSpacePost) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateSpacePost) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateSpacePost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateSpacePost(val *PrivateSpacePost) *NullablePrivateSpacePost {
	return &NullablePrivateSpacePost{value: val, isSet: true}
}

func (v NullablePrivateSpacePost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateSpacePost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


