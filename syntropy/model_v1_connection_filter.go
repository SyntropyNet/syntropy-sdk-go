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

// checks if the V1ConnectionFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ConnectionFilter{}

// V1ConnectionFilter struct for V1ConnectionFilter
type V1ConnectionFilter struct {
	AgentConnectionGroupId []int32             `json:"agent_connection_group_id,omitempty"`
	AgentId                []int32             `json:"agent_id,omitempty"`
	AgentPair              []V1AgentPairFilter `json:"agent_pair,omitempty"`
}

// NewV1ConnectionFilter instantiates a new V1ConnectionFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ConnectionFilter() *V1ConnectionFilter {
	this := V1ConnectionFilter{}
	return &this
}

// NewV1ConnectionFilterWithDefaults instantiates a new V1ConnectionFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ConnectionFilterWithDefaults() *V1ConnectionFilter {
	this := V1ConnectionFilter{}
	return &this
}

// GetAgentConnectionGroupId returns the AgentConnectionGroupId field value if set, zero value otherwise.
func (o *V1ConnectionFilter) GetAgentConnectionGroupId() []int32 {
	if o == nil || isNil(o.AgentConnectionGroupId) {
		var ret []int32
		return ret
	}
	return o.AgentConnectionGroupId
}

// GetAgentConnectionGroupIdOk returns a tuple with the AgentConnectionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConnectionFilter) GetAgentConnectionGroupIdOk() ([]int32, bool) {
	if o == nil || isNil(o.AgentConnectionGroupId) {
		return nil, false
	}
	return o.AgentConnectionGroupId, true
}

// HasAgentConnectionGroupId returns a boolean if a field has been set.
func (o *V1ConnectionFilter) HasAgentConnectionGroupId() bool {
	if o != nil && !isNil(o.AgentConnectionGroupId) {
		return true
	}

	return false
}

// SetAgentConnectionGroupId gets a reference to the given []int32 and assigns it to the AgentConnectionGroupId field.
func (o *V1ConnectionFilter) SetAgentConnectionGroupId(v []int32) {
	o.AgentConnectionGroupId = v
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *V1ConnectionFilter) GetAgentId() []int32 {
	if o == nil || isNil(o.AgentId) {
		var ret []int32
		return ret
	}
	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConnectionFilter) GetAgentIdOk() ([]int32, bool) {
	if o == nil || isNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *V1ConnectionFilter) HasAgentId() bool {
	if o != nil && !isNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given []int32 and assigns it to the AgentId field.
func (o *V1ConnectionFilter) SetAgentId(v []int32) {
	o.AgentId = v
}

// GetAgentPair returns the AgentPair field value if set, zero value otherwise.
func (o *V1ConnectionFilter) GetAgentPair() []V1AgentPairFilter {
	if o == nil || isNil(o.AgentPair) {
		var ret []V1AgentPairFilter
		return ret
	}
	return o.AgentPair
}

// GetAgentPairOk returns a tuple with the AgentPair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConnectionFilter) GetAgentPairOk() ([]V1AgentPairFilter, bool) {
	if o == nil || isNil(o.AgentPair) {
		return nil, false
	}
	return o.AgentPair, true
}

// HasAgentPair returns a boolean if a field has been set.
func (o *V1ConnectionFilter) HasAgentPair() bool {
	if o != nil && !isNil(o.AgentPair) {
		return true
	}

	return false
}

// SetAgentPair gets a reference to the given []V1AgentPairFilter and assigns it to the AgentPair field.
func (o *V1ConnectionFilter) SetAgentPair(v []V1AgentPairFilter) {
	o.AgentPair = v
}

func (o V1ConnectionFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ConnectionFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AgentConnectionGroupId) {
		toSerialize["agent_connection_group_id"] = o.AgentConnectionGroupId
	}
	if !isNil(o.AgentId) {
		toSerialize["agent_id"] = o.AgentId
	}
	if !isNil(o.AgentPair) {
		toSerialize["agent_pair"] = o.AgentPair
	}
	return toSerialize, nil
}

type NullableV1ConnectionFilter struct {
	value *V1ConnectionFilter
	isSet bool
}

func (v NullableV1ConnectionFilter) Get() *V1ConnectionFilter {
	return v.value
}

func (v *NullableV1ConnectionFilter) Set(val *V1ConnectionFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ConnectionFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ConnectionFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ConnectionFilter(val *V1ConnectionFilter) *NullableV1ConnectionFilter {
	return &NullableV1ConnectionFilter{value: val, isSet: true}
}

func (v NullableV1ConnectionFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ConnectionFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
