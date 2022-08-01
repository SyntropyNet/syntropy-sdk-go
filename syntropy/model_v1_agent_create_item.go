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

// V1AgentCreateItem struct for V1AgentCreateItem
type V1AgentCreateItem struct {
	AgentServicesDefaultStatus bool                `json:"agent_services_default_status"`
	ApiKeyId                   int32               `json:"api_key_id"`
	AgentPublicIpv4            string              `json:"agent_public_ipv4"`
	AgentLocationLat           NullableFloat32     `json:"agent_location_lat"`
	AgentLocationLon           NullableFloat32     `json:"agent_location_lon"`
	AgentLocationCity          NullableString      `json:"agent_location_city"`
	AgentLocationCountry       NullableString      `json:"agent_location_country"`
	AgentDeviceId              string              `json:"agent_device_id"`
	AgentName                  string              `json:"agent_name"`
	AgentStatus                NullableAgentStatus `json:"agent_status"`
	AgentVersion               NullableString      `json:"agent_version"`
	AgentProviderId            NullableInt32       `json:"agent_provider_id,omitempty"`
	AgentLockedFields          AgentLockedFields   `json:"agent_locked_fields"`
	AgentModifiedAt            NullableTime        `json:"agent_modified_at"`
	AgentIsVirtual             bool                `json:"agent_is_virtual"`
	AgentType                  AgentType           `json:"agent_type"`
	AgentInternalSubnet        string              `json:"agent_internal_subnet"`
	AgentCreatedAt             time.Time           `json:"agent_created_at"`
	AgentUpdatedAt             time.Time           `json:"agent_updated_at"`
	AgentId                    int32               `json:"agent_id"`
}

// NewV1AgentCreateItem instantiates a new V1AgentCreateItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AgentCreateItem(agentServicesDefaultStatus bool, apiKeyId int32, agentPublicIpv4 string, agentLocationLat NullableFloat32, agentLocationLon NullableFloat32, agentLocationCity NullableString, agentLocationCountry NullableString, agentDeviceId string, agentName string, agentStatus NullableAgentStatus, agentVersion NullableString, agentLockedFields AgentLockedFields, agentModifiedAt NullableTime, agentIsVirtual bool, agentType AgentType, agentInternalSubnet string, agentCreatedAt time.Time, agentUpdatedAt time.Time, agentId int32) *V1AgentCreateItem {
	this := V1AgentCreateItem{}
	this.AgentServicesDefaultStatus = agentServicesDefaultStatus
	this.ApiKeyId = apiKeyId
	this.AgentPublicIpv4 = agentPublicIpv4
	this.AgentLocationLat = agentLocationLat
	this.AgentLocationLon = agentLocationLon
	this.AgentLocationCity = agentLocationCity
	this.AgentLocationCountry = agentLocationCountry
	this.AgentDeviceId = agentDeviceId
	this.AgentName = agentName
	this.AgentStatus = agentStatus
	this.AgentVersion = agentVersion
	this.AgentLockedFields = agentLockedFields
	this.AgentModifiedAt = agentModifiedAt
	this.AgentIsVirtual = agentIsVirtual
	this.AgentType = agentType
	this.AgentInternalSubnet = agentInternalSubnet
	this.AgentCreatedAt = agentCreatedAt
	this.AgentUpdatedAt = agentUpdatedAt
	this.AgentId = agentId
	return &this
}

// NewV1AgentCreateItemWithDefaults instantiates a new V1AgentCreateItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AgentCreateItemWithDefaults() *V1AgentCreateItem {
	this := V1AgentCreateItem{}
	return &this
}

// GetAgentServicesDefaultStatus returns the AgentServicesDefaultStatus field value
func (o *V1AgentCreateItem) GetAgentServicesDefaultStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AgentServicesDefaultStatus
}

// GetAgentServicesDefaultStatusOk returns a tuple with the AgentServicesDefaultStatus field value
// and a boolean to check if the value has been set.
func (o *V1AgentCreateItem) GetAgentServicesDefaultStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentServicesDefaultStatus, true
}

// SetAgentServicesDefaultStatus sets field value
func (o *V1AgentCreateItem) SetAgentServicesDefaultStatus(v bool) {
	o.AgentServicesDefaultStatus = v
}

// GetApiKeyId returns the ApiKeyId field value
func (o *V1AgentCreateItem) GetApiKeyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value
// and a boolean to check if the value has been set.
func (o *V1AgentCreateItem) GetApiKeyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKeyId, true
}

// SetApiKeyId sets field value
func (o *V1AgentCreateItem) SetApiKeyId(v int32) {
	o.ApiKeyId = v
}

// GetAgentPublicIpv4 returns the AgentPublicIpv4 field value
func (o *V1AgentCreateItem) GetAgentPublicIpv4() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentPublicIpv4
}

// GetAgentPublicIpv4Ok returns a tuple with the AgentPublicIpv4 field value
// and a boolean to check if the value has been set.
func (o *V1AgentCreateItem) GetAgentPublicIpv4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentPublicIpv4, true
}

// SetAgentPublicIpv4 sets field value
func (o *V1AgentCreateItem) SetAgentPublicIpv4(v string) {
	o.AgentPublicIpv4 = v
}

// GetAgentLocationLat returns the AgentLocationLat field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *V1AgentCreateItem) GetAgentLocationLat() float32 {
	if o == nil || o.AgentLocationLat.Get() == nil {
		var ret float32
		return ret
	}

	return *o.AgentLocationLat.Get()
}

// GetAgentLocationLatOk returns a tuple with the AgentLocationLat field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V1AgentCreateItem) GetAgentLocationLatOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentLocationLat.Get(), o.AgentLocationLat.IsSet()
}

// SetAgentLocationLat sets field value
func (o *V1AgentCreateItem) SetAgentLocationLat(v float32) {
	o.AgentLocationLat.Set(&v)
}

// GetAgentLocationLon returns the AgentLocationLon field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *V1AgentCreateItem) GetAgentLocationLon() float32 {
	if o == nil || o.AgentLocationLon.Get() == nil {
		var ret float32
		return ret
	}

	return *o.AgentLocationLon.Get()
}

// GetAgentLocationLonOk returns a tuple with the AgentLocationLon field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V1AgentCreateItem) GetAgentLocationLonOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentLocationLon.Get(), o.AgentLocationLon.IsSet()
}

// SetAgentLocationLon sets field value
func (o *V1AgentCreateItem) SetAgentLocationLon(v float32) {
	o.AgentLocationLon.Set(&v)
}

// GetAgentLocationCity returns the AgentLocationCity field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V1AgentCreateItem) GetAgentLocationCity() string {
	if o == nil || o.AgentLocationCity.Get() == nil {
		var ret string
		return ret
	}

	return *o.AgentLocationCity.Get()
}

// GetAgentLocationCityOk returns a tuple with the AgentLocationCity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V1AgentCreateItem) GetAgentLocationCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentLocationCity.Get(), o.AgentLocationCity.IsSet()
}

// SetAgentLocationCity sets field value
func (o *V1AgentCreateItem) SetAgentLocationCity(v string) {
	o.AgentLocationCity.Set(&v)
}

// GetAgentLocationCountry returns the AgentLocationCountry field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V1AgentCreateItem) GetAgentLocationCountry() string {
	if o == nil || o.AgentLocationCountry.Get() == nil {
		var ret string
		return ret
	}

	return *o.AgentLocationCountry.Get()
}

// GetAgentLocationCountryOk returns a tuple with the AgentLocationCountry field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V1AgentCreateItem) GetAgentLocationCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentLocationCountry.Get(), o.AgentLocationCountry.IsSet()
}

// SetAgentLocationCountry sets field value
func (o *V1AgentCreateItem) SetAgentLocationCountry(v string) {
	o.AgentLocationCountry.Set(&v)
}

// GetAgentDeviceId returns the AgentDeviceId field value
func (o *V1AgentCreateItem) GetAgentDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentDeviceId
}

// GetAgentDeviceIdOk returns a tuple with the AgentDeviceId field value
// and a boolean to check if the value has been set.
func (o *V1AgentCreateItem) GetAgentDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentDeviceId, true
}

// SetAgentDeviceId sets field value
func (o *V1AgentCreateItem) SetAgentDeviceId(v string) {
	o.AgentDeviceId = v
}

// GetAgentName returns the AgentName field value
func (o *V1AgentCreateItem) GetAgentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentName
}

// GetAgentNameOk returns a tuple with the AgentName field value
// and a boolean to check if the value has been set.
func (o *V1AgentCreateItem) GetAgentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentName, true
}

// SetAgentName sets field value
func (o *V1AgentCreateItem) SetAgentName(v string) {
	o.AgentName = v
}

// GetAgentStatus returns the AgentStatus field value
// If the value is explicit nil, the zero value for AgentStatus will be returned
func (o *V1AgentCreateItem) GetAgentStatus() AgentStatus {
	if o == nil || o.AgentStatus.Get() == nil {
		var ret AgentStatus
		return ret
	}

	return *o.AgentStatus.Get()
}

// GetAgentStatusOk returns a tuple with the AgentStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V1AgentCreateItem) GetAgentStatusOk() (*AgentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentStatus.Get(), o.AgentStatus.IsSet()
}

// SetAgentStatus sets field value
func (o *V1AgentCreateItem) SetAgentStatus(v AgentStatus) {
	o.AgentStatus.Set(&v)
}

// GetAgentVersion returns the AgentVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V1AgentCreateItem) GetAgentVersion() string {
	if o == nil || o.AgentVersion.Get() == nil {
		var ret string
		return ret
	}

	return *o.AgentVersion.Get()
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V1AgentCreateItem) GetAgentVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentVersion.Get(), o.AgentVersion.IsSet()
}

// SetAgentVersion sets field value
func (o *V1AgentCreateItem) SetAgentVersion(v string) {
	o.AgentVersion.Set(&v)
}

// GetAgentProviderId returns the AgentProviderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *V1AgentCreateItem) GetAgentProviderId() int32 {
	if o == nil || o.AgentProviderId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AgentProviderId.Get()
}

// GetAgentProviderIdOk returns a tuple with the AgentProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V1AgentCreateItem) GetAgentProviderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentProviderId.Get(), o.AgentProviderId.IsSet()
}

// HasAgentProviderId returns a boolean if a field has been set.
func (o *V1AgentCreateItem) HasAgentProviderId() bool {
	if o != nil && o.AgentProviderId.IsSet() {
		return true
	}

	return false
}

// SetAgentProviderId gets a reference to the given NullableInt32 and assigns it to the AgentProviderId field.
func (o *V1AgentCreateItem) SetAgentProviderId(v int32) {
	o.AgentProviderId.Set(&v)
}

// SetAgentProviderIdNil sets the value for AgentProviderId to be an explicit nil
func (o *V1AgentCreateItem) SetAgentProviderIdNil() {
	o.AgentProviderId.Set(nil)
}

// UnsetAgentProviderId ensures that no value is present for AgentProviderId, not even an explicit nil
func (o *V1AgentCreateItem) UnsetAgentProviderId() {
	o.AgentProviderId.Unset()
}

// GetAgentLockedFields returns the AgentLockedFields field value
func (o *V1AgentCreateItem) GetAgentLockedFields() AgentLockedFields {
	if o == nil {
		var ret AgentLockedFields
		return ret
	}

	return o.AgentLockedFields
}

// GetAgentLockedFieldsOk returns a tuple with the AgentLockedFields field value
// and a boolean to check if the value has been set.
func (o *V1AgentCreateItem) GetAgentLockedFieldsOk() (*AgentLockedFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentLockedFields, true
}

// SetAgentLockedFields sets field value
func (o *V1AgentCreateItem) SetAgentLockedFields(v AgentLockedFields) {
	o.AgentLockedFields = v
}

// GetAgentModifiedAt returns the AgentModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *V1AgentCreateItem) GetAgentModifiedAt() time.Time {
	if o == nil || o.AgentModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.AgentModifiedAt.Get()
}

// GetAgentModifiedAtOk returns a tuple with the AgentModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V1AgentCreateItem) GetAgentModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentModifiedAt.Get(), o.AgentModifiedAt.IsSet()
}

// SetAgentModifiedAt sets field value
func (o *V1AgentCreateItem) SetAgentModifiedAt(v time.Time) {
	o.AgentModifiedAt.Set(&v)
}

// GetAgentIsVirtual returns the AgentIsVirtual field value
func (o *V1AgentCreateItem) GetAgentIsVirtual() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AgentIsVirtual
}

// GetAgentIsVirtualOk returns a tuple with the AgentIsVirtual field value
// and a boolean to check if the value has been set.
func (o *V1AgentCreateItem) GetAgentIsVirtualOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentIsVirtual, true
}

// SetAgentIsVirtual sets field value
func (o *V1AgentCreateItem) SetAgentIsVirtual(v bool) {
	o.AgentIsVirtual = v
}

// GetAgentType returns the AgentType field value
func (o *V1AgentCreateItem) GetAgentType() AgentType {
	if o == nil {
		var ret AgentType
		return ret
	}

	return o.AgentType
}

// GetAgentTypeOk returns a tuple with the AgentType field value
// and a boolean to check if the value has been set.
func (o *V1AgentCreateItem) GetAgentTypeOk() (*AgentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentType, true
}

// SetAgentType sets field value
func (o *V1AgentCreateItem) SetAgentType(v AgentType) {
	o.AgentType = v
}

// GetAgentInternalSubnet returns the AgentInternalSubnet field value
func (o *V1AgentCreateItem) GetAgentInternalSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentInternalSubnet
}

// GetAgentInternalSubnetOk returns a tuple with the AgentInternalSubnet field value
// and a boolean to check if the value has been set.
func (o *V1AgentCreateItem) GetAgentInternalSubnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentInternalSubnet, true
}

// SetAgentInternalSubnet sets field value
func (o *V1AgentCreateItem) SetAgentInternalSubnet(v string) {
	o.AgentInternalSubnet = v
}

// GetAgentCreatedAt returns the AgentCreatedAt field value
func (o *V1AgentCreateItem) GetAgentCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AgentCreatedAt
}

// GetAgentCreatedAtOk returns a tuple with the AgentCreatedAt field value
// and a boolean to check if the value has been set.
func (o *V1AgentCreateItem) GetAgentCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentCreatedAt, true
}

// SetAgentCreatedAt sets field value
func (o *V1AgentCreateItem) SetAgentCreatedAt(v time.Time) {
	o.AgentCreatedAt = v
}

// GetAgentUpdatedAt returns the AgentUpdatedAt field value
func (o *V1AgentCreateItem) GetAgentUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AgentUpdatedAt
}

// GetAgentUpdatedAtOk returns a tuple with the AgentUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *V1AgentCreateItem) GetAgentUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentUpdatedAt, true
}

// SetAgentUpdatedAt sets field value
func (o *V1AgentCreateItem) SetAgentUpdatedAt(v time.Time) {
	o.AgentUpdatedAt = v
}

// GetAgentId returns the AgentId field value
func (o *V1AgentCreateItem) GetAgentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value
// and a boolean to check if the value has been set.
func (o *V1AgentCreateItem) GetAgentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentId, true
}

// SetAgentId sets field value
func (o *V1AgentCreateItem) SetAgentId(v int32) {
	o.AgentId = v
}

func (o V1AgentCreateItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["agent_services_default_status"] = o.AgentServicesDefaultStatus
	}
	if true {
		toSerialize["api_key_id"] = o.ApiKeyId
	}
	if true {
		toSerialize["agent_public_ipv4"] = o.AgentPublicIpv4
	}
	if true {
		toSerialize["agent_location_lat"] = o.AgentLocationLat.Get()
	}
	if true {
		toSerialize["agent_location_lon"] = o.AgentLocationLon.Get()
	}
	if true {
		toSerialize["agent_location_city"] = o.AgentLocationCity.Get()
	}
	if true {
		toSerialize["agent_location_country"] = o.AgentLocationCountry.Get()
	}
	if true {
		toSerialize["agent_device_id"] = o.AgentDeviceId
	}
	if true {
		toSerialize["agent_name"] = o.AgentName
	}
	if true {
		toSerialize["agent_status"] = o.AgentStatus.Get()
	}
	if true {
		toSerialize["agent_version"] = o.AgentVersion.Get()
	}
	if o.AgentProviderId.IsSet() {
		toSerialize["agent_provider_id"] = o.AgentProviderId.Get()
	}
	if true {
		toSerialize["agent_locked_fields"] = o.AgentLockedFields
	}
	if true {
		toSerialize["agent_modified_at"] = o.AgentModifiedAt.Get()
	}
	if true {
		toSerialize["agent_is_virtual"] = o.AgentIsVirtual
	}
	if true {
		toSerialize["agent_type"] = o.AgentType
	}
	if true {
		toSerialize["agent_internal_subnet"] = o.AgentInternalSubnet
	}
	if true {
		toSerialize["agent_created_at"] = o.AgentCreatedAt
	}
	if true {
		toSerialize["agent_updated_at"] = o.AgentUpdatedAt
	}
	if true {
		toSerialize["agent_id"] = o.AgentId
	}
	return json.Marshal(toSerialize)
}

type NullableV1AgentCreateItem struct {
	value *V1AgentCreateItem
	isSet bool
}

func (v NullableV1AgentCreateItem) Get() *V1AgentCreateItem {
	return v.value
}

func (v *NullableV1AgentCreateItem) Set(val *V1AgentCreateItem) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AgentCreateItem) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AgentCreateItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AgentCreateItem(val *V1AgentCreateItem) *NullableV1AgentCreateItem {
	return &NullableV1AgentCreateItem{value: val, isSet: true}
}

func (v NullableV1AgentCreateItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AgentCreateItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
