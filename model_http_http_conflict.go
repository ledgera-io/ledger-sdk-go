/*
ledgera

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the HttpHttpConflict type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpHttpConflict{}

// HttpHttpConflict struct for HttpHttpConflict
type HttpHttpConflict struct {
	Data *HttpConflictData `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewHttpHttpConflict instantiates a new HttpHttpConflict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpHttpConflict() *HttpHttpConflict {
	this := HttpHttpConflict{}
	return &this
}

// NewHttpHttpConflictWithDefaults instantiates a new HttpHttpConflict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpHttpConflictWithDefaults() *HttpHttpConflict {
	this := HttpHttpConflict{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HttpHttpConflict) GetData() HttpConflictData {
	if o == nil || IsNil(o.Data) {
		var ret HttpConflictData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpHttpConflict) GetDataOk() (*HttpConflictData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HttpHttpConflict) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given HttpConflictData and assigns it to the Data field.
func (o *HttpHttpConflict) SetData(v HttpConflictData) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HttpHttpConflict) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpHttpConflict) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HttpHttpConflict) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HttpHttpConflict) SetMessage(v string) {
	o.Message = &v
}

func (o HttpHttpConflict) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpHttpConflict) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableHttpHttpConflict struct {
	value *HttpHttpConflict
	isSet bool
}

func (v NullableHttpHttpConflict) Get() *HttpHttpConflict {
	return v.value
}

func (v *NullableHttpHttpConflict) Set(val *HttpHttpConflict) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpHttpConflict) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpHttpConflict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpHttpConflict(val *HttpHttpConflict) *NullableHttpHttpConflict {
	return &NullableHttpHttpConflict{value: val, isSet: true}
}

func (v NullableHttpHttpConflict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpHttpConflict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


