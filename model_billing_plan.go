/*
API ledgera

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the BillingPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingPlan{}

// BillingPlan struct for BillingPlan
type BillingPlan struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	TransactionType *string `json:"transactionType,omitempty"`
	EntryType *string `json:"entryType,omitempty"`
	AmountType *string `json:"amountType,omitempty"`
	Amount *string `json:"amount,omitempty"`
}

// NewBillingPlan instantiates a new BillingPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPlan() *BillingPlan {
	this := BillingPlan{}
	return &this
}

// NewBillingPlanWithDefaults instantiates a new BillingPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPlanWithDefaults() *BillingPlan {
	this := BillingPlan{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillingPlan) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlan) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillingPlan) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BillingPlan) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillingPlan) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlan) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BillingPlan) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillingPlan) SetName(v string) {
	o.Name = &v
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *BillingPlan) GetTransactionType() string {
	if o == nil || IsNil(o.TransactionType) {
		var ret string
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlan) GetTransactionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionType) {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *BillingPlan) HasTransactionType() bool {
	if o != nil && !IsNil(o.TransactionType) {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given string and assigns it to the TransactionType field.
func (o *BillingPlan) SetTransactionType(v string) {
	o.TransactionType = &v
}

// GetEntryType returns the EntryType field value if set, zero value otherwise.
func (o *BillingPlan) GetEntryType() string {
	if o == nil || IsNil(o.EntryType) {
		var ret string
		return ret
	}
	return *o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlan) GetEntryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntryType) {
		return nil, false
	}
	return o.EntryType, true
}

// HasEntryType returns a boolean if a field has been set.
func (o *BillingPlan) HasEntryType() bool {
	if o != nil && !IsNil(o.EntryType) {
		return true
	}

	return false
}

// SetEntryType gets a reference to the given string and assigns it to the EntryType field.
func (o *BillingPlan) SetEntryType(v string) {
	o.EntryType = &v
}

// GetAmountType returns the AmountType field value if set, zero value otherwise.
func (o *BillingPlan) GetAmountType() string {
	if o == nil || IsNil(o.AmountType) {
		var ret string
		return ret
	}
	return *o.AmountType
}

// GetAmountTypeOk returns a tuple with the AmountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlan) GetAmountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AmountType) {
		return nil, false
	}
	return o.AmountType, true
}

// HasAmountType returns a boolean if a field has been set.
func (o *BillingPlan) HasAmountType() bool {
	if o != nil && !IsNil(o.AmountType) {
		return true
	}

	return false
}

// SetAmountType gets a reference to the given string and assigns it to the AmountType field.
func (o *BillingPlan) SetAmountType(v string) {
	o.AmountType = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BillingPlan) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlan) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BillingPlan) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *BillingPlan) SetAmount(v string) {
	o.Amount = &v
}

func (o BillingPlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TransactionType) {
		toSerialize["transactionType"] = o.TransactionType
	}
	if !IsNil(o.EntryType) {
		toSerialize["entryType"] = o.EntryType
	}
	if !IsNil(o.AmountType) {
		toSerialize["amountType"] = o.AmountType
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableBillingPlan struct {
	value *BillingPlan
	isSet bool
}

func (v NullableBillingPlan) Get() *BillingPlan {
	return v.value
}

func (v *NullableBillingPlan) Set(val *BillingPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPlan(val *BillingPlan) *NullableBillingPlan {
	return &NullableBillingPlan{value: val, isSet: true}
}

func (v NullableBillingPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


