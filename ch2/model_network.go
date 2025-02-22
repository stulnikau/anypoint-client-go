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

// checks if the Network type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Network{}

// Network Private space network configuration 
type Network struct {
	// The region of the Private Space. Required when create a Private Space network.
	Region *string `json:"region,omitempty"`
	// The existing VPC where the Private Space will be created. Required when create a Private Space in existing Anypoint VPC. 
	VpcId *string `json:"vpcId,omitempty"`
	// The CIDR block of the VPC where the Private Space will be created. Required when create a Private Space in a new VPC. Default is 10.0.0.0/16. 
	CidrBlock *string `json:"cidrBlock,omitempty"`
	InternalDns *InternalDns `json:"internalDns,omitempty"`
	// The inbound static IPs of the Private Space network. 
	InboundStaticIps []string `json:"inboundStaticIps,omitempty"`
	// The outbound static IPs of the Private Space network. 
	OutboundStaticIps []string `json:"outboundStaticIps,omitempty"`
	// The dns target of the Private Space network. 
	DnsTarget *string `json:"dnsTarget,omitempty"`
	// The internal dns target of the VPC that the Private Space network uses. 
	InternalDnsTarget *string `json:"internalDnsTarget,omitempty"`
}

// NewNetwork instantiates a new Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetwork() *Network {
	this := Network{}
	var cidrBlock string = "10.0.0.0/16"
	this.CidrBlock = &cidrBlock
	return &this
}

// NewNetworkWithDefaults instantiates a new Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkWithDefaults() *Network {
	this := Network{}
	var cidrBlock string = "10.0.0.0/16"
	this.CidrBlock = &cidrBlock
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Network) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Network) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Network) SetRegion(v string) {
	o.Region = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *Network) GetVpcId() string {
	if o == nil || IsNil(o.VpcId) {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetVpcIdOk() (*string, bool) {
	if o == nil || IsNil(o.VpcId) {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *Network) HasVpcId() bool {
	if o != nil && !IsNil(o.VpcId) {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *Network) SetVpcId(v string) {
	o.VpcId = &v
}

// GetCidrBlock returns the CidrBlock field value if set, zero value otherwise.
func (o *Network) GetCidrBlock() string {
	if o == nil || IsNil(o.CidrBlock) {
		var ret string
		return ret
	}
	return *o.CidrBlock
}

// GetCidrBlockOk returns a tuple with the CidrBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetCidrBlockOk() (*string, bool) {
	if o == nil || IsNil(o.CidrBlock) {
		return nil, false
	}
	return o.CidrBlock, true
}

// HasCidrBlock returns a boolean if a field has been set.
func (o *Network) HasCidrBlock() bool {
	if o != nil && !IsNil(o.CidrBlock) {
		return true
	}

	return false
}

// SetCidrBlock gets a reference to the given string and assigns it to the CidrBlock field.
func (o *Network) SetCidrBlock(v string) {
	o.CidrBlock = &v
}

// GetInternalDns returns the InternalDns field value if set, zero value otherwise.
func (o *Network) GetInternalDns() InternalDns {
	if o == nil || IsNil(o.InternalDns) {
		var ret InternalDns
		return ret
	}
	return *o.InternalDns
}

// GetInternalDnsOk returns a tuple with the InternalDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetInternalDnsOk() (*InternalDns, bool) {
	if o == nil || IsNil(o.InternalDns) {
		return nil, false
	}
	return o.InternalDns, true
}

// HasInternalDns returns a boolean if a field has been set.
func (o *Network) HasInternalDns() bool {
	if o != nil && !IsNil(o.InternalDns) {
		return true
	}

	return false
}

// SetInternalDns gets a reference to the given InternalDns and assigns it to the InternalDns field.
func (o *Network) SetInternalDns(v InternalDns) {
	o.InternalDns = &v
}

// GetInboundStaticIps returns the InboundStaticIps field value if set, zero value otherwise.
func (o *Network) GetInboundStaticIps() []string {
	if o == nil || IsNil(o.InboundStaticIps) {
		var ret []string
		return ret
	}
	return o.InboundStaticIps
}

// GetInboundStaticIpsOk returns a tuple with the InboundStaticIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetInboundStaticIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.InboundStaticIps) {
		return nil, false
	}
	return o.InboundStaticIps, true
}

// HasInboundStaticIps returns a boolean if a field has been set.
func (o *Network) HasInboundStaticIps() bool {
	if o != nil && !IsNil(o.InboundStaticIps) {
		return true
	}

	return false
}

// SetInboundStaticIps gets a reference to the given []string and assigns it to the InboundStaticIps field.
func (o *Network) SetInboundStaticIps(v []string) {
	o.InboundStaticIps = v
}

// GetOutboundStaticIps returns the OutboundStaticIps field value if set, zero value otherwise.
func (o *Network) GetOutboundStaticIps() []string {
	if o == nil || IsNil(o.OutboundStaticIps) {
		var ret []string
		return ret
	}
	return o.OutboundStaticIps
}

// GetOutboundStaticIpsOk returns a tuple with the OutboundStaticIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetOutboundStaticIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.OutboundStaticIps) {
		return nil, false
	}
	return o.OutboundStaticIps, true
}

// HasOutboundStaticIps returns a boolean if a field has been set.
func (o *Network) HasOutboundStaticIps() bool {
	if o != nil && !IsNil(o.OutboundStaticIps) {
		return true
	}

	return false
}

// SetOutboundStaticIps gets a reference to the given []string and assigns it to the OutboundStaticIps field.
func (o *Network) SetOutboundStaticIps(v []string) {
	o.OutboundStaticIps = v
}

// GetDnsTarget returns the DnsTarget field value if set, zero value otherwise.
func (o *Network) GetDnsTarget() string {
	if o == nil || IsNil(o.DnsTarget) {
		var ret string
		return ret
	}
	return *o.DnsTarget
}

// GetDnsTargetOk returns a tuple with the DnsTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetDnsTargetOk() (*string, bool) {
	if o == nil || IsNil(o.DnsTarget) {
		return nil, false
	}
	return o.DnsTarget, true
}

// HasDnsTarget returns a boolean if a field has been set.
func (o *Network) HasDnsTarget() bool {
	if o != nil && !IsNil(o.DnsTarget) {
		return true
	}

	return false
}

// SetDnsTarget gets a reference to the given string and assigns it to the DnsTarget field.
func (o *Network) SetDnsTarget(v string) {
	o.DnsTarget = &v
}

// GetInternalDnsTarget returns the InternalDnsTarget field value if set, zero value otherwise.
func (o *Network) GetInternalDnsTarget() string {
	if o == nil || IsNil(o.InternalDnsTarget) {
		var ret string
		return ret
	}
	return *o.InternalDnsTarget
}

// GetInternalDnsTargetOk returns a tuple with the InternalDnsTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetInternalDnsTargetOk() (*string, bool) {
	if o == nil || IsNil(o.InternalDnsTarget) {
		return nil, false
	}
	return o.InternalDnsTarget, true
}

// HasInternalDnsTarget returns a boolean if a field has been set.
func (o *Network) HasInternalDnsTarget() bool {
	if o != nil && !IsNil(o.InternalDnsTarget) {
		return true
	}

	return false
}

// SetInternalDnsTarget gets a reference to the given string and assigns it to the InternalDnsTarget field.
func (o *Network) SetInternalDnsTarget(v string) {
	o.InternalDnsTarget = &v
}

func (o Network) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Network) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.VpcId) {
		toSerialize["vpcId"] = o.VpcId
	}
	if !IsNil(o.CidrBlock) {
		toSerialize["cidrBlock"] = o.CidrBlock
	}
	if !IsNil(o.InternalDns) {
		toSerialize["internalDns"] = o.InternalDns
	}
	if !IsNil(o.InboundStaticIps) {
		toSerialize["inboundStaticIps"] = o.InboundStaticIps
	}
	if !IsNil(o.OutboundStaticIps) {
		toSerialize["outboundStaticIps"] = o.OutboundStaticIps
	}
	if !IsNil(o.DnsTarget) {
		toSerialize["dnsTarget"] = o.DnsTarget
	}
	if !IsNil(o.InternalDnsTarget) {
		toSerialize["internalDnsTarget"] = o.InternalDnsTarget
	}
	return toSerialize, nil
}

type NullableNetwork struct {
	value *Network
	isSet bool
}

func (v NullableNetwork) Get() *Network {
	return v.value
}

func (v *NullableNetwork) Set(val *Network) {
	v.value = val
	v.isSet = true
}

func (v NullableNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetwork(val *Network) *NullableNetwork {
	return &NullableNetwork{value: val, isSet: true}
}

func (v NullableNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


