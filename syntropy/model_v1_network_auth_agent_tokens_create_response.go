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

// V1NetworkAuthAgentTokensCreateResponse struct for V1NetworkAuthAgentTokensCreateResponse
type V1NetworkAuthAgentTokensCreateResponse struct {
	Data V1AuthApiKeysCreateItem `json:"data"`
}

// NewV1NetworkAuthAgentTokensCreateResponse instantiates a new V1NetworkAuthAgentTokensCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NetworkAuthAgentTokensCreateResponse(data V1AuthApiKeysCreateItem) *V1NetworkAuthAgentTokensCreateResponse {
	this := V1NetworkAuthAgentTokensCreateResponse{}
	this.Data = data
	return &this
}

// NewV1NetworkAuthAgentTokensCreateResponseWithDefaults instantiates a new V1NetworkAuthAgentTokensCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NetworkAuthAgentTokensCreateResponseWithDefaults() *V1NetworkAuthAgentTokensCreateResponse {
	this := V1NetworkAuthAgentTokensCreateResponse{}
	return &this
}

// GetData returns the Data field value
func (o *V1NetworkAuthAgentTokensCreateResponse) GetData() V1AuthApiKeysCreateItem {
	if o == nil {
		var ret V1AuthApiKeysCreateItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V1NetworkAuthAgentTokensCreateResponse) GetDataOk() (*V1AuthApiKeysCreateItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *V1NetworkAuthAgentTokensCreateResponse) SetData(v V1AuthApiKeysCreateItem) {
	o.Data = v
}

func (o V1NetworkAuthAgentTokensCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableV1NetworkAuthAgentTokensCreateResponse struct {
	value *V1NetworkAuthAgentTokensCreateResponse
	isSet bool
}

func (v NullableV1NetworkAuthAgentTokensCreateResponse) Get() *V1NetworkAuthAgentTokensCreateResponse {
	return v.value
}

func (v *NullableV1NetworkAuthAgentTokensCreateResponse) Set(val *V1NetworkAuthAgentTokensCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NetworkAuthAgentTokensCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NetworkAuthAgentTokensCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NetworkAuthAgentTokensCreateResponse(val *V1NetworkAuthAgentTokensCreateResponse) *NullableV1NetworkAuthAgentTokensCreateResponse {
	return &NullableV1NetworkAuthAgentTokensCreateResponse{value: val, isSet: true}
}

func (v NullableV1NetworkAuthAgentTokensCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NetworkAuthAgentTokensCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
