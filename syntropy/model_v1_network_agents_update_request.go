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

// checks if the V1NetworkAgentsUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1NetworkAgentsUpdateRequest{}

// V1NetworkAgentsUpdateRequest struct for V1NetworkAgentsUpdateRequest
type V1NetworkAgentsUpdateRequest struct {
	AgentTags       []string                 `json:"agent_tags,omitempty"`
	AgentName       *string                  `json:"agent_name,omitempty"`
	AgentProviderId *int32                   `json:"agent_provider_id,omitempty"`
	Network         *AgentInterfacesMetadata `json:"network,omitempty"`
}

// NewV1NetworkAgentsUpdateRequest instantiates a new V1NetworkAgentsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NetworkAgentsUpdateRequest() *V1NetworkAgentsUpdateRequest {
	this := V1NetworkAgentsUpdateRequest{}
	return &this
}

// NewV1NetworkAgentsUpdateRequestWithDefaults instantiates a new V1NetworkAgentsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NetworkAgentsUpdateRequestWithDefaults() *V1NetworkAgentsUpdateRequest {
	this := V1NetworkAgentsUpdateRequest{}
	return &this
}

// GetAgentTags returns the AgentTags field value if set, zero value otherwise.
func (o *V1NetworkAgentsUpdateRequest) GetAgentTags() []string {
	if o == nil || isNil(o.AgentTags) {
		var ret []string
		return ret
	}
	return o.AgentTags
}

// GetAgentTagsOk returns a tuple with the AgentTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NetworkAgentsUpdateRequest) GetAgentTagsOk() ([]string, bool) {
	if o == nil || isNil(o.AgentTags) {
		return nil, false
	}
	return o.AgentTags, true
}

// HasAgentTags returns a boolean if a field has been set.
func (o *V1NetworkAgentsUpdateRequest) HasAgentTags() bool {
	if o != nil && !isNil(o.AgentTags) {
		return true
	}

	return false
}

// SetAgentTags gets a reference to the given []string and assigns it to the AgentTags field.
func (o *V1NetworkAgentsUpdateRequest) SetAgentTags(v []string) {
	o.AgentTags = v
}

// GetAgentName returns the AgentName field value if set, zero value otherwise.
func (o *V1NetworkAgentsUpdateRequest) GetAgentName() string {
	if o == nil || isNil(o.AgentName) {
		var ret string
		return ret
	}
	return *o.AgentName
}

// GetAgentNameOk returns a tuple with the AgentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NetworkAgentsUpdateRequest) GetAgentNameOk() (*string, bool) {
	if o == nil || isNil(o.AgentName) {
		return nil, false
	}
	return o.AgentName, true
}

// HasAgentName returns a boolean if a field has been set.
func (o *V1NetworkAgentsUpdateRequest) HasAgentName() bool {
	if o != nil && !isNil(o.AgentName) {
		return true
	}

	return false
}

// SetAgentName gets a reference to the given string and assigns it to the AgentName field.
func (o *V1NetworkAgentsUpdateRequest) SetAgentName(v string) {
	o.AgentName = &v
}

// GetAgentProviderId returns the AgentProviderId field value if set, zero value otherwise.
func (o *V1NetworkAgentsUpdateRequest) GetAgentProviderId() int32 {
	if o == nil || isNil(o.AgentProviderId) {
		var ret int32
		return ret
	}
	return *o.AgentProviderId
}

// GetAgentProviderIdOk returns a tuple with the AgentProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NetworkAgentsUpdateRequest) GetAgentProviderIdOk() (*int32, bool) {
	if o == nil || isNil(o.AgentProviderId) {
		return nil, false
	}
	return o.AgentProviderId, true
}

// HasAgentProviderId returns a boolean if a field has been set.
func (o *V1NetworkAgentsUpdateRequest) HasAgentProviderId() bool {
	if o != nil && !isNil(o.AgentProviderId) {
		return true
	}

	return false
}

// SetAgentProviderId gets a reference to the given int32 and assigns it to the AgentProviderId field.
func (o *V1NetworkAgentsUpdateRequest) SetAgentProviderId(v int32) {
	o.AgentProviderId = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *V1NetworkAgentsUpdateRequest) GetNetwork() AgentInterfacesMetadata {
	if o == nil || isNil(o.Network) {
		var ret AgentInterfacesMetadata
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1NetworkAgentsUpdateRequest) GetNetworkOk() (*AgentInterfacesMetadata, bool) {
	if o == nil || isNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *V1NetworkAgentsUpdateRequest) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given AgentInterfacesMetadata and assigns it to the Network field.
func (o *V1NetworkAgentsUpdateRequest) SetNetwork(v AgentInterfacesMetadata) {
	o.Network = &v
}

func (o V1NetworkAgentsUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1NetworkAgentsUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AgentTags) {
		toSerialize["agent_tags"] = o.AgentTags
	}
	if !isNil(o.AgentName) {
		toSerialize["agent_name"] = o.AgentName
	}
	if !isNil(o.AgentProviderId) {
		toSerialize["agent_provider_id"] = o.AgentProviderId
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	return toSerialize, nil
}

type NullableV1NetworkAgentsUpdateRequest struct {
	value *V1NetworkAgentsUpdateRequest
	isSet bool
}

func (v NullableV1NetworkAgentsUpdateRequest) Get() *V1NetworkAgentsUpdateRequest {
	return v.value
}

func (v *NullableV1NetworkAgentsUpdateRequest) Set(val *V1NetworkAgentsUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NetworkAgentsUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NetworkAgentsUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NetworkAgentsUpdateRequest(val *V1NetworkAgentsUpdateRequest) *NullableV1NetworkAgentsUpdateRequest {
	return &NullableV1NetworkAgentsUpdateRequest{value: val, isSet: true}
}

func (v NullableV1NetworkAgentsUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NetworkAgentsUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
