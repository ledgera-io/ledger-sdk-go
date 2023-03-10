/*
API ledgera

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

ledgersdkgo

import (
	"encoding/json"
)

// checks if the NotFound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotFound{}

// NotFound struct for NotFound
type NotFound struct {
	Message *string `json:"message,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

// NewNotFound instantiates a new NotFound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotFound() *NotFound {
	this := NotFound{}
	return &this
}

// NewNotFoundWithDefaults instantiates a new NotFound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotFoundWithDefaults() *NotFound {
	this := NotFound{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NotFound) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotFound) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NotFound) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NotFound) SetMessage(v string) {
	o.Message = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *NotFound) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotFound) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *NotFound) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *NotFound) SetResource(v string) {
	o.Resource = &v
}

func (o NotFound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotFound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	return toSerialize, nil
}

type NullableNotFound struct {
	value *NotFound
	isSet bool
}

func (v NullableNotFound) Get() *NotFound {
	return v.value
}

func (v *NullableNotFound) Set(val *NotFound) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFound) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFound(val *NotFound) *NullableNotFound {
	return &NullableNotFound{value: val, isSet: true}
}

func (v NullableNotFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


