/*
Ledgera core API

Ledgera servers.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
	"fmt"
)

// LedgerJournalQueryType the model 'LedgerJournalQueryType'
type LedgerJournalQueryType string

// List of ledger.JournalQueryType
const (
	JournalQueryTypeDebit LedgerJournalQueryType = "DEBIT"
	JournalQueryTypeCredit LedgerJournalQueryType = "CREDIT"
)

// All allowed values of LedgerJournalQueryType enum
var AllowedLedgerJournalQueryTypeEnumValues = []LedgerJournalQueryType{
	"DEBIT",
	"CREDIT",
}

func (v *LedgerJournalQueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LedgerJournalQueryType(value)
	for _, existing := range AllowedLedgerJournalQueryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LedgerJournalQueryType", value)
}

// NewLedgerJournalQueryTypeFromValue returns a pointer to a valid LedgerJournalQueryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLedgerJournalQueryTypeFromValue(v string) (*LedgerJournalQueryType, error) {
	ev := LedgerJournalQueryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LedgerJournalQueryType: valid values are %v", v, AllowedLedgerJournalQueryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LedgerJournalQueryType) IsValid() bool {
	for _, existing := range AllowedLedgerJournalQueryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ledger.JournalQueryType value
func (v LedgerJournalQueryType) Ptr() *LedgerJournalQueryType {
	return &v
}

type NullableLedgerJournalQueryType struct {
	value *LedgerJournalQueryType
	isSet bool
}

func (v NullableLedgerJournalQueryType) Get() *LedgerJournalQueryType {
	return v.value
}

func (v *NullableLedgerJournalQueryType) Set(val *LedgerJournalQueryType) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerJournalQueryType) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerJournalQueryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerJournalQueryType(val *LedgerJournalQueryType) *NullableLedgerJournalQueryType {
	return &NullableLedgerJournalQueryType{value: val, isSet: true}
}

func (v NullableLedgerJournalQueryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerJournalQueryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
