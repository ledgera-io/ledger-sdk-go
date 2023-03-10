/*
API ledgera

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the AccountPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountPostRequest{}

// AccountPostRequest struct for AccountPostRequest
type AccountPostRequest struct {
	Type AccountType `json:"type"`
	Label string `json:"label"`
}

// NewAccountPostRequest instantiates a new AccountPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountPostRequest(type_ AccountType, label string) *AccountPostRequest {
	this := AccountPostRequest{}
	this.Type = type_
	this.Label = label
	return &this
}

// NewAccountPostRequestWithDefaults instantiates a new AccountPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountPostRequestWithDefaults() *AccountPostRequest {
	this := AccountPostRequest{}
	return &this
}

// GetType returns the Type field value
func (o *AccountPostRequest) GetType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountPostRequest) GetTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountPostRequest) SetType(v AccountType) {
	o.Type = v
}

// GetLabel returns the Label field value
func (o *AccountPostRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *AccountPostRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *AccountPostRequest) SetLabel(v string) {
	o.Label = v
}

func (o AccountPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

type NullableAccountPostRequest struct {
	value *AccountPostRequest
	isSet bool
}

func (v NullableAccountPostRequest) Get() *AccountPostRequest {
	return v.value
}

func (v *NullableAccountPostRequest) Set(val *AccountPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountPostRequest(val *AccountPostRequest) *NullableAccountPostRequest {
	return &NullableAccountPostRequest{value: val, isSet: true}
}

func (v NullableAccountPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


