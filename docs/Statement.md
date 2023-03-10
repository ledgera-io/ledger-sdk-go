# Statement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SeqNum** | Pointer to **int32** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**EntryId** | Pointer to **string** |  | [optional] 
**AccountLabel** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**EntryLabel** | Pointer to **string** |  | [optional] 
**EntryOrder** | Pointer to **int32** |  | [optional] 
**BalanceType** | Pointer to [**BalanceType**](BalanceType.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**ConciliationId** | Pointer to **string** |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 
**ProcessType** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**BalanceAfter** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**SettledAt** | Pointer to **string** |  | [optional] 

## Methods

### NewStatement

`func NewStatement() *Statement`

NewStatement instantiates a new Statement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementWithDefaults

`func NewStatementWithDefaults() *Statement`

NewStatementWithDefaults instantiates a new Statement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Statement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Statement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Statement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Statement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSeqNum

`func (o *Statement) GetSeqNum() int32`

GetSeqNum returns the SeqNum field if non-nil, zero value otherwise.

### GetSeqNumOk

`func (o *Statement) GetSeqNumOk() (*int32, bool)`

GetSeqNumOk returns a tuple with the SeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNum

`func (o *Statement) SetSeqNum(v int32)`

SetSeqNum sets SeqNum field to given value.

### HasSeqNum

`func (o *Statement) HasSeqNum() bool`

HasSeqNum returns a boolean if a field has been set.

### GetTransactionId

`func (o *Statement) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Statement) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Statement) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Statement) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetEntryId

`func (o *Statement) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *Statement) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *Statement) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *Statement) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetAccountLabel

`func (o *Statement) GetAccountLabel() string`

GetAccountLabel returns the AccountLabel field if non-nil, zero value otherwise.

### GetAccountLabelOk

`func (o *Statement) GetAccountLabelOk() (*string, bool)`

GetAccountLabelOk returns a tuple with the AccountLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLabel

`func (o *Statement) SetAccountLabel(v string)`

SetAccountLabel sets AccountLabel field to given value.

### HasAccountLabel

`func (o *Statement) HasAccountLabel() bool`

HasAccountLabel returns a boolean if a field has been set.

### GetAccountType

`func (o *Statement) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *Statement) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *Statement) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *Statement) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAccountId

`func (o *Statement) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Statement) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Statement) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Statement) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetEntryLabel

`func (o *Statement) GetEntryLabel() string`

GetEntryLabel returns the EntryLabel field if non-nil, zero value otherwise.

### GetEntryLabelOk

`func (o *Statement) GetEntryLabelOk() (*string, bool)`

GetEntryLabelOk returns a tuple with the EntryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryLabel

`func (o *Statement) SetEntryLabel(v string)`

SetEntryLabel sets EntryLabel field to given value.

### HasEntryLabel

`func (o *Statement) HasEntryLabel() bool`

HasEntryLabel returns a boolean if a field has been set.

### GetEntryOrder

`func (o *Statement) GetEntryOrder() int32`

GetEntryOrder returns the EntryOrder field if non-nil, zero value otherwise.

### GetEntryOrderOk

`func (o *Statement) GetEntryOrderOk() (*int32, bool)`

GetEntryOrderOk returns a tuple with the EntryOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryOrder

`func (o *Statement) SetEntryOrder(v int32)`

SetEntryOrder sets EntryOrder field to given value.

### HasEntryOrder

`func (o *Statement) HasEntryOrder() bool`

HasEntryOrder returns a boolean if a field has been set.

### GetBalanceType

`func (o *Statement) GetBalanceType() BalanceType`

GetBalanceType returns the BalanceType field if non-nil, zero value otherwise.

### GetBalanceTypeOk

`func (o *Statement) GetBalanceTypeOk() (*BalanceType, bool)`

GetBalanceTypeOk returns a tuple with the BalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceType

`func (o *Statement) SetBalanceType(v BalanceType)`

SetBalanceType sets BalanceType field to given value.

### HasBalanceType

`func (o *Statement) HasBalanceType() bool`

HasBalanceType returns a boolean if a field has been set.

### GetCurrency

`func (o *Statement) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Statement) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Statement) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Statement) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *Statement) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Statement) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Statement) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Statement) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetConciliationId

`func (o *Statement) GetConciliationId() string`

GetConciliationId returns the ConciliationId field if non-nil, zero value otherwise.

### GetConciliationIdOk

`func (o *Statement) GetConciliationIdOk() (*string, bool)`

GetConciliationIdOk returns a tuple with the ConciliationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConciliationId

`func (o *Statement) SetConciliationId(v string)`

SetConciliationId sets ConciliationId field to given value.

### HasConciliationId

`func (o *Statement) HasConciliationId() bool`

HasConciliationId returns a boolean if a field has been set.

### GetTransactionType

`func (o *Statement) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *Statement) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *Statement) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *Statement) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetProcessType

`func (o *Statement) GetProcessType() string`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *Statement) GetProcessTypeOk() (*string, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *Statement) SetProcessType(v string)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *Statement) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetMetadata

`func (o *Statement) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Statement) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Statement) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Statement) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetBalanceAfter

`func (o *Statement) GetBalanceAfter() string`

GetBalanceAfter returns the BalanceAfter field if non-nil, zero value otherwise.

### GetBalanceAfterOk

`func (o *Statement) GetBalanceAfterOk() (*string, bool)`

GetBalanceAfterOk returns a tuple with the BalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfter

`func (o *Statement) SetBalanceAfter(v string)`

SetBalanceAfter sets BalanceAfter field to given value.

### HasBalanceAfter

`func (o *Statement) HasBalanceAfter() bool`

HasBalanceAfter returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Statement) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Statement) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Statement) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Statement) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSettledAt

`func (o *Statement) GetSettledAt() string`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *Statement) GetSettledAtOk() (*string, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAt

`func (o *Statement) SetSettledAt(v string)`

SetSettledAt sets SettledAt field to given value.

### HasSettledAt

`func (o *Statement) HasSettledAt() bool`

HasSettledAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


