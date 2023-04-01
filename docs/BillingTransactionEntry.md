# BillingTransactionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**AmountType** | Pointer to [**BillingBillingTransactionAmountType**](BillingBillingTransactionAmountType.md) |  | [optional] 
**EntryType** | Pointer to **string** |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingTransactionEntry

`func NewBillingTransactionEntry() *BillingTransactionEntry`

NewBillingTransactionEntry instantiates a new BillingTransactionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingTransactionEntryWithDefaults

`func NewBillingTransactionEntryWithDefaults() *BillingTransactionEntry`

NewBillingTransactionEntryWithDefaults instantiates a new BillingTransactionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillingTransactionEntry) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingTransactionEntry) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingTransactionEntry) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingTransactionEntry) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountType

`func (o *BillingTransactionEntry) GetAmountType() BillingBillingTransactionAmountType`

GetAmountType returns the AmountType field if non-nil, zero value otherwise.

### GetAmountTypeOk

`func (o *BillingTransactionEntry) GetAmountTypeOk() (*BillingBillingTransactionAmountType, bool)`

GetAmountTypeOk returns a tuple with the AmountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountType

`func (o *BillingTransactionEntry) SetAmountType(v BillingBillingTransactionAmountType)`

SetAmountType sets AmountType field to given value.

### HasAmountType

`func (o *BillingTransactionEntry) HasAmountType() bool`

HasAmountType returns a boolean if a field has been set.

### GetEntryType

`func (o *BillingTransactionEntry) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *BillingTransactionEntry) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *BillingTransactionEntry) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *BillingTransactionEntry) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### GetTransactionType

`func (o *BillingTransactionEntry) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *BillingTransactionEntry) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *BillingTransactionEntry) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *BillingTransactionEntry) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


