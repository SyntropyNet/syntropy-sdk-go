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

// V1NetworkConnectionsUpdateRequest struct for V1NetworkConnectionsUpdateRequest
type V1NetworkConnectionsUpdateRequest struct {
	Changes []V1ConnectionUpdateChange `json:"changes"`
}

// NewV1NetworkConnectionsUpdateRequest instantiates a new V1NetworkConnectionsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NetworkConnectionsUpdateRequest(changes []V1ConnectionUpdateChange) *V1NetworkConnectionsUpdateRequest {
	this := V1NetworkConnectionsUpdateRequest{}
	this.Changes = changes
	return &this
}

// NewV1NetworkConnectionsUpdateRequestWithDefaults instantiates a new V1NetworkConnectionsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NetworkConnectionsUpdateRequestWithDefaults() *V1NetworkConnectionsUpdateRequest {
	this := V1NetworkConnectionsUpdateRequest{}
	return &this
}

// GetChanges returns the Changes field value
func (o *V1NetworkConnectionsUpdateRequest) GetChanges() []V1ConnectionUpdateChange {
	if o == nil {
		var ret []V1ConnectionUpdateChange
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *V1NetworkConnectionsUpdateRequest) GetChangesOk() ([]V1ConnectionUpdateChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *V1NetworkConnectionsUpdateRequest) SetChanges(v []V1ConnectionUpdateChange) {
	o.Changes = v
}

func (o V1NetworkConnectionsUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["changes"] = o.Changes
	}
	return json.Marshal(toSerialize)
}

type NullableV1NetworkConnectionsUpdateRequest struct {
	value *V1NetworkConnectionsUpdateRequest
	isSet bool
}

func (v NullableV1NetworkConnectionsUpdateRequest) Get() *V1NetworkConnectionsUpdateRequest {
	return v.value
}

func (v *NullableV1NetworkConnectionsUpdateRequest) Set(val *V1NetworkConnectionsUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NetworkConnectionsUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NetworkConnectionsUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NetworkConnectionsUpdateRequest(val *V1NetworkConnectionsUpdateRequest) *NullableV1NetworkConnectionsUpdateRequest {
	return &NullableV1NetworkConnectionsUpdateRequest{value: val, isSet: true}
}

func (v NullableV1NetworkConnectionsUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NetworkConnectionsUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
