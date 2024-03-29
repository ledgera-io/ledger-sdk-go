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

// BillingBillingPlanStatus the model 'BillingBillingPlanStatus'
type BillingBillingPlanStatus string

// List of billing.BillingPlanStatus
const (
	BillingPlanActive BillingBillingPlanStatus = "ACTIVE"
	BillingPlanInactive BillingBillingPlanStatus = "INACTIVE"
)

// All allowed values of BillingBillingPlanStatus enum
var AllowedBillingBillingPlanStatusEnumValues = []BillingBillingPlanStatus{
	"ACTIVE",
	"INACTIVE",
}

func (v *BillingBillingPlanStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BillingBillingPlanStatus(value)
	for _, existing := range AllowedBillingBillingPlanStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BillingBillingPlanStatus", value)
}

// NewBillingBillingPlanStatusFromValue returns a pointer to a valid BillingBillingPlanStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBillingBillingPlanStatusFromValue(v string) (*BillingBillingPlanStatus, error) {
	ev := BillingBillingPlanStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BillingBillingPlanStatus: valid values are %v", v, AllowedBillingBillingPlanStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BillingBillingPlanStatus) IsValid() bool {
	for _, existing := range AllowedBillingBillingPlanStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to billing.BillingPlanStatus value
func (v BillingBillingPlanStatus) Ptr() *BillingBillingPlanStatus {
	return &v
}

type NullableBillingBillingPlanStatus struct {
	value *BillingBillingPlanStatus
	isSet bool
}

func (v NullableBillingBillingPlanStatus) Get() *BillingBillingPlanStatus {
	return v.value
}

func (v *NullableBillingBillingPlanStatus) Set(val *BillingBillingPlanStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingBillingPlanStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingBillingPlanStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingBillingPlanStatus(val *BillingBillingPlanStatus) *NullableBillingBillingPlanStatus {
	return &NullableBillingBillingPlanStatus{value: val, isSet: true}
}

func (v NullableBillingBillingPlanStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingBillingPlanStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

