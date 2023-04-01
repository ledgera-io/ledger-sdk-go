# LedgerTransactionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** |  | 
**Currency** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ParamAccount1** | Pointer to **string** |  | [optional] 
**ParamAccount2** | Pointer to **string** |  | [optional] 
**TransactionId** | **string** |  | 
**TransactionProcess** | **string** |  | 
**TransactionType** | **string** |  | 

## Methods

### NewLedgerTransactionParams

`func NewLedgerTransactionParams(amount string, currency string, transactionId string, transactionProcess string, transactionType string, ) *LedgerTransactionParams`

NewLedgerTransactionParams instantiates a new LedgerTransactionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerTransactionParamsWithDefaults

`func NewLedgerTransactionParamsWithDefaults() *LedgerTransactionParams`

NewLedgerTransactionParamsWithDefaults instantiates a new LedgerTransactionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *LedgerTransactionParams) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LedgerTransactionParams) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LedgerTransactionParams) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *LedgerTransactionParams) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LedgerTransactionParams) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LedgerTransactionParams) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetMetadata

`func (o *LedgerTransactionParams) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LedgerTransactionParams) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LedgerTransactionParams) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LedgerTransactionParams) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetParamAccount1

`func (o *LedgerTransactionParams) GetParamAccount1() string`

GetParamAccount1 returns the ParamAccount1 field if non-nil, zero value otherwise.

### GetParamAccount1Ok

`func (o *LedgerTransactionParams) GetParamAccount1Ok() (*string, bool)`

GetParamAccount1Ok returns a tuple with the ParamAccount1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamAccount1

`func (o *LedgerTransactionParams) SetParamAccount1(v string)`

SetParamAccount1 sets ParamAccount1 field to given value.

### HasParamAccount1

`func (o *LedgerTransactionParams) HasParamAccount1() bool`

HasParamAccount1 returns a boolean if a field has been set.

### GetParamAccount2

`func (o *LedgerTransactionParams) GetParamAccount2() string`

GetParamAccount2 returns the ParamAccount2 field if non-nil, zero value otherwise.

### GetParamAccount2Ok

`func (o *LedgerTransactionParams) GetParamAccount2Ok() (*string, bool)`

GetParamAccount2Ok returns a tuple with the ParamAccount2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamAccount2

`func (o *LedgerTransactionParams) SetParamAccount2(v string)`

SetParamAccount2 sets ParamAccount2 field to given value.

### HasParamAccount2

`func (o *LedgerTransactionParams) HasParamAccount2() bool`

HasParamAccount2 returns a boolean if a field has been set.

### GetTransactionId

`func (o *LedgerTransactionParams) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *LedgerTransactionParams) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *LedgerTransactionParams) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTransactionProcess

`func (o *LedgerTransactionParams) GetTransactionProcess() string`

GetTransactionProcess returns the TransactionProcess field if non-nil, zero value otherwise.

### GetTransactionProcessOk

`func (o *LedgerTransactionParams) GetTransactionProcessOk() (*string, bool)`

GetTransactionProcessOk returns a tuple with the TransactionProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionProcess

`func (o *LedgerTransactionParams) SetTransactionProcess(v string)`

SetTransactionProcess sets TransactionProcess field to given value.


### GetTransactionType

`func (o *LedgerTransactionParams) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *LedgerTransactionParams) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *LedgerTransactionParams) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


