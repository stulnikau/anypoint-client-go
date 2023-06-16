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

// checks if the Environments4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Environments4{}

// Environments4 Environment association for the Private Space 
type Environments4 struct {
	// The type of environments. Possible values are all, sandbox, production
	Type string `json:"type"`
	// The Ids of business groups.
	BusinessGroups []string `json:"businessGroups"`
}

// NewEnvironments4 instantiates a new Environments4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironments4(type_ string, businessGroups []string) *Environments4 {
	this := Environments4{}
	this.Type = type_
	this.BusinessGroups = businessGroups
	return &this
}

// NewEnvironments4WithDefaults instantiates a new Environments4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironments4WithDefaults() *Environments4 {
	this := Environments4{}
	return &this
}

// GetType returns the Type field value
func (o *Environments4) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Environments4) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Environments4) SetType(v string) {
	o.Type = v
}

// GetBusinessGroups returns the BusinessGroups field value
func (o *Environments4) GetBusinessGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BusinessGroups
}

// GetBusinessGroupsOk returns a tuple with the BusinessGroups field value
// and a boolean to check if the value has been set.
func (o *Environments4) GetBusinessGroupsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessGroups, true
}

// SetBusinessGroups sets field value
func (o *Environments4) SetBusinessGroups(v []string) {
	o.BusinessGroups = v
}

func (o Environments4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Environments4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["businessGroups"] = o.BusinessGroups
	return toSerialize, nil
}

type NullableEnvironments4 struct {
	value *Environments4
	isSet bool
}

func (v NullableEnvironments4) Get() *Environments4 {
	return v.value
}

func (v *NullableEnvironments4) Set(val *Environments4) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironments4) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironments4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironments4(val *Environments4) *NullableEnvironments4 {
	return &NullableEnvironments4{value: val, isSet: true}
}

func (v NullableEnvironments4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironments4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


