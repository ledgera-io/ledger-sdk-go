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

// checks if the LedgerTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LedgerTransaction{}

// LedgerTransaction struct for LedgerTransaction
type LedgerTransaction struct {
	ConciliationId *string `json:"conciliationId,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Id *string `json:"id,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	TransactionType *string `json:"transactionType,omitempty"`
}

// NewLedgerTransaction instantiates a new LedgerTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLedgerTransaction() *LedgerTransaction {
	this := LedgerTransaction{}
	return &this
}

// NewLedgerTransactionWithDefaults instantiates a new LedgerTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLedgerTransactionWithDefaults() *LedgerTransaction {
	this := LedgerTransaction{}
	return &this
}

// GetConciliationId returns the ConciliationId field value if set, zero value otherwise.
func (o *LedgerTransaction) GetConciliationId() string {
	if o == nil || IsNil(o.ConciliationId) {
		var ret string
		return ret
	}
	return *o.ConciliationId
}

// GetConciliationIdOk returns a tuple with the ConciliationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransaction) GetConciliationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConciliationId) {
		return nil, false
	}
	return o.ConciliationId, true
}

// HasConciliationId returns a boolean if a field has been set.
func (o *LedgerTransaction) HasConciliationId() bool {
	if o != nil && !IsNil(o.ConciliationId) {
		return true
	}

	return false
}

// SetConciliationId gets a reference to the given string and assigns it to the ConciliationId field.
func (o *LedgerTransaction) SetConciliationId(v string) {
	o.ConciliationId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LedgerTransaction) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransaction) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LedgerTransaction) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *LedgerTransaction) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LedgerTransaction) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransaction) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LedgerTransaction) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LedgerTransaction) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LedgerTransaction) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransaction) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LedgerTransaction) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LedgerTransaction) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *LedgerTransaction) GetTransactionType() string {
	if o == nil || IsNil(o.TransactionType) {
		var ret string
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransaction) GetTransactionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionType) {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *LedgerTransaction) HasTransactionType() bool {
	if o != nil && !IsNil(o.TransactionType) {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given string and assigns it to the TransactionType field.
func (o *LedgerTransaction) SetTransactionType(v string) {
	o.TransactionType = &v
}

func (o LedgerTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LedgerTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConciliationId) {
		toSerialize["conciliationId"] = o.ConciliationId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.TransactionType) {
		toSerialize["transactionType"] = o.TransactionType
	}
	return toSerialize, nil
}

type NullableLedgerTransaction struct {
	value *LedgerTransaction
	isSet bool
}

func (v NullableLedgerTransaction) Get() *LedgerTransaction {
	return v.value
}

func (v *NullableLedgerTransaction) Set(val *LedgerTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerTransaction(val *LedgerTransaction) *NullableLedgerTransaction {
	return &NullableLedgerTransaction{value: val, isSet: true}
}

func (v NullableLedgerTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


