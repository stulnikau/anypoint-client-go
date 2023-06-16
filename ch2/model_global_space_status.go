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

// checks if the GlobalSpaceStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalSpaceStatus{}

// GlobalSpaceStatus Global Status object for the space.
type GlobalSpaceStatus struct {
	Space Space `json:"space"`
	// The status of the space cluster component.
	Cluster []GlobalClusterStatus `json:"cluster"`
	Network Network8 `json:"network"`
}

// NewGlobalSpaceStatus instantiates a new GlobalSpaceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalSpaceStatus(space Space, cluster []GlobalClusterStatus, network Network8) *GlobalSpaceStatus {
	this := GlobalSpaceStatus{}
	this.Space = space
	this.Cluster = cluster
	this.Network = network
	return &this
}

// NewGlobalSpaceStatusWithDefaults instantiates a new GlobalSpaceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalSpaceStatusWithDefaults() *GlobalSpaceStatus {
	this := GlobalSpaceStatus{}
	return &this
}

// GetSpace returns the Space field value
func (o *GlobalSpaceStatus) GetSpace() Space {
	if o == nil {
		var ret Space
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *GlobalSpaceStatus) GetSpaceOk() (*Space, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *GlobalSpaceStatus) SetSpace(v Space) {
	o.Space = v
}

// GetCluster returns the Cluster field value
func (o *GlobalSpaceStatus) GetCluster() []GlobalClusterStatus {
	if o == nil {
		var ret []GlobalClusterStatus
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *GlobalSpaceStatus) GetClusterOk() ([]GlobalClusterStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster, true
}

// SetCluster sets field value
func (o *GlobalSpaceStatus) SetCluster(v []GlobalClusterStatus) {
	o.Cluster = v
}

// GetNetwork returns the Network field value
func (o *GlobalSpaceStatus) GetNetwork() Network8 {
	if o == nil {
		var ret Network8
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *GlobalSpaceStatus) GetNetworkOk() (*Network8, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *GlobalSpaceStatus) SetNetwork(v Network8) {
	o.Network = v
}

func (o GlobalSpaceStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalSpaceStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["space"] = o.Space
	toSerialize["cluster"] = o.Cluster
	toSerialize["network"] = o.Network
	return toSerialize, nil
}

type NullableGlobalSpaceStatus struct {
	value *GlobalSpaceStatus
	isSet bool
}

func (v NullableGlobalSpaceStatus) Get() *GlobalSpaceStatus {
	return v.value
}

func (v *NullableGlobalSpaceStatus) Set(val *GlobalSpaceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalSpaceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalSpaceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalSpaceStatus(val *GlobalSpaceStatus) *NullableGlobalSpaceStatus {
	return &NullableGlobalSpaceStatus{value: val, isSet: true}
}

func (v NullableGlobalSpaceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalSpaceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


