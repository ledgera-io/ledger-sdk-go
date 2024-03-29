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

// checks if the AccountPlanCreateAccountPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountPlanCreateAccountPlan{}

// AccountPlanCreateAccountPlan struct for AccountPlanCreateAccountPlan
type AccountPlanCreateAccountPlan struct {
	BalanceType AccountPlanBalanceType `json:"balanceType"`
	BalanceValidation string `json:"balanceValidation"`
	IsUnique *bool `json:"isUnique,omitempty"`
	SubType string `json:"subType"`
	Type string `json:"type"`
}

// NewAccountPlanCreateAccountPlan instantiates a new AccountPlanCreateAccountPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountPlanCreateAccountPlan(balanceType AccountPlanBalanceType, balanceValidation string, subType string, type_ string) *AccountPlanCreateAccountPlan {
	this := AccountPlanCreateAccountPlan{}
	this.BalanceType = balanceType
	this.BalanceValidation = balanceValidation
	this.SubType = subType
	this.Type = type_
	return &this
}

// NewAccountPlanCreateAccountPlanWithDefaults instantiates a new AccountPlanCreateAccountPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountPlanCreateAccountPlanWithDefaults() *AccountPlanCreateAccountPlan {
	this := AccountPlanCreateAccountPlan{}
	return &this
}

// GetBalanceType returns the BalanceType field value
func (o *AccountPlanCreateAccountPlan) GetBalanceType() AccountPlanBalanceType {
	if o == nil {
		var ret AccountPlanBalanceType
		return ret
	}

	return o.BalanceType
}

// GetBalanceTypeOk returns a tuple with the BalanceType field value
// and a boolean to check if the value has been set.
func (o *AccountPlanCreateAccountPlan) GetBalanceTypeOk() (*AccountPlanBalanceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceType, true
}

// SetBalanceType sets field value
func (o *AccountPlanCreateAccountPlan) SetBalanceType(v AccountPlanBalanceType) {
	o.BalanceType = v
}

// GetBalanceValidation returns the BalanceValidation field value
func (o *AccountPlanCreateAccountPlan) GetBalanceValidation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalanceValidation
}

// GetBalanceValidationOk returns a tuple with the BalanceValidation field value
// and a boolean to check if the value has been set.
func (o *AccountPlanCreateAccountPlan) GetBalanceValidationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceValidation, true
}

// SetBalanceValidation sets field value
func (o *AccountPlanCreateAccountPlan) SetBalanceValidation(v string) {
	o.BalanceValidation = v
}

// GetIsUnique returns the IsUnique field value if set, zero value otherwise.
func (o *AccountPlanCreateAccountPlan) GetIsUnique() bool {
	if o == nil || IsNil(o.IsUnique) {
		var ret bool
		return ret
	}
	return *o.IsUnique
}

// GetIsUniqueOk returns a tuple with the IsUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPlanCreateAccountPlan) GetIsUniqueOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUnique) {
		return nil, false
	}
	return o.IsUnique, true
}

// HasIsUnique returns a boolean if a field has been set.
func (o *AccountPlanCreateAccountPlan) HasIsUnique() bool {
	if o != nil && !IsNil(o.IsUnique) {
		return true
	}

	return false
}

// SetIsUnique gets a reference to the given bool and assigns it to the IsUnique field.
func (o *AccountPlanCreateAccountPlan) SetIsUnique(v bool) {
	o.IsUnique = &v
}

// GetSubType returns the SubType field value
func (o *AccountPlanCreateAccountPlan) GetSubType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value
// and a boolean to check if the value has been set.
func (o *AccountPlanCreateAccountPlan) GetSubTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubType, true
}

// SetSubType sets field value
func (o *AccountPlanCreateAccountPlan) SetSubType(v string) {
	o.SubType = v
}

// GetType returns the Type field value
func (o *AccountPlanCreateAccountPlan) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountPlanCreateAccountPlan) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountPlanCreateAccountPlan) SetType(v string) {
	o.Type = v
}

func (o AccountPlanCreateAccountPlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountPlanCreateAccountPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balanceType"] = o.BalanceType
	toSerialize["balanceValidation"] = o.BalanceValidation
	if !IsNil(o.IsUnique) {
		toSerialize["isUnique"] = o.IsUnique
	}
	toSerialize["subType"] = o.SubType
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAccountPlanCreateAccountPlan struct {
	value *AccountPlanCreateAccountPlan
	isSet bool
}

func (v NullableAccountPlanCreateAccountPlan) Get() *AccountPlanCreateAccountPlan {
	return v.value
}

func (v *NullableAccountPlanCreateAccountPlan) Set(val *AccountPlanCreateAccountPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountPlanCreateAccountPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountPlanCreateAccountPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountPlanCreateAccountPlan(val *AccountPlanCreateAccountPlan) *NullableAccountPlanCreateAccountPlan {
	return &NullableAccountPlanCreateAccountPlan{value: val, isSet: true}
}

func (v NullableAccountPlanCreateAccountPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountPlanCreateAccountPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


