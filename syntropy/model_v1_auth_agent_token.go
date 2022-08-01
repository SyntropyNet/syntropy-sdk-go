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
	"time"
)

// V1AuthAgentToken struct for V1AuthAgentToken
type V1AuthAgentToken struct {
	AgentTokenName        *string        `json:"agent_token_name,omitempty"`
	AgentTokenDescription NullableString `json:"agent_token_description,omitempty"`
	AgentTokenValidUntil  *time.Time     `json:"agent_token_valid_until,omitempty"`
	AgentTokenStatus      *bool          `json:"agent_token_status,omitempty"`
	AgentTokenId          *int32         `json:"agent_token_id,omitempty"`
	AgentTokenCreatedAt   *time.Time     `json:"agent_token_created_at,omitempty"`
	AgentTokenUpdatedAt   *time.Time     `json:"agent_token_updated_at,omitempty"`
}

// NewV1AuthAgentToken instantiates a new V1AuthAgentToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AuthAgentToken() *V1AuthAgentToken {
	this := V1AuthAgentToken{}
	return &this
}

// NewV1AuthAgentTokenWithDefaults instantiates a new V1AuthAgentToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AuthAgentTokenWithDefaults() *V1AuthAgentToken {
	this := V1AuthAgentToken{}
	return &this
}

// GetAgentTokenName returns the AgentTokenName field value if set, zero value otherwise.
func (o *V1AuthAgentToken) GetAgentTokenName() string {
	if o == nil || o.AgentTokenName == nil {
		var ret string
		return ret
	}
	return *o.AgentTokenName
}

// GetAgentTokenNameOk returns a tuple with the AgentTokenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AuthAgentToken) GetAgentTokenNameOk() (*string, bool) {
	if o == nil || o.AgentTokenName == nil {
		return nil, false
	}
	return o.AgentTokenName, true
}

// HasAgentTokenName returns a boolean if a field has been set.
func (o *V1AuthAgentToken) HasAgentTokenName() bool {
	if o != nil && o.AgentTokenName != nil {
		return true
	}

	return false
}

// SetAgentTokenName gets a reference to the given string and assigns it to the AgentTokenName field.
func (o *V1AuthAgentToken) SetAgentTokenName(v string) {
	o.AgentTokenName = &v
}

// GetAgentTokenDescription returns the AgentTokenDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *V1AuthAgentToken) GetAgentTokenDescription() string {
	if o == nil || o.AgentTokenDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.AgentTokenDescription.Get()
}

// GetAgentTokenDescriptionOk returns a tuple with the AgentTokenDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V1AuthAgentToken) GetAgentTokenDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentTokenDescription.Get(), o.AgentTokenDescription.IsSet()
}

// HasAgentTokenDescription returns a boolean if a field has been set.
func (o *V1AuthAgentToken) HasAgentTokenDescription() bool {
	if o != nil && o.AgentTokenDescription.IsSet() {
		return true
	}

	return false
}

// SetAgentTokenDescription gets a reference to the given NullableString and assigns it to the AgentTokenDescription field.
func (o *V1AuthAgentToken) SetAgentTokenDescription(v string) {
	o.AgentTokenDescription.Set(&v)
}

// SetAgentTokenDescriptionNil sets the value for AgentTokenDescription to be an explicit nil
func (o *V1AuthAgentToken) SetAgentTokenDescriptionNil() {
	o.AgentTokenDescription.Set(nil)
}

// UnsetAgentTokenDescription ensures that no value is present for AgentTokenDescription, not even an explicit nil
func (o *V1AuthAgentToken) UnsetAgentTokenDescription() {
	o.AgentTokenDescription.Unset()
}

// GetAgentTokenValidUntil returns the AgentTokenValidUntil field value if set, zero value otherwise.
func (o *V1AuthAgentToken) GetAgentTokenValidUntil() time.Time {
	if o == nil || o.AgentTokenValidUntil == nil {
		var ret time.Time
		return ret
	}
	return *o.AgentTokenValidUntil
}

// GetAgentTokenValidUntilOk returns a tuple with the AgentTokenValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AuthAgentToken) GetAgentTokenValidUntilOk() (*time.Time, bool) {
	if o == nil || o.AgentTokenValidUntil == nil {
		return nil, false
	}
	return o.AgentTokenValidUntil, true
}

// HasAgentTokenValidUntil returns a boolean if a field has been set.
func (o *V1AuthAgentToken) HasAgentTokenValidUntil() bool {
	if o != nil && o.AgentTokenValidUntil != nil {
		return true
	}

	return false
}

// SetAgentTokenValidUntil gets a reference to the given time.Time and assigns it to the AgentTokenValidUntil field.
func (o *V1AuthAgentToken) SetAgentTokenValidUntil(v time.Time) {
	o.AgentTokenValidUntil = &v
}

// GetAgentTokenStatus returns the AgentTokenStatus field value if set, zero value otherwise.
func (o *V1AuthAgentToken) GetAgentTokenStatus() bool {
	if o == nil || o.AgentTokenStatus == nil {
		var ret bool
		return ret
	}
	return *o.AgentTokenStatus
}

// GetAgentTokenStatusOk returns a tuple with the AgentTokenStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AuthAgentToken) GetAgentTokenStatusOk() (*bool, bool) {
	if o == nil || o.AgentTokenStatus == nil {
		return nil, false
	}
	return o.AgentTokenStatus, true
}

// HasAgentTokenStatus returns a boolean if a field has been set.
func (o *V1AuthAgentToken) HasAgentTokenStatus() bool {
	if o != nil && o.AgentTokenStatus != nil {
		return true
	}

	return false
}

// SetAgentTokenStatus gets a reference to the given bool and assigns it to the AgentTokenStatus field.
func (o *V1AuthAgentToken) SetAgentTokenStatus(v bool) {
	o.AgentTokenStatus = &v
}

// GetAgentTokenId returns the AgentTokenId field value if set, zero value otherwise.
func (o *V1AuthAgentToken) GetAgentTokenId() int32 {
	if o == nil || o.AgentTokenId == nil {
		var ret int32
		return ret
	}
	return *o.AgentTokenId
}

// GetAgentTokenIdOk returns a tuple with the AgentTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AuthAgentToken) GetAgentTokenIdOk() (*int32, bool) {
	if o == nil || o.AgentTokenId == nil {
		return nil, false
	}
	return o.AgentTokenId, true
}

// HasAgentTokenId returns a boolean if a field has been set.
func (o *V1AuthAgentToken) HasAgentTokenId() bool {
	if o != nil && o.AgentTokenId != nil {
		return true
	}

	return false
}

// SetAgentTokenId gets a reference to the given int32 and assigns it to the AgentTokenId field.
func (o *V1AuthAgentToken) SetAgentTokenId(v int32) {
	o.AgentTokenId = &v
}

// GetAgentTokenCreatedAt returns the AgentTokenCreatedAt field value if set, zero value otherwise.
func (o *V1AuthAgentToken) GetAgentTokenCreatedAt() time.Time {
	if o == nil || o.AgentTokenCreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AgentTokenCreatedAt
}

// GetAgentTokenCreatedAtOk returns a tuple with the AgentTokenCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AuthAgentToken) GetAgentTokenCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.AgentTokenCreatedAt == nil {
		return nil, false
	}
	return o.AgentTokenCreatedAt, true
}

// HasAgentTokenCreatedAt returns a boolean if a field has been set.
func (o *V1AuthAgentToken) HasAgentTokenCreatedAt() bool {
	if o != nil && o.AgentTokenCreatedAt != nil {
		return true
	}

	return false
}

// SetAgentTokenCreatedAt gets a reference to the given time.Time and assigns it to the AgentTokenCreatedAt field.
func (o *V1AuthAgentToken) SetAgentTokenCreatedAt(v time.Time) {
	o.AgentTokenCreatedAt = &v
}

// GetAgentTokenUpdatedAt returns the AgentTokenUpdatedAt field value if set, zero value otherwise.
func (o *V1AuthAgentToken) GetAgentTokenUpdatedAt() time.Time {
	if o == nil || o.AgentTokenUpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AgentTokenUpdatedAt
}

// GetAgentTokenUpdatedAtOk returns a tuple with the AgentTokenUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AuthAgentToken) GetAgentTokenUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.AgentTokenUpdatedAt == nil {
		return nil, false
	}
	return o.AgentTokenUpdatedAt, true
}

// HasAgentTokenUpdatedAt returns a boolean if a field has been set.
func (o *V1AuthAgentToken) HasAgentTokenUpdatedAt() bool {
	if o != nil && o.AgentTokenUpdatedAt != nil {
		return true
	}

	return false
}

// SetAgentTokenUpdatedAt gets a reference to the given time.Time and assigns it to the AgentTokenUpdatedAt field.
func (o *V1AuthAgentToken) SetAgentTokenUpdatedAt(v time.Time) {
	o.AgentTokenUpdatedAt = &v
}

func (o V1AuthAgentToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AgentTokenName != nil {
		toSerialize["agent_token_name"] = o.AgentTokenName
	}
	if o.AgentTokenDescription.IsSet() {
		toSerialize["agent_token_description"] = o.AgentTokenDescription.Get()
	}
	if o.AgentTokenValidUntil != nil {
		toSerialize["agent_token_valid_until"] = o.AgentTokenValidUntil
	}
	if o.AgentTokenStatus != nil {
		toSerialize["agent_token_status"] = o.AgentTokenStatus
	}
	if o.AgentTokenId != nil {
		toSerialize["agent_token_id"] = o.AgentTokenId
	}
	if o.AgentTokenCreatedAt != nil {
		toSerialize["agent_token_created_at"] = o.AgentTokenCreatedAt
	}
	if o.AgentTokenUpdatedAt != nil {
		toSerialize["agent_token_updated_at"] = o.AgentTokenUpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableV1AuthAgentToken struct {
	value *V1AuthAgentToken
	isSet bool
}

func (v NullableV1AuthAgentToken) Get() *V1AuthAgentToken {
	return v.value
}

func (v *NullableV1AuthAgentToken) Set(val *V1AuthAgentToken) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AuthAgentToken) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AuthAgentToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AuthAgentToken(val *V1AuthAgentToken) *NullableV1AuthAgentToken {
	return &NullableV1AuthAgentToken{value: val, isSet: true}
}

func (v NullableV1AuthAgentToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AuthAgentToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}