/*
Ledgera

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the LedgerBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LedgerBalance{}

// LedgerBalance struct for LedgerBalance
type LedgerBalance struct {
	AccountId string `json:"accountId"`
	Amount float32 `json:"amount"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Currency string `json:"currency"`
	SeqNum int32 `json:"seqNum"`
	Type string `json:"type"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewLedgerBalance instantiates a new LedgerBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLedgerBalance(accountId string, amount float32, currency string, seqNum int32, type_ string) *LedgerBalance {
	this := LedgerBalance{}
	this.AccountId = accountId
	this.Amount = amount
	this.Currency = currency
	this.SeqNum = seqNum
	this.Type = type_
	return &this
}

// NewLedgerBalanceWithDefaults instantiates a new LedgerBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLedgerBalanceWithDefaults() *LedgerBalance {
	this := LedgerBalance{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *LedgerBalance) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *LedgerBalance) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *LedgerBalance) SetAccountId(v string) {
	o.AccountId = v
}

// GetAmount returns the Amount field value
func (o *LedgerBalance) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *LedgerBalance) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *LedgerBalance) SetAmount(v float32) {
	o.Amount = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LedgerBalance) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerBalance) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LedgerBalance) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *LedgerBalance) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCurrency returns the Currency field value
func (o *LedgerBalance) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *LedgerBalance) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *LedgerBalance) SetCurrency(v string) {
	o.Currency = v
}

// GetSeqNum returns the SeqNum field value
func (o *LedgerBalance) GetSeqNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SeqNum
}

// GetSeqNumOk returns a tuple with the SeqNum field value
// and a boolean to check if the value has been set.
func (o *LedgerBalance) GetSeqNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeqNum, true
}

// SetSeqNum sets field value
func (o *LedgerBalance) SetSeqNum(v int32) {
	o.SeqNum = v
}

// GetType returns the Type field value
func (o *LedgerBalance) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LedgerBalance) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LedgerBalance) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LedgerBalance) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerBalance) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LedgerBalance) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *LedgerBalance) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o LedgerBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LedgerBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountId"] = o.AccountId
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	toSerialize["currency"] = o.Currency
	toSerialize["seqNum"] = o.SeqNum
	toSerialize["type"] = o.Type
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableLedgerBalance struct {
	value *LedgerBalance
	isSet bool
}

func (v NullableLedgerBalance) Get() *LedgerBalance {
	return v.value
}

func (v *NullableLedgerBalance) Set(val *LedgerBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerBalance(val *LedgerBalance) *NullableLedgerBalance {
	return &NullableLedgerBalance{value: val, isSet: true}
}

func (v NullableLedgerBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


