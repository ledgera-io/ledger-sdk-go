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

// checks if the LedgerTransactionProcessed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LedgerTransactionProcessed{}

// LedgerTransactionProcessed struct for LedgerTransactionProcessed
type LedgerTransactionProcessed struct {
	CreatedAt *string `json:"createdAt,omitempty"`
	Entries []LedgerJournalEntry `json:"entries,omitempty"`
	Id *string `json:"id,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewLedgerTransactionProcessed instantiates a new LedgerTransactionProcessed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLedgerTransactionProcessed() *LedgerTransactionProcessed {
	this := LedgerTransactionProcessed{}
	return &this
}

// NewLedgerTransactionProcessedWithDefaults instantiates a new LedgerTransactionProcessed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLedgerTransactionProcessedWithDefaults() *LedgerTransactionProcessed {
	this := LedgerTransactionProcessed{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LedgerTransactionProcessed) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionProcessed) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LedgerTransactionProcessed) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *LedgerTransactionProcessed) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *LedgerTransactionProcessed) GetEntries() []LedgerJournalEntry {
	if o == nil || IsNil(o.Entries) {
		var ret []LedgerJournalEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionProcessed) GetEntriesOk() ([]LedgerJournalEntry, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *LedgerTransactionProcessed) HasEntries() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []LedgerJournalEntry and assigns it to the Entries field.
func (o *LedgerTransactionProcessed) SetEntries(v []LedgerJournalEntry) {
	o.Entries = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LedgerTransactionProcessed) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionProcessed) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LedgerTransactionProcessed) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LedgerTransactionProcessed) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LedgerTransactionProcessed) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionProcessed) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LedgerTransactionProcessed) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LedgerTransactionProcessed) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o LedgerTransactionProcessed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LedgerTransactionProcessed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableLedgerTransactionProcessed struct {
	value *LedgerTransactionProcessed
	isSet bool
}

func (v NullableLedgerTransactionProcessed) Get() *LedgerTransactionProcessed {
	return v.value
}

func (v *NullableLedgerTransactionProcessed) Set(val *LedgerTransactionProcessed) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerTransactionProcessed) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerTransactionProcessed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerTransactionProcessed(val *LedgerTransactionProcessed) *NullableLedgerTransactionProcessed {
	return &NullableLedgerTransactionProcessed{value: val, isSet: true}
}

func (v NullableLedgerTransactionProcessed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerTransactionProcessed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


