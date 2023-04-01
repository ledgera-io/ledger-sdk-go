/*
ledgera

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the HelperPaginatedArrayLedgerCurrency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelperPaginatedArrayLedgerCurrency{}

// HelperPaginatedArrayLedgerCurrency struct for HelperPaginatedArrayLedgerCurrency
type HelperPaginatedArrayLedgerCurrency struct {
	Data []LedgerCurrency `json:"data,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewHelperPaginatedArrayLedgerCurrency instantiates a new HelperPaginatedArrayLedgerCurrency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelperPaginatedArrayLedgerCurrency() *HelperPaginatedArrayLedgerCurrency {
	this := HelperPaginatedArrayLedgerCurrency{}
	return &this
}

// NewHelperPaginatedArrayLedgerCurrencyWithDefaults instantiates a new HelperPaginatedArrayLedgerCurrency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelperPaginatedArrayLedgerCurrencyWithDefaults() *HelperPaginatedArrayLedgerCurrency {
	this := HelperPaginatedArrayLedgerCurrency{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HelperPaginatedArrayLedgerCurrency) GetData() []LedgerCurrency {
	if o == nil || IsNil(o.Data) {
		var ret []LedgerCurrency
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelperPaginatedArrayLedgerCurrency) GetDataOk() ([]LedgerCurrency, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HelperPaginatedArrayLedgerCurrency) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []LedgerCurrency and assigns it to the Data field.
func (o *HelperPaginatedArrayLedgerCurrency) SetData(v []LedgerCurrency) {
	o.Data = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HelperPaginatedArrayLedgerCurrency) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelperPaginatedArrayLedgerCurrency) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HelperPaginatedArrayLedgerCurrency) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *HelperPaginatedArrayLedgerCurrency) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o HelperPaginatedArrayLedgerCurrency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelperPaginatedArrayLedgerCurrency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableHelperPaginatedArrayLedgerCurrency struct {
	value *HelperPaginatedArrayLedgerCurrency
	isSet bool
}

func (v NullableHelperPaginatedArrayLedgerCurrency) Get() *HelperPaginatedArrayLedgerCurrency {
	return v.value
}

func (v *NullableHelperPaginatedArrayLedgerCurrency) Set(val *HelperPaginatedArrayLedgerCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullableHelperPaginatedArrayLedgerCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableHelperPaginatedArrayLedgerCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelperPaginatedArrayLedgerCurrency(val *HelperPaginatedArrayLedgerCurrency) *NullableHelperPaginatedArrayLedgerCurrency {
	return &NullableHelperPaginatedArrayLedgerCurrency{value: val, isSet: true}
}

func (v NullableHelperPaginatedArrayLedgerCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelperPaginatedArrayLedgerCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


