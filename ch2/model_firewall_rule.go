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

// checks if the FirewallRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallRule{}

// FirewallRule struct for FirewallRule
type FirewallRule struct {
	CidrBlock string `json:"cidrBlock"`
	Protocol string `json:"protocol"`
	FromPort int32 `json:"fromPort"`
	ToPort int32 `json:"toPort"`
	// Type of the firewall rule. Allowed values are [inbound, outbound]
	Type string `json:"type"`
}

// NewFirewallRule instantiates a new FirewallRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallRule(cidrBlock string, protocol string, fromPort int32, toPort int32, type_ string) *FirewallRule {
	this := FirewallRule{}
	this.CidrBlock = cidrBlock
	this.Protocol = protocol
	this.FromPort = fromPort
	this.ToPort = toPort
	this.Type = type_
	return &this
}

// NewFirewallRuleWithDefaults instantiates a new FirewallRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallRuleWithDefaults() *FirewallRule {
	this := FirewallRule{}
	var type_ string = "inbound"
	this.Type = type_
	return &this
}

// GetCidrBlock returns the CidrBlock field value
func (o *FirewallRule) GetCidrBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CidrBlock
}

// GetCidrBlockOk returns a tuple with the CidrBlock field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetCidrBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CidrBlock, true
}

// SetCidrBlock sets field value
func (o *FirewallRule) SetCidrBlock(v string) {
	o.CidrBlock = v
}

// GetProtocol returns the Protocol field value
func (o *FirewallRule) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *FirewallRule) SetProtocol(v string) {
	o.Protocol = v
}

// GetFromPort returns the FromPort field value
func (o *FirewallRule) GetFromPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FromPort
}

// GetFromPortOk returns a tuple with the FromPort field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetFromPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromPort, true
}

// SetFromPort sets field value
func (o *FirewallRule) SetFromPort(v int32) {
	o.FromPort = v
}

// GetToPort returns the ToPort field value
func (o *FirewallRule) GetToPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToPort
}

// GetToPortOk returns a tuple with the ToPort field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetToPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToPort, true
}

// SetToPort sets field value
func (o *FirewallRule) SetToPort(v int32) {
	o.ToPort = v
}

// GetType returns the Type field value
func (o *FirewallRule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FirewallRule) SetType(v string) {
	o.Type = v
}

func (o FirewallRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cidrBlock"] = o.CidrBlock
	toSerialize["protocol"] = o.Protocol
	toSerialize["fromPort"] = o.FromPort
	toSerialize["toPort"] = o.ToPort
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableFirewallRule struct {
	value *FirewallRule
	isSet bool
}

func (v NullableFirewallRule) Get() *FirewallRule {
	return v.value
}

func (v *NullableFirewallRule) Set(val *FirewallRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallRule(val *FirewallRule) *NullableFirewallRule {
	return &NullableFirewallRule{value: val, isSet: true}
}

func (v NullableFirewallRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


