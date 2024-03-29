/*
Syntropy network API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@syntropynet.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syntropy

import (
	"encoding/json"
)

// checks if the V1ConnectionUpdateChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ConnectionUpdateChange{}

// V1ConnectionUpdateChange struct for V1ConnectionUpdateChange
type V1ConnectionUpdateChange struct {
	ConnectionGroupId int32 `json:"connection_group_id"`
	IsSdnEnabled      bool  `json:"is_sdn_enabled"`
}

// NewV1ConnectionUpdateChange instantiates a new V1ConnectionUpdateChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ConnectionUpdateChange(connectionGroupId int32, isSdnEnabled bool) *V1ConnectionUpdateChange {
	this := V1ConnectionUpdateChange{}
	this.ConnectionGroupId = connectionGroupId
	this.IsSdnEnabled = isSdnEnabled
	return &this
}

// NewV1ConnectionUpdateChangeWithDefaults instantiates a new V1ConnectionUpdateChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ConnectionUpdateChangeWithDefaults() *V1ConnectionUpdateChange {
	this := V1ConnectionUpdateChange{}
	return &this
}

// GetConnectionGroupId returns the ConnectionGroupId field value
func (o *V1ConnectionUpdateChange) GetConnectionGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConnectionGroupId
}

// GetConnectionGroupIdOk returns a tuple with the ConnectionGroupId field value
// and a boolean to check if the value has been set.
func (o *V1ConnectionUpdateChange) GetConnectionGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionGroupId, true
}

// SetConnectionGroupId sets field value
func (o *V1ConnectionUpdateChange) SetConnectionGroupId(v int32) {
	o.ConnectionGroupId = v
}

// GetIsSdnEnabled returns the IsSdnEnabled field value
func (o *V1ConnectionUpdateChange) GetIsSdnEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSdnEnabled
}

// GetIsSdnEnabledOk returns a tuple with the IsSdnEnabled field value
// and a boolean to check if the value has been set.
func (o *V1ConnectionUpdateChange) GetIsSdnEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSdnEnabled, true
}

// SetIsSdnEnabled sets field value
func (o *V1ConnectionUpdateChange) SetIsSdnEnabled(v bool) {
	o.IsSdnEnabled = v
}

func (o V1ConnectionUpdateChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ConnectionUpdateChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connection_group_id"] = o.ConnectionGroupId
	toSerialize["is_sdn_enabled"] = o.IsSdnEnabled
	return toSerialize, nil
}

type NullableV1ConnectionUpdateChange struct {
	value *V1ConnectionUpdateChange
	isSet bool
}

func (v NullableV1ConnectionUpdateChange) Get() *V1ConnectionUpdateChange {
	return v.value
}

func (v *NullableV1ConnectionUpdateChange) Set(val *V1ConnectionUpdateChange) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ConnectionUpdateChange) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ConnectionUpdateChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ConnectionUpdateChange(val *V1ConnectionUpdateChange) *NullableV1ConnectionUpdateChange {
	return &NullableV1ConnectionUpdateChange{value: val, isSet: true}
}

func (v NullableV1ConnectionUpdateChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ConnectionUpdateChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
