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

// checks if the Fabric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Fabric{}

// Fabric The last known status of the space fabric component
type Fabric struct {
	// The last known status
	Status *string `json:"status,omitempty"`
	// Message describing the last known status
	Message *string `json:"message,omitempty"`
	// Last known timestamp of the status in EPOCH
	LastSeenTimeStamp *float32 `json:"lastSeenTimeStamp,omitempty"`
}

// NewFabric instantiates a new Fabric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabric() *Fabric {
	this := Fabric{}
	return &this
}

// NewFabricWithDefaults instantiates a new Fabric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricWithDefaults() *Fabric {
	this := Fabric{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Fabric) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fabric) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Fabric) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Fabric) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Fabric) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fabric) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Fabric) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Fabric) SetMessage(v string) {
	o.Message = &v
}

// GetLastSeenTimeStamp returns the LastSeenTimeStamp field value if set, zero value otherwise.
func (o *Fabric) GetLastSeenTimeStamp() float32 {
	if o == nil || IsNil(o.LastSeenTimeStamp) {
		var ret float32
		return ret
	}
	return *o.LastSeenTimeStamp
}

// GetLastSeenTimeStampOk returns a tuple with the LastSeenTimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fabric) GetLastSeenTimeStampOk() (*float32, bool) {
	if o == nil || IsNil(o.LastSeenTimeStamp) {
		return nil, false
	}
	return o.LastSeenTimeStamp, true
}

// HasLastSeenTimeStamp returns a boolean if a field has been set.
func (o *Fabric) HasLastSeenTimeStamp() bool {
	if o != nil && !IsNil(o.LastSeenTimeStamp) {
		return true
	}

	return false
}

// SetLastSeenTimeStamp gets a reference to the given float32 and assigns it to the LastSeenTimeStamp field.
func (o *Fabric) SetLastSeenTimeStamp(v float32) {
	o.LastSeenTimeStamp = &v
}

func (o Fabric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Fabric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.LastSeenTimeStamp) {
		toSerialize["lastSeenTimeStamp"] = o.LastSeenTimeStamp
	}
	return toSerialize, nil
}

type NullableFabric struct {
	value *Fabric
	isSet bool
}

func (v NullableFabric) Get() *Fabric {
	return v.value
}

func (v *NullableFabric) Set(val *Fabric) {
	v.value = val
	v.isSet = true
}

func (v NullableFabric) IsSet() bool {
	return v.isSet
}

func (v *NullableFabric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabric(val *Fabric) *NullableFabric {
	return &NullableFabric{value: val, isSet: true}
}

func (v NullableFabric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


