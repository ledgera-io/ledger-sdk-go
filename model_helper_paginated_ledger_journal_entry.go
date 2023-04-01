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

// checks if the HelperPaginatedLedgerJournalEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelperPaginatedLedgerJournalEntry{}

// HelperPaginatedLedgerJournalEntry struct for HelperPaginatedLedgerJournalEntry
type HelperPaginatedLedgerJournalEntry struct {
	Data *LedgerJournalEntry `json:"data,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewHelperPaginatedLedgerJournalEntry instantiates a new HelperPaginatedLedgerJournalEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelperPaginatedLedgerJournalEntry() *HelperPaginatedLedgerJournalEntry {
	this := HelperPaginatedLedgerJournalEntry{}
	return &this
}

// NewHelperPaginatedLedgerJournalEntryWithDefaults instantiates a new HelperPaginatedLedgerJournalEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelperPaginatedLedgerJournalEntryWithDefaults() *HelperPaginatedLedgerJournalEntry {
	this := HelperPaginatedLedgerJournalEntry{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HelperPaginatedLedgerJournalEntry) GetData() LedgerJournalEntry {
	if o == nil || IsNil(o.Data) {
		var ret LedgerJournalEntry
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelperPaginatedLedgerJournalEntry) GetDataOk() (*LedgerJournalEntry, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HelperPaginatedLedgerJournalEntry) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LedgerJournalEntry and assigns it to the Data field.
func (o *HelperPaginatedLedgerJournalEntry) SetData(v LedgerJournalEntry) {
	o.Data = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HelperPaginatedLedgerJournalEntry) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelperPaginatedLedgerJournalEntry) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HelperPaginatedLedgerJournalEntry) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *HelperPaginatedLedgerJournalEntry) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o HelperPaginatedLedgerJournalEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelperPaginatedLedgerJournalEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableHelperPaginatedLedgerJournalEntry struct {
	value *HelperPaginatedLedgerJournalEntry
	isSet bool
}

func (v NullableHelperPaginatedLedgerJournalEntry) Get() *HelperPaginatedLedgerJournalEntry {
	return v.value
}

func (v *NullableHelperPaginatedLedgerJournalEntry) Set(val *HelperPaginatedLedgerJournalEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableHelperPaginatedLedgerJournalEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableHelperPaginatedLedgerJournalEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelperPaginatedLedgerJournalEntry(val *HelperPaginatedLedgerJournalEntry) *NullableHelperPaginatedLedgerJournalEntry {
	return &NullableHelperPaginatedLedgerJournalEntry{value: val, isSet: true}
}

func (v NullableHelperPaginatedLedgerJournalEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelperPaginatedLedgerJournalEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


