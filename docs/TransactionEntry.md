# TransactionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionType** | **string** |  | 
**EntryType** | **string** |  | 

## Methods

### NewTransactionEntry

`func NewTransactionEntry(transactionType string, entryType string, ) *TransactionEntry`

NewTransactionEntry instantiates a new TransactionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEntryWithDefaults

`func NewTransactionEntryWithDefaults() *TransactionEntry`

NewTransactionEntryWithDefaults instantiates a new TransactionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionType

`func (o *TransactionEntry) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *TransactionEntry) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *TransactionEntry) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.


### GetEntryType

`func (o *TransactionEntry) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *TransactionEntry) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *TransactionEntry) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


