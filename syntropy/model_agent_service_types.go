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
	"fmt"
)

// AgentServiceTypes the model 'AgentServiceTypes'
type AgentServiceTypes string

// List of AgentServiceTypes
const (
	AGENTSERVICETYPES_DOCKER     AgentServiceTypes = "DOCKER"
	AGENTSERVICETYPES_HOST       AgentServiceTypes = "HOST"
	AGENTSERVICETYPES_KUBERNETES AgentServiceTypes = "KUBERNETES"
	AGENTSERVICETYPES_MANUAL     AgentServiceTypes = "MANUAL"
)

// All allowed values of AgentServiceTypes enum
var AllowedAgentServiceTypesEnumValues = []AgentServiceTypes{
	"DOCKER",
	"HOST",
	"KUBERNETES",
	"MANUAL",
}

func (v *AgentServiceTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AgentServiceTypes(value)
	for _, existing := range AllowedAgentServiceTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AgentServiceTypes", value)
}

// NewAgentServiceTypesFromValue returns a pointer to a valid AgentServiceTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentServiceTypesFromValue(v string) (*AgentServiceTypes, error) {
	ev := AgentServiceTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentServiceTypes: valid values are %v", v, AllowedAgentServiceTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentServiceTypes) IsValid() bool {
	for _, existing := range AllowedAgentServiceTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentServiceTypes value
func (v AgentServiceTypes) Ptr() *AgentServiceTypes {
	return &v
}

type NullableAgentServiceTypes struct {
	value *AgentServiceTypes
	isSet bool
}

func (v NullableAgentServiceTypes) Get() *AgentServiceTypes {
	return v.value
}

func (v *NullableAgentServiceTypes) Set(val *AgentServiceTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentServiceTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentServiceTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentServiceTypes(val *AgentServiceTypes) *NullableAgentServiceTypes {
	return &NullableAgentServiceTypes{value: val, isSet: true}
}

func (v NullableAgentServiceTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentServiceTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
