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

// checks if the IngressConfigurationLogs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngressConfigurationLogs{}

// IngressConfigurationLogs Ingress log level configuration
type IngressConfigurationLogs struct {
	Filters []PrivateSpaceIngressConfigurationLogsFilter `json:"filters,omitempty"`
	// Log levels
	PortLogLevel *string `json:"portLogLevel,omitempty"`
}

// NewIngressConfigurationLogs instantiates a new IngressConfigurationLogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngressConfigurationLogs() *IngressConfigurationLogs {
	this := IngressConfigurationLogs{}
	var portLogLevel string = "ERROR"
	this.PortLogLevel = &portLogLevel
	return &this
}

// NewIngressConfigurationLogsWithDefaults instantiates a new IngressConfigurationLogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngressConfigurationLogsWithDefaults() *IngressConfigurationLogs {
	this := IngressConfigurationLogs{}
	var portLogLevel string = "ERROR"
	this.PortLogLevel = &portLogLevel
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *IngressConfigurationLogs) GetFilters() []PrivateSpaceIngressConfigurationLogsFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []PrivateSpaceIngressConfigurationLogsFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngressConfigurationLogs) GetFiltersOk() ([]PrivateSpaceIngressConfigurationLogsFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *IngressConfigurationLogs) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []PrivateSpaceIngressConfigurationLogsFilter and assigns it to the Filters field.
func (o *IngressConfigurationLogs) SetFilters(v []PrivateSpaceIngressConfigurationLogsFilter) {
	o.Filters = v
}

// GetPortLogLevel returns the PortLogLevel field value if set, zero value otherwise.
func (o *IngressConfigurationLogs) GetPortLogLevel() string {
	if o == nil || IsNil(o.PortLogLevel) {
		var ret string
		return ret
	}
	return *o.PortLogLevel
}

// GetPortLogLevelOk returns a tuple with the PortLogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngressConfigurationLogs) GetPortLogLevelOk() (*string, bool) {
	if o == nil || IsNil(o.PortLogLevel) {
		return nil, false
	}
	return o.PortLogLevel, true
}

// HasPortLogLevel returns a boolean if a field has been set.
func (o *IngressConfigurationLogs) HasPortLogLevel() bool {
	if o != nil && !IsNil(o.PortLogLevel) {
		return true
	}

	return false
}

// SetPortLogLevel gets a reference to the given string and assigns it to the PortLogLevel field.
func (o *IngressConfigurationLogs) SetPortLogLevel(v string) {
	o.PortLogLevel = &v
}

func (o IngressConfigurationLogs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngressConfigurationLogs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.PortLogLevel) {
		toSerialize["portLogLevel"] = o.PortLogLevel
	}
	return toSerialize, nil
}

type NullableIngressConfigurationLogs struct {
	value *IngressConfigurationLogs
	isSet bool
}

func (v NullableIngressConfigurationLogs) Get() *IngressConfigurationLogs {
	return v.value
}

func (v *NullableIngressConfigurationLogs) Set(val *IngressConfigurationLogs) {
	v.value = val
	v.isSet = true
}

func (v NullableIngressConfigurationLogs) IsSet() bool {
	return v.isSet
}

func (v *NullableIngressConfigurationLogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngressConfigurationLogs(val *IngressConfigurationLogs) *NullableIngressConfigurationLogs {
	return &NullableIngressConfigurationLogs{value: val, isSet: true}
}

func (v NullableIngressConfigurationLogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngressConfigurationLogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


