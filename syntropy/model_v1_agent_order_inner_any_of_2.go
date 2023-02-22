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

// checks if the V1AgentOrderInnerAnyOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1AgentOrderInnerAnyOf2{}

// V1AgentOrderInnerAnyOf2 struct for V1AgentOrderInnerAnyOf2
type V1AgentOrderInnerAnyOf2 struct {
	AgentType Order `json:"agent_type"`
}

// NewV1AgentOrderInnerAnyOf2 instantiates a new V1AgentOrderInnerAnyOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AgentOrderInnerAnyOf2(agentType Order) *V1AgentOrderInnerAnyOf2 {
	this := V1AgentOrderInnerAnyOf2{}
	this.AgentType = agentType
	return &this
}

// NewV1AgentOrderInnerAnyOf2WithDefaults instantiates a new V1AgentOrderInnerAnyOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AgentOrderInnerAnyOf2WithDefaults() *V1AgentOrderInnerAnyOf2 {
	this := V1AgentOrderInnerAnyOf2{}
	return &this
}

// GetAgentType returns the AgentType field value
func (o *V1AgentOrderInnerAnyOf2) GetAgentType() Order {
	if o == nil {
		var ret Order
		return ret
	}

	return o.AgentType
}

// GetAgentTypeOk returns a tuple with the AgentType field value
// and a boolean to check if the value has been set.
func (o *V1AgentOrderInnerAnyOf2) GetAgentTypeOk() (*Order, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentType, true
}

// SetAgentType sets field value
func (o *V1AgentOrderInnerAnyOf2) SetAgentType(v Order) {
	o.AgentType = v
}

func (o V1AgentOrderInnerAnyOf2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1AgentOrderInnerAnyOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agent_type"] = o.AgentType
	return toSerialize, nil
}

type NullableV1AgentOrderInnerAnyOf2 struct {
	value *V1AgentOrderInnerAnyOf2
	isSet bool
}

func (v NullableV1AgentOrderInnerAnyOf2) Get() *V1AgentOrderInnerAnyOf2 {
	return v.value
}

func (v *NullableV1AgentOrderInnerAnyOf2) Set(val *V1AgentOrderInnerAnyOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AgentOrderInnerAnyOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AgentOrderInnerAnyOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AgentOrderInnerAnyOf2(val *V1AgentOrderInnerAnyOf2) *NullableV1AgentOrderInnerAnyOf2 {
	return &NullableV1AgentOrderInnerAnyOf2{value: val, isSet: true}
}

func (v NullableV1AgentOrderInnerAnyOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AgentOrderInnerAnyOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
