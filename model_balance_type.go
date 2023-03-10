/*
API ledgera

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// BalanceType the model 'BalanceType'
type BalanceType string

// List of balanceType
const (
	AVAILABLE BalanceType = "available"
	BLOCKED BalanceType = "blocked"
	PENDING BalanceType = "pending"
)

// All allowed values of BalanceType enum
var AllowedBalanceTypeEnumValues = []BalanceType{
	"available",
	"blocked",
	"pending",
}

func (v *BalanceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BalanceType(value)
	for _, existing := range AllowedBalanceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BalanceType", value)
}

// NewBalanceTypeFromValue returns a pointer to a valid BalanceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBalanceTypeFromValue(v string) (*BalanceType, error) {
	ev := BalanceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BalanceType: valid values are %v", v, AllowedBalanceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BalanceType) IsValid() bool {
	for _, existing := range AllowedBalanceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to balanceType value
func (v BalanceType) Ptr() *BalanceType {
	return &v
}

type NullableBalanceType struct {
	value *BalanceType
	isSet bool
}

func (v NullableBalanceType) Get() *BalanceType {
	return v.value
}

func (v *NullableBalanceType) Set(val *BalanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceType(val *BalanceType) *NullableBalanceType {
	return &NullableBalanceType{value: val, isSet: true}
}

func (v NullableBalanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

