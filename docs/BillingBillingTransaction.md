# BillingBillingTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**EntryType** | Pointer to **string** |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**BillingBillingTransactionAmountType**](BillingBillingTransactionAmountType.md) |  | [optional] 

## Methods

### NewBillingBillingTransaction

`func NewBillingBillingTransaction() *BillingBillingTransaction`

NewBillingBillingTransaction instantiates a new BillingBillingTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingBillingTransactionWithDefaults

`func NewBillingBillingTransactionWithDefaults() *BillingBillingTransaction`

NewBillingBillingTransactionWithDefaults instantiates a new BillingBillingTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillingBillingTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingBillingTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingBillingTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingBillingTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetEntryType

`func (o *BillingBillingTransaction) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *BillingBillingTransaction) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *BillingBillingTransaction) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *BillingBillingTransaction) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### GetTransactionType

`func (o *BillingBillingTransaction) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *BillingBillingTransaction) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *BillingBillingTransaction) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *BillingBillingTransaction) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetType

`func (o *BillingBillingTransaction) GetType() BillingBillingTransactionAmountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingBillingTransaction) GetTypeOk() (*BillingBillingTransactionAmountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingBillingTransaction) SetType(v BillingBillingTransactionAmountType)`

SetType sets Type field to given value.

### HasType

`func (o *BillingBillingTransaction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


