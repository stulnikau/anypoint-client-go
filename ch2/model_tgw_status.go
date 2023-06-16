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

// checks if the TgwStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TgwStatus{}

// TgwStatus struct for TgwStatus
type TgwStatus struct {
	Gateway string `json:"gateway"`
	Attachment string `json:"attachment"`
	TgwResource string `json:"tgwResource"`
	Routes []string `json:"routes"`
}

// NewTgwStatus instantiates a new TgwStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTgwStatus(gateway string, attachment string, tgwResource string, routes []string) *TgwStatus {
	this := TgwStatus{}
	this.Gateway = gateway
	this.Attachment = attachment
	this.TgwResource = tgwResource
	this.Routes = routes
	return &this
}

// NewTgwStatusWithDefaults instantiates a new TgwStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTgwStatusWithDefaults() *TgwStatus {
	this := TgwStatus{}
	return &this
}

// GetGateway returns the Gateway field value
func (o *TgwStatus) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *TgwStatus) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *TgwStatus) SetGateway(v string) {
	o.Gateway = v
}

// GetAttachment returns the Attachment field value
func (o *TgwStatus) GetAttachment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value
// and a boolean to check if the value has been set.
func (o *TgwStatus) GetAttachmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attachment, true
}

// SetAttachment sets field value
func (o *TgwStatus) SetAttachment(v string) {
	o.Attachment = v
}

// GetTgwResource returns the TgwResource field value
func (o *TgwStatus) GetTgwResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TgwResource
}

// GetTgwResourceOk returns a tuple with the TgwResource field value
// and a boolean to check if the value has been set.
func (o *TgwStatus) GetTgwResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TgwResource, true
}

// SetTgwResource sets field value
func (o *TgwStatus) SetTgwResource(v string) {
	o.TgwResource = v
}

// GetRoutes returns the Routes field value
func (o *TgwStatus) GetRoutes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value
// and a boolean to check if the value has been set.
func (o *TgwStatus) GetRoutesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Routes, true
}

// SetRoutes sets field value
func (o *TgwStatus) SetRoutes(v []string) {
	o.Routes = v
}

func (o TgwStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TgwStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gateway"] = o.Gateway
	toSerialize["attachment"] = o.Attachment
	toSerialize["tgwResource"] = o.TgwResource
	toSerialize["routes"] = o.Routes
	return toSerialize, nil
}

type NullableTgwStatus struct {
	value *TgwStatus
	isSet bool
}

func (v NullableTgwStatus) Get() *TgwStatus {
	return v.value
}

func (v *NullableTgwStatus) Set(val *TgwStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTgwStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTgwStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTgwStatus(val *TgwStatus) *NullableTgwStatus {
	return &NullableTgwStatus{value: val, isSet: true}
}

func (v NullableTgwStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTgwStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


