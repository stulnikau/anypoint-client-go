/*
 * Anypoint MQ Exchange Binding specfication
 *
 * Anypoint MQ Exchange Binding API specification
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ame_binding

import (
	"encoding/json"
)

// AMEBindingRuleBody struct for AMEBindingRuleBody
type AMEBindingRuleBody struct {
	RoutingRules *[]map[string]interface{} `json:"routingRules,omitempty"`
}

// NewAMEBindingRuleBody instantiates a new AMEBindingRuleBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAMEBindingRuleBody() *AMEBindingRuleBody {
	this := AMEBindingRuleBody{}
	return &this
}

// NewAMEBindingRuleBodyWithDefaults instantiates a new AMEBindingRuleBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAMEBindingRuleBodyWithDefaults() *AMEBindingRuleBody {
	this := AMEBindingRuleBody{}
	return &this
}

// GetRoutingRules returns the RoutingRules field value if set, zero value otherwise.
func (o *AMEBindingRuleBody) GetRoutingRules() []map[string]interface{} {
	if o == nil || o.RoutingRules == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.RoutingRules
}

// GetRoutingRulesOk returns a tuple with the RoutingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AMEBindingRuleBody) GetRoutingRulesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.RoutingRules == nil {
		return nil, false
	}
	return o.RoutingRules, true
}

// HasRoutingRules returns a boolean if a field has been set.
func (o *AMEBindingRuleBody) HasRoutingRules() bool {
	if o != nil && o.RoutingRules != nil {
		return true
	}

	return false
}

// SetRoutingRules gets a reference to the given []map[string]interface{} and assigns it to the RoutingRules field.
func (o *AMEBindingRuleBody) SetRoutingRules(v []map[string]interface{}) {
	o.RoutingRules = &v
}

func (o AMEBindingRuleBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoutingRules != nil {
		toSerialize["routingRules"] = o.RoutingRules
	}
	return json.Marshal(toSerialize)
}

type NullableAMEBindingRuleBody struct {
	value *AMEBindingRuleBody
	isSet bool
}

func (v NullableAMEBindingRuleBody) Get() *AMEBindingRuleBody {
	return v.value
}

func (v *NullableAMEBindingRuleBody) Set(val *AMEBindingRuleBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAMEBindingRuleBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAMEBindingRuleBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAMEBindingRuleBody(val *AMEBindingRuleBody) *NullableAMEBindingRuleBody {
	return &NullableAMEBindingRuleBody{value: val, isSet: true}
}

func (v NullableAMEBindingRuleBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAMEBindingRuleBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


