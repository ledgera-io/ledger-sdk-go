# HelperPaginatedLedgerJournalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**LedgerJournalEntry**](LedgerJournalEntry.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewHelperPaginatedLedgerJournalEntry

`func NewHelperPaginatedLedgerJournalEntry() *HelperPaginatedLedgerJournalEntry`

NewHelperPaginatedLedgerJournalEntry instantiates a new HelperPaginatedLedgerJournalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelperPaginatedLedgerJournalEntryWithDefaults

`func NewHelperPaginatedLedgerJournalEntryWithDefaults() *HelperPaginatedLedgerJournalEntry`

NewHelperPaginatedLedgerJournalEntryWithDefaults instantiates a new HelperPaginatedLedgerJournalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HelperPaginatedLedgerJournalEntry) GetData() LedgerJournalEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HelperPaginatedLedgerJournalEntry) GetDataOk() (*LedgerJournalEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HelperPaginatedLedgerJournalEntry) SetData(v LedgerJournalEntry)`

SetData sets Data field to given value.

### HasData

`func (o *HelperPaginatedLedgerJournalEntry) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMetadata

`func (o *HelperPaginatedLedgerJournalEntry) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HelperPaginatedLedgerJournalEntry) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HelperPaginatedLedgerJournalEntry) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HelperPaginatedLedgerJournalEntry) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


