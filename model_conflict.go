/*
API ledgera

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Conflict type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Conflict{}

// Conflict struct for Conflict
type Conflict struct {
	TransactionType *string `json:"transactionType,omitempty"`
	TransactionProcess *string `json:"transactionProcess,omitempty"`
	ExpectedAmount *string `json:"expectedAmount,omitempty"`
	AccountId *string `json:"accountId,omitempty"`
}

// NewConflict instantiates a new Conflict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConflict() *Conflict {
	this := Conflict{}
	return &this
}

// NewConflictWithDefaults instantiates a new Conflict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConflictWithDefaults() *Conflict {
	this := Conflict{}
	return &this
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *Conflict) GetTransactionType() string {
	if o == nil || IsNil(o.TransactionType) {
		var ret string
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conflict) GetTransactionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionType) {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *Conflict) HasTransactionType() bool {
	if o != nil && !IsNil(o.TransactionType) {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given string and assigns it to the TransactionType field.
func (o *Conflict) SetTransactionType(v string) {
	o.TransactionType = &v
}

// GetTransactionProcess returns the TransactionProcess field value if set, zero value otherwise.
func (o *Conflict) GetTransactionProcess() string {
	if o == nil || IsNil(o.TransactionProcess) {
		var ret string
		return ret
	}
	return *o.TransactionProcess
}

// GetTransactionProcessOk returns a tuple with the TransactionProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conflict) GetTransactionProcessOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionProcess) {
		return nil, false
	}
	return o.TransactionProcess, true
}

// HasTransactionProcess returns a boolean if a field has been set.
func (o *Conflict) HasTransactionProcess() bool {
	if o != nil && !IsNil(o.TransactionProcess) {
		return true
	}

	return false
}

// SetTransactionProcess gets a reference to the given string and assigns it to the TransactionProcess field.
func (o *Conflict) SetTransactionProcess(v string) {
	o.TransactionProcess = &v
}

// GetExpectedAmount returns the ExpectedAmount field value if set, zero value otherwise.
func (o *Conflict) GetExpectedAmount() string {
	if o == nil || IsNil(o.ExpectedAmount) {
		var ret string
		return ret
	}
	return *o.ExpectedAmount
}

// GetExpectedAmountOk returns a tuple with the ExpectedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conflict) GetExpectedAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ExpectedAmount) {
		return nil, false
	}
	return o.ExpectedAmount, true
}

// HasExpectedAmount returns a boolean if a field has been set.
func (o *Conflict) HasExpectedAmount() bool {
	if o != nil && !IsNil(o.ExpectedAmount) {
		return true
	}

	return false
}

// SetExpectedAmount gets a reference to the given string and assigns it to the ExpectedAmount field.
func (o *Conflict) SetExpectedAmount(v string) {
	o.ExpectedAmount = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Conflict) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conflict) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Conflict) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Conflict) SetAccountId(v string) {
	o.AccountId = &v
}

func (o Conflict) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Conflict) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionType) {
		toSerialize["transactionType"] = o.TransactionType
	}
	if !IsNil(o.TransactionProcess) {
		toSerialize["transactionProcess"] = o.TransactionProcess
	}
	if !IsNil(o.ExpectedAmount) {
		toSerialize["expectedAmount"] = o.ExpectedAmount
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	return toSerialize, nil
}

type NullableConflict struct {
	value *Conflict
	isSet bool
}

func (v NullableConflict) Get() *Conflict {
	return v.value
}

func (v *NullableConflict) Set(val *Conflict) {
	v.value = val
	v.isSet = true
}

func (v NullableConflict) IsSet() bool {
	return v.isSet
}

func (v *NullableConflict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConflict(val *Conflict) *NullableConflict {
	return &NullableConflict{value: val, isSet: true}
}

func (v NullableConflict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConflict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


