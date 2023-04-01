/*
Ledgera

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
	"fmt"
)

// LedgerMessageType the model 'LedgerMessageType'
type LedgerMessageType string

// List of ledger.MessageType
const (
	MessageTypeSingle LedgerMessageType = "single"
	MessageTypeDual LedgerMessageType = "dual"
)

// All allowed values of LedgerMessageType enum
var AllowedLedgerMessageTypeEnumValues = []LedgerMessageType{
	"single",
	"dual",
}

func (v *LedgerMessageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LedgerMessageType(value)
	for _, existing := range AllowedLedgerMessageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LedgerMessageType", value)
}

// NewLedgerMessageTypeFromValue returns a pointer to a valid LedgerMessageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLedgerMessageTypeFromValue(v string) (*LedgerMessageType, error) {
	ev := LedgerMessageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LedgerMessageType: valid values are %v", v, AllowedLedgerMessageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LedgerMessageType) IsValid() bool {
	for _, existing := range AllowedLedgerMessageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ledger.MessageType value
func (v LedgerMessageType) Ptr() *LedgerMessageType {
	return &v
}

type NullableLedgerMessageType struct {
	value *LedgerMessageType
	isSet bool
}

func (v NullableLedgerMessageType) Get() *LedgerMessageType {
	return v.value
}

func (v *NullableLedgerMessageType) Set(val *LedgerMessageType) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerMessageType) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerMessageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerMessageType(val *LedgerMessageType) *NullableLedgerMessageType {
	return &NullableLedgerMessageType{value: val, isSet: true}
}

func (v NullableLedgerMessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerMessageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

