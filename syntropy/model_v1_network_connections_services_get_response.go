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

// V1NetworkConnectionsServicesGetResponse struct for V1NetworkConnectionsServicesGetResponse
type V1NetworkConnectionsServicesGetResponse struct {
	Data []V1ConnectionService `json:"data"`
}

// NewV1NetworkConnectionsServicesGetResponse instantiates a new V1NetworkConnectionsServicesGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NetworkConnectionsServicesGetResponse(data []V1ConnectionService) *V1NetworkConnectionsServicesGetResponse {
	this := V1NetworkConnectionsServicesGetResponse{}
	this.Data = data
	return &this
}

// NewV1NetworkConnectionsServicesGetResponseWithDefaults instantiates a new V1NetworkConnectionsServicesGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NetworkConnectionsServicesGetResponseWithDefaults() *V1NetworkConnectionsServicesGetResponse {
	this := V1NetworkConnectionsServicesGetResponse{}
	return &this
}

// GetData returns the Data field value
func (o *V1NetworkConnectionsServicesGetResponse) GetData() []V1ConnectionService {
	if o == nil {
		var ret []V1ConnectionService
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V1NetworkConnectionsServicesGetResponse) GetDataOk() ([]V1ConnectionService, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *V1NetworkConnectionsServicesGetResponse) SetData(v []V1ConnectionService) {
	o.Data = v
}

func (o V1NetworkConnectionsServicesGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableV1NetworkConnectionsServicesGetResponse struct {
	value *V1NetworkConnectionsServicesGetResponse
	isSet bool
}

func (v NullableV1NetworkConnectionsServicesGetResponse) Get() *V1NetworkConnectionsServicesGetResponse {
	return v.value
}

func (v *NullableV1NetworkConnectionsServicesGetResponse) Set(val *V1NetworkConnectionsServicesGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NetworkConnectionsServicesGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NetworkConnectionsServicesGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NetworkConnectionsServicesGetResponse(val *V1NetworkConnectionsServicesGetResponse) *NullableV1NetworkConnectionsServicesGetResponse {
	return &NullableV1NetworkConnectionsServicesGetResponse{value: val, isSet: true}
}

func (v NullableV1NetworkConnectionsServicesGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NetworkConnectionsServicesGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
