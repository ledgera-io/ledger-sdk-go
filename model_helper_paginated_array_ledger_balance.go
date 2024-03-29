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

// checks if the HelperPaginatedArrayLedgerBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelperPaginatedArrayLedgerBalance{}

// HelperPaginatedArrayLedgerBalance struct for HelperPaginatedArrayLedgerBalance
type HelperPaginatedArrayLedgerBalance struct {
	Data []LedgerBalance `json:"data,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewHelperPaginatedArrayLedgerBalance instantiates a new HelperPaginatedArrayLedgerBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelperPaginatedArrayLedgerBalance() *HelperPaginatedArrayLedgerBalance {
	this := HelperPaginatedArrayLedgerBalance{}
	return &this
}

// NewHelperPaginatedArrayLedgerBalanceWithDefaults instantiates a new HelperPaginatedArrayLedgerBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelperPaginatedArrayLedgerBalanceWithDefaults() *HelperPaginatedArrayLedgerBalance {
	this := HelperPaginatedArrayLedgerBalance{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HelperPaginatedArrayLedgerBalance) GetData() []LedgerBalance {
	if o == nil || IsNil(o.Data) {
		var ret []LedgerBalance
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelperPaginatedArrayLedgerBalance) GetDataOk() ([]LedgerBalance, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HelperPaginatedArrayLedgerBalance) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []LedgerBalance and assigns it to the Data field.
func (o *HelperPaginatedArrayLedgerBalance) SetData(v []LedgerBalance) {
	o.Data = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HelperPaginatedArrayLedgerBalance) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelperPaginatedArrayLedgerBalance) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HelperPaginatedArrayLedgerBalance) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *HelperPaginatedArrayLedgerBalance) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o HelperPaginatedArrayLedgerBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelperPaginatedArrayLedgerBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableHelperPaginatedArrayLedgerBalance struct {
	value *HelperPaginatedArrayLedgerBalance
	isSet bool
}

func (v NullableHelperPaginatedArrayLedgerBalance) Get() *HelperPaginatedArrayLedgerBalance {
	return v.value
}

func (v *NullableHelperPaginatedArrayLedgerBalance) Set(val *HelperPaginatedArrayLedgerBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableHelperPaginatedArrayLedgerBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableHelperPaginatedArrayLedgerBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelperPaginatedArrayLedgerBalance(val *HelperPaginatedArrayLedgerBalance) *NullableHelperPaginatedArrayLedgerBalance {
	return &NullableHelperPaginatedArrayLedgerBalance{value: val, isSet: true}
}

func (v NullableHelperPaginatedArrayLedgerBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelperPaginatedArrayLedgerBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


