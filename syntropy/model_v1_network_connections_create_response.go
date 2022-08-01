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

// V1NetworkConnectionsCreateResponse struct for V1NetworkConnectionsCreateResponse
type V1NetworkConnectionsCreateResponse struct {
	Data []V1ConnectionCreateItem `json:"data"`
}

// NewV1NetworkConnectionsCreateResponse instantiates a new V1NetworkConnectionsCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NetworkConnectionsCreateResponse(data []V1ConnectionCreateItem) *V1NetworkConnectionsCreateResponse {
	this := V1NetworkConnectionsCreateResponse{}
	this.Data = data
	return &this
}

// NewV1NetworkConnectionsCreateResponseWithDefaults instantiates a new V1NetworkConnectionsCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NetworkConnectionsCreateResponseWithDefaults() *V1NetworkConnectionsCreateResponse {
	this := V1NetworkConnectionsCreateResponse{}
	return &this
}

// GetData returns the Data field value
func (o *V1NetworkConnectionsCreateResponse) GetData() []V1ConnectionCreateItem {
	if o == nil {
		var ret []V1ConnectionCreateItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V1NetworkConnectionsCreateResponse) GetDataOk() ([]V1ConnectionCreateItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *V1NetworkConnectionsCreateResponse) SetData(v []V1ConnectionCreateItem) {
	o.Data = v
}

func (o V1NetworkConnectionsCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableV1NetworkConnectionsCreateResponse struct {
	value *V1NetworkConnectionsCreateResponse
	isSet bool
}

func (v NullableV1NetworkConnectionsCreateResponse) Get() *V1NetworkConnectionsCreateResponse {
	return v.value
}

func (v *NullableV1NetworkConnectionsCreateResponse) Set(val *V1NetworkConnectionsCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NetworkConnectionsCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NetworkConnectionsCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NetworkConnectionsCreateResponse(val *V1NetworkConnectionsCreateResponse) *NullableV1NetworkConnectionsCreateResponse {
	return &NullableV1NetworkConnectionsCreateResponse{value: val, isSet: true}
}

func (v NullableV1NetworkConnectionsCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NetworkConnectionsCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}