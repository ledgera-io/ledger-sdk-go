# LedgerTransactionRuleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditAccountBalanceType** | **string** |  | 
**CreditAccountSubType** | **string** |  | 
**CreditAccountType** | **string** |  | 
**DebitAccountBalanceType** | **string** |  | 
**DebitAccountSubType** | **string** |  | 
**DebitAccountType** | **string** |  | 
**EntryOrder** | **int32** |  | 
**EntryType** | **string** |  | 
**MessageType** | [**LedgerMessageType**](LedgerMessageType.md) |  | 
**ParamAccount1Subtype** | Pointer to **string** |  | [optional] 
**ParamAccount1Type** | Pointer to **string** |  | [optional] 
**ParamAccount2Subtype** | Pointer to **string** |  | [optional] 
**ParamAccount2Type** | Pointer to **string** |  | [optional] 
**ProcessType** | [**LedgerProcessType**](LedgerProcessType.md) |  | 
**TransactionType** | **string** |  | 

## Methods

### NewLedgerTransactionRuleParams

`func NewLedgerTransactionRuleParams(creditAccountBalanceType string, creditAccountSubType string, creditAccountType string, debitAccountBalanceType string, debitAccountSubType string, debitAccountType string, entryOrder int32, entryType string, messageType LedgerMessageType, processType LedgerProcessType, transactionType string, ) *LedgerTransactionRuleParams`

NewLedgerTransactionRuleParams instantiates a new LedgerTransactionRuleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerTransactionRuleParamsWithDefaults

`func NewLedgerTransactionRuleParamsWithDefaults() *LedgerTransactionRuleParams`

NewLedgerTransactionRuleParamsWithDefaults instantiates a new LedgerTransactionRuleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditAccountBalanceType

`func (o *LedgerTransactionRuleParams) GetCreditAccountBalanceType() string`

GetCreditAccountBalanceType returns the CreditAccountBalanceType field if non-nil, zero value otherwise.

### GetCreditAccountBalanceTypeOk

`func (o *LedgerTransactionRuleParams) GetCreditAccountBalanceTypeOk() (*string, bool)`

GetCreditAccountBalanceTypeOk returns a tuple with the CreditAccountBalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountBalanceType

`func (o *LedgerTransactionRuleParams) SetCreditAccountBalanceType(v string)`

SetCreditAccountBalanceType sets CreditAccountBalanceType field to given value.


### GetCreditAccountSubType

`func (o *LedgerTransactionRuleParams) GetCreditAccountSubType() string`

GetCreditAccountSubType returns the CreditAccountSubType field if non-nil, zero value otherwise.

### GetCreditAccountSubTypeOk

`func (o *LedgerTransactionRuleParams) GetCreditAccountSubTypeOk() (*string, bool)`

GetCreditAccountSubTypeOk returns a tuple with the CreditAccountSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountSubType

`func (o *LedgerTransactionRuleParams) SetCreditAccountSubType(v string)`

SetCreditAccountSubType sets CreditAccountSubType field to given value.


### GetCreditAccountType

`func (o *LedgerTransactionRuleParams) GetCreditAccountType() string`

GetCreditAccountType returns the CreditAccountType field if non-nil, zero value otherwise.

### GetCreditAccountTypeOk

`func (o *LedgerTransactionRuleParams) GetCreditAccountTypeOk() (*string, bool)`

GetCreditAccountTypeOk returns a tuple with the CreditAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccountType

`func (o *LedgerTransactionRuleParams) SetCreditAccountType(v string)`

SetCreditAccountType sets CreditAccountType field to given value.


### GetDebitAccountBalanceType

`func (o *LedgerTransactionRuleParams) GetDebitAccountBalanceType() string`

GetDebitAccountBalanceType returns the DebitAccountBalanceType field if non-nil, zero value otherwise.

### GetDebitAccountBalanceTypeOk

`func (o *LedgerTransactionRuleParams) GetDebitAccountBalanceTypeOk() (*string, bool)`

GetDebitAccountBalanceTypeOk returns a tuple with the DebitAccountBalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountBalanceType

`func (o *LedgerTransactionRuleParams) SetDebitAccountBalanceType(v string)`

SetDebitAccountBalanceType sets DebitAccountBalanceType field to given value.


### GetDebitAccountSubType

`func (o *LedgerTransactionRuleParams) GetDebitAccountSubType() string`

GetDebitAccountSubType returns the DebitAccountSubType field if non-nil, zero value otherwise.

### GetDebitAccountSubTypeOk

`func (o *LedgerTransactionRuleParams) GetDebitAccountSubTypeOk() (*string, bool)`

GetDebitAccountSubTypeOk returns a tuple with the DebitAccountSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountSubType

`func (o *LedgerTransactionRuleParams) SetDebitAccountSubType(v string)`

SetDebitAccountSubType sets DebitAccountSubType field to given value.


### GetDebitAccountType

`func (o *LedgerTransactionRuleParams) GetDebitAccountType() string`

GetDebitAccountType returns the DebitAccountType field if non-nil, zero value otherwise.

### GetDebitAccountTypeOk

`func (o *LedgerTransactionRuleParams) GetDebitAccountTypeOk() (*string, bool)`

GetDebitAccountTypeOk returns a tuple with the DebitAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAccountType

`func (o *LedgerTransactionRuleParams) SetDebitAccountType(v string)`

SetDebitAccountType sets DebitAccountType field to given value.


### GetEntryOrder

`func (o *LedgerTransactionRuleParams) GetEntryOrder() int32`

GetEntryOrder returns the EntryOrder field if non-nil, zero value otherwise.

### GetEntryOrderOk

`func (o *LedgerTransactionRuleParams) GetEntryOrderOk() (*int32, bool)`

GetEntryOrderOk returns a tuple with the EntryOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryOrder

`func (o *LedgerTransactionRuleParams) SetEntryOrder(v int32)`

SetEntryOrder sets EntryOrder field to given value.


### GetEntryType

`func (o *LedgerTransactionRuleParams) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *LedgerTransactionRuleParams) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *LedgerTransactionRuleParams) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.


### GetMessageType

`func (o *LedgerTransactionRuleParams) GetMessageType() LedgerMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LedgerTransactionRuleParams) GetMessageTypeOk() (*LedgerMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LedgerTransactionRuleParams) SetMessageType(v LedgerMessageType)`

SetMessageType sets MessageType field to given value.


### GetParamAccount1Subtype

`func (o *LedgerTransactionRuleParams) GetParamAccount1Subtype() string`

GetParamAccount1Subtype returns the ParamAccount1Subtype field if non-nil, zero value otherwise.

### GetParamAccount1SubtypeOk

`func (o *LedgerTransactionRuleParams) GetParamAccount1SubtypeOk() (*string, bool)`

GetParamAccount1SubtypeOk returns a tuple with the ParamAccount1Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamAccount1Subtype

`func (o *LedgerTransactionRuleParams) SetParamAccount1Subtype(v string)`

SetParamAccount1Subtype sets ParamAccount1Subtype field to given value.

### HasParamAccount1Subtype

`func (o *LedgerTransactionRuleParams) HasParamAccount1Subtype() bool`

HasParamAccount1Subtype returns a boolean if a field has been set.

### GetParamAccount1Type

`func (o *LedgerTransactionRuleParams) GetParamAccount1Type() string`

GetParamAccount1Type returns the ParamAccount1Type field if non-nil, zero value otherwise.

### GetParamAccount1TypeOk

`func (o *LedgerTransactionRuleParams) GetParamAccount1TypeOk() (*string, bool)`

GetParamAccount1TypeOk returns a tuple with the ParamAccount1Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamAccount1Type

`func (o *LedgerTransactionRuleParams) SetParamAccount1Type(v string)`

SetParamAccount1Type sets ParamAccount1Type field to given value.

### HasParamAccount1Type

`func (o *LedgerTransactionRuleParams) HasParamAccount1Type() bool`

HasParamAccount1Type returns a boolean if a field has been set.

### GetParamAccount2Subtype

`func (o *LedgerTransactionRuleParams) GetParamAccount2Subtype() string`

GetParamAccount2Subtype returns the ParamAccount2Subtype field if non-nil, zero value otherwise.

### GetParamAccount2SubtypeOk

`func (o *LedgerTransactionRuleParams) GetParamAccount2SubtypeOk() (*string, bool)`

GetParamAccount2SubtypeOk returns a tuple with the ParamAccount2Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamAccount2Subtype

`func (o *LedgerTransactionRuleParams) SetParamAccount2Subtype(v string)`

SetParamAccount2Subtype sets ParamAccount2Subtype field to given value.

### HasParamAccount2Subtype

`func (o *LedgerTransactionRuleParams) HasParamAccount2Subtype() bool`

HasParamAccount2Subtype returns a boolean if a field has been set.

### GetParamAccount2Type

`func (o *LedgerTransactionRuleParams) GetParamAccount2Type() string`

GetParamAccount2Type returns the ParamAccount2Type field if non-nil, zero value otherwise.

### GetParamAccount2TypeOk

`func (o *LedgerTransactionRuleParams) GetParamAccount2TypeOk() (*string, bool)`

GetParamAccount2TypeOk returns a tuple with the ParamAccount2Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamAccount2Type

`func (o *LedgerTransactionRuleParams) SetParamAccount2Type(v string)`

SetParamAccount2Type sets ParamAccount2Type field to given value.

### HasParamAccount2Type

`func (o *LedgerTransactionRuleParams) HasParamAccount2Type() bool`

HasParamAccount2Type returns a boolean if a field has been set.

### GetProcessType

`func (o *LedgerTransactionRuleParams) GetProcessType() LedgerProcessType`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *LedgerTransactionRuleParams) GetProcessTypeOk() (*LedgerProcessType, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *LedgerTransactionRuleParams) SetProcessType(v LedgerProcessType)`

SetProcessType sets ProcessType field to given value.


### GetTransactionType

`func (o *LedgerTransactionRuleParams) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *LedgerTransactionRuleParams) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *LedgerTransactionRuleParams) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


