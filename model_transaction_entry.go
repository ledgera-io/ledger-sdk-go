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

// checks if the TransactionEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionEntry{}

// TransactionEntry struct for TransactionEntry
type TransactionEntry struct {
	TransactionType string `json:"transactionType"`
	EntryType string `json:"entryType"`
}

// NewTransactionEntry instantiates a new TransactionEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionEntry(transactionType string, entryType string) *TransactionEntry {
	this := TransactionEntry{}
	this.TransactionType = transactionType
	this.EntryType = entryType
	return &this
}

// NewTransactionEntryWithDefaults instantiates a new TransactionEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionEntryWithDefaults() *TransactionEntry {
	this := TransactionEntry{}
	return &this
}

// GetTransactionType returns the TransactionType field value
func (o *TransactionEntry) GetTransactionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value
// and a boolean to check if the value has been set.
func (o *TransactionEntry) GetTransactionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionType, true
}

// SetTransactionType sets field value
func (o *TransactionEntry) SetTransactionType(v string) {
	o.TransactionType = v
}

// GetEntryType returns the EntryType field value
func (o *TransactionEntry) GetEntryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value
// and a boolean to check if the value has been set.
func (o *TransactionEntry) GetEntryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryType, true
}

// SetEntryType sets field value
func (o *TransactionEntry) SetEntryType(v string) {
	o.EntryType = v
}

func (o TransactionEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactionType"] = o.TransactionType
	toSerialize["entryType"] = o.EntryType
	return toSerialize, nil
}

type NullableTransactionEntry struct {
	value *TransactionEntry
	isSet bool
}

func (v NullableTransactionEntry) Get() *TransactionEntry {
	return v.value
}

func (v *NullableTransactionEntry) Set(val *TransactionEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionEntry(val *TransactionEntry) *NullableTransactionEntry {
	return &NullableTransactionEntry{value: val, isSet: true}
}

func (v NullableTransactionEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


