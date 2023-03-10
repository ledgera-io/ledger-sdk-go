/*
API ledgera

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledger-sdk

import (
	"encoding/json"
	"fmt"
)

// AccountType the model 'AccountType'
type AccountType string

// List of accountType
const (
	LIABILITY AccountType = "liability"
	ASSET AccountType = "asset"
)

// All allowed values of AccountType enum
var AllowedAccountTypeEnumValues = []AccountType{
	"liability",
	"asset",
}

func (v *AccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountType(value)
	for _, existing := range AllowedAccountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountType", value)
}

// NewAccountTypeFromValue returns a pointer to a valid AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountTypeFromValue(v string) (*AccountType, error) {
	ev := AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountType: valid values are %v", v, AllowedAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountType) IsValid() bool {
	for _, existing := range AllowedAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to accountType value
func (v AccountType) Ptr() *AccountType {
	return &v
}

type NullableAccountType struct {
	value *AccountType
	isSet bool
}

func (v NullableAccountType) Get() *AccountType {
	return v.value
}

func (v *NullableAccountType) Set(val *AccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountType(val *AccountType) *NullableAccountType {
	return &NullableAccountType{value: val, isSet: true}
}

func (v NullableAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

