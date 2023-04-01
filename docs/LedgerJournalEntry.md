# LedgerJournalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**AccountSubType** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
**BalanceAfter** | Pointer to **float32** |  | [optional] 
**BalanceType** | Pointer to **string** |  | [optional] 
**ConciliationId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EntryId** | Pointer to **string** |  | [optional] 
**EntryLabel** | Pointer to **string** |  | [optional] 
**EntryOrder** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ProcessType** | Pointer to **string** |  | [optional] 
**SeqNum** | Pointer to **int32** |  | [optional] 
**SettledAt** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 

## Methods

### NewLedgerJournalEntry

`func NewLedgerJournalEntry() *LedgerJournalEntry`

NewLedgerJournalEntry instantiates a new LedgerJournalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerJournalEntryWithDefaults

`func NewLedgerJournalEntryWithDefaults() *LedgerJournalEntry`

NewLedgerJournalEntryWithDefaults instantiates a new LedgerJournalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *LedgerJournalEntry) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LedgerJournalEntry) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LedgerJournalEntry) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *LedgerJournalEntry) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountSubType

`func (o *LedgerJournalEntry) GetAccountSubType() string`

GetAccountSubType returns the AccountSubType field if non-nil, zero value otherwise.

### GetAccountSubTypeOk

`func (o *LedgerJournalEntry) GetAccountSubTypeOk() (*string, bool)`

GetAccountSubTypeOk returns a tuple with the AccountSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSubType

`func (o *LedgerJournalEntry) SetAccountSubType(v string)`

SetAccountSubType sets AccountSubType field to given value.

### HasAccountSubType

`func (o *LedgerJournalEntry) HasAccountSubType() bool`

HasAccountSubType returns a boolean if a field has been set.

### GetAccountType

`func (o *LedgerJournalEntry) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *LedgerJournalEntry) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *LedgerJournalEntry) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *LedgerJournalEntry) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAmount

`func (o *LedgerJournalEntry) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LedgerJournalEntry) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LedgerJournalEntry) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *LedgerJournalEntry) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalanceAfter

`func (o *LedgerJournalEntry) GetBalanceAfter() float32`

GetBalanceAfter returns the BalanceAfter field if non-nil, zero value otherwise.

### GetBalanceAfterOk

`func (o *LedgerJournalEntry) GetBalanceAfterOk() (*float32, bool)`

GetBalanceAfterOk returns a tuple with the BalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfter

`func (o *LedgerJournalEntry) SetBalanceAfter(v float32)`

SetBalanceAfter sets BalanceAfter field to given value.

### HasBalanceAfter

`func (o *LedgerJournalEntry) HasBalanceAfter() bool`

HasBalanceAfter returns a boolean if a field has been set.

### GetBalanceType

`func (o *LedgerJournalEntry) GetBalanceType() string`

GetBalanceType returns the BalanceType field if non-nil, zero value otherwise.

### GetBalanceTypeOk

`func (o *LedgerJournalEntry) GetBalanceTypeOk() (*string, bool)`

GetBalanceTypeOk returns a tuple with the BalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceType

`func (o *LedgerJournalEntry) SetBalanceType(v string)`

SetBalanceType sets BalanceType field to given value.

### HasBalanceType

`func (o *LedgerJournalEntry) HasBalanceType() bool`

HasBalanceType returns a boolean if a field has been set.

### GetConciliationId

`func (o *LedgerJournalEntry) GetConciliationId() string`

GetConciliationId returns the ConciliationId field if non-nil, zero value otherwise.

### GetConciliationIdOk

`func (o *LedgerJournalEntry) GetConciliationIdOk() (*string, bool)`

GetConciliationIdOk returns a tuple with the ConciliationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConciliationId

`func (o *LedgerJournalEntry) SetConciliationId(v string)`

SetConciliationId sets ConciliationId field to given value.

### HasConciliationId

`func (o *LedgerJournalEntry) HasConciliationId() bool`

HasConciliationId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LedgerJournalEntry) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LedgerJournalEntry) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LedgerJournalEntry) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LedgerJournalEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *LedgerJournalEntry) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LedgerJournalEntry) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LedgerJournalEntry) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *LedgerJournalEntry) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEntryId

`func (o *LedgerJournalEntry) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *LedgerJournalEntry) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *LedgerJournalEntry) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *LedgerJournalEntry) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetEntryLabel

`func (o *LedgerJournalEntry) GetEntryLabel() string`

GetEntryLabel returns the EntryLabel field if non-nil, zero value otherwise.

### GetEntryLabelOk

`func (o *LedgerJournalEntry) GetEntryLabelOk() (*string, bool)`

GetEntryLabelOk returns a tuple with the EntryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryLabel

`func (o *LedgerJournalEntry) SetEntryLabel(v string)`

SetEntryLabel sets EntryLabel field to given value.

### HasEntryLabel

`func (o *LedgerJournalEntry) HasEntryLabel() bool`

HasEntryLabel returns a boolean if a field has been set.

### GetEntryOrder

`func (o *LedgerJournalEntry) GetEntryOrder() int32`

GetEntryOrder returns the EntryOrder field if non-nil, zero value otherwise.

### GetEntryOrderOk

`func (o *LedgerJournalEntry) GetEntryOrderOk() (*int32, bool)`

GetEntryOrderOk returns a tuple with the EntryOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryOrder

`func (o *LedgerJournalEntry) SetEntryOrder(v int32)`

SetEntryOrder sets EntryOrder field to given value.

### HasEntryOrder

`func (o *LedgerJournalEntry) HasEntryOrder() bool`

HasEntryOrder returns a boolean if a field has been set.

### GetId

`func (o *LedgerJournalEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerJournalEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedgerJournalEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LedgerJournalEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *LedgerJournalEntry) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LedgerJournalEntry) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LedgerJournalEntry) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LedgerJournalEntry) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProcessType

`func (o *LedgerJournalEntry) GetProcessType() string`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *LedgerJournalEntry) GetProcessTypeOk() (*string, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *LedgerJournalEntry) SetProcessType(v string)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *LedgerJournalEntry) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetSeqNum

`func (o *LedgerJournalEntry) GetSeqNum() int32`

GetSeqNum returns the SeqNum field if non-nil, zero value otherwise.

### GetSeqNumOk

`func (o *LedgerJournalEntry) GetSeqNumOk() (*int32, bool)`

GetSeqNumOk returns a tuple with the SeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNum

`func (o *LedgerJournalEntry) SetSeqNum(v int32)`

SetSeqNum sets SeqNum field to given value.

### HasSeqNum

`func (o *LedgerJournalEntry) HasSeqNum() bool`

HasSeqNum returns a boolean if a field has been set.

### GetSettledAt

`func (o *LedgerJournalEntry) GetSettledAt() string`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *LedgerJournalEntry) GetSettledAtOk() (*string, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAt

`func (o *LedgerJournalEntry) SetSettledAt(v string)`

SetSettledAt sets SettledAt field to given value.

### HasSettledAt

`func (o *LedgerJournalEntry) HasSettledAt() bool`

HasSettledAt returns a boolean if a field has been set.

### GetTransactionId

`func (o *LedgerJournalEntry) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *LedgerJournalEntry) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *LedgerJournalEntry) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *LedgerJournalEntry) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionType

`func (o *LedgerJournalEntry) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *LedgerJournalEntry) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *LedgerJournalEntry) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *LedgerJournalEntry) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


