/*
Syntropy network API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: support@syntropynet.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// V1NetworkConnectionsCreateP2PRequestAgentPairs struct for V1NetworkConnectionsCreateP2PRequestAgentPairs
type V1NetworkConnectionsCreateP2PRequestAgentPairs struct {
	Agent2Id int32 `json:"agent_2_id"`
	Agent1Id int32 `json:"agent_1_id"`
}

// NewV1NetworkConnectionsCreateP2PRequestAgentPairs instantiates a new V1NetworkConnectionsCreateP2PRequestAgentPairs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NetworkConnectionsCreateP2PRequestAgentPairs(agent2Id int32, agent1Id int32) *V1NetworkConnectionsCreateP2PRequestAgentPairs {
	this := V1NetworkConnectionsCreateP2PRequestAgentPairs{}
	this.Agent2Id = agent2Id
	this.Agent1Id = agent1Id
	return &this
}

// NewV1NetworkConnectionsCreateP2PRequestAgentPairsWithDefaults instantiates a new V1NetworkConnectionsCreateP2PRequestAgentPairs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NetworkConnectionsCreateP2PRequestAgentPairsWithDefaults() *V1NetworkConnectionsCreateP2PRequestAgentPairs {
	this := V1NetworkConnectionsCreateP2PRequestAgentPairs{}
	return &this
}

// GetAgent2Id returns the Agent2Id field value
func (o *V1NetworkConnectionsCreateP2PRequestAgentPairs) GetAgent2Id() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Agent2Id
}

// GetAgent2IdOk returns a tuple with the Agent2Id field value
// and a boolean to check if the value has been set.
func (o *V1NetworkConnectionsCreateP2PRequestAgentPairs) GetAgent2IdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agent2Id, true
}

// SetAgent2Id sets field value
func (o *V1NetworkConnectionsCreateP2PRequestAgentPairs) SetAgent2Id(v int32) {
	o.Agent2Id = v
}

// GetAgent1Id returns the Agent1Id field value
func (o *V1NetworkConnectionsCreateP2PRequestAgentPairs) GetAgent1Id() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Agent1Id
}

// GetAgent1IdOk returns a tuple with the Agent1Id field value
// and a boolean to check if the value has been set.
func (o *V1NetworkConnectionsCreateP2PRequestAgentPairs) GetAgent1IdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agent1Id, true
}

// SetAgent1Id sets field value
func (o *V1NetworkConnectionsCreateP2PRequestAgentPairs) SetAgent1Id(v int32) {
	o.Agent1Id = v
}

func (o V1NetworkConnectionsCreateP2PRequestAgentPairs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["agent_2_id"] = o.Agent2Id
	}
	if true {
		toSerialize["agent_1_id"] = o.Agent1Id
	}
	return json.Marshal(toSerialize)
}

type NullableV1NetworkConnectionsCreateP2PRequestAgentPairs struct {
	value *V1NetworkConnectionsCreateP2PRequestAgentPairs
	isSet bool
}

func (v NullableV1NetworkConnectionsCreateP2PRequestAgentPairs) Get() *V1NetworkConnectionsCreateP2PRequestAgentPairs {
	return v.value
}

func (v *NullableV1NetworkConnectionsCreateP2PRequestAgentPairs) Set(val *V1NetworkConnectionsCreateP2PRequestAgentPairs) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NetworkConnectionsCreateP2PRequestAgentPairs) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NetworkConnectionsCreateP2PRequestAgentPairs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NetworkConnectionsCreateP2PRequestAgentPairs(val *V1NetworkConnectionsCreateP2PRequestAgentPairs) *NullableV1NetworkConnectionsCreateP2PRequestAgentPairs {
	return &NullableV1NetworkConnectionsCreateP2PRequestAgentPairs{value: val, isSet: true}
}

func (v NullableV1NetworkConnectionsCreateP2PRequestAgentPairs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NetworkConnectionsCreateP2PRequestAgentPairs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
