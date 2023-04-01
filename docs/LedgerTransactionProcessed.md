# LedgerTransactionProcessed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Entries** | Pointer to [**[]LedgerJournalEntry**](LedgerJournalEntry.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewLedgerTransactionProcessed

`func NewLedgerTransactionProcessed() *LedgerTransactionProcessed`

NewLedgerTransactionProcessed instantiates a new LedgerTransactionProcessed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerTransactionProcessedWithDefaults

`func NewLedgerTransactionProcessedWithDefaults() *LedgerTransactionProcessed`

NewLedgerTransactionProcessedWithDefaults instantiates a new LedgerTransactionProcessed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LedgerTransactionProcessed) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LedgerTransactionProcessed) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LedgerTransactionProcessed) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LedgerTransactionProcessed) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEntries

`func (o *LedgerTransactionProcessed) GetEntries() []LedgerJournalEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *LedgerTransactionProcessed) GetEntriesOk() (*[]LedgerJournalEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *LedgerTransactionProcessed) SetEntries(v []LedgerJournalEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *LedgerTransactionProcessed) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetId

`func (o *LedgerTransactionProcessed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerTransactionProcessed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedgerTransactionProcessed) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LedgerTransactionProcessed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *LedgerTransactionProcessed) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LedgerTransactionProcessed) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LedgerTransactionProcessed) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LedgerTransactionProcessed) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


