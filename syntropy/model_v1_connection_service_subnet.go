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

// checks if the V1ConnectionServiceSubnet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ConnectionServiceSubnet{}

// V1ConnectionServiceSubnet struct for V1ConnectionServiceSubnet
type V1ConnectionServiceSubnet struct {
	AgentServiceSubnetId           int32                         `json:"agent_service_subnet_id"`
	AgentConnectionSubnetId        int32                         `json:"agent_connection_subnet_id"`
	AgentConnectionSubnetIsEnabled bool                          `json:"agent_connection_subnet_is_enabled"`
	AgentConnectionSubnetError     NullableString                `json:"agent_connection_subnet_error"`
	AgentConnectionSubnetStatus    AgentConnectionSubnetStatuses `json:"agent_connection_subnet_status"`
}

// NewV1ConnectionServiceSubnet instantiates a new V1ConnectionServiceSubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ConnectionServiceSubnet(agentServiceSubnetId int32, agentConnectionSubnetId int32, agentConnectionSubnetIsEnabled bool, agentConnectionSubnetError NullableString, agentConnectionSubnetStatus AgentConnectionSubnetStatuses) *V1ConnectionServiceSubnet {
	this := V1ConnectionServiceSubnet{}
	this.AgentServiceSubnetId = agentServiceSubnetId
	this.AgentConnectionSubnetId = agentConnectionSubnetId
	this.AgentConnectionSubnetIsEnabled = agentConnectionSubnetIsEnabled
	this.AgentConnectionSubnetError = agentConnectionSubnetError
	this.AgentConnectionSubnetStatus = agentConnectionSubnetStatus
	return &this
}

// NewV1ConnectionServiceSubnetWithDefaults instantiates a new V1ConnectionServiceSubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ConnectionServiceSubnetWithDefaults() *V1ConnectionServiceSubnet {
	this := V1ConnectionServiceSubnet{}
	return &this
}

// GetAgentServiceSubnetId returns the AgentServiceSubnetId field value
func (o *V1ConnectionServiceSubnet) GetAgentServiceSubnetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AgentServiceSubnetId
}

// GetAgentServiceSubnetIdOk returns a tuple with the AgentServiceSubnetId field value
// and a boolean to check if the value has been set.
func (o *V1ConnectionServiceSubnet) GetAgentServiceSubnetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentServiceSubnetId, true
}

// SetAgentServiceSubnetId sets field value
func (o *V1ConnectionServiceSubnet) SetAgentServiceSubnetId(v int32) {
	o.AgentServiceSubnetId = v
}

// GetAgentConnectionSubnetId returns the AgentConnectionSubnetId field value
func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AgentConnectionSubnetId
}

// GetAgentConnectionSubnetIdOk returns a tuple with the AgentConnectionSubnetId field value
// and a boolean to check if the value has been set.
func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentConnectionSubnetId, true
}

// SetAgentConnectionSubnetId sets field value
func (o *V1ConnectionServiceSubnet) SetAgentConnectionSubnetId(v int32) {
	o.AgentConnectionSubnetId = v
}

// GetAgentConnectionSubnetIsEnabled returns the AgentConnectionSubnetIsEnabled field value
func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AgentConnectionSubnetIsEnabled
}

// GetAgentConnectionSubnetIsEnabledOk returns a tuple with the AgentConnectionSubnetIsEnabled field value
// and a boolean to check if the value has been set.
func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentConnectionSubnetIsEnabled, true
}

// SetAgentConnectionSubnetIsEnabled sets field value
func (o *V1ConnectionServiceSubnet) SetAgentConnectionSubnetIsEnabled(v bool) {
	o.AgentConnectionSubnetIsEnabled = v
}

// GetAgentConnectionSubnetError returns the AgentConnectionSubnetError field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetError() string {
	if o == nil || o.AgentConnectionSubnetError.Get() == nil {
		var ret string
		return ret
	}

	return *o.AgentConnectionSubnetError.Get()
}

// GetAgentConnectionSubnetErrorOk returns a tuple with the AgentConnectionSubnetError field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentConnectionSubnetError.Get(), o.AgentConnectionSubnetError.IsSet()
}

// SetAgentConnectionSubnetError sets field value
func (o *V1ConnectionServiceSubnet) SetAgentConnectionSubnetError(v string) {
	o.AgentConnectionSubnetError.Set(&v)
}

// GetAgentConnectionSubnetStatus returns the AgentConnectionSubnetStatus field value
func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetStatus() AgentConnectionSubnetStatuses {
	if o == nil {
		var ret AgentConnectionSubnetStatuses
		return ret
	}

	return o.AgentConnectionSubnetStatus
}

// GetAgentConnectionSubnetStatusOk returns a tuple with the AgentConnectionSubnetStatus field value
// and a boolean to check if the value has been set.
func (o *V1ConnectionServiceSubnet) GetAgentConnectionSubnetStatusOk() (*AgentConnectionSubnetStatuses, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentConnectionSubnetStatus, true
}

// SetAgentConnectionSubnetStatus sets field value
func (o *V1ConnectionServiceSubnet) SetAgentConnectionSubnetStatus(v AgentConnectionSubnetStatuses) {
	o.AgentConnectionSubnetStatus = v
}

func (o V1ConnectionServiceSubnet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ConnectionServiceSubnet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agent_service_subnet_id"] = o.AgentServiceSubnetId
	toSerialize["agent_connection_subnet_id"] = o.AgentConnectionSubnetId
	toSerialize["agent_connection_subnet_is_enabled"] = o.AgentConnectionSubnetIsEnabled
	toSerialize["agent_connection_subnet_error"] = o.AgentConnectionSubnetError.Get()
	toSerialize["agent_connection_subnet_status"] = o.AgentConnectionSubnetStatus
	return toSerialize, nil
}

type NullableV1ConnectionServiceSubnet struct {
	value *V1ConnectionServiceSubnet
	isSet bool
}

func (v NullableV1ConnectionServiceSubnet) Get() *V1ConnectionServiceSubnet {
	return v.value
}

func (v *NullableV1ConnectionServiceSubnet) Set(val *V1ConnectionServiceSubnet) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ConnectionServiceSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ConnectionServiceSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ConnectionServiceSubnet(val *V1ConnectionServiceSubnet) *NullableV1ConnectionServiceSubnet {
	return &NullableV1ConnectionServiceSubnet{value: val, isSet: true}
}

func (v NullableV1ConnectionServiceSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ConnectionServiceSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
