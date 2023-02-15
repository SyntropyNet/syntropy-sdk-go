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
	"time"
)

// checks if the V1AgentFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1AgentFilter{}

// V1AgentFilter struct for V1AgentFilter
type V1AgentFilter struct {
	AgentId              []int32                  `json:"agent_id,omitempty"`
	AgentProviderId      []int32                  `json:"agent_provider_id,omitempty"`
	AgentTagId           []int32                  `json:"agent_tag_id,omitempty"`
	AgentType            []AgentType              `json:"agent_type,omitempty"`
	AgentVersion         []string                 `json:"agent_version,omitempty"`
	AgentTagName         []string                 `json:"agent_tag_name,omitempty"`
	AgentStatus          []AgentFilterAgentStatus `json:"agent_status,omitempty"`
	AgentLocationCountry []string                 `json:"agent_location_country,omitempty"`
	AgentModifiedAtFrom  *time.Time               `json:"agent_modified_at_from,omitempty"`
	AgentModifiedAtTo    *time.Time               `json:"agent_modified_at_to,omitempty"`
	AgentName            *string                  `json:"agent_name,omitempty"`
}

// NewV1AgentFilter instantiates a new V1AgentFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AgentFilter() *V1AgentFilter {
	this := V1AgentFilter{}
	return &this
}

// NewV1AgentFilterWithDefaults instantiates a new V1AgentFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AgentFilterWithDefaults() *V1AgentFilter {
	this := V1AgentFilter{}
	return &this
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *V1AgentFilter) GetAgentId() []int32 {
	if o == nil || isNil(o.AgentId) {
		var ret []int32
		return ret
	}
	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AgentFilter) GetAgentIdOk() ([]int32, bool) {
	if o == nil || isNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *V1AgentFilter) HasAgentId() bool {
	if o != nil && !isNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given []int32 and assigns it to the AgentId field.
func (o *V1AgentFilter) SetAgentId(v []int32) {
	o.AgentId = v
}

// GetAgentProviderId returns the AgentProviderId field value if set, zero value otherwise.
func (o *V1AgentFilter) GetAgentProviderId() []int32 {
	if o == nil || isNil(o.AgentProviderId) {
		var ret []int32
		return ret
	}
	return o.AgentProviderId
}

// GetAgentProviderIdOk returns a tuple with the AgentProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AgentFilter) GetAgentProviderIdOk() ([]int32, bool) {
	if o == nil || isNil(o.AgentProviderId) {
		return nil, false
	}
	return o.AgentProviderId, true
}

// HasAgentProviderId returns a boolean if a field has been set.
func (o *V1AgentFilter) HasAgentProviderId() bool {
	if o != nil && !isNil(o.AgentProviderId) {
		return true
	}

	return false
}

// SetAgentProviderId gets a reference to the given []int32 and assigns it to the AgentProviderId field.
func (o *V1AgentFilter) SetAgentProviderId(v []int32) {
	o.AgentProviderId = v
}

// GetAgentTagId returns the AgentTagId field value if set, zero value otherwise.
func (o *V1AgentFilter) GetAgentTagId() []int32 {
	if o == nil || isNil(o.AgentTagId) {
		var ret []int32
		return ret
	}
	return o.AgentTagId
}

// GetAgentTagIdOk returns a tuple with the AgentTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AgentFilter) GetAgentTagIdOk() ([]int32, bool) {
	if o == nil || isNil(o.AgentTagId) {
		return nil, false
	}
	return o.AgentTagId, true
}

// HasAgentTagId returns a boolean if a field has been set.
func (o *V1AgentFilter) HasAgentTagId() bool {
	if o != nil && !isNil(o.AgentTagId) {
		return true
	}

	return false
}

// SetAgentTagId gets a reference to the given []int32 and assigns it to the AgentTagId field.
func (o *V1AgentFilter) SetAgentTagId(v []int32) {
	o.AgentTagId = v
}

// GetAgentType returns the AgentType field value if set, zero value otherwise.
func (o *V1AgentFilter) GetAgentType() []AgentType {
	if o == nil || isNil(o.AgentType) {
		var ret []AgentType
		return ret
	}
	return o.AgentType
}

// GetAgentTypeOk returns a tuple with the AgentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AgentFilter) GetAgentTypeOk() ([]AgentType, bool) {
	if o == nil || isNil(o.AgentType) {
		return nil, false
	}
	return o.AgentType, true
}

// HasAgentType returns a boolean if a field has been set.
func (o *V1AgentFilter) HasAgentType() bool {
	if o != nil && !isNil(o.AgentType) {
		return true
	}

	return false
}

// SetAgentType gets a reference to the given []AgentType and assigns it to the AgentType field.
func (o *V1AgentFilter) SetAgentType(v []AgentType) {
	o.AgentType = v
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *V1AgentFilter) GetAgentVersion() []string {
	if o == nil || isNil(o.AgentVersion) {
		var ret []string
		return ret
	}
	return o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AgentFilter) GetAgentVersionOk() ([]string, bool) {
	if o == nil || isNil(o.AgentVersion) {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *V1AgentFilter) HasAgentVersion() bool {
	if o != nil && !isNil(o.AgentVersion) {
		return true
	}

	return false
}

// SetAgentVersion gets a reference to the given []string and assigns it to the AgentVersion field.
func (o *V1AgentFilter) SetAgentVersion(v []string) {
	o.AgentVersion = v
}

// GetAgentTagName returns the AgentTagName field value if set, zero value otherwise.
func (o *V1AgentFilter) GetAgentTagName() []string {
	if o == nil || isNil(o.AgentTagName) {
		var ret []string
		return ret
	}
	return o.AgentTagName
}

// GetAgentTagNameOk returns a tuple with the AgentTagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AgentFilter) GetAgentTagNameOk() ([]string, bool) {
	if o == nil || isNil(o.AgentTagName) {
		return nil, false
	}
	return o.AgentTagName, true
}

// HasAgentTagName returns a boolean if a field has been set.
func (o *V1AgentFilter) HasAgentTagName() bool {
	if o != nil && !isNil(o.AgentTagName) {
		return true
	}

	return false
}

// SetAgentTagName gets a reference to the given []string and assigns it to the AgentTagName field.
func (o *V1AgentFilter) SetAgentTagName(v []string) {
	o.AgentTagName = v
}

// GetAgentStatus returns the AgentStatus field value if set, zero value otherwise.
func (o *V1AgentFilter) GetAgentStatus() []AgentFilterAgentStatus {
	if o == nil || isNil(o.AgentStatus) {
		var ret []AgentFilterAgentStatus
		return ret
	}
	return o.AgentStatus
}

// GetAgentStatusOk returns a tuple with the AgentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AgentFilter) GetAgentStatusOk() ([]AgentFilterAgentStatus, bool) {
	if o == nil || isNil(o.AgentStatus) {
		return nil, false
	}
	return o.AgentStatus, true
}

// HasAgentStatus returns a boolean if a field has been set.
func (o *V1AgentFilter) HasAgentStatus() bool {
	if o != nil && !isNil(o.AgentStatus) {
		return true
	}

	return false
}

// SetAgentStatus gets a reference to the given []AgentFilterAgentStatus and assigns it to the AgentStatus field.
func (o *V1AgentFilter) SetAgentStatus(v []AgentFilterAgentStatus) {
	o.AgentStatus = v
}

// GetAgentLocationCountry returns the AgentLocationCountry field value if set, zero value otherwise.
func (o *V1AgentFilter) GetAgentLocationCountry() []string {
	if o == nil || isNil(o.AgentLocationCountry) {
		var ret []string
		return ret
	}
	return o.AgentLocationCountry
}

// GetAgentLocationCountryOk returns a tuple with the AgentLocationCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AgentFilter) GetAgentLocationCountryOk() ([]string, bool) {
	if o == nil || isNil(o.AgentLocationCountry) {
		return nil, false
	}
	return o.AgentLocationCountry, true
}

// HasAgentLocationCountry returns a boolean if a field has been set.
func (o *V1AgentFilter) HasAgentLocationCountry() bool {
	if o != nil && !isNil(o.AgentLocationCountry) {
		return true
	}

	return false
}

// SetAgentLocationCountry gets a reference to the given []string and assigns it to the AgentLocationCountry field.
func (o *V1AgentFilter) SetAgentLocationCountry(v []string) {
	o.AgentLocationCountry = v
}

// GetAgentModifiedAtFrom returns the AgentModifiedAtFrom field value if set, zero value otherwise.
func (o *V1AgentFilter) GetAgentModifiedAtFrom() time.Time {
	if o == nil || isNil(o.AgentModifiedAtFrom) {
		var ret time.Time
		return ret
	}
	return *o.AgentModifiedAtFrom
}

// GetAgentModifiedAtFromOk returns a tuple with the AgentModifiedAtFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AgentFilter) GetAgentModifiedAtFromOk() (*time.Time, bool) {
	if o == nil || isNil(o.AgentModifiedAtFrom) {
		return nil, false
	}
	return o.AgentModifiedAtFrom, true
}

// HasAgentModifiedAtFrom returns a boolean if a field has been set.
func (o *V1AgentFilter) HasAgentModifiedAtFrom() bool {
	if o != nil && !isNil(o.AgentModifiedAtFrom) {
		return true
	}

	return false
}

// SetAgentModifiedAtFrom gets a reference to the given time.Time and assigns it to the AgentModifiedAtFrom field.
func (o *V1AgentFilter) SetAgentModifiedAtFrom(v time.Time) {
	o.AgentModifiedAtFrom = &v
}

// GetAgentModifiedAtTo returns the AgentModifiedAtTo field value if set, zero value otherwise.
func (o *V1AgentFilter) GetAgentModifiedAtTo() time.Time {
	if o == nil || isNil(o.AgentModifiedAtTo) {
		var ret time.Time
		return ret
	}
	return *o.AgentModifiedAtTo
}

// GetAgentModifiedAtToOk returns a tuple with the AgentModifiedAtTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AgentFilter) GetAgentModifiedAtToOk() (*time.Time, bool) {
	if o == nil || isNil(o.AgentModifiedAtTo) {
		return nil, false
	}
	return o.AgentModifiedAtTo, true
}

// HasAgentModifiedAtTo returns a boolean if a field has been set.
func (o *V1AgentFilter) HasAgentModifiedAtTo() bool {
	if o != nil && !isNil(o.AgentModifiedAtTo) {
		return true
	}

	return false
}

// SetAgentModifiedAtTo gets a reference to the given time.Time and assigns it to the AgentModifiedAtTo field.
func (o *V1AgentFilter) SetAgentModifiedAtTo(v time.Time) {
	o.AgentModifiedAtTo = &v
}

// GetAgentName returns the AgentName field value if set, zero value otherwise.
func (o *V1AgentFilter) GetAgentName() string {
	if o == nil || isNil(o.AgentName) {
		var ret string
		return ret
	}
	return *o.AgentName
}

// GetAgentNameOk returns a tuple with the AgentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AgentFilter) GetAgentNameOk() (*string, bool) {
	if o == nil || isNil(o.AgentName) {
		return nil, false
	}
	return o.AgentName, true
}

// HasAgentName returns a boolean if a field has been set.
func (o *V1AgentFilter) HasAgentName() bool {
	if o != nil && !isNil(o.AgentName) {
		return true
	}

	return false
}

// SetAgentName gets a reference to the given string and assigns it to the AgentName field.
func (o *V1AgentFilter) SetAgentName(v string) {
	o.AgentName = &v
}

func (o V1AgentFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1AgentFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AgentId) {
		toSerialize["agent_id"] = o.AgentId
	}
	if !isNil(o.AgentProviderId) {
		toSerialize["agent_provider_id"] = o.AgentProviderId
	}
	if !isNil(o.AgentTagId) {
		toSerialize["agent_tag_id"] = o.AgentTagId
	}
	if !isNil(o.AgentType) {
		toSerialize["agent_type"] = o.AgentType
	}
	if !isNil(o.AgentVersion) {
		toSerialize["agent_version"] = o.AgentVersion
	}
	if !isNil(o.AgentTagName) {
		toSerialize["agent_tag_name"] = o.AgentTagName
	}
	if !isNil(o.AgentStatus) {
		toSerialize["agent_status"] = o.AgentStatus
	}
	if !isNil(o.AgentLocationCountry) {
		toSerialize["agent_location_country"] = o.AgentLocationCountry
	}
	if !isNil(o.AgentModifiedAtFrom) {
		toSerialize["agent_modified_at_from"] = o.AgentModifiedAtFrom
	}
	if !isNil(o.AgentModifiedAtTo) {
		toSerialize["agent_modified_at_to"] = o.AgentModifiedAtTo
	}
	if !isNil(o.AgentName) {
		toSerialize["agent_name"] = o.AgentName
	}
	return toSerialize, nil
}

type NullableV1AgentFilter struct {
	value *V1AgentFilter
	isSet bool
}

func (v NullableV1AgentFilter) Get() *V1AgentFilter {
	return v.value
}

func (v *NullableV1AgentFilter) Set(val *V1AgentFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AgentFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AgentFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AgentFilter(val *V1AgentFilter) *NullableV1AgentFilter {
	return &NullableV1AgentFilter{value: val, isSet: true}
}

func (v NullableV1AgentFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AgentFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
