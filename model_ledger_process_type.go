/*
Ledgera API

teste

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
	"fmt"
)

// LedgerProcessType the model 'LedgerProcessType'
type LedgerProcessType string

// List of ledger.ProcessType
const (
	ProcessTypeExecution LedgerProcessType = "execution"
	ProcessTypeAuthorization LedgerProcessType = "authorization"
	ProcessTypeConfirmation LedgerProcessType = "confirmation"
	ProcessTypeReversal LedgerProcessType = "reversal"
)

// All allowed values of LedgerProcessType enum
var AllowedLedgerProcessTypeEnumValues = []LedgerProcessType{
	"execution",
	"authorization",
	"confirmation",
	"reversal",
}

func (v *LedgerProcessType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LedgerProcessType(value)
	for _, existing := range AllowedLedgerProcessTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LedgerProcessType", value)
}

// NewLedgerProcessTypeFromValue returns a pointer to a valid LedgerProcessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLedgerProcessTypeFromValue(v string) (*LedgerProcessType, error) {
	ev := LedgerProcessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LedgerProcessType: valid values are %v", v, AllowedLedgerProcessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LedgerProcessType) IsValid() bool {
	for _, existing := range AllowedLedgerProcessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ledger.ProcessType value
func (v LedgerProcessType) Ptr() *LedgerProcessType {
	return &v
}

type NullableLedgerProcessType struct {
	value *LedgerProcessType
	isSet bool
}

func (v NullableLedgerProcessType) Get() *LedgerProcessType {
	return v.value
}

func (v *NullableLedgerProcessType) Set(val *LedgerProcessType) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerProcessType) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerProcessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerProcessType(val *LedgerProcessType) *NullableLedgerProcessType {
	return &NullableLedgerProcessType{value: val, isSet: true}
}

func (v NullableLedgerProcessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerProcessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
