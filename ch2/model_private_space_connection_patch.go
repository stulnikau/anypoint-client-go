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

// checks if the PrivateSpaceConnectionPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateSpaceConnectionPatch{}

// PrivateSpaceConnectionPatch struct for PrivateSpaceConnectionPatch
type PrivateSpaceConnectionPatch struct {
	Id string `json:"id"`
	Name string `json:"name"`
	// The CIDRs that the Cloudhub VPC will send traffic to via the VPN. 
	StaticRoutes []string `json:"staticRoutes,omitempty"`
}

// NewPrivateSpaceConnectionPatch instantiates a new PrivateSpaceConnectionPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateSpaceConnectionPatch(id string, name string) *PrivateSpaceConnectionPatch {
	this := PrivateSpaceConnectionPatch{}
	this.Id = id
	this.Name = name
	return &this
}

// NewPrivateSpaceConnectionPatchWithDefaults instantiates a new PrivateSpaceConnectionPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateSpaceConnectionPatchWithDefaults() *PrivateSpaceConnectionPatch {
	this := PrivateSpaceConnectionPatch{}
	return &this
}

// GetId returns the Id field value
func (o *PrivateSpaceConnectionPatch) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrivateSpaceConnectionPatch) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrivateSpaceConnectionPatch) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PrivateSpaceConnectionPatch) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrivateSpaceConnectionPatch) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrivateSpaceConnectionPatch) SetName(v string) {
	o.Name = v
}

// GetStaticRoutes returns the StaticRoutes field value if set, zero value otherwise.
func (o *PrivateSpaceConnectionPatch) GetStaticRoutes() []string {
	if o == nil || IsNil(o.StaticRoutes) {
		var ret []string
		return ret
	}
	return o.StaticRoutes
}

// GetStaticRoutesOk returns a tuple with the StaticRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSpaceConnectionPatch) GetStaticRoutesOk() ([]string, bool) {
	if o == nil || IsNil(o.StaticRoutes) {
		return nil, false
	}
	return o.StaticRoutes, true
}

// HasStaticRoutes returns a boolean if a field has been set.
func (o *PrivateSpaceConnectionPatch) HasStaticRoutes() bool {
	if o != nil && !IsNil(o.StaticRoutes) {
		return true
	}

	return false
}

// SetStaticRoutes gets a reference to the given []string and assigns it to the StaticRoutes field.
func (o *PrivateSpaceConnectionPatch) SetStaticRoutes(v []string) {
	o.StaticRoutes = v
}

func (o PrivateSpaceConnectionPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateSpaceConnectionPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.StaticRoutes) {
		toSerialize["staticRoutes"] = o.StaticRoutes
	}
	return toSerialize, nil
}

type NullablePrivateSpaceConnectionPatch struct {
	value *PrivateSpaceConnectionPatch
	isSet bool
}

func (v NullablePrivateSpaceConnectionPatch) Get() *PrivateSpaceConnectionPatch {
	return v.value
}

func (v *NullablePrivateSpaceConnectionPatch) Set(val *PrivateSpaceConnectionPatch) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateSpaceConnectionPatch) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateSpaceConnectionPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateSpaceConnectionPatch(val *PrivateSpaceConnectionPatch) *NullablePrivateSpaceConnectionPatch {
	return &NullablePrivateSpaceConnectionPatch{value: val, isSet: true}
}

func (v NullablePrivateSpaceConnectionPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateSpaceConnectionPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


