# LedgerTransactionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreditAccountBalanceType** | Pointer to **string** |  | [optional] 
**CreditAccountSubtype** | Pointer to **string** |  | [optional] 
**CreditAccountType** | Pointer to **string** |  | [optional] 
**DebitAccountBalanceType** | Pointer to **string** |  | [optional] 
**DebitAccountSubtype** | Pointer to **string** |  | [optional] 
**DebitAccountType** | Pointer to **string** |  | [optional] 
**EntryOrder** | Pointer to **int32** |  | [optional] 
**EntryType** | Pointer to **string** |  | [optional] 
**MessageType** | Pointer to [**LedgerMessageType**](LedgerMessageType.md) |  | [optional] 
**ParamAccount1Subtype** | Pointer to **string** |  | [optional] 
**ParamAccount1Type** | Pointer to **string** |  | [optional] 
**ParamAccount2Subtype** | Pointer to **string** |  | [optional] 
**ParamAccount2Type** | Pointer to **string** |  | [optional] 
**ProcessType** | Pointer to [**LedgerProcessType**](LedgerProcessType.md) |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 

## Methods

### NewLedgerTransactionRule

`func NewLedgerTransactionRule() *LedgerTransactionRule`

NewLedgerTransactionRule instantiates a new LedgerTransactionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerTransactionRuleWithDefaults

`func NewLedgerTransactionRuleWithDefaults() *LedgerTransactionRule`

NewLedgerTransactionRuleWithDefaults instantiates a new LedgerTransactionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LedgerTransactionRule) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LedgerTransactionRule) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LedgerTransactionRule) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LedgerTransactionRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreditAccountBalanceType

`func (o *LedgerTransactionRule) GetCreditAccountBalanceType() string`

GetCreditAccountBalanceType returns the CreditAccountBalanceType field if non-nil, zero value otherwise.

### GetCreditAccountBalanceTypeOk

`func (o *LedgerTransactionRule) GetCreditAccountBalanceTypeOk() (*string, bool)`

GetCreditAccountBalanceTypeOk returns a tuple with the CreditAccountBalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountBalanceType

`func (o *LedgerTransactionRule) SetCreditAccountBalanceType(v string)`

SetCreditAccountBalanceType sets CreditAccountBalanceType field to given value.

### HasCreditAccountBalanceType

`func (o *LedgerTransactionRule) HasCreditAccountBalanceType() bool`

HasCreditAccountBalanceType returns a boolean if a field has been set.

### GetCreditAccountSubtype

`func (o *LedgerTransactionRule) GetCreditAccountSubtype() string`

GetCreditAccountSubtype returns the CreditAccountSubtype field if non-nil, zero value otherwise.

### GetCreditAccountSubtypeOk

`func (o *LedgerTransactionRule) GetCreditAccountSubtypeOk() (*string, bool)`

GetCreditAccountSubtypeOk returns a tuple with the CreditAccountSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountSubtype

`func (o *LedgerTransactionRule) SetCreditAccountSubtype(v string)`

SetCreditAccountSubtype sets CreditAccountSubtype field to given value.

### HasCreditAccountSubtype

`func (o *LedgerTransactionRule) HasCreditAccountSubtype() bool`

HasCreditAccountSubtype returns a boolean if a field has been set.

### GetCreditAccountType

`func (o *LedgerTransactionRule) GetCreditAccountType() string`

GetCreditAccountType returns the CreditAccountType field if non-nil, zero value otherwise.

### GetCreditAccountTypeOk

`func (o *LedgerTransactionRule) GetCreditAccountTypeOk() (*string, bool)`

GetCreditAccountTypeOk returns a tuple with the CreditAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountType

`func (o *LedgerTransactionRule) SetCreditAccountType(v string)`

SetCreditAccountType sets CreditAccountType field to given value.

### HasCreditAccountType

`func (o *LedgerTransactionRule) HasCreditAccountType() bool`

HasCreditAccountType returns a boolean if a field has been set.

### GetDebitAccountBalanceType

`func (o *LedgerTransactionRule) GetDebitAccountBalanceType() string`

GetDebitAccountBalanceType returns the DebitAccountBalanceType field if non-nil, zero value otherwise.

### GetDebitAccountBalanceTypeOk

`func (o *LedgerTransactionRule) GetDebitAccountBalanceTypeOk() (*string, bool)`

GetDebitAccountBalanceTypeOk returns a tuple with the DebitAccountBalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountBalanceType

`func (o *LedgerTransactionRule) SetDebitAccountBalanceType(v string)`

SetDebitAccountBalanceType sets DebitAccountBalanceType field to given value.

### HasDebitAccountBalanceType

`func (o *LedgerTransactionRule) HasDebitAccountBalanceType() bool`

HasDebitAccountBalanceType returns a boolean if a field has been set.

### GetDebitAccountSubtype

`func (o *LedgerTransactionRule) GetDebitAccountSubtype() string`

GetDebitAccountSubtype returns the DebitAccountSubtype field if non-nil, zero value otherwise.

### GetDebitAccountSubtypeOk

`func (o *LedgerTransactionRule) GetDebitAccountSubtypeOk() (*string, bool)`

GetDebitAccountSubtypeOk returns a tuple with the DebitAccountSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountSubtype

`func (o *LedgerTransactionRule) SetDebitAccountSubtype(v string)`

SetDebitAccountSubtype sets DebitAccountSubtype field to given value.

### HasDebitAccountSubtype

`func (o *LedgerTransactionRule) HasDebitAccountSubtype() bool`

HasDebitAccountSubtype returns a boolean if a field has been set.

### GetDebitAccountType

`func (o *LedgerTransactionRule) GetDebitAccountType() string`

GetDebitAccountType returns the DebitAccountType field if non-nil, zero value otherwise.

### GetDebitAccountTypeOk

`func (o *LedgerTransactionRule) GetDebitAccountTypeOk() (*string, bool)`

GetDebitAccountTypeOk returns a tuple with the DebitAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountType

`func (o *LedgerTransactionRule) SetDebitAccountType(v string)`

SetDebitAccountType sets DebitAccountType field to given value.

### HasDebitAccountType

`func (o *LedgerTransactionRule) HasDebitAccountType() bool`

HasDebitAccountType returns a boolean if a field has been set.

### GetEntryOrder

`func (o *LedgerTransactionRule) GetEntryOrder() int32`

GetEntryOrder returns the EntryOrder field if non-nil, zero value otherwise.

### GetEntryOrderOk

`func (o *LedgerTransactionRule) GetEntryOrderOk() (*int32, bool)`

GetEntryOrderOk returns a tuple with the EntryOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryOrder

`func (o *LedgerTransactionRule) SetEntryOrder(v int32)`

SetEntryOrder sets EntryOrder field to given value.

### HasEntryOrder

`func (o *LedgerTransactionRule) HasEntryOrder() bool`

HasEntryOrder returns a boolean if a field has been set.

### GetEntryType

`func (o *LedgerTransactionRule) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *LedgerTransactionRule) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *LedgerTransactionRule) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *LedgerTransactionRule) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### GetMessageType

`func (o *LedgerTransactionRule) GetMessageType() LedgerMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LedgerTransactionRule) GetMessageTypeOk() (*LedgerMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LedgerTransactionRule) SetMessageType(v LedgerMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LedgerTransactionRule) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetParamAccount1Subtype

`func (o *LedgerTransactionRule) GetParamAccount1Subtype() string`

GetParamAccount1Subtype returns the ParamAccount1Subtype field if non-nil, zero value otherwise.

### GetParamAccount1SubtypeOk

`func (o *LedgerTransactionRule) GetParamAccount1SubtypeOk() (*string, bool)`

GetParamAccount1SubtypeOk returns a tuple with the ParamAccount1Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamAccount1Subtype

`func (o *LedgerTransactionRule) SetParamAccount1Subtype(v string)`

SetParamAccount1Subtype sets ParamAccount1Subtype field to given value.

### HasParamAccount1Subtype

`func (o *LedgerTransactionRule) HasParamAccount1Subtype() bool`

HasParamAccount1Subtype returns a boolean if a field has been set.

### GetParamAccount1Type

`func (o *LedgerTransactionRule) GetParamAccount1Type() string`

GetParamAccount1Type returns the ParamAccount1Type field if non-nil, zero value otherwise.

### GetParamAccount1TypeOk

`func (o *LedgerTransactionRule) GetParamAccount1TypeOk() (*string, bool)`

GetParamAccount1TypeOk returns a tuple with the ParamAccount1Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamAccount1Type

`func (o *LedgerTransactionRule) SetParamAccount1Type(v string)`

SetParamAccount1Type sets ParamAccount1Type field to given value.

### HasParamAccount1Type

`func (o *LedgerTransactionRule) HasParamAccount1Type() bool`

HasParamAccount1Type returns a boolean if a field has been set.

### GetParamAccount2Subtype

`func (o *LedgerTransactionRule) GetParamAccount2Subtype() string`

GetParamAccount2Subtype returns the ParamAccount2Subtype field if non-nil, zero value otherwise.

### GetParamAccount2SubtypeOk

`func (o *LedgerTransactionRule) GetParamAccount2SubtypeOk() (*string, bool)`

GetParamAccount2SubtypeOk returns a tuple with the ParamAccount2Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamAccount2Subtype

`func (o *LedgerTransactionRule) SetParamAccount2Subtype(v string)`

SetParamAccount2Subtype sets ParamAccount2Subtype field to given value.

### HasParamAccount2Subtype

`func (o *LedgerTransactionRule) HasParamAccount2Subtype() bool`

HasParamAccount2Subtype returns a boolean if a field has been set.

### GetParamAccount2Type

`func (o *LedgerTransactionRule) GetParamAccount2Type() string`

GetParamAccount2Type returns the ParamAccount2Type field if non-nil, zero value otherwise.

### GetParamAccount2TypeOk

`func (o *LedgerTransactionRule) GetParamAccount2TypeOk() (*string, bool)`

GetParamAccount2TypeOk returns a tuple with the ParamAccount2Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamAccount2Type

`func (o *LedgerTransactionRule) SetParamAccount2Type(v string)`

SetParamAccount2Type sets ParamAccount2Type field to given value.

### HasParamAccount2Type

`func (o *LedgerTransactionRule) HasParamAccount2Type() bool`

HasParamAccount2Type returns a boolean if a field has been set.

### GetProcessType

`func (o *LedgerTransactionRule) GetProcessType() LedgerProcessType`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *LedgerTransactionRule) GetProcessTypeOk() (*LedgerProcessType, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *LedgerTransactionRule) SetProcessType(v LedgerProcessType)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *LedgerTransactionRule) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetTransactionType

`func (o *LedgerTransactionRule) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *LedgerTransactionRule) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *LedgerTransactionRule) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *LedgerTransactionRule) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


