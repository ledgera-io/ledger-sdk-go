/*
API ledgera

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledger-sdk-go

import (
	"encoding/json"
	"time"
)

// checks if the Balance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Balance{}

// Balance struct for Balance
type Balance struct {
	AccountId *string `json:"accountId,omitempty"`
	Type *string `json:"type,omitempty"`
	Currency *string `json:"currency,omitempty"`
	SeqNum *int32 `json:"seqNum,omitempty"`
	Amount *string `json:"amount,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewBalance instantiates a new Balance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalance() *Balance {
	this := Balance{}
	return &this
}

// NewBalanceWithDefaults instantiates a new Balance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceWithDefaults() *Balance {
	this := Balance{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Balance) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balance) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Balance) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Balance) SetAccountId(v string) {
	o.AccountId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Balance) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balance) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Balance) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Balance) SetType(v string) {
	o.Type = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Balance) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balance) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Balance) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Balance) SetCurrency(v string) {
	o.Currency = &v
}

// GetSeqNum returns the SeqNum field value if set, zero value otherwise.
func (o *Balance) GetSeqNum() int32 {
	if o == nil || IsNil(o.SeqNum) {
		var ret int32
		return ret
	}
	return *o.SeqNum
}

// GetSeqNumOk returns a tuple with the SeqNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balance) GetSeqNumOk() (*int32, bool) {
	if o == nil || IsNil(o.SeqNum) {
		return nil, false
	}
	return o.SeqNum, true
}

// HasSeqNum returns a boolean if a field has been set.
func (o *Balance) HasSeqNum() bool {
	if o != nil && !IsNil(o.SeqNum) {
		return true
	}

	return false
}

// SetSeqNum gets a reference to the given int32 and assigns it to the SeqNum field.
func (o *Balance) SetSeqNum(v int32) {
	o.SeqNum = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Balance) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balance) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Balance) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *Balance) SetAmount(v string) {
	o.Amount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Balance) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balance) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Balance) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Balance) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Balance) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balance) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Balance) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Balance) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Balance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Balance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.SeqNum) {
		toSerialize["seqNum"] = o.SeqNum
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableBalance struct {
	value *Balance
	isSet bool
}

func (v NullableBalance) Get() *Balance {
	return v.value
}

func (v *NullableBalance) Set(val *Balance) {
	v.value = val
	v.isSet = true
}

func (v NullableBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalance(val *Balance) *NullableBalance {
	return &NullableBalance{value: val, isSet: true}
}

func (v NullableBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


