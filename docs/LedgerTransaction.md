# LedgerTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConciliationId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 

## Methods

### NewLedgerTransaction

`func NewLedgerTransaction() *LedgerTransaction`

NewLedgerTransaction instantiates a new LedgerTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerTransactionWithDefaults

`func NewLedgerTransactionWithDefaults() *LedgerTransaction`

NewLedgerTransactionWithDefaults instantiates a new LedgerTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConciliationId

`func (o *LedgerTransaction) GetConciliationId() string`

GetConciliationId returns the ConciliationId field if non-nil, zero value otherwise.

### GetConciliationIdOk

`func (o *LedgerTransaction) GetConciliationIdOk() (*string, bool)`

GetConciliationIdOk returns a tuple with the ConciliationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConciliationId

`func (o *LedgerTransaction) SetConciliationId(v string)`

SetConciliationId sets ConciliationId field to given value.

### HasConciliationId

`func (o *LedgerTransaction) HasConciliationId() bool`

HasConciliationId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LedgerTransaction) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LedgerTransaction) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LedgerTransaction) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LedgerTransaction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *LedgerTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedgerTransaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LedgerTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *LedgerTransaction) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LedgerTransaction) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LedgerTransaction) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LedgerTransaction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTransactionType

`func (o *LedgerTransaction) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *LedgerTransaction) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *LedgerTransaction) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *LedgerTransaction) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


