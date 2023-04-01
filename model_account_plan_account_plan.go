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

// checks if the AccountPlanAccountPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountPlanAccountPlan{}

// AccountPlanAccountPlan struct for AccountPlanAccountPlan
type AccountPlanAccountPlan struct {
	BalanceType *AccountPlanBalanceType `json:"balanceType,omitempty"`
	BalanceValidation *AccountPlanBalanceValidation `json:"balanceValidation,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	DeletedAt *string `json:"deletedAt,omitempty"`
	Id *string `json:"id,omitempty"`
	IsUnique *bool `json:"isUnique,omitempty"`
	SubType *string `json:"subType,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewAccountPlanAccountPlan instantiates a new AccountPlanAccountPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountPlanAccountPlan() *AccountPlanAccountPlan {
	this := AccountPlanAccountPlan{}
	return &this
}

// NewAccountPlanAccountPlanWithDefaults instantiates a new AccountPlanAccountPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountPlanAccountPlanWithDefaults() *AccountPlanAccountPlan {
	this := AccountPlanAccountPlan{}
	return &this
}

// GetBalanceType returns the BalanceType field value if set, zero value otherwise.
func (o *AccountPlanAccountPlan) GetBalanceType() AccountPlanBalanceType {
	if o == nil || IsNil(o.BalanceType) {
		var ret AccountPlanBalanceType
		return ret
	}
	return *o.BalanceType
}

// GetBalanceTypeOk returns a tuple with the BalanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPlanAccountPlan) GetBalanceTypeOk() (*AccountPlanBalanceType, bool) {
	if o == nil || IsNil(o.BalanceType) {
		return nil, false
	}
	return o.BalanceType, true
}

// HasBalanceType returns a boolean if a field has been set.
func (o *AccountPlanAccountPlan) HasBalanceType() bool {
	if o != nil && !IsNil(o.BalanceType) {
		return true
	}

	return false
}

// SetBalanceType gets a reference to the given AccountPlanBalanceType and assigns it to the BalanceType field.
func (o *AccountPlanAccountPlan) SetBalanceType(v AccountPlanBalanceType) {
	o.BalanceType = &v
}

// GetBalanceValidation returns the BalanceValidation field value if set, zero value otherwise.
func (o *AccountPlanAccountPlan) GetBalanceValidation() AccountPlanBalanceValidation {
	if o == nil || IsNil(o.BalanceValidation) {
		var ret AccountPlanBalanceValidation
		return ret
	}
	return *o.BalanceValidation
}

// GetBalanceValidationOk returns a tuple with the BalanceValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPlanAccountPlan) GetBalanceValidationOk() (*AccountPlanBalanceValidation, bool) {
	if o == nil || IsNil(o.BalanceValidation) {
		return nil, false
	}
	return o.BalanceValidation, true
}

// HasBalanceValidation returns a boolean if a field has been set.
func (o *AccountPlanAccountPlan) HasBalanceValidation() bool {
	if o != nil && !IsNil(o.BalanceValidation) {
		return true
	}

	return false
}

// SetBalanceValidation gets a reference to the given AccountPlanBalanceValidation and assigns it to the BalanceValidation field.
func (o *AccountPlanAccountPlan) SetBalanceValidation(v AccountPlanBalanceValidation) {
	o.BalanceValidation = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AccountPlanAccountPlan) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPlanAccountPlan) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AccountPlanAccountPlan) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AccountPlanAccountPlan) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AccountPlanAccountPlan) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPlanAccountPlan) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AccountPlanAccountPlan) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AccountPlanAccountPlan) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountPlanAccountPlan) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPlanAccountPlan) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountPlanAccountPlan) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountPlanAccountPlan) SetId(v string) {
	o.Id = &v
}

// GetIsUnique returns the IsUnique field value if set, zero value otherwise.
func (o *AccountPlanAccountPlan) GetIsUnique() bool {
	if o == nil || IsNil(o.IsUnique) {
		var ret bool
		return ret
	}
	return *o.IsUnique
}

// GetIsUniqueOk returns a tuple with the IsUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPlanAccountPlan) GetIsUniqueOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUnique) {
		return nil, false
	}
	return o.IsUnique, true
}

// HasIsUnique returns a boolean if a field has been set.
func (o *AccountPlanAccountPlan) HasIsUnique() bool {
	if o != nil && !IsNil(o.IsUnique) {
		return true
	}

	return false
}

// SetIsUnique gets a reference to the given bool and assigns it to the IsUnique field.
func (o *AccountPlanAccountPlan) SetIsUnique(v bool) {
	o.IsUnique = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *AccountPlanAccountPlan) GetSubType() string {
	if o == nil || IsNil(o.SubType) {
		var ret string
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPlanAccountPlan) GetSubTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubType) {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *AccountPlanAccountPlan) HasSubType() bool {
	if o != nil && !IsNil(o.SubType) {
		return true
	}

	return false
}

// SetSubType gets a reference to the given string and assigns it to the SubType field.
func (o *AccountPlanAccountPlan) SetSubType(v string) {
	o.SubType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountPlanAccountPlan) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPlanAccountPlan) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountPlanAccountPlan) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccountPlanAccountPlan) SetType(v string) {
	o.Type = &v
}

func (o AccountPlanAccountPlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountPlanAccountPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BalanceType) {
		toSerialize["balanceType"] = o.BalanceType
	}
	if !IsNil(o.BalanceValidation) {
		toSerialize["balanceValidation"] = o.BalanceValidation
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsUnique) {
		toSerialize["isUnique"] = o.IsUnique
	}
	if !IsNil(o.SubType) {
		toSerialize["subType"] = o.SubType
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAccountPlanAccountPlan struct {
	value *AccountPlanAccountPlan
	isSet bool
}

func (v NullableAccountPlanAccountPlan) Get() *AccountPlanAccountPlan {
	return v.value
}

func (v *NullableAccountPlanAccountPlan) Set(val *AccountPlanAccountPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountPlanAccountPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountPlanAccountPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountPlanAccountPlan(val *AccountPlanAccountPlan) *NullableAccountPlanAccountPlan {
	return &NullableAccountPlanAccountPlan{value: val, isSet: true}
}

func (v NullableAccountPlanAccountPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountPlanAccountPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


