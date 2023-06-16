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

// checks if the IAMRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IAMRole{}

// IAMRole struct for IAMRole
type IAMRole struct {
	// IAM roles associated with the space
	Roles []string `json:"roles"`
	// The parent organization ID of the space
	OrganizationId string `json:"organizationId"`
	// id of the space.
	SpaceId string `json:"spaceId"`
	// IAM roles associated with the space
	Message *string `json:"message,omitempty"`
}

// NewIAMRole instantiates a new IAMRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIAMRole(roles []string, organizationId string, spaceId string) *IAMRole {
	this := IAMRole{}
	this.Roles = roles
	this.OrganizationId = organizationId
	this.SpaceId = spaceId
	return &this
}

// NewIAMRoleWithDefaults instantiates a new IAMRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIAMRoleWithDefaults() *IAMRole {
	this := IAMRole{}
	return &this
}

// GetRoles returns the Roles field value
func (o *IAMRole) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *IAMRole) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *IAMRole) SetRoles(v []string) {
	o.Roles = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *IAMRole) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *IAMRole) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *IAMRole) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetSpaceId returns the SpaceId field value
func (o *IAMRole) GetSpaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value
// and a boolean to check if the value has been set.
func (o *IAMRole) GetSpaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceId, true
}

// SetSpaceId sets field value
func (o *IAMRole) SetSpaceId(v string) {
	o.SpaceId = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IAMRole) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IAMRole) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IAMRole) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *IAMRole) SetMessage(v string) {
	o.Message = &v
}

func (o IAMRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IAMRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roles"] = o.Roles
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["spaceId"] = o.SpaceId
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableIAMRole struct {
	value *IAMRole
	isSet bool
}

func (v NullableIAMRole) Get() *IAMRole {
	return v.value
}

func (v *NullableIAMRole) Set(val *IAMRole) {
	v.value = val
	v.isSet = true
}

func (v NullableIAMRole) IsSet() bool {
	return v.isSet
}

func (v *NullableIAMRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIAMRole(val *IAMRole) *NullableIAMRole {
	return &NullableIAMRole{value: val, isSet: true}
}

func (v NullableIAMRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIAMRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


