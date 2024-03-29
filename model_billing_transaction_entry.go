/*
Ledgera core API

Ledgera servers.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the BillingTransactionEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingTransactionEntry{}

// BillingTransactionEntry struct for BillingTransactionEntry
type BillingTransactionEntry struct {
	Amount *string `json:"amount,omitempty"`
	AmountType *BillingBillingTransactionAmountType `json:"amountType,omitempty"`
	EntryType *string `json:"entryType,omitempty"`
	TransactionType *string `json:"transactionType,omitempty"`
}

// NewBillingTransactionEntry instantiates a new BillingTransactionEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingTransactionEntry() *BillingTransactionEntry {
	this := BillingTransactionEntry{}
	return &this
}

// NewBillingTransactionEntryWithDefaults instantiates a new BillingTransactionEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingTransactionEntryWithDefaults() *BillingTransactionEntry {
	this := BillingTransactionEntry{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BillingTransactionEntry) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTransactionEntry) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BillingTransactionEntry) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *BillingTransactionEntry) SetAmount(v string) {
	o.Amount = &v
}

// GetAmountType returns the AmountType field value if set, zero value otherwise.
func (o *BillingTransactionEntry) GetAmountType() BillingBillingTransactionAmountType {
	if o == nil || IsNil(o.AmountType) {
		var ret BillingBillingTransactionAmountType
		return ret
	}
	return *o.AmountType
}

// GetAmountTypeOk returns a tuple with the AmountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTransactionEntry) GetAmountTypeOk() (*BillingBillingTransactionAmountType, bool) {
	if o == nil || IsNil(o.AmountType) {
		return nil, false
	}
	return o.AmountType, true
}

// HasAmountType returns a boolean if a field has been set.
func (o *BillingTransactionEntry) HasAmountType() bool {
	if o != nil && !IsNil(o.AmountType) {
		return true
	}

	return false
}

// SetAmountType gets a reference to the given BillingBillingTransactionAmountType and assigns it to the AmountType field.
func (o *BillingTransactionEntry) SetAmountType(v BillingBillingTransactionAmountType) {
	o.AmountType = &v
}

// GetEntryType returns the EntryType field value if set, zero value otherwise.
func (o *BillingTransactionEntry) GetEntryType() string {
	if o == nil || IsNil(o.EntryType) {
		var ret string
		return ret
	}
	return *o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTransactionEntry) GetEntryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntryType) {
		return nil, false
	}
	return o.EntryType, true
}

// HasEntryType returns a boolean if a field has been set.
func (o *BillingTransactionEntry) HasEntryType() bool {
	if o != nil && !IsNil(o.EntryType) {
		return true
	}

	return false
}

// SetEntryType gets a reference to the given string and assigns it to the EntryType field.
func (o *BillingTransactionEntry) SetEntryType(v string) {
	o.EntryType = &v
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *BillingTransactionEntry) GetTransactionType() string {
	if o == nil || IsNil(o.TransactionType) {
		var ret string
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingTransactionEntry) GetTransactionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionType) {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *BillingTransactionEntry) HasTransactionType() bool {
	if o != nil && !IsNil(o.TransactionType) {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given string and assigns it to the TransactionType field.
func (o *BillingTransactionEntry) SetTransactionType(v string) {
	o.TransactionType = &v
}

func (o BillingTransactionEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingTransactionEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AmountType) {
		toSerialize["amountType"] = o.AmountType
	}
	if !IsNil(o.EntryType) {
		toSerialize["entryType"] = o.EntryType
	}
	if !IsNil(o.TransactionType) {
		toSerialize["transactionType"] = o.TransactionType
	}
	return toSerialize, nil
}

type NullableBillingTransactionEntry struct {
	value *BillingTransactionEntry
	isSet bool
}

func (v NullableBillingTransactionEntry) Get() *BillingTransactionEntry {
	return v.value
}

func (v *NullableBillingTransactionEntry) Set(val *BillingTransactionEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingTransactionEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingTransactionEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingTransactionEntry(val *BillingTransactionEntry) *NullableBillingTransactionEntry {
	return &NullableBillingTransactionEntry{value: val, isSet: true}
}

func (v NullableBillingTransactionEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingTransactionEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


