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

// V1NetworkConnectionsRemoveRequest struct for V1NetworkConnectionsRemoveRequest
type V1NetworkConnectionsRemoveRequest struct {
	AgentConnectionGroupIds []int32 `json:"agent_connection_group_ids"`
}

// NewV1NetworkConnectionsRemoveRequest instantiates a new V1NetworkConnectionsRemoveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NetworkConnectionsRemoveRequest(agentConnectionGroupIds []int32) *V1NetworkConnectionsRemoveRequest {
	this := V1NetworkConnectionsRemoveRequest{}
	this.AgentConnectionGroupIds = agentConnectionGroupIds
	return &this
}

// NewV1NetworkConnectionsRemoveRequestWithDefaults instantiates a new V1NetworkConnectionsRemoveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NetworkConnectionsRemoveRequestWithDefaults() *V1NetworkConnectionsRemoveRequest {
	this := V1NetworkConnectionsRemoveRequest{}
	return &this
}

// GetAgentConnectionGroupIds returns the AgentConnectionGroupIds field value
func (o *V1NetworkConnectionsRemoveRequest) GetAgentConnectionGroupIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.AgentConnectionGroupIds
}

// GetAgentConnectionGroupIdsOk returns a tuple with the AgentConnectionGroupIds field value
// and a boolean to check if the value has been set.
func (o *V1NetworkConnectionsRemoveRequest) GetAgentConnectionGroupIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentConnectionGroupIds, true
}

// SetAgentConnectionGroupIds sets field value
func (o *V1NetworkConnectionsRemoveRequest) SetAgentConnectionGroupIds(v []int32) {
	o.AgentConnectionGroupIds = v
}

func (o V1NetworkConnectionsRemoveRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["agent_connection_group_ids"] = o.AgentConnectionGroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableV1NetworkConnectionsRemoveRequest struct {
	value *V1NetworkConnectionsRemoveRequest
	isSet bool
}

func (v NullableV1NetworkConnectionsRemoveRequest) Get() *V1NetworkConnectionsRemoveRequest {
	return v.value
}

func (v *NullableV1NetworkConnectionsRemoveRequest) Set(val *V1NetworkConnectionsRemoveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NetworkConnectionsRemoveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NetworkConnectionsRemoveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NetworkConnectionsRemoveRequest(val *V1NetworkConnectionsRemoveRequest) *NullableV1NetworkConnectionsRemoveRequest {
	return &NullableV1NetworkConnectionsRemoveRequest{value: val, isSet: true}
}

func (v NullableV1NetworkConnectionsRemoveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NetworkConnectionsRemoveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
