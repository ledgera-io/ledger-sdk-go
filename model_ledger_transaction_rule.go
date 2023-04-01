/*
Ledgera API

teste

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the LedgerTransactionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LedgerTransactionRule{}

// LedgerTransactionRule struct for LedgerTransactionRule
type LedgerTransactionRule struct {
	CreatedAt *string `json:"createdAt,omitempty"`
	CreditAccountBalanceType *string `json:"creditAccountBalanceType,omitempty"`
	CreditAccountSubtype *string `json:"creditAccountSubtype,omitempty"`
	CreditAccountType *string `json:"creditAccountType,omitempty"`
	DebitAccountBalanceType *string `json:"debitAccountBalanceType,omitempty"`
	DebitAccountSubtype *string `json:"debitAccountSubtype,omitempty"`
	DebitAccountType *string `json:"debitAccountType,omitempty"`
	EntryOrder *int32 `json:"entryOrder,omitempty"`
	EntryType *string `json:"entryType,omitempty"`
	MessageType *LedgerMessageType `json:"messageType,omitempty"`
	ParamAccount1Subtype *string `json:"paramAccount1Subtype,omitempty"`
	ParamAccount1Type *string `json:"paramAccount1Type,omitempty"`
	ParamAccount2Subtype *string `json:"paramAccount2Subtype,omitempty"`
	ParamAccount2Type *string `json:"paramAccount2Type,omitempty"`
	ProcessType *LedgerProcessType `json:"processType,omitempty"`
	TransactionType *string `json:"transactionType,omitempty"`
}

// NewLedgerTransactionRule instantiates a new LedgerTransactionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLedgerTransactionRule() *LedgerTransactionRule {
	this := LedgerTransactionRule{}
	return &this
}

// NewLedgerTransactionRuleWithDefaults instantiates a new LedgerTransactionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLedgerTransactionRuleWithDefaults() *LedgerTransactionRule {
	this := LedgerTransactionRule{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *LedgerTransactionRule) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreditAccountBalanceType returns the CreditAccountBalanceType field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetCreditAccountBalanceType() string {
	if o == nil || IsNil(o.CreditAccountBalanceType) {
		var ret string
		return ret
	}
	return *o.CreditAccountBalanceType
}

// GetCreditAccountBalanceTypeOk returns a tuple with the CreditAccountBalanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetCreditAccountBalanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CreditAccountBalanceType) {
		return nil, false
	}
	return o.CreditAccountBalanceType, true
}

// HasCreditAccountBalanceType returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasCreditAccountBalanceType() bool {
	if o != nil && !IsNil(o.CreditAccountBalanceType) {
		return true
	}

	return false
}

// SetCreditAccountBalanceType gets a reference to the given string and assigns it to the CreditAccountBalanceType field.
func (o *LedgerTransactionRule) SetCreditAccountBalanceType(v string) {
	o.CreditAccountBalanceType = &v
}

// GetCreditAccountSubtype returns the CreditAccountSubtype field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetCreditAccountSubtype() string {
	if o == nil || IsNil(o.CreditAccountSubtype) {
		var ret string
		return ret
	}
	return *o.CreditAccountSubtype
}

// GetCreditAccountSubtypeOk returns a tuple with the CreditAccountSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetCreditAccountSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.CreditAccountSubtype) {
		return nil, false
	}
	return o.CreditAccountSubtype, true
}

// HasCreditAccountSubtype returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasCreditAccountSubtype() bool {
	if o != nil && !IsNil(o.CreditAccountSubtype) {
		return true
	}

	return false
}

// SetCreditAccountSubtype gets a reference to the given string and assigns it to the CreditAccountSubtype field.
func (o *LedgerTransactionRule) SetCreditAccountSubtype(v string) {
	o.CreditAccountSubtype = &v
}

// GetCreditAccountType returns the CreditAccountType field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetCreditAccountType() string {
	if o == nil || IsNil(o.CreditAccountType) {
		var ret string
		return ret
	}
	return *o.CreditAccountType
}

// GetCreditAccountTypeOk returns a tuple with the CreditAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetCreditAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CreditAccountType) {
		return nil, false
	}
	return o.CreditAccountType, true
}

// HasCreditAccountType returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasCreditAccountType() bool {
	if o != nil && !IsNil(o.CreditAccountType) {
		return true
	}

	return false
}

// SetCreditAccountType gets a reference to the given string and assigns it to the CreditAccountType field.
func (o *LedgerTransactionRule) SetCreditAccountType(v string) {
	o.CreditAccountType = &v
}

// GetDebitAccountBalanceType returns the DebitAccountBalanceType field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetDebitAccountBalanceType() string {
	if o == nil || IsNil(o.DebitAccountBalanceType) {
		var ret string
		return ret
	}
	return *o.DebitAccountBalanceType
}

// GetDebitAccountBalanceTypeOk returns a tuple with the DebitAccountBalanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetDebitAccountBalanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DebitAccountBalanceType) {
		return nil, false
	}
	return o.DebitAccountBalanceType, true
}

// HasDebitAccountBalanceType returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasDebitAccountBalanceType() bool {
	if o != nil && !IsNil(o.DebitAccountBalanceType) {
		return true
	}

	return false
}

// SetDebitAccountBalanceType gets a reference to the given string and assigns it to the DebitAccountBalanceType field.
func (o *LedgerTransactionRule) SetDebitAccountBalanceType(v string) {
	o.DebitAccountBalanceType = &v
}

// GetDebitAccountSubtype returns the DebitAccountSubtype field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetDebitAccountSubtype() string {
	if o == nil || IsNil(o.DebitAccountSubtype) {
		var ret string
		return ret
	}
	return *o.DebitAccountSubtype
}

// GetDebitAccountSubtypeOk returns a tuple with the DebitAccountSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetDebitAccountSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.DebitAccountSubtype) {
		return nil, false
	}
	return o.DebitAccountSubtype, true
}

// HasDebitAccountSubtype returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasDebitAccountSubtype() bool {
	if o != nil && !IsNil(o.DebitAccountSubtype) {
		return true
	}

	return false
}

// SetDebitAccountSubtype gets a reference to the given string and assigns it to the DebitAccountSubtype field.
func (o *LedgerTransactionRule) SetDebitAccountSubtype(v string) {
	o.DebitAccountSubtype = &v
}

// GetDebitAccountType returns the DebitAccountType field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetDebitAccountType() string {
	if o == nil || IsNil(o.DebitAccountType) {
		var ret string
		return ret
	}
	return *o.DebitAccountType
}

// GetDebitAccountTypeOk returns a tuple with the DebitAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetDebitAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DebitAccountType) {
		return nil, false
	}
	return o.DebitAccountType, true
}

// HasDebitAccountType returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasDebitAccountType() bool {
	if o != nil && !IsNil(o.DebitAccountType) {
		return true
	}

	return false
}

// SetDebitAccountType gets a reference to the given string and assigns it to the DebitAccountType field.
func (o *LedgerTransactionRule) SetDebitAccountType(v string) {
	o.DebitAccountType = &v
}

// GetEntryOrder returns the EntryOrder field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetEntryOrder() int32 {
	if o == nil || IsNil(o.EntryOrder) {
		var ret int32
		return ret
	}
	return *o.EntryOrder
}

// GetEntryOrderOk returns a tuple with the EntryOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetEntryOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.EntryOrder) {
		return nil, false
	}
	return o.EntryOrder, true
}

// HasEntryOrder returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasEntryOrder() bool {
	if o != nil && !IsNil(o.EntryOrder) {
		return true
	}

	return false
}

// SetEntryOrder gets a reference to the given int32 and assigns it to the EntryOrder field.
func (o *LedgerTransactionRule) SetEntryOrder(v int32) {
	o.EntryOrder = &v
}

// GetEntryType returns the EntryType field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetEntryType() string {
	if o == nil || IsNil(o.EntryType) {
		var ret string
		return ret
	}
	return *o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetEntryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntryType) {
		return nil, false
	}
	return o.EntryType, true
}

// HasEntryType returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasEntryType() bool {
	if o != nil && !IsNil(o.EntryType) {
		return true
	}

	return false
}

// SetEntryType gets a reference to the given string and assigns it to the EntryType field.
func (o *LedgerTransactionRule) SetEntryType(v string) {
	o.EntryType = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetMessageType() LedgerMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret LedgerMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetMessageTypeOk() (*LedgerMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given LedgerMessageType and assigns it to the MessageType field.
func (o *LedgerTransactionRule) SetMessageType(v LedgerMessageType) {
	o.MessageType = &v
}

// GetParamAccount1Subtype returns the ParamAccount1Subtype field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetParamAccount1Subtype() string {
	if o == nil || IsNil(o.ParamAccount1Subtype) {
		var ret string
		return ret
	}
	return *o.ParamAccount1Subtype
}

// GetParamAccount1SubtypeOk returns a tuple with the ParamAccount1Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetParamAccount1SubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParamAccount1Subtype) {
		return nil, false
	}
	return o.ParamAccount1Subtype, true
}

// HasParamAccount1Subtype returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasParamAccount1Subtype() bool {
	if o != nil && !IsNil(o.ParamAccount1Subtype) {
		return true
	}

	return false
}

// SetParamAccount1Subtype gets a reference to the given string and assigns it to the ParamAccount1Subtype field.
func (o *LedgerTransactionRule) SetParamAccount1Subtype(v string) {
	o.ParamAccount1Subtype = &v
}

// GetParamAccount1Type returns the ParamAccount1Type field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetParamAccount1Type() string {
	if o == nil || IsNil(o.ParamAccount1Type) {
		var ret string
		return ret
	}
	return *o.ParamAccount1Type
}

// GetParamAccount1TypeOk returns a tuple with the ParamAccount1Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetParamAccount1TypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParamAccount1Type) {
		return nil, false
	}
	return o.ParamAccount1Type, true
}

// HasParamAccount1Type returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasParamAccount1Type() bool {
	if o != nil && !IsNil(o.ParamAccount1Type) {
		return true
	}

	return false
}

// SetParamAccount1Type gets a reference to the given string and assigns it to the ParamAccount1Type field.
func (o *LedgerTransactionRule) SetParamAccount1Type(v string) {
	o.ParamAccount1Type = &v
}

// GetParamAccount2Subtype returns the ParamAccount2Subtype field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetParamAccount2Subtype() string {
	if o == nil || IsNil(o.ParamAccount2Subtype) {
		var ret string
		return ret
	}
	return *o.ParamAccount2Subtype
}

// GetParamAccount2SubtypeOk returns a tuple with the ParamAccount2Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetParamAccount2SubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParamAccount2Subtype) {
		return nil, false
	}
	return o.ParamAccount2Subtype, true
}

// HasParamAccount2Subtype returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasParamAccount2Subtype() bool {
	if o != nil && !IsNil(o.ParamAccount2Subtype) {
		return true
	}

	return false
}

// SetParamAccount2Subtype gets a reference to the given string and assigns it to the ParamAccount2Subtype field.
func (o *LedgerTransactionRule) SetParamAccount2Subtype(v string) {
	o.ParamAccount2Subtype = &v
}

// GetParamAccount2Type returns the ParamAccount2Type field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetParamAccount2Type() string {
	if o == nil || IsNil(o.ParamAccount2Type) {
		var ret string
		return ret
	}
	return *o.ParamAccount2Type
}

// GetParamAccount2TypeOk returns a tuple with the ParamAccount2Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetParamAccount2TypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParamAccount2Type) {
		return nil, false
	}
	return o.ParamAccount2Type, true
}

// HasParamAccount2Type returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasParamAccount2Type() bool {
	if o != nil && !IsNil(o.ParamAccount2Type) {
		return true
	}

	return false
}

// SetParamAccount2Type gets a reference to the given string and assigns it to the ParamAccount2Type field.
func (o *LedgerTransactionRule) SetParamAccount2Type(v string) {
	o.ParamAccount2Type = &v
}

// GetProcessType returns the ProcessType field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetProcessType() LedgerProcessType {
	if o == nil || IsNil(o.ProcessType) {
		var ret LedgerProcessType
		return ret
	}
	return *o.ProcessType
}

// GetProcessTypeOk returns a tuple with the ProcessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetProcessTypeOk() (*LedgerProcessType, bool) {
	if o == nil || IsNil(o.ProcessType) {
		return nil, false
	}
	return o.ProcessType, true
}

// HasProcessType returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasProcessType() bool {
	if o != nil && !IsNil(o.ProcessType) {
		return true
	}

	return false
}

// SetProcessType gets a reference to the given LedgerProcessType and assigns it to the ProcessType field.
func (o *LedgerTransactionRule) SetProcessType(v LedgerProcessType) {
	o.ProcessType = &v
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *LedgerTransactionRule) GetTransactionType() string {
	if o == nil || IsNil(o.TransactionType) {
		var ret string
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionRule) GetTransactionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionType) {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *LedgerTransactionRule) HasTransactionType() bool {
	if o != nil && !IsNil(o.TransactionType) {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given string and assigns it to the TransactionType field.
func (o *LedgerTransactionRule) SetTransactionType(v string) {
	o.TransactionType = &v
}

func (o LedgerTransactionRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LedgerTransactionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.CreditAccountBalanceType) {
		toSerialize["creditAccountBalanceType"] = o.CreditAccountBalanceType
	}
	if !IsNil(o.CreditAccountSubtype) {
		toSerialize["creditAccountSubtype"] = o.CreditAccountSubtype
	}
	if !IsNil(o.CreditAccountType) {
		toSerialize["creditAccountType"] = o.CreditAccountType
	}
	if !IsNil(o.DebitAccountBalanceType) {
		toSerialize["debitAccountBalanceType"] = o.DebitAccountBalanceType
	}
	if !IsNil(o.DebitAccountSubtype) {
		toSerialize["debitAccountSubtype"] = o.DebitAccountSubtype
	}
	if !IsNil(o.DebitAccountType) {
		toSerialize["debitAccountType"] = o.DebitAccountType
	}
	if !IsNil(o.EntryOrder) {
		toSerialize["entryOrder"] = o.EntryOrder
	}
	if !IsNil(o.EntryType) {
		toSerialize["entryType"] = o.EntryType
	}
	if !IsNil(o.MessageType) {
		toSerialize["messageType"] = o.MessageType
	}
	if !IsNil(o.ParamAccount1Subtype) {
		toSerialize["paramAccount1Subtype"] = o.ParamAccount1Subtype
	}
	if !IsNil(o.ParamAccount1Type) {
		toSerialize["paramAccount1Type"] = o.ParamAccount1Type
	}
	if !IsNil(o.ParamAccount2Subtype) {
		toSerialize["paramAccount2Subtype"] = o.ParamAccount2Subtype
	}
	if !IsNil(o.ParamAccount2Type) {
		toSerialize["paramAccount2Type"] = o.ParamAccount2Type
	}
	if !IsNil(o.ProcessType) {
		toSerialize["processType"] = o.ProcessType
	}
	if !IsNil(o.TransactionType) {
		toSerialize["transactionType"] = o.TransactionType
	}
	return toSerialize, nil
}

type NullableLedgerTransactionRule struct {
	value *LedgerTransactionRule
	isSet bool
}

func (v NullableLedgerTransactionRule) Get() *LedgerTransactionRule {
	return v.value
}

func (v *NullableLedgerTransactionRule) Set(val *LedgerTransactionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerTransactionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerTransactionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerTransactionRule(val *LedgerTransactionRule) *NullableLedgerTransactionRule {
	return &NullableLedgerTransactionRule{value: val, isSet: true}
}

func (v NullableLedgerTransactionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerTransactionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


