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

// checks if the V1ErrorResponseErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ErrorResponseErrorsInner{}

// V1ErrorResponseErrorsInner struct for V1ErrorResponseErrorsInner
type V1ErrorResponseErrorsInner struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Type    *string `json:"type,omitempty"`
}

// NewV1ErrorResponseErrorsInner instantiates a new V1ErrorResponseErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ErrorResponseErrorsInner() *V1ErrorResponseErrorsInner {
	this := V1ErrorResponseErrorsInner{}
	return &this
}

// NewV1ErrorResponseErrorsInnerWithDefaults instantiates a new V1ErrorResponseErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ErrorResponseErrorsInnerWithDefaults() *V1ErrorResponseErrorsInner {
	this := V1ErrorResponseErrorsInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *V1ErrorResponseErrorsInner) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ErrorResponseErrorsInner) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *V1ErrorResponseErrorsInner) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *V1ErrorResponseErrorsInner) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1ErrorResponseErrorsInner) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ErrorResponseErrorsInner) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1ErrorResponseErrorsInner) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1ErrorResponseErrorsInner) SetMessage(v string) {
	o.Message = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1ErrorResponseErrorsInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ErrorResponseErrorsInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1ErrorResponseErrorsInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1ErrorResponseErrorsInner) SetType(v string) {
	o.Type = &v
}

func (o V1ErrorResponseErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ErrorResponseErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableV1ErrorResponseErrorsInner struct {
	value *V1ErrorResponseErrorsInner
	isSet bool
}

func (v NullableV1ErrorResponseErrorsInner) Get() *V1ErrorResponseErrorsInner {
	return v.value
}

func (v *NullableV1ErrorResponseErrorsInner) Set(val *V1ErrorResponseErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ErrorResponseErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ErrorResponseErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ErrorResponseErrorsInner(val *V1ErrorResponseErrorsInner) *NullableV1ErrorResponseErrorsInner {
	return &NullableV1ErrorResponseErrorsInner{value: val, isSet: true}
}

func (v NullableV1ErrorResponseErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ErrorResponseErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
