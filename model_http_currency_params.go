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

// checks if the HttpCurrencyParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpCurrencyParams{}

// HttpCurrencyParams struct for HttpCurrencyParams
type HttpCurrencyParams struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// NewHttpCurrencyParams instantiates a new HttpCurrencyParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpCurrencyParams(code string, name string) *HttpCurrencyParams {
	this := HttpCurrencyParams{}
	this.Code = code
	this.Name = name
	return &this
}

// NewHttpCurrencyParamsWithDefaults instantiates a new HttpCurrencyParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpCurrencyParamsWithDefaults() *HttpCurrencyParams {
	this := HttpCurrencyParams{}
	return &this
}

// GetCode returns the Code field value
func (o *HttpCurrencyParams) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *HttpCurrencyParams) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *HttpCurrencyParams) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *HttpCurrencyParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HttpCurrencyParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HttpCurrencyParams) SetName(v string) {
	o.Name = v
}

func (o HttpCurrencyParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpCurrencyParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableHttpCurrencyParams struct {
	value *HttpCurrencyParams
	isSet bool
}

func (v NullableHttpCurrencyParams) Get() *HttpCurrencyParams {
	return v.value
}

func (v *NullableHttpCurrencyParams) Set(val *HttpCurrencyParams) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpCurrencyParams) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpCurrencyParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpCurrencyParams(val *HttpCurrencyParams) *NullableHttpCurrencyParams {
	return &NullableHttpCurrencyParams{value: val, isSet: true}
}

func (v NullableHttpCurrencyParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpCurrencyParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


