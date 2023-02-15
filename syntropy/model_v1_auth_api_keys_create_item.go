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

// checks if the V1AuthApiKeysCreateItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1AuthApiKeysCreateItem{}

// V1AuthApiKeysCreateItem struct for V1AuthApiKeysCreateItem
type V1AuthApiKeysCreateItem struct {
	ApiKeyName            string     `json:"api_key_name"`
	ApiKeyDescription     *string    `json:"api_key_description,omitempty"`
	ApiKeyValidUntil      *time.Time `json:"api_key_valid_until,omitempty"`
	ApiKeyAllowedTagNames []string   `json:"api_key_allowed_tag_names,omitempty"`
	ApiKeySecret          string     `json:"api_key_secret"`
	UserId                int32      `json:"user_id"`
	ApiKeyId              int32      `json:"api_key_id"`
	ApiKeyCreatedAt       time.Time  `json:"api_key_created_at"`
	ApiKeyUpdatedAt       time.Time  `json:"api_key_updated_at"`
}

// NewV1AuthApiKeysCreateItem instantiates a new V1AuthApiKeysCreateItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AuthApiKeysCreateItem(apiKeyName string, apiKeySecret string, userId int32, apiKeyId int32, apiKeyCreatedAt time.Time, apiKeyUpdatedAt time.Time) *V1AuthApiKeysCreateItem {
	this := V1AuthApiKeysCreateItem{}
	this.ApiKeyName = apiKeyName
	this.ApiKeySecret = apiKeySecret
	this.UserId = userId
	this.ApiKeyId = apiKeyId
	this.ApiKeyCreatedAt = apiKeyCreatedAt
	this.ApiKeyUpdatedAt = apiKeyUpdatedAt
	return &this
}

// NewV1AuthApiKeysCreateItemWithDefaults instantiates a new V1AuthApiKeysCreateItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AuthApiKeysCreateItemWithDefaults() *V1AuthApiKeysCreateItem {
	this := V1AuthApiKeysCreateItem{}
	return &this
}

// GetApiKeyName returns the ApiKeyName field value
func (o *V1AuthApiKeysCreateItem) GetApiKeyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKeyName
}

// GetApiKeyNameOk returns a tuple with the ApiKeyName field value
// and a boolean to check if the value has been set.
func (o *V1AuthApiKeysCreateItem) GetApiKeyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKeyName, true
}

// SetApiKeyName sets field value
func (o *V1AuthApiKeysCreateItem) SetApiKeyName(v string) {
	o.ApiKeyName = v
}

// GetApiKeyDescription returns the ApiKeyDescription field value if set, zero value otherwise.
func (o *V1AuthApiKeysCreateItem) GetApiKeyDescription() string {
	if o == nil || isNil(o.ApiKeyDescription) {
		var ret string
		return ret
	}
	return *o.ApiKeyDescription
}

// GetApiKeyDescriptionOk returns a tuple with the ApiKeyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AuthApiKeysCreateItem) GetApiKeyDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.ApiKeyDescription) {
		return nil, false
	}
	return o.ApiKeyDescription, true
}

// HasApiKeyDescription returns a boolean if a field has been set.
func (o *V1AuthApiKeysCreateItem) HasApiKeyDescription() bool {
	if o != nil && !isNil(o.ApiKeyDescription) {
		return true
	}

	return false
}

// SetApiKeyDescription gets a reference to the given string and assigns it to the ApiKeyDescription field.
func (o *V1AuthApiKeysCreateItem) SetApiKeyDescription(v string) {
	o.ApiKeyDescription = &v
}

// GetApiKeyValidUntil returns the ApiKeyValidUntil field value if set, zero value otherwise.
func (o *V1AuthApiKeysCreateItem) GetApiKeyValidUntil() time.Time {
	if o == nil || isNil(o.ApiKeyValidUntil) {
		var ret time.Time
		return ret
	}
	return *o.ApiKeyValidUntil
}

// GetApiKeyValidUntilOk returns a tuple with the ApiKeyValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AuthApiKeysCreateItem) GetApiKeyValidUntilOk() (*time.Time, bool) {
	if o == nil || isNil(o.ApiKeyValidUntil) {
		return nil, false
	}
	return o.ApiKeyValidUntil, true
}

// HasApiKeyValidUntil returns a boolean if a field has been set.
func (o *V1AuthApiKeysCreateItem) HasApiKeyValidUntil() bool {
	if o != nil && !isNil(o.ApiKeyValidUntil) {
		return true
	}

	return false
}

// SetApiKeyValidUntil gets a reference to the given time.Time and assigns it to the ApiKeyValidUntil field.
func (o *V1AuthApiKeysCreateItem) SetApiKeyValidUntil(v time.Time) {
	o.ApiKeyValidUntil = &v
}

// GetApiKeyAllowedTagNames returns the ApiKeyAllowedTagNames field value if set, zero value otherwise.
func (o *V1AuthApiKeysCreateItem) GetApiKeyAllowedTagNames() []string {
	if o == nil || isNil(o.ApiKeyAllowedTagNames) {
		var ret []string
		return ret
	}
	return o.ApiKeyAllowedTagNames
}

// GetApiKeyAllowedTagNamesOk returns a tuple with the ApiKeyAllowedTagNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AuthApiKeysCreateItem) GetApiKeyAllowedTagNamesOk() ([]string, bool) {
	if o == nil || isNil(o.ApiKeyAllowedTagNames) {
		return nil, false
	}
	return o.ApiKeyAllowedTagNames, true
}

// HasApiKeyAllowedTagNames returns a boolean if a field has been set.
func (o *V1AuthApiKeysCreateItem) HasApiKeyAllowedTagNames() bool {
	if o != nil && !isNil(o.ApiKeyAllowedTagNames) {
		return true
	}

	return false
}

// SetApiKeyAllowedTagNames gets a reference to the given []string and assigns it to the ApiKeyAllowedTagNames field.
func (o *V1AuthApiKeysCreateItem) SetApiKeyAllowedTagNames(v []string) {
	o.ApiKeyAllowedTagNames = v
}

// GetApiKeySecret returns the ApiKeySecret field value
func (o *V1AuthApiKeysCreateItem) GetApiKeySecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKeySecret
}

// GetApiKeySecretOk returns a tuple with the ApiKeySecret field value
// and a boolean to check if the value has been set.
func (o *V1AuthApiKeysCreateItem) GetApiKeySecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKeySecret, true
}

// SetApiKeySecret sets field value
func (o *V1AuthApiKeysCreateItem) SetApiKeySecret(v string) {
	o.ApiKeySecret = v
}

// GetUserId returns the UserId field value
func (o *V1AuthApiKeysCreateItem) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *V1AuthApiKeysCreateItem) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *V1AuthApiKeysCreateItem) SetUserId(v int32) {
	o.UserId = v
}

// GetApiKeyId returns the ApiKeyId field value
func (o *V1AuthApiKeysCreateItem) GetApiKeyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value
// and a boolean to check if the value has been set.
func (o *V1AuthApiKeysCreateItem) GetApiKeyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKeyId, true
}

// SetApiKeyId sets field value
func (o *V1AuthApiKeysCreateItem) SetApiKeyId(v int32) {
	o.ApiKeyId = v
}

// GetApiKeyCreatedAt returns the ApiKeyCreatedAt field value
func (o *V1AuthApiKeysCreateItem) GetApiKeyCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ApiKeyCreatedAt
}

// GetApiKeyCreatedAtOk returns a tuple with the ApiKeyCreatedAt field value
// and a boolean to check if the value has been set.
func (o *V1AuthApiKeysCreateItem) GetApiKeyCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKeyCreatedAt, true
}

// SetApiKeyCreatedAt sets field value
func (o *V1AuthApiKeysCreateItem) SetApiKeyCreatedAt(v time.Time) {
	o.ApiKeyCreatedAt = v
}

// GetApiKeyUpdatedAt returns the ApiKeyUpdatedAt field value
func (o *V1AuthApiKeysCreateItem) GetApiKeyUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ApiKeyUpdatedAt
}

// GetApiKeyUpdatedAtOk returns a tuple with the ApiKeyUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *V1AuthApiKeysCreateItem) GetApiKeyUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKeyUpdatedAt, true
}

// SetApiKeyUpdatedAt sets field value
func (o *V1AuthApiKeysCreateItem) SetApiKeyUpdatedAt(v time.Time) {
	o.ApiKeyUpdatedAt = v
}

func (o V1AuthApiKeysCreateItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1AuthApiKeysCreateItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["api_key_name"] = o.ApiKeyName
	if !isNil(o.ApiKeyDescription) {
		toSerialize["api_key_description"] = o.ApiKeyDescription
	}
	if !isNil(o.ApiKeyValidUntil) {
		toSerialize["api_key_valid_until"] = o.ApiKeyValidUntil
	}
	if !isNil(o.ApiKeyAllowedTagNames) {
		toSerialize["api_key_allowed_tag_names"] = o.ApiKeyAllowedTagNames
	}
	toSerialize["api_key_secret"] = o.ApiKeySecret
	toSerialize["user_id"] = o.UserId
	toSerialize["api_key_id"] = o.ApiKeyId
	toSerialize["api_key_created_at"] = o.ApiKeyCreatedAt
	toSerialize["api_key_updated_at"] = o.ApiKeyUpdatedAt
	return toSerialize, nil
}

type NullableV1AuthApiKeysCreateItem struct {
	value *V1AuthApiKeysCreateItem
	isSet bool
}

func (v NullableV1AuthApiKeysCreateItem) Get() *V1AuthApiKeysCreateItem {
	return v.value
}

func (v *NullableV1AuthApiKeysCreateItem) Set(val *V1AuthApiKeysCreateItem) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AuthApiKeysCreateItem) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AuthApiKeysCreateItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AuthApiKeysCreateItem(val *V1AuthApiKeysCreateItem) *NullableV1AuthApiKeysCreateItem {
	return &NullableV1AuthApiKeysCreateItem{value: val, isSet: true}
}

func (v NullableV1AuthApiKeysCreateItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AuthApiKeysCreateItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
