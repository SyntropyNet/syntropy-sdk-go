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

// V1NetworkAgentsCreateRequest struct for V1NetworkAgentsCreateRequest
type V1NetworkAgentsCreateRequest struct {
	AgentTags       []string `json:"agent_tags,omitempty"`
	AgentToken      string   `json:"agent_token"`
	AgentName       string   `json:"agent_name"`
	AgentProviderId *int32   `json:"agent_provider_id,omitempty"`
}

// NewV1NetworkAgentsCreateRequest instantiates a new V1NetworkAgentsCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NetworkAgentsCreateRequest(agentToken string, agentName string) *V1NetworkAgentsCreateRequest {
	this := V1NetworkAgentsCreateRequest{}
	this.AgentToken = agentToken
	this.AgentName = agentName
	return &this
}

// NewV1NetworkAgentsCreateRequestWithDefaults instantiates a new V1NetworkAgentsCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NetworkAgentsCreateRequestWithDefaults() *V1NetworkAgentsCreateRequest {
	this := V1NetworkAgentsCreateRequest{}
	return &this
}

// GetAgentTags returns the AgentTags field value if set, zero value otherwise.
func (o *V1NetworkAgentsCreateRequest) GetAgentTags() []string {
	if o == nil || o.AgentTags == nil {
		var ret []string
		return ret
	}
	return o.AgentTags
}

// GetAgentTagsOk returns a tuple with the AgentTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NetworkAgentsCreateRequest) GetAgentTagsOk() ([]string, bool) {
	if o == nil || o.AgentTags == nil {
		return nil, false
	}
	return o.AgentTags, true
}

// HasAgentTags returns a boolean if a field has been set.
func (o *V1NetworkAgentsCreateRequest) HasAgentTags() bool {
	if o != nil && o.AgentTags != nil {
		return true
	}

	return false
}

// SetAgentTags gets a reference to the given []string and assigns it to the AgentTags field.
func (o *V1NetworkAgentsCreateRequest) SetAgentTags(v []string) {
	o.AgentTags = v
}

// GetAgentToken returns the AgentToken field value
func (o *V1NetworkAgentsCreateRequest) GetAgentToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentToken
}

// GetAgentTokenOk returns a tuple with the AgentToken field value
// and a boolean to check if the value has been set.
func (o *V1NetworkAgentsCreateRequest) GetAgentTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentToken, true
}

// SetAgentToken sets field value
func (o *V1NetworkAgentsCreateRequest) SetAgentToken(v string) {
	o.AgentToken = v
}

// GetAgentName returns the AgentName field value
func (o *V1NetworkAgentsCreateRequest) GetAgentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentName
}

// GetAgentNameOk returns a tuple with the AgentName field value
// and a boolean to check if the value has been set.
func (o *V1NetworkAgentsCreateRequest) GetAgentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentName, true
}

// SetAgentName sets field value
func (o *V1NetworkAgentsCreateRequest) SetAgentName(v string) {
	o.AgentName = v
}

// GetAgentProviderId returns the AgentProviderId field value if set, zero value otherwise.
func (o *V1NetworkAgentsCreateRequest) GetAgentProviderId() int32 {
	if o == nil || o.AgentProviderId == nil {
		var ret int32
		return ret
	}
	return *o.AgentProviderId
}

// GetAgentProviderIdOk returns a tuple with the AgentProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NetworkAgentsCreateRequest) GetAgentProviderIdOk() (*int32, bool) {
	if o == nil || o.AgentProviderId == nil {
		return nil, false
	}
	return o.AgentProviderId, true
}

// HasAgentProviderId returns a boolean if a field has been set.
func (o *V1NetworkAgentsCreateRequest) HasAgentProviderId() bool {
	if o != nil && o.AgentProviderId != nil {
		return true
	}

	return false
}

// SetAgentProviderId gets a reference to the given int32 and assigns it to the AgentProviderId field.
func (o *V1NetworkAgentsCreateRequest) SetAgentProviderId(v int32) {
	o.AgentProviderId = &v
}

func (o V1NetworkAgentsCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AgentTags != nil {
		toSerialize["agent_tags"] = o.AgentTags
	}
	if true {
		toSerialize["agent_token"] = o.AgentToken
	}
	if true {
		toSerialize["agent_name"] = o.AgentName
	}
	if o.AgentProviderId != nil {
		toSerialize["agent_provider_id"] = o.AgentProviderId
	}
	return json.Marshal(toSerialize)
}

type NullableV1NetworkAgentsCreateRequest struct {
	value *V1NetworkAgentsCreateRequest
	isSet bool
}

func (v NullableV1NetworkAgentsCreateRequest) Get() *V1NetworkAgentsCreateRequest {
	return v.value
}

func (v *NullableV1NetworkAgentsCreateRequest) Set(val *V1NetworkAgentsCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NetworkAgentsCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NetworkAgentsCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NetworkAgentsCreateRequest(val *V1NetworkAgentsCreateRequest) *NullableV1NetworkAgentsCreateRequest {
	return &NullableV1NetworkAgentsCreateRequest{value: val, isSet: true}
}

func (v NullableV1NetworkAgentsCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NetworkAgentsCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}